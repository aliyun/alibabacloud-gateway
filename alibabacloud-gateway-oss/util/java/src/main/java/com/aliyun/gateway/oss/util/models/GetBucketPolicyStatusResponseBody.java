// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketPolicyStatusResponseBody extends TeaModel {
    @NameInMap("IsPublic")
    public Boolean isPublic;

    public static GetBucketPolicyStatusResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketPolicyStatusResponseBody self = new GetBucketPolicyStatusResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketPolicyStatusResponseBody setIsPublic(Boolean isPublic) {
        this.isPublic = isPublic;
        return this;
    }
    public Boolean getIsPublic() {
        return this.isPublic;
    }

}
