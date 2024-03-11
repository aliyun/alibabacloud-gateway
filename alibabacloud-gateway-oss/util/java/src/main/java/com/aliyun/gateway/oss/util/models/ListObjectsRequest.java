// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ListObjectsRequest extends TeaModel {
    /**
     * <p>The character that is used to group objects by name. If you specify delimiter in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("delimiter")
    public String delimiter;

    /**
     * <p>The encoding format of the content in the response.</p>
     * <br>
     * <p>>  The value of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the values of Delimiter, Marker, Prefix, NextMarker, and Key contain a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>The name of the object after which the GetBucket (ListObjects) operation begins. If this parameter is specified, objects whose names are alphabetically after the value of marker are returned.\</p>
     * <p>The objects are returned by page based on marker. The value of marker can be up to 1,024 bytes.\</p>
     * <p>If the value of marker does not exist in the list when you perform a conditional query, the GetBucket (ListObjects) operation starts from the object whose name is alphabetically after the value of marker.</p>
     */
    @NameInMap("marker")
    public String marker;

    /**
     * <p>The maximum number of objects that can be returned. If the number of objects to be returned exceeds the value of max-keys specified in the request, NextMarker is included in the returned response. The value of NextMarker is used as the value of marker for the next request.\</p>
     * <p>Valid values: 1 to 999.\</p>
     * <p>Default value: 100.</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>The prefix that must be contained in names of the returned objects.</p>
     * <br>
     * <p>*   The value of prefix can be up to 1,024 bytes in length.</p>
     * <p>*   If you specify prefix, the names of the returned objects contain the prefix.</p>
     * <br>
     * <p>If you set prefix to a directory name, the object whose names start with this prefix are listed. The objects consist of all recursive objects and subdirectories in this directory.\</p>
     * <p>If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are listed in CommonPrefixes. Recursive objects and subdirectories in the subdirectories are not listed.\</p>
     * <p>For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.</p>
     */
    @NameInMap("prefix")
    public String prefix;

    public static ListObjectsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListObjectsRequest self = new ListObjectsRequest();
        return TeaModel.build(map, self);
    }

    public ListObjectsRequest setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public ListObjectsRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListObjectsRequest setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListObjectsRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListObjectsRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
