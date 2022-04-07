// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewayOss
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {
        protected List<string> _default_signed_params;
        protected List<string> _except_signed_params;

        public Client(): base()
        {
            this._default_signed_params = new List<string>
            {
                "location",
                "cors",
                "objectMeta",
                "uploadId",
                "partNumber",
                "security-token",
                "position",
                "img",
                "style",
                "styleName",
                "replication",
                "replicationProgress",
                "replicationLocation",
                "cname",
                "qos",
                "startTime",
                "endTime",
                "symlink",
                "x-oss-process",
                "response-content-type",
                "response-content-language",
                "response-expires",
                "response-cache-control",
                "response-content-disposition",
                "response-content-encoding",
                "udf",
                "udfName",
                "udfImage",
                "udfId",
                "udfImageDesc",
                "udfApplication",
                "udfApplicationLog",
                "restore",
                "callback",
                "callback-var",
                "policy",
                "encryption",
                "versions",
                "versioning",
                "versionId"
            };
            this._except_signed_params = new List<string>
            {
                "list-type",
                "regions"
            };
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = GetEndpoint(config.RegionId, config.Network, config.Endpoint);
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = await GetEndpointAsync(config.RegionId, config.Network, config.Endpoint);
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            Dictionary<string, string> hostMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.HostMap))
            {
                hostMap = request.HostMap;
            }
            string bucketName = hostMap.Get("bucket");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(bucketName))
            {
                bucketName = "";
            }
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = credential.GetAccessKeyId();
            string accessKeySecret = credential.GetAccessKeySecret();
            string securityToken = credential.GetSecurityToken();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-oss-security-token"] = securityToken;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "xml"))
                {
                    Dictionary<string, object> reqBodyMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    request.Stream = TeaCore.BytesReadable(AlibabaCloud.TeaXML.Client.ToXML(reqBodyMap));
                    request.Headers["content-type"] = "application/xml";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json"))
                {
                    string reqBodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Stream = TeaCore.BytesReadable(reqBodyStr);
                    request.Headers["content-type"] = "application/json; charset=utf-8";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    Dictionary<string, object> reqBodyForm = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    request.Stream = TeaCore.BytesReadable(AlibabaCloud.OpenApiUtil.Client.ToForm(reqBodyForm));
                    request.Headers["content-type"] = "application/x-www-form-urlencoded";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "binary"))
                {
                    request.Stream = AlibabaCloud.OSSUtil.Common.Inject(request.Stream, attributeMap.Key);
                    request.Headers["content-type"] = "application/octet-stream";
                }
            }
            string host = GetHost(config.EndpointType, bucketName, config.Endpoint);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", host},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"user-agent", request.UserAgent},
                },
                request.Headers
            );
            request.Headers["authorization"] = GetAuthorization(request.SignatureVersion, bucketName, request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            Dictionary<string, string> hostMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.HostMap))
            {
                hostMap = request.HostMap;
            }
            string bucketName = hostMap.Get("bucket");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(bucketName))
            {
                bucketName = "";
            }
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = await credential.GetAccessKeyIdAsync();
            string accessKeySecret = await credential.GetAccessKeySecretAsync();
            string securityToken = await credential.GetSecurityTokenAsync();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-oss-security-token"] = securityToken;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "xml"))
                {
                    Dictionary<string, object> reqBodyMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    request.Stream = TeaCore.BytesReadable(AlibabaCloud.TeaXML.Client.ToXML(reqBodyMap));
                    request.Headers["content-type"] = "application/xml";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json"))
                {
                    string reqBodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Stream = TeaCore.BytesReadable(reqBodyStr);
                    request.Headers["content-type"] = "application/json; charset=utf-8";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    Dictionary<string, object> reqBodyForm = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    request.Stream = TeaCore.BytesReadable(AlibabaCloud.OpenApiUtil.Client.ToForm(reqBodyForm));
                    request.Headers["content-type"] = "application/x-www-form-urlencoded";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "binary"))
                {
                    request.Stream = AlibabaCloud.OSSUtil.Common.Inject(request.Stream, attributeMap.Key);
                    request.Headers["content-type"] = "application/octet-stream";
                }
            }
            string host = await GetHostAsync(config.EndpointType, bucketName, config.Endpoint);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", host},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"user-agent", request.UserAgent},
                },
                request.Headers
            );
            request.Headers["authorization"] = await GetAuthorizationAsync(request.SignatureVersion, bucketName, request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            string bodyStr = null;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                Dictionary<string, object> respMap = AlibabaCloud.TeaXML.Client.ParseXml(bodyStr, null);
                List<string> errors = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(respMap);
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(errors), 1))
                {
                    string error = errors[0];
                    respMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(respMap.Get(error));
                }
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", respMap.Get("Code")},
                    {"message", respMap.Get("Message")},
                    {"data", new Dictionary<string, object>
                    {
                        {"httpCode", response.StatusCode},
                        {"requestId", respMap.Get("RequestId")},
                        {"hostId", respMap.Get("HostId")},
                    }},
                });
            }
            Dictionary<string, string> ctx = attributeMap.Key;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx))
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("crc")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-oss-hash-crc64ecma")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("crc"), response.Headers.Get("x-oss-hash-crc64ecma")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "CrcNotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientCrc", ctx.Get("crc")},
                            {"serverCrc", response.Headers.Get("x-oss-hash-crc64ecma")},
                        }},
                    });
                }
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("md5")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("content-md5")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("md5"), response.Headers.Get("content-md5")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "MD5NotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientMD5", ctx.Get("md5")},
                            {"serverMD5", response.Headers.Get("content-md5")},
                        }},
                    });
                }
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.BodyType, "xml"))
                {
                    bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                    Dictionary<string, object> result = AlibabaCloud.TeaXML.Client.ParseXml(bodyStr, null);
                    List<string> list = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(result);
                    if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(list), 1))
                    {
                        string tmp = list[0];
                        try
                        {
                            response.DeserializedBody = AlibabaCloud.TeaUtil.Common.AssertAsMap(result.Get(tmp));
                        }
                        catch (TeaException error)
                        {
                            response.DeserializedBody = result;
                        }
                        catch (Exception _error)
                        {
                            TeaException error = new TeaException(new Dictionary<string, object>
                            {
                                { "message", _error.Message }
                            });
                            response.DeserializedBody = result;
                        }
                    }
                    else
                    {
                        response.DeserializedBody = result;
                    }
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
                {
                    response.DeserializedBody = response.Body;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "byte"))
                {
                    byte[] byt = AlibabaCloud.TeaUtil.Common.ReadAsBytes(response.Body);
                    response.DeserializedBody = byt;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "string"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
                {
                    object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> res = AlibabaCloud.TeaUtil.Common.AssertAsMap(obj);
                    response.DeserializedBody = res;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                }
                else
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
            }
        }

        public async Task ModifyResponseAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            string bodyStr = null;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                Dictionary<string, object> respMap = AlibabaCloud.TeaXML.Client.ParseXml(bodyStr, null);
                List<string> errors = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(respMap);
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(errors), 1))
                {
                    string error = errors[0];
                    respMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(respMap.Get(error));
                }
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", respMap.Get("Code")},
                    {"message", respMap.Get("Message")},
                    {"data", new Dictionary<string, object>
                    {
                        {"httpCode", response.StatusCode},
                        {"requestId", respMap.Get("RequestId")},
                        {"hostId", respMap.Get("HostId")},
                    }},
                });
            }
            Dictionary<string, string> ctx = attributeMap.Key;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx))
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("crc")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-oss-hash-crc64ecma")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("crc"), response.Headers.Get("x-oss-hash-crc64ecma")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "CrcNotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientCrc", ctx.Get("crc")},
                            {"serverCrc", response.Headers.Get("x-oss-hash-crc64ecma")},
                        }},
                    });
                }
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(ctx.Get("md5")) && !AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("content-md5")) && !AlibabaCloud.DarabonbaString.StringUtil.Equals(ctx.Get("md5"), response.Headers.Get("content-md5")))
                {
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "MD5NotMatched"},
                        {"data", new Dictionary<string, string>
                        {
                            {"clientMD5", ctx.Get("md5")},
                            {"serverMD5", response.Headers.Get("content-md5")},
                        }},
                    });
                }
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.BodyType, "xml"))
                {
                    bodyStr = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                    Dictionary<string, object> result = AlibabaCloud.TeaXML.Client.ParseXml(bodyStr, null);
                    List<string> list = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(result);
                    if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(list), 1))
                    {
                        string tmp = list[0];
                        try
                        {
                            response.DeserializedBody = AlibabaCloud.TeaUtil.Common.AssertAsMap(result.Get(tmp));
                        }
                        catch (TeaException error)
                        {
                            response.DeserializedBody = result;
                        }
                        catch (Exception _error)
                        {
                            TeaException error = new TeaException(new Dictionary<string, object>
                            {
                                { "message", _error.Message }
                            });
                            response.DeserializedBody = result;
                        }
                    }
                    else
                    {
                        response.DeserializedBody = result;
                    }
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
                {
                    response.DeserializedBody = response.Body;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "byte"))
                {
                    byte[] byt = AlibabaCloud.TeaUtil.Common.ReadAsBytes(response.Body);
                    response.DeserializedBody = byt;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "string"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
                {
                    object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> res = AlibabaCloud.TeaUtil.Common.AssertAsMap(obj);
                    response.DeserializedBody = res;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                }
                else
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                }
            }
        }

        public string GetEndpoint(string regionId, string network, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = "cn-hangzhou";
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(network))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "internal"))
                {
                    return "oss-" + regionId + "-internal.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "ipv6"))
                {
                    return "" + regionId + "oss.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "accelerate"))
                {
                    return "oss-" + network + ".aliyuncs.com";
                }
            }
            return "oss-" + regionId + ".aliyuncs.com";
        }

        public async Task<string> GetEndpointAsync(string regionId, string network, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = "cn-hangzhou";
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(network))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "internal"))
                {
                    return "oss-" + regionId + "-internal.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "ipv6"))
                {
                    return "" + regionId + "oss.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Contains(network, "accelerate"))
                {
                    return "oss-" + network + ".aliyuncs.com";
                }
            }
            return "oss-" + regionId + ".aliyuncs.com";
        }

        public string GetHost(string endpointType, string bucketName, string endpoint)
        {
            if (AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                return endpoint;
            }
            string host = "" + bucketName + "." + endpoint;
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpointType))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "ip"))
                {
                    host = "" + endpoint + "/" + bucketName;
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "cname"))
                {
                    host = endpoint;
                }
            }
            return host;
        }

        public async Task<string> GetHostAsync(string endpointType, string bucketName, string endpoint)
        {
            if (AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                return endpoint;
            }
            string host = "" + bucketName + "." + endpoint;
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpointType))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "ip"))
                {
                    host = "" + endpoint + "/" + bucketName;
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(endpointType, "cname"))
                {
                    host = endpoint;
                }
            }
            return host;
        }

        public string GetAuthorization(string signatureVersion, string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = "";
            if (AlibabaCloud.TeaUtil.Common.IsUnset(signatureVersion) || AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v1"))
            {
                sign = GetSignatureV1(bucketName, pathname, method, query, headers, secret);
                return "OSS " + ak + ":" + sign;
            }
            else
            {
                sign = GetSignatureV2(bucketName, pathname, method, query, headers, secret);
                return "OSS2 AccessKeyId:" + ak + ",Signature:" + sign;
            }
        }

        public async Task<string> GetAuthorizationAsync(string signatureVersion, string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = "";
            if (AlibabaCloud.TeaUtil.Common.IsUnset(signatureVersion) || AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v1"))
            {
                sign = await GetSignatureV1Async(bucketName, pathname, method, query, headers, secret);
                return "OSS " + ak + ":" + sign;
            }
            else
            {
                sign = await GetSignatureV2Async(bucketName, pathname, method, query, headers, secret);
                return "OSS2 AccessKeyId:" + ak + ",Signature:" + sign;
            }
        }

        public string GetSignatureV1(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = "";
            string stringToSign = "";
            if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                resource = "/" + bucketName;
            }
            resource = "" + resource + pathname;
            string canonicalizedResource = BuildCanonicalizedResource(resource, query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureV1Async(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = "";
            string stringToSign = "";
            if (!AlibabaCloud.TeaUtil.Common.Empty(bucketName))
            {
                resource = "/" + bucketName;
            }
            resource = "" + resource + pathname;
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(resource, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public string BuildCanonicalizedResource(string pathname, Dictionary<string, string> query)
        {
            Dictionary<string, string> subResourcesMap = new Dictionary<string, string>(){};
            string canonicalizedResource = pathname;
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(pathname, "?", 2);
                canonicalizedResource = paths[0];
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
                {
                    List<string> subResources = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", 0);

                    foreach (var sub in subResources) {
                        bool? hasExcepts = false;

                        foreach (var excepts in _except_signed_params) {
                            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(sub, excepts))
                            {
                                hasExcepts = true;
                            }
                        }
                        if (!hasExcepts)
                        {
                            List<string> item = AlibabaCloud.DarabonbaString.StringUtil.Split(sub, "=", 0);
                            string key = item[0];
                            string value = null;
                            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(item), 2))
                            {
                                value = item[1];
                            }
                            subResourcesMap[key] = value;
                        }
                    }
                }
            }
            List<string> subResourcesArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(subResourcesMap);
            List<string> newQueryList = subResourcesArray;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryList = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                newQueryList = AlibabaCloud.DarabonbaArray.ArrayUtil.Concat(subResourcesArray, queryList);
            }
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(newQueryList);
            string separator = "?";

            foreach (var paramName in sortedParams) {
                if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(_default_signed_params, paramName))
                {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query) && !AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(paramName);
                    }
                    else if (!AlibabaCloud.TeaUtil.Common.IsUnset(subResourcesMap.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + subResourcesMap.Get(paramName);
                    }
                }
                else if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(subResourcesArray, paramName))
                {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(subResourcesMap.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + subResourcesMap.Get(paramName);
                    }
                }
                separator = "&";
            }
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceAsync(string pathname, Dictionary<string, string> query)
        {
            Dictionary<string, string> subResourcesMap = new Dictionary<string, string>(){};
            string canonicalizedResource = pathname;
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(pathname, "?", 2);
                canonicalizedResource = paths[0];
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
                {
                    List<string> subResources = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", 0);

                    foreach (var sub in subResources) {
                        bool? hasExcepts = false;

                        foreach (var excepts in _except_signed_params) {
                            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(sub, excepts))
                            {
                                hasExcepts = true;
                            }
                        }
                        if (!hasExcepts)
                        {
                            List<string> item = AlibabaCloud.DarabonbaString.StringUtil.Split(sub, "=", 0);
                            string key = item[0];
                            string value = null;
                            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(item), 2))
                            {
                                value = item[1];
                            }
                            subResourcesMap[key] = value;
                        }
                    }
                }
            }
            List<string> subResourcesArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(subResourcesMap);
            List<string> newQueryList = subResourcesArray;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryList = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                newQueryList = AlibabaCloud.DarabonbaArray.ArrayUtil.Concat(subResourcesArray, queryList);
            }
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(newQueryList);
            string separator = "?";

            foreach (var paramName in sortedParams) {
                if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(_default_signed_params, paramName))
                {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query) && !AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(paramName);
                    }
                    else if (!AlibabaCloud.TeaUtil.Common.IsUnset(subResourcesMap.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + subResourcesMap.Get(paramName);
                    }
                }
                else if (AlibabaCloud.DarabonbaArray.ArrayUtil.Contains(subResourcesArray, paramName))
                {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(subResourcesMap.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + subResourcesMap.Get(paramName);
                    }
                }
                separator = "&";
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeaders(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            string contentType = headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string contentMd5 = headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            canonicalizedHeaders = "" + canonicalizedHeaders + contentMd5 + "\n" + contentType + "\n" + headers.Get("date") + "\n";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-oss-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersAsync(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            string contentType = headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string contentMd5 = headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            canonicalizedHeaders = "" + canonicalizedHeaders + contentMd5 + "\n" + contentType + "\n" + headers.Get("date") + "\n";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-oss-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public string GetSignatureV2(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            return "";
        }

        public async Task<string> GetSignatureV2Async(string bucketName, string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            return "";
        }

    }
}
