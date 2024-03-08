// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class RestoreObjectRequest extends TeaModel {
    /**
     * <p>Request information of the restoration.</p>
     */
    @NameInMap("body")
    public RestoreRequest body;

    /**
     * <p>The version number of the object that you want to restore.</p>
     */
    @NameInMap("versionId")
    public String versionId;

    public static RestoreObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        RestoreObjectRequest self = new RestoreObjectRequest();
        return TeaModel.build(map, self);
    }

    public RestoreObjectRequest setBody(RestoreRequest body) {
        this.body = body;
        return this;
    }
    public RestoreRequest getBody() {
        return this.body;
    }

    public RestoreObjectRequest setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
