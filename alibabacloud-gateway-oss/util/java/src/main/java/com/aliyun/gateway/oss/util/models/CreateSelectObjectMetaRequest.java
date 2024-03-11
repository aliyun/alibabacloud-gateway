// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class CreateSelectObjectMetaRequest extends TeaModel {
    /**
     * <p>The container that stores SelectMetaRequest information.</p>
     */
    @NameInMap("SelectMetaRequest")
    public SelectMetaRequest selectMetaRequest;

    public static CreateSelectObjectMetaRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateSelectObjectMetaRequest self = new CreateSelectObjectMetaRequest();
        return TeaModel.build(map, self);
    }

    public CreateSelectObjectMetaRequest setSelectMetaRequest(SelectMetaRequest selectMetaRequest) {
        this.selectMetaRequest = selectMetaRequest;
        return this;
    }
    public SelectMetaRequest getSelectMetaRequest() {
        return this.selectMetaRequest;
    }

}
