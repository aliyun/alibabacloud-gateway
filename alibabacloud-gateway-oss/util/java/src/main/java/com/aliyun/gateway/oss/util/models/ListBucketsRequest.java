// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketsRequest extends TeaModel {
    /**
     * <p>The name of the bucket from which the buckets start to return. The buckets whose names are alphabetically after the value of marker are returned. If this parameter is not specified, all results are returned. By default, this parameter is left empty.</p>
     * 
     * <strong>example:</strong>
     * <p>mybucket10</p>
     */
    @NameInMap("marker")
    public String marker;

    /**
     * <p>The maximum number of buckets that can be returned. Valid values: 1 to 1000. Default value: 100</p>
     * 
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("max-keys")
    public Long maxKeys;

    /**
     * <p>The prefix that the names of returned buckets must contain. If this parameter is not specified, prefixes are not used to filter returned buckets. By default, this parameter is left empty.</p>
     * 
     * <strong>example:</strong>
     * <p>my</p>
     */
    @NameInMap("prefix")
    public String prefix;

    @NameInMap("tag-key")
    public String tagKey;

    @NameInMap("tag-value")
    public String tagValue;

    @NameInMap("tagging")
    public String tagging;

    public static ListBucketsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListBucketsRequest self = new ListBucketsRequest();
        return TeaModel.build(map, self);
    }

    public ListBucketsRequest setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListBucketsRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListBucketsRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public ListBucketsRequest setTagKey(String tagKey) {
        this.tagKey = tagKey;
        return this;
    }
    public String getTagKey() {
        return this.tagKey;
    }

    public ListBucketsRequest setTagValue(String tagValue) {
        this.tagValue = tagValue;
        return this;
    }
    public String getTagValue() {
        return this.tagValue;
    }

    public ListBucketsRequest setTagging(String tagging) {
        this.tagging = tagging;
        return this;
    }
    public String getTagging() {
        return this.tagging;
    }

}
