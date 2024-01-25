// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutObjectAclHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The name of the bucket.</p>
     */
    @NameInMap("x-oss-object-acl")
    public String xOssObjectAcl;

    public static PutObjectAclHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutObjectAclHeaders self = new PutObjectAclHeaders();
        return TeaModel.build(map, self);
    }

    public PutObjectAclHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutObjectAclHeaders setXOssObjectAcl(String xOssObjectAcl) {
        this.xOssObjectAcl = xOssObjectAcl;
        return this;
    }
    public String getXOssObjectAcl() {
        return this.xOssObjectAcl;
    }

}
