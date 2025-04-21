// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ResourcePoolBucketGroupQoSInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-group</p>
     */
    @NameInMap("GroupName")
    public String groupName;

    @NameInMap("QoSConfiguration")
    public QoSConfiguration qoSConfiguration;

    public static ResourcePoolBucketGroupQoSInfo build(java.util.Map<String, ?> map) throws Exception {
        ResourcePoolBucketGroupQoSInfo self = new ResourcePoolBucketGroupQoSInfo();
        return TeaModel.build(map, self);
    }

    public ResourcePoolBucketGroupQoSInfo setGroupName(String groupName) {
        this.groupName = groupName;
        return this;
    }
    public String getGroupName() {
        return this.groupName;
    }

    public ResourcePoolBucketGroupQoSInfo setQoSConfiguration(QoSConfiguration qoSConfiguration) {
        this.qoSConfiguration = qoSConfiguration;
        return this;
    }
    public QoSConfiguration getQoSConfiguration() {
        return this.qoSConfiguration;
    }

}
