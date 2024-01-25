// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketEncryptionResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketEncryptionResponseBody body;

    public static GetBucketEncryptionResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketEncryptionResponse self = new GetBucketEncryptionResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketEncryptionResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketEncryptionResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketEncryptionResponse setBody(GetBucketEncryptionResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketEncryptionResponseBody getBody() {
        return this.body;
    }

}
