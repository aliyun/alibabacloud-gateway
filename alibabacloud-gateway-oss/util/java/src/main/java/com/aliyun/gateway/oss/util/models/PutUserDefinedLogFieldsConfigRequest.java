// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutUserDefinedLogFieldsConfigRequest extends TeaModel {
    @NameInMap("body")
    public UserDefinedLogFieldsConfiguration body;

    public static PutUserDefinedLogFieldsConfigRequest build(java.util.Map<String, ?> map) throws Exception {
        PutUserDefinedLogFieldsConfigRequest self = new PutUserDefinedLogFieldsConfigRequest();
        return TeaModel.build(map, self);
    }

    public PutUserDefinedLogFieldsConfigRequest setBody(UserDefinedLogFieldsConfiguration body) {
        this.body = body;
        return this;
    }
    public UserDefinedLogFieldsConfiguration getBody() {
        return this.body;
    }

}
