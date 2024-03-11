// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListMultipartUploadsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the response to the ListMultipartUpload request.</p>
     */
    @NameInMap("ListMultipartUploadsResult")
    public ListMultipartUploadsResult listMultipartUploadsResult;

    public static ListMultipartUploadsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListMultipartUploadsResponseBody self = new ListMultipartUploadsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListMultipartUploadsResponseBody setListMultipartUploadsResult(ListMultipartUploadsResult listMultipartUploadsResult) {
        this.listMultipartUploadsResult = listMultipartUploadsResult;
        return this;
    }
    public ListMultipartUploadsResult getListMultipartUploadsResult() {
        return this.listMultipartUploadsResult;
    }

    public static class ListMultipartUploadsResult extends TeaModel {
        /**
         * <p>The name of the bucket.</p>
         */
        @NameInMap("Bucket")
        public String bucket;

        /**
         * <p>If the delimiter parameter is specified in the request, the response contains the CommonPrefixes parameter. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.</p>
         */
        @NameInMap("CommonPrefixes")
        public java.util.List<CommonPrefix> commonPrefixes;

        /**
         * <p>The character used to group objects by name. If you specify the Delimiter parameter in the request, the response contains the CommonPrefixes element. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in</p>
         */
        @NameInMap("Delimiter")
        public String delimiter;

        /**
         * <p>The method used to encode the object name in the response. If encoding-type is specified in the request, values of those elements including Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key are encoded in the returned result.</p>
         */
        @NameInMap("EncodingType")
        public String encodingType;

        /**
         * <p>Indicates whether the list of multipart upload tasks returned in the response is truncated. Default value: false. Valid values:</p>
         * <br>
         * <p>- true: Only part of the results are returned this time.</p>
         * <br>
         * <p>- false: All results are returned.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The name of the object that corresponds to the multipart upload task after which the list begins.</p>
         */
        @NameInMap("KeyMarker")
        public String keyMarker;

        /**
         * <p>The maximum number of multipart upload tasks returned by OSS.</p>
         */
        @NameInMap("MaxUploads")
        public Long maxUploads;

        /**
         * <p>The object name marker in the response for the next request to return the remaining results.</p>
         */
        @NameInMap("NextKeyMarker")
        public String nextKeyMarker;

        /**
         * <p>The NextUploadMarker value that is used for the UploadMarker value in the next request if the response does not contain all required results.</p>
         */
        @NameInMap("NextUploadIdMarker")
        public String nextUploadIdMarker;

        /**
         * <p>The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.</p>
         */
        @NameInMap("Prefix")
        public String prefix;

        /**
         * <p>The ID list of the multipart upload tasks.</p>
         */
        @NameInMap("Upload")
        public java.util.List<Upload> uploads;

        /**
         * <p>The upload ID of the multipart upload task after which the list begins.</p>
         */
        @NameInMap("UploadIdMarker")
        public String uploadIdMarker;

        public static ListMultipartUploadsResult build(java.util.Map<String, ?> map) throws Exception {
            ListMultipartUploadsResult self = new ListMultipartUploadsResult();
            return TeaModel.build(map, self);
        }

        public ListMultipartUploadsResult setBucket(String bucket) {
            this.bucket = bucket;
            return this;
        }
        public String getBucket() {
            return this.bucket;
        }

        public ListMultipartUploadsResult setCommonPrefixes(java.util.List<CommonPrefix> commonPrefixes) {
            this.commonPrefixes = commonPrefixes;
            return this;
        }
        public java.util.List<CommonPrefix> getCommonPrefixes() {
            return this.commonPrefixes;
        }

        public ListMultipartUploadsResult setDelimiter(String delimiter) {
            this.delimiter = delimiter;
            return this;
        }
        public String getDelimiter() {
            return this.delimiter;
        }

        public ListMultipartUploadsResult setEncodingType(String encodingType) {
            this.encodingType = encodingType;
            return this;
        }
        public String getEncodingType() {
            return this.encodingType;
        }

        public ListMultipartUploadsResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListMultipartUploadsResult setKeyMarker(String keyMarker) {
            this.keyMarker = keyMarker;
            return this;
        }
        public String getKeyMarker() {
            return this.keyMarker;
        }

        public ListMultipartUploadsResult setMaxUploads(Long maxUploads) {
            this.maxUploads = maxUploads;
            return this;
        }
        public Long getMaxUploads() {
            return this.maxUploads;
        }

        public ListMultipartUploadsResult setNextKeyMarker(String nextKeyMarker) {
            this.nextKeyMarker = nextKeyMarker;
            return this;
        }
        public String getNextKeyMarker() {
            return this.nextKeyMarker;
        }

        public ListMultipartUploadsResult setNextUploadIdMarker(String nextUploadIdMarker) {
            this.nextUploadIdMarker = nextUploadIdMarker;
            return this;
        }
        public String getNextUploadIdMarker() {
            return this.nextUploadIdMarker;
        }

        public ListMultipartUploadsResult setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

        public ListMultipartUploadsResult setUploads(java.util.List<Upload> uploads) {
            this.uploads = uploads;
            return this;
        }
        public java.util.List<Upload> getUploads() {
            return this.uploads;
        }

        public ListMultipartUploadsResult setUploadIdMarker(String uploadIdMarker) {
            this.uploadIdMarker = uploadIdMarker;
            return this;
        }
        public String getUploadIdMarker() {
            return this.uploadIdMarker;
        }

    }

}
