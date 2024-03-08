// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ListObjectsV2Request extends TeaModel {
    /**
     * <p>The token from which the list operation starts. You can obtain the token from NextContinuationToken in the response of the ListObjectsV2 request.</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    /**
     * <p>The character that is used to group objects by name. If you specify delimiter in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.</p>
     */
    @NameInMap("delimiter")
    public String delimiter;

    /**
     * <p>The encoding format of the returned objects in the response.</p>
     * <br>
     * <p>>  The values of Delimiter, StartAfter, Prefix, NextContinuationToken, and Key are UTF-8 encoded. If the value of Delimiter, StartAfter, Prefix, NextContinuationToken, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>Specifies whether to include the information about the bucket owner in the response. Valid values:</p>
     * <br>
     * <p>*   true</p>
     * <p>*   false</p>
     */
    @NameInMap("fetch-owner")
    public Boolean fetchOwner;

    /**
     * <p>The maximum number of objects to be returned.\</p>
     * <p>Valid values: 1 to 999.\</p>
     * <p>Default value: 100.</p>
     * <br>
     * <p>>  If the number of returned objects exceeds the value of max-keys, the response contains NextContinuationToken.Use the value of NextContinuationToken as the value of continuation-token in the next request.</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>The prefix that must be contained in names of the returned objects.\</p>
     * <br>
     * <br>
     * <p>*   The value of prefix can be up to 1,024 bytes in length.</p>
     * <p>*   If you specify prefix, the names of the returned objects contain the prefix.</p>
     * <br>
     * <p>If you set prefix to a directory name, the objects whose names start with this prefix are listed. The objects consist of all objects and subdirectories in this directory.\</p>
     * <p>If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.\</p>
     * <p>For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.</p>
     */
    @NameInMap("prefix")
    public String prefix;

    /**
     * <p>The name of the object after which the list operation begins. If this parameter is specified, objects whose names are alphabetically after the value of start-after are returned.\</p>
     * <p>The objects are returned by page based on start-after. The value of start-after can be up to 1,024 bytes in length.\</p>
     * <p>If the value of start-after does not exist when you perform a conditional query, the list starts from the object whose name is alphabetically after the value of start-after.</p>
     */
    @NameInMap("start-after")
    public String startAfter;

    public static ListObjectsV2Request build(java.util.Map<String, ?> map) throws Exception {
        ListObjectsV2Request self = new ListObjectsV2Request();
        return TeaModel.build(map, self);
    }

    public ListObjectsV2Request setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListObjectsV2Request setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public ListObjectsV2Request setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListObjectsV2Request setFetchOwner(Boolean fetchOwner) {
        this.fetchOwner = fetchOwner;
        return this;
    }
    public Boolean getFetchOwner() {
        return this.fetchOwner;
    }

    public ListObjectsV2Request setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListObjectsV2Request setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public ListObjectsV2Request setStartAfter(String startAfter) {
        this.startAfter = startAfter;
        return this;
    }
    public String getStartAfter() {
        return this.startAfter;
    }

}
