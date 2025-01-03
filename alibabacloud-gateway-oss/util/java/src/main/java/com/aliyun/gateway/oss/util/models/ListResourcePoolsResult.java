// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListResourcePoolsResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>abcd</p>
     */
    @NameInMap("ContionuationToken")
    public String contionuationToken;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <strong>example:</strong>
     * <p>xyz</p>
     */
    @NameInMap("NextContionuationToken")
    public String nextContionuationToken;

    /**
     * <strong>example:</strong>
     * <p>1032307xxxx72056</p>
     */
    @NameInMap("Owner")
    public String owner;

    /**
     * <strong>example:</strong>
     * <p>oss-cn-shanghai</p>
     */
    @NameInMap("Region")
    public String region;

    @NameInMap("ResourcePool")
    public java.util.List<ResourcePoolSimpleInfo> resourcePool;

    public static ListResourcePoolsResult build(java.util.Map<String, ?> map) throws Exception {
        ListResourcePoolsResult self = new ListResourcePoolsResult();
        return TeaModel.build(map, self);
    }

    public ListResourcePoolsResult setContionuationToken(String contionuationToken) {
        this.contionuationToken = contionuationToken;
        return this;
    }
    public String getContionuationToken() {
        return this.contionuationToken;
    }

    public ListResourcePoolsResult setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListResourcePoolsResult setNextContionuationToken(String nextContionuationToken) {
        this.nextContionuationToken = nextContionuationToken;
        return this;
    }
    public String getNextContionuationToken() {
        return this.nextContionuationToken;
    }

    public ListResourcePoolsResult setOwner(String owner) {
        this.owner = owner;
        return this;
    }
    public String getOwner() {
        return this.owner;
    }

    public ListResourcePoolsResult setRegion(String region) {
        this.region = region;
        return this;
    }
    public String getRegion() {
        return this.region;
    }

    public ListResourcePoolsResult setResourcePool(java.util.List<ResourcePoolSimpleInfo> resourcePool) {
        this.resourcePool = resourcePool;
        return this;
    }
    public java.util.List<ResourcePoolSimpleInfo> getResourcePool() {
        return this.resourcePool;
    }

}
