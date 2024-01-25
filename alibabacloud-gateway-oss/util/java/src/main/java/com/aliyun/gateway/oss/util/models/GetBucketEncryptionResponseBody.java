// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketEncryptionResponseBody extends TeaModel {
    /**
     * <p>The container that stores the default server-side encryption method.</p>
     */
    @NameInMap("ApplyServerSideEncryptionByDefault")
    public ApplyServerSideEncryptionByDefault applyServerSideEncryptionByDefault;

    public static GetBucketEncryptionResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketEncryptionResponseBody self = new GetBucketEncryptionResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketEncryptionResponseBody setApplyServerSideEncryptionByDefault(ApplyServerSideEncryptionByDefault applyServerSideEncryptionByDefault) {
        this.applyServerSideEncryptionByDefault = applyServerSideEncryptionByDefault;
        return this;
    }
    public ApplyServerSideEncryptionByDefault getApplyServerSideEncryptionByDefault() {
        return this.applyServerSideEncryptionByDefault;
    }

}
