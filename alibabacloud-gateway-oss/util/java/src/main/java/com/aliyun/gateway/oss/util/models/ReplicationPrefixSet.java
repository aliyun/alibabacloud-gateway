// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ReplicationPrefixSet extends TeaModel {
    @NameInMap("Prefix")
    public java.util.List<String> prefixs;

    public static ReplicationPrefixSet build(java.util.Map<String, ?> map) throws Exception {
        ReplicationPrefixSet self = new ReplicationPrefixSet();
        return TeaModel.build(map, self);
    }

    public ReplicationPrefixSet setPrefixs(java.util.List<String> prefixs) {
        this.prefixs = prefixs;
        return this;
    }
    public java.util.List<String> getPrefixs() {
        return this.prefixs;
    }

}
