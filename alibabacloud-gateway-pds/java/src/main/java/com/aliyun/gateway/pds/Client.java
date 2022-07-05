// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.pds;

import com.aliyun.tea.*;
import com.aliyun.gateway.spi.*;
import com.aliyun.gateway.spi.models.*;
import com.aliyun.credentials.*;
import com.aliyun.teautil.*;
import com.aliyun.openapiutil.*;
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
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        config.endpoint = this.getEndpoint(config.network, config.endpoint);
    }

    public void modifyRequest(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("date", com.aliyun.teautil.Common.getDateUTCString()),
                new TeaPair("host", config.endpoint),
                new TeaPair("x-acs-version", request.version),
                new TeaPair("x-acs-action", request.action),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("x-acs-signature-nonce", com.aliyun.teautil.Common.getNonce()),
                new TeaPair("x-acs-signature-method", "HMAC-SHA1"),
                new TeaPair("x-acs-signature-version", "1.0"),
                new TeaPair("accept", "application/json")
            ),
            request.headers
        );
        if (!com.aliyun.teautil.Common.isUnset(request.stream)) {
            byte[] tmp = com.aliyun.teautil.Common.readAsBytes(request.stream);
            request.stream = Tea.toReadable(tmp);
            request.headers.put("content-type", "application/octet-stream");
        } else {
            if (!com.aliyun.teautil.Common.isUnset(request.body)) {
                if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "json")) {
                    String jsonObj = com.aliyun.teautil.Common.toJSONString(request.body);
                    request.stream = Tea.toReadable(jsonObj);
                    request.headers.put("content-type", "application/json; charset=utf-8");
                } else {
                    java.util.Map<String, Object> m = com.aliyun.teautil.Common.assertAsMap(request.body);
                    String formObj = com.aliyun.openapiutil.Client.toForm(m);
                    request.stream = Tea.toReadable(formObj);
                    request.headers.put("content-type", "application/x-www-form-urlencoded");
                }

            }

        }

        if (!com.aliyun.teautil.Common.equalString(request.authType, "Anonymous")) {
            com.aliyun.credentials.Client credential = request.credential;
            String accessKeyId = credential.getAccessKeyId();
            String accessKeySecret = credential.getAccessKeySecret();
            String securityToken = credential.getSecurityToken();
            String bearerToken = credential.getBearerToken();
            if (!com.aliyun.teautil.Common.empty(bearerToken)) {
                request.headers.put("x-acs-bearer-token", bearerToken);
                request.headers.put("Authorization", "Bearer " + bearerToken + "");
            } else {
                if (!com.aliyun.teautil.Common.empty(securityToken)) {
                    request.headers.put("x-acs-security-token", securityToken);
                }

                request.headers.put("Authorization", this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret));
            }

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

        if (!com.aliyun.teautil.Common.isUnset(response.body) && !com.aliyun.teautil.Common.equalNumber(response.statusCode, 204)) {
            if (com.aliyun.teautil.Common.equalString(request.bodyType, "binary")) {
                response.deserializedBody = response.body;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "byte")) {
                byte[] byt = com.aliyun.teautil.Common.readAsBytes(response.body);
                response.deserializedBody = byt;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "string")) {
                String str = com.aliyun.teautil.Common.readAsString(response.body);
                response.deserializedBody = str;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "json")) {
                response.deserializedBody = com.aliyun.teautil.Common.readAsJSON(response.body);
            } else {
                response.deserializedBody = com.aliyun.teautil.Common.readAsString(response.body);
            }

        }

    }

    public String getEndpoint(String network, String endpoint) throws Exception {
        String realEndpoint = "api.aliyunpds.com";
        if (!com.aliyun.teautil.Common.empty(endpoint)) {
            realEndpoint = endpoint;
        }

        if (!com.aliyun.teautil.Common.empty(network) && com.aliyun.darabonbastring.Client.equals(network, "vpc")) {
            realEndpoint = com.aliyun.darabonbastring.Client.replace(realEndpoint, "api.aliyunpds.com", "api-vpc.aliyunpds.com", null);
            realEndpoint = com.aliyun.darabonbastring.Client.replace(realEndpoint, "admin.aliyunpds.com", "admin-vpc.aliyunpds.com", null);
        }

        return realEndpoint;
    }

    public Object defaultAny(Object inputValue, Object defaultValue) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(inputValue)) {
            return defaultValue;
        }

        return inputValue;
    }

    public String getAuthorization(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String ak, String secret) throws Exception {
        String signature = this.getSignature(pathname, method, query, headers, secret);
        return "acs " + ak + ":" + signature + "";
    }

    public String getSignature(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String secret) throws Exception {
        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResource(pathname, query);
        String canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
        stringToSign = "" + method + "\n" + canonicalizedHeaders + "" + canonicalizedResource + "";
        byte[] signature = com.aliyun.darabonba.signature.Signer.HmacSHA1Sign(stringToSign, secret);
        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(signature);
    }

    public String buildCanonicalizedResource(String pathname, java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = pathname;
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryArray = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedQueryArray = com.aliyun.darabonba.array.Client.ascSort(queryArray);
            String separator = "?";
            for (String key : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + key + "";
                if (!com.aliyun.teautil.Common.empty(query.get(key))) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + query.get(key) + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeaders(java.util.Map<String, String> headers) throws Exception {
        String accept = headers.get("accept");
        if (com.aliyun.teautil.Common.isUnset(accept)) {
            accept = "";
        }

        String contentMd5 = headers.get("content-md5");
        if (com.aliyun.teautil.Common.isUnset(contentMd5)) {
            contentMd5 = "";
        }

        String contentType = headers.get("content-type");
        if (com.aliyun.teautil.Common.isUnset(contentType)) {
            contentType = "";
        }

        String date = headers.get("date");
        if (com.aliyun.teautil.Common.isUnset(date)) {
            date = "";
        }

        String canonicalizedHeaders = "" + accept + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n";
        java.util.List<String> sortedHeaders = this.getSignedHeaders(headers);
        for (String header : sortedHeaders) {
            String value = headers.get(header);
            String valueTrim = com.aliyun.darabonbastring.Client.trim(value);
            canonicalizedHeaders = "" + canonicalizedHeaders + "" + header + ":" + valueTrim + "\n";
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
            if (com.aliyun.darabonbastring.Client.hasPrefix(lowerKey, "x-acs-")) {
                if (!com.aliyun.darabonbastring.Client.contains(tmp, lowerKey)) {
                    tmp = "" + tmp + "" + separator + "" + lowerKey + "";
                    separator = ";";
                }

            }

        }
        return com.aliyun.darabonbastring.Client.split(tmp, ";", 0);
    }
}
