// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketAccessMonitorResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetBucketAccessMonitorResponseBody body;

    public static GetBucketAccessMonitorResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketAccessMonitorResponse self = new GetBucketAccessMonitorResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketAccessMonitorResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketAccessMonitorResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketAccessMonitorResponse setBody(GetBucketAccessMonitorResponseBody body) {
        this.body = body;
        return this;
    }
    public GetBucketAccessMonitorResponseBody getBody() {
        return this.body;
    }

}
