// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ObjectHashConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("DisplayObjectHash")
    public Boolean displayObjectHash;

    /**
     * <strong>example:</strong>
     * <p>SHA-256</p>
     */
    @NameInMap("ObjectHashFunction")
    public String objectHashFunction;

    public static ObjectHashConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ObjectHashConfiguration self = new ObjectHashConfiguration();
        return TeaModel.build(map, self);
    }

    public ObjectHashConfiguration setDisplayObjectHash(Boolean displayObjectHash) {
        this.displayObjectHash = displayObjectHash;
        return this;
    }
    public Boolean getDisplayObjectHash() {
        return this.displayObjectHash;
    }

    public ObjectHashConfiguration setObjectHashFunction(String objectHashFunction) {
        this.objectHashFunction = objectHashFunction;
        return this;
    }
    public String getObjectHashFunction() {
        return this.objectHashFunction;
    }

}
