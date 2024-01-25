// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetAccessPointPublicAccessBlockResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PublicAccessBlockConfiguration body;

    public static GetAccessPointPublicAccessBlockResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointPublicAccessBlockResponse self = new GetAccessPointPublicAccessBlockResponse();
        return TeaModel.build(map, self);
    }

    public GetAccessPointPublicAccessBlockResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAccessPointPublicAccessBlockResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetAccessPointPublicAccessBlockResponse setBody(PublicAccessBlockConfiguration body) {
        this.body = body;
        return this;
    }
    public PublicAccessBlockConfiguration getBody() {
        return this.body;
    }

}
