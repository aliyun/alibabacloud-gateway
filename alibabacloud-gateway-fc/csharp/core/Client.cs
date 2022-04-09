// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;

using AlibabaCloud.GatewayFc.Models;

namespace AlibabaCloud.GatewayFc
{
    public class Client : AlibabaCloud.GatewaySpi.Client
    {

        public Client(): base()
        {
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
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            if (!AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(config.Endpoint, "aliyuncs.com"))
            {
                SignRequestForFc(context);
            }
            else
            {
                SignRequestForPop(context);
            }
        }

        public async Task ModifyRequestAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            if (!AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(config.Endpoint, "aliyuncs.com"))
            {
                await SignRequestForFcAsync(context);
            }
            else
            {
                await SignRequestForPopAsync(context);
            }
        }

        public void ModifyResponse(AlibabaCloud.GatewaySpi.Models.InterceptorContext context, AlibabaCloud.GatewaySpi.Models.AttributeMap attributeMap)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(config.Endpoint, "fc.") && AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(config.Endpoint, ".aliyuncs.com"))
                {
                    object popRes = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> popErr = AlibabaCloud.TeaUtil.Common.AssertAsMap(popRes);
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "" + DefaultAny(popErr.Get("Code"), popErr.Get("code"))},
                        {"message", "code: " + response.StatusCode + ", " + DefaultAny(popErr.Get("Message"), popErr.Get("message")) + " request id: " + DefaultAny(popErr.Get("RequestId"), popErr.Get("requestId"))},
                        {"data", popErr},
                    });
                }
                else
                {
                    Dictionary<string, object> _headers = AlibabaCloud.TeaUtil.Common.AssertAsMap(response.Headers);
                    object fcRes = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> fcErr = AlibabaCloud.TeaUtil.Common.AssertAsMap(fcRes);
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", fcErr.Get("ErrorCode")},
                        {"message", "code: " + response.StatusCode + ", " + fcErr.Get("ErrorMessage") + " request id: " + _headers.Get("x-fc-request-id")},
                        {"data", fcErr},
                    });
                }
            }
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
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextResponse response = context.Response;
            if (AlibabaCloud.TeaUtil.Common.Is4xx(response.StatusCode) || AlibabaCloud.TeaUtil.Common.Is5xx(response.StatusCode))
            {
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(config.Endpoint, "fc.") && AlibabaCloud.DarabonbaString.StringUtil.HasSuffix(config.Endpoint, ".aliyuncs.com"))
                {
                    object popRes = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> popErr = AlibabaCloud.TeaUtil.Common.AssertAsMap(popRes);
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", "" + DefaultAny(popErr.Get("Code"), popErr.Get("code"))},
                        {"message", "code: " + response.StatusCode + ", " + DefaultAny(popErr.Get("Message"), popErr.Get("message")) + " request id: " + DefaultAny(popErr.Get("RequestId"), popErr.Get("requestId"))},
                        {"data", popErr},
                    });
                }
                else
                {
                    Dictionary<string, object> _headers = AlibabaCloud.TeaUtil.Common.AssertAsMap(response.Headers);
                    object fcRes = AlibabaCloud.TeaUtil.Common.ReadAsJSON(response.Body);
                    Dictionary<string, object> fcErr = AlibabaCloud.TeaUtil.Common.AssertAsMap(fcRes);
                    throw new TeaException(new Dictionary<string, object>
                    {
                        {"code", fcErr.Get("ErrorCode")},
                        {"message", "code: " + response.StatusCode + ", " + fcErr.Get("ErrorMessage") + " request id: " + _headers.Get("x-fc-request-id")},
                        {"data", fcErr},
                    });
                }
            }
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

        public void SignRequestForFc(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"accept", "application/json"},
                    {"user-agent", request.UserAgent},
                },
                request.Headers
            );
            request.Headers["content-type"] = "application/json";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Stream))
            {
                byte[] tmp = AlibabaCloud.TeaUtil.Common.ReadAsBytes(request.Stream);
                request.Stream = TeaCore.BytesReadable(tmp);
                request.Headers["content-type"] = "application/octet-stream";
                request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5SignForBytes(tmp));
            }
            else
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
                {
                    if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "json"))
                    {
                        string jsonObj = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                        request.Stream = TeaCore.BytesReadable(jsonObj);
                        request.Headers["content-type"] = "application/json";
                        request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(jsonObj));
                    }
                    else
                    {
                        Dictionary<string, object> m = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                        string formObj = AlibabaCloud.OpenApiUtil.Client.ToForm(m);
                        request.Stream = TeaCore.BytesReadable(formObj);
                        request.Headers["content-type"] = "application/x-www-form-urlencoded";
                        request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(formObj));
                    }
                }
            }
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = credential.GetAccessKeyId();
            string accessKeySecret = credential.GetAccessKeySecret();
            string securityToken = credential.GetSecurityToken();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-fc-security-token"] = securityToken;
            }
            request.Headers["Authorization"] = GetAuthorizationForFc(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public async Task SignRequestForFcAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"date", AlibabaCloud.TeaUtil.Common.GetDateUTCString()},
                    {"accept", "application/json"},
                    {"user-agent", request.UserAgent},
                },
                request.Headers
            );
            request.Headers["content-type"] = "application/json";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Stream))
            {
                byte[] tmp = AlibabaCloud.TeaUtil.Common.ReadAsBytes(request.Stream);
                request.Stream = TeaCore.BytesReadable(tmp);
                request.Headers["content-type"] = "application/octet-stream";
                request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5SignForBytes(tmp));
            }
            else
            {
                if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
                {
                    if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "json"))
                    {
                        string jsonObj = AlibabaCloud.TeaUtil.Common.ToJSONString(request.Body);
                        request.Stream = TeaCore.BytesReadable(jsonObj);
                        request.Headers["content-type"] = "application/json";
                        request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(jsonObj));
                    }
                    else
                    {
                        Dictionary<string, object> m = AlibabaCloud.TeaUtil.Common.AssertAsMap(request.Body);
                        string formObj = AlibabaCloud.OpenApiUtil.Client.ToForm(m);
                        request.Stream = TeaCore.BytesReadable(formObj);
                        request.Headers["content-type"] = "application/x-www-form-urlencoded";
                        request.Headers["content-md5"] = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.MD5Sign(formObj));
                    }
                }
            }
            Aliyun.Credentials.Client credential = request.Credential;
            string accessKeyId = await credential.GetAccessKeyIdAsync();
            string accessKeySecret = await credential.GetAccessKeySecretAsync();
            string securityToken = await credential.GetSecurityTokenAsync();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                request.Headers["x-fc-security-token"] = securityToken;
            }
            request.Headers["Authorization"] = await GetAuthorizationForFcAsync(request.Pathname, request.Method, request.Query, request.Headers, accessKeyId, accessKeySecret);
        }

        public void SignRequestForPop(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-date", AlibabaCloud.OpenApiUtil.Client.GetTimestamp()},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            string signatureAlgorithm = "ACS3-HMAC-SHA256";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.SignatureAlgorithm))
            {
                signatureAlgorithm = request.SignatureAlgorithm;
            }
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
            request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
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
                request.Headers["Authorization"] = GetAuthorizationForPop(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret);
            }
        }

        public async Task SignRequestForPopAsync(AlibabaCloud.GatewaySpi.Models.InterceptorContext context)
        {
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextRequest request = context.Request;
            AlibabaCloud.GatewaySpi.Models.InterceptorContext.InterceptorContextConfiguration config = context.Configuration;
            request.Headers = TeaConverter.merge<string>
            (
                new Dictionary<string, string>()
                {
                    {"host", config.Endpoint},
                    {"x-acs-version", request.Version},
                    {"x-acs-action", request.Action},
                    {"user-agent", request.UserAgent},
                    {"x-acs-date", AlibabaCloud.OpenApiUtil.Client.GetTimestamp()},
                    {"x-acs-signature-nonce", AlibabaCloud.TeaUtil.Common.GetNonce()},
                    {"accept", "application/json"},
                },
                request.Headers
            );
            string signatureAlgorithm = "ACS3-HMAC-SHA256";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.SignatureAlgorithm))
            {
                signatureAlgorithm = request.SignatureAlgorithm;
            }
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
            request.Headers["x-acs-content-sha256"] = hashedRequestPayload;
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
                request.Headers["Authorization"] = await GetAuthorizationForPopAsync(request.Pathname, request.Method, request.Query, request.Headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret);
            }
        }

        public string GetAuthorizationForFc(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = GetSignatureForFc(pathname, method, query, headers, secret);
            return "FC " + ak + ":" + sign;
        }

        public async Task<string> GetAuthorizationForFcAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string ak, string secret)
        {
            string sign = await GetSignatureForFcAsync(pathname, method, query, headers, secret);
            return "FC " + ak + ":" + sign;
        }

        public string GetSignatureForFc(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = pathname;
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
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResourceForFc(resource, query);
            string canonicalizedHeaders = BuildCanonicalizedHeadersForFc(headers);
            stringToSign = "" + method + "\n" + contentMd5 + "\n" + contentType + "\n" + headers.Get("date") + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret));
        }

        public async Task<string> GetSignatureForFcAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string secret)
        {
            string resource = pathname;
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
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceForFcAsync(resource, query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersForFcAsync(headers);
            stringToSign = "" + method + "\n" + contentMd5 + "\n" + contentType + "\n" + headers.Get("date") + "\n" + canonicalizedHeaders + canonicalizedResource;
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret));
        }

        public string BuildCanonicalizedResourceForFc(string pathname, Dictionary<string, string> query)
        {
            List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(pathname, "?", 2);
            string canonicalizedResource = paths[0];
            List<string> resources = new List<string>
            {
            };
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
            {
                resources = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", 0);
            }
            List<string> subResources = new List<string>
            {
            };
            string tmp = "";
            string separator = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryList = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);

                foreach (var paramName in queryList) {
                    tmp = "" + tmp + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(paramName)))
                    {
                        tmp = "" + tmp + "=" + query.Get(paramName);
                    }
                    separator = ";";
                }
                subResources = AlibabaCloud.DarabonbaString.StringUtil.Split(tmp, ";", 0);
            }
            List<string> result = AlibabaCloud.DarabonbaArray.ArrayUtil.Concat(subResources, resources);
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(result);
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(sortedParams), 0))
            {
                return "" + canonicalizedResource + "\n";
            }
            return "" + canonicalizedResource + "\n" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(sortedParams, "\n");
        }

        public async Task<string> BuildCanonicalizedResourceForFcAsync(string pathname, Dictionary<string, string> query)
        {
            List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(pathname, "?", 2);
            string canonicalizedResource = paths[0];
            List<string> resources = new List<string>
            {
            };
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
            {
                resources = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", 0);
            }
            List<string> subResources = new List<string>
            {
            };
            string tmp = "";
            string separator = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryList = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);

                foreach (var paramName in queryList) {
                    tmp = "" + tmp + separator + paramName;
                    if (!AlibabaCloud.TeaUtil.Common.IsUnset(query.Get(paramName)))
                    {
                        tmp = "" + tmp + "=" + query.Get(paramName);
                    }
                    separator = ";";
                }
                subResources = AlibabaCloud.DarabonbaString.StringUtil.Split(tmp, ";", 0);
            }
            List<string> result = AlibabaCloud.DarabonbaArray.ArrayUtil.Concat(subResources, resources);
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(result);
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(sortedParams), 0))
            {
                return "" + canonicalizedResource + "\n";
            }
            return "" + canonicalizedResource + "\n" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(sortedParams, "\n");
        }

        public string BuildCanonicalizedHeadersForFc(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-fc-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + AlibabaCloud.DarabonbaString.StringUtil.ToLower(header) + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersForFcAsync(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-fc-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + AlibabaCloud.DarabonbaString.StringUtil.ToLower(header) + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public string GetAuthorizationForPop(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, string secret)
        {
            string signature = GetSignatureForPop(pathname, method, query, headers, signatureAlgorithm, payload, secret);
            return "" + signatureAlgorithm + "  Credential=" + ak + ",SignedHeaders=" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(GetSignedHeaders(headers), ";") + ",Signature=" + signature;
        }

        public async Task<string> GetAuthorizationForPopAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string ak, string secret)
        {
            string signature = await GetSignatureForPopAsync(pathname, method, query, headers, signatureAlgorithm, payload, secret);
            return "" + signatureAlgorithm + "  Credential=" + ak + ",SignedHeaders=" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(await GetSignedHeadersAsync(headers), ";") + ",Signature=" + signature;
        }

        public string GetSignatureForPop(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string secret)
        {
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                canonicalURI = pathname;
            }
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResourceForPop(query);
            string canonicalizedHeaders = BuildCanonicalizedHeadersForPop(headers);
            List<string> signedHeaders = GetSignedHeaders(headers);
            stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";") + "\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureAlgorithm, "ACS3-HMAC-SHA256"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret);
            }
            else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureAlgorithm, "ACS3-HMAC-SM3"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3Sign(stringToSign, secret);
            }
            else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureAlgorithm, "ACS3-RSA-SHA256"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.SHA256withRSASign(stringToSign, secret);
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public async Task<string> GetSignatureForPopAsync(string pathname, string method, Dictionary<string, string> query, Dictionary<string, string> headers, string signatureAlgorithm, string payload, string secret)
        {
            string canonicalURI = "/";
            if (!AlibabaCloud.TeaUtil.Common.Empty(pathname))
            {
                canonicalURI = pathname;
            }
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceForPopAsync(query);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersForPopAsync(headers);
            List<string> signedHeaders = await GetSignedHeadersAsync(headers);
            stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(signedHeaders, ";") + "\n" + payload;
            string hex = AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(AlibabaCloud.DarabonbaEncodeUtil.Encoder.Hash(AlibabaCloud.TeaUtil.Common.ToBytes(stringToSign), signatureAlgorithm));
            stringToSign = "" + signatureAlgorithm + "\n" + hex;
            byte[] signature = AlibabaCloud.TeaUtil.Common.ToBytes("");
            if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureAlgorithm, "ACS3-HMAC-SHA256"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, secret);
            }
            else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureAlgorithm, "ACS3-HMAC-SM3"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSM3Sign(stringToSign, secret);
            }
            else if (AlibabaCloud.DarabonbaString.StringUtil.Equals(signatureAlgorithm, "ACS3-RSA-SHA256"))
            {
                signature = AlibabaCloud.DarabonbaSignatureUtil.Signer.SHA256withRSASign(stringToSign, secret);
            }
            return AlibabaCloud.DarabonbaEncodeUtil.Encoder.HexEncode(signature);
        }

        public string BuildCanonicalizedResourceForPop(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + "&" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(key);
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(key));
                    }
                }
            }
            return canonicalizedResource;
        }

        public async Task<string> BuildCanonicalizedResourceForPopAsync(Dictionary<string, string> query)
        {
            string canonicalizedResource = "";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(query))
            {
                List<string> queryArray = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(query);
                List<string> sortedQueryArray = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(queryArray);

                foreach (var key in sortedQueryArray) {
                    canonicalizedResource = "" + canonicalizedResource + "&" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(key);
                    if (!AlibabaCloud.TeaUtil.Common.Empty(query.Get(key)))
                    {
                        canonicalizedResource = "" + canonicalizedResource + "=" + AlibabaCloud.DarabonbaEncodeUtil.Encoder.PercentEncode(query.Get(key));
                    }
                }
            }
            return canonicalizedResource;
        }

        public string BuildCanonicalizedHeadersForPop(Dictionary<string, string> headers)
        {
            string canonicalizedHeaders = "";
            List<string> sortedHeaders = GetSignedHeaders(headers);

            foreach (var header in sortedHeaders) {
                canonicalizedHeaders = "" + canonicalizedHeaders + header + ":" + AlibabaCloud.DarabonbaString.StringUtil.Trim(headers.Get(header)) + "\n";
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersForPopAsync(Dictionary<string, string> headers)
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
                if (AlibabaCloud.DarabonbaString.StringUtil.HasPrefix(lowerKey, "x-acs-") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "host") || AlibabaCloud.DarabonbaString.StringUtil.Equals(lowerKey, "content-type"))
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

        public Dictionary<string, object> SignRequest(HttpRequest request, Aliyun.Credentials.Client credential)
        {
            HttpRequest httpRequest = new HttpRequest
            {
                Method = request.Method,
                Path = request.Path,
                Headers = request.Headers,
                Body = request.Body,
                ReqBodyType = request.ReqBodyType,
            };
            httpRequest.Headers["date"] = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            httpRequest.Headers["accept"] = "application/json";
            httpRequest.Headers["content-type"] = "application/json";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "json"))
                {
                    httpRequest.Headers["content-type"] = "application/json";
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "form"))
                {
                    httpRequest.Headers["content-type"] = "application/x-www-form-urlencoded";
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "binary"))
                {
                    httpRequest.Headers["content-type"] = "application/octet-stream";
                }
            }
            string accessKeyId = credential.GetAccessKeyId();
            string accessKeySecret = credential.GetAccessKeySecret();
            string securityToken = credential.GetSecurityToken();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                httpRequest.Headers["x-fc-security-token"] = securityToken;
            }
            string resource = request.Path;
            object contentMd5 = httpRequest.Headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            object contentType = httpRequest.Headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string stringToSign = "";
            string canonicalizedResource = BuildCanonicalizedResource(resource);
            string canonicalizedHeaders = BuildCanonicalizedHeaders(httpRequest.Headers);
            stringToSign = "" + request.Method + "\n" + contentMd5 + "\n" + contentType + "\n" + httpRequest.Headers.Get("date") + "\n" + canonicalizedHeaders + canonicalizedResource;
            string signature = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, accessKeySecret));
            httpRequest.Headers["Authorization"] = "FC " + accessKeyId + ":" + signature;
            return httpRequest.Headers;
        }

        public async Task<Dictionary<string, object>> SignRequestAsync(HttpRequest request, Aliyun.Credentials.Client credential)
        {
            HttpRequest httpRequest = new HttpRequest
            {
                Method = request.Method,
                Path = request.Path,
                Headers = request.Headers,
                Body = request.Body,
                ReqBodyType = request.ReqBodyType,
            };
            httpRequest.Headers["date"] = AlibabaCloud.TeaUtil.Common.GetDateUTCString();
            httpRequest.Headers["accept"] = "application/json";
            httpRequest.Headers["content-type"] = "application/json";
            if (!AlibabaCloud.TeaUtil.Common.IsUnset(request.Body))
            {
                if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "json"))
                {
                    httpRequest.Headers["content-type"] = "application/json";
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "form"))
                {
                    httpRequest.Headers["content-type"] = "application/x-www-form-urlencoded";
                }
                else if (AlibabaCloud.TeaUtil.Common.EqualString(request.ReqBodyType, "binary"))
                {
                    httpRequest.Headers["content-type"] = "application/octet-stream";
                }
            }
            string accessKeyId = await credential.GetAccessKeyIdAsync();
            string accessKeySecret = await credential.GetAccessKeySecretAsync();
            string securityToken = await credential.GetSecurityTokenAsync();
            if (!AlibabaCloud.TeaUtil.Common.Empty(securityToken))
            {
                httpRequest.Headers["x-fc-security-token"] = securityToken;
            }
            string resource = request.Path;
            object contentMd5 = httpRequest.Headers.Get("content-md5");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentMd5))
            {
                contentMd5 = "";
            }
            object contentType = httpRequest.Headers.Get("content-type");
            if (AlibabaCloud.TeaUtil.Common.IsUnset(contentType))
            {
                contentType = "";
            }
            string stringToSign = "";
            string canonicalizedResource = await BuildCanonicalizedResourceAsync(resource);
            string canonicalizedHeaders = await BuildCanonicalizedHeadersAsync(httpRequest.Headers);
            stringToSign = "" + request.Method + "\n" + contentMd5 + "\n" + contentType + "\n" + httpRequest.Headers.Get("date") + "\n" + canonicalizedHeaders + canonicalizedResource;
            string signature = AlibabaCloud.DarabonbaEncodeUtil.Encoder.Base64EncodeToString(AlibabaCloud.DarabonbaSignatureUtil.Signer.HmacSHA256Sign(stringToSign, accessKeySecret));
            httpRequest.Headers["Authorization"] = "FC " + accessKeyId + ":" + signature;
            return httpRequest.Headers;
        }

        public string BuildCanonicalizedResource(string pathname)
        {
            List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(pathname, "?", 2);
            string canonicalizedResource = paths[0];
            List<string> resources = new List<string>
            {
            };
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
            {
                resources = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", 0);
            }
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(resources);
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(sortedParams), 0))
            {
                return "" + canonicalizedResource + "\n";
            }
            return "" + canonicalizedResource + "\n" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(sortedParams, "\n");
        }

        public async Task<string> BuildCanonicalizedResourceAsync(string pathname)
        {
            List<string> paths = AlibabaCloud.DarabonbaString.StringUtil.Split(pathname, "?", 2);
            string canonicalizedResource = paths[0];
            List<string> resources = new List<string>
            {
            };
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(paths), 2))
            {
                resources = AlibabaCloud.DarabonbaString.StringUtil.Split(paths[1], "&", 0);
            }
            List<string> sortedParams = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(resources);
            if (AlibabaCloud.TeaUtil.Common.EqualNumber(AlibabaCloud.DarabonbaArray.ArrayUtil.Size(sortedParams), 0))
            {
                return "" + canonicalizedResource + "\n";
            }
            return "" + canonicalizedResource + "\n" + AlibabaCloud.DarabonbaArray.ArrayUtil.Join(sortedParams, "\n");
        }

        public string BuildCanonicalizedHeaders(Dictionary<string, object> headers)
        {
            string canonicalizedHeaders = "";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-fc-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + AlibabaCloud.DarabonbaString.StringUtil.ToLower(header) + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

        public async Task<string> BuildCanonicalizedHeadersAsync(Dictionary<string, object> headers)
        {
            string canonicalizedHeaders = "";
            List<string> keys = AlibabaCloud.DarabonbaMap.MapUtil.KeySet(headers);
            List<string> sortedHeaders = AlibabaCloud.DarabonbaArray.ArrayUtil.AscSort(keys);

            foreach (var header in sortedHeaders) {
                if (AlibabaCloud.DarabonbaString.StringUtil.Contains(AlibabaCloud.DarabonbaString.StringUtil.ToLower(header), "x-fc-"))
                {
                    canonicalizedHeaders = "" + canonicalizedHeaders + AlibabaCloud.DarabonbaString.StringUtil.ToLower(header) + ":" + headers.Get(header) + "\n";
                }
            }
            return canonicalizedHeaders;
        }

    }
}
