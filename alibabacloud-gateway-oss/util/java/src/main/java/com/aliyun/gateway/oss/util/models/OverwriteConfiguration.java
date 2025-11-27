// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class OverwriteConfiguration extends TeaModel {
    @NameInMap("Rule")
    public java.util.List<Rule> rule;

    public static OverwriteConfiguration build(java.util.Map<String, ?> map) throws Exception {
        OverwriteConfiguration self = new OverwriteConfiguration();
        return TeaModel.build(map, self);
    }

    public OverwriteConfiguration setRule(java.util.List<Rule> rule) {
        this.rule = rule;
        return this;
    }
    public java.util.List<Rule> getRule() {
        return this.rule;
    }

    public static class Principals extends TeaModel {
        @NameInMap("Principal")
        public java.util.List<String> principal;

        public static Principals build(java.util.Map<String, ?> map) throws Exception {
            Principals self = new Principals();
            return TeaModel.build(map, self);
        }

        public Principals setPrincipal(java.util.List<String> principal) {
            this.principal = principal;
            return this;
        }
        public java.util.List<String> getPrincipal() {
            return this.principal;
        }

    }

    public static class Rule extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>forbid</p>
         */
        @NameInMap("Action")
        public String action;

        /**
         * <strong>example:</strong>
         * <p>forbid-write-rule1</p>
         */
        @NameInMap("ID")
        public String ID;

        /**
         * <strong>example:</strong>
         * <p>a/</p>
         */
        @NameInMap("Prefix")
        public String prefix;

        @NameInMap("Principals")
        public Principals principals;

        /**
         * <strong>example:</strong>
         * <p>.txt</p>
         */
        @NameInMap("Suffix")
        public String suffix;

        public static Rule build(java.util.Map<String, ?> map) throws Exception {
            Rule self = new Rule();
            return TeaModel.build(map, self);
        }

        public Rule setAction(String action) {
            this.action = action;
            return this;
        }
        public String getAction() {
            return this.action;
        }

        public Rule setID(String ID) {
            this.ID = ID;
            return this;
        }
        public String getID() {
            return this.ID;
        }

        public Rule setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

        public Rule setPrincipals(Principals principals) {
            this.principals = principals;
            return this;
        }
        public Principals getPrincipals() {
            return this.principals;
        }

        public Rule setSuffix(String suffix) {
            this.suffix = suffix;
            return this;
        }
        public String getSuffix() {
            return this.suffix;
        }

    }

}
