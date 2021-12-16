// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewaySls
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {

        public Client(): base()
        {
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
            string project = hostMap.Get("project");
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = credential.GetAccessKeyId();
            string accessKeySecret = credential.GetAccessKeySecret();
            string securityToken = credential.GetSecurityToken();
            if (!AlibabaCloud.TeaUtil.Common.Empty(accessKeyId))
            {
                request.Headers["x-log-signaturemethod"] = "hmac-sha1";
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-acs-security-token"] = securityToken;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "protobuf"))
                {
                    // var bodyMap = Util.assertAsMap(request.body);
                    // 缺少body的Content-MD5计算，以及protobuf处理
                    request.Headers["content-type"] = "application/x-protobuf";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json"))
                {
                    string bodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaString.StringUtil.ToUpper(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(bodyStr)));
                    request.Stream = TeaCore.BytesReadable(bodyStr);
                    request.Headers["content-type"] = "application/json";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    string str = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaString.StringUtil.ToUpper(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(str)));
                    request.Stream = TeaCore.BytesReadable(str);
                    request.Headers["content-type"] = "application/json";
                }
            }
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"accept", "application/json"},
                    {"host", GetHost(config.Network, project, config.Endpoint)},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"user-agent", request.UserAgent},
                    {"x-log-apiversion", "0.6.0"},
                    {"x-log-bodyrawsize", "0"},
                },
                request.Headers
            );
            request.Headers["authorization"] = GetAuthorization(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            Dictionary<string, string> hostMap = new Dictionary<string, string>(){};
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.HostMap))
            {
                hostMap = request.HostMap;
            }
            string project = hostMap.Get("project");
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = await credential.GetAccessKeyIdAsync();
            string accessKeySecret = await credential.GetAccessKeySecretAsync();
            string securityToken = await credential.GetSecurityTokenAsync();
            if (!AlibabaCloud.TeaUtil.Common.Empty(accessKeyId))
            {
                request.Headers["x-log-signaturemethod"] = "hmac-sha1";
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-acs-security-token"] = securityToken;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "protobuf"))
                {
                    // var bodyMap = Util.assertAsMap(request.body);
                    // 缺少body的Content-MD5计算，以及protobuf处理
                    request.Headers["content-type"] = "application/x-protobuf";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json"))
                {
                    string bodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaString.StringUtil.ToUpper(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(bodyStr)));
                    request.Stream = TeaCore.BytesReadable(bodyStr);
                    request.Headers["content-type"] = "application/json";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    string str = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaString.StringUtil.ToUpper(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(str)));
                    request.Stream = TeaCore.BytesReadable(str);
                    request.Headers["content-type"] = "application/json";
                }
            }
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"accept", "application/json"},
                    {"host", await GetHostAsync(config.Network, project, config.Endpoint)},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"user-agent", request.UserAgent},
                    {"x-log-apiversion", "0.6.0"},
                    {"x-log-bodyrawsize", "0"},
                },
                request.Headers
            );
            request.Headers["authorization"] = await GetAuthorizationAsync(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                object error = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> resMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(error);
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", resMap.Get("errorCode")},
                    {"message", resMap.Get("errorMessage")},
                    {"data", new Dictionary<string, object>
                    {
                        {"httpCode", response.StatusCode},
                        {"requestId", response.Headers.Get("x-log-requestid")},
                    }},
                });
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
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
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                object error = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> resMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(error);
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", resMap.Get("errorCode")},
                    {"message", resMap.Get("errorMessage")},
                    {"data", new Dictionary<string, object>
                    {
                        {"httpCode", response.StatusCode},
                        {"requestId", response.Headers.Get("x-log-requestid")},
                    }},
                });
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
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
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "intranet"))
                {
                    return "" + regionId + "-intranet.log.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "accelerate"))
                {
                    return "log-global.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "share"))
                {
                    if (AlibabaCloud.DarabonbaString.StringUtil.Equals(regionId, "cn-hangzhou-corp") || AlibabaCloud.DarabonbaString.StringUtil.Equals(regionId, "cn-shanghai-corp"))
                    {
                        return "" + regionId + ".sls.aliyuncs.com";
                    }
                    else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(regionId, "cn-zhangbei-corp"))
                    {
                        return "zhangbei-corp-share.log.aliyuncs.com";
                    }
                    return "" + regionId + "-share.log.aliyuncs.com";
                }
            }
            return "" + regionId + ".log.aliyuncs.com";
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
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "intranet"))
                {
                    return "" + regionId + "-intranet.log.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "accelerate"))
                {
                    return "log-global.aliyuncs.com";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "share"))
                {
                    if (AlibabaCloud.DarabonbaString.StringUtil.Equals(regionId, "cn-hangzhou-corp") || AlibabaCloud.DarabonbaString.StringUtil.Equals(regionId, "cn-shanghai-corp"))
                    {
                        return "" + regionId + ".sls.aliyuncs.com";
                    }
                    else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(regionId, "cn-zhangbei-corp"))
                    {
                        return "zhangbei-corp-share.log.aliyuncs.com";
                    }
                    return "" + regionId + "-share.log.aliyuncs.com";
                }
            }
            return "" + regionId + ".log.aliyuncs.com";
        }

        public string GetHost(string network, string project, string endpoint)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(project))
            {
                return endpoint;
            }
            return "" + project + "." + endpoint;
        }

        public async Task<string> GetHostAsync(string network, string project, string endpoint)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(project))
            {
                return endpoint;
            }
            return "" + project + "." + endpoint;
        }

        public string GetAuthorization(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            return "LOG " + ak + ":" + GetSignature(pathname, method, query, headers, secret);
        }

        public async Task<string> GetAuthorizationAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            return "LOG " + ak + ":" + await GetSignatureAsync(pathname, method, query, headers, secret);
        }

        public string GetSignature(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = pathname;
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(resource, query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = pathname;
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(resource, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public string BuildCanonicalizedResource(string pathname, Dictionary<string, string> query)
        {
            string canonicalizedResource = pathname;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryList = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryList);
                string separator = "?";

                foreach (var paramName in sortedParams) {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(paramName);
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceAsync(string pathname, Dictionary<string, string> query)
        {
            string canonicalizedResource = pathname;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryList = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryList);
                string separator = "?";

                foreach (var paramName in sortedParams) {
                    canonicalizedResource = "" + canonicalizedResource + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(paramName)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(paramName);
                    }
                    separator = "&";
                }
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
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-log-") || AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-acs-"))
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
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-log-") || AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-acs-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

    }
}
