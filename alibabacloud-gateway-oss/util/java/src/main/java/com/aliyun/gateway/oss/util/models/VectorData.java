// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class VectorData extends TeaModel {
    @NameInMap("float32")
    public java.util.List<Float> float32;

    public static VectorData build(java.util.Map<String, ?> map) throws Exception {
        VectorData self = new VectorData();
        return TeaModel.build(map, self);
    }

    public VectorData setFloat32(java.util.List<Float> float32) {
        this.float32 = float32;
        return this;
    }
    public java.util.List<Float> getFloat32() {
        return this.float32;
    }

}
