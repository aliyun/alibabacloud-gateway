// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ResourcePoolSimpleInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>2024-07-24T08:42:32.000Z</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <strong>example:</strong>
     * <p>rp-for-api</p>
     */
    @NameInMap("Name")
    public String name;

    public static ResourcePoolSimpleInfo build(java.util.Map<String, ?> map) throws Exception {
        ResourcePoolSimpleInfo self = new ResourcePoolSimpleInfo();
        return TeaModel.build(map, self);
    }

    public ResourcePoolSimpleInfo setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public ResourcePoolSimpleInfo setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

}
