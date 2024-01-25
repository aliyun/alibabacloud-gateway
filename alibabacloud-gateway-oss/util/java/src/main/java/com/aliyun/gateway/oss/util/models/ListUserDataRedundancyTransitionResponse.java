// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListUserDataRedundancyTransitionResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    @NameInMap("body")
    public ListUserDataRedundancyTransitionResponseBody body;

    public static ListUserDataRedundancyTransitionResponse build(java.util.Map<String, ?> map) throws Exception {
        ListUserDataRedundancyTransitionResponse self = new ListUserDataRedundancyTransitionResponse();
        return TeaModel.build(map, self);
    }

    public ListUserDataRedundancyTransitionResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListUserDataRedundancyTransitionResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

    public ListUserDataRedundancyTransitionResponse setBody(ListUserDataRedundancyTransitionResponseBody body) {
        this.body = body;
        return this;
    }
    public ListUserDataRedundancyTransitionResponseBody getBody() {
        return this.body;
    }

}
