// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ResourcePoolBucketGroup extends TeaModel {
    @NameInMap("GroupBucketInfo")
    public java.util.List<GroupBucketInfo> groupBucketInfo;

    /**
     * <strong>example:</strong>
     * <p>test-group-1</p>
     */
    @NameInMap("Name")
    public String name;

    public static ResourcePoolBucketGroup build(java.util.Map<String, ?> map) throws Exception {
        ResourcePoolBucketGroup self = new ResourcePoolBucketGroup();
        return TeaModel.build(map, self);
    }

    public ResourcePoolBucketGroup setGroupBucketInfo(java.util.List<GroupBucketInfo> groupBucketInfo) {
        this.groupBucketInfo = groupBucketInfo;
        return this;
    }
    public java.util.List<GroupBucketInfo> getGroupBucketInfo() {
        return this.groupBucketInfo;
    }

    public ResourcePoolBucketGroup setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

}
