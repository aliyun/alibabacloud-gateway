// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VectorIndexMetaData extends TeaModel {
    @NameInMap("nonFilterableMetadataKeys")
    public java.util.List<String> nonFilterableMetadataKeys;

    public static VectorIndexMetaData build(java.util.Map<String, ?> map) throws Exception {
        VectorIndexMetaData self = new VectorIndexMetaData();
        return TeaModel.build(map, self);
    }

    public VectorIndexMetaData setNonFilterableMetadataKeys(java.util.List<String> nonFilterableMetadataKeys) {
        this.nonFilterableMetadataKeys = nonFilterableMetadataKeys;
        return this;
    }
    public java.util.List<String> getNonFilterableMetadataKeys() {
        return this.nonFilterableMetadataKeys;
    }

}
