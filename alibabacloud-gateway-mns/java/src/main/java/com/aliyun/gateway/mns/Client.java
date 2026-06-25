// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.mns;

import com.aliyun.tea.*;

public class Client extends com.aliyun.gateway.spi.Client {

    public String _signPrefix;
    public String _signSuffix;
    public String _authPrefix;
    public Client() throws Exception {
        super();
        this._signPrefix = "aliyun_v4";
        this._signSuffix = "aliyun_v4_request";
        this._authPrefix = "MNS4-HMAC-SHA256";
    }


    public void modifyConfiguration(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        config.endpoint = this.getEndpoint(config.regionId, config.endpoint);
    }

    public void modifyRequest(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        if (!com.aliyun.teautil.Common.isUnset(request.signatureVersion) && com.aliyun.darabonbastring.Client.equals(request.signatureVersion, "v2")) {
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "UnsupportedSignatureVersion"),
                new TeaPair("message", "MNS gateway does not support signature version v2, please use v4")
            ));
        }

        if (!com.aliyun.teautil.Common.isUnset(request.body)) {
            if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "xml")) {
                java.util.Map<String, Object> reqBodyMap = com.aliyun.teautil.Common.assertAsMap(request.body);
                String xmlStr = com.aliyun.teaxml.Client.toXML(reqBodyMap);
                request.stream = Tea.toReadable(xmlStr);
                request.headers.put("content-type", "text/xml;charset=UTF-8");
                request.headers.put("content-md5", com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.MD5Sign(xmlStr)));
            } else if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "json") || com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "formData")) {
                String bodyStr = com.aliyun.teautil.Common.toJSONString(request.body);
                request.stream = Tea.toReadable(bodyStr);
                request.headers.put("content-type", "application/json");
                request.headers.put("content-md5", com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.MD5Sign(bodyStr)));
            } else if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "byte") || com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "binary")) {
                byte[] bodyBytes = com.aliyun.teautil.Common.assertAsBytes(request.body);
                request.stream = Tea.toReadable(bodyBytes);
                request.headers.put("content-md5", com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.MD5SignForBytes(bodyBytes)));
            }

        }

        this.buildRequest(context);
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("host", config.endpoint),
                new TeaPair("x-mns-version", request.version),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("accept", "application/json")
            ),
            request.headers
        );
        if (!com.aliyun.teautil.Common.equalString(request.authType, "Anonymous")) {
            com.aliyun.credentials.Client credential = request.credential;
            if (com.aliyun.teautil.Common.isUnset(credential)) {
                throw new TeaException(TeaConverter.buildMap(
                    new TeaPair("code", "ParameterMissing"),
                    new TeaPair("message", "'config.credential' can not be unset")
                ));
            }

            com.aliyun.credentials.models.CredentialModel credentialModel = credential.getCredential();
            String authType = credentialModel.type;
            if (com.aliyun.teautil.Common.equalString(authType, "bearer")) {
                String bearerToken = credentialModel.bearerToken;
                request.headers.put("x-acs-bearer-token", bearerToken);
                request.headers.put("x-acs-signature-type", "BEARERTOKEN");
                request.headers.put("Authorization", "Bearer " + bearerToken + "");
            } else {
                String accessKeyId = credentialModel.accessKeyId;
                String accessKeySecret = credentialModel.accessKeySecret;
                String securityToken = credentialModel.securityToken;
                if (!com.aliyun.teautil.Common.empty(securityToken)) {
                    request.headers.put("security-token", securityToken);
                }

                request.headers.put("date", com.aliyun.teautil.Common.getDateUTCString());
                String date = this.getDateISO8601();
                request.headers.put("authorization", this.getAuthorizationV4(context, date, accessKeyId, accessKeySecret));

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
            if (!com.aliyun.teautil.Common.isUnset(response.headers.get("x-mns-request-id"))) {
                requestId = response.headers.get("x-mns-request-id");
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

    public String getEndpoint(String regionId, String endpoint) throws Exception {
        if (!com.aliyun.teautil.Common.empty(endpoint)) {
            return endpoint;
        }

        if (com.aliyun.teautil.Common.empty(regionId)) {
            regionId = "cn-hangzhou";
        }

        return "" + regionId + ".mns.aliyuncs.com";
    }

    public Object defaultAny(Object inputValue, Object defaultValue) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(inputValue)) {
            return defaultValue;
        }

        return inputValue;
    }

    public String defaultString(String inputValue, String defaultValue) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(inputValue) || com.aliyun.teautil.Common.empty(inputValue)) {
            return defaultValue;
        }

        return inputValue;
    }

    public static String base64Encode(String input) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(input)) {
            return "";
        }

        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.teautil.Common.toBytes(input));
    }

    public static String base64Decode(String input) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(input)) {
            return "";
        }

        return com.aliyun.teautil.Common.toString(com.aliyun.darabonba.encode.Encoder.base64Decode(input));
    }

    public void buildRequest(com.aliyun.gateway.spi.models.InterceptorContext context) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        String resource = request.pathname;
        if (!com.aliyun.teautil.Common.empty(resource)) {
            java.util.List<String> paths = com.aliyun.darabonbastring.Client.split(resource, "?", 2);
            resource = paths.get(0);
            if (com.aliyun.teautil.Common.equalNumber(com.aliyun.darabonba.array.Client.size(paths), 2)) {
                java.util.List<String> params = com.aliyun.darabonbastring.Client.split(paths.get(1), "&", null);
                for (String sub : params) {
                    java.util.List<String> item = com.aliyun.darabonbastring.Client.split(sub, "=", null);
                    String key = item.get(0);
                    String value = null;
                    if (com.aliyun.teautil.Common.equalNumber(com.aliyun.darabonba.array.Client.size(item), 2)) {
                        value = item.get(1);
                    }

                    request.query.put(key, value);
                }
            }

        }

        request.pathname = resource;
    }

    public String getAuthorizationV2(String pathname, String method, java.util.Map<String, String> headers, String ak, String secret) throws Exception {
        String sign = this.getSignatureV2(pathname, method, headers, secret);
        return "MNS " + ak + ":" + sign + "";
    }

    public String getSignatureV2(String pathname, String method, java.util.Map<String, String> headers, String secret) throws Exception {
        String canonicalizedResource = this.defaultString(pathname, "/");
        String canonicalizedHeaders = this.buildCanonicalizedHeadersV2(headers);
        String stringToSign = "" + method + "\n" + canonicalizedHeaders + "" + canonicalizedResource + "";
        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.HmacSHA1Sign(stringToSign, secret));
    }

    public String buildCanonicalizedHeadersV2(java.util.Map<String, String> headers) throws Exception {
        String contentMd5 = this.defaultString(headers.get("content-md5"), "");
        String contentType = this.defaultString(headers.get("content-type"), "");
        String date = this.defaultString(headers.get("date"), "");
        String canonicalizedHeaders = "" + contentMd5 + "\n" + contentType + "\n" + date + "\n";
        java.util.Map<String, String> mnsHeaders = new java.util.HashMap<>();
        for (String header : com.aliyun.darabonba.map.Client.keySet(headers)) {
            String lowerHeader = com.aliyun.darabonbastring.Client.toLower(header);
            if (com.aliyun.darabonbastring.Client.hasPrefix(lowerHeader, "x-mns-")) {
                mnsHeaders.put(lowerHeader, headers.get(header));
            }

        }
        for (String header : com.aliyun.darabonba.array.Client.ascSort(com.aliyun.darabonba.map.Client.keySet(mnsHeaders))) {
            canonicalizedHeaders = "" + canonicalizedHeaders + "" + header + ":" + mnsHeaders.get(header) + "\n";
        }
        return canonicalizedHeaders;
    }

    public String getAuthorizationV4(com.aliyun.gateway.spi.models.InterceptorContext context, String date, String accessKeyId, String accessKeySecret) throws Exception {
        String region = this.getRegion(context);
        String dateShort = com.aliyun.darabonbastring.Client.subString(date, 0, 8);
        dateShort = com.aliyun.darabonbastring.Client.replace(dateShort, "T", "", null);
        String scope = "" + dateShort + "/" + region + "/mns/" + _signSuffix + "";
        byte[] signingkey = this.getSigningkeyV4(accessKeySecret, region, dateShort);
        String signature = this.getSignatureV4(context, date, scope, signingkey);
        return "" + _authPrefix + " Credential=" + accessKeyId + "/" + scope + ",Signature=" + signature + "";
    }

    public byte[] getSigningkeyV4(String secret, String region, String date) throws Exception {
        String sc1 = "" + _signPrefix + "" + secret + "";
        byte[] sc2 = com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(date, sc1);
        byte[] sc3 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(region, sc2);
        byte[] sc4 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("mns", sc3);
        return com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(_signSuffix, sc4);
    }

    public String getSignatureV4(com.aliyun.gateway.spi.models.InterceptorContext context, String date, String scope, byte[] signingkey) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        String canonicalString = this.buildCanonicalRequestV4(request.pathname, request.method, request.query, request.headers);
        String stringToSign = "" + _authPrefix + "\n" + date + "\n" + scope + "\n" + canonicalString + "";
        byte[] signature = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
        return com.aliyun.darabonba.encode.Encoder.hexEncode(signature);
    }

    public String buildCanonicalRequestV4(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers) throws Exception {
        String canonicalURI = "/";
        if (!com.aliyun.teautil.Common.empty(pathname)) {
            canonicalURI = pathname;
            if (!com.aliyun.darabonbastring.Client.hasPrefix(canonicalURI, "/")) {
                canonicalURI = "/" + canonicalURI + "";
            }

        }

        String suffix = "";
        if (!com.aliyun.darabonbastring.Client.equals(canonicalURI, "/") && com.aliyun.darabonbastring.Client.hasSuffix(canonicalURI, "/")) {
            suffix = "/";
        }

        canonicalURI = "" + com.aliyun.openapiutil.Client.getEncodePath(canonicalURI) + "" + suffix + "";
        String resources = this.buildCanonicalizedResourceV4(query);
        String canonicalHeaders = this.buildCanonicalizedHeadersV4(headers);
        return "" + method + "\n" + canonicalURI + "\n" + resources + "\n" + canonicalHeaders + "\n";
    }

    public String buildCanonicalizedResourceV4(java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.Map<String, String> queryMap = new java.util.HashMap<>();
            for (String key : com.aliyun.darabonba.map.Client.keySet(query)) {
                String encodedKey = this.percentEncodeMns(com.aliyun.darabonbastring.Client.toLower(com.aliyun.darabonbastring.Client.trim(key)));
                String encodedValue = "";
                if (!com.aliyun.teautil.Common.isUnset(query.get(key)) && !com.aliyun.teautil.Common.empty(query.get(key))) {
                    encodedValue = this.percentEncodeMns(com.aliyun.darabonbastring.Client.trim(query.get(key)));
                }

                queryMap.put(encodedKey, encodedValue);
            }
            java.util.List<String> queryArray = com.aliyun.darabonba.map.Client.keySet(queryMap);
            java.util.List<String> sortedQueryArray = com.aliyun.darabonba.array.Client.ascSort(queryArray);
            String separator = "";
            for (String encodedKey : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + encodedKey + "";
                String encodedValue = queryMap.get(encodedKey);
                if (!com.aliyun.teautil.Common.empty(encodedValue)) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + encodedValue + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeadersV4(java.util.Map<String, String> headers) throws Exception {
        java.util.Map<String, String> signedHeaders = new java.util.HashMap<>();
        for (String key : com.aliyun.darabonba.map.Client.keySet(headers)) {
            String lowerKey = com.aliyun.darabonbastring.Client.toLower(key);
            if (this.hasSignedHeaderV4(lowerKey)) {
                signedHeaders.put(lowerKey, com.aliyun.darabonbastring.Client.trim(headers.get(key)));
            }

        }
        String canonicalizedHeaders = "";
        for (String lowerKey : com.aliyun.darabonba.array.Client.ascSort(com.aliyun.darabonba.map.Client.keySet(signedHeaders))) {
            canonicalizedHeaders = "" + canonicalizedHeaders + "" + lowerKey + ":" + signedHeaders.get(lowerKey) + "\n";
        }
        return canonicalizedHeaders;
    }

    public Boolean hasSignedHeaderV4(String header) throws Exception {
        if (com.aliyun.darabonbastring.Client.equals(header, "content-type") || com.aliyun.darabonbastring.Client.equals(header, "content-md5")) {
            return true;
        }

        return com.aliyun.darabonbastring.Client.hasPrefix(header, "x-mns-");
    }

    public String percentEncodeMns(String value) throws Exception {
        String encoded = com.aliyun.darabonba.encode.Encoder.percentEncode(value);
        return com.aliyun.darabonbastring.Client.replace(encoded, "+", "%20", null);
    }

    public String getRegion(com.aliyun.gateway.spi.models.InterceptorContext context) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        if (!com.aliyun.teautil.Common.isUnset(config.regionId) && !com.aliyun.teautil.Common.empty(config.regionId)) {
            return config.regionId;
        }

        String region = com.aliyun.darabonbastring.Client.replace(config.endpoint, ".mns.aliyuncs.com", "", null);
        if (com.aliyun.darabonbastring.Client.equals(region, config.endpoint)) {
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "ClientConfigError"),
                new TeaPair("message", "The regionId configuration of mns client is missing.")
            ));
        }

        return region;
    }

    public String getDateISO8601() throws Exception {
        String date = com.aliyun.openapiutil.Client.getTimestamp();
        date = com.aliyun.darabonbastring.Client.replace(date, "-", "", null);
        return com.aliyun.darabonbastring.Client.replace(date, ":", "", null);
    }
}
