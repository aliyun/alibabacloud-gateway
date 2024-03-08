// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutStyleRequest extends TeaModel {
    /**
     * <p>The container that stores the content information about the image style.</p>
     */
    @NameInMap("body")
    public Style body;

    /**
     * <p>The name of the image style.</p>
     */
    @NameInMap("styleName")
    public String styleName;

    public static PutStyleRequest build(java.util.Map<String, ?> map) throws Exception {
        PutStyleRequest self = new PutStyleRequest();
        return TeaModel.build(map, self);
    }

    public PutStyleRequest setBody(Style body) {
        this.body = body;
        return this;
    }
    public Style getBody() {
        return this.body;
    }

    public PutStyleRequest setStyleName(String styleName) {
        this.styleName = styleName;
        return this;
    }
    public String getStyleName() {
        return this.styleName;
    }

}
