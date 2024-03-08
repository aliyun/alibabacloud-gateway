// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteObjectTaggingRequest extends TeaModel {
    /**
     * <p>The version ID of the object that you want to delete.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static DeleteObjectTaggingRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteObjectTaggingRequest self = new DeleteObjectTaggingRequest();
        return TeaModel.build(map, self);
    }

    public DeleteObjectTaggingRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
