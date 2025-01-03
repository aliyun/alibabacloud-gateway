// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RoutingRule extends TeaModel {
    @NameInMap("Condition")
    public RoutingRuleCondition condition;

    @NameInMap("LuaConfig")
    public RoutingRuleLuaConfig luaConfig;

    @NameInMap("Redirect")
    public RoutingRuleRedirect redirect;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("RuleNumber")
    public Long ruleNumber;

    public static RoutingRule build(java.util.Map<String, ?> map) throws Exception {
        RoutingRule self = new RoutingRule();
        return TeaModel.build(map, self);
    }

    public RoutingRule setCondition(RoutingRuleCondition condition) {
        this.condition = condition;
        return this;
    }
    public RoutingRuleCondition getCondition() {
        return this.condition;
    }

    public RoutingRule setLuaConfig(RoutingRuleLuaConfig luaConfig) {
        this.luaConfig = luaConfig;
        return this;
    }
    public RoutingRuleLuaConfig getLuaConfig() {
        return this.luaConfig;
    }

    public RoutingRule setRedirect(RoutingRuleRedirect redirect) {
        this.redirect = redirect;
        return this;
    }
    public RoutingRuleRedirect getRedirect() {
        return this.redirect;
    }

    public RoutingRule setRuleNumber(Long ruleNumber) {
        this.ruleNumber = ruleNumber;
        return this;
    }
    public Long getRuleNumber() {
        return this.ruleNumber;
    }

}
