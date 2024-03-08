// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketCallbackPolicyResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CallbackPolicy body;

    public static PutBucketCallbackPolicyResponse build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCallbackPolicyResponse self = new PutBucketCallbackPolicyResponse();
        return TeaModel.build(map, self);
    }

    public PutBucketCallbackPolicyResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutBucketCallbackPolicyResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutBucketCallbackPolicyResponse setBody(CallbackPolicy body) {
        this.body = body;
        return this;
    }
    public CallbackPolicy getBody() {
        return this.body;
    }

}
