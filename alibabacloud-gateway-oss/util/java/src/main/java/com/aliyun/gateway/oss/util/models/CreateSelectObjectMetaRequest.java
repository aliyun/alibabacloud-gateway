// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateSelectObjectMetaRequest extends TeaModel {
    /**
     * <p>The container that stores SelectMetaRequest information.</p>
     */
    @NameInMap("body")
    public SelectMetaRequest body;

    public static CreateSelectObjectMetaRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateSelectObjectMetaRequest self = new CreateSelectObjectMetaRequest();
        return TeaModel.build(map, self);
    }

    public CreateSelectObjectMetaRequest setBody(SelectMetaRequest body) {
        this.body = body;
        return this;
    }
    public SelectMetaRequest getBody() {
        return this.body;
    }

}
