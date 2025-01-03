// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketCommonHeaderResponseBody extends TeaModel {
    @NameInMap("CommonHeaders")
    public CommonHeaders commonHeaders;

    public static GetBucketCommonHeaderResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCommonHeaderResponseBody self = new GetBucketCommonHeaderResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketCommonHeaderResponseBody setCommonHeaders(CommonHeaders commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public CommonHeaders getCommonHeaders() {
        return this.commonHeaders;
    }

}
