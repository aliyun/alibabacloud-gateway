// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;


namespace AlibabaCloud.GatewayPop
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {
        protected string _sha256;
        protected string _sm3;

        public Client(): base()
        {
            this._sha256 = "ACS4-HMAC-SHA256";
            this._sm3 = "ACS4-HMAC-SM3";
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = GetEndpoint(request.ProductId, config.RegionId, config.EndpointRule, config.Network, config.Suffix, config.EndpointMap, config.Endpoint);
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            config.Endpoint = GetEndpoint(request.ProductId, config.RegionId, config.EndpointRule, config.Network, config.Suffix, config.EndpointMap, config.Endpoint);
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string date = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-date", date},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            string signatureAlgorithm = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureAlgorithm, _sha256);
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
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                request.Headers["x-acs-content-sm3"] = hashedRequestPayload;
            }
            else
            {
                request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
            }
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
                string authType = credential.GetType();
                if (AlibabaCloud.TeaUtil.Common.EqualString(authType, "bearer"))
                {
                    string bearerToken = credential.GetBearerToken();
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else
                {
                    string accessKeyId = credential.GetAccessKeyId();
                    string accessKeySecret = credential.GetAccessKeySecret();
                    string securityToken = credential.GetSecurityToken();
                    if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                    {
                        request.Headers["x-acs-accesskey-id"] = accessKeyId;
                        request.Headers["x-acs-security-token"] = securityToken;
                    }
                    string dateNew = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 10);
                    dateNew = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateNew, "-", "", null);
                    string region = GetRegion(request.ProductId, config.Endpoint);
                    byte[] signingkey = GetSigningkey(signatureAlgorithm, accessKeySecret, request.ProductId, region, dateNew);
                    request.Headers["Authorization"] = GetAuthorization(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew);
                }
            }
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string date = AlibabaCloud.OpenApiUtil.Client.GetTimestamp();
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-date", date},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            string signatureAlgorithm = AlibabaCloud.TeaUtil.Common.DefaultString(request.SignatureAlgorithm, _sha256);
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
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                request.Headers["x-acs-content-sm3"] = hashedRequestPayload;
            }
            else
            {
                request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
            }
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
                string authType = credential.GetType();
                if (AlibabaCloud.TeaUtil.Common.EqualString(authType, "bearer"))
                {
                    string bearerToken = credential.GetBearerToken();
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else
                {
                    string accessKeyId = await credential.GetAccessKeyIdAsync();
                    string accessKeySecret = await credential.GetAccessKeySecretAsync();
                    string securityToken = await credential.GetSecurityTokenAsync();
                    if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
                    {
                        request.Headers["x-acs-accesskey-id"] = accessKeyId;
                        request.Headers["x-acs-security-token"] = securityToken;
                    }
                    string dateNew = AlibabaCloud.DarabonbaString.StringUtil.SubString(date, 0, 10);
                    dateNew = AlibabaCloud.DarabonbaString.StringUtil.Replace(dateNew, "-", "", null);
                    string region = GetRegion(request.ProductId, config.Endpoint);
                    byte[] signingkey = await GetSigningkeyAsync(signatureAlgorithm, accessKeySecret, request.ProductId, region, dateNew);
                    request.Headers["Authorization"] = await GetAuthorizationAsync(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew);
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
                object requestId = DefaultAny(err.Get("RequestId"), err.Get("requestId"));
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-acs-request-id")))
                {
                    requestId = response.Headers.Get("x-acs-request-id");
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
                object _res = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                Dictionary<string, object> err = AlibabaCloud.TeaUtil.Common.AssertAsMap(_res);
                object requestId = DefaultAny(err.Get("RequestId"), err.Get("requestId"));
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(response.Headers.Get("x-acs-request-id")))
                {
                    requestId = response.Headers.Get("x-acs-request-id");
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

        public string GetEndpoint(string productId, string regionId, string endpointRule, string network, string suffix, Dictionary<string, string> endpointMap, string endpoint)
        {
            if (!AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return endpoint;
            }
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(endpointMap) && !AlibabaCloud.TeaUtil.Common.Empty(endpointMap.Get(regionId)))
            {
                return endpointMap.Get(regionId);
            }
            return AlibabaCloud.EndpointUtil.Common.GetEndpointRules(productId, regionId, endpointRule, network, suffix);
        }

        public object DefaultAny(object inputValue, object defaultValue)
        {
            if (AlibabaCloud.TeaUtil.Common.IsUnset(inputValue))
            {
                return defaultValue;
            }
            return inputValue;
        }

        public string GetAuthorization(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, byte[] signingkey, string product, string region, string date)
        {
            string signature = GetSignature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
            List<string> signedHeaders = GetSignedHeaders(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/aliyun_v4_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature;
        }

        public async Task<string> GetAuthorizationAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, byte[] signingkey, string product, string region, string date)
        {
            string signature = await GetSignatureAsync(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
            List<string> signedHeaders = await GetSignedHeadersAsync(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/aliyun_v4_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature;
        }

        public string GetSignature(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, byte[] signingkey)
        {
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                canonicalURI = pathname;
            }
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            List<string> signedHeaders = GetSignedHeaders(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(stringToSign, signingkey);
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public async Task<string> GetSignatureAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, byte[] signingkey)
        {
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                canonicalURI = pathname;
            }
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(headers);
            List<string> signedHeaders = await GetSignedHeadersAsync(headers);
            string signedHeadersStr = AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";");
            stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(stringToSign, signingkey);
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public byte[] GetSigningkey(string signatureAlgorithm, string secret, string product, string region, string date)
        {
            string sc1 = "aliyun_v4" + secret;
            byte[] sc2 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3Sign(date, sc1);
            }
            byte[] sc3 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(region, sc2);
            }
            byte[] sc4 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(product, sc3);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(product, sc3);
            }
            byte[] hmac = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes("aliyun_v4_request", sc4);
            }
            return hmac;
        }

        public async Task<byte[]> GetSigningkeyAsync(string signatureAlgorithm, string secret, string product, string region, string date)
        {
            string sc1 = "aliyun_v4" + secret;
            byte[] sc2 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(date, sc1);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                sc2 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3Sign(date, sc1);
            }
            byte[] sc3 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(region, sc2);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                sc3 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(region, sc2);
            }
            byte[] sc4 = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes(product, sc3);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                sc4 = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes(product, sc3);
            }
            byte[] hmac = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sha256))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
            }
            else if (AlibabaCloud.TeaUtil.Common.EqualString(signatureAlgorithm, _sm3))
            {
                hmac = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3SignByBytes("aliyun_v4_request", sc4);
            }
            return hmac;
        }

        public string GetRegion(string product, string endpoint)
        {
            string region = "center";
            if (AlibabaCloud.TeaUtil.Common.Empty(product) || AlibabaCloud.TeaUtil.Common.Empty(endpoint))
            {
                return region;
            }
            List<string> strs = AlibabaCloud.DarabonbaString.StringUtil.Split(endpoint, ":", null);
            string withoutPort = strs[0];
            string preRegion = AlibabaCloud.DarabonbaString.StringUtil.Replace(withoutPort, ".aliyuncs.com", "", null);
            List<string> nodes = AlibabaCloud.DarabonbaString.StringUtil.Split(preRegion, ".", null);
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(nodes), 2))
            {
                region = nodes[1];
            }
            return region;
        }

        public string BuildCanonicalizedResource(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(key);
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(key));
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceAsync(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);
                string separator = "";

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + separator + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(key);
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(key));
                    }
                    separator = "&";
                }
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeaders(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> sortedHeaders = GetSignedHeaders(headers);

            foreach (var header in sortedHeaders) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(header)) + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersAsync(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> sortedHeaders = await GetSignedHeadersAsync(headers);

            foreach (var header in sortedHeaders) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(header)) + "\n";
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
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-acs-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "host") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type"))
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
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-acs-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "host") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type"))
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

    }
}
