// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListObjectVersionsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListObjectVersionsResponseBody body;

    public static ListObjectVersionsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListObjectVersionsResponse self = new ListObjectVersionsResponse();
        return TeaModel.build(map, self);
    }

    public ListObjectVersionsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListObjectVersionsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListObjectVersionsResponse setBody(ListObjectVersionsResponseBody body) {
        this.body = body;
        return this;
    }
    public ListObjectVersionsResponseBody getBody() {
        return this.body;
    }

}
