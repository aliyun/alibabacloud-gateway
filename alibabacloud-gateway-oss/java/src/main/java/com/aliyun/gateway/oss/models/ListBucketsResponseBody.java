// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListBucketsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the result of ListBuckets(GetService) request.</p>
     */
    @NameInMap("ListAllMyBucketsResult")
    public ListAllMyBucketsResult listAllMyBucketsResult;

    public static ListBucketsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketsResponseBody self = new ListBucketsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketsResponseBody setListAllMyBucketsResult(ListAllMyBucketsResult listAllMyBucketsResult) {
        this.listAllMyBucketsResult = listAllMyBucketsResult;
        return this;
    }
    public ListAllMyBucketsResult getListAllMyBucketsResult() {
        return this.listAllMyBucketsResult;
    }

    public static class Buckets extends TeaModel {
        /**
         * <p>The container that stores the information list of multiple buckets.</p>
         */
        @NameInMap("Bucket")
        public java.util.List<Bucket> bucket;

        public static Buckets build(java.util.Map<String, ?> map) throws Exception {
            Buckets self = new Buckets();
            return TeaModel.build(map, self);
        }

        public Buckets setBucket(java.util.List<Bucket> bucket) {
            this.bucket = bucket;
            return this;
        }
        public java.util.List<Bucket> getBucket() {
            return this.bucket;
        }

    }

    public static class ListAllMyBucketsResult extends TeaModel {
        /**
         * <p>The container that stores the information about multiple buckets.</p>
         */
        @NameInMap("Buckets")
        public Buckets buckets;

        /**
         * <p>Indicates whether all results are returned. Valid values:</p>
         * <p>- true: All results are not returned in the response. </p>
         * <p>- false: All results are returned in the response.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The name of the bucket from which the buckets are returned.</p>
         */
        @NameInMap("Marker")
        public String marker;

        /**
         * <p>The maximum number of buckets that can be returned.</p>
         */
        @NameInMap("MaxKeys")
        public Long maxKeys;

        /**
         * <p>The marker for the next ListBuckets (GetService) request. You can use the value of this parameter as the value of marker in the next ListBuckets (GetService) request to retrieve the unreturned results.</p>
         */
        @NameInMap("NextMarker")
        public String nextMarker;

        /**
         * <p>The container that stores the information about the bucket owner.</p>
         */
        @NameInMap("Owner")
        public Owner owner;

        /**
         * <p>The prefix contained in the names of returned buckets.</p>
         */
        @NameInMap("Prefix")
        public String prefix;

        public static ListAllMyBucketsResult build(java.util.Map<String, ?> map) throws Exception {
            ListAllMyBucketsResult self = new ListAllMyBucketsResult();
            return TeaModel.build(map, self);
        }

        public ListAllMyBucketsResult setBuckets(Buckets buckets) {
            this.buckets = buckets;
            return this;
        }
        public Buckets getBuckets() {
            return this.buckets;
        }

        public ListAllMyBucketsResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListAllMyBucketsResult setMarker(String marker) {
            this.marker = marker;
            return this;
        }
        public String getMarker() {
            return this.marker;
        }

        public ListAllMyBucketsResult setMaxKeys(Long maxKeys) {
            this.maxKeys = maxKeys;
            return this;
        }
        public Long getMaxKeys() {
            return this.maxKeys;
        }

        public ListAllMyBucketsResult setNextMarker(String nextMarker) {
            this.nextMarker = nextMarker;
            return this;
        }
        public String getNextMarker() {
            return this.nextMarker;
        }

        public ListAllMyBucketsResult setOwner(Owner owner) {
            this.owner = owner;
            return this;
        }
        public Owner getOwner() {
            return this.owner;
        }

        public ListAllMyBucketsResult setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

    }

}
