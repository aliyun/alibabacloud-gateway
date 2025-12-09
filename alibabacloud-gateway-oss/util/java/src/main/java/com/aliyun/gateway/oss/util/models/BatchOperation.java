// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperation extends TeaModel {
    @NameInMap("CopyObject")
    public BatchCopyObject copyObject;

    @NameInMap("DeleteObjectTagging")
    public Object deleteObjectTagging;

    @NameInMap("PutObjectAcl")
    public BatchPutObjectAcl putObjectAcl;

    @NameInMap("PutObjectTagging")
    public BatchPutObjectTagging putObjectTagging;

    @NameInMap("RestoreObject")
    public BatchRestoreObject restoreObject;

    public static BatchOperation build(java.util.Map<String, ?> map) throws Exception {
        BatchOperation self = new BatchOperation();
        return TeaModel.build(map, self);
    }

    public BatchOperation setCopyObject(BatchCopyObject copyObject) {
        this.copyObject = copyObject;
        return this;
    }
    public BatchCopyObject getCopyObject() {
        return this.copyObject;
    }

    public BatchOperation setDeleteObjectTagging(Object deleteObjectTagging) {
        this.deleteObjectTagging = deleteObjectTagging;
        return this;
    }
    public Object getDeleteObjectTagging() {
        return this.deleteObjectTagging;
    }

    public BatchOperation setPutObjectAcl(BatchPutObjectAcl putObjectAcl) {
        this.putObjectAcl = putObjectAcl;
        return this;
    }
    public BatchPutObjectAcl getPutObjectAcl() {
        return this.putObjectAcl;
    }

    public BatchOperation setPutObjectTagging(BatchPutObjectTagging putObjectTagging) {
        this.putObjectTagging = putObjectTagging;
        return this;
    }
    public BatchPutObjectTagging getPutObjectTagging() {
        return this.putObjectTagging;
    }

    public BatchOperation setRestoreObject(BatchRestoreObject restoreObject) {
        this.restoreObject = restoreObject;
        return this;
    }
    public BatchRestoreObject getRestoreObject() {
        return this.restoreObject;
    }

}
