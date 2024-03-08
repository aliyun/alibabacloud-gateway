// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DoMetaQueryRequest extends TeaModel {
    /**
     * <p>The container that stores the query conditions.</p>
     */
    @NameInMap("body")
    public MetaQuery body;

    public static DoMetaQueryRequest build(java.util.Map<String, ?> map) throws Exception {
        DoMetaQueryRequest self = new DoMetaQueryRequest();
        return TeaModel.build(map, self);
    }

    public DoMetaQueryRequest setBody(MetaQuery body) {
        this.body = body;
        return this;
    }
    public MetaQuery getBody() {
        return this.body;
    }

}
