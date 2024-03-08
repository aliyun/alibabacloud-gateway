// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketResourceGroupResponseBody extends TeaModel {
    /**
     * <p>The ID of the resource group to which the bucket belongs.</p>
     */
    @NameInMap("ResourceGroupId")
    public String resourceGroupId;

    public static GetBucketResourceGroupResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketResourceGroupResponseBody self = new GetBucketResourceGroupResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketResourceGroupResponseBody setResourceGroupId(String resourceGroupId) {
        this.resourceGroupId = resourceGroupId;
        return this;
    }
    public String getResourceGroupId() {
        return this.resourceGroupId;
    }

}
