// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateAccessPointForObjectProcessRequest extends TeaModel {
    @NameInMap("AccessPointName")
    public String accessPointName;

    @NameInMap("ObjectProcessConfiguration")
    public ObjectProcessConfiguration objectProcessConfiguration;

    public static CreateAccessPointForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointForObjectProcessRequest self = new CreateAccessPointForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointForObjectProcessRequest setAccessPointName(String accessPointName) {
        this.accessPointName = accessPointName;
        return this;
    }
    public String getAccessPointName() {
        return this.accessPointName;
    }

    public CreateAccessPointForObjectProcessRequest setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
        this.objectProcessConfiguration = objectProcessConfiguration;
        return this;
    }
    public ObjectProcessConfiguration getObjectProcessConfiguration() {
        return this.objectProcessConfiguration;
    }

}
