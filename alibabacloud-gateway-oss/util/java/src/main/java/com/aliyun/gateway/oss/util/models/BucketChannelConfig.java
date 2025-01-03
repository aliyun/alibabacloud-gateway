// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BucketChannelConfig extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>testinfo</p>
     */
    @NameInMap("DebugInfo")
    public String debugInfo;

    @NameInMap("RuleList")
    public RuleList ruleList;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("version")
    public Integer version;

    public static BucketChannelConfig build(java.util.Map<String, ?> map) throws Exception {
        BucketChannelConfig self = new BucketChannelConfig();
        return TeaModel.build(map, self);
    }

    public BucketChannelConfig setDebugInfo(String debugInfo) {
        this.debugInfo = debugInfo;
        return this;
    }
    public String getDebugInfo() {
        return this.debugInfo;
    }

    public BucketChannelConfig setRuleList(RuleList ruleList) {
        this.ruleList = ruleList;
        return this;
    }
    public RuleList getRuleList() {
        return this.ruleList;
    }

    public BucketChannelConfig setVersion(Integer version) {
        this.version = version;
        return this;
    }
    public Integer getVersion() {
        return this.version;
    }

    public static class Rule extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>a</p>
         */
        @NameInMap("FrontContent")
        public String frontContent;

        /**
         * <strong>example:</strong>
         * <p>rule1</p>
         */
        @NameInMap("RuleName")
        public String ruleName;

        /**
         * <strong>example:</strong>
         * <p>\.webp$</p>
         */
        @NameInMap("RuleRegex")
        public String ruleRegex;

        public static Rule build(java.util.Map<String, ?> map) throws Exception {
            Rule self = new Rule();
            return TeaModel.build(map, self);
        }

        public Rule setFrontContent(String frontContent) {
            this.frontContent = frontContent;
            return this;
        }
        public String getFrontContent() {
            return this.frontContent;
        }

        public Rule setRuleName(String ruleName) {
            this.ruleName = ruleName;
            return this;
        }
        public String getRuleName() {
            return this.ruleName;
        }

        public Rule setRuleRegex(String ruleRegex) {
            this.ruleRegex = ruleRegex;
            return this;
        }
        public String getRuleRegex() {
            return this.ruleRegex;
        }

    }

    public static class RuleList extends TeaModel {
        @NameInMap("Rule")
        public java.util.List<Rule> rule;

        public static RuleList build(java.util.Map<String, ?> map) throws Exception {
            RuleList self = new RuleList();
            return TeaModel.build(map, self);
        }

        public RuleList setRule(java.util.List<Rule> rule) {
            this.rule = rule;
            return this;
        }
        public java.util.List<Rule> getRule() {
            return this.rule;
        }

    }

}
