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
        protected string _endpointSuffix;
        protected string _signatureTypePrefix;
        protected string _signPrefix;
        protected string _sha256;
        protected string _sm3;

        public Client(): base()
        {
            // CLOUD4-
            this._signatureTypePrefix = "ACS4-";
            // cloud_v4
            this._signPrefix = "aliyun_v4";
            this._endpointSuffix = "aliyuncs.com";
            this._sha256 = "" + _signatureTypePrefix + "HMAC-SHA256";
            this._sm3 = "" + _signatureTypePrefix + "HMAC-SM3";
        }


        public void ModifyConfiguration(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string endpointSuffix = ResolveEndpointSuffix(attributeMap.Key);
            config.Endpoint = GetEndpointWithSuffix(request.ProductId, config.RegionId, config.EndpointRule, config.Network, config.Suffix, config.EndpointMap, config.Endpoint, endpointSuffix);
        }

        public async Task ModifyConfigurationAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string endpointSuffix = ResolveEndpointSuffix(attributeMap.Key);
            config.Endpoint = GetEndpointWithSuffix(request.ProductId, config.RegionId, config.EndpointRule, config.Network, config.Suffix, config.EndpointMap, config.Endpoint, endpointSuffix);
        }

        internal string[] SignatureNames(Dictionary<string, string> attributes)
        {
            string prefix = "ACS4-";
            string signPrefix = "aliyun_v4";
            if (!Runtime.IsUnset(attributes))
            {
                string requested = attributes.Get("signatureTypePrefix");
                if (!Runtime.Empty(requested))
                {
                    prefix = requested;
                }
                string requestedSign = attributes.Get("signPrefix");
                if (!Runtime.Empty(requestedSign))
                {
                    signPrefix = requestedSign;
                }
            }
            return new string[] { signPrefix, prefix + "HMAC-SHA256", prefix + "HMAC-SM3" };
        }

        internal string CheckedEndpointSuffix(string suffix)
        {
            if (suffix == null)
            {
                return "";
            }
            suffix = suffix.Trim().ToLowerInvariant();
            if (suffix.Length == 0)
            {
                return "";
            }
            if ((suffix != "aliyuncs.com" && suffix != "aliyuncs.alibaba") || !System.Text.RegularExpressions.Regex.IsMatch(suffix, "^[a-z0-9]([a-z0-9-]*[a-z0-9])?(\\.[a-z0-9]([a-z0-9-]*[a-z0-9])?)+$"))
            {
                throw new TeaException(new Dictionary<string, object>
                {
                    {"code", "InvalidEndpointSuffix"},
                    {"message", "endpointSuffix must be aliyuncs.com or aliyuncs.alibaba"},
                });
            }
            return suffix;
        }

        internal string ResolveEndpointSuffix(Dictionary<string, string> attributes)
        {
            if (!Runtime.IsUnset(attributes))
            {
                string suffix = CheckedEndpointSuffix(attributes.Get("endpointSuffix"));
                if (!Runtime.Empty(suffix))
                {
                    return suffix;
                }
            }
            string envSuffix = CheckedEndpointSuffix(Environment.GetEnvironmentVariable("ALIBABA_CLOUD_ENDPOINT_SUFFIX"));
            if (!Runtime.Empty(envSuffix))
            {
                return envSuffix;
            }
            return "aliyuncs.com";
        }

        internal string WithEndpointSuffix(string endpoint, string endpointSuffix = null)
        {
            string suffix = endpointSuffix == null ? this._endpointSuffix : endpointSuffix;
            suffix = CheckedEndpointSuffix(suffix);
            if (Runtime.Empty(suffix))
            {
                suffix = "aliyuncs.com";
            }
            if (Runtime.Empty(endpoint) || suffix == "aliyuncs.com")
            {
                return endpoint;
            }
            string old = ".aliyuncs.com";
            if (endpoint.EndsWith(old))
            {
                return endpoint.Substring(0, endpoint.Length - old.Length) + "." + suffix;
            }
            return endpoint;
        }

        public void ModifyRequest(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string date = Runtime.GetTimestamp();
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"user-agent", request.UserAgent},
                    {"x-acs-date", date},
                    {"x-acs-signature-nonce", Runtime.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            if (!Runtime.Empty(request.Action))
            {
                request.Headers["x-acs-action"] = request.Action;
            }
            string[] names = SignatureNames(attributeMap.Key);
            string signatureAlgorithm = Runtime.DefaultString(request.SignatureAlgorithm, names[1]);
            string hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(""), signatureAlgorithm));
            if (!Runtime.IsUnset(request.Stream))
            {
                byte[] tmp = Runtime.ReadAsBytes(request.Stream);
                hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(tmp, signatureAlgorithm));
                request.Stream = TeaCore.BytesReadable(tmp);
                request.Headers["content-type"] = "application/octet-stream";
            }
            else
            {
                if (!Runtime.IsUnset(request.Body))
                {
                    if (Runtime.EqualString(request.ReqBodyType, "byte"))
                    {
                        byte[] byteObj = Runtime.AssertAsBytes(request.Body);
                        hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(byteObj, signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(byteObj);
                    }
                    else if (Runtime.EqualString(request.ReqBodyType, "json"))
                    {
                        string jsonObj = Runtime.ToJSONString(request.Body);
                        hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(jsonObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(jsonObj);
                        request.Headers["content-type"] = "application/json; charset=utf-8";
                    }
                    else
                    {
                        Dictionary<string, object> m = Runtime.AssertAsMap(request.Body);
                        string formObj = Runtime.ToForm(m);
                        hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(formObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(formObj);
                        request.Headers["content-type"] = "application/x-www-form-urlencoded";
                    }
                }
            }
            if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                request.Headers["x-acs-content-sm3"] = hashedRequestPayload;
            }
            else
            {
                request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
            }
            if (!Runtime.EqualString(request.AuthType, "Anonymous"))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                if (Runtime.IsUnset(credential))
                {
                    throw new TeaException(new Dictionary<string, string>
                    {
                        {"code", "ParameterMissing"},
                        {"message", "'config.credential' can not be unset"},
                    });
                }
                Aliyun.Credentials.Models.CredentialModel credentialModel = credential.GetCredential();
                if (!Runtime.Empty(credentialModel.ProviderName))
                {
                    request.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName;
                }
                string authType = credentialModel.Type;
                if (Runtime.EqualString(authType, "bearer"))
                {
                    string bearerToken = credentialModel.BearerToken;
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["x-acs-signature-type"] = "BEARERTOKEN";
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else if (Runtime.EqualString(authType, "id_token"))
                {
                    string idToken = credentialModel.SecurityToken;
                    request.Headers["x-acs-zero-trust-idtoken"] = idToken;
                }
                else
                {
                    string accessKeyId = credentialModel.AccessKeyId;
                    string accessKeySecret = credentialModel.AccessKeySecret;
                    string securityToken = credentialModel.SecurityToken;
                    if (!Runtime.Empty(securityToken))
                    {
                        request.Headers["x-acs-accesskey-id"] = accessKeyId;
                        request.Headers["x-acs-security-token"] = securityToken;
                    }
                    string dateNew = Runtime.SubString(date, 0, 10);
                    dateNew = Runtime.Replace(dateNew, "-", "", null);
                    string region = GetRegionWithSuffix(request.ProductId, config.Endpoint, config.RegionId, this._endpointSuffix);
                    byte[] signingkey = SigningKey(signatureAlgorithm, accessKeySecret, request.ProductId, region, dateNew, names);
                    request.Headers["Authorization"] = Authorization(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew, names);
                }
            }
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            string date = Runtime.GetTimestamp();
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"user-agent", request.UserAgent},
                    {"x-acs-date", date},
                    {"x-acs-signature-nonce", Runtime.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            if (!Runtime.Empty(request.Action))
            {
                request.Headers["x-acs-action"] = request.Action;
            }
            string[] names = SignatureNames(attributeMap.Key);
            string signatureAlgorithm = Runtime.DefaultString(request.SignatureAlgorithm, names[1]);
            string hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(""), signatureAlgorithm));
            if (!Runtime.IsUnset(request.Stream))
            {
                byte[] tmp = Runtime.ReadAsBytes(request.Stream);
                hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(tmp, signatureAlgorithm));
                request.Stream = TeaCore.BytesReadable(tmp);
                request.Headers["content-type"] = "application/octet-stream";
            }
            else
            {
                if (!Runtime.IsUnset(request.Body))
                {
                    if (Runtime.EqualString(request.ReqBodyType, "byte"))
                    {
                        byte[] byteObj = Runtime.AssertAsBytes(request.Body);
                        hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(byteObj, signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(byteObj);
                    }
                    else if (Runtime.EqualString(request.ReqBodyType, "json"))
                    {
                        string jsonObj = Runtime.ToJSONString(request.Body);
                        hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(jsonObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(jsonObj);
                        request.Headers["content-type"] = "application/json; charset=utf-8";
                    }
                    else
                    {
                        Dictionary<string, object> m = Runtime.AssertAsMap(request.Body);
                        string formObj = Runtime.ToForm(m);
                        hashedRequestPayload = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(formObj), signatureAlgorithm));
                        request.Stream = TeaCore.BytesReadable(formObj);
                        request.Headers["content-type"] = "application/x-www-form-urlencoded";
                    }
                }
            }
            if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                request.Headers["x-acs-content-sm3"] = hashedRequestPayload;
            }
            else
            {
                request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
            }
            if (!Runtime.EqualString(request.AuthType, "Anonymous"))
            {
                Aliyun.Credentials.Client credential = request.Credential;
                if (Runtime.IsUnset(credential))
                {
                    throw new TeaException(new Dictionary<string, string>
                    {
                        {"code", "ParameterMissing"},
                        {"message", "'config.credential' can not be unset"},
                    });
                }
                Aliyun.Credentials.Models.CredentialModel credentialModel = await credential.GetCredentialAsync();
                if (!Runtime.Empty(credentialModel.ProviderName))
                {
                    request.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName;
                }
                string authType = credentialModel.Type;
                if (Runtime.EqualString(authType, "bearer"))
                {
                    string bearerToken = credentialModel.BearerToken;
                    request.Headers["x-acs-bearer-token"] = bearerToken;
                    request.Headers["x-acs-signature-type"] = "BEARERTOKEN";
                    request.Headers["Authorization"] = "Bearer " + bearerToken;
                }
                else if (Runtime.EqualString(authType, "id_token"))
                {
                    string idToken = credentialModel.SecurityToken;
                    request.Headers["x-acs-zero-trust-idtoken"] = idToken;
                }
                else
                {
                    string accessKeyId = credentialModel.AccessKeyId;
                    string accessKeySecret = credentialModel.AccessKeySecret;
                    string securityToken = credentialModel.SecurityToken;
                    if (!Runtime.Empty(securityToken))
                    {
                        request.Headers["x-acs-accesskey-id"] = accessKeyId;
                        request.Headers["x-acs-security-token"] = securityToken;
                    }
                    string dateNew = Runtime.SubString(date, 0, 10);
                    dateNew = Runtime.Replace(dateNew, "-", "", null);
                    string region = GetRegionWithSuffix(request.ProductId, config.Endpoint, config.RegionId, this._endpointSuffix);
                    byte[] signingkey = SigningKey(signatureAlgorithm, accessKeySecret, request.ProductId, region, dateNew, names);
                    request.Headers["Authorization"] = Authorization(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.ProductId, region, dateNew, names);
                }
            }
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (Runtime.Is4xx(response.StatusCode) || Runtime.Is5xx(response.StatusCode))
            {
                Dictionary<string, object> err = new Dictionary<string, object>(){};
                if (!Runtime.IsUnset(response.Headers.Get("content-type")) && Runtime.Contains(response.Headers.Get("content-type"), "text/xml"))
                {
                    string _str = Runtime.ReadAsString(response.Body);
                    Dictionary<string, object> respMap = Runtime.ParseXml(_str);
                    err = Runtime.AssertAsMap(respMap.Get("Error"));
                }
                else
                {
                    object _res = Runtime.ReadAsJSON(response.Body);
                    err = Runtime.AssertAsMap(_res);
                }
                object requestId = DefaultAny(err.Get("RequestId"), err.Get("requestId"));
                if (!Runtime.IsUnset(response.Headers.Get("x-acs-request-id")))
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
            if (Runtime.EqualNumber(response.StatusCode, 204))
            {
                Runtime.ReadAsString(response.Body);
            }
            else if (Runtime.EqualString(request.BodyType, "binary"))
            {
                response.DeserializedBody = response.Body;
            }
            else if (Runtime.EqualString(request.BodyType, "byte"))
            {
                byte[] byt = Runtime.ReadAsBytes(response.Body);
                response.DeserializedBody = byt;
            }
            else if (Runtime.EqualString(request.BodyType, "string"))
            {
                string str = Runtime.ReadAsString(response.Body);
                response.DeserializedBody = str;
            }
            else if (Runtime.EqualString(request.BodyType, "json"))
            {
                object obj = Runtime.ReadAsJSON(response.Body);
                Dictionary<string, object> res = Runtime.AssertAsMap(obj);
                response.DeserializedBody = res;
            }
            else if (Runtime.EqualString(request.BodyType, "array"))
            {
                object arr = Runtime.ReadAsJSON(response.Body);
                response.DeserializedBody = arr;
            }
            else
            {
                response.DeserializedBody = Runtime.ReadAsString(response.Body);
            }
        }

        public async Task ModifyResponseAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (Runtime.Is4xx(response.StatusCode) || Runtime.Is5xx(response.StatusCode))
            {
                Dictionary<string, object> err = new Dictionary<string, object>(){};
                if (!Runtime.IsUnset(response.Headers.Get("content-type")) && Runtime.Contains(response.Headers.Get("content-type"), "text/xml"))
                {
                    string _str = Runtime.ReadAsString(response.Body);
                    Dictionary<string, object> respMap = Runtime.ParseXml(_str);
                    err = Runtime.AssertAsMap(respMap.Get("Error"));
                }
                else
                {
                    object _res = Runtime.ReadAsJSON(response.Body);
                    err = Runtime.AssertAsMap(_res);
                }
                object requestId = DefaultAny(err.Get("RequestId"), err.Get("requestId"));
                if (!Runtime.IsUnset(response.Headers.Get("x-acs-request-id")))
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
            if (Runtime.EqualNumber(response.StatusCode, 204))
            {
                Runtime.ReadAsString(response.Body);
            }
            else if (Runtime.EqualString(request.BodyType, "binary"))
            {
                response.DeserializedBody = response.Body;
            }
            else if (Runtime.EqualString(request.BodyType, "byte"))
            {
                byte[] byt = Runtime.ReadAsBytes(response.Body);
                response.DeserializedBody = byt;
            }
            else if (Runtime.EqualString(request.BodyType, "string"))
            {
                string str = Runtime.ReadAsString(response.Body);
                response.DeserializedBody = str;
            }
            else if (Runtime.EqualString(request.BodyType, "json"))
            {
                object obj = Runtime.ReadAsJSON(response.Body);
                Dictionary<string, object> res = Runtime.AssertAsMap(obj);
                response.DeserializedBody = res;
            }
            else if (Runtime.EqualString(request.BodyType, "array"))
            {
                object arr = Runtime.ReadAsJSON(response.Body);
                response.DeserializedBody = arr;
            }
            else
            {
                response.DeserializedBody = Runtime.ReadAsString(response.Body);
            }
        }

        public string GetEndpoint(string productId, string regionId, string endpointRule, string network, string suffix, Dictionary<string, string> endpointMap, string endpoint)
        {
            return GetEndpointWithSuffix(productId, regionId, endpointRule, network, suffix, endpointMap, endpoint, this._endpointSuffix);
        }

        internal string GetEndpointWithSuffix(string productId, string regionId, string endpointRule, string network, string suffix, Dictionary<string, string> endpointMap, string endpoint, string endpointSuffix)
        {
            if (!Runtime.Empty(endpoint))
            {
                return WithEndpointSuffix(endpoint, endpointSuffix);
            }
            if (!Runtime.IsUnset(endpointMap) && !Runtime.Empty(endpointMap.Get(regionId)))
            {
                return WithEndpointSuffix(endpointMap.Get(regionId), endpointSuffix);
            }
            return WithEndpointSuffix(Runtime.GetEndpointRules(productId, regionId, endpointRule, network, suffix), endpointSuffix);
        }

        public object DefaultAny(object inputValue, object defaultValue)
        {
            if (Runtime.IsUnset(inputValue))
            {
                return defaultValue;
            }
            return inputValue;
        }

        public string GetAuthorization(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, byte[] signingkey, string product, string region, string date)
        {
            return Authorization(pathname, method, query, headers, signatureAlgorithm, payload, ak, signingkey, product, region, date, new string[] { _signPrefix, _sha256, _sm3 });
        }

        internal string Authorization(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, byte[] signingkey, string product, string region, string date, string[] names)
        {
            string signature = Signature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey, names);
            List<string> signedHeaders = GetSignedHeaders(headers);
            string signedHeadersStr = Runtime.Join(signedHeaders, ";");
            return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/" + names[0] + "_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature;
        }

        public string GetSignature(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, byte[] signingkey)
        {
            return Signature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey, new string[] { _signPrefix, _sha256, _sm3 });
        }

        internal string Signature(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, byte[] signingkey, string[] names)
        {
            string canonicalURI = "/";
            if (!Runtime.Empty(pathname))
            {
                canonicalURI = pathname;
            }
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(query);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(headers);
            List<string> signedHeaders = GetSignedHeaders(headers);
            ValidateSignedHeaders(signedHeaders);
            string signedHeadersStr = Runtime.Join(signedHeaders, ";");
            stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload;
            string hex = Runtime.HexEncode(Runtime.Hash(Runtime.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = Runtime.ToBytes("");
            if (Runtime.EqualString(signatureAlgorithm, names[1]))
            {
                signature = Runtime.HmacSHA256SignByBytes(stringToSign, signingkey);
            }
            else if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                signature = Runtime.HmacSM3SignByBytes(stringToSign, signingkey);
            }
            return Runtime.HexEncode(signature);
        }

        public byte[] GetSigningkey(string signatureAlgorithm, string secret, string product, string region, string date)
        {
            return SigningKey(signatureAlgorithm, secret, product, region, date, new string[] { _signPrefix, _sha256, _sm3 });
        }

        internal byte[] SigningKey(string signatureAlgorithm, string secret, string product, string region, string date, string[] names)
        {
            string sc1 = "" + names[0] + secret;
            byte[] sc2 = Runtime.ToBytes("");
            if (Runtime.EqualString(signatureAlgorithm, names[1]))
            {
                sc2 = Runtime.HmacSHA256Sign(date, sc1);
            }
            else if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                sc2 = Runtime.HmacSM3Sign(date, sc1);
            }
            byte[] sc3 = Runtime.ToBytes("");
            if (Runtime.EqualString(signatureAlgorithm, names[1]))
            {
                sc3 = Runtime.HmacSHA256SignByBytes(region, sc2);
            }
            else if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                sc3 = Runtime.HmacSM3SignByBytes(region, sc2);
            }
            byte[] sc4 = Runtime.ToBytes("");
            if (Runtime.EqualString(signatureAlgorithm, names[1]))
            {
                sc4 = Runtime.HmacSHA256SignByBytes(product, sc3);
            }
            else if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                sc4 = Runtime.HmacSM3SignByBytes(product, sc3);
            }
            byte[] hmac = Runtime.ToBytes("");
            if (Runtime.EqualString(signatureAlgorithm, names[1]))
            {
                hmac = Runtime.HmacSHA256SignByBytes("" + names[0] + "_request", sc4);
            }
            else if (Runtime.EqualString(signatureAlgorithm, names[2]))
            {
                hmac = Runtime.HmacSM3SignByBytes("" + names[0] + "_request", sc4);
            }
            return hmac;
        }

        public string GetRegion(string product, string endpoint, string regionId)
        {
            return GetRegionWithSuffix(product, endpoint, regionId, _endpointSuffix);
        }

        internal string GetRegionWithSuffix(string product, string endpoint, string regionId, string endpointSuffix)
        {
            if (!Runtime.Empty(regionId))
            {
                return regionId;
            }
            string region = "center";
            if (Runtime.Empty(product) || Runtime.Empty(endpoint))
            {
                return region;
            }
            List<string> strs = Runtime.Split(endpoint, ":", null);
            string withoutPort = strs[0];
            string preRegion = HostWithoutDomain(withoutPort, endpointSuffix);
            List<string> nodes = Runtime.Split(preRegion, ".", null);
            if (Runtime.EqualNumber(Runtime.Size(nodes), 2))
            {
                region = nodes[1];
            }
            return region;
        }

        internal string HostWithoutDomain(string host, string suffix)
        {
            string[] domains = new string[] { suffix, "aliyuncs.com", "aliyuncs.alibaba" };
            Dictionary<string, bool> seen = new Dictionary<string, bool>();
            foreach (string domain in domains)
            {
                if (Runtime.Empty(domain) || seen.ContainsKey(domain))
                {
                    continue;
                }
                seen[domain] = true;
                string needle = "." + domain;
                if (!Runtime.Empty(host) && host.EndsWith(needle))
                {
                    return host.Substring(0, host.Length - needle.Length);
                }
            }
            return host;
        }

        public string BuildCanonicalizedResource(Dictionary<string, string> query)
        {
            var encoded = new SortedDictionary<string, string>(StringComparer.Ordinal);
            if (query != null)
            {
                foreach (var entry in query)
                {
                    if (entry.Key != "Signature")
                    {
                        encoded.Add(Runtime.PercentEncode(entry.Key), Runtime.PercentEncode(entry.Value ?? ""));
                    }
                }
            }
            var pairs = new List<string>();
            foreach (var entry in encoded)
            {
                pairs.Add(entry.Key + "=" + entry.Value);
            }
            return string.Join("&", pairs);
        }

        public string BuildCanonicalizedHeaders(Dictionary<string, string> headers)
        {
            // lower header key
            List<string> headersArray = Runtime.KeySet(headers);
            Dictionary<string, string> newHeaders = new Dictionary<string, string>(){};
            foreach (var key in headersArray) {
                string lowerKey = Runtime.ToLower(key);
                string value = headers.Get(key);
                if (!Runtime.IsUnset(value))
                {
                    if (Runtime.IsUnset(newHeaders.Get(lowerKey)) || Runtime.Equals(lowerKey, "host"))
                    {
                        newHeaders[lowerKey] = Runtime.Trim(value);
                    }
                    else
                    {
                        newHeaders[lowerKey] = "" + newHeaders.Get(lowerKey) + "," + Runtime.Trim(value);
                    }
                }
            }
            string canonicalizedHeaders = "";
            List<string> sortedHeaders = GetSignedHeaders(headers);

            foreach (var header in sortedHeaders) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + newHeaders.Get(header) + "\n";
            }
            return canonicalizedHeaders;
        }

        public void ValidateSignedHeaders(List<string> signedHeaders)
        {
            if (!Runtime.ListContains(signedHeaders, "host") || !Runtime.ListContains(signedHeaders, "x-acs-date"))
            {
                throw new TeaException(new Dictionary<string, object>
                {
                    { "code", "InvalidSignedHeaders" },
                    { "message", "signed headers must include host and x-acs-date" }
                });
            }
        }

        public List<string> GetSignedHeaders(Dictionary<string, string> headers)
        {
            List<string> headersArray = Runtime.KeySet(headers);
            List<string> newHeadersArray = new List<string>
            {
            };

            foreach (var key in headersArray) {
                string lowerKey = Runtime.ToLower(key);
                string value = headers.Get(key);
                if (!Runtime.IsUnset(value))
                {
                    Runtime.Append(newHeadersArray, lowerKey);
                }
            }
            List<string> sortedHeadersArray = Runtime.AscSort(newHeadersArray);
            List<string> result = new List<string>
            {
            };

            foreach (var key in sortedHeadersArray) {
                if (Runtime.HasPrefix(key, "x-acs-") || Runtime.Equals(key, "host") || Runtime.Equals(key, "content-type"))
                {
                    if (!Runtime.ListContains(result, key))
                    {
                        Runtime.Append(result, key);
                    }
                }
            }
            return result;
        }

    }
}
