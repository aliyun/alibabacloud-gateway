// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.pop;

import com.aliyun.tea.*;
import com.aliyun.gateway.spi.*;
import com.aliyun.gateway.spi.models.*;
import com.aliyun.credentials.*;
import com.aliyun.teautil.*;
import com.aliyun.openapiutil.*;
import com.aliyun.endpointutil.*;

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
        com.aliyun.credentials.Client credential = request.credential;
        request.query = TeaConverter.merge(String.class,
            TeaConverter.buildMap(
                new TeaPair("Action", request.action),
                new TeaPair("Format", "json"),
                new TeaPair("Version", request.version),
                new TeaPair("Timestamp", com.aliyun.openapiutil.Client.getTimestamp()),
                new TeaPair("SignatureNonce", com.aliyun.teautil.Common.getNonce())
            ),
            request.query
        );
        request.headers = TeaConverter.buildMap(
            new TeaPair("host", config.endpoint),
            new TeaPair("x-acs-version", request.version),
            new TeaPair("x-acs-action", request.action),
            new TeaPair("user-agent", request.userAgent)
        );
        java.util.Map<String, Object> t = null;
        if (!com.aliyun.teautil.Common.isUnset(request.body)) {
            t = com.aliyun.teautil.Common.assertAsMap(request.body);
            java.util.Map<String, Object> tmp = com.aliyun.teautil.Common.anyifyMapValue(com.aliyun.openapiutil.Client.query(t));
            request.stream = Tea.toReadable(com.aliyun.teautil.Common.toFormString(tmp));
            request.headers.put("content-type", "application/x-www-form-urlencoded");
        }

        if (!com.aliyun.teautil.Common.equalString(request.authType, "Anonymous")) {
            String accessKeyId = credential.getAccessKeyId();
            String accessKeySecret = credential.getAccessKeySecret();
            String securityToken = credential.getSecurityToken();
            if (!com.aliyun.teautil.Common.empty(securityToken)) {
                request.query.put("SecurityToken", securityToken);
            }

            request.query.put("SignatureMethod", "HMAC-SHA1");
            request.query.put("SignatureVersion", "1.0");
            request.query.put("AccessKeyId", accessKeyId);
            java.util.Map<String, String> signedParam = TeaConverter.merge(String.class,
                request.query,
                com.aliyun.openapiutil.Client.query(t)
            );
            request.query.put("Signature", com.aliyun.openapiutil.Client.getRPCSignature(signedParam, request.method, accessKeySecret));
        }

    }

    public void modifyResponse(InterceptorContext context, AttributeMap attributeMap) throws Exception {
        InterceptorContext.InterceptorContextRequest request = context.request;
        InterceptorContext.InterceptorContextConfiguration config = context.configuration;
        InterceptorContext.InterceptorContextResponse response = context.response;
        if (com.aliyun.teautil.Common.is4xx(response.statusCode) || com.aliyun.teautil.Common.is5xx(response.statusCode)) {
            Object _res = com.aliyun.teautil.Common.readAsJSON(response.body);
            java.util.Map<String, Object> err = com.aliyun.teautil.Common.assertAsMap(_res);
            Object requestId = this.defaultAny(err.get("RequestId"), err.get("requestId"));
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
}
