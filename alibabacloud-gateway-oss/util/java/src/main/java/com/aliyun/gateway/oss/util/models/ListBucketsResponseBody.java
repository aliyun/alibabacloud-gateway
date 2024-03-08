// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListBucketsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about multiple buckets.</p>
     */
    @NameInMap("buckets")
    public java.util.List<Bucket> buckets;

    /**
     * <p>Indicates whether all results are returned. Valid values: true: All results are not returned in the response. false: All results are returned in the response.</p>
     */
    @NameInMap("isTruncated")
    public Boolean isTruncated;

    /**
     * <p>The name of the bucket from which the buckets are returned.</p>
     */
    @NameInMap("marker")
    public String marker;

    /**
     * <p>The maximum number of buckets that can be returned.</p>
     */
    @NameInMap("maxKeys")
    public Long maxKeys;

    /**
     * <p>The marker for the next ListBuckets (GetService) request. You can use the value of this parameter as the value of marker in the next ListBuckets (GetService) request to retrieve the unreturned results.</p>
     */
    @NameInMap("nextMarker")
    public String nextMarker;

    /**
     * <p>The container that stores the information about the bucket owner.</p>
     */
    @NameInMap("owner")
    public Owner owner;

    /**
     * <p>The prefix contained in the names of returned buckets.</p>
     */
    @NameInMap("prefix")
    public String prefix;

    public static ListBucketsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketsResponseBody self = new ListBucketsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketsResponseBody setBuckets(java.util.List<Bucket> buckets) {
        this.buckets = buckets;
        return this;
    }
    public java.util.List<Bucket> getBuckets() {
        return this.buckets;
    }

    public ListBucketsResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListBucketsResponseBody setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListBucketsResponseBody setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListBucketsResponseBody setNextMarker(String nextMarker) {
        this.nextMarker = nextMarker;
        return this;
    }
    public String getNextMarker() {
        return this.nextMarker;
    }

    public ListBucketsResponseBody setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public ListBucketsResponseBody setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
