// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchCopyObject extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>REPLACE</p>
     */
    @NameInMap("MetadataDirective")
    public String metadataDirective;

    @NameInMap("NewObjectMetadata")
    public String newObjectMetadata;

    /**
     * <strong>example:</strong>
     * <p>TagA=A&amp;TagB=B</p>
     */
    @NameInMap("NewObjectTagging")
    public String newObjectTagging;

    @NameInMap("ObjectAcl")
    public String objectAcl;

    /**
     * <strong>example:</strong>
     * <p>SM4</p>
     */
    @NameInMap("ServerSideDataEncryption")
    public String serverSideDataEncryption;

    /**
     * <strong>example:</strong>
     * <p>AES256</p>
     */
    @NameInMap("ServerSideEncryption")
    public String serverSideEncryption;

    @NameInMap("StorageClass")
    public String storageClass;

    /**
     * <strong>example:</strong>
     * <p>REPLACE</p>
     */
    @NameInMap("TaggingDirective")
    public String taggingDirective;

    /**
     * <strong>example:</strong>
     * <p>/xxx</p>
     */
    @NameInMap("TargetKeyPrefix")
    public String targetKeyPrefix;

    /**
     * <strong>example:</strong>
     * <p>bucketname</p>
     */
    @NameInMap("TargetResource")
    public String targetResource;

    public static BatchCopyObject build(java.util.Map<String, ?> map) throws Exception {
        BatchCopyObject self = new BatchCopyObject();
        return TeaModel.build(map, self);
    }

    public BatchCopyObject setMetadataDirective(String metadataDirective) {
        this.metadataDirective = metadataDirective;
        return this;
    }
    public String getMetadataDirective() {
        return this.metadataDirective;
    }

    public BatchCopyObject setNewObjectMetadata(String newObjectMetadata) {
        this.newObjectMetadata = newObjectMetadata;
        return this;
    }
    public String getNewObjectMetadata() {
        return this.newObjectMetadata;
    }

    public BatchCopyObject setNewObjectTagging(String newObjectTagging) {
        this.newObjectTagging = newObjectTagging;
        return this;
    }
    public String getNewObjectTagging() {
        return this.newObjectTagging;
    }

    public BatchCopyObject setObjectAcl(String objectAcl) {
        this.objectAcl = objectAcl;
        return this;
    }
    public String getObjectAcl() {
        return this.objectAcl;
    }

    public BatchCopyObject setServerSideDataEncryption(String serverSideDataEncryption) {
        this.serverSideDataEncryption = serverSideDataEncryption;
        return this;
    }
    public String getServerSideDataEncryption() {
        return this.serverSideDataEncryption;
    }

    public BatchCopyObject setServerSideEncryption(String serverSideEncryption) {
        this.serverSideEncryption = serverSideEncryption;
        return this;
    }
    public String getServerSideEncryption() {
        return this.serverSideEncryption;
    }

    public BatchCopyObject setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public BatchCopyObject setTaggingDirective(String taggingDirective) {
        this.taggingDirective = taggingDirective;
        return this;
    }
    public String getTaggingDirective() {
        return this.taggingDirective;
    }

    public BatchCopyObject setTargetKeyPrefix(String targetKeyPrefix) {
        this.targetKeyPrefix = targetKeyPrefix;
        return this;
    }
    public String getTargetKeyPrefix() {
        return this.targetKeyPrefix;
    }

    public BatchCopyObject setTargetResource(String targetResource) {
        this.targetResource = targetResource;
        return this;
    }
    public String getTargetResource() {
        return this.targetResource;
    }

}
