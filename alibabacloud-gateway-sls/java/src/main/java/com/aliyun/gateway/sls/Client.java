// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.sls;

import com.aliyun.tea.*;
import com.aliyun.gateway.spi.*;
import com.aliyun.gateway.spi.models.*;
import com.aliyun.credentials.*;
import com.aliyun.teautil.*;
import com.aliyun.openapiutil.*;
import com.aliyun.darabonbastring.*;
import com.aliyun.darabonba.map.*;
import com.aliyun.darabonba.array.*;
import com.aliyun.darabonba.encode.*;
import com.aliyun.darabonba.signature.*;

public class Client extends com.aliyun.gateway.spi.Client {

    public Client() throws Exception {
        super();
    }


    public void modifyConfiguration(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        config.endpoint = this.getEndpoint(config.regionId, config.network, config.endpoint);
    }

    public void modifyRequest(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        java.util.Map<String, String> hostMap = new java.util.HashMap<>();
        if (com.aliyun.teautil.Common.isUnset(request.hostMap)) {
            hostMap = request.hostMap;
        }

        String project = hostMap.get("project");
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        com.aliyun.credentials.Client credential = request.credential;
        String accessKeyId = credential.getAccessKeyId();
        String accessKeySecret = credential.getAccessKeySecret();
        String securityToken = credential.getSecurityToken();
        if (!com.aliyun.teautil.Common.empty(accessKeyId)) {
            request.headers.put("x-log-signaturemethod", "hmac-sha1");
        }

        if (!com.aliyun.teautil.Common.empty(securityToken)) {
            request.headers.put("x-acs-security-token", securityToken);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.body)) {
            if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "protobuf")) {
                java.util.Map<String, Object> bodyMap = com.aliyun.teautil.Common.assertAsMap(request.body);
                // 缺少body的Content-MD5计算，以及protobuf处理
                request.headers.put("content-type", "application/x-protobuf");
            } else if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "json")) {
                String bodyStr = com.aliyun.teautil.Common.toJSONString(request.body);
                request.headers.put("content-md5", com.aliyun.darabonbastring.Client.toUpper(com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.signature.Signer.MD5Sign(bodyStr))));
                request.stream = Tea.toReadable(bodyStr);
                request.headers.put("content-type", "application/json");
            } else if (com.aliyun.darabonbastring.Client.equals(request.reqBodyType, "formData")) {
                String str = com.aliyun.teautil.Common.toJSONString(request.body);
                request.headers.put("content-md5", com.aliyun.darabonbastring.Client.toUpper(com.aliyun.darabonba.encode.Encoder.hexEncode(com.aliyun.darabonba.signature.Signer.MD5Sign(str))));
                request.stream = Tea.toReadable(str);
                request.headers.put("content-type", "application/json");
            }

        }

        request.headers = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("accept", "application/json"),
                new TeaPair("host", this.getHost(config.network, project, config.endpoint)),
                new TeaPair("date", com.aliyun.teautil.Common.getDateUTCString()),
                new TeaPair("user-agent", request.userAgent),
                new TeaPair("x-log-apiversion", "0.6.0"),
                new TeaPair("x-log-bodyrawsize", "0")
            ),
            request.headers
        );
        request.headers.put("authorization", this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret));
    }

    public void modifyResponse(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            Object error = com.aliyun.teautil.Common.readAsJSON(response.body);
            java.util.Map<String, Object> resMap = com.aliyun.teautil.Common.assertAsMap(error);
            throw new TeaException(TeaConverter.buildMap(
                new TeaPair("code", resMap.get("errorCode")),
                new TeaPair("message", resMap.get("errorMessage")),
                new TeaPair("data", TeaConverter.buildMap(
                    new TeaPair("httpCode", response.statusCode),
                    new TeaPair("requestId", resMap.get("x-log-requestid"))
                ))
            ));
        }

        if (!com.aliyun.teautil.Common.isUnset(response.body)) {
            if (com.aliyun.teautil.Common.equalString(request.bodyType, "binary")) {
                response.deserializedBody = response.body;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "byte")) {
                byte[] byt = com.aliyun.teautil.Common.readAsBytes(response.body);
                response.deserializedBody = byt;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "string")) {
                response.deserializedBody = com.aliyun.teautil.Common.readAsString(response.body);
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "json")) {
                Object obj = com.aliyun.teautil.Common.readAsJSON(response.body);
                java.util.Map<String, Object> res = com.aliyun.teautil.Common.assertAsMap(obj);
                response.deserializedBody = res;
            } else if (com.aliyun.teautil.Common.equalString(request.bodyType, "array")) {
                response.deserializedBody = com.aliyun.teautil.Common.readAsJSON(response.body);
            } else {
                response.deserializedBody = com.aliyun.teautil.Common.readAsString(response.body);
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
        return "LOG " + ak + ":" + this.getSignature(pathname, method, query, headers, secret) + "";
    }

    public String getSignature(String pathname, String method, java.util.Map<String, String> query, java.util.Map<String, String> headers, String secret) throws Exception {
        String resource = pathname;
        String stringToSign = "";
        String canonicalizedResource = this.buildCanonicalizedResource(resource, query);
        String canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
        stringToSign = "" + method + "\n" + canonicalizedHeaders + "" + canonicalizedResource + "";
        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(com.aliyun.darabonba.signature.Signer.HmacSHA1Sign(stringToSign, secret));
    }

    public String buildCanonicalizedResource(String pathname, java.util.Map<String, String> query) throws Exception {
        String canonicalizedResource = pathname;
        if (!com.aliyun.teautil.Common.isUnset(query)) {
            java.util.List<String> queryList = com.aliyun.darabonba.map.Client.keySet(query);
            java.util.List<String> sortedParams = com.aliyun.darabonba.array.Client.ascSort(queryList);
            String separator = "?";
            for (String paramName : sortedParams) {
                canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + paramName + "";
                if (!com.aliyun.teautil.Common.isUnset(query.get(paramName))) {
                    canonicalizedResource = "" + canonicalizedResource + "=" + query.get(paramName) + "";
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
}
