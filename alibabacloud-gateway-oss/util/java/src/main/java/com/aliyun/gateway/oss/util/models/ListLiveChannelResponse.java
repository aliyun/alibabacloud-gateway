// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListLiveChannelResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListLiveChannelResponseBody body;

    public static ListLiveChannelResponse build(java.util.Map<String, ?> map) throws Exception {
        ListLiveChannelResponse self = new ListLiveChannelResponse();
        return TeaModel.build(map, self);
    }

    public ListLiveChannelResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListLiveChannelResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListLiveChannelResponse setBody(ListLiveChannelResponseBody body) {
        this.body = body;
        return this;
    }
    public ListLiveChannelResponseBody getBody() {
        return this.body;
    }

}
