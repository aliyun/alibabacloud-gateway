// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewayPds
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {

        public Client(): base()
        {
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = GetEndpoint(config.Network, config.Endpoint);
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = await GetEndpointAsync(config.Network, config.Endpoint);
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string date = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"date", date},
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            string signatureAlgorithm = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureAlgorithm, "ACS4-HMAC-SHA256");
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v1");
            string hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(""), signatureAlgorithm));
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Stream))
            {
                byte[] tmp = AlibabaCloud.TeaUtil.Common.ReadAsBytes(request.Stream);
                hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(tmp, signatureAlgorithm));
                request.Stream = TeaCore.BytesReadable(tmp);
                request.Headers["content-type"] = "application/octet-stream";
            }
            else
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
                {
                    if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "json"))
                    {
                        string jsonObj = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                        hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(jsonObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(jsonObj);
                        request.Headers["content-type"] = "application/json; charset=utf-8";
                    }
                    else
                    {
                        Dictionary<string, object> m = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                        string formObj = AlibabaCloud.OpenApiUtil.Client.ToForm(m);
                        hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(formObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(formObj);
                        request.Headers["content-type"] = "application/x-www-form-urlencoded";
                    }
                }
            }
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
                {
                    request.Headers["x-acs-content-sm3"] = hashedRequestPayload;
                }
                else
                {
                    request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
                }
            }
            else
            {
                request.Headers["x-acs-signature-method"] = "HMAC-SHA1";
                request.Headers["x-acs-signature-version"] = "1.0";
            }
            if (!AlibabaCloud.TeaUtil.Common.EqualString(request.AuthType, "Anonymous") && !AlibabaCloud.TeaUtil.Common.IsUnset(request.Credential))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                Aliyun.Credentials.Models.CredentialModel credentialModel = credential.GetCredential();
                string authType = credentialModel.Type;
                if (AlibabaCloud.TeaUtil.Common.EqualString(authType, "bearer"))
                {
                    string bearerToken = credentialModel.BearerToken;
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else
                {
                    string accessKeyId = credentialModel.AccessKeyId;
                    string accessKeySecret = credentialModel.AccessKeySecret;
                    string securityToken = credentialModel.SecurityToken;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                    {
                        request.Headers["x-acs-security-token"] = securityToken;
                    }
                    if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
                    {
                        string dateNew = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 10);
                        string region = GetRegion(config.Endpoint);
                        byte[] signingkey = GetSigningkey(signatureAlgorithm, accessKeySecret, region, dateNew);
                        request.Headers["Authorization"] = GetAuthorizationV4(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew);
                    }
                    else
                    {
                        request.Headers["Authorization"] = GetAuthorization(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
                    }
                }
            }
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string date = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"date", date},
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            string signatureAlgorithm = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureAlgorithm, "ACS4-HMAC-SHA256");
            string signatureVersion = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureVersion, "v1");
            string hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(""), signatureAlgorithm));
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Stream))
            {
                byte[] tmp = AlibabaCloud.TeaUtil.Common.ReadAsBytes(request.Stream);
                hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(tmp, signatureAlgorithm));
                request.Stream = TeaCore.BytesReadable(tmp);
                request.Headers["content-type"] = "application/octet-stream";
            }
            else
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
                {
                    if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "json"))
                    {
                        string jsonObj = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                        hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(jsonObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(jsonObj);
                        request.Headers["content-type"] = "application/json; charset=utf-8";
                    }
                    else
                    {
                        Dictionary<string, object> m = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                        string formObj = AlibabaCloud.OpenApiUtil.Client.ToForm(m);
                        hashedRequestPayload = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(formObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(formObj);
                        request.Headers["content-type"] = "application/x-www-form-urlencoded";
                    }
                }
            }
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
                {
                    request.Headers["x-acs-content-sm3"] = hashedRequestPayload;
                }
                else
                {
                    request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
                }
            }
            else
            {
                request.Headers["x-acs-signature-method"] = "HMAC-SHA1";
                request.Headers["x-acs-signature-version"] = "1.0";
            }
            if (!AlibabaCloud.TeaUtil.Common.EqualString(request.AuthType, "Anonymous") && !AlibabaCloud.TeaUtil.Common.IsUnset(request.Credential))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                Aliyun.Credentials.Models.CredentialModel credentialModel = await credential.GetCredentialAsync();
                string authType = credentialModel.Type;
                if (AlibabaCloud.TeaUtil.Common.EqualString(authType, "bearer"))
                {
                    string bearerToken = credentialModel.BearerToken;
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else
                {
                    string accessKeyId = credentialModel.AccessKeyId;
                    string accessKeySecret = credentialModel.AccessKeySecret;
                    string securityToken = credentialModel.SecurityToken;
                    if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                    {
                        request.Headers["x-acs-security-token"] = securityToken;
                    }
                    if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureVersion, "v4"))
                    {
                        string dateNew = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 10);
                        string region = GetRegion(config.Endpoint);
                        byte[] signingkey = await GetSigningkeyAsync(signatureAlgorithm, accessKeySecret, region, dateNew);
                        request.Headers["Authorization"] = await GetAuthorizationV4Async(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew);
                    }
                    else
                    {
                        request.Headers["Authorization"] = await GetAuthorizationAsync(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
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
                object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                Dictionary<string, string> headers = response.Headers;
                string requestId = headers.Get("x-ca-request-id");
                err["statusCode"] = response.StatusCode;
                err["requestId"] = requestId;
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "" + DefaultAny(err.Get("Code"), err.Get("code"))},
                    {"message", "" + DefaultAny(err.Get("Message"), err.Get("message"))},
                    {"data", err},
                });
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
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
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
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
                object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                Dictionary<string, string> headers = response.Headers;
                string requestId = headers.Get("x-ca-request-id");
                err["statusCode"] = response.StatusCode;
                err["requestId"] = requestId;
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "" + DefaultAny(err.Get("Code"), err.Get("code"))},
                    {"message", "" + DefaultAny(err.Get("Message"), err.Get("message"))},
                    {"data", err},
                });
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Body))
            {
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
                    response.DeserializedBody = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
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

        public string GetEndpoint(string network, string endpoint)
        {
            string realEndpoint = "api.aliyunpds.com";
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                realEndpoint = endpoint;
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(network) && AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "vpc"))
            {
                realEndpoint = AlibabaCloud.DarabonbaString.StringUtil.Replace(realEndpoint, "api.aliyunpds.com", "api-vpc.aliyunpds.com", null);
                realEndpoint = AlibabaCloud.DarabonbaString.StringUtil.Replace(realEndpoint, "admin.aliyunpds.com", "admin-vpc.aliyunpds.com", null);
            }
            return realEndpoint;
        }

        public async Task<string> GetEndpointAsync(string network, string endpoint)
        {
            string realEndpoint = "api.aliyunpds.com";
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                realEndpoint = endpoint;
            }
            if (!AlibabaCloud.TeaUtil.Common.Empty(network) && AlibabaCloud.DarabonbaString.StringUtil.Equals(network, "vpc"))
            {
                realEndpoint = AlibabaCloud.DarabonbaString.StringUtil.Replace(realEndpoint, "api.aliyunpds.com", "api-vpc.aliyunpds.com", null);
                realEndpoint = AlibabaCloud.DarabonbaString.StringUtil.Replace(realEndpoint, "admin.aliyunpds.com", "admin-vpc.aliyunpds.com", null);
            }
            return realEndpoint;
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
            return AlibabaCloud.DarabonbaString.StringUtil.Split(tmp, ";", null);
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
            return AlibabaCloud.DarabonbaString.StringUtil.Split(tmp, ";", null);
        }

        public string GetRegion(string endpoint)
        {
            string region = "center";
            if (AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return region;
            }
            if (AlibabaCloud.DarabonbaString.StringUtil.Contains(endpoint, ".admin.aliyunpds.com"))
            {
                region = AlibabaCloud.DarabonbaString.StringUtil.Replace(endpoint, ".admin.aliyunpds.com", "", null);
            }
            return region;
        }

        public byte[] GetSigningkey(string signatureAlgorithm, string secret, string region, string date)
        {
            string sc1 = "aliyun_v4" + secret;
            byte[] sc2 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3Sign(date, sc1);
            }
            byte[] sc3 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(region, sc2);
            }
            byte[] sc4 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("pds", sc3);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes("pds", sc3);
            }
            byte[] hmac = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes("aliyun_v4_request", sc4);
            }
            return hmac;
        }

        public async Task<byte[]> GetSigningkeyAsync(string signatureAlgorithm, string secret, string region, string date)
        {
            string sc1 = "aliyun_v4" + secret;
            byte[] sc2 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3Sign(date, sc1);
            }
            byte[] sc3 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(region, sc2);
            }
            byte[] sc4 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("pds", sc3);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes("pds", sc3);
            }
            byte[] hmac = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes("aliyun_v4_request", sc4);
            }
            return hmac;
        }

        public string GetAuthorizationV4(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, byte[] signingkey, string product, string region, string date)
        {
            string signature = GetSignatureV4(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
            List<string> signedHeaders = GetSignedHeaders(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/aliyun_v4_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature;
        }

        public async Task<string> GetAuthorizationV4Async(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, byte[] signingkey, string product, string region, string date)
        {
            string signature = await GetSignatureV4Async(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
            List<string> signedHeaders = await GetSignedHeadersAsync(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/aliyun_v4_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature;
        }

        public string GetSignatureV4(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, byte[] signingkey)
        {
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(pathname, query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            List<string> signedHeaders = GetSignedHeaders(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            stringToSign = "" + method + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(stringToSign, signingkey);
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public async Task<string> GetSignatureV4Async(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, byte[] signingkey)
        {
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(pathname, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            List<string> signedHeaders = await GetSignedHeadersAsync(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            stringToSign = "" + method + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SHA256"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, "ACS4-HMAC-SM3"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(stringToSign, signingkey);
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

    }
}
