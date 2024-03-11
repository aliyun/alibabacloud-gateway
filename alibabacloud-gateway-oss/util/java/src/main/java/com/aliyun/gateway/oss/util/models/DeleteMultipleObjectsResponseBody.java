// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class DeleteMultipleObjectsResponseBody extends TeaModel {
    @NameInMap("DeleteResult")
    public DeleteResult deleteResult;

    public static DeleteMultipleObjectsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DeleteMultipleObjectsResponseBody self = new DeleteMultipleObjectsResponseBody();
        return TeaModel.build(map, self);
    }

    public DeleteMultipleObjectsResponseBody setDeleteResult(DeleteResult deleteResult) {
        this.deleteResult = deleteResult;
        return this;
    }
    public DeleteResult getDeleteResult() {
        return this.deleteResult;
    }

    public static class DeleteResult extends TeaModel {
        @NameInMap("Deleted")
        public java.util.List<DeletedObject> deleted;

        @NameInMap("EncodingType")
        public String encodingType;

        public static DeleteResult build(java.util.Map<String, ?> map) throws Exception {
            DeleteResult self = new DeleteResult();
            return TeaModel.build(map, self);
        }

        public DeleteResult setDeleted(java.util.List<DeletedObject> deleted) {
            this.deleted = deleted;
            return this;
        }
        public java.util.List<DeletedObject> getDeleted() {
            return this.deleted;
        }

        public DeleteResult setEncodingType(String encodingType) {
            this.encodingType = encodingType;
            return this;
        }
        public String getEncodingType() {
            return this.encodingType;
        }

    }

}
