// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.fc;

import com.aliyun.tea.*;
import com.aliyun.gateway.spi.*;
import com.aliyun.gateway.spi.models.*;
import com.aliyun.credentials.*;
import com.aliyun.teautil.*;
import com.aliyun.openapiutil.*;
import com.aliyun.endpointutil.*;
import com.aliyun.darabonba.encode.*;
import com.aliyun.darabonba.signature.*;
import com.aliyun.darabonbastring.*;
import com.aliyun.darabonba.map.*;
import com.aliyun.darabonba.array.*;

public class Client extends com.aliyun.gateway.spi.Client {

    public Client() throws Exception {
        super();
    }


    public void modifyConfiguration(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        config.endpoint = this.getEndpoint(request.productId, config.regionId, config.endpointRule, config.network, config.suffix, config.endpointMap, config.endpoint);
    }

    public void modifyRequest(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        if (!com.aliyun.darabonbastring.Client.hasSuffix(config.endpoint, "aliyuncs.com")) {
            this.signRequestForFc(context);
        } else {
            this.signRequestForPop(context);
        }

    }

    public void modifyResponse(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            if (com.aliyun.darabonbastring.Client.hasPrefix(config.endpoint, "fc.") && com.aliyun.darabonbastring.Client.hasSuffix(config.endpoint, ".aliyuncs.com")) {
                Object popRes = com.aliyun.teautil.Common.readAsJSON(response.body);
                java.util.Map<String, Object> popErr = com.aliyun.teautil.Common.assertAsMap(popRes);
                throw new TeaException(TeaConverter.buildMap(
                    new TeaPair("code", "" + this.defaultAny(popErr.get("Code"), popErr.get("code")) + ""),
                    new TeaPair("message", "code: " + response.statusCode + ", " + this.defaultAny(popErr.get("Message"), popErr.get("message")) + " request id: " + this.defaultAny(popErr.get("RequestId"), popErr.get("requestId")) + ""),
                    new TeaPair("data", popErr)
                ));
            } else {
                java.util.Map<String, Object> _headers = com.aliyun.teautil.Common.assertAsMap(response.headers);
                Object fcRes = com.aliyun.teautil.Common.readAsJSON(response.body);
                java.util.Map<String, Object> fcErr = com.aliyun.teautil.Common.assertAsMap(fcRes);
                throw new TeaException(TeaConverter.buildMap(
                    new TeaPair("code", fcErr.get("ErrorCode")),
                    new TeaPair("message", "code: " + response.statusCode + ", " + fcErr.get("ErrorMessage") + " request id: " + _headers.get("x-fc-request-id") + ""),
                    new TeaPair("data", fcErr)
                ));
            }

        }

        if (com.aliyun.teautil.Common.equalString(request.bodyType, "binary")) {
            response.deserializedBody = response.body;
        } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "byte")) {
            byte[] byt = com.aliyun.teautil.Common.readAsBytes(response.body);
            response.deserializedBody = byt;
        } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "string")) {
            String str = com.aliyun.teautil.Common.readAsString(response.body);
            response.deserializedBody = str;
        } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "json")) {
            Object obj = com.aliyun.teautil.Common.readAsJSON(response.body);
            java.util.Map<String, Object> res = com.aliyun.teautil.Common.assertAsMap(obj);
            response.deserializedBody = res;
        } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "array")) {
            Object arr = com.aliyun.teautil.Common.readAsJSON(response.body);
            response.deserializedBody = arr;
        } else {
            response.deserializedBody = com.aliyun.teautil.Common.readAsString(response.body);
        }

    }

    public String getEndpoint(String productId, String regionId, String endpointRule, String network, String suffix, java.util.Map<String, String> endpointMap, String endpoint) throws Exception {
        if (!com.aliyun.teautil.Common.empty(endpoint)) {
            return endpoint;
        }

        if (!com.aliyun.teautil.Common.isUnset(endpointMap) && !com.aliyun.teautil.Common.empty(endpointMap.get(regionId))) {
            return endpointMap.get(regionId);
        }

        return com.aliyun.endpointutil.Client.getEndpointRules(productId, regionId, endpointRule, network, suffix);
    }

    public Object defaultAny(Object inputValue, Object defaultValue) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(inputValue)) {
            return defaultValue;
        }

        return inputValue;
    }

    public void signRequestForFc(InterceptorContext context) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("host", config.endpoint),
                new TeaPair("date", com.aliyun.teautil.Common.getDateUTCString()),
                new TeaPair("accept", "application/json"),
                new TeaPair("user-agent", request.userAgent)
            ),
            request.headers
        );
        request.headers.put("content-type", "application/json");
        if (!com.aliyun.teautil.Common.isUnset(request.stream)) {
            byte[] tmp = com.aliyun.teautil.Common.readAsBytes(request.stream);
            request.stream = Tea.toReadable(tmp);
            request.headers.put("content-type", "application/octet-stream");
            request.headers.put("content-md5", com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.MD5Sign(com.aliyun.teautil.Common.toString(tmp))));
        } else {
            if (!com.aliyun.teautil.Common.isUnset(request.body)) {
                if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "json")) {
                    String jsonObj = com.aliyun.teautil.Common.toJSONString(request.body);
                    request.stream = Tea.toReadable(jsonObj);
                    request.headers.put("content-type", "application/json");
                    request.headers.put("content-md5", com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.MD5Sign(jsonObj)));
                } else {
                    java.util.Map<String, Object> m = com.aliyun.teautil.Common.assertAsMap(request.body);
                    String formObj = com.aliyun.openapiutil.Client.toForm(m);
                    request.stream = Tea.toReadable(formObj);
                    request.headers.put("content-type", "application/x-www-form-urlencoded");
                    request.headers.put("content-md5", com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.MD5Sign(formObj)));
                }

            }

        }

        com.aliyun.credentials.Client credential = request.credential;
        String accessKeyId = credential.getAccessKeyId();
        String accessKeySecret = credential.getAccessKeySecret();
        String securityToken = credential.getSecurityToken();
        if (!com.aliyun.teautil.Common.empty(securityToken)) {
            request.headers.put("x-fc-security-token", securityToken);
        }

        request.headers.put("Authorization", this.getAuthorizationForFc(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret));
    }

    public void signRequestForPop(InterceptorContext context) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("host", config.endpoint),
                new TeaPair("x-acs-version", request.version),
                new TeaPair("x-acs-action", request.action),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("x-acs-date", com.aliyun.openapiutil.Client.getTimestamp()),
                new TeaPair("x-acs-signature-nonce", com.aliyun.teautil.Common.getNonce()),
                new TeaPair("accept", "application/json")
            ),
            request.headers
        );
        String signatureAlgorithm = "ACS3-HMAC-SHA256";
        if (!com.aliyun.teautil.Common.isUnset(request.signatureAlgorithm)) {
            signatureAlgorithm = request.signatureAlgorithm;
        }

        String hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(""), signatureAlgorithm));
        if (!com.aliyun.teautil.Common.isUnset(request.stream)) {
            byte[] tmp = com.aliyun.teautil.Common.readAsBytes(request.stream);
            hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(tmp, signatureAlgorithm));
            request.stream = Tea.toReadable(tmp);
            request.headers.put("content-type", "application/octet-stream");
        } else {
            if (!com.aliyun.teautil.Common.isUnset(request.body)) {
                if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "json")) {
                    String jsonObj = com.aliyun.teautil.Common.toJSONString(request.body);
                    hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(jsonObj), signatureAlgorithm));
                    request.stream = Tea.toReadable(jsonObj);
                    request.headers.put("content-type", "application/json; charset=utf-8");
                } else {
                    java.util.Map<String, Object> m = com.aliyun.teautil.Common.assertAsMap(request.body);
                    String formObj = com.aliyun.openapiutil.Client.toForm(m);
                    hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(formObj), signatureAlgorithm));
                    request.stream = Tea.toReadable(formObj);
                    request.headers.put("content-type", "application/x-www-form-urlencoded");
                }

            }

        }

        request.headers.put("x-acs-content-sha256", hashedRequestPayload);
        if (!com.aliyun.teautil.Common.equalString(request.authType, "Anonymous")) {
            com.aliyun.credentials.Client credential = request.credential;
            String accessKeyId = credential.getAccessKeyId();
            String accessKeySecret = credential.getAccessKeySecret();
            String securityToken = credential.getSecurityToken();
            if (!com.aliyun.teautil.Common.empty(securityToken)) {
                request.headers.put("x-acs-accesskey-id", accessKeyId);
                request.headers.put("x-acs-security-token", securityToken);
            }

            request.headers.put("Authorization", this.getAuthorizationForPop(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret));
        }

    }

    public String getAuthorizationForFc(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String ak, String secret) throws Exception {
        String sign = this.getSignatureForFc(pathname, method, query, headers, secret);
        return "FC " + ak + ":" + sign + "";
    }

    public String getSignatureForFc(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String secret) throws Exception {
        String resource = pathname;
        String contentMd5 = headers.get("content-md5");
        if (com.aliyun.teautil.Common.isUnset(contentMd5)) {
            contentMd5 = "";
        }

        String contentType = headers.get("content-type");
        if (com.aliyun.teautil.Common.isUnset(contentType)) {
            contentType = "";
        }

        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResourceForFc(resource, query);
        String canonicalizedHeaders = this.buildCanonicalizedHeadersForFc(headers);
        stringToSign = "" + method + "\n" + contentMd5 + "\n" + contentType + "\n" + headers.get("date") + "\n" + canonicalizedHeaders + "" + canonicalizedResource + "";
        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(stringToSign, secret));
    }

    public String buildCanonicalizedResourceForFc(String pathname, java.util.Map<String, String> query) throws Exception {
        java.util.List<String> paths = com.aliyun.darabonbastring.Client.split(pathname, "?", 2);
        String canonicalizedResource = paths.get(0);
        java.util.List<String> resources = new java.util.ArrayList<>();
        if (com.aliyun.teautil.Common.equalNumber(com.aliyun.darabonba.array.Client.size(paths), 2)) {
            resources = com.aliyun.darabonbastring.Client.split(paths.get(1), "&", 0);
        }

        java.util.List<String> subResources = new java.util.ArrayList<>();
        String tmp = "";
        String separator = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryList = com.aliyun.darabonba.map.Client.keySet(query);
            for (String paramName : queryList) {
                tmp = "" + tmp + "" + separator + "" + paramName + "";
                if (!com.aliyun.teautil.Common.isUnset(query.get(paramName))) {
                    tmp = "" + tmp + "=" + query.get(paramName) + "";
                }

                separator = ";";
            }
            subResources = com.aliyun.darabonbastring.Client.split(tmp, ";", 0);
        }

        java.util.List<String> result = com.aliyun.darabonba.array.Client.concat(subResources, resources);
        java.util.List<String> sortedParams = com.aliyun.darabonba.array.Client.ascSort(result);
        if (com.aliyun.teautil.Common.equalNumber(com.aliyun.darabonba.array.Client.size(sortedParams), 0)) {
            return "" + canonicalizedResource + "\n";
        }

        return "" + canonicalizedResource + "\n" + com.aliyun.darabonba.array.Client.join(sortedParams, "\n") + "";
    }

    public String buildCanonicalizedHeadersForFc(java.util.Map<String, String> headers) throws Exception {
        String canonicalizedHeaders = "";
        java.util.List<String> keys = com.aliyun.darabonba.map.Client.keySet(headers);
        java.util.List<String> sortedHeaders = com.aliyun.darabonba.array.Client.ascSort(keys);
        for (String header : sortedHeaders) {
            if (com.aliyun.darabonbastring.Client.contains(com.aliyun.darabonbastring.Client.toLower(header), "x-fc-")) {
                canonicalizedHeaders = "" + canonicalizedHeaders + "" + com.aliyun.darabonbastring.Client.toLower(header) + ":" + headers.get(header) + "\n";
            }

        }
        return canonicalizedHeaders;
    }

    public String getAuthorizationForPop(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, String ak, String secret) throws Exception {
        String signature = this.getSignatureForPop(pathname, method, query, headers, signatureAlgorithm, payload, secret);
        return "" + signatureAlgorithm + "  Credential=" + ak + ",SignedHeaders=" + com.aliyun.darabonba.array.Client.join(this.getSignedHeaders(headers), ";") + ",Signature=" + signature + "";
    }

    public String getSignatureForPop(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, String secret) throws Exception {
        String canonicalURI = "/";
        if (!com.aliyun.teautil.Common.empty(pathname)) {
            canonicalURI = pathname;
        }

        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResourceForPop(query);
        String canonicalizedHeaders = this.buildCanonicalizedHeadersForPop(headers);
        java.util.List<String> signedHeaders = this.getSignedHeaders(headers);
        stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + com.aliyun.darabonba.array.Client.join(signedHeaders, ";") + "\n" + payload + "";
        String hex = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(stringToSign), signatureAlgorithm));
        stringToSign = "" + signatureAlgorithm + "\n" + hex + "";
        byte[] signature = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.darabonbastring.Client.equals(signatureAlgorithm, "ACS3-HMAC-SHA256")) {
            signature = com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(stringToSign, secret);
        } else if (com.aliyun.darabonbastring.Client.equals(signatureAlgorithm, "ACS3-HMAC-SM3")) {
            signature = com.aliyun.darabonba.signature.Signer.HmacSM3Sign(stringToSign, secret);
        } else if (com.aliyun.darabonbastring.Client.equals(signatureAlgorithm, "ACS3-RSA-SHA256")) {
            signature = com.aliyun.darabonba.signature.Signer.SHA256withRSASign(stringToSign, secret);
        }

        return com.aliyun.darabonba.encode.Encoder.hexEncode(signature);
    }

    public String buildCanonicalizedResourceForPop(java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryArray = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedQueryArray = com.aliyun.darabonba.array.Client.ascSort(queryArray);
            for (String key : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "&" + com.aliyun.darabonba.encode.Encoder.percentEncode(key) + "";
                if (!com.aliyun.teautil.Common.empty(query.get(key))) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + com.aliyun.darabonba.encode.Encoder.percentEncode(query.get(key)) + "";
                }

            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeadersForPop(java.util.Map<String, String> headers) throws Exception {
        String canonicalizedHeaders = "";
        java.util.List<String> sortedHeaders = this.getSignedHeaders(headers);
        for (String header : sortedHeaders) {
            canonicalizedHeaders = "" + canonicalizedHeaders + "" + header + ":" + com.aliyun.darabonbastring.Client.trim(headers.get(header)) + "\n";
        }
        return canonicalizedHeaders;
    }

    public java.util.List<String> getSignedHeaders(java.util.Map<String, String> headers) throws Exception {
        java.util.List<String> headersArray = com.aliyun.darabonba.map.Client.keySet(headers);
        java.util.List<String> sortedHeadersArray = com.aliyun.darabonba.array.Client.ascSort(headersArray);
        String tmp = "";
        String separator = "";
        for (String key : sortedHeadersArray) {
            String lowerKey = com.aliyun.darabonbastring.Client.toLower(key);
            if (com.aliyun.darabonbastring.Client.hasPrefix(lowerKey, "x-acs-") || com.aliyun.darabonbastring.Client.equals(lowerKey, "host") || com.aliyun.darabonbastring.Client.equals(lowerKey, "content-type")) {
                if (!com.aliyun.darabonbastring.Client.contains(tmp, lowerKey)) {
                    tmp = "" + tmp + "" + separator + "" + lowerKey + "";
                    separator = ";";
                }

            }

        }
        return com.aliyun.darabonbastring.Client.split(tmp, ";", 0);
    }
}
