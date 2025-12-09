// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListVectorBucketsRequest extends TeaModel {
    @NameInMap("marker")
    public String marker;

    @NameInMap("max-keys")
    public Long maxKeys;

    @NameInMap("prefix")
    public String prefix;

    public static ListVectorBucketsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListVectorBucketsRequest self = new ListVectorBucketsRequest();
        return TeaModel.build(map, self);
    }

    public ListVectorBucketsRequest setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListVectorBucketsRequest setMaxKeys(Long maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Long getMaxKeys() {
        return this.maxKeys;
    }

    public ListVectorBucketsRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
