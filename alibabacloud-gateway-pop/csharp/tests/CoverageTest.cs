using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using System.Threading.Tasks;
using AlibabaCloud.GatewayPop;
using AlibabaCloud.GatewaySpi.Models;
using Aliyun.Credentials.Models;
using Tea;
using Xunit;

namespace tests
{
    public class CoverageTest
    {
        private class BoomStream : MemoryStream
        {
            public override int Read(byte[] buffer, int offset, int count)
            {
                throw new IOException("read failed");
            }
        }

        private static Aliyun.Credentials.Client AccessKey()
        {
            return new Aliyun.Credentials.Client(new Config
            {
                Type = "access_key",
                AccessKeyId = "ak",
                AccessKeySecret = "secret"
            });
        }

        private static Aliyun.Credentials.Client Credential(string type, string accessKeyId, string accessKeySecret, string securityToken, string bearerToken)
        {
            return new Aliyun.Credentials.Client(new Config
            {
                Type = type,
                AccessKeyId = accessKeyId,
                AccessKeySecret = accessKeySecret,
                SecurityToken = securityToken,
                BearerToken = bearerToken
            });
        }

        private static InterceptorContext.InterceptorContextRequest Request(Aliyun.Credentials.Client credential)
        {
            return new InterceptorContext.InterceptorContextRequest
            {
                Headers = new Dictionary<string, string>(),
                Query = new Dictionary<string, string>(),
                Pathname = "/",
                ProductId = "ecs",
                Action = "DescribeRegions",
                Version = "2014-05-26",
                Protocol = "HTTPS",
                Method = "POST",
                AuthType = "AK",
                BodyType = "json",
                ReqBodyType = "json",
                Credential = credential,
                UserAgent = "test-agent"
            };
        }

        private static InterceptorContext.InterceptorContextConfiguration Config(string endpoint)
        {
            return new InterceptorContext.InterceptorContextConfiguration
            {
                RegionId = "cn-hangzhou",
                Endpoint = endpoint
            };
        }

        private static InterceptorContext Context(InterceptorContext.InterceptorContextRequest request, InterceptorContext.InterceptorContextConfiguration config)
        {
            return new InterceptorContext
            {
                Request = request,
                Configuration = config,
                Response = new InterceptorContext.InterceptorContextResponse()
            };
        }

        private static void Call(bool asyncCall, Action sync, Func<Task> asyncMethod)
        {
            if (asyncCall)
            {
                asyncMethod().GetAwaiter().GetResult();
            }
            else
            {
                sync();
            }
        }

        [Fact]
        public void EndpointAndConfiguration()
        {
            Client client = new Client();
            Assert.Equal("ecs.example.com", client.GetEndpoint("ecs", "cn-hangzhou", "regional", "", "", null, "ecs.example.com"));
            Assert.Equal("ecs.mapped.example.com", client.GetEndpoint("ecs", "cn-hangzhou", "regional", "", "", new Dictionary<string, string>
            {
                { "cn-hangzhou", "ecs.mapped.example.com" }
            }, null));
            Assert.Equal("ecs.cn-hangzhou.aliyuncs.com", client.GetEndpoint("ecs", "cn-hangzhou", "regional", "", "", null, null));
            Assert.ThrowsAny<Exception>(() => client.GetEndpoint("ecs", "", "regional", "", "", null, null));

            var built = new InterceptorContext.InterceptorContextConfiguration
            {
                RegionId = "cn-hangzhou",
                EndpointRule = "regional",
                Network = "",
                Suffix = ""
            };
            client.ModifyConfiguration(Context(Request(AccessKey()), built), new AttributeMap
            {
                Key = new Dictionary<string, string>
                {
                    { "signatureTypePrefix", "ACS4-" },
                    { "signPrefix", "aliyun_v4" },
                    { "endpointSuffix", "aliyuncs.com" }
                }
            });
            Assert.Equal("ecs.cn-hangzhou.aliyuncs.com", built.Endpoint);
            var builtAsync = new InterceptorContext.InterceptorContextConfiguration
            {
                RegionId = "cn-hangzhou",
                EndpointRule = "regional",
                Network = "",
                Suffix = ""
            };
            client.ModifyConfigurationAsync(Context(Request(AccessKey()), builtAsync), new AttributeMap
            {
                Key = new Dictionary<string, string>
                {
                    { "signatureTypePrefix", "ACS4-" },
                    { "signPrefix", "aliyun_v4" },
                    { "endpointSuffix", "aliyuncs.com" }
                }
            }).GetAwaiter().GetResult();
            Assert.Equal("ecs.cn-hangzhou.aliyuncs.com", builtAsync.Endpoint);
            var kept = Config("already.example.com");
            client.ModifyConfiguration(Context(Request(AccessKey()), kept), new AttributeMap());
            Assert.Equal("already.example.com", kept.Endpoint);
        }

        [Fact]
        public void CanonicalAndSigning()
        {
            Client client = new Client();
            Assert.Equal("", client.BuildCanonicalizedResource(null));
            string resource = client.BuildCanonicalizedResource(new Dictionary<string, string>
            {
                { "b", "x y" },
                { "a", "m*n" },
                { "empty", "" }
            });
            Assert.Contains("a=m%2An", resource);
            Assert.Contains("b=x%20y", resource);
            string canonical = client.BuildCanonicalizedHeaders(new Dictionary<string, string>
            {
                { "Host", "  example.com  " },
                { "X-Acs-Action", "One" },
                { "x-acs-action", "Two" }
            });
            Assert.Contains("host:example.com", canonical);
            Assert.Equal("fallback", client.DefaultAny(null, "fallback"));
            Assert.Equal("kept", client.DefaultAny("kept", "fallback"));
            byte[] signingkey = client.GetSigningkey("ACS4-HMAC-SHA256", "secret", "ecs", "cn-hangzhou", "20240101");
            Assert.NotEmpty(signingkey);
            byte[] unknown = client.GetSigningkey("OTHER", "secret", "ecs", "cn-hangzhou", "20240101");
            Assert.Empty(unknown);
            var headers = new Dictionary<string, string>
            {
                { "host", "example.com" },
                { "x-acs-date", "2024-01-01T00:00:00Z" }
            };
            string signature = client.GetSignature("", "GET", new Dictionary<string, string> { { "q", "a b" } }, headers, "ACS4-HMAC-SHA256", "payload", signingkey);
            Assert.Equal(64, signature.Length);
            string authorization = client.GetAuthorization("/", "GET", null, headers, "ACS4-HMAC-SHA256", "payload", "ak", signingkey, "ecs", "cn-hangzhou", "20240101");
            Assert.Contains("ACS4-HMAC-SHA256 Credential=ak/20240101/cn-hangzhou/ecs/aliyun_v4_request,SignedHeaders=", authorization);
        }

        [Theory]
        [InlineData(false)]
        [InlineData(true)]
        public void ModifyRequestBodiesAndAuth(bool asyncCall)
        {
            Client client = new Client();
            var json = Request(AccessKey());
            json.Body = new Dictionary<string, object> { { "name", "a b" } };
            json.Query["InstanceId"] = "i-1";
            Call(asyncCall,
                () => client.ModifyRequest(Context(json, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(json, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.Equal("application/json; charset=utf-8", json.Headers["content-type"]);
            Assert.Contains("ACS4-HMAC-SHA256 Credential=ak/", json.Headers["Authorization"]);
            Assert.False(string.IsNullOrEmpty(json.Headers["x-acs-content-sha256"]));

            var bytes = Request(AccessKey());
            bytes.ReqBodyType = "byte";
            bytes.Body = Encoding.UTF8.GetBytes("hello");
            bytes.Action = "";
            Call(asyncCall,
                () => client.ModifyRequest(Context(bytes, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(bytes, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.False(bytes.Headers.ContainsKey("x-acs-action"));

            var form = Request(AccessKey());
            form.ReqBodyType = "form";
            form.Body = new Dictionary<string, object> { { "b", "x y" }, { "a", new List<string> { "m", "n" } } };
            Call(asyncCall,
                () => client.ModifyRequest(Context(form, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(form, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.Equal("application/x-www-form-urlencoded", form.Headers["content-type"]);

            var stream = Request(AccessKey());
            stream.Stream = new MemoryStream(Encoding.UTF8.GetBytes("stream-body"));
            Call(asyncCall,
                () => client.ModifyRequest(Context(stream, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(stream, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.Equal("application/octet-stream", stream.Headers["content-type"]);

            var bearer = Request(Credential("bearer", null, null, null, "bearer-token"));
            Call(asyncCall,
                () => client.ModifyRequest(Context(bearer, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(bearer, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.Equal("Bearer bearer-token", bearer.Headers["Authorization"]);

            var sts = Request(Credential("sts", "sts-ak", "secret", "token", null));
            Call(asyncCall,
                () => client.ModifyRequest(Context(sts, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(sts, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.Equal("token", sts.Headers["x-acs-security-token"]);

            var anonymous = Request(null);
            anonymous.AuthType = "Anonymous";
            Call(asyncCall,
                () => client.ModifyRequest(Context(anonymous, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()),
                () => client.ModifyRequestAsync(Context(anonymous, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.False(anonymous.Headers.ContainsKey("Authorization"));

            var missing = Assert.Throws<TeaException>(() => client.ModifyRequest(Context(Request(null), Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            Assert.Equal("ParameterMissing", missing.Code);
            var missingAsync = Assert.Throws<TeaException>(() => client.ModifyRequestAsync(Context(Request(null), Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()).GetAwaiter().GetResult());
            Assert.Equal("ParameterMissing", missingAsync.Code);

            var badBytes = Request(AccessKey());
            badBytes.ReqBodyType = "byte";
            badBytes.Body = "not-bytes";
            Assert.ThrowsAny<Exception>(() => client.ModifyRequest(Context(badBytes, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
            var badStream = Request(AccessKey());
            badStream.Stream = new BoomStream();
            Assert.ThrowsAny<Exception>(() => client.ModifyRequest(Context(badStream, Config("ecs.cn-hangzhou.aliyuncs.com")), new AttributeMap()));
        }

        [Theory]
        [InlineData(false)]
        [InlineData(true)]
        public void ModifyResponseBodies(bool asyncCall)
        {
            Client client = new Client();
            var error = Assert.Throws<TeaException>(() => client.ModifyResponse(Response(400, "application/json", "{\"Code\":\"Invalid\",\"Message\":\"no\",\"RequestId\":\"req-1\",\"Description\":\"d\"}", "json"), new AttributeMap()));
            Assert.Equal("Invalid", error.Code);
            Assert.Contains("request id: req-1", error.Message);
            var headerError = Response(400, "application/json", "{\"Code\":\"Denied\",\"Message\":\"no\",\"RequestId\":\"body-id\"}", "json");
            headerError.Response.Headers["x-acs-request-id"] = "header-id";
            var denied = Assert.Throws<TeaException>(() => client.ModifyResponse(headerError, new AttributeMap()));
            Assert.Contains("request id: header-id", denied.Message);

            var xml = Assert.Throws<TeaException>(() => client.ModifyResponse(Response(500, "text/xml", "<Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error>", "xml"), new AttributeMap()));
            Assert.Equal("XmlCode", xml.Code);
            var asyncError = Assert.Throws<TeaException>(() => client.ModifyResponseAsync(Response(400, "application/json", "{\"Code\":\"Invalid\",\"Message\":\"no\",\"RequestId\":\"req-1\",\"Description\":\"d\"}", "json"), new AttributeMap()).GetAwaiter().GetResult());
            Assert.Equal("Invalid", asyncError.Code);
            var asyncXml = Assert.Throws<TeaException>(() => client.ModifyResponseAsync(Response(500, "text/xml", "<Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error>", "xml"), new AttributeMap()).GetAwaiter().GetResult());
            Assert.Equal("XmlCode", asyncXml.Code);

            Call(asyncCall,
                () => client.ModifyResponse(Response(204, null, "", "string"), new AttributeMap()),
                () => client.ModifyResponseAsync(Response(204, null, "", "string"), new AttributeMap()));
            var binary = Response(200, null, "raw", "binary");
            Call(asyncCall,
                () => client.ModifyResponse(binary, new AttributeMap()),
                () => client.ModifyResponseAsync(binary, new AttributeMap()));
            Assert.NotNull(binary.Response.DeserializedBody);
            var text = Response(200, null, "hello", "string");
            Call(asyncCall,
                () => client.ModifyResponse(text, new AttributeMap()),
                () => client.ModifyResponseAsync(text, new AttributeMap()));
            Assert.Equal("hello", text.Response.DeserializedBody);
            var parsed = Response(200, null, "{\"ok\":true}", "json");
            Call(asyncCall,
                () => client.ModifyResponse(parsed, new AttributeMap()),
                () => client.ModifyResponseAsync(parsed, new AttributeMap()));
            var other = Response(200, null, "plain", "other");
            Call(asyncCall,
                () => client.ModifyResponse(other, new AttributeMap()),
                () => client.ModifyResponseAsync(other, new AttributeMap()));
            Assert.Equal("plain", other.Response.DeserializedBody);
            var bytes = Response(200, null, "raw", "byte");
            Call(asyncCall,
                () => client.ModifyResponse(bytes, new AttributeMap()),
                () => client.ModifyResponseAsync(bytes, new AttributeMap()));
            Assert.NotNull(bytes.Response.DeserializedBody);
            var array = Response(200, null, "[1,2]", "array");
            Call(asyncCall,
                () => client.ModifyResponse(array, new AttributeMap()),
                () => client.ModifyResponseAsync(array, new AttributeMap()));
            Assert.NotNull(array.Response.DeserializedBody);
        }

        private static InterceptorContext Response(int status, string contentType, string body, string bodyType)
        {
            var context = Context(Request(AccessKey()), Config("ecs.cn-hangzhou.aliyuncs.com"));
            context.Request.BodyType = bodyType;
            context.Response.StatusCode = status;
            context.Response.Headers = new Dictionary<string, string>();
            if (contentType != null)
            {
                context.Response.Headers["content-type"] = contentType;
            }
            context.Response.Body = new MemoryStream(Encoding.UTF8.GetBytes(body));
            return context;
        }
    }
}
