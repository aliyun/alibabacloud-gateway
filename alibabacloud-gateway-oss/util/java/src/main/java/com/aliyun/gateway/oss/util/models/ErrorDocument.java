// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ErrorDocument extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>200</p>
     */
    @NameInMap("HttpStatus")
    public Long httpStatus;

    /**
     * <strong>example:</strong>
     * <p>error.html</p>
     */
    @NameInMap("Key")
    public String key;

    public static ErrorDocument build(java.util.Map<String, ?> map) throws Exception {
        ErrorDocument self = new ErrorDocument();
        return TeaModel.build(map, self);
    }

    public ErrorDocument setHttpStatus(Long httpStatus) {
        this.httpStatus = httpStatus;
        return this;
    }
    public Long getHttpStatus() {
        return this.httpStatus;
    }

    public ErrorDocument setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

}
