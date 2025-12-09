// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class Vector extends TeaModel {
    @NameInMap("data")
    public VectorData data;

    /**
     * <strong>example:</strong>
     * <p>test-key</p>
     */
    @NameInMap("key")
    public String key;

    /**
     * <strong>example:</strong>
     * <p>{&quot;key1&quot;:&quot;value1&quot;}</p>
     */
    @NameInMap("metadata")
    public Object metadata;

    public static Vector build(java.util.Map<String, ?> map) throws Exception {
        Vector self = new Vector();
        return TeaModel.build(map, self);
    }

    public Vector setData(VectorData data) {
        this.data = data;
        return this;
    }
    public VectorData getData() {
        return this.data;
    }

    public Vector setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public Vector setMetadata(Object metadata) {
        this.metadata = metadata;
        return this;
    }
    public Object getMetadata() {
        return this.metadata;
    }

}
