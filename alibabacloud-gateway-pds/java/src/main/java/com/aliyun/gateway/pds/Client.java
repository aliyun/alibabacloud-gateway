// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.pds;

import com.aliyun.tea.*;

public class Client extends com.aliyun.gateway.spi.Client {

    public Client() throws Exception {
        super();
    }


    public void modifyConfiguration(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        config.endpoint = this.getEndpoint(config.network, config.endpoint);
    }

    public void modifyRequest(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        String date = com.aliyun.teautil.Common.getDateUTCString();
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("date", date),
                new TeaPair("host", config.endpoint),
                new TeaPair("x-acs-version", request.version),
                new TeaPair("x-acs-action", request.action),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("x-acs-signature-nonce", com.aliyun.teautil.Common.getNonce()),
                new TeaPair("accept", "application/json")
            ),
            request.headers
        );
        String signatureAlgorithm = com.aliyun.teautil.Common.defaultString(request.signatureAlgorithm, "ACS4-HMAC-SHA256");
        String signatureVersion = com.aliyun.teautil.Common.defaultString(request.signatureVersion, "v1");
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

        if (com.aliyun.darabonbastring.Client.equals(signatureVersion, "v4")) {
            if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
                request.headers.put("x-acs-content-sm3", hashedRequestPayload);
            } else {
                request.headers.put("x-acs-content-sha256", hashedRequestPayload);
            }

        } else {
            request.headers.put("x-acs-signature-method", "HMAC-SHA1");
            request.headers.put("x-acs-signature-version", "1.0");
        }

        if (!com.aliyun.teautil.Common.equalString(request.authType, "Anonymous") && !com.aliyun.teautil.Common.isUnset(request.credential)) {
            com.aliyun.credentials.Client credential = request.credential;
            com.aliyun.credentials.models.CredentialModel credentialModel = credential.getCredential();
            String authType = credentialModel.type;
            if (com.aliyun.teautil.Common.equalString(authType, "bearer")) {
                String bearerToken = credentialModel.bearerToken;
                request.headers.put("x-acs-bearer-token", bearerToken);
                request.headers.put("Authorization", "Bearer " + bearerToken + "");
            } else {
                String accessKeyId = credentialModel.accessKeyId;
                String accessKeySecret = credentialModel.accessKeySecret;
                String securityToken = credentialModel.securityToken;
                if (!com.aliyun.teautil.Common.empty(securityToken)) {
                    request.headers.put("x-acs-security-token", securityToken);
                }

                if (com.aliyun.darabonbastring.Client.equals(signatureVersion, "v4")) {
                    String dateNew = com.aliyun.darabonbastring.Client.subString(date, 0, 10);
                    String region = this.getRegion(config.endpoint);
                    byte[] signingkey = this.getSigningkey(signatureAlgorithm, accessKeySecret, region, dateNew);
                    request.headers.put("Authorization", this.getAuthorizationV4(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.productId, region, dateNew));
                } else {
                    request.headers.put("Authorization", this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret));
                }

            }

        }
    }

    public void modifyResponse(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            Object _res = com.aliyun.teautil.Common.readAsJSON(response.body);
            java.util.Map<String, Object> err = com.aliyun.teautil.Common.assertAsMap(_res);
            java.util.Map<String, String> headers = response.headers;
            String requestId = headers.get("x-ca-request-id");
            err.put("statusCode", response.statusCode);
            err.put("requestId", requestId);
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "" + this.defaultAny(err.get("Code"), err.get("code")) + ""),
                new TeaPair("message", "" + this.defaultAny(err.get("Message"), err.get("message")) + ""),
                new TeaPair("data", err)
            ));
        }

        if (!com.aliyun.teautil.Common.isUnset(response.body)) {
            if (com.aliyun.teautil.Common.equalNumber(response.statusCode, 204)) {
                com.aliyun.teautil.Common.readAsString(response.body);
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "binary")) {
                response.deserializedBody = response.body;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "byte")) {
                byte[] byt = com.aliyun.teautil.Common.readAsBytes(response.body);
                response.deserializedBody = byt;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "string")) {
                String str = com.aliyun.teautil.Common.readAsString(response.body);
                response.deserializedBody = str;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "json")) {
                response.deserializedBody = com.aliyun.teautil.Common.readAsJSON(response.body);
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "array")) {
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
        return com.aliyun.darabonbastring.Client.split(tmp, ";", null);
    }

    public String getRegion(String endpoint) throws Exception {
        String region = "center";
        if (com.aliyun.teautil.Common.empty(endpoint)) {
            return region;
        }

        if (com.aliyun.darabonbastring.Client.contains(endpoint, ".admin.aliyunpds.com")) {
            region = com.aliyun.darabonbastring.Client.replace(endpoint, ".admin.aliyunpds.com", "", null);
        }

        return region;
    }

    public byte[] getSigningkey(String signatureAlgorithm, String secret, String region, String date) throws Exception {
        String sc1 = "aliyun_v4" + secret + "";
        byte[] sc2 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            sc2 = com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(date, sc1);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
            sc2 = com.aliyun.darabonba.signature.Signer.HmacSM3Sign(date, sc1);
        }

        byte[] sc3 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            sc3 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(region, sc2);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
            sc3 = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes(region, sc2);
        }

        byte[] sc4 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            sc4 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("pds", sc3);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
            sc4 = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes("pds", sc3);
        }

        byte[] hmac = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            hmac = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
            hmac = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes("aliyun_v4_request", sc4);
        }

        return hmac;
    }

    public String getAuthorizationV4(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, String ak, byte[] signingkey, String product, String region, String date) throws Exception {
        String signature = this.getSignatureV4(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
        java.util.List<String> signedHeaders = this.getSignedHeaders(headers);
        String signedHeadersStr = com.aliyun.darabonba.array.Client.join(signedHeaders, ";");
        return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/aliyun_v4_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature + "";
    }

    public String getSignatureV4(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, byte[] signingkey) throws Exception {
        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResource(pathname, query);
        String canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
        java.util.List<String> signedHeaders = this.getSignedHeaders(headers);
        String signedHeadersStr = com.aliyun.darabonba.array.Client.join(signedHeaders, ";");
        stringToSign = "" + method + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload + "";
        String hex = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(stringToSign), signatureAlgorithm));
        stringToSign = "" + signatureAlgorithm + "\n" + hex + "";
        byte[] signature = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
            signature = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
            signature = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes(stringToSign, signingkey);
        }

        return com.aliyun.darabonba.encode.Encoder.hexEncode(signature);
    }
}
