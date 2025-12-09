// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketReplicationBandwidthResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PutBucketReplicationBandwidthResponseBody body;

    public static PutBucketReplicationBandwidthResponse build(java.util.Map<String, ?> map) throws Exception {
        PutBucketReplicationBandwidthResponse self = new PutBucketReplicationBandwidthResponse();
        return TeaModel.build(map, self);
    }

    public PutBucketReplicationBandwidthResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutBucketReplicationBandwidthResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutBucketReplicationBandwidthResponse setBody(PutBucketReplicationBandwidthResponseBody body) {
        this.body = body;
        return this;
    }
    public PutBucketReplicationBandwidthResponseBody getBody() {
        return this.body;
    }

}
