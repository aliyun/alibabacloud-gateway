// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutChannelHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("bucket")
    public String bucket;

    public static PutChannelHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutChannelHeaders self = new PutChannelHeaders();
        return TeaModel.build(map, self);
    }

    public PutChannelHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutChannelHeaders setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

}
