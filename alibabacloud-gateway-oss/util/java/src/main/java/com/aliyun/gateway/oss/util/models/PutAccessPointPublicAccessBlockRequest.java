// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutAccessPointPublicAccessBlockRequest extends TeaModel {
    @NameInMap("body")
    public PublicAccessBlockConfiguration body;

    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static PutAccessPointPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointPublicAccessBlockRequest self = new PutAccessPointPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointPublicAccessBlockRequest setBody(PublicAccessBlockConfiguration body) {
        this.body = body;
        return this;
    }
    public PublicAccessBlockConfiguration getBody() {
        return this.body;
    }

    public PutAccessPointPublicAccessBlockRequest setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
