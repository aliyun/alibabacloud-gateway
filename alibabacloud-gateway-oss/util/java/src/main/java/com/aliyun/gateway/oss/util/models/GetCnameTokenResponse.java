// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetCnameTokenResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public CnameToken body;

    public static GetCnameTokenResponse build(java.util.Map<String, ?> map) throws Exception {
        GetCnameTokenResponse self = new GetCnameTokenResponse();
        return TeaModel.build(map, self);
    }

    public GetCnameTokenResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetCnameTokenResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public GetCnameTokenResponse setBody(CnameToken body) {
        this.body = body;
        return this;
    }
    public CnameToken getBody() {
        return this.body;
    }

}
