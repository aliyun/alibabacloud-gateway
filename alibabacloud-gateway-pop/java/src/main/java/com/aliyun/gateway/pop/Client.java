// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.pop;

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

            request.headers.put("Authorization", this.getAuthorization(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret));
        }

    }

    public void modifyResponse(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            Object _res = com.aliyun.teautil.Common.readAsJSON(response.body);
            java.util.Map<String, Object> err = com.aliyun.teautil.Common.assertAsMap(_res);
            Object requestId = this.defaultAny(err.get("RequestId"), err.get("requestId"));
            err.put("statusCode", response.statusCode);
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "" + this.defaultAny(err.get("Code"), err.get("code")) + ""),
                new TeaPair("message", "code: " + response.statusCode + ", " + this.defaultAny(err.get("Message"), err.get("message")) + " request id: " + requestId + ""),
                new TeaPair("data", err)
            ));
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

    public String getAuthorization(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, String ak, String secret) throws Exception {
        String signature = this.getSignature(pathname, method, query, headers, signatureAlgorithm, payload, secret);
        java.util.List<String> signedHeaders = this.getSignedHeaders(headers);
        String signedHeadersStr = com.aliyun.darabonba.array.Client.join(signedHeaders, ";");
        return "" + signatureAlgorithm + " Credential=" + ak + ",SignedHeaders=" + signedHeadersStr + ",Signature=" + signature + "";
    }

    public String getSignature(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, String secret) throws Exception {
        String canonicalURI = "/";
        if (!com.aliyun.teautil.Common.empty(pathname)) {
            canonicalURI = pathname;
        }

        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResource(query);
        String canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
        java.util.List<String> signedHeaders = this.getSignedHeaders(headers);
        String signedHeadersStr = com.aliyun.darabonba.array.Client.join(signedHeaders, ";");
        stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload + "";
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

    public String buildCanonicalizedResource(java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryArray = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedQueryArray = com.aliyun.darabonba.array.Client.ascSort(queryArray);
            String separator = "";
            for (String key : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + com.aliyun.darabonba.encode.Encoder.percentEncode(key) + "";
                if (!com.aliyun.teautil.Common.empty(query.get(key))) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + com.aliyun.darabonba.encode.Encoder.percentEncode(query.get(key)) + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeaders(java.util.Map<String, String> headers) throws Exception {
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
