// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class InventoryFilter extends TeaModel {
    @NameInMap("Prefix")
    public String prefix;

    public static InventoryFilter build(java.util.Map<String, ?> map) throws Exception {
        InventoryFilter self = new InventoryFilter();
        return TeaModel.build(map, self);
    }

    public InventoryFilter setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

}
