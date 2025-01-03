// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectsReq extends TeaModel {
    @NameInMap("Object")
    public java.util.List<Object> object;

    public static GetObjectsReq build(java.util.Map<String, ?> map) throws Exception {
        GetObjectsReq self = new GetObjectsReq();
        return TeaModel.build(map, self);
    }

    public GetObjectsReq setObject(java.util.List<Object> object) {
        this.object = object;
        return this;
    }
    public java.util.List<Object> getObject() {
        return this.object;
    }

    public static class Object extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>test.txt</p>
         */
        @NameInMap("ObjectName")
        public String objectName;

        /**
         * <strong>example:</strong>
         * <p>bytes=20-60</p>
         */
        @NameInMap("Range")
        public String range;

        /**
         * <strong>example:</strong>
         * <p>2</p>
         */
        @NameInMap("RefId")
        public Integer refId;

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

        public Object setRefId(Integer refId) {
            this.refId = refId;
            return this;
        }
        public Integer getRefId() {
            return this.refId;
        }

    }

}
