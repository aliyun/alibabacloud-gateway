// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class Retention extends TeaModel {
    @NameInMap("Mode")
    public String mode;

    @NameInMap("RetainUtilDateInMs")
    public String retainUtilDateInMs;

    public static Retention build(java.util.Map<String, ?> map) throws Exception {
        Retention self = new Retention();
        return TeaModel.build(map, self);
    }

    public Retention setMode(String mode) {
        this.mode = mode;
        return this;
    }
    public String getMode() {
        return this.mode;
    }

    public Retention setRetainUtilDateInMs(String retainUtilDateInMs) {
        this.retainUtilDateInMs = retainUtilDateInMs;
        return this;
    }
    public String getRetainUtilDateInMs() {
        return this.retainUtilDateInMs;
    }

}
