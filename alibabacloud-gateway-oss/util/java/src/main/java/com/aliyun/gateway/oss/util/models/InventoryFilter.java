// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InventoryFilter extends TeaModel {
    @NameInMap("LastModifyBeginTimeStamp")
    public Long lastModifyBeginTimeStamp;

    @NameInMap("LastModifyEndTimeStamp")
    public Long lastModifyEndTimeStamp;

    @NameInMap("LowerSizeBound")
    public Long lowerSizeBound;

    @NameInMap("Prefix")
    public String prefix;

    @NameInMap("StorageClass")
    public String storageClass;

    @NameInMap("Tags")
    public String tags;

    @NameInMap("TagsCondition")
    public String tagsCondition;

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
