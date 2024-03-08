// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListObjectsResponseBody extends TeaModel {
    /**
     * <p>If delimiter is specified in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("CommonPrefixes")
    public java.util.List<CommonPrefix> commonPrefixes;

    /**
     * <p>The container that stores the metadata of the returned objects.</p>
     */
    @NameInMap("Contents")
    public java.util.List<ObjectSummary> contents;

    /**
     * <p>The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("Delimiter")
    public String delimiter;

    /**
     * <p>The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.</p>
     */
    @NameInMap("EncodingType")
    public String encodingType;

    /**
     * <p>Indicates whether the returned list in the result is truncated. Valid values:</p>
     * <br>
     * <p>*   true</p>
     * <p>*   false</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>The name of the object after which the GetBucket (ListObjects) operation begins.</p>
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
     * <p>If not all results are returned, NextMarker is included in the response to indicate the value of marker in the next request.</p>
     */
    @NameInMap("NextMarker")
    public String nextMarker;

    /**
     * <p>The prefix in the names of the returned objects.</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    public static ListObjectsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListObjectsResponseBody self = new ListObjectsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListObjectsResponseBody setCommonPrefixes(java.util.List<CommonPrefix> commonPrefixes) {
        this.commonPrefixes = commonPrefixes;
        return this;
    }
    public java.util.List<CommonPrefix> getCommonPrefixes() {
        return this.commonPrefixes;
    }

    public ListObjectsResponseBody setContents(java.util.List<ObjectSummary> contents) {
        this.contents = contents;
        return this;
    }
    public java.util.List<ObjectSummary> getContents() {
        return this.contents;
    }

    public ListObjectsResponseBody setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public ListObjectsResponseBody setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListObjectsResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListObjectsResponseBody setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListObjectsResponseBody setMaxKeys(Integer maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Integer getMaxKeys() {
        return this.maxKeys;
    }

    public ListObjectsResponseBody setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public ListObjectsResponseBody setNextMarker(String nextMarker) {
        this.nextMarker = nextMarker;
        return this;
    }
    public String getNextMarker() {
        return this.nextMarker;
    }

    public ListObjectsResponseBody setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
