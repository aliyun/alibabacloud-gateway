// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketRequest extends TeaModel {
    /**
     * <p>The character that is used to group objects by name. If you specify prefix and delimiter in the request, the response contains CommonPrefixes. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.\</p>
     * <p>By default, this parameter is left empty.</p>
     */
    @NameInMap("delimiter")
    public String delimiter;

    /**
     * <p>The encoding type of the content in the response.\</p>
     * <p>By default, this parameter is left empty.\</p>
     * <p>Set the value to URL.\</p>
     * <p>The values of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the value of Delimiter, Marker, Prefix, NextMarker, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>The name of the object after which the GetBucket (ListObjects) operation begins. If this parameter is specified, objects whose name is alphabetically after the value of marker are returned.\</p>
     * <p>The objects are returned by page based on marker. The value of marker must be less than 1,024 bytes in length.\</p>
     * <p>If the value of marker does not exist in the list when you perform a conditional query, the GetBucket (ListObjects) operation also starts from the object name that is alphabetically after the value of marker.\</p>
     * <p>By default, this parameter is left empty.</p>
     */
    @NameInMap("marker")
    public String marker;

    /**
     * <p>The maximum number of objects to be returned. If the number of returned objects exceeds the value of max-keys, use the value of NextMarker returned in the response as the value of marker in the next request.\</p>
     * <p>Valid values: 1 to 999.\</p>
     * <p>Default value: 100.</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>The prefix that the names of returned objects must contain.</p>
     * <br>
     * <p>*   The value of prefix must be less than 1,024 bytes in length.</p>
     * <p>*   If you specify prefix, the names of the returned objects contain the prefix.\</p>
     * <br>
     * <br>
     * <p>If you set prefix to a directory name, the objects whose name starts with the prefix are listed. The listed objects consist of all objects and subdirectories in the directory.\</p>
     * <p>If you specify prefix and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.\</p>
     * <p>For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If you set prefix to fun/, the three objects are returned. If you set prefix to fun/and delimiter to a forward slash (/), fun/test.jpg and fun/movie/are returned.\</p>
     * <p>By default, this parameter is left empty.</p>
     */
    @NameInMap("prefix")
    public String prefix;

    public static GetBucketRequest build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequest self = new GetBucketRequest();
        return TeaModel.build(map, self);
    }

    public GetBucketRequest setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public GetBucketRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public GetBucketRequest setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public GetBucketRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public GetBucketRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
