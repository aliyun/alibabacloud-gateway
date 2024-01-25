// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetObjectMetaRequest extends TeaModel {
    /**
     * <p>The version ID of the object. This element is displayed only if you query the metadata of an object of a specific version.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static GetObjectMetaRequest build(java.util.Map<String, ?> map) throws Exception {
        GetObjectMetaRequest self = new GetObjectMetaRequest();
        return TeaModel.build(map, self);
    }

    public GetObjectMetaRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
