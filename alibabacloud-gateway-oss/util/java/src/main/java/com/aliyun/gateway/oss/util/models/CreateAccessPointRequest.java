// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateAccessPointRequest extends TeaModel {
    /**
     * <p>The container that stores the information about the access point.</p>
     */
    @NameInMap("body")
    public CreateAccessPointConfiguration body;

    public static CreateAccessPointRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointRequest self = new CreateAccessPointRequest();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointRequest setBody(CreateAccessPointConfiguration body) {
        this.body = body;
        return this;
    }
    public CreateAccessPointConfiguration getBody() {
        return this.body;
    }

}
