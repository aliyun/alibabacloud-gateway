// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewayMns
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {
        protected string _signPrefix;
        protected string _signSuffix;
        protected string _authPrefix;

        public Client(): base()
        {
            this._signPrefix = "aliyun_v4";
            this._signSuffix = "aliyun_v4_request";
            this._authPrefix = "MNS4-HMAC-SHA256";
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = GetEndpoint(config.RegionId, config.Endpoint);
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = await GetEndpointAsync(config.RegionId, config.Endpoint);
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v2");
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "xml"))
                {
                    Dictionary<string, object> reqBodyMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    string xmlStr = AlibabaCloud.TeaXML.Client.ToXML(reqBodyMap);
                    request.Stream = TeaCore.BytesReadable(xmlStr);
                    request.Headers["content-type"] = "text/xml;charset=UTF-8";
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(xmlStr));
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json") || AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    string bodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Stream = TeaCore.BytesReadable(bodyStr);
                    request.Headers["content-type"] = "application/json";
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(bodyStr));
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "byte") || AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "binary"))
                {
                    byte[] bodyBytes = AlibabaCloud.TeaUtil.Common.AssertAsBytes(request.Body);
                    request.Stream = TeaCore.BytesReadable(bodyBytes);
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5SignForBytes(bodyBytes));
                }
            }
            BuildRequest(context);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-mns-version", request.Version},
                    {"user-agent", request.UserAgent},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            if (!AlibabaCloud.TeaUtil.Common.EqualString(request.AuthType, "Anonymous"))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                if (AlibabaCloud.TeaUtil.Common.IsUnset(credential))
                {
                    throw new TeaException(new Dictionary<string, string>
                    {
                        {"code", "ParameterMissing"},
                        {"message", "'config.credential' can not be unset"},
                    });
                }
                Aliyun.Credentials.Models.CredentialModel credentialModel = credential.GetCredential();
                string authType = credentialModel.Type;
                if (AlibabaCloud.TeaUtil.Common.EqualString(authType, "bearer"))
                {
                    string bearerToken = credentialModel.BearerToken;
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["x-acs-signature-type"] = "BEARERTOKEN";
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else
                {
                    string accessKeyId = credentialModel.AccessKeyId;
                    string accessKeySecret = credentialModel.AccessKeySecret;
                    string securityToken = credentialModel.SecurityToken;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                    {
                        request.Headers["security-token"] = securityToken;
                    }
                    request.Headers["date"] = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
                    if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
                    {
                        string date = GetDateISO8601();
                        request.Headers["authorization"] = GetAuthorizationV4(context, date, accessKeyId, accessKeySecret);
                    }
                    else
                    {
                        request.Headers["authorization"] = GetAuthorizationV2(request.Pathname, request.Method, request.Headers, accessKeyId, accessKeySecret);
                    }
                }
            }
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v2");
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "xml"))
                {
                    Dictionary<string, object> reqBodyMap = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                    string xmlStr = AlibabaCloud.TeaXML.Client.ToXML(reqBodyMap);
                    request.Stream = TeaCore.BytesReadable(xmlStr);
                    request.Headers["content-type"] = "text/xml;charset=UTF-8";
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(xmlStr));
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "json") || AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "formData"))
                {
                    string bodyStr = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                    request.Stream = TeaCore.BytesReadable(bodyStr);
                    request.Headers["content-type"] = "application/json";
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(bodyStr));
                }
                else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "byte") || AlibabaCloud.DarabonbaString.StringUtil.Equals(request.ReqBodyType, "binary"))
                {
                    byte[] bodyBytes = AlibabaCloud.TeaUtil.Common.AssertAsBytes(request.Body);
                    request.Stream = TeaCore.BytesReadable(bodyBytes);
                    request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5SignForBytes(bodyBytes));
                }
            }
            await BuildRequestAsync(context);
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-mns-version", request.Version},
                    {"user-agent", request.UserAgent},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            if (!AlibabaCloud.TeaUtil.Common.EqualString(request.AuthType, "Anonymous"))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                if (AlibabaCloud.TeaUtil.Common.IsUnset(credential))
                {
                    throw new TeaException(new Dictionary<string, string>
                    {
                        {"code", "ParameterMissing"},
                        {"message", "'config.credential' can not be unset"},
                    });
                }
                Aliyun.Credentials.Models.CredentialModel credentialModel = await credential.GetCredentialAsync();
                string authType = credentialModel.Type;
                if (AlibabaCloud.TeaUtil.Common.EqualString(authType, "bearer"))
                {
                    string bearerToken = credentialModel.BearerToken;
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["x-acs-signature-type"] = "BEARERTOKEN";
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else
                {
                    string accessKeyId = credentialModel.AccessKeyId;
                    string accessKeySecret = credentialModel.AccessKeySecret;
                    string securityToken = credentialModel.SecurityToken;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                    {
                        request.Headers["security-token"] = securityToken;
                    }
                    request.Headers["date"] = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
                    if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
                    {
                        string date = await GetDateISO8601Async();
                        request.Headers["authorization"] = await GetAuthorizationV4Async(context, date, accessKeyId, accessKeySecret);
                    }
                    else
                    {
                        request.Headers["authorization"] = await GetAuthorizationV2Async(request.Pathname, request.Method, request.Headers, accessKeyId, accessKeySecret);
                    }
                }
            }
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                Dictionary<string, object> err = new Dictionary<string, object>(){};
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("content-type")) && AlibabaCloud.DarabonbaString.StringUtil.Contains(response.Headers.Get("content-type"), "text/xml"))
                {
                    string _str = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                    Dictionary<string, object> respMap = AlibabaCloud.TeaXML.Client.ParseXml(_str, null);
                    err = AlibabaCloud.TeaUtil.Common.AssertAsMap(respMap.Get("Error"));
                }
                else
                {
                    object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                }
                object requestId = DefaultAny(err.Get("RequestId"), err.Get("requestId"));
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-mns-request-id")))
                {
                    requestId = response.Headers.Get("x-mns-request-id");
                }
                err["statusCode"] = response.StatusCode;
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "" + DefaultAny(err.Get("Code"), err.Get("code"))},
                    {"message", "code: " + response.StatusCode + ", " + DefaultAny(err.Get("Message"), err.Get("message")) + " request id: " + requestId},
                    {"data", err},
                    {"description", "" + DefaultAny(err.Get("Description"), err.Get("description"))},
                    {"accessDeniedDetail", DefaultAny(err.Get("AccessDeniedDetail"), err.Get("accessDeniedDetail"))},
                });
            }
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(response.StatusCode, 204))
            {
                AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
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
                string str = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                response.DeserializedBody = str;
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
            {
                object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> res = AlibabaCloud.TeaUtil.Common.AssertAsMap(obj);
                response.DeserializedBody = res;
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
            {
                object arr = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                response.DeserializedBody = arr;
            }
            else
            {
                response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
            }
        }

        public async Task ModifyResponseAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                Dictionary<string, object> err = new Dictionary<string, object>(){};
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("content-type")) && AlibabaCloud.DarabonbaString.StringUtil.Contains(response.Headers.Get("content-type"), "text/xml"))
                {
                    string _str = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                    Dictionary<string, object> respMap = AlibabaCloud.TeaXML.Client.ParseXml(_str, null);
                    err = AlibabaCloud.TeaUtil.Common.AssertAsMap(respMap.Get("Error"));
                }
                else
                {
                    object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                }
                object requestId = DefaultAny(err.Get("RequestId"), err.Get("requestId"));
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-mns-request-id")))
                {
                    requestId = response.Headers.Get("x-mns-request-id");
                }
                err["statusCode"] = response.StatusCode;
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "" + DefaultAny(err.Get("Code"), err.Get("code"))},
                    {"message", "code: " + response.StatusCode + ", " + DefaultAny(err.Get("Message"), err.Get("message")) + " request id: " + requestId},
                    {"data", err},
                    {"description", "" + DefaultAny(err.Get("Description"), err.Get("description"))},
                    {"accessDeniedDetail", DefaultAny(err.Get("AccessDeniedDetail"), err.Get("accessDeniedDetail"))},
                });
            }
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(response.StatusCode, 204))
            {
                AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
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
                string str = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
                response.DeserializedBody = str;
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "json"))
            {
                object obj = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> res = AlibabaCloud.TeaUtil.Common.AssertAsMap(obj);
                response.DeserializedBody = res;
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(request.BodyType, "array"))
            {
                object arr = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                response.DeserializedBody = arr;
            }
            else
            {
                response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsString(response.Body);
            }
        }

        public string GetEndpoint(string regionId, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = "cn-hangzhou";
            }
            return "" + regionId + ".mns.aliyuncs.com";
        }

        public async Task<string> GetEndpointAsync(string regionId, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (AlibabaCloud.TeaUtil.Common.Empty(regionId))
            {
                regionId = "cn-hangzhou";
            }
            return "" + regionId + ".mns.aliyuncs.com";
        }

        public object DefaultAny(object inputValue, object defaultValue)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(inputValue))
            {
                return defaultValue;
            }
            return inputValue;
        }

        public string DefaultString(string inputValue, string defaultValue)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(inputValue) || AlibabaCloud.TeaUtil.Common.Empty(inputValue))
            {
                return defaultValue;
            }
            return inputValue;
        }

        public string Base64Encode(string input)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(input))
            {
                return "";
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.TeaUtil.Common.ToBytes(input));
        }

        public async Task<string> Base64EncodeAsync(string input)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(input))
            {
                return "";
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.TeaUtil.Common.ToBytes(input));
        }

        public string Base64Decode(string input)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(input))
            {
                return "";
            }
            return AlibabaCloud.TeaUtil.Common.ToString(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64Decode(input));
        }

        public async Task<string> Base64DecodeAsync(string input)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(input))
            {
                return "";
            }
            return AlibabaCloud.TeaUtil.Common.ToString(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64Decode(input));
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

        public string GetAuthorizationV2(string pathname, string method, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = GetSignatureV2(pathname, method, headers, secret);
            return "MNS " + ak + ":" + sign;
        }

        public async Task<string> GetAuthorizationV2Async(string pathname, string method, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = await GetSignatureV2Async(pathname, method, headers, secret);
            return "MNS " + ak + ":" + sign;
        }

        public string GetSignatureV2(string pathname, string method, Dictionary<string, string> headers, string secret)
        {
            string canonicalizedResource = DefaultString(pathname, "/");
            string canonicalizedHeaders = BuildCanonicalizedHeadersV2(headers);
            string stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureV2Async(string pathname, string method, Dictionary<string, string> headers, string secret)
        {
            string canonicalizedResource = DefaultString(pathname, "/");
            string canonicalizedHeaders = await BuildCanonicalizedHeadersV2Async(headers);
            string stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret));
        }

        public string BuildCanonicalizedHeadersV2(Dictionary<string, string> headers)
        {
            string contentMd5 = DefaultString(headers.Get("content-md5"), "");
            string contentType = DefaultString(headers.Get("content-type"), "");
            string date = DefaultString(headers.Get("date"), "");
            string canonicalizedHeaders = "" + contentMd5 + "\n" + contentType + "\n" + date + "\n";
            Dictionary<string, string> mnsHeaders = new Dictionary<string, string>(){};

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerHeader = AlibabaCloud.DarabonbaString.StringUtil.ToLower(header);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerHeader, "x-mns-"))
                {
                    mnsHeaders[lowerHeader] = headers.Get(header);
                }
            }

            foreach (var header in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(mnsHeaders))) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + mnsHeaders.Get(header) + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersV2Async(Dictionary<string, string> headers)
        {
            string contentMd5 = DefaultString(headers.Get("content-md5"), "");
            string contentType = DefaultString(headers.Get("content-type"), "");
            string date = DefaultString(headers.Get("date"), "");
            string canonicalizedHeaders = "" + contentMd5 + "\n" + contentType + "\n" + date + "\n";
            Dictionary<string, string> mnsHeaders = new Dictionary<string, string>(){};

            foreach (var header in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerHeader = AlibabaCloud.DarabonbaString.StringUtil.ToLower(header);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerHeader, "x-mns-"))
                {
                    mnsHeaders[lowerHeader] = headers.Get(header);
                }
            }

            foreach (var header in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(mnsHeaders))) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + mnsHeaders.Get(header) + "\n";
            }
            return canonicalizedHeaders;
        }

        public string GetAuthorizationV4(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string date, string accessKeyId, string accessKeySecret)
        {
            string region = GetRegion(context);
            string dateShort = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 8);
            dateShort = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateShort, "T", "", null);
            string scope = "" + dateShort + "/" + region + "/mns/" + _signSuffix;
            byte[] signingkey = GetSigningkeyV4(accessKeySecret, region, dateShort);
            string signature = GetSignatureV4(context, date, scope, signingkey);
            return "" + _authPrefix + " Credential=" + accessKeyId + "/" + scope + ",Signature=" + signature;
        }

        public async Task<string> GetAuthorizationV4Async(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string date, string accessKeyId, string accessKeySecret)
        {
            string region = await GetRegionAsync(context);
            string dateShort = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 8);
            dateShort = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateShort, "T", "", null);
            string scope = "" + dateShort + "/" + region + "/mns/" + _signSuffix;
            byte[] signingkey = await GetSigningkeyV4Async(accessKeySecret, region, dateShort);
            string signature = await GetSignatureV4Async(context, date, scope, signingkey);
            return "" + _authPrefix + " Credential=" + accessKeyId + "/" + scope + ",Signature=" + signature;
        }

        public byte[] GetSigningkeyV4(string secret, string region, string date)
        {
            string sc1 = "" + _signPrefix + secret;
            byte[] sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            byte[] sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            byte[] sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("mns", sc3);
            return AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(_signSuffix, sc4);
        }

        public async Task<byte[]> GetSigningkeyV4Async(string secret, string region, string date)
        {
            string sc1 = "" + _signPrefix + secret;
            byte[] sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            byte[] sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            byte[] sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("mns", sc3);
            return AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(_signSuffix, sc4);
        }

        public string GetSignatureV4(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string date, string scope, byte[] signingkey)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            string canonicalString = BuildCanonicalRequestV4(request.Pathname, request.Method, request.Query, request.Headers);
            string stringToSign = "" + _authPrefix + "\n" + date + "\n" + scope + "\n" + canonicalString;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public async Task<string> GetSignatureV4Async(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, string date, string scope, byte[] signingkey)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            string canonicalString = await BuildCanonicalRequestV4Async(request.Pathname, request.Method, request.Query, request.Headers);
            string stringToSign = "" + _authPrefix + "\n" + date + "\n" + scope + "\n" + canonicalString;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public string BuildCanonicalRequestV4(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers)
        {
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                canonicalURI = pathname;
                if (!AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(canonicalURI, "/"))
                {
                    canonicalURI = "/" + canonicalURI;
                }
            }
            string suffix = "";
            if (!AlibabaCloud.DarabonbaString.StringUtil.Equals(canonicalURI, "/") && AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(canonicalURI, "/"))
            {
                suffix = "/";
            }
            canonicalURI = "" + AlibabaCloud.OpenApiUtil.Client.GetEncodePath(canonicalURI) + suffix;
            string resources = BuildCanonicalizedResourceV4(query);
            string canonicalHeaders = BuildCanonicalizedHeadersV4(headers);
            return "" + method + "\n" + canonicalURI + "\n" + resources + "\n" + canonicalHeaders + "\n";
        }

        public async Task<string> BuildCanonicalRequestV4Async(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers)
        {
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                canonicalURI = pathname;
                if (!AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(canonicalURI, "/"))
                {
                    canonicalURI = "/" + canonicalURI;
                }
            }
            string suffix = "";
            if (!AlibabaCloud.DarabonbaString.StringUtil.Equals(canonicalURI, "/") && AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(canonicalURI, "/"))
            {
                suffix = "/";
            }
            canonicalURI = "" + AlibabaCloud.OpenApiUtil.Client.GetEncodePath(canonicalURI) + suffix;
            string resources = await BuildCanonicalizedResourceV4Async(query);
            string canonicalHeaders = await BuildCanonicalizedHeadersV4Async(headers);
            return "" + method + "\n" + canonicalURI + "\n" + resources + "\n" + canonicalHeaders + "\n";
        }

        public string BuildCanonicalizedResourceV4(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                Dictionary<string, string> queryMap = new Dictionary<string, string>(){};

                foreach (var key in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query)) {
                    string encodedKey = PercentEncodeMns(AlibabaCloud.DarabonbaString.StringUtil.ToLower(AlibabaCloud.DarabonbaString.StringUtil.Trim(key)));
                    string encodedValue = "";
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(key)) && !AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        encodedValue = PercentEncodeMns(AlibabaCloud.DarabonbaString.StringUtil.Trim(query.Get(key)));
                    }
                    queryMap[encodedKey] = encodedValue;
                }
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var encodedKey in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + encodedKey;
                    string encodedValue = queryMap.Get(encodedKey);
                    if (!AlibabaCloud.TeaUtil.Common.Empty(encodedValue))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + encodedValue;
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
                Dictionary<string, string> queryMap = new Dictionary<string, string>(){};

                foreach (var key in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query)) {
                    string encodedKey = PercentEncodeMns(AlibabaCloud.DarabonbaString.StringUtil.ToLower(AlibabaCloud.DarabonbaString.StringUtil.Trim(key)));
                    string encodedValue = "";
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(key)) && !AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        encodedValue = PercentEncodeMns(AlibabaCloud.DarabonbaString.StringUtil.Trim(query.Get(key)));
                    }
                    queryMap[encodedKey] = encodedValue;
                }
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(queryMap);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var encodedKey in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + encodedKey;
                    string encodedValue = queryMap.Get(encodedKey);
                    if (!AlibabaCloud.TeaUtil.Common.Empty(encodedValue))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + encodedValue;
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeadersV4(Dictionary<string, string> headers)
        {
            Dictionary<string, string> signedHeaders = new Dictionary<string, string>(){};

            foreach (var key in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (HasSignedHeaderV4(lowerKey))
                {
                    signedHeaders[lowerKey] = AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(key));
                }
            }
            string canonicalizedHeaders = "";

            foreach (var lowerKey in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(signedHeaders))) {
                canonicalizedHeaders = "" + canonicalizedHeaders + lowerKey + ":" + signedHeaders.Get(lowerKey) + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersV4Async(Dictionary<string, string> headers)
        {
            Dictionary<string, string> signedHeaders = new Dictionary<string, string>(){};

            foreach (var key in AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers)) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (HasSignedHeaderV4(lowerKey))
                {
                    signedHeaders[lowerKey] = AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(key));
                }
            }
            string canonicalizedHeaders = "";

            foreach (var lowerKey in AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(AlibabaCloud.DarabonbaMap.MapUtil.KeySet(signedHeaders))) {
                canonicalizedHeaders = "" + canonicalizedHeaders + lowerKey + ":" + signedHeaders.Get(lowerKey) + "\n";
            }
            return canonicalizedHeaders;
        }

        public bool HasSignedHeaderV4(string header)
        {
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(header, "content-type") || AlibabaCloud.DarabonbaString.StringUtil.Equals(header, "content-md5"))
            {
                return true;
            }
            return AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(header, "x-mns-");
        }

        public string PercentEncodeMns(string value)
        {
            string encoded = AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(value);
            return AlibabaCloud.DarabonbaString.StringUtil.Replace(encoded, "+", "%20", null);
        }

        public string GetRegion(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(config.RegionId) && !AlibabaCloud.TeaUtil.Common.Empty(config.RegionId))
            {
                return config.RegionId;
            }
            string region = AlibabaCloud.DarabonbaString.StringUtil.Replace(config.Endpoint, ".mns.aliyuncs.com", "", null);
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(region, config.Endpoint))
            {
                throw new TeaException(new Dictionary<string, string>
                {
                    {"code", "ClientConfigError"},
                    {"message", "The regionId configuration of mns client is missing."},
                });
            }
            return region;
        }

        public async Task<string> GetRegionAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(config.RegionId) && !AlibabaCloud.TeaUtil.Common.Empty(config.RegionId))
            {
                return config.RegionId;
            }
            string region = AlibabaCloud.DarabonbaString.StringUtil.Replace(config.Endpoint, ".mns.aliyuncs.com", "", null);
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(region, config.Endpoint))
            {
                throw new TeaException(new Dictionary<string, string>
                {
                    {"code", "ClientConfigError"},
                    {"message", "The regionId configuration of mns client is missing."},
                });
            }
            return region;
        }

        public string GetDateISO8601()
        {
            string date = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            date = AlibabaCloud.DarabonbaString.StringUtil.Replace(date, "-", "", null);
            return AlibabaCloud.DarabonbaString.StringUtil.Replace(date, ":", "", null);
        }

        public async Task<string> GetDateISO8601Async()
        {
            string date = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            date = AlibabaCloud.DarabonbaString.StringUtil.Replace(date, "-", "", null);
            return AlibabaCloud.DarabonbaString.StringUtil.Replace(date, ":", "", null);
        }

    }
}
