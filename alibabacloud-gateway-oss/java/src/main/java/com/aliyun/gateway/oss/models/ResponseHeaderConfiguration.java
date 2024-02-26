// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class ResponseHeaderConfiguration extends TeaModel {
    @NameInMap("Rule")
    public java.util.List<Rule> rule;

    public static ResponseHeaderConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ResponseHeaderConfiguration self = new ResponseHeaderConfiguration();
        return TeaModel.build(map, self);
    }

    public ResponseHeaderConfiguration setRule(java.util.List<Rule> rule) {
        this.rule = rule;
        return this;
    }
    public java.util.List<Rule> getRule() {
        return this.rule;
    }

    public static class Filters extends TeaModel {
        @NameInMap("Operation")
        public java.util.List<String> operation;

        public static Filters build(java.util.Map<String, ?> map) throws Exception {
            Filters self = new Filters();
            return TeaModel.build(map, self);
        }

        public Filters setOperation(java.util.List<String> operation) {
            this.operation = operation;
            return this;
        }
        public java.util.List<String> getOperation() {
            return this.operation;
        }

    }

    public static class HideHeaders extends TeaModel {
        @NameInMap("Header")
        public java.util.List<String> header;

        public static HideHeaders build(java.util.Map<String, ?> map) throws Exception {
            HideHeaders self = new HideHeaders();
            return TeaModel.build(map, self);
        }

        public HideHeaders setHeader(java.util.List<String> header) {
            this.header = header;
            return this;
        }
        public java.util.List<String> getHeader() {
            return this.header;
        }

    }

    public static class Rule extends TeaModel {
        @NameInMap("Filters")
        public Filters filters;

        @NameInMap("HideHeaders")
        public HideHeaders hideHeaders;

        @NameInMap("Name")
        public String name;

        public static Rule build(java.util.Map<String, ?> map) throws Exception {
            Rule self = new Rule();
            return TeaModel.build(map, self);
        }

        public Rule setFilters(Filters filters) {
            this.filters = filters;
            return this;
        }
        public Filters getFilters() {
            return this.filters;
        }

        public Rule setHideHeaders(HideHeaders hideHeaders) {
            this.hideHeaders = hideHeaders;
            return this;
        }
        public HideHeaders getHideHeaders() {
            return this.hideHeaders;
        }

        public Rule setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

    }

}
