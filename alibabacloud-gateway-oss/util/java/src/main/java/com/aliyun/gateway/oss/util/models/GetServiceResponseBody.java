// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetServiceResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about buckets.</p>
     */
    @NameInMap("Buckets")
    public GetServiceResponseBodyBuckets buckets;

    /**
     * <p>Indicates whether all results are returned. Valid values:</p>
     * <br>
     * <p>*   true</p>
     * <p>*   false</p>
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
     * <p>The container that stores the information about the bucket owner.</p>
     */
    @NameInMap("Owner")
    public Owner owner;

    /**
     * <p>The prefix that the names of returned buckets contain.</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    public static GetServiceResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetServiceResponseBody self = new GetServiceResponseBody();
        return TeaModel.build(map, self);
    }

    public GetServiceResponseBody setBuckets(GetServiceResponseBodyBuckets buckets) {
        this.buckets = buckets;
        return this;
    }
    public GetServiceResponseBodyBuckets getBuckets() {
        return this.buckets;
    }

    public GetServiceResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public GetServiceResponseBody setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public GetServiceResponseBody setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public GetServiceResponseBody setNextMarker(String nextMarker) {
        this.nextMarker = nextMarker;
        return this;
    }
    public String getNextMarker() {
        return this.nextMarker;
    }

    public GetServiceResponseBody setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public GetServiceResponseBody setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public static class GetServiceResponseBodyBuckets extends TeaModel {
        /**
         * <p>The container that stores the information about a single bucket.</p>
         */
        @NameInMap("Bucket")
        public java.util.List<Bucket> bucket;

        public static GetServiceResponseBodyBuckets build(java.util.Map<String, ?> map) throws Exception {
            GetServiceResponseBodyBuckets self = new GetServiceResponseBodyBuckets();
            return TeaModel.build(map, self);
        }

        public GetServiceResponseBodyBuckets setBucket(java.util.List<Bucket> bucket) {
            this.bucket = bucket;
            return this;
        }
        public java.util.List<Bucket> getBucket() {
            return this.bucket;
        }

    }

}
