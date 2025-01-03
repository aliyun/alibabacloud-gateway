// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectsResult extends TeaModel {
    @NameInMap("Failed")
    public Failed failed;

    @NameInMap("Success")
    public Success success;

    public static CopyObjectsResult build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectsResult self = new CopyObjectsResult();
        return TeaModel.build(map, self);
    }

    public CopyObjectsResult setFailed(Failed failed) {
        this.failed = failed;
        return this;
    }
    public Failed getFailed() {
        return this.failed;
    }

    public CopyObjectsResult setSuccess(Success success) {
        this.success = success;
        return this;
    }
    public Success getSuccess() {
        return this.success;
    }

    public static class Failed extends TeaModel {
        @NameInMap("Object")
        public java.util.List<CopyObjectsResultFailedObject> object;

        public static Failed build(java.util.Map<String, ?> map) throws Exception {
            Failed self = new Failed();
            return TeaModel.build(map, self);
        }

        public Failed setObject(java.util.List<CopyObjectsResultFailedObject> object) {
            this.object = object;
            return this;
        }
        public java.util.List<CopyObjectsResultFailedObject> getObject() {
            return this.object;
        }

    }

    public static class Success extends TeaModel {
        @NameInMap("Object")
        public java.util.List<CopyObjectsResultSuccessObject> object;

        public static Success build(java.util.Map<String, ?> map) throws Exception {
            Success self = new Success();
            return TeaModel.build(map, self);
        }

        public Success setObject(java.util.List<CopyObjectsResultSuccessObject> object) {
            this.object = object;
            return this;
        }
        public java.util.List<CopyObjectsResultSuccessObject> getObject() {
            return this.object;
        }

    }

}
