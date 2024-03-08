// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteAccessPointPolicyHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the access point.</p>
     */
    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static DeleteAccessPointPolicyHeaders build(java.util.Map<String, ?> map) throws Exception {
        DeleteAccessPointPolicyHeaders self = new DeleteAccessPointPolicyHeaders();
        return TeaModel.build(map, self);
    }

    public DeleteAccessPointPolicyHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public DeleteAccessPointPolicyHeaders setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
