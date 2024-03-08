// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketCorsResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketCorsResponseBody body;

    public static GetBucketCorsResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCorsResponse self = new GetBucketCorsResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketCorsResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketCorsResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketCorsResponse setBody(GetBucketCorsResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketCorsResponseBody getBody() {
        return this.body;
    }

}
