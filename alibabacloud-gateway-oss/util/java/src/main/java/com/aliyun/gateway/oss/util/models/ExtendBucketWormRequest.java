// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ExtendBucketWormRequest extends TeaModel {
    /**
     * <p>The parameters for WORM extension.</p>
     */
    @NameInMap("ExtendWormConfiguration")
    public ExtendWormConfiguration extendWormConfiguration;

    /**
     * <p>The ID of the retention policy.</p>
     * <blockquote>
     * <p> If the ID of the retention policy that specifies the number of days for which objects can be retained does not exist, the HTTP status code 404 is returned.</p>
     * </blockquote>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>1666E2CFB2B3418****</p>
     */
    @NameInMap("wormId")
    public String wormId;

    public static ExtendBucketWormRequest build(java.util.Map<String, ?> map) throws Exception {
        ExtendBucketWormRequest self = new ExtendBucketWormRequest();
        return TeaModel.build(map, self);
    }

    public ExtendBucketWormRequest setExtendWormConfiguration(ExtendWormConfiguration extendWormConfiguration) {
        this.extendWormConfiguration = extendWormConfiguration;
        return this;
    }
    public ExtendWormConfiguration getExtendWormConfiguration() {
        return this.extendWormConfiguration;
    }

    public ExtendBucketWormRequest setWormId(String wormId) {
        this.wormId = wormId;
        return this;
    }
    public String getWormId() {
        return this.wormId;
    }

}
