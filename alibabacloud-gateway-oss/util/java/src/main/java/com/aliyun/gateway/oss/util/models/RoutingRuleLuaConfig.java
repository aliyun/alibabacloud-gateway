// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RoutingRuleLuaConfig extends TeaModel {
    @NameInMap("Script")
    public String script;

    public static RoutingRuleLuaConfig build(java.util.Map<String, ?> map) throws Exception {
        RoutingRuleLuaConfig self = new RoutingRuleLuaConfig();
        return TeaModel.build(map, self);
    }

    public RoutingRuleLuaConfig setScript(String script) {
        this.script = script;
        return this;
    }
    public String getScript() {
        return this.script;
    }

}
