// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteMultipleObjectsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the deleted objects.</p>
     */
    @NameInMap("Deleted")
    public java.util.List<DeletedObject> deleted;

    /**
     * <p>The encoding type for the returned results. If encoding-type is specified in the request, the object name is encoded in the returned result.</p>
     */
    @NameInMap("EncodingType")
    public String encodingType;

    public static DeleteMultipleObjectsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DeleteMultipleObjectsResponseBody self = new DeleteMultipleObjectsResponseBody();
        return TeaModel.build(map, self);
    }

    public DeleteMultipleObjectsResponseBody setDeleted(java.util.List<DeletedObject> deleted) {
        this.deleted = deleted;
        return this;
    }
    public java.util.List<DeletedObject> getDeleted() {
        return this.deleted;
    }

    public DeleteMultipleObjectsResponseBody setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

}
