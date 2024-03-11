// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListObjectsV2Response extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListObjectsV2ResponseBody body;

    public static ListObjectsV2Response build(java.util.Map<String, ?> map) throws Exception {
        ListObjectsV2Response self = new ListObjectsV2Response();
        return TeaModel.build(map, self);
    }

    public ListObjectsV2Response setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListObjectsV2Response setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListObjectsV2Response setBody(ListObjectsV2ResponseBody body) {
        this.body = body;
        return this;
    }
    public ListObjectsV2ResponseBody getBody() {
        return this.body;
    }

}
