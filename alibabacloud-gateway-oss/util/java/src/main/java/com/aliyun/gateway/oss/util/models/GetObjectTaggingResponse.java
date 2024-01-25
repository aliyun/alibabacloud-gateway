// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetObjectTaggingResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetObjectTaggingResponseBody body;

    public static GetObjectTaggingResponse build(java.util.Map<String, ?> map) throws Exception {
        GetObjectTaggingResponse self = new GetObjectTaggingResponse();
        return TeaModel.build(map, self);
    }

    public GetObjectTaggingResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetObjectTaggingResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetObjectTaggingResponse setBody(GetObjectTaggingResponseBody body) {
        this.body = body;
        return this;
    }
    public GetObjectTaggingResponseBody getBody() {
        return this.body;
    }

}
