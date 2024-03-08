// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetAccessPointConfigForObjectProcessResponseBody extends TeaModel {
    @NameInMap("ObjectProcessConfiguration")
    public ObjectProcessConfiguration objectProcessConfiguration;

    public static GetAccessPointConfigForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointConfigForObjectProcessResponseBody self = new GetAccessPointConfigForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public GetAccessPointConfigForObjectProcessResponseBody setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
        this.objectProcessConfiguration = objectProcessConfiguration;
        return this;
    }
    public ObjectProcessConfiguration getObjectProcessConfiguration() {
        return this.objectProcessConfiguration;
    }

}
