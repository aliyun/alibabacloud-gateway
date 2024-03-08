// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ReplicationPrefixSet extends TeaModel {
    @NameInMap("Prefix")
    public java.util.List<String> prefix;

    public static ReplicationPrefixSet build(java.util.Map<String, ?> map) throws Exception {
        ReplicationPrefixSet self = new ReplicationPrefixSet();
        return TeaModel.build(map, self);
    }

    public ReplicationPrefixSet setPrefix(java.util.List<String> prefix) {
        this.prefix = prefix;
        return this;
    }
    public java.util.List<String> getPrefix() {
        return this.prefix;
    }

}
