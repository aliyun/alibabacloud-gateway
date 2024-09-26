// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls;

import com.aliyun.tea.*;

public class Client extends com.aliyun.gateway.spi.Client {

    public Client() throws Exception {
        super();
    }


    public void modifyConfiguration(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        config.endpoint = this.getEndpoint(config.regionId, config.network, config.endpoint);
    }

    public void modifyRequest(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        java.util.Map<String, String> hostMap = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.hostMap)) {
            hostMap = request.hostMap;
        }

        String project = hostMap.get("project");
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        com.aliyun.credentials.Client credential = request.credential;
        com.aliyun.credentials.models.CredentialModel credentialModel = credential.getCredential();
        String accessKeyId = credentialModel.accessKeyId;
        String accessKeySecret = credentialModel.accessKeySecret;
        String securityToken = credentialModel.securityToken;
        if (!com.aliyun.teautil.Common.empty(securityToken)) {
            request.headers.put("x-acs-security-token", securityToken);
        }

        String signatureVersion = com.aliyun.teautil.Common.defaultString(request.signatureVersion, "v1");
        String contentHash = "";
        if (!com.aliyun.teautil.Common.isUnset(request.body)) {
            if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "json")) {
                String bodyStr = com.aliyun.teautil.Common.toJSONString(request.body);
                contentHash = this.MakeContentHash(com.aliyun.teautil.Common.toBytes(bodyStr), signatureVersion);
                request.stream = Tea.toReadable(bodyStr);
                request.headers.put("content-type", "application/json");
            } else if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "formData")) {
                String str = com.aliyun.teautil.Common.toJSONString(request.body);
                contentHash = this.MakeContentHash(com.aliyun.teautil.Common.toBytes(str), signatureVersion);
                request.stream = Tea.toReadable(str);
                request.headers.put("content-type", "application/json");
            } else if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "binary")) {
                // content-type: application/octet-stream
                byte[] bodyBytes = com.aliyun.teautil.Common.assertAsBytes(request.body);
                contentHash = this.MakeContentHash(bodyBytes, signatureVersion);
                request.stream = Tea.toReadable(bodyBytes);
            }

        }

        String host = this.getHost(config.network, project, config.endpoint);
        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("accept", "application/json"),
                new TeaPair("host", host),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("x-log-apiversion", "0.6.0"),
                new TeaPair("x-log-bodyrawsize", "0")
            ),
            request.headers
        );
        this.buildRequest(context);
        // move param in path to query
        if (com.aliyun.darabonbastring.Client.equals(signatureVersion, "v4")) {
            if (com.aliyun.teautil.Common.empty(contentHash)) {
                contentHash = "e3b0c44298fc1c149afbf4c8996fb9242a7e41e4649b934ca495991b7852b855";
            }

            String date = this.getDateISO8601();
            request.headers.put("x-log-date", date);
            request.headers.put("x-log-content-sha256", contentHash);
            request.headers.put("authorization", this.getAuthorizationV4(context, date, contentHash, accessKeyId, accessKeySecret));
            return ;
        }

        if (!com.aliyun.teautil.Common.empty(accessKeyId)) {
            request.headers.put("x-log-signaturemethod", "hmac-sha256");
        }

        request.headers.put("date", com.aliyun.teautil.Common.getDateUTCString());
        request.headers.put("content-md5", contentHash);
        request.headers.put("authorization", this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret));
    }

    public String MakeContentHash(byte[] content, String signatureVersion) throws Exception {
        if (com.aliyun.darabonbastring.Client.equals(signatureVersion, "v4")) {
            // TODO: 这里应当检查 length == 0，但是还不支持。通常情况下也不会出现 body 设置了但是长度为 0
            if (com.aliyun.teautil.Common.isUnset(content)) {
                return "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
            }

            return com.aliyun.darabonbastring.Client.toLower(com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(content, "SLS4-HMAC-SHA256")));
        }

        return com.aliyun.darabonbastring.Client.toUpper(com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.signature.Signer.MD5SignForBytes(content)));
    }

    public void modifyResponse(com.aliyun.gateway.spi.models.InterceptorContext context, com.aliyun.gateway.spi.models.AttributeMap attributeMap) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            Object error = com.aliyun.teautil.Common.readAsJSON(response.body);
            java.util.Map<String, Object> resMap = com.aliyun.teautil.Common.assertAsMap(error);
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", resMap.get("errorCode")),
                new TeaPair("message", resMap.get("errorMessage")),
                new TeaPair("accessDeniedDetail", resMap.get("accessDeniedDetail")),
                new TeaPair("data", TeaConverter.buildMap(
                    new TeaPair("httpCode", response.statusCode),
                    new TeaPair("requestId", response.headers.get("x-log-requestid")),
                    new TeaPair("statusCode", response.statusCode)
                ))
            ));
        }

        if (!com.aliyun.teautil.Common.isUnset(response.body)) {
            String bodyrawSize = response.headers.get("x-log-bodyrawsize");
            String compressType = response.headers.get("x-log-compresstype");
            java.io.InputStream uncompressedData = response.body;
            if (!com.aliyun.teautil.Common.isUnset(bodyrawSize) && !com.aliyun.teautil.Common.isUnset(compressType)) {
                uncompressedData = com.aliyun.gateway.sls.util.Client.readAndUncompressBlock(response.body, compressType, bodyrawSize);
            }

            if (com.aliyun.teautil.Common.equalString(request.bodyType, "binary")) {
                response.deserializedBody = uncompressedData;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "byte")) {
                byte[] byt = com.aliyun.teautil.Common.readAsBytes(uncompressedData);
                response.deserializedBody = byt;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "string")) {
                response.deserializedBody = com.aliyun.teautil.Common.readAsString(uncompressedData);
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "json")) {
                Object obj = com.aliyun.teautil.Common.readAsJSON(uncompressedData);
                // var res = Util.assertAsMap(obj);
                response.deserializedBody = obj;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "array")) {
                response.deserializedBody = com.aliyun.teautil.Common.readAsJSON(uncompressedData);
            } else {
                response.deserializedBody = com.aliyun.teautil.Common.readAsString(uncompressedData);
            }

        }

    }

    public String getEndpoint(String regionId, String network, String endpoint) throws Exception {
        if (!com.aliyun.teautil.Common.empty(endpoint)) {
            return endpoint;
        }

        if (com.aliyun.teautil.Common.empty(regionId)) {
            regionId = "cn-hangzhou";
        }

        if (!com.aliyun.teautil.Common.empty(network)) {
            if (com.aliyun.darabonbastring.Client.equals(network, "intranet")) {
                return "" + regionId + "-intranet.log.aliyuncs.com";
            } else if (com.aliyun.darabonbastring.Client.equals(network, "accelerate")) {
                return "log-global.aliyuncs.com";
            } else if (com.aliyun.darabonbastring.Client.equals(network, "share")) {
                if (com.aliyun.darabonbastring.Client.equals(regionId, "cn-hangzhou-corp") || com.aliyun.darabonbastring.Client.equals(regionId, "cn-shanghai-corp")) {
                    return "" + regionId + ".sls.aliyuncs.com";
                } else if (com.aliyun.darabonbastring.Client.equals(regionId, "cn-zhangbei-corp")) {
                    return "zhangbei-corp-share.log.aliyuncs.com";
                }

                return "" + regionId + "-share.log.aliyuncs.com";
            }

        }

        return "" + regionId + ".log.aliyuncs.com";
    }

    public String getHost(String network, String project, String endpoint) throws Exception {
        if (com.aliyun.teautil.Common.isUnset(project)) {
            return endpoint;
        }

        return "" + project + "." + endpoint + "";
    }

    public String getAuthorization(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String ak, String secret) throws Exception {
        String sign = this.getSignature(pathname, method, query, headers, secret);
        return "LOG " + ak + ":" + sign + "";
    }

    public String getSignature(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String secret) throws Exception {
        String resource = pathname;
        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResource(resource, query);
        String canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
        stringToSign = "" + method + "\n" + canonicalizedHeaders + "" + canonicalizedResource + "";
        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(stringToSign, secret));
    }

    public String buildCanonicalizedResource(String pathname, java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = pathname;
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryList = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedParams = com.aliyun.darabonba.array.Client.ascSort(queryList);
            String separator = "?";
            for (String paramName : sortedParams) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + paramName + "";
                String paramValue = query.get(paramName);
                if (!com.aliyun.teautil.Common.isUnset(paramValue)) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + paramValue + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeaders(java.util.Map<String, String> headers) throws Exception {
        String canonicalizedHeaders = "";
        String contentType = headers.get("content-type");
        if (com.aliyun.teautil.Common.isUnset(contentType)) {
            contentType = "";
        }

        String contentMd5 = headers.get("content-md5");
        if (com.aliyun.teautil.Common.isUnset(contentMd5)) {
            contentMd5 = "";
        }

        canonicalizedHeaders = "" + canonicalizedHeaders + "" + contentMd5 + "\n" + contentType + "\n" + headers.get("date") + "\n";
        java.util.List<String> keys = com.aliyun.darabonba.map.Client.keySet(headers);
        java.util.List<String> sortedHeaders = com.aliyun.darabonba.array.Client.ascSort(keys);
        for (String header : sortedHeaders) {
            if (com.aliyun.darabonbastring.Client.contains(com.aliyun.darabonbastring.Client.toLower(header), "x-log-") || com.aliyun.darabonbastring.Client.contains(com.aliyun.darabonbastring.Client.toLower(header), "x-acs-")) {
                canonicalizedHeaders = "" + canonicalizedHeaders + "" + header + ":" + headers.get(header) + "\n";
            }

        }
        return canonicalizedHeaders;
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

    public String getAuthorizationV4(com.aliyun.gateway.spi.models.InterceptorContext context, String date, String contentHash, String accessKeyId, String accessKeySecret) throws Exception {
        String region = this.getRegion(context);
        String headerStr = this.getSignedHeaderStrV4(context.request.headers);
        String dateShort = com.aliyun.darabonbastring.Client.subString(date, 0, 8);
        dateShort = com.aliyun.darabonbastring.Client.replace(dateShort, "T", "", null);
        // for fix php sdk bug
        String scope = "" + dateShort + "/" + region + "/sls/aliyun_v4_request";
        byte[] signingkey = this.getSigningkeyV4("SLS4-HMAC-SHA256", accessKeySecret, region, dateShort);
        String signature = this.getSignatureV4(context, "SLS4-HMAC-SHA256", headerStr, date, scope, contentHash, signingkey);
        return "" + "SLS4-HMAC-SHA256" + " Credential=" + accessKeyId + "/" + scope + ",Signature=" + signature + "";
    }

    public byte[] getSigningkeyV4(String signatureAlgorithm, String secret, String region, String date) throws Exception {
        String sc1 = "aliyun_v4" + secret + "";
        byte[] sc2 = com.aliyun.darabonba.signature.Signer.HmacSHA256Sign(date, sc1);
        byte[] sc3 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(region, sc2);
        byte[] sc4 = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("sls", sc3);
        return com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
    }

    public String getSignatureV4(com.aliyun.gateway.spi.models.InterceptorContext context, String signatureAlgorithm, String signedHeaderStr, String date, String scope, String contentSha256, byte[] signingkey) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextRequest request = context.request;
        String canonicalURI = "/";
        if (!com.aliyun.teautil.Common.empty(request.pathname)) {
            canonicalURI = request.pathname;
        }

        String resources = this.buildCanonicalizedResourceV4(request.query);
        String headers = this.buildCanonicalizedHeadersV4(request.headers, signedHeaderStr);
        String stringToHash = "" + request.method + "\n" + canonicalURI + "\n" + resources + "\n" + headers + "\n" + signedHeaderStr + "\n" + contentSha256 + "";
        String hex = com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.encode.Encoder.hash(com.aliyun.teautil.Common.toBytes(stringToHash), signatureAlgorithm));
        String stringToSign = "" + signatureAlgorithm + "\n" + date + "\n" + scope + "\n" + hex + "";
        byte[] signature = com.aliyun.darabonba.signature.Signer.HmacSHA256SignByBytes(stringToSign, signingkey);
        return com.aliyun.darabonba.encode.Encoder.hexEncode(signature);
    }

    public String buildCanonicalizedResourceV4(java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = "";
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryArray = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedQueryArray = com.aliyun.darabonba.array.Client.ascSort(queryArray);
            String separator = "";
            for (String key : sortedQueryArray) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + key + "";
                if (!com.aliyun.teautil.Common.empty(query.get(key))) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + com.aliyun.darabonba.encode.Encoder.percentEncode(query.get(key)) + "";
                }

                separator = "&";
            }
        }

        return canonicalizedResource;
    }

    public String buildCanonicalizedHeadersV4(java.util.Map<String, String> headers, String signedHeaderStr) throws Exception {
        String canonicalizedHeaders = "";
        java.util.List<String> signedHeaders = com.aliyun.darabonbastring.Client.split(signedHeaderStr, ";", null);
        for (String header : signedHeaders) {
            canonicalizedHeaders = "" + canonicalizedHeaders + "" + header + ":" + com.aliyun.darabonbastring.Client.trim(headers.get(header)) + "\n";
        }
        return canonicalizedHeaders;
    }

    public String getSignedHeaderStrV4(java.util.Map<String, String> headers) throws Exception {
        java.util.List<String> headersArray = com.aliyun.darabonba.map.Client.keySet(headers);
        java.util.List<String> sortedHeadersArray = com.aliyun.darabonba.array.Client.ascSort(headersArray);
        String tmp = "";
        String separator = "";
        for (String key : sortedHeadersArray) {
            String lowerKey = com.aliyun.darabonbastring.Client.toLower(key);
            if (com.aliyun.darabonbastring.Client.hasPrefix(lowerKey, "x-log-") || com.aliyun.darabonbastring.Client.hasPrefix(lowerKey, "x-acs-") || com.aliyun.darabonbastring.Client.equals(lowerKey, "host") || com.aliyun.darabonbastring.Client.equals(lowerKey, "content-type")) {
                tmp = "" + tmp + "" + separator + "" + lowerKey + "";
                separator = ";";
            }

        }
        return tmp;
    }

    public String getRegion(com.aliyun.gateway.spi.models.InterceptorContext context) throws Exception {
        com.aliyun.gateway.spi.models.InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        if (!com.aliyun.teautil.Common.isUnset(config.regionId) && !com.aliyun.teautil.Common.empty(config.regionId)) {
            return config.regionId;
        }

        // try parse region from endpoint
        // do not use String.subString, subString has bug in php implementation
        String region = com.aliyun.darabonbastring.Client.replace(config.endpoint, ".log.aliyuncs.com", "", null);
        region = com.aliyun.darabonbastring.Client.replace(region, ".sls.aliyuncs.com", "", null);
        if (com.aliyun.darabonbastring.Client.equals(region, config.endpoint)) {
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", "ClientConfigError"),
                new TeaPair("message", "The regionId configuration of sls client is missing.")
            ));
        }

        region = com.aliyun.darabonbastring.Client.replace(region, "-share", "", null);
        region = com.aliyun.darabonbastring.Client.replace(region, "-intranet", "", null);
        region = com.aliyun.darabonbastring.Client.replace(region, "-vpc", "", null);
        return region;
    }

    // format: YYYYMMDDTHHMMSSZ
    public String getDateISO8601() throws Exception {
        String date = com.aliyun.openapiutil.Client.getTimestamp();
        // 2024-02-04T11:31:58Z
        date = com.aliyun.darabonbastring.Client.replace(date, "-", "", null);
        return com.aliyun.darabonbastring.Client.replace(date, ":", "", null);
    }
}
