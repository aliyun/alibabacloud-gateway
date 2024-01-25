// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListPartsResponseBody extends TeaModel {
    /**
     * <p>The name of the bucket.</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <p>Indicates whether the list of parts returned in the response has been truncated. "true" indicates that the response does not contain all required results. "false" indicates that the response contains all required results.</p>
     * <p><br>Valid values: true and false</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>The full path of the object.</p>
     */
    @NameInMap("Key")
    public String key;

    /**
     * <p>The maximum number of parts in this response.</p>
     */
    @NameInMap("MaxParts")
    public Long maxParts;

    /**
     * <p>If the response does not contain all required results, all parts whose part numbers are greater than the value of this parameter are listed.</p>
     */
    @NameInMap("NextPartNumberMarker")
    public Long nextPartNumberMarker;

    /**
     * <p>The container that stores part information.</p>
     */
    @NameInMap("Part")
    public java.util.List<Part> part;

    /**
     * <p>The number of the part after which the list operation begins. All parts whose numbers are greater than the value of this parameter are listed.</p>
     */
    @NameInMap("PartNumberMarker")
    public Long partNumberMarker;

    /**
     * <p>The ID of the upload task.</p>
     */
    @NameInMap("UploadId")
    public String uploadId;

    public static ListPartsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListPartsResponseBody self = new ListPartsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListPartsResponseBody setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public ListPartsResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListPartsResponseBody setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public ListPartsResponseBody setMaxParts(Long maxParts) {
        this.maxParts = maxParts;
        return this;
    }
    public Long getMaxParts() {
        return this.maxParts;
    }

    public ListPartsResponseBody setNextPartNumberMarker(Long nextPartNumberMarker) {
        this.nextPartNumberMarker = nextPartNumberMarker;
        return this;
    }
    public Long getNextPartNumberMarker() {
        return this.nextPartNumberMarker;
    }

    public ListPartsResponseBody setPart(java.util.List<Part> part) {
        this.part = part;
        return this;
    }
    public java.util.List<Part> getPart() {
        return this.part;
    }

    public ListPartsResponseBody setPartNumberMarker(Long partNumberMarker) {
        this.partNumberMarker = partNumberMarker;
        return this;
    }
    public Long getPartNumberMarker() {
        return this.partNumberMarker;
    }

    public ListPartsResponseBody setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
