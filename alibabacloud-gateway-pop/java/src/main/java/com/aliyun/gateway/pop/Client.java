// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.pop;

import com.aliyun.tea.*;

public class Client extends com.aliyun.gateway.spi.Client {

    public String _endpointSuffix;
    public String _signatureTypePrefix;
    public String _signPrefix;
    public String _sha256;
    public String _sm3;
    public Client() throws Exception {
        super();
        // CLOUD4-
        this._signatureTypePrefix = "ACS4-";
        // cloud_v4
        this._signPrefix = "aliyun_v4";
        this._endpointSuffix = "aliyuncs.com";
        this._sha256 = "" + _signatureTypePrefix + "HMAC-SHA256";
        this._sm3 = "" + _signatureTypePrefix + "HMAC-SM3";
    }


    public void modifyConfiguration(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        java.util.Map<String, String> attributes = attributeMap.key;
        if (!com.aliyun.teautil.Common.isUnset(attributes)) {
            this._signatureTypePrefix = attributes.get("signatureTypePrefix");
            this._signPrefix = attributes.get("signPrefix");
            this._endpointSuffix = attributes.get("endpointSuffix");
            this._sha256 = "" + _signatureTypePrefix + "HMAC-SHA256";
            this._sm3 = "" + _signatureTypePrefix + "HMAC-SM3";
        }

        config.endpoint = this.getEndpoint(request.productId, config.regionId, config.endpointRule, config.network, config.suffix, config.endpointMap, config.endpoint);
    }

    public void modifyRequest(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        String date = com.aliyun.openapiutil.Client.getTimestamp();
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("host", config.endpoint),
                new TeaPair("x-acs-version", request.version),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("x-acs-date", date),
                new TeaPair("x-acs-signature-nonce", com.aliyun.teautil.Common.getNonce()),
                new TeaPair("accept", "application/json")
            ),
            request.headers
        );
        if (!com.aliyun.teautil.Common.empty(request.action)) {
            request.headers.put("x-acs-action", request.action);
        }

        String signatureAlgorithm = com.aliyun.teautil.Common.defaultString(request.signatureAlgorithm, _sha256);
        String hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(""), signatureAlgorithm));
        if (!com.aliyun.teautil.Common.isUnset(request.stream)) {
            byte[] tmp = com.aliyun.teautil.Common.readAsBytes(request.stream);
            hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(tmp, signatureAlgorithm));
            request.stream = Tea.toReadable(tmp);
            request.headers.put("content-type", "application/octet-stream");
        } else {
            if (!com.aliyun.teautil.Common.isUnset(request.body)) {
                if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "byte")) {
                    byte[] byteObj = com.aliyun.teautil.Common.assertAsBytes(request.body);
                    hashedRequestPayload = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(byteObj, signatureAlgorithm));
                    request.stream = Tea.toReadable(byteObj);
                } else if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "json")) {
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

        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            request.headers.put("x-acs-content-sm3", hashedRequestPayload);
        } else {
            request.headers.put("x-acs-content-sha256", hashedRequestPayload);
        }

        if (!com.aliyun.teautil.Common.equalString(request.authType, "Anonymous")) {
            com.aliyun.credentials.Client credential = request.credential;
            if (com.aliyun.teautil.Common.isUnset(credential)) {
                throw new TeaException(TeaConverter.buildMap(
                    new TeaPair("code", "ParameterMissing"),
                    new TeaPair("message", "'config.credential' can not be unset")
                ));
            }

            com.aliyun.credentials.models.CredentialModel credentialModel = credential.getCredential();
            if (!com.aliyun.teautil.Common.empty(credentialModel.providerName)) {
                request.headers.put("x-acs-credentials-provider", credentialModel.providerName);
            }

            String authType = credentialModel.type;
            if (com.aliyun.teautil.Common.equalString(authType, "bearer")) {
                String bearerToken = credentialModel.bearerToken;
                request.headers.put("x-acs-bearer-token", bearerToken);
                request.headers.put("x-acs-signature-type", "BEARERTOKEN");
                request.headers.put("Authorization", "Bearer " + bearerToken + "");
            } else if (com.aliyun.teautil.Common.equalString(authType, "id_token")) {
                String idToken = credentialModel.securityToken;
                request.headers.put("x-acs-zero-trust-idtoken", idToken);
            } else {
                String accessKeyId = credentialModel.accessKeyId;
                String accessKeySecret = credentialModel.accessKeySecret;
                String securityToken = credentialModel.securityToken;
                if (!com.aliyun.teautil.Common.empty(securityToken)) {
                    request.headers.put("x-acs-accesskey-id", accessKeyId);
                    request.headers.put("x-acs-security-token", securityToken);
                }

                String dateNew = com.aliyun.darabonbastring.Client.subString(date, 0, 10);
                dateNew = com.aliyun.darabonbastring.Client.replace(dateNew, "-", "", null);
                String region = this.getRegion(request.productId, config.endpoint, config.regionId);
                byte[] signingkey = this.getSigningkey(signatureAlgorithm, accessKeySecret, request.productId, region, dateNew);
                request.headers.put("Authorization", this.getAuthorization(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.productId, region, dateNew));
            }

        }

    }

    public void modifyResponse(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            java.util.Map<String, Object> err = new java.util.HashMap<>();
            if (!com.aliyun.teautil.Common.isUnset(response.headers.get("content-type")) && com.aliyun.darabonbastring.Client.contains(response.headers.get("content-type"), "text/xml")) {
                String _str = com.aliyun.teautil.Common.readAsString(response.body);
                java.util.Map<String, Object> respMap = com.aliyun.teaxml.Client.parseXml(_str, null);
                err = com.aliyun.teautil.Common.assertAsMap(respMap.get("Error"));
            } else {
                Object _res = com.aliyun.teautil.Common.readAsJSON(response.body);
                err = com.aliyun.teautil.Common.assertAsMap(_res);
            }

            Object requestId = this.defaultAny(err.get("RequestId"), err.get("requestId"));
            if (!com.aliyun.teautil.Common.isUnset(response.headers.get("x-acs-request-id"))) {
                requestId = response.headers.get("x-acs-request-id");
            }

            err.put("statusCode", response.statusCode);
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "" + this.defaultAny(err.get("Code"), err.get("code")) + ""),
                new TeaPair("message", "code: " + response.statusCode + ", " + this.defaultAny(err.get("Message"), err.get("message")) + " request id: " + requestId + ""),
                new TeaPair("data", err),
                new TeaPair("description", "" + this.defaultAny(err.get("Description"), err.get("description")) + ""),
                new TeaPair("accessDeniedDetail", this.defaultAny(err.get("AccessDeniedDetail"), err.get("accessDeniedDetail")))
            ));
        }

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

    public String getAuthorization(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, String ak, byte[] signingkey, String product, String region, String date) throws Exception {
        String signature = this.getSignature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
        java.util.List<String> signedHeaders = this.getSignedHeaders(headers);
        String signedHeadersStr = com.aliyun.darabonba.array.Client.join(signedHeaders, ";");
        return "" + signatureAlgorithm + " Credential=" + ak + "/" + date + "/" + region + "/" + product + "/" + _signPrefix + "_request,SignedHeaders=" + signedHeadersStr + ",Signature=" + signature + "";
    }

    public String getSignature(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String signatureAlgorithm, String payload, byte[] signingkey) throws Exception {
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
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            signature = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            signature = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes(stringToSign, signingkey);
        }

        return com.aliyun.darabonba.encode.Encoder.hexEncode(signature);
    }

    public byte[] getSigningkey(String signatureAlgorithm, String secret, String product, String region, String date) throws Exception {
        String sc1 = "" + _signPrefix + "" + secret + "";
        byte[] sc2 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            sc2 = com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(date, sc1);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            sc2 = com.aliyun.darabonba.signature.Signer.HmacSM3Sign(date, sc1);
        }

        byte[] sc3 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            sc3 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(region, sc2);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            sc3 = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes(region, sc2);
        }

        byte[] sc4 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            sc4 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(product, sc3);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            sc4 = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes(product, sc3);
        }

        byte[] hmac = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            hmac = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("" + _signPrefix + "_request", sc4);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            hmac = com.aliyun.darabonba.signature.Signer.HmacSM3SignByBytes("" + _signPrefix + "_request", sc4);
        }

        return hmac;
    }

    public String getRegion(String product, String endpoint, String regionId) throws Exception {
        if (!com.aliyun.teautil.Common.empty(regionId)) {
            return regionId;
        }

        String region = "center";
        if (com.aliyun.teautil.Common.empty(product) || com.aliyun.teautil.Common.empty(endpoint)) {
            return region;
        }

        java.util.List<String> strs = com.aliyun.darabonbastring.Client.split(endpoint, ":", null);
        String withoutPort = strs.get(0);
        String preRegion = com.aliyun.darabonbastring.Client.replace(withoutPort, "." + _endpointSuffix + "", "", null);
        java.util.List<String> nodes = com.aliyun.darabonbastring.Client.split(preRegion, ".", null);
        if (com.aliyun.teautil.Common.equalNumber(com.aliyun.darabonba.array.Client.size(nodes), 2)) {
            region = nodes.get(1);
        }

        return region;
    }

    public String buildCanonicalizedResource(java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryArray = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedQueryArray = com.aliyun.darabonba.array.Client.ascSort(queryArray);
            String separator = "";
            for (String key : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + com.aliyun.darabonba.encode.Encoder.percentEncode(key) + "=";
                if (!com.aliyun.teautil.Common.empty(query.get(key))) {
                    canonicalizedResource = "" + canonicalizedResource + "" + com.aliyun.darabonba.encode.Encoder.percentEncode(query.get(key)) + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeaders(java.util.Map<String, String> headers) throws Exception {
        // lower header key
        java.util.List<String> headersArray = com.aliyun.darabonba.map.Client.keySet(headers);
        java.util.Map<String, String> newHeaders = new java.util.HashMap<>();
        String tmp = "";
        for (String key : headersArray) {
            String lowerKey = com.aliyun.darabonbastring.Client.toLower(key);
            String value = headers.get(key);
            if (!com.aliyun.teautil.Common.isUnset(value)) {
                if (!com.aliyun.darabonbastring.Client.contains(tmp, lowerKey)) {
                    tmp = "" + tmp + "," + lowerKey + "";
                    newHeaders.put(lowerKey, com.aliyun.darabonbastring.Client.trim(value));
                } else {
                    newHeaders.put(lowerKey, "" + newHeaders.get(lowerKey) + "," + com.aliyun.darabonbastring.Client.trim(value) + "");
                }

            }

        }
        String canonicalizedHeaders = "";
        java.util.List<String> sortedHeaders = this.getSignedHeaders(headers);
        for (String header : sortedHeaders) {
            canonicalizedHeaders = "" + canonicalizedHeaders + "" + header + ":" + newHeaders.get(header) + "\n";
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
                String value = headers.get(key);
                if (!com.aliyun.teautil.Common.isUnset(value) && !com.aliyun.darabonbastring.Client.contains(tmp, lowerKey)) {
                    tmp = "" + tmp + "" + separator + "" + lowerKey + "";
                    separator = ";";
                }

            }

        }
        return com.aliyun.darabonbastring.Client.split(tmp, ";", null);
    }
}
