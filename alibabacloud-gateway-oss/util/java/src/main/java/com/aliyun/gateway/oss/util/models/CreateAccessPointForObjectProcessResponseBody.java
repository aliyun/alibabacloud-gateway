// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CreateAccessPointForObjectProcessResponseBody extends TeaModel {
    @NameInMap("AccessPointForObjectProcessAlias")
    public String accessPointForObjectProcessAlias;

    @NameInMap("AccessPointForObjectProcessArn")
    public String accessPointForObjectProcessArn;

    public static CreateAccessPointForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointForObjectProcessResponseBody self = new CreateAccessPointForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointForObjectProcessResponseBody setAccessPointForObjectProcessAlias(String accessPointForObjectProcessAlias) {
        this.accessPointForObjectProcessAlias = accessPointForObjectProcessAlias;
        return this;
    }
    public String getAccessPointForObjectProcessAlias() {
        return this.accessPointForObjectProcessAlias;
    }

    public CreateAccessPointForObjectProcessResponseBody setAccessPointForObjectProcessArn(String accessPointForObjectProcessArn) {
        this.accessPointForObjectProcessArn = accessPointForObjectProcessArn;
        return this;
    }
    public String getAccessPointForObjectProcessArn() {
        return this.accessPointForObjectProcessArn;
    }

}
