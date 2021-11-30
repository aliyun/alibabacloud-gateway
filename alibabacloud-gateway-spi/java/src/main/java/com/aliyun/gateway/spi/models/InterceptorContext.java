// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.spi.models;

import com.aliyun.tea.*;

public class InterceptorContext extends TeaModel {
    @NameInMap("request")
    @Validation(required = true)
    public InterceptorContextRequest request;

    @NameInMap("configuration")
    @Validation(required = true)
    public InterceptorContextConfiguration configuration;

    @NameInMap("response")
    @Validation(required = true)
    public InterceptorContextResponse response;

    public static InterceptorContext build(java.util.Map<String, ?> map) {
        InterceptorContext self = new InterceptorContext();
        return TeaModel.build(map, self);
    }

    public InterceptorContext setRequest(InterceptorContextRequest request) {
        this.request = request;
        return this;
    }
    public InterceptorContextRequest getRequest() {
        return this.request;
    }

    public InterceptorContext setConfiguration(InterceptorContextConfiguration configuration) {
        this.configuration = configuration;
        return this;
    }
    public InterceptorContextConfiguration getConfiguration() {
        return this.configuration;
    }

    public InterceptorContext setResponse(InterceptorContextResponse response) {
        this.response = response;
        return this;
    }
    public InterceptorContextResponse getResponse() {
        return this.response;
    }

    public static class InterceptorContextRequest extends TeaModel {
        @NameInMap("headers")
        public java.util.Map<String, String> headers;

        @NameInMap("query")
        public java.util.Map<String, String> query;

        @NameInMap("body")
        public Object body;

        @NameInMap("stream")
        public java.io.InputStream stream;

        @NameInMap("hostMap")
        public java.util.Map<String, String> hostMap;

        @NameInMap("pathname")
        @Validation(required = true)
        public String pathname;

        @NameInMap("productId")
        @Validation(required = true)
        public String productId;

        @NameInMap("action")
        @Validation(required = true)
        public String action;

        @NameInMap("version")
        @Validation(required = true)
        public String version;

        @NameInMap("protocol")
        @Validation(required = true)
        public String protocol;

        @NameInMap("method")
        @Validation(required = true)
        public String method;

        @NameInMap("authType")
        @Validation(required = true)
        public String authType;

        @NameInMap("bodyType")
        @Validation(required = true)
        public String bodyType;

        @NameInMap("reqBodyType")
        @Validation(required = true)
        public String reqBodyType;

        @NameInMap("style")
        public String style;

        @NameInMap("credential")
        @Validation(required = true)
        public com.aliyun.credentials.Client credential;

        @NameInMap("signatureVersion")
        public String signatureVersion;

        @NameInMap("signatureAlgorithm")
        public String signatureAlgorithm;

        @NameInMap("userAgent")
        @Validation(required = true)
        public String userAgent;

        public static InterceptorContextRequest build(java.util.Map<String, ?> map) {
            InterceptorContextRequest self = new InterceptorContextRequest();
            return TeaModel.build(map, self);
        }

        public InterceptorContextRequest setHeaders(java.util.Map<String, String> headers) {
            this.headers = headers;
            return this;
        }
        public java.util.Map<String, String> getHeaders() {
            return this.headers;
        }

        public InterceptorContextRequest setQuery(java.util.Map<String, String> query) {
            this.query = query;
            return this;
        }
        public java.util.Map<String, String> getQuery() {
            return this.query;
        }

        public InterceptorContextRequest setBody(Object body) {
            this.body = body;
            return this;
        }
        public Object getBody() {
            return this.body;
        }

        public InterceptorContextRequest setStream(java.io.InputStream stream) {
            this.stream = stream;
            return this;
        }
        public java.io.InputStream getStream() {
            return this.stream;
        }

        public InterceptorContextRequest setHostMap(java.util.Map<String, String> hostMap) {
            this.hostMap = hostMap;
            return this;
        }
        public java.util.Map<String, String> getHostMap() {
            return this.hostMap;
        }

        public InterceptorContextRequest setPathname(String pathname) {
            this.pathname = pathname;
            return this;
        }
        public String getPathname() {
            return this.pathname;
        }

        public InterceptorContextRequest setProductId(String productId) {
            this.productId = productId;
            return this;
        }
        public String getProductId() {
            return this.productId;
        }

        public InterceptorContextRequest setAction(String action) {
            this.action = action;
            return this;
        }
        public String getAction() {
            return this.action;
        }

        public InterceptorContextRequest setVersion(String version) {
            this.version = version;
            return this;
        }
        public String getVersion() {
            return this.version;
        }

        public InterceptorContextRequest setProtocol(String protocol) {
            this.protocol = protocol;
            return this;
        }
        public String getProtocol() {
            return this.protocol;
        }

        public InterceptorContextRequest setMethod(String method) {
            this.method = method;
            return this;
        }
        public String getMethod() {
            return this.method;
        }

        public InterceptorContextRequest setAuthType(String authType) {
            this.authType = authType;
            return this;
        }
        public String getAuthType() {
            return this.authType;
        }

        public InterceptorContextRequest setBodyType(String bodyType) {
            this.bodyType = bodyType;
            return this;
        }
        public String getBodyType() {
            return this.bodyType;
        }

        public InterceptorContextRequest setReqBodyType(String reqBodyType) {
            this.reqBodyType = reqBodyType;
            return this;
        }
        public String getReqBodyType() {
            return this.reqBodyType;
        }

        public InterceptorContextRequest setStyle(String style) {
            this.style = style;
            return this;
        }
        public String getStyle() {
            return this.style;
        }

        public InterceptorContextRequest setCredential(com.aliyun.credentials.Client credential) {
            this.credential = credential;
            return this;
        }
        public com.aliyun.credentials.Client getCredential() {
            return this.credential;
        }

        public InterceptorContextRequest setSignatureVersion(String signatureVersion) {
            this.signatureVersion = signatureVersion;
            return this;
        }
        public String getSignatureVersion() {
            return this.signatureVersion;
        }

        public InterceptorContextRequest setSignatureAlgorithm(String signatureAlgorithm) {
            this.signatureAlgorithm = signatureAlgorithm;
            return this;
        }
        public String getSignatureAlgorithm() {
            return this.signatureAlgorithm;
        }

        public InterceptorContextRequest setUserAgent(String userAgent) {
            this.userAgent = userAgent;
            return this;
        }
        public String getUserAgent() {
            return this.userAgent;
        }

    }

    public static class InterceptorContextConfiguration extends TeaModel {
        @NameInMap("regionId")
        @Validation(required = true)
        public String regionId;

        @NameInMap("endpoint")
        public String endpoint;

        @NameInMap("endpointRule")
        public String endpointRule;

        @NameInMap("endpointMap")
        public java.util.Map<String, String> endpointMap;

        @NameInMap("endpointType")
        public String endpointType;

        @NameInMap("network")
        public String network;

        @NameInMap("suffix")
        public String suffix;

        public static InterceptorContextConfiguration build(java.util.Map<String, ?> map) {
            InterceptorContextConfiguration self = new InterceptorContextConfiguration();
            return TeaModel.build(map, self);
        }

        public InterceptorContextConfiguration setRegionId(String regionId) {
            this.regionId = regionId;
            return this;
        }
        public String getRegionId() {
            return this.regionId;
        }

        public InterceptorContextConfiguration setEndpoint(String endpoint) {
            this.endpoint = endpoint;
            return this;
        }
        public String getEndpoint() {
            return this.endpoint;
        }

        public InterceptorContextConfiguration setEndpointRule(String endpointRule) {
            this.endpointRule = endpointRule;
            return this;
        }
        public String getEndpointRule() {
            return this.endpointRule;
        }

        public InterceptorContextConfiguration setEndpointMap(java.util.Map<String, String> endpointMap) {
            this.endpointMap = endpointMap;
            return this;
        }
        public java.util.Map<String, String> getEndpointMap() {
            return this.endpointMap;
        }

        public InterceptorContextConfiguration setEndpointType(String endpointType) {
            this.endpointType = endpointType;
            return this;
        }
        public String getEndpointType() {
            return this.endpointType;
        }

        public InterceptorContextConfiguration setNetwork(String network) {
            this.network = network;
            return this;
        }
        public String getNetwork() {
            return this.network;
        }

        public InterceptorContextConfiguration setSuffix(String suffix) {
            this.suffix = suffix;
            return this;
        }
        public String getSuffix() {
            return this.suffix;
        }

    }

    public static class InterceptorContextResponse extends TeaModel {
        @NameInMap("statusCode")
        public Number statusCode;

        @NameInMap("headers")
        public java.util.Map<String, String> headers;

        @NameInMap("body")
        public java.io.InputStream body;

        @NameInMap("deserializedBody")
        public Object deserializedBody;

        public static InterceptorContextResponse build(java.util.Map<String, ?> map) {
            InterceptorContextResponse self = new InterceptorContextResponse();
            return TeaModel.build(map, self);
        }

        public InterceptorContextResponse setStatusCode(Number statusCode) {
            this.statusCode = statusCode;
            return this;
        }
        public Number getStatusCode() {
            return this.statusCode;
        }

        public InterceptorContextResponse setHeaders(java.util.Map<String, String> headers) {
            this.headers = headers;
            return this;
        }
        public java.util.Map<String, String> getHeaders() {
            return this.headers;
        }

        public InterceptorContextResponse setBody(java.io.InputStream body) {
            this.body = body;
            return this;
        }
        public java.io.InputStream getBody() {
            return this.body;
        }

        public InterceptorContextResponse setDeserializedBody(Object deserializedBody) {
            this.deserializedBody = deserializedBody;
            return this;
        }
        public Object getDeserializedBody() {
            return this.deserializedBody;
        }

    }

}
