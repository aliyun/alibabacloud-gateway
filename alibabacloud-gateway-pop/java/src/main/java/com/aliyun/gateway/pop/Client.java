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
        String hashedRequestPayload = Codec.hexEncode(Codec.hash(com.aliyun.teautil.Common.toBytes(""), signatureAlgorithm));
        if (!com.aliyun.teautil.Common.isUnset(request.stream)) {
            byte[] tmp = com.aliyun.teautil.Common.readAsBytes(request.stream);
            hashedRequestPayload = Codec.hexEncode(Codec.hash(tmp, signatureAlgorithm));
            request.stream = Tea.toReadable(tmp);
            request.headers.put("content-type", "application/octet-stream");
        } else {
            if (!com.aliyun.teautil.Common.isUnset(request.body)) {
                if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "byte")) {
                    byte[] byteObj = com.aliyun.teautil.Common.assertAsBytes(request.body);
                    hashedRequestPayload = Codec.hexEncode(Codec.hash(byteObj, signatureAlgorithm));
                    request.stream = Tea.toReadable(byteObj);
                } else if (com.aliyun.teautil.Common.equalString(request.reqBodyType, "json")) {
                    String jsonObj = com.aliyun.teautil.Common.toJSONString(request.body);
                    hashedRequestPayload = Codec.hexEncode(Codec.hash(com.aliyun.teautil.Common.toBytes(jsonObj), signatureAlgorithm));
                    request.stream = Tea.toReadable(jsonObj);
                    request.headers.put("content-type", "application/json; charset=utf-8");
                } else {
                    java.util.Map<String, Object> m = com.aliyun.teautil.Common.assertAsMap(request.body);
                    String formObj = com.aliyun.openapiutil.Client.toForm(m);
                    hashedRequestPayload = Codec.hexEncode(Codec.hash(com.aliyun.teautil.Common.toBytes(formObj), signatureAlgorithm));
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

                String dateNew = Codec.subString(date, 0, 10);
                dateNew = Codec.replace(dateNew, "-", "", null);
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
            if (!com.aliyun.teautil.Common.isUnset(response.headers.get("content-type")) && Codec.contains(response.headers.get("content-type"), "text/xml")) {
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
        String signedHeadersStr = Codec.join(signedHeaders, ";");
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
        this.validateSignedHeaders(signedHeaders);
        String signedHeadersStr = Codec.join(signedHeaders, ";");
        stringToSign = "" + method + "\n" + canonicalURI + "\n" + canonicalizedResource + "\n" + canonicalizedHeaders + "\n" + signedHeadersStr + "\n" + payload + "";
        String hex = Codec.hexEncode(Codec.hash(com.aliyun.teautil.Common.toBytes(stringToSign), signatureAlgorithm));
        stringToSign = "" + signatureAlgorithm + "\n" + hex + "";
        byte[] signature = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            signature = Codec.hmacSha256ByBytes(stringToSign, signingkey);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            signature = Codec.hmacSm3ByBytes(stringToSign, signingkey);
        }

        return Codec.hexEncode(signature);
    }

    public byte[] getSigningkey(String signatureAlgorithm, String secret, String product, String region, String date) throws Exception {
        String sc1 = "" + _signPrefix + "" + secret + "";
        byte[] sc2 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            sc2 = Codec.hmacSha256(date, sc1);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            sc2 = Codec.hmacSm3(date, sc1);
        }

        byte[] sc3 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            sc3 = Codec.hmacSha256ByBytes(region, sc2);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            sc3 = Codec.hmacSm3ByBytes(region, sc2);
        }

        byte[] sc4 = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            sc4 = Codec.hmacSha256ByBytes(product, sc3);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            sc4 = Codec.hmacSm3ByBytes(product, sc3);
        }

        byte[] hmac = com.aliyun.teautil.Common.toBytes("");
        if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sha256)) {
            hmac = Codec.hmacSha256ByBytes("" + _signPrefix + "_request", sc4);
        } else if (com.aliyun.teautil.Common.equalString(signatureAlgorithm, _sm3)) {
            hmac = Codec.hmacSm3ByBytes("" + _signPrefix + "_request", sc4);
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

        java.util.List<String> strs = Codec.split(endpoint, ":", null);
        String withoutPort = strs.get(0);
        String preRegion = Codec.replace(withoutPort, "." + _endpointSuffix + "", "", null);
        java.util.List<String> nodes = Codec.split(preRegion, ".", null);
        if (com.aliyun.teautil.Common.equalNumber(Codec.size(nodes), 2)) {
            region = nodes.get(1);
        }

        return region;
    }

    public String buildCanonicalizedResource(java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryArray = Codec.keySet(query);
            java.util.List<String> sortedQueryArray = Codec.ascSort(queryArray);
            String separator = "";
            for (String key : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + Codec.percentEncode(key) + "=";
                if (!com.aliyun.teautil.Common.empty(query.get(key))) {
                    canonicalizedResource = "" + canonicalizedResource + "" + Codec.percentEncode(query.get(key)) + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeaders(java.util.Map<String, String> headers) throws Exception {
        // lower header key
        java.util.List<String> headersArray = Codec.keySet(headers);
        java.util.Map<String, String> newHeaders = new java.util.HashMap<>();
        for (String key : headersArray) {
            String lowerKey = Codec.toLower(key);
            String value = headers.get(key);
            if (!com.aliyun.teautil.Common.isUnset(value)) {
                if (com.aliyun.teautil.Common.isUnset(newHeaders.get(lowerKey)) || Codec.equals(lowerKey, "host")) {
                    newHeaders.put(lowerKey, Codec.trim(value));
                } else {
                    newHeaders.put(lowerKey, "" + newHeaders.get(lowerKey) + "," + Codec.trim(value) + "");
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

    public void validateSignedHeaders(java.util.List<String> signedHeaders) throws Exception {
        if (!Codec.listContains(signedHeaders, "host") || !Codec.listContains(signedHeaders, "x-acs-date")) {
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "InvalidSignedHeaders"),
                new TeaPair("message", "signed headers must include host and x-acs-date")
            ));
        }
    }

    public java.util.List<String> getSignedHeaders(java.util.Map<String, String> headers) throws Exception {
        java.util.List<String> headersArray = Codec.keySet(headers);
        java.util.List<String> newHeadersArray = new java.util.ArrayList<>();
        for (String key : headersArray) {
            String lowerKey = Codec.toLower(key);
            String value = headers.get(key);
            if (!com.aliyun.teautil.Common.isUnset(value)) {
                Codec.append(newHeadersArray, lowerKey);
            }

        }
        java.util.List<String> sortedHeadersArray = Codec.ascSort(newHeadersArray);
        java.util.List<String> result = new java.util.ArrayList<>();
        for (String key : sortedHeadersArray) {
            if (Codec.hasPrefix(key, "x-acs-") || Codec.equals(key, "host") || Codec.equals(key, "content-type")) {
                if (!Codec.listContains(result, key)) {
                    Codec.append(result, key);
                }

            }

        }
        return result;
    }
}
