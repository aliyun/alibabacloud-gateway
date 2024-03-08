// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketResponseBody extends TeaModel {
    /**
     * <p>If delimiter is specified in the request, the response contains CommonPrefixes. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("CommonPrefixes")
    public java.util.List<CommonPrefix> commonPrefixes;

    /**
     * <p>The container that stores the metadata of the returned objects.</p>
     */
    @NameInMap("Contents")
    public java.util.List<ObjectSummary> contents;

    /**
     * <p>The character that is used to group objects by name. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("Delimiter")
    public String delimiter;

    /**
     * <p>Indicates whether the returned results are truncated.</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>The name of the object from which the GetBucket (ListObjects) operation begins.</p>
     */
    @NameInMap("Marker")
    public String marker;

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
     * <p>The prefix in the names of the returned objects.</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    public static GetBucketResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResponseBody self = new GetBucketResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketResponseBody setCommonPrefixes(java.util.List<CommonPrefix> commonPrefixes) {
        this.commonPrefixes = commonPrefixes;
        return this;
    }
    public java.util.List<CommonPrefix> getCommonPrefixes() {
        return this.commonPrefixes;
    }

    public GetBucketResponseBody setContents(java.util.List<ObjectSummary> contents) {
        this.contents = contents;
        return this;
    }
    public java.util.List<ObjectSummary> getContents() {
        return this.contents;
    }

    public GetBucketResponseBody setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public GetBucketResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public GetBucketResponseBody setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public GetBucketResponseBody setMaxKeys(Integer maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Integer getMaxKeys() {
        return this.maxKeys;
    }

    public GetBucketResponseBody setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public GetBucketResponseBody setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
