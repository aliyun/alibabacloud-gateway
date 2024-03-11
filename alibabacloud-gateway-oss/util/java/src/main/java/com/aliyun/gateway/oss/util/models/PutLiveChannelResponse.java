// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class PutLiveChannelResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PutLiveChannelResponseBody body;

    public static PutLiveChannelResponse build(java.util.Map<String, ?> map) throws Exception {
        PutLiveChannelResponse self = new PutLiveChannelResponse();
        return TeaModel.build(map, self);
    }

    public PutLiveChannelResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutLiveChannelResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutLiveChannelResponse setBody(PutLiveChannelResponseBody body) {
        this.body = body;
        return this;
    }
    public PutLiveChannelResponseBody getBody() {
        return this.body;
    }

}
