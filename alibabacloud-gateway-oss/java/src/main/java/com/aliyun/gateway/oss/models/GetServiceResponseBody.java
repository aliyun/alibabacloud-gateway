// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetServiceResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the bucket owner.</p>
     */
    @NameInMap("ListAllMyBucketsResult")
    public ListAllMyBucketsResult listAllMyBucketsResult;

    public static GetServiceResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetServiceResponseBody self = new GetServiceResponseBody();
        return TeaModel.build(map, self);
    }

    public GetServiceResponseBody setListAllMyBucketsResult(ListAllMyBucketsResult listAllMyBucketsResult) {
        this.listAllMyBucketsResult = listAllMyBucketsResult;
        return this;
    }
    public ListAllMyBucketsResult getListAllMyBucketsResult() {
        return this.listAllMyBucketsResult;
    }

    public static class Buckets extends TeaModel {
        /**
         * <p>The container that stores the information about a single bucket.</p>
         */
        @NameInMap("Bucket")
        public java.util.List<Bucket> buckets;

        public static Buckets build(java.util.Map<String, ?> map) throws Exception {
            Buckets self = new Buckets();
            return TeaModel.build(map, self);
        }

        public Buckets setBuckets(java.util.List<Bucket> buckets) {
            this.buckets = buckets;
            return this;
        }
        public java.util.List<Bucket> getBuckets() {
            return this.buckets;
        }

    }

    public static class ListAllMyBucketsResult extends TeaModel {
        /**
         * <p>The container that stores the information about buckets.</p>
         */
        @NameInMap("Buckets")
        public Buckets buckets;

        /**
         * <p>Indicates whether all results are returned. Valid values:</p>
         * <br>
         * <p>- true</p>
         * <p>- false</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The pagination token used in the current request.</p>
         */
        @NameInMap("Marker")
        public String marker;

        /**
         * <p>The maximum number of buckets that can be returned.</p>
         */
        @NameInMap("MaxKeys")
        public Long maxKeys;

        /**
         * <p>A pagination token. It can be used in the next request to retrieve a new page of results.</p>
         */
        @NameInMap("NextMarker")
        public String nextMarker;

        /**
         * <p>The prefix that the names of returned buckets contain.</p>
         */
        @NameInMap("Owner")
        public Owner owner;

        /**
         * <p>The prefix that the names of returned buckets contain.</p>
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
