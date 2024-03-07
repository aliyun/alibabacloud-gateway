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
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-acs-security-token"] = securityToken;
            }
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v1");
            string contentHash = "";
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
                    contentHash = MakeContentHash(bodyStr, signatureVersion);
                    request.Stream = TeaCore.BytesReadable(bodyStr);
                    request.Headers["content-type"] = "application/json";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    string str = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    contentHash = MakeContentHash(str, signatureVersion);
                    request.Stream = TeaCore.BytesReadable(str);
                    request.Headers["content-type"] = "application/json";
                }
            }
            string host = GetHost(config.Network, project, config.Endpoint);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"accept", "application/json"},
                    {"host", host},
                    {"user-agent", request.UserAgent},
                    {"x-log-apiversion", "0.6.0"},
                    {"x-log-bodyrawsize", "0"},
                },
                request.Headers
            );
            BuildRequest(context);
            // move param in path to query
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
            {
                if (AlibabaCloud.TeaUtil.Common.Empty(contentHash))
                {
                    contentHash = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
                }
                string date = GetDateISO8601();
                request.Headers["x-log-date"] = date;
                request.Headers["x-log-content-sha256"] = contentHash;
                request.Headers["authorization"] = GetAuthorizationV4(context, date, contentHash, accessKeyId, accessKeySecret);
                return ;
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(accessKeyId))
            {
                request.Headers["x-log-signaturemethod"] = "hmac-sha256";
            }
            request.Headers["date"] = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            request.Headers["content-md5"] = contentHash;
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
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-acs-security-token"] = securityToken;
            }
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v1");
            string contentHash = "";
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
                    contentHash = await MakeContentHashAsync(bodyStr, signatureVersion);
                    request.Stream = TeaCore.BytesReadable(bodyStr);
                    request.Headers["content-type"] = "application/json";
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    string str = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    contentHash = await MakeContentHashAsync(str, signatureVersion);
                    request.Stream = TeaCore.BytesReadable(str);
                    request.Headers["content-type"] = "application/json";
                }
            }
            string host = await GetHostAsync(config.Network, project, config.Endpoint);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"accept", "application/json"},
                    {"host", host},
                    {"user-agent", request.UserAgent},
                    {"x-log-apiversion", "0.6.0"},
                    {"x-log-bodyrawsize", "0"},
                },
                request.Headers
            );
            await BuildRequestAsync(context);
            // move param in path to query
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
            {
                if (AlibabaCloud.TeaUtil.Common.Empty(contentHash))
                {
                    contentHash = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
                }
                string date = await GetDateISO8601Async();
                request.Headers["x-log-date"] = date;
                request.Headers["x-log-content-sha256"] = contentHash;
                request.Headers["authorization"] = await GetAuthorizationV4Async(context, date, contentHash, accessKeyId, accessKeySecret);
                return ;
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(accessKeyId))
            {
                request.Headers["x-log-signaturemethod"] = "hmac-sha256";
            }
            request.Headers["date"] = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            request.Headers["content-md5"] = contentHash;
            request.Headers["authorization"] = await GetAuthorizationAsync(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public string MakeContentHash(string content, string signatureVersion)
        {
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
            {
                if (AlibabaCloud.TeaUtil.Common.Empty(content))
                {
                    return "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
                }
                return AlibabaCloud.DarabonbaString.StringUtil.ToLower(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(content), "SLS4-HMAC-SHA256")));
            }
            return AlibabaCloud.DarabonbaString.StringUtil.ToUpper(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(content)));
        }

        public async Task<string> MakeContentHashAsync(string content, string signatureVersion)
        {
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
            {
                if (AlibabaCloud.TeaUtil.Common.Empty(content))
                {
                    return "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
                }
                return AlibabaCloud.DarabonbaString.StringUtil.ToLower(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(content), "SLS4-HMAC-SHA256")));
            }
            return AlibabaCloud.DarabonbaString.StringUtil.ToUpper(AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(content)));
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
                        {"statusCode", response.StatusCode},
                    }},
                });
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                string bodyrawSize = response.Headers.Get("x-log-bodyrawsize");
                string compressType = response.Headers.Get("x-log-compresstype");
                Stream uncompressedData = response.Body;
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(bodyrawSize) && !AlibabaCloud.TeaUtil.Common.IsUnset(compressType))
                {
                    uncompressedData = AlibabaCloud.GatewaySls_Util.Common.ReadAndUncompressBlock(response.Body, compressType, bodyrawSize);
                }
                if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
                {
                    response.DeserializedBody = uncompressedData;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "byte"))
                {
                    byte[] byt = AlibabaCloud.TeaUtil.Common.ReadAsBytes(uncompressedData);
                    response.DeserializedBody = byt;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "string"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(uncompressedData);
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
                {
                    object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(uncompressedData);
                    // var res = Util.assertAsMap(obj);
                    response.DeserializedBody = obj;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(uncompressedData);
                }
                else
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(uncompressedData);
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
                        {"statusCode", response.StatusCode},
                    }},
                });
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
                string bodyrawSize = response.Headers.Get("x-log-bodyrawsize");
                string compressType = response.Headers.Get("x-log-compresstype");
                Stream uncompressedData = response.Body;
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(bodyrawSize) && !AlibabaCloud.TeaUtil.Common.IsUnset(compressType))
                {
                    uncompressedData = AlibabaCloud.GatewaySls_Util.Common.ReadAndUncompressBlock(response.Body, compressType, bodyrawSize);
                }
                if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "binary"))
                {
                    response.DeserializedBody = uncompressedData;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "byte"))
                {
                    byte[] byt = AlibabaCloud.TeaUtil.Common.ReadAsBytes(uncompressedData);
                    response.DeserializedBody = byt;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "string"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(uncompressedData);
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
                {
                    object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(uncompressedData);
                    // var res = Util.assertAsMap(obj);
                    response.DeserializedBody = obj;
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(uncompressedData);
                }
                else
                {
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(uncompressedData);
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
            string sign = GetSignature(pathname, method, query, headers, secret);
            return "LOG " + ak + ":" + sign;
        }

        public async Task<string> GetAuthorizationAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = await GetSignatureAsync(pathname, method, query, headers, secret);
            return "LOG " + ak + ":" + sign;
        }

        public string GetSignature(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = pathname;
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(resource, query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = pathname;
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(resource, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret));
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
                    string paramValue = query.Get(paramName);
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(paramValue))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + paramValue;
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
                    string paramValue = query.Get(paramName);
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(paramValue))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + paramValue;
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

        public void BuildRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            string resource = request.Pathname;
            if (!AlibabaCloud.TeaUtil.Common.Empty(resource))
            {
                List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(resource, "?", 2);
                resource = paths[0];
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
                {
                    List<string> params_ = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", null);

                    foreach (var sub in params_) {
                        List<string> item = AlibabaCloud.DarabonbaString.StringUtil.Split(sub, "=", null);
                        string key = item[0];
                        string value = null;
                        if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(item), 2))
                        {
                            value = item[1];
                        }
                        request.Query[key] = value;
                    }
                }
            }
            request.Pathname = resource;
        }

        public async Task BuildRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            string resource = request.Pathname;
            if (!AlibabaCloud.TeaUtil.Common.Empty(resource))
            {
                List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(resource, "?", 2);
                resource = paths[0];
                if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
                {
                    List<string> params_ = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", null);

                    foreach (var sub in params_) {
                        List<string> item = AlibabaCloud.DarabonbaString.StringUtil.Split(sub, "=", null);
                        string key = item[0];
                        string value = null;
                        if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(item), 2))
                        {
                            value = item[1];
                        }
                        request.Query[key] = value;
                    }
                }
            }
            request.Pathname = resource;
        }

        public string GetAuthorizationV4(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string date, string contentHash, string accessKeyId, string accessKeySecret)
        {
            string region = GetRegion(context);
            string headerStr = GetSignedHeaderStrV4(context.Request.Headers);
            string dateShort = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 8);
            dateShort = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateShort, "T", "", null);
            // for fix php sdk bug
            string scope = "" + dateShort + "/" + region + "/sls/aliyun_v4_request";
            byte[] signingkey = GetSigningkeyV4("SLS4-HMAC-SHA256", accessKeySecret, region, dateShort);
            string signature = GetSignatureV4(context, "SLS4-HMAC-SHA256", headerStr, date, scope, contentHash, signingkey);
            return "" + "SLS4-HMAC-SHA256" + " Credential=" + accessKeyId + "/" + scope + ",Signature=" + signature;
        }

        public async Task<string> GetAuthorizationV4Async(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string date, string contentHash, string accessKeyId, string accessKeySecret)
        {
            string region = await GetRegionAsync(context);
            string headerStr = await GetSignedHeaderStrV4Async(context.Request.Headers);
            string dateShort = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 8);
            dateShort = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateShort, "T", "", null);
            // for fix php sdk bug
            string scope = "" + dateShort + "/" + region + "/sls/aliyun_v4_request";
            byte[] signingkey = await GetSigningkeyV4Async("SLS4-HMAC-SHA256", accessKeySecret, region, dateShort);
            string signature = await GetSignatureV4Async(context, "SLS4-HMAC-SHA256", headerStr, date, scope, contentHash, signingkey);
            return "" + "SLS4-HMAC-SHA256" + " Credential=" + accessKeyId + "/" + scope + ",Signature=" + signature;
        }

        public byte[] GetSigningkeyV4(string signatureAlgorithm, string secret, string region, string date)
        {
            string sc1 = "aliyun_v4" + secret;
            byte[] sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            byte[] sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            byte[] sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("sls", sc3);
            return AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
        }

        public async Task<byte[]> GetSigningkeyV4Async(string signatureAlgorithm, string secret, string region, string date)
        {
            string sc1 = "aliyun_v4" + secret;
            byte[] sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            byte[] sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            byte[] sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("sls", sc3);
            return AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
        }

        public string GetSignatureV4(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string signatureAlgorithm, string signedHeaderStr, string date, string scope, string contentSha256, byte[] signingkey)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(request.Pathname))
            {
                canonicalURI = request.Pathname;
            }
            string resources = BuildCanonicalizedResourceV4(request.Query);
            string headers = BuildCanonicalizedHeadersV4(request.Headers, signedHeaderStr);
            string stringToHash = "" + request.Method + "\n" + canonicalURI + "\n" + resources + "\n" + headers + "\n" + signedHeaderStr + "\n" + contentSha256;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToHash), signatureAlgorithm));
            string stringToSign = "" + signatureAlgorithm + "\n" + date + "\n" + scope + "\n" + hex;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public async Task<string> GetSignatureV4Async(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string signatureAlgorithm, string signedHeaderStr, string date, string scope, string contentSha256, byte[] signingkey)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(request.Pathname))
            {
                canonicalURI = request.Pathname;
            }
            string resources = await BuildCanonicalizedResourceV4Async(request.Query);
            string headers = await BuildCanonicalizedHeadersV4Async(request.Headers, signedHeaderStr);
            string stringToHash = "" + request.Method + "\n" + canonicalURI + "\n" + resources + "\n" + headers + "\n" + signedHeaderStr + "\n" + contentSha256;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToHash), signatureAlgorithm));
            string stringToSign = "" + signatureAlgorithm + "\n" + date + "\n" + scope + "\n" + hex;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public string BuildCanonicalizedResourceV4(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + key;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(key));
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceV4Async(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + key;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(key));
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeadersV4(Dictionary<string, string> headers, string signedHeaderStr)
        {
            string canonicalizedHeaders = "";
            List<string> signedHeaders = AlibabaCloud.DarabonbaString.StringUtil.Split(signedHeaderStr, ";", null);

            foreach (var header in signedHeaders) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(header)) + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersV4Async(Dictionary<string, string> headers, string signedHeaderStr)
        {
            string canonicalizedHeaders = "";
            List<string> signedHeaders = AlibabaCloud.DarabonbaString.StringUtil.Split(signedHeaderStr, ";", null);

            foreach (var header in signedHeaders) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(header)) + "\n";
            }
            return canonicalizedHeaders;
        }

        public string GetSignedHeaderStrV4(Dictionary<string, string> headers)
        {
            List<string> headersArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeadersArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(headersArray);
            string tmp = "";
            string separator = "";

            foreach (var key in sortedHeadersArray) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-log-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "host") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type"))
                {
                    if (!AlibabaCloud.DarabonbaString.StringUtil.Contains(tmp, lowerKey))
                    {
                        tmp = "" + tmp + separator + lowerKey;
                        separator = ";";
                    }
                }
            }
            return tmp;
        }

        public async Task<string> GetSignedHeaderStrV4Async(Dictionary<string, string> headers)
        {
            List<string> headersArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeadersArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(headersArray);
            string tmp = "";
            string separator = "";

            foreach (var key in sortedHeadersArray) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-log-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "host") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type"))
                {
                    if (!AlibabaCloud.DarabonbaString.StringUtil.Contains(tmp, lowerKey))
                    {
                        tmp = "" + tmp + separator + lowerKey;
                        separator = ";";
                    }
                }
            }
            return tmp;
        }

        public string GetRegion(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(config.RegionId) && !AlibabaCloud.TeaUtil.Common.Empty(config.RegionId))
            {
                return config.RegionId;
            }
            // try parse region from endpoint
            // do not use String.subString, subString has bug in php implementation
            string region = AlibabaCloud.DarabonbaString.StringUtil.Replace(config.Endpoint, ".log.aliyuncs.com", "", null);
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, ".sls.aliyuncs.com", "", null);
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(region, config.Endpoint))
            {
                throw new TeaException(new Dictionary<string, string>
                {
                    {"code", "ClientConfigError"},
                    {"message", "The regionId configuration of sls client is missing."},
                });
            }
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, "-share", "", null);
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, "-intranet", "", null);
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, "-vpc", "", null);
            return region;
        }

        public async Task<string> GetRegionAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(config.RegionId) && !AlibabaCloud.TeaUtil.Common.Empty(config.RegionId))
            {
                return config.RegionId;
            }
            // try parse region from endpoint
            // do not use String.subString, subString has bug in php implementation
            string region = AlibabaCloud.DarabonbaString.StringUtil.Replace(config.Endpoint, ".log.aliyuncs.com", "", null);
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, ".sls.aliyuncs.com", "", null);
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(region, config.Endpoint))
            {
                throw new TeaException(new Dictionary<string, string>
                {
                    {"code", "ClientConfigError"},
                    {"message", "The regionId configuration of sls client is missing."},
                });
            }
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, "-share", "", null);
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, "-intranet", "", null);
            region = AlibabaCloud.DarabonbaString.StringUtil.Replace(region, "-vpc", "", null);
            return region;
        }

        // format: YYYYMMDDTHHMMSSZ
        public string GetDateISO8601()
        {
            string date = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            // 2024-02-04T11:31:58Z
            date = AlibabaCloud.DarabonbaString.StringUtil.Replace(date, "-", "", null);
            return AlibabaCloud.DarabonbaString.StringUtil.Replace(date, ":", "", null);
        }

        // format: YYYYMMDDTHHMMSSZ
        public async Task<string> GetDateISO8601Async()
        {
            string date = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            // 2024-02-04T11:31:58Z
            date = AlibabaCloud.DarabonbaString.StringUtil.Replace(date, "-", "", null);
            return AlibabaCloud.DarabonbaString.StringUtil.Replace(date, ":", "", null);
        }

    }
}
