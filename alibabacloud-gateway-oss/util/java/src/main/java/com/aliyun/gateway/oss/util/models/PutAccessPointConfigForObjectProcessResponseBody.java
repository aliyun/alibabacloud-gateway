// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutAccessPointConfigForObjectProcessResponseBody extends TeaModel {
    @NameInMap("ObjectProcessConfiguration")
    public ObjectProcessConfiguration objectProcessConfiguration;

    public static PutAccessPointConfigForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointConfigForObjectProcessResponseBody self = new PutAccessPointConfigForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public PutAccessPointConfigForObjectProcessResponseBody setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
        this.objectProcessConfiguration = objectProcessConfiguration;
        return this;
    }
    public ObjectProcessConfiguration getObjectProcessConfiguration() {
        return this.objectProcessConfiguration;
    }

}
