// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ObjectSummary extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>5B3C1A2E053D763E1B002CC607C5A0FE1****</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>fun/test.jpg</p>
     */
    @NameInMap("Key")
    public String key;

    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2012-02-24T08:42:32.000Z</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    @NameInMap("Owner")
    public Owner owner;

    /**
     * <strong>example:</strong>
     * <p>ongoing-request=&quot;true”</p>
     */
    @NameInMap("RestoreInfo")
    public String restoreInfo;

    /**
     * <strong>example:</strong>
     * <p>344606</p>
     */
    @NameInMap("Size")
    public Long size;

    @NameInMap("StorageClass")
    public String storageClass;

    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2012-02-24T08:42:32.000Z</p>
     */
    @NameInMap("TransitionTime")
    public String transitionTime;

    /**
     * <strong>example:</strong>
     * <p>Normal</p>
     */
    @NameInMap("Type")
    public String type;

    public static ObjectSummary build(java.util.Map<String, ?> map) throws Exception {
        ObjectSummary self = new ObjectSummary();
        return TeaModel.build(map, self);
    }

    public ObjectSummary setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public ObjectSummary setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public ObjectSummary setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

    public ObjectSummary setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public ObjectSummary setRestoreInfo(String restoreInfo) {
        this.restoreInfo = restoreInfo;
        return this;
    }
    public String getRestoreInfo() {
        return this.restoreInfo;
    }

    public ObjectSummary setSize(Long size) {
        this.size = size;
        return this;
    }
    public Long getSize() {
        return this.size;
    }

    public ObjectSummary setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public ObjectSummary setTransitionTime(String transitionTime) {
        this.transitionTime = transitionTime;
        return this;
    }
    public String getTransitionTime() {
        return this.transitionTime;
    }

    public ObjectSummary setType(String type) {
        this.type = type;
        return this;
    }
    public String getType() {
        return this.type;
    }

}
