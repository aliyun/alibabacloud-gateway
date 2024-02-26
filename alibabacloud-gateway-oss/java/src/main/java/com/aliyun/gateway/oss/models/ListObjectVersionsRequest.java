// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListObjectVersionsRequest extends TeaModel {
    /**
     * <p>The character that is used to group objects by name. If you specify prefix and delimiter in the request, the response contains CommonPrefixes. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes. If you specify prefix and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.</p>
     * <br>
     * <p>By default, this parameter is left empty.</p>
     */
    @NameInMap("delimiter")
    public String delimiter;

    /**
     * <p>The encoding type of the content in the response. By default, this parameter is left empty. Set the value to URL.</p>
     * <br>
     * <p>>  The values of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the value of Delimiter, Marker, Prefix, NextMarker, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>The name of the object after which the GetBucketVersions (ListObjectVersions) operation begins. If this parameter is specified, objects whose name is alphabetically after the value of key-marker are returned. Use key-marker and version-id-marker in combination. The value of key-marker must be less than 1,024 bytes in length.</p>
     * <br>
     * <p>By default, this parameter is left empty.</p>
     * <br>
     * <p>>  You must also specify key-marker if you specify version-id-marker.</p>
     */
    @NameInMap("key-marker")
    public String keyMarker;

    /**
     * <p>The maximum number of objects to be returned. If the number of returned objects exceeds the value of max-keys, the response contains `NextKeyMarker` and `NextVersionIdMarker`. Specify the values of `NextKeyMarker` and `NextVersionIdMarker` as the markers for the next request. Valid values: 1 to 999. Default value: 100.</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>The prefix that the names of returned objects must contain.</p>
     * <br>
     * <p>*   The value of prefix must be less than 1,024 bytes in length.</p>
     * <p>*   If you specify prefix, the names of the returned objects contain the prefix.</p>
     * <br>
     * <p>If you set prefix to a directory name, the objects whose name starts with the prefix are listed. The returned objects consist of all objects and subdirectories in the directory.</p>
     * <br>
     * <p>By default, this parameter is left empty.</p>
     */
    @NameInMap("prefix")
    public String prefix;

    /**
     * <p>The version ID of the object specified in key-marker after which the GetBucketVersions (ListObjectVersions) operation begins. The versions are returned from the latest version to the earliest version. If version-id-marker is not specified, the GetBucketVersions (ListObjectVersions) operation starts from the latest version of the object whose name is alphabetically after the value of key-marker by default.</p>
     * <br>
     * <p>By default, this parameter is left empty.</p>
     * <br>
     * <p>Valid values: version IDs.</p>
     */
    @NameInMap("version-id-marker")
    public String versionIdMarker;

    public static ListObjectVersionsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListObjectVersionsRequest self = new ListObjectVersionsRequest();
        return TeaModel.build(map, self);
    }

    public ListObjectVersionsRequest setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public ListObjectVersionsRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListObjectVersionsRequest setKeyMarker(String keyMarker) {
        this.keyMarker = keyMarker;
        return this;
    }
    public String getKeyMarker() {
        return this.keyMarker;
    }

    public ListObjectVersionsRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListObjectVersionsRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public ListObjectVersionsRequest setVersionIdMarker(String versionIdMarker) {
        this.versionIdMarker = versionIdMarker;
        return this;
    }
    public String getVersionIdMarker() {
        return this.versionIdMarker;
    }

}
