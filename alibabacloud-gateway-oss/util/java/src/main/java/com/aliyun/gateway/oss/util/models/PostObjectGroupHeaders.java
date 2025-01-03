// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostObjectGroupHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-file-group")
    public String xOssFileGroup;

    public static PostObjectGroupHeaders build(java.util.Map<String, ?> map) throws Exception {
        PostObjectGroupHeaders self = new PostObjectGroupHeaders();
        return TeaModel.build(map, self);
    }

    public PostObjectGroupHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PostObjectGroupHeaders setXOssFileGroup(String xOssFileGroup) {
        this.xOssFileGroup = xOssFileGroup;
        return this;
    }
    public String getXOssFileGroup() {
        return this.xOssFileGroup;
    }

}
