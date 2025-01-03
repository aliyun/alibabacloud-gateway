// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InventoryFilter extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>1637883649</p>
     */
    @NameInMap("LastModifyBeginTimeStamp")
    public Long lastModifyBeginTimeStamp;

    /**
     * <strong>example:</strong>
     * <p>1638347592</p>
     */
    @NameInMap("LastModifyEndTimeStamp")
    public Long lastModifyEndTimeStamp;

    /**
     * <strong>example:</strong>
     * <p>1024</p>
     */
    @NameInMap("LowerSizeBound")
    public Long lowerSizeBound;

    /**
     * <strong>example:</strong>
     * <p>Pics/</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    /**
     * <strong>example:</strong>
     * <p>All</p>
     */
    @NameInMap("StorageClass")
    public String storageClass;

    /**
     * <strong>example:</strong>
     * <p>tag1#val1;tag2#val2</p>
     */
    @NameInMap("Tags")
    public String tags;

    /**
     * <strong>example:</strong>
     * <p>OR_FILTER</p>
     */
    @NameInMap("TagsCondition")
    public String tagsCondition;

    /**
     * <strong>example:</strong>
     * <p>1048576</p>
     */
    @NameInMap("UpperSizeBound")
    public Long upperSizeBound;

    public static InventoryFilter build(java.util.Map<String, ?> map) throws Exception {
        InventoryFilter self = new InventoryFilter();
        return TeaModel.build(map, self);
    }

    public InventoryFilter setLastModifyBeginTimeStamp(Long lastModifyBeginTimeStamp) {
        this.lastModifyBeginTimeStamp = lastModifyBeginTimeStamp;
        return this;
    }
    public Long getLastModifyBeginTimeStamp() {
        return this.lastModifyBeginTimeStamp;
    }

    public InventoryFilter setLastModifyEndTimeStamp(Long lastModifyEndTimeStamp) {
        this.lastModifyEndTimeStamp = lastModifyEndTimeStamp;
        return this;
    }
    public Long getLastModifyEndTimeStamp() {
        return this.lastModifyEndTimeStamp;
    }

    public InventoryFilter setLowerSizeBound(Long lowerSizeBound) {
        this.lowerSizeBound = lowerSizeBound;
        return this;
    }
    public Long getLowerSizeBound() {
        return this.lowerSizeBound;
    }

    public InventoryFilter setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public InventoryFilter setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public InventoryFilter setTags(String tags) {
        this.tags = tags;
        return this;
    }
    public String getTags() {
        return this.tags;
    }

    public InventoryFilter setTagsCondition(String tagsCondition) {
        this.tagsCondition = tagsCondition;
        return this;
    }
    public String getTagsCondition() {
        return this.tagsCondition;
    }

    public InventoryFilter setUpperSizeBound(Long upperSizeBound) {
        this.upperSizeBound = upperSizeBound;
        return this;
    }
    public Long getUpperSizeBound() {
        return this.upperSizeBound;
    }

}
