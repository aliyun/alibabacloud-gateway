// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorBucketsResponseBody extends TeaModel {
    @NameInMap("ListAllMyBucketsResult")
    public ListAllMyBucketsResult listAllMyBucketsResult;

    public static ListVectorBucketsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListVectorBucketsResponseBody self = new ListVectorBucketsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListVectorBucketsResponseBody setListAllMyBucketsResult(ListAllMyBucketsResult listAllMyBucketsResult) {
        this.listAllMyBucketsResult = listAllMyBucketsResult;
        return this;
    }
    public ListAllMyBucketsResult getListAllMyBucketsResult() {
        return this.listAllMyBucketsResult;
    }

    public static class Buckets extends TeaModel {
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
        @NameInMap("Buckets")
        public Buckets buckets;

        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        @NameInMap("Marker")
        public String marker;

        @NameInMap("MaxKeys")
        public Long maxKeys;

        @NameInMap("NextMarker")
        public String nextMarker;

        @NameInMap("Owner")
        public Owner owner;

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
