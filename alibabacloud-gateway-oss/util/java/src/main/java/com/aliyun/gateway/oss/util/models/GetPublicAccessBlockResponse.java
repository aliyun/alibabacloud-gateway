// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetPublicAccessBlockResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PublicAccessBlockConfiguration body;

    public static GetPublicAccessBlockResponse build(java.util.Map<String, ?> map) throws Exception {
        GetPublicAccessBlockResponse self = new GetPublicAccessBlockResponse();
        return TeaModel.build(map, self);
    }

    public GetPublicAccessBlockResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetPublicAccessBlockResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetPublicAccessBlockResponse setBody(PublicAccessBlockConfiguration body) {
        this.body = body;
        return this;
    }
    public PublicAccessBlockConfiguration getBody() {
        return this.body;
    }

}
