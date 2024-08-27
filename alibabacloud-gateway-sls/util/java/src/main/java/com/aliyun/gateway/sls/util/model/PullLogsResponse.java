package com.aliyun.gateway.sls.util.model;

import com.aliyun.tea.NameInMap;
import com.aliyun.tea.TeaModel;

public class PullLogsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PostLogStoreLogsResponse.PullLogsResponseBody body;

    public PostLogStoreLogsResponse.PullLogsResponseBody getBody() {
        return body;
    }
    public void setBody(PostLogStoreLogsResponse.PullLogsResponseBody body) {
        this.body = body;
    }
    public static PullLogsResponse build(java.util.Map<String, ?> map) throws Exception {
        PullLogsResponse self = new PullLogsResponse();
        return TeaModel.build(map, self);
    }

    public PullLogsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PullLogsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }
}
