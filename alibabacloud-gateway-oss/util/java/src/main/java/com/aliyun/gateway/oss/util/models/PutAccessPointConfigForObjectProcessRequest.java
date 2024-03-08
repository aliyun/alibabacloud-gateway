// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutAccessPointConfigForObjectProcessRequest extends TeaModel {
    @NameInMap("ObjectProcessConfiguration")
    public ObjectProcessConfiguration objectProcessConfiguration;

    public static PutAccessPointConfigForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointConfigForObjectProcessRequest self = new PutAccessPointConfigForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointConfigForObjectProcessRequest setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
        this.objectProcessConfiguration = objectProcessConfiguration;
        return this;
    }
    public ObjectProcessConfiguration getObjectProcessConfiguration() {
        return this.objectProcessConfiguration;
    }

}
