// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListObjectVersionsResponseBody extends TeaModel {
    /**
     * <p>If delimiter is specified in the request, the response contains CommonPrefixes. Objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("CommonPrefixes")
    public java.util.List<CommonPrefix> commonPrefixes;

    /**
     * <p>The container that stores the delete markers.</p>
     */
    @NameInMap("DeleteMarker")
    public java.util.List<DeleteMarkerEntry> deleteMarker;

    /**
     * <p>The character that is used to group objects by name. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("Delimiter")
    public String delimiter;

    /**
     * <p>The encoding type of the content in the response. If encoding-type is specified in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.</p>
     */
    @NameInMap("EncodingType")
    public String encodingType;

    /**
     * <p>Indicates whether the returned results are truncated.</p>
     * <br>
     * <p>*   true</p>
     * <p>*   false</p>
     * <br>
     * <p>Valid values: true and false.</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>The name of the object from which the GetBucketVersions (ListObjectVersions) operation begins.</p>
     */
    @NameInMap("KeyMarker")
    public String keyMarker;

    /**
     * <p>The maximum number of returned objects in the response.</p>
     */
    @NameInMap("MaxKeys")
    public Long maxKeys;

    /**
     * <p>The name of the bucket.</p>
     */
    @NameInMap("Name")
    public String name;

    /**
     * <p>If not all results are returned for the request, the response contains NextKeyMarker. Use the value of NextKeyMarker as the value of key-marker in the next request.</p>
     */
    @NameInMap("NextKeyMarker")
    public String nextKeyMarker;

    /**
     * <p>If not all results are returned for the request, the response contains NextVersionIdMarker. Use the value of NextVersionIdMarker as the value of version-id-marker in the next request.</p>
     */
    @NameInMap("NextVersionIdMarker")
    public String nextVersionIdMarker;

    /**
     * <p>The prefix in the names of the returned objects.</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    /**
     * <p>The container that stores the versions of objects, excluding the delete markers.</p>
     */
    @NameInMap("Version")
    public java.util.List<ObjectVersion> version;

    /**
     * <p>The version from which the GetBucketVersions (ListObjectVersions) operation begins. VersionIdMarker is returned together with KeyMarker.</p>
     */
    @NameInMap("VersionIdMarker")
    public String versionIdMarker;

    public static ListObjectVersionsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListObjectVersionsResponseBody self = new ListObjectVersionsResponseBody();
        return TeaModel.build(map, self);
    }

    public ListObjectVersionsResponseBody setCommonPrefixes(java.util.List<CommonPrefix> commonPrefixes) {
        this.commonPrefixes = commonPrefixes;
        return this;
    }
    public java.util.List<CommonPrefix> getCommonPrefixes() {
        return this.commonPrefixes;
    }

    public ListObjectVersionsResponseBody setDeleteMarker(java.util.List<DeleteMarkerEntry> deleteMarker) {
        this.deleteMarker = deleteMarker;
        return this;
    }
    public java.util.List<DeleteMarkerEntry> getDeleteMarker() {
        return this.deleteMarker;
    }

    public ListObjectVersionsResponseBody setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public ListObjectVersionsResponseBody setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListObjectVersionsResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListObjectVersionsResponseBody setKeyMarker(String keyMarker) {
        this.keyMarker = keyMarker;
        return this;
    }
    public String getKeyMarker() {
        return this.keyMarker;
    }

    public ListObjectVersionsResponseBody setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListObjectVersionsResponseBody setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public ListObjectVersionsResponseBody setNextKeyMarker(String nextKeyMarker) {
        this.nextKeyMarker = nextKeyMarker;
        return this;
    }
    public String getNextKeyMarker() {
        return this.nextKeyMarker;
    }

    public ListObjectVersionsResponseBody setNextVersionIdMarker(String nextVersionIdMarker) {
        this.nextVersionIdMarker = nextVersionIdMarker;
        return this;
    }
    public String getNextVersionIdMarker() {
        return this.nextVersionIdMarker;
    }

    public ListObjectVersionsResponseBody setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public ListObjectVersionsResponseBody setVersion(java.util.List<ObjectVersion> version) {
        this.version = version;
        return this;
    }
    public java.util.List<ObjectVersion> getVersion() {
        return this.version;
    }

    public ListObjectVersionsResponseBody setVersionIdMarker(String versionIdMarker) {
        this.versionIdMarker = versionIdMarker;
        return this;
    }
    public String getVersionIdMarker() {
        return this.versionIdMarker;
    }

}
