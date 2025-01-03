// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ResourcePoolBucket extends TeaModel {
    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2024-07-24T08:42:32.000Z</p>
     */
    @NameInMap("JoinTime")
    public String joinTime;

    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Name")
    public String name;

    public static ResourcePoolBucket build(java.util.Map<String, ?> map) throws Exception {
        ResourcePoolBucket self = new ResourcePoolBucket();
        return TeaModel.build(map, self);
    }

    public ResourcePoolBucket setJoinTime(String joinTime) {
        this.joinTime = joinTime;
        return this;
    }
    public String getJoinTime() {
        return this.joinTime;
    }

    public ResourcePoolBucket setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

}
