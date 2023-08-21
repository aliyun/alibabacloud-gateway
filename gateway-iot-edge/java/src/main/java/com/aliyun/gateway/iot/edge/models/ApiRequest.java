// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.iot.edge.models;

import com.aliyun.tea.*;

public class ApiRequest extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("pathname")
    public String pathname;

    @NameInMap("body")
    public java.util.Map<String, ?> body;

    @NameInMap("action")
    public String action;

    public static ApiRequest build(java.util.Map<String, ?> map) throws Exception {
        ApiRequest self = new ApiRequest();
        return TeaModel.build(map, self);
    }

    public ApiRequest setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ApiRequest setPathname(String pathname) {
        this.pathname = pathname;
        return this;
    }
    public String getPathname() {
        return this.pathname;
    }

    public ApiRequest setBody(java.util.Map<String, ?> body) {
        this.body = body;
        return this;
    }
    public java.util.Map<String, ?> getBody() {
        return this.body;
    }

    public ApiRequest setAction(String action) {
        this.action = action;
        return this;
    }
    public String getAction() {
        return this.action;
    }

}
