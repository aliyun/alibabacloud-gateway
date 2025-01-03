// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PromoteDataLakeCacheReq extends TeaModel {
    @NameInMap("Object")
    public Object object;

    public static PromoteDataLakeCacheReq build(java.util.Map<String, ?> map) throws Exception {
        PromoteDataLakeCacheReq self = new PromoteDataLakeCacheReq();
        return TeaModel.build(map, self);
    }

    public PromoteDataLakeCacheReq setObject(Object object) {
        this.object = object;
        return this;
    }
    public Object getObject() {
        return this.object;
    }

    public static class Object extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>test.data</p>
         */
        @NameInMap("ObjectName")
        public String objectName;

        /**
         * <strong>example:</strong>
         * <p>1024-2048</p>
         */
        @NameInMap("Range")
        public String range;

        public static Object build(java.util.Map<String, ?> map) throws Exception {
            Object self = new Object();
            return TeaModel.build(map, self);
        }

        public Object setObjectName(String objectName) {
            this.objectName = objectName;
            return this;
        }
        public String getObjectName() {
            return this.objectName;
        }

        public Object setRange(String range) {
            this.range = range;
            return this;
        }
        public String getRange() {
            return this.range;
        }

    }

}
