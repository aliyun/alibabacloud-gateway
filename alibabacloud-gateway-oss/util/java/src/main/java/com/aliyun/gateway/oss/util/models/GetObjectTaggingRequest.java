// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetObjectTaggingRequest extends TeaModel {
    /**
     * <p>The version ID of the object that you want to query.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static GetObjectTaggingRequest build(java.util.Map<String, ?> map) throws Exception {
        GetObjectTaggingRequest self = new GetObjectTaggingRequest();
        return TeaModel.build(map, self);
    }

    public GetObjectTaggingRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
