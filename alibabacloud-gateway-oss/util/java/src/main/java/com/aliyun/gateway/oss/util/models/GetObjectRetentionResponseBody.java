// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class GetObjectRetentionResponseBody extends TeaModel {
    @NameInMap("Retention")
    public Retention retention;

    public static GetObjectRetentionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectRetentionResponseBody self = new GetObjectRetentionResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectRetentionResponseBody setRetention(Retention retention) {
        this.retention = retention;
        return this;
    }
    public Retention getRetention() {
        return this.retention;
    }

}
