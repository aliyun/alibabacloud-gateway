// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutAccessPointConfigForObjectProcessResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public PutAccessPointConfigForObjectProcessResponseBody body;

    public static PutAccessPointConfigForObjectProcessResponse build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointConfigForObjectProcessResponse self = new PutAccessPointConfigForObjectProcessResponse();
        return TeaModel.build(map, self);
    }

    public PutAccessPointConfigForObjectProcessResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public PutAccessPointConfigForObjectProcessResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public PutAccessPointConfigForObjectProcessResponse setBody(PutAccessPointConfigForObjectProcessResponseBody body) {
        this.body = body;
        return this;
    }
    public PutAccessPointConfigForObjectProcessResponseBody getBody() {
        return this.body;
    }

}
