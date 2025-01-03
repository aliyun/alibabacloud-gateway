// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BucketStat extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("ArchiveObjectCount")
    public Long archiveObjectCount;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("ArchiveRealStorage")
    public Long archiveRealStorage;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("ArchiveStorage")
    public Long archiveStorage;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("ColdArchiveObjectCount")
    public Long coldArchiveObjectCount;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("ColdArchiveRealStorage")
    public Long coldArchiveRealStorage;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("ColdArchiveStorage")
    public Long coldArchiveStorage;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("DeepColdArchiveObjectCount")
    public Long deepColdArchiveObjectCount;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("DeepColdArchiveRealStorage")
    public Long deepColdArchiveRealStorage;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("DeepColdArchiveStorage")
    public Long deepColdArchiveStorage;

    /**
     * <strong>example:</strong>
     * <p>12</p>
     */
    @NameInMap("DeleteMarkerCount")
    public Long deleteMarkerCount;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("InfrequentAccessObjectCount")
    public Long infrequentAccessObjectCount;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("InfrequentAccessRealStorage")
    public Long infrequentAccessRealStorage;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("InfrequentAccessStorage")
    public Long infrequentAccessStorage;

    /**
     * <strong>example:</strong>
     * <p>1709724731</p>
     */
    @NameInMap("LastModifiedTime")
    public Long lastModifiedTime;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("LiveChannelCount")
    public Long liveChannelCount;

    /**
     * <strong>example:</strong>
     * <p>128</p>
     */
    @NameInMap("MultipartPartCount")
    public Long multipartPartCount;

    /**
     * <strong>example:</strong>
     * <p>27</p>
     */
    @NameInMap("MultipartUploadCount")
    public Long multipartUploadCount;

    /**
     * <strong>example:</strong>
     * <p>32</p>
     */
    @NameInMap("ObjectCount")
    public Long objectCount;

    /**
     * <strong>example:</strong>
     * <p>18</p>
     */
    @NameInMap("StandardObjectCount")
    public Long standardObjectCount;

    /**
     * <strong>example:</strong>
     * <p>1990</p>
     */
    @NameInMap("StandardStorage")
    public Long standardStorage;

    /**
     * <strong>example:</strong>
     * <p>1994</p>
     */
    @NameInMap("Storage")
    public Long storage;

    public static BucketStat build(java.util.Map<String, ?> map) throws Exception {
        BucketStat self = new BucketStat();
        return TeaModel.build(map, self);
    }

    public BucketStat setArchiveObjectCount(Long archiveObjectCount) {
        this.archiveObjectCount = archiveObjectCount;
        return this;
    }
    public Long getArchiveObjectCount() {
        return this.archiveObjectCount;
    }

    public BucketStat setArchiveRealStorage(Long archiveRealStorage) {
        this.archiveRealStorage = archiveRealStorage;
        return this;
    }
    public Long getArchiveRealStorage() {
        return this.archiveRealStorage;
    }

    public BucketStat setArchiveStorage(Long archiveStorage) {
        this.archiveStorage = archiveStorage;
        return this;
    }
    public Long getArchiveStorage() {
        return this.archiveStorage;
    }

    public BucketStat setColdArchiveObjectCount(Long coldArchiveObjectCount) {
        this.coldArchiveObjectCount = coldArchiveObjectCount;
        return this;
    }
    public Long getColdArchiveObjectCount() {
        return this.coldArchiveObjectCount;
    }

    public BucketStat setColdArchiveRealStorage(Long coldArchiveRealStorage) {
        this.coldArchiveRealStorage = coldArchiveRealStorage;
        return this;
    }
    public Long getColdArchiveRealStorage() {
        return this.coldArchiveRealStorage;
    }

    public BucketStat setColdArchiveStorage(Long coldArchiveStorage) {
        this.coldArchiveStorage = coldArchiveStorage;
        return this;
    }
    public Long getColdArchiveStorage() {
        return this.coldArchiveStorage;
    }

    public BucketStat setDeepColdArchiveObjectCount(Long deepColdArchiveObjectCount) {
        this.deepColdArchiveObjectCount = deepColdArchiveObjectCount;
        return this;
    }
    public Long getDeepColdArchiveObjectCount() {
        return this.deepColdArchiveObjectCount;
    }

    public BucketStat setDeepColdArchiveRealStorage(Long deepColdArchiveRealStorage) {
        this.deepColdArchiveRealStorage = deepColdArchiveRealStorage;
        return this;
    }
    public Long getDeepColdArchiveRealStorage() {
        return this.deepColdArchiveRealStorage;
    }

    public BucketStat setDeepColdArchiveStorage(Long deepColdArchiveStorage) {
        this.deepColdArchiveStorage = deepColdArchiveStorage;
        return this;
    }
    public Long getDeepColdArchiveStorage() {
        return this.deepColdArchiveStorage;
    }

    public BucketStat setDeleteMarkerCount(Long deleteMarkerCount) {
        this.deleteMarkerCount = deleteMarkerCount;
        return this;
    }
    public Long getDeleteMarkerCount() {
        return this.deleteMarkerCount;
    }

    public BucketStat setInfrequentAccessObjectCount(Long infrequentAccessObjectCount) {
        this.infrequentAccessObjectCount = infrequentAccessObjectCount;
        return this;
    }
    public Long getInfrequentAccessObjectCount() {
        return this.infrequentAccessObjectCount;
    }

    public BucketStat setInfrequentAccessRealStorage(Long infrequentAccessRealStorage) {
        this.infrequentAccessRealStorage = infrequentAccessRealStorage;
        return this;
    }
    public Long getInfrequentAccessRealStorage() {
        return this.infrequentAccessRealStorage;
    }

    public BucketStat setInfrequentAccessStorage(Long infrequentAccessStorage) {
        this.infrequentAccessStorage = infrequentAccessStorage;
        return this;
    }
    public Long getInfrequentAccessStorage() {
        return this.infrequentAccessStorage;
    }

    public BucketStat setLastModifiedTime(Long lastModifiedTime) {
        this.lastModifiedTime = lastModifiedTime;
        return this;
    }
    public Long getLastModifiedTime() {
        return this.lastModifiedTime;
    }

    public BucketStat setLiveChannelCount(Long liveChannelCount) {
        this.liveChannelCount = liveChannelCount;
        return this;
    }
    public Long getLiveChannelCount() {
        return this.liveChannelCount;
    }

    public BucketStat setMultipartPartCount(Long multipartPartCount) {
        this.multipartPartCount = multipartPartCount;
        return this;
    }
    public Long getMultipartPartCount() {
        return this.multipartPartCount;
    }

    public BucketStat setMultipartUploadCount(Long multipartUploadCount) {
        this.multipartUploadCount = multipartUploadCount;
        return this;
    }
    public Long getMultipartUploadCount() {
        return this.multipartUploadCount;
    }

    public BucketStat setObjectCount(Long objectCount) {
        this.objectCount = objectCount;
        return this;
    }
    public Long getObjectCount() {
        return this.objectCount;
    }

    public BucketStat setStandardObjectCount(Long standardObjectCount) {
        this.standardObjectCount = standardObjectCount;
        return this;
    }
    public Long getStandardObjectCount() {
        return this.standardObjectCount;
    }

    public BucketStat setStandardStorage(Long standardStorage) {
        this.standardStorage = standardStorage;
        return this;
    }
    public Long getStandardStorage() {
        return this.standardStorage;
    }

    public BucketStat setStorage(Long storage) {
        this.storage = storage;
        return this;
    }
    public Long getStorage() {
        return this.storage;
    }

}
