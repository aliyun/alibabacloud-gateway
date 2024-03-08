// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketCallbackPolicyResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CallbackPolicy body;

    public static GetBucketCallbackPolicyResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCallbackPolicyResponse self = new GetBucketCallbackPolicyResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketCallbackPolicyResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketCallbackPolicyResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketCallbackPolicyResponse setBody(CallbackPolicy body) {
        this.body = body;
        return this;
    }
    public CallbackPolicy getBody() {
        return this.body;
    }

}
