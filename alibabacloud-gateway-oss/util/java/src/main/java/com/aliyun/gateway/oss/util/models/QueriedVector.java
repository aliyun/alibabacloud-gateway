// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class QueriedVector extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>32</p>
     */
    @NameInMap("distance")
    public Long distance;

    /**
     * <strong>example:</strong>
     * <p>test-vetor</p>
     */
    @NameInMap("key")
    public String key;

    /**
     * <strong>example:</strong>
     * <p>{&quot;key1&quot;:&quot;value&quot;}</p>
     */
    @NameInMap("metadata")
    public Object metadata;

    public static QueriedVector build(java.util.Map<String, ?> map) throws Exception {
        QueriedVector self = new QueriedVector();
        return TeaModel.build(map, self);
    }

    public QueriedVector setDistance(Long distance) {
        this.distance = distance;
        return this;
    }
    public Long getDistance() {
        return this.distance;
    }

    public QueriedVector setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public QueriedVector setMetadata(Object metadata) {
        this.metadata = metadata;
        return this;
    }
    public Object getMetadata() {
        return this.metadata;
    }

}
