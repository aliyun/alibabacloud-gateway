// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateAccessPointResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the access point.</p>
     */
    @NameInMap("CreateAccessPointResult")
    public CreateAccessPointResult createAccessPointResult;

    public static CreateAccessPointResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointResponseBody self = new CreateAccessPointResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointResponseBody setCreateAccessPointResult(CreateAccessPointResult createAccessPointResult) {
        this.createAccessPointResult = createAccessPointResult;
        return this;
    }
    public CreateAccessPointResult getCreateAccessPointResult() {
        return this.createAccessPointResult;
    }

}
