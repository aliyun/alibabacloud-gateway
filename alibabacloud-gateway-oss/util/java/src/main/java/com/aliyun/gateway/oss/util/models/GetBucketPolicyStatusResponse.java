// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketPolicyStatusResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketPolicyStatusResponseBody body;

    public static GetBucketPolicyStatusResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketPolicyStatusResponse self = new GetBucketPolicyStatusResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketPolicyStatusResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketPolicyStatusResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketPolicyStatusResponse setBody(GetBucketPolicyStatusResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketPolicyStatusResponseBody getBody() {
        return this.body;
    }

}
