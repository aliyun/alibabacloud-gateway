// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteUserDefinedLogFieldsConfigResponse extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("statusCode")
    public Integer statusCode;

    public static DeleteUserDefinedLogFieldsConfigResponse build(java.util.Map<String, ?> map) throws Exception {
        DeleteUserDefinedLogFieldsConfigResponse self = new DeleteUserDefinedLogFieldsConfigResponse();
        return TeaModel.build(map, self);
    }

    public DeleteUserDefinedLogFieldsConfigResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public DeleteUserDefinedLogFieldsConfigResponse setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public Integer getStatusCode() {
        return this.statusCode;
    }

}
