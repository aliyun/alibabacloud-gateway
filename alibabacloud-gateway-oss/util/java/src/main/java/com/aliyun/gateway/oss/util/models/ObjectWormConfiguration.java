// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class ObjectWormConfiguration extends TeaModel {
    @NameInMap("ObjectWormEnabled")
    public String objectWormEnabled;

    @NameInMap("Rule")
    public Rule rule;

    public static ObjectWormConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ObjectWormConfiguration self = new ObjectWormConfiguration();
        return TeaModel.build(map, self);
    }

    public ObjectWormConfiguration setObjectWormEnabled(String objectWormEnabled) {
        this.objectWormEnabled = objectWormEnabled;
        return this;
    }
    public String getObjectWormEnabled() {
        return this.objectWormEnabled;
    }

    public ObjectWormConfiguration setRule(Rule rule) {
        this.rule = rule;
        return this;
    }
    public Rule getRule() {
        return this.rule;
    }

    public static class DefaultRetention extends TeaModel {
        @NameInMap("Days")
        public Long days;

        @NameInMap("Mode")
        public String mode;

        @NameInMap("Years")
        public Long years;

        public static DefaultRetention build(java.util.Map<String, ?> map) throws Exception {
            DefaultRetention self = new DefaultRetention();
            return TeaModel.build(map, self);
        }

        public DefaultRetention setDays(Long days) {
            this.days = days;
            return this;
        }
        public Long getDays() {
            return this.days;
        }

        public DefaultRetention setMode(String mode) {
            this.mode = mode;
            return this;
        }
        public String getMode() {
            return this.mode;
        }

        public DefaultRetention setYears(Long years) {
            this.years = years;
            return this;
        }
        public Long getYears() {
            return this.years;
        }

    }

    public static class Rule extends TeaModel {
        @NameInMap("DefaultRetention")
        public java.util.List<DefaultRetention> defaultRetention;

        public static Rule build(java.util.Map<String, ?> map) throws Exception {
            Rule self = new Rule();
            return TeaModel.build(map, self);
        }

        public Rule setDefaultRetention(java.util.List<DefaultRetention> defaultRetention) {
            this.defaultRetention = defaultRetention;
            return this;
        }
        public java.util.List<DefaultRetention> getDefaultRetention() {
            return this.defaultRetention;
        }

    }

}
