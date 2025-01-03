// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsCopy extends TeaModel {
    @NameInMap("Object")
    public java.util.List<Object> object;

    public static CopyObjectsCopy build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsCopy self = new CopyObjectsCopy();
        return TeaModel.build(map, self);
    }

    public CopyObjectsCopy setObject(java.util.List<Object> object) {
        this.object = object;
        return this;
    }
    public java.util.List<Object> getObject() {
        return this.object;
    }

    public static class Object extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>abc/source.txt</p>
         */
        @NameInMap("SourceKey")
        public String sourceKey;

        /**
         * <strong>example:</strong>
         * <p>def/target.txt</p>
         */
        @NameInMap("TargetKey")
        public String targetKey;

        public static Object build(java.util.Map<String, ?> map) throws Exception {
            Object self = new Object();
            return TeaModel.build(map, self);
        }

        public Object setSourceKey(String sourceKey) {
            this.sourceKey = sourceKey;
            return this;
        }
        public String getSourceKey() {
            return this.sourceKey;
        }

        public Object setTargetKey(String targetKey) {
            this.targetKey = targetKey;
            return this;
        }
        public String getTargetKey() {
            return this.targetKey;
        }

    }

}
