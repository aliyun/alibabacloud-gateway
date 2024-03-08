// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutObjectTaggingRequest extends TeaModel {
    /**
     * <p>The container used to store Tagset.</p>
     */
    @NameInMap("body")
    public Tagging body;

    /**
     * <p>The ID of the version.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static PutObjectTaggingRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectTaggingRequest self = new PutObjectTaggingRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectTaggingRequest setBody(Tagging body) {
        this.body = body;
        return this;
    }
    public Tagging getBody() {
        return this.body;
    }

    public PutObjectTaggingRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
