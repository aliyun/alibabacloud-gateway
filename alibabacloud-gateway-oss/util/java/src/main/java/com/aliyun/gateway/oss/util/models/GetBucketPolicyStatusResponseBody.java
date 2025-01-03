// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketPolicyStatusResponseBody extends TeaModel {
    /**
     * <p>The container that stores public access information.</p>
     */
    @NameInMap("PolicyStatus")
    public PolicyStatus policyStatus;

    public static GetBucketPolicyStatusResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketPolicyStatusResponseBody self = new GetBucketPolicyStatusResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketPolicyStatusResponseBody setPolicyStatus(PolicyStatus policyStatus) {
        this.policyStatus = policyStatus;
        return this;
    }
    public PolicyStatus getPolicyStatus() {
        return this.policyStatus;
    }

    public static class PolicyStatus extends TeaModel {
        /**
         * <p>Indicates whether the current bucket policy allows public access.</p>
         * <p>true</p>
         * <p>false</p>
         * 
         * <strong>example:</strong>
         * <p>true</p>
         */
        @NameInMap("IsPublic")
        public Boolean isPublic;

        public static PolicyStatus build(java.util.Map<String, ?> map) throws Exception {
            PolicyStatus self = new PolicyStatus();
            return TeaModel.build(map, self);
        }

        public PolicyStatus setIsPublic(Boolean isPublic) {
            this.isPublic = isPublic;
            return this;
        }
        public Boolean getIsPublic() {
            return this.isPublic;
        }

    }

}
