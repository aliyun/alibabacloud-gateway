// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateAccessPointRequest extends TeaModel {
    /**
     * <p>The container that stores the information about the access point.</p>
     */
    @NameInMap("CreateAccessPointConfiguration")
    public CreateAccessPointConfiguration createAccessPointConfiguration;

    public static CreateAccessPointRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointRequest self = new CreateAccessPointRequest();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointRequest setCreateAccessPointConfiguration(CreateAccessPointConfiguration createAccessPointConfiguration) {
        this.createAccessPointConfiguration = createAccessPointConfiguration;
        return this;
    }
    public CreateAccessPointConfiguration getCreateAccessPointConfiguration() {
        return this.createAccessPointConfiguration;
    }

}
