// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketArchiveDirectReadRequest extends TeaModel {
    @NameInMap("body")
    public ArchiveDirectReadConfiguration body;

    public static PutBucketArchiveDirectReadRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketArchiveDirectReadRequest self = new PutBucketArchiveDirectReadRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketArchiveDirectReadRequest setBody(ArchiveDirectReadConfiguration body) {
        this.body = body;
        return this;
    }
    public ArchiveDirectReadConfiguration getBody() {
        return this.body;
    }

}
