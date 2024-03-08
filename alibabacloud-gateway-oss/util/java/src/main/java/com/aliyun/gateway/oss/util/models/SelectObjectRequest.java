// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class SelectObjectRequest extends TeaModel {
    /**
     * <p>The container that stores the SelectObject request.</p>
     */
    @NameInMap("body")
    public SelectRequest body;

    public static SelectObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        SelectObjectRequest self = new SelectObjectRequest();
        return TeaModel.build(map, self);
    }

    public SelectObjectRequest setBody(SelectRequest body) {
        this.body = body;
        return this;
    }
    public SelectRequest getBody() {
        return this.body;
    }

}
