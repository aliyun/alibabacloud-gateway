// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketHttpsConfigResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public HttpsConfiguration body;

    public static GetBucketHttpsConfigResponse build(java.util.Map<String, ?> map) throws Exception {
        GetBucketHttpsConfigResponse self = new GetBucketHttpsConfigResponse();
        return TeaModel.build(map, self);
    }

    public GetBucketHttpsConfigResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetBucketHttpsConfigResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetBucketHttpsConfigResponse setBody(HttpsConfiguration body) {
        this.body = body;
        return this;
    }
    public HttpsConfiguration getBody() {
        return this.body;
    }

}
