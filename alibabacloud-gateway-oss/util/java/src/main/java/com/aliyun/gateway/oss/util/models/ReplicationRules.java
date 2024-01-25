// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ReplicationRules extends TeaModel {
    @NameInMap("ID")
    public java.util.List<String> ID;

    public static ReplicationRules build(java.util.Map<String, ?> map) throws Exception {
        ReplicationRules self = new ReplicationRules();
        return TeaModel.build(map, self);
    }

    public ReplicationRules setID(java.util.List<String> ID) {
        this.ID = ID;
        return this;
    }
    public java.util.List<String> getID() {
        return this.ID;
    }

}
