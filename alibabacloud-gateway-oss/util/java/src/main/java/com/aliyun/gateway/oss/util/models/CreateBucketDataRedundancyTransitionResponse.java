// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateBucketDataRedundancyTransitionResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CreateBucketDataRedundancyTransitionResponseBody body;

    public static CreateBucketDataRedundancyTransitionResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateBucketDataRedundancyTransitionResponse self = new CreateBucketDataRedundancyTransitionResponse();
        return TeaModel.build(map, self);
    }

    public CreateBucketDataRedundancyTransitionResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CreateBucketDataRedundancyTransitionResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CreateBucketDataRedundancyTransitionResponse setBody(CreateBucketDataRedundancyTransitionResponseBody body) {
        this.body = body;
        return this;
    }
    public CreateBucketDataRedundancyTransitionResponseBody getBody() {
        return this.body;
    }

}
