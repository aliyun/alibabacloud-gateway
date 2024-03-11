// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListPartsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the response of the ListParts request.</p>
     */
    @NameInMap("ListPartResult")
    public ListPartResult listPartResult;

    public static ListPartsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListPartsResponseBody self = new ListPartsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListPartsResponseBody setListPartResult(ListPartResult listPartResult) {
        this.listPartResult = listPartResult;
        return this;
    }
    public ListPartResult getListPartResult() {
        return this.listPartResult;
    }

    public static class ListPartResult extends TeaModel {
        /**
         * <p>The name of the bucket.</p>
         */
        @NameInMap("Bucket")
        public String bucket;

        /**
         * <p>Indicates whether the list of parts returned in the response has been truncated. A value of true indicates that the response does not contain all required results. A value of false indicates that the response contains all required results.</p>
         * <br>
         * <p>Valid values: true and false.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The name of the object.</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <p>The maximum number of parts in the response.</p>
         */
        @NameInMap("MaxParts")
        public Long maxParts;

        /**
         * <p>The NextPartNumberMarker value that is used for the PartNumberMarker value in a subsequent request when the response does not contain all required results.</p>
         */
        @NameInMap("NextPartNumberMarker")
        public Long nextPartNumberMarker;

        /**
         * <p>The list of all parts.</p>
         */
        @NameInMap("Part")
        public java.util.List<Part> part;

        /**
         * <p>The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.</p>
         */
        @NameInMap("PartNumberMarker")
        public Long partNumberMarker;

        /**
         * <p>The ID of the upload task.</p>
         */
        @NameInMap("UploadId")
        public String uploadId;

        public static ListPartResult build(java.util.Map<String, ?> map) throws Exception {
            ListPartResult self = new ListPartResult();
            return TeaModel.build(map, self);
        }

        public ListPartResult setBucket(String bucket) {
            this.bucket = bucket;
            return this;
        }
        public String getBucket() {
            return this.bucket;
        }

        public ListPartResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListPartResult setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public ListPartResult setMaxParts(Long maxParts) {
            this.maxParts = maxParts;
            return this;
        }
        public Long getMaxParts() {
            return this.maxParts;
        }

        public ListPartResult setNextPartNumberMarker(Long nextPartNumberMarker) {
            this.nextPartNumberMarker = nextPartNumberMarker;
            return this;
        }
        public Long getNextPartNumberMarker() {
            return this.nextPartNumberMarker;
        }

        public ListPartResult setPart(java.util.List<Part> part) {
            this.part = part;
            return this;
        }
        public java.util.List<Part> getPart() {
            return this.part;
        }

        public ListPartResult setPartNumberMarker(Long partNumberMarker) {
            this.partNumberMarker = partNumberMarker;
            return this;
        }
        public Long getPartNumberMarker() {
            return this.partNumberMarker;
        }

        public ListPartResult setUploadId(String uploadId) {
            this.uploadId = uploadId;
            return this;
        }
        public String getUploadId() {
            return this.uploadId;
        }

    }

}
