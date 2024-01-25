// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetAccessPointResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public GetAccessPointResult body;

    public static GetAccessPointResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointResponse self = new GetAccessPointResponse();
        return TeaModel.build(map, self);
    }

    public GetAccessPointResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAccessPointResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetAccessPointResponse setBody(GetAccessPointResult body) {
        this.body = body;
        return this;
    }
    public GetAccessPointResult getBody() {
        return this.body;
    }

}
