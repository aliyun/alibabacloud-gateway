// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListObjectsV2ResponseBody extends TeaModel {
    /**
     * <p>The container that stores the metadata of returned objects.</p>
     */
    @NameInMap("ListBucketResult")
    public ListBucketResult listBucketResult;

    public static ListObjectsV2ResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListObjectsV2ResponseBody self = new ListObjectsV2ResponseBody();
        return TeaModel.build(map, self);
    }

    public ListObjectsV2ResponseBody setListBucketResult(ListBucketResult listBucketResult) {
        this.listBucketResult = listBucketResult;
        return this;
    }
    public ListBucketResult getListBucketResult() {
        return this.listBucketResult;
    }

    public static class ListBucketResult extends TeaModel {
        /**
         * <p>Objects whose names contain the same string that ranges from the prefix to the next occurrence of the delimiter are grouped as a single result element</p>
         */
        @NameInMap("CommonPrefixes")
        public java.util.List<CommonPrefix> commonPrefixes;

        /**
         * <p>The container that stores the metadata of the returned objects.</p>
         */
        @NameInMap("Contents")
        public java.util.List<ObjectSummary> contents;

        /**
         * <p>If continuation-token is specified in the request, the response contains ContinuationToken.</p>
         */
        @NameInMap("ContinuationToken")
        public String continuationToken;

        /**
         * <p>The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
         */
        @NameInMap("Delimiter")
        public String delimiter;

        /**
         * <p>The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, StartAfter, Prefix, NextContinuationToken, and Key are encoded in the response.</p>
         */
        @NameInMap("EncodingType")
        public String encodingType;

        /**
         * <p>Indicates whether the returned results are truncated. Valid values:</p>
         * <br>
         * <p>- true</p>
         * <p>- false</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The number of objects returned for this request. If delimiter is specified in the request, the value of KeyCount is the sum of the values of Key and CommonPrefixes.</p>
         */
        @NameInMap("KeyCount")
        public Integer keyCount;

        /**
         * <p>The maximum number of returned objects in the response.</p>
         */
        @NameInMap("MaxKeys")
        public Integer maxKeys;

        /**
         * <p>The name of the bucket.</p>
         */
        @NameInMap("Name")
        public String name;

        /**
         * <p>The token from which the next list operation starts. Use the value of NextContinuationToken as the value of continuation-token in the next request.</p>
         */
        @NameInMap("NextContinuationToken")
        public String nextContinuationToken;

        /**
         * <p>The prefix in the names of the returned objects.</p>
         */
        @NameInMap("Prefix")
        public String prefix;

        /**
         * <p>If start-after is specified in the request, the response contains StartAfter.</p>
         */
        @NameInMap("StartAfter")
        public String startAfter;

        public static ListBucketResult build(java.util.Map<String, ?> map) throws Exception {
            ListBucketResult self = new ListBucketResult();
            return TeaModel.build(map, self);
        }

        public ListBucketResult setCommonPrefixes(java.util.List<CommonPrefix> commonPrefixes) {
            this.commonPrefixes = commonPrefixes;
            return this;
        }
        public java.util.List<CommonPrefix> getCommonPrefixes() {
            return this.commonPrefixes;
        }

        public ListBucketResult setContents(java.util.List<ObjectSummary> contents) {
            this.contents = contents;
            return this;
        }
        public java.util.List<ObjectSummary> getContents() {
            return this.contents;
        }

        public ListBucketResult setContinuationToken(String continuationToken) {
            this.continuationToken = continuationToken;
            return this;
        }
        public String getContinuationToken() {
            return this.continuationToken;
        }

        public ListBucketResult setDelimiter(String delimiter) {
            this.delimiter = delimiter;
            return this;
        }
        public String getDelimiter() {
            return this.delimiter;
        }

        public ListBucketResult setEncodingType(String encodingType) {
            this.encodingType = encodingType;
            return this;
        }
        public String getEncodingType() {
            return this.encodingType;
        }

        public ListBucketResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListBucketResult setKeyCount(Integer keyCount) {
            this.keyCount = keyCount;
            return this;
        }
        public Integer getKeyCount() {
            return this.keyCount;
        }

        public ListBucketResult setMaxKeys(Integer maxKeys) {
            this.maxKeys = maxKeys;
            return this;
        }
        public Integer getMaxKeys() {
            return this.maxKeys;
        }

        public ListBucketResult setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

        public ListBucketResult setNextContinuationToken(String nextContinuationToken) {
            this.nextContinuationToken = nextContinuationToken;
            return this;
        }
        public String getNextContinuationToken() {
            return this.nextContinuationToken;
        }

        public ListBucketResult setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

        public ListBucketResult setStartAfter(String startAfter) {
            this.startAfter = startAfter;
            return this;
        }
        public String getStartAfter() {
            return this.startAfter;
        }

    }

}
