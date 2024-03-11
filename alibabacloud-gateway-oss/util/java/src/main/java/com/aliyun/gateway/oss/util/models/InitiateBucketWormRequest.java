// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class InitiateBucketWormRequest extends TeaModel {
    /**
     * <p>The parameters for WORM initialization.</p>
     */
    @NameInMap("InitiateWormConfiguration")
    public InitiateWormConfiguration initiateWormConfiguration;

    public static InitiateBucketWormRequest build(java.util.Map<String, ?> map) throws Exception {
        InitiateBucketWormRequest self = new InitiateBucketWormRequest();
        return TeaModel.build(map, self);
    }

    public InitiateBucketWormRequest setInitiateWormConfiguration(InitiateWormConfiguration initiateWormConfiguration) {
        this.initiateWormConfiguration = initiateWormConfiguration;
        return this;
    }
    public InitiateWormConfiguration getInitiateWormConfiguration() {
        return this.initiateWormConfiguration;
    }

}
