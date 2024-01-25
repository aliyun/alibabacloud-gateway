// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateAccessPointResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CreateAccessPointResult body;

    public static CreateAccessPointResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointResponse self = new CreateAccessPointResponse();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public CreateAccessPointResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public CreateAccessPointResponse setBody(CreateAccessPointResult body) {
        this.body = body;
        return this;
    }
    public CreateAccessPointResult getBody() {
        return this.body;
    }

}
