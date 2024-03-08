// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListCnameResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListCnameResponseBody body;

    public static ListCnameResponse build(java.util.Map<String, ?> map) throws Exception {
        ListCnameResponse self = new ListCnameResponse();
        return TeaModel.build(map, self);
    }

    public ListCnameResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListCnameResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListCnameResponse setBody(ListCnameResponseBody body) {
        this.body = body;
        return this;
    }
    public ListCnameResponseBody getBody() {
        return this.body;
    }

}
