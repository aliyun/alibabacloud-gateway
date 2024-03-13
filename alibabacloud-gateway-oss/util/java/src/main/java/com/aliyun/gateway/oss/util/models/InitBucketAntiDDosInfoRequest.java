// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InitBucketAntiDDosInfoRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of Anti-DDoS instances.</p>
     */
    @NameInMap("AntiDDOSConfiguration")
    public BucketAntiDDOSConfiguration antiDDOSConfiguration;

    public static InitBucketAntiDDosInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        InitBucketAntiDDosInfoRequest self = new InitBucketAntiDDosInfoRequest();
        return TeaModel.build(map, self);
    }

    public InitBucketAntiDDosInfoRequest setAntiDDOSConfiguration(BucketAntiDDOSConfiguration antiDDOSConfiguration) {
        this.antiDDOSConfiguration = antiDDOSConfiguration;
        return this;
    }
    public BucketAntiDDOSConfiguration getAntiDDOSConfiguration() {
        return this.antiDDOSConfiguration;
    }

}
