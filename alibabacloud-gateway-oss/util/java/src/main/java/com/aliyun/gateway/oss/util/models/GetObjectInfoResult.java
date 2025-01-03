// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectInfoResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>text/plain</p>
     */
    @NameInMap("Content-Type")
    public String contentType;

    /**
     * <strong>example:</strong>
     * <p>&quot;EB327B57BB20D17C293A966115FE27BD&quot;</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("EncryptFlag")
    public Long encryptFlag;

    /**
     * <strong>example:</strong>
     * <p>7866203970082294162</p>
     */
    @NameInMap("HashCrc64ecma")
    public String hashCrc64ecma;

    /**
     * <strong>example:</strong>
     * <p>test.txt</p>
     */
    @NameInMap("Key")
    public String key;

    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2024-11-08T10:02:11.000Z</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    /**
     * <strong>example:</strong>
     * <p>78</p>
     */
    @NameInMap("Size")
    public String size;

    /**
     * <strong>example:</strong>
     * <p>Standard</p>
     */
    @NameInMap("StorageClass")
    public String storageClass;

    /**
     * <strong>example:</strong>
     * <p>Multipart</p>
     */
    @NameInMap("Type")
    public String type;

    /**
     * <strong>example:</strong>
     * <p>C5DD87C5E7CD482F8F2C3D63904DA228</p>
     */
    @NameInMap("UploadId")
    public String uploadId;

    public static GetObjectInfoResult build(java.util.Map<String, ?> map) throws Exception {
        GetObjectInfoResult self = new GetObjectInfoResult();
        return TeaModel.build(map, self);
    }

    public GetObjectInfoResult setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public GetObjectInfoResult setContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public String getContentType() {
        return this.contentType;
    }

    public GetObjectInfoResult setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public GetObjectInfoResult setEncryptFlag(Long encryptFlag) {
        this.encryptFlag = encryptFlag;
        return this;
    }
    public Long getEncryptFlag() {
        return this.encryptFlag;
    }

    public GetObjectInfoResult setHashCrc64ecma(String hashCrc64ecma) {
        this.hashCrc64ecma = hashCrc64ecma;
        return this;
    }
    public String getHashCrc64ecma() {
        return this.hashCrc64ecma;
    }

    public GetObjectInfoResult setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public GetObjectInfoResult setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

    public GetObjectInfoResult setSize(String size) {
        this.size = size;
        return this;
    }
    public String getSize() {
        return this.size;
    }

    public GetObjectInfoResult setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public GetObjectInfoResult setType(String type) {
        this.type = type;
        return this;
    }
    public String getType() {
        return this.type;
    }

    public GetObjectInfoResult setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
