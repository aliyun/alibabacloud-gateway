// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class UpdateBucketAntiDDosInfoRequest extends TeaModel {
    /**
     * <p>The container that stores the configurations of Anti-DDoS instances.</p>
     */
    @NameInMap("AntiDDOSConfiguration")
    public BucketAntiDDOSConfiguration antiDDOSConfiguration;

    public static UpdateBucketAntiDDosInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        UpdateBucketAntiDDosInfoRequest self = new UpdateBucketAntiDDosInfoRequest();
        return TeaModel.build(map, self);
    }

    public UpdateBucketAntiDDosInfoRequest setAntiDDOSConfiguration(BucketAntiDDOSConfiguration antiDDOSConfiguration) {
        this.antiDDOSConfiguration = antiDDOSConfiguration;
        return this;
    }
    public BucketAntiDDOSConfiguration getAntiDDOSConfiguration() {
        return this.antiDDOSConfiguration;
    }

}
