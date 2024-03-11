// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListObjectVersionsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the results of the ListObjectVersions (GetBucketVersions) request.</p>
     */
    @NameInMap("ListVersionsResult")
    public ListVersionsResult listVersionsResult;

    public static ListObjectVersionsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListObjectVersionsResponseBody self = new ListObjectVersionsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListObjectVersionsResponseBody setListVersionsResult(ListVersionsResult listVersionsResult) {
        this.listVersionsResult = listVersionsResult;
        return this;
    }
    public ListVersionsResult getListVersionsResult() {
        return this.listVersionsResult;
    }

    public static class ListVersionsResult extends TeaModel {
        /**
         * <p>Objects whose names contain the same string that ranges from the prefix to the next occurrence of the delimiter are grouped as a single result element</p>
         */
        @NameInMap("CommonPrefixes")
        public java.util.List<CommonPrefix> commonPrefixes;

        /**
         * <p>The container that stores delete markers</p>
         */
        @NameInMap("DeleteMarker")
        public java.util.List<DeleteMarkerEntry> deleteMarkers;

        /**
         * <p>The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result parameter in CommonPrefixes.</p>
         */
        @NameInMap("Delimiter")
        public String delimiter;

        /**
         * <p>The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.</p>
         */
        @NameInMap("EncodingType")
        public String encodingType;

        /**
         * <p>Indicates whether the returned results are truncated.</p>
         * <br>
         * <p>- true: indicates that not all results are returned for the request.</p>
         * <p>- false: indicates that all results are returned for the request.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>Indicates the object from which the ListObjectVersions (GetBucketVersions) operation starts.</p>
         */
        @NameInMap("KeyMarker")
        public String keyMarker;

        /**
         * <p>The maximum number of objects that can be returned in the response.</p>
         */
        @NameInMap("MaxKeys")
        public Long maxKeys;

        /**
         * <p>The bucket name</p>
         */
        @NameInMap("Name")
        public String name;

        /**
         * <p>If not all results are returned for the request, the NextKeyMarker parameter is included in the response to indicate the key-marker value of the next ListObjectVersions (GetBucketVersions) request.</p>
         */
        @NameInMap("NextKeyMarker")
        public String nextKeyMarker;

        /**
         * <p>If not all results are returned for the request, the NextVersionIdMarker parameter is included in the response to indicate the version-id-marker value of the next ListObjectVersions (GetBucketVersions) request.</p>
         */
        @NameInMap("NextVersionIdMarker")
        public String nextVersionIdMarker;

        /**
         * <p>The prefix contained in the names of the returned objects.</p>
         */
        @NameInMap("Prefix")
        public String prefix;

        /**
         * <p>The container that stores the versions of objects except for delete markers</p>
         */
        @NameInMap("Version")
        public java.util.List<ObjectVersion> versions;

        /**
         * <p>The version from which the ListObjectVersions (GetBucketVersions) operation starts. This parameter is used together with KeyMarker.</p>
         */
        @NameInMap("VersionIdMarker")
        public String versionIdMarker;

        public static ListVersionsResult build(java.util.Map<String, ?> map) throws Exception {
            ListVersionsResult self = new ListVersionsResult();
            return TeaModel.build(map, self);
        }

        public ListVersionsResult setCommonPrefixes(java.util.List<CommonPrefix> commonPrefixes) {
            this.commonPrefixes = commonPrefixes;
            return this;
        }
        public java.util.List<CommonPrefix> getCommonPrefixes() {
            return this.commonPrefixes;
        }

        public ListVersionsResult setDeleteMarkers(java.util.List<DeleteMarkerEntry> deleteMarkers) {
            this.deleteMarkers = deleteMarkers;
            return this;
        }
        public java.util.List<DeleteMarkerEntry> getDeleteMarkers() {
            return this.deleteMarkers;
        }

        public ListVersionsResult setDelimiter(String delimiter) {
            this.delimiter = delimiter;
            return this;
        }
        public String getDelimiter() {
            return this.delimiter;
        }

        public ListVersionsResult setEncodingType(String encodingType) {
            this.encodingType = encodingType;
            return this;
        }
        public String getEncodingType() {
            return this.encodingType;
        }

        public ListVersionsResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListVersionsResult setKeyMarker(String keyMarker) {
            this.keyMarker = keyMarker;
            return this;
        }
        public String getKeyMarker() {
            return this.keyMarker;
        }

        public ListVersionsResult setMaxKeys(Long maxKeys) {
            this.maxKeys = maxKeys;
            return this;
        }
        public Long getMaxKeys() {
            return this.maxKeys;
        }

        public ListVersionsResult setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

        public ListVersionsResult setNextKeyMarker(String nextKeyMarker) {
            this.nextKeyMarker = nextKeyMarker;
            return this;
        }
        public String getNextKeyMarker() {
            return this.nextKeyMarker;
        }

        public ListVersionsResult setNextVersionIdMarker(String nextVersionIdMarker) {
            this.nextVersionIdMarker = nextVersionIdMarker;
            return this;
        }
        public String getNextVersionIdMarker() {
            return this.nextVersionIdMarker;
        }

        public ListVersionsResult setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

        public ListVersionsResult setVersions(java.util.List<ObjectVersion> versions) {
            this.versions = versions;
            return this;
        }
        public java.util.List<ObjectVersion> getVersions() {
            return this.versions;
        }

        public ListVersionsResult setVersionIdMarker(String versionIdMarker) {
            this.versionIdMarker = versionIdMarker;
            return this;
        }
        public String getVersionIdMarker() {
            return this.versionIdMarker;
        }

    }

}
