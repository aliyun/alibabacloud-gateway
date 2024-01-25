// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketInfoResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public BucketInfo body;

    public static GetBucketInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketInfoResponse self = new GetBucketInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketInfoResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketInfoResponse setBody(BucketInfo body) {
        this.body = body;
        return this;
    }
    public BucketInfo getBody() {
        return this.body;
    }

}
