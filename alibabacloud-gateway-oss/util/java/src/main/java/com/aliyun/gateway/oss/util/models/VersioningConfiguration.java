// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class VersioningConfiguration extends TeaModel {
    @NameInMap("Status")
    public String status;

    public static VersioningConfiguration build(java.util.Map<String, ?> map) throws Exception {
        VersioningConfiguration self = new VersioningConfiguration();
        return TeaModel.build(map, self);
    }

    public VersioningConfiguration setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

}
