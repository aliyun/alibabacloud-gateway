// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ObjectVersion extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>&quot;250F8A0AE989679A22926A875F0A2****&quot;</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("IsLatest")
    public Boolean isLatest;

    /**
     * <strong>example:</strong>
     * <p>dic/test.txt</p>
     */
    @NameInMap("Key")
    public String key;

    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2023-04-09T07:27:28.000Z</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    @NameInMap("Owner")
    public Owner owner;

    /**
     * <strong>example:</strong>
     * <p>ongoing-request=&quot;true&quot;</p>
     */
    @NameInMap("RestoreInfo")
    public String restoreInfo;

    /**
     * <strong>example:</strong>
     * <p>93731</p>
     */
    @NameInMap("Size")
    public Long size;

    @NameInMap("StorageClass")
    public String storageClass;

    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2023-04-09T07:27:28.000Z</p>
     */
    @NameInMap("TransitionTime")
    public String transitionTime;

    /**
     * <strong>example:</strong>
     * <p>CAEQMxiBgMDNoP2D0BYiIDE3MWUxNzgxZDQxNTRiODI5OGYwZGMwNGY3MzZjN****</p>
     */
    @NameInMap("VersionId")
    public String versionId;

    public static ObjectVersion build(java.util.Map<String, ?> map) throws Exception {
        ObjectVersion self = new ObjectVersion();
        return TeaModel.build(map, self);
    }

    public ObjectVersion setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public ObjectVersion setIsLatest(Boolean isLatest) {
        this.isLatest = isLatest;
        return this;
    }
    public Boolean getIsLatest() {
        return this.isLatest;
    }

    public ObjectVersion setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public ObjectVersion setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

    public ObjectVersion setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public ObjectVersion setRestoreInfo(String restoreInfo) {
        this.restoreInfo = restoreInfo;
        return this;
    }
    public String getRestoreInfo() {
        return this.restoreInfo;
    }

    public ObjectVersion setSize(Long size) {
        this.size = size;
        return this;
    }
    public Long getSize() {
        return this.size;
    }

    public ObjectVersion setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public ObjectVersion setTransitionTime(String transitionTime) {
        this.transitionTime = transitionTime;
        return this;
    }
    public String getTransitionTime() {
        return this.transitionTime;
    }

    public ObjectVersion setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
