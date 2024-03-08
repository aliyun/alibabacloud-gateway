// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PublicAccessBlockConfiguration extends TeaModel {
    @NameInMap("BlockPublicAccess")
    public Boolean blockPublicAccess;

    public static PublicAccessBlockConfiguration build(java.util.Map<String, ?> map) throws Exception {
        PublicAccessBlockConfiguration self = new PublicAccessBlockConfiguration();
        return TeaModel.build(map, self);
    }

    public PublicAccessBlockConfiguration setBlockPublicAccess(Boolean blockPublicAccess) {
        this.blockPublicAccess = blockPublicAccess;
        return this;
    }
    public Boolean getBlockPublicAccess() {
        return this.blockPublicAccess;
    }

}
