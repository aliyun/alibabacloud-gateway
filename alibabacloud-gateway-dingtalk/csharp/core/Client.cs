// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewayDingTalk
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {

        public Client(): base()
        {
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"x-acs-signature-method", "HMAC-SHA1"},
                    {"x-acs-signature-version", "1.0"},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                string jsonObj = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                request.Stream = TeaCore.BytesReadable(jsonObj);
                request.Headers["content-type"] = "application/json; charset=utf-8";
            }
            if (!AlibabaCloud.TeaUtil.Common.EqualString(request.AuthType, "Anonymous"))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                string accessKeyId = credential.GetAccessKeyId();
                string accessKeySecret = credential.GetAccessKeySecret();
                string securityToken = credential.GetSecurityToken();
                if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                {
                    request.Headers["x-acs-accesskey-id"] = accessKeyId;
                    request.Headers["x-acs-security-token"] = securityToken;
                }
                request.Headers["Authorization"] = GetAuthorization(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
            }
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"x-acs-signature-method", "HMAC-SHA1"},
                    {"x-acs-signature-version", "1.0"},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                string jsonObj = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                request.Stream = TeaCore.BytesReadable(jsonObj);
                request.Headers["content-type"] = "application/json; charset=utf-8";
            }
            if (!AlibabaCloud.TeaUtil.Common.EqualString(request.AuthType, "Anonymous"))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                string accessKeyId = await credential.GetAccessKeyIdAsync();
                string accessKeySecret = await credential.GetAccessKeySecretAsync();
                string securityToken = await credential.GetSecurityTokenAsync();
                if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                {
                    request.Headers["x-acs-accesskey-id"] = accessKeyId;
                    request.Headers["x-acs-security-token"] = securityToken;
                }
                request.Headers["Authorization"] = await GetAuthorizationAsync(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
            }
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                err["statusCode"] = response.StatusCode;
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "" + DefaultAny(err.Get("Code"), err.Get("code"))},
                    {"message", "code: " + response.StatusCode + ", " + DefaultAny(err.Get("Message"), err.Get("message")) + " request id: " + DefaultAny(err.Get("requestId"), err.Get("requestid"))},
                    {"data", err},
                    {"description", "" + DefaultAny(err.Get("Description"), err.Get("description"))},
                    {"accessDeniedDetail", DefaultAny(err.Get("accessDeniedDetail"), err.Get("accessdenieddetail"))},
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
                object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                err["statusCode"] = response.StatusCode;
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "" + DefaultAny(err.Get("Code"), err.Get("code"))},
                    {"message", "code: " + response.StatusCode + ", " + DefaultAny(err.Get("Message"), err.Get("message")) + " request id: " + DefaultAny(err.Get("requestId"), err.Get("requestid"))},
                    {"data", err},
                    {"description", "" + DefaultAny(err.Get("Description"), err.Get("description"))},
                    {"accessDeniedDetail", DefaultAny(err.Get("accessDeniedDetail"), err.Get("accessdenieddetail"))},
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

        public object DefaultAny(object inputValue, object defaultValue)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(inputValue))
            {
                return defaultValue;
            }
            return inputValue;
        }

        public string GetAuthorization(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string signature = GetSignature(pathname, method, query, headers, secret);
            return "acs " + ak + ":" + signature;
        }

        public async Task<string> GetAuthorizationAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string signature = await GetSignatureAsync(pathname, method, query, headers, secret);
            return "acs " + ak + ":" + signature;
        }

        public string GetSignature(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(pathname, query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(signature);
        }

        public async Task<string> GetSignatureAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(pathname, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            stringToSign = "" + method + "\n" + canonicalizedHeaders + canonicalizedResource;
            byte[] signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA1Sign(stringToSign, secret);
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(signature);
        }

        public string BuildCanonicalizedResource(string pathname, Dictionary<string, string> query)
        {
            string canonicalizedResource = pathname;
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "?";

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + key;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(key);
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
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "?";

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + key;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + query.Get(key);
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeaders(Dictionary<string, string> headers)
        {
            string accept = headers.Get("accept");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(accept))
            {
                accept = "";
            }
            string contentMd5 = headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            string contentType = headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string date = headers.Get("date");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(date))
            {
                date = "";
            }
            string canonicalizedHeaders = "" + accept + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n";
            List<string> sortedHeaders = GetSignedHeaders(headers);

            foreach (var header in sortedHeaders) {
                string value = headers.Get(header);
                string valueTrim = AlibabaCloud.DarabonbaString.StringUtil.Trim(value);
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + valueTrim + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersAsync(Dictionary<string, string> headers)
        {
            string accept = headers.Get("accept");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(accept))
            {
                accept = "";
            }
            string contentMd5 = headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            string contentType = headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string date = headers.Get("date");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(date))
            {
                date = "";
            }
            string canonicalizedHeaders = "" + accept + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n";
            List<string> sortedHeaders = await GetSignedHeadersAsync(headers);

            foreach (var header in sortedHeaders) {
                string value = headers.Get(header);
                string valueTrim = AlibabaCloud.DarabonbaString.StringUtil.Trim(value);
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + valueTrim + "\n";
            }
            return canonicalizedHeaders;
        }

        public List<string> GetSignedHeaders(Dictionary<string, string> headers)
        {
            List<string> headersArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeadersArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(headersArray);
            string tmp = "";
            string separator = "";

            foreach (var key in sortedHeadersArray) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-acs-"))
                {
                    if (!AlibabaCloud.DarabonbaString.StringUtil.Contains(tmp, lowerKey))
                    {
                        tmp = "" + tmp + separator + lowerKey;
                        separator = ";";
                    }
                }
            }
            return AlibabaCloud.DarabonbaString.StringUtil.Split(tmp, ";", 0);
        }

        public async Task<List<string>> GetSignedHeadersAsync(Dictionary<string, string> headers)
        {
            List<string> headersArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeadersArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(headersArray);
            string tmp = "";
            string separator = "";

            foreach (var key in sortedHeadersArray) {
                string lowerKey = AlibabaCloud.DarabonbaString.StringUtil.ToLower(key);
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-acs-"))
                {
                    if (!AlibabaCloud.DarabonbaString.StringUtil.Contains(tmp, lowerKey))
                    {
                        tmp = "" + tmp + separator + lowerKey;
                        separator = ";";
                    }
                }
            }
            return AlibabaCloud.DarabonbaString.StringUtil.Split(tmp, ";", 0);
        }

    }
}
