// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketCommonHeaderRequest extends TeaModel {
    /**
     * <p>User-defined response headers configuration</p>
     */
    @NameInMap("CommonHeaders")
    public CommonHeaders commonHeaders;

    public static PutBucketCommonHeaderRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCommonHeaderRequest self = new PutBucketCommonHeaderRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketCommonHeaderRequest setCommonHeaders(CommonHeaders commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public CommonHeaders getCommonHeaders() {
        return this.commonHeaders;
    }

}
