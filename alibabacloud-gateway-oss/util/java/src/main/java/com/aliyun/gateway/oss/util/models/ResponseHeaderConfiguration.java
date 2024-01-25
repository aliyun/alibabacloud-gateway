// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ResponseHeaderConfiguration extends TeaModel {
    @NameInMap("Rule")
    public java.util.List<ResponseHeaderConfigurationRule> rule;

    public static ResponseHeaderConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ResponseHeaderConfiguration self = new ResponseHeaderConfiguration();
        return TeaModel.build(map, self);
    }

    public ResponseHeaderConfiguration setRule(java.util.List<ResponseHeaderConfigurationRule> rule) {
        this.rule = rule;
        return this;
    }
    public java.util.List<ResponseHeaderConfigurationRule> getRule() {
        return this.rule;
    }

    public static class ResponseHeaderConfigurationRuleFilters extends TeaModel {
        @NameInMap("Operation")
        public java.util.List<String> operation;

        public static ResponseHeaderConfigurationRuleFilters build(java.util.Map<String, ?> map) throws Exception {
            ResponseHeaderConfigurationRuleFilters self = new ResponseHeaderConfigurationRuleFilters();
            return TeaModel.build(map, self);
        }

        public ResponseHeaderConfigurationRuleFilters setOperation(java.util.List<String> operation) {
            this.operation = operation;
            return this;
        }
        public java.util.List<String> getOperation() {
            return this.operation;
        }

    }

    public static class ResponseHeaderConfigurationRuleHideHeaders extends TeaModel {
        @NameInMap("Header")
        public java.util.List<String> header;

        public static ResponseHeaderConfigurationRuleHideHeaders build(java.util.Map<String, ?> map) throws Exception {
            ResponseHeaderConfigurationRuleHideHeaders self = new ResponseHeaderConfigurationRuleHideHeaders();
            return TeaModel.build(map, self);
        }

        public ResponseHeaderConfigurationRuleHideHeaders setHeader(java.util.List<String> header) {
            this.header = header;
            return this;
        }
        public java.util.List<String> getHeader() {
            return this.header;
        }

    }

    public static class ResponseHeaderConfigurationRule extends TeaModel {
        @NameInMap("Filters")
        public ResponseHeaderConfigurationRuleFilters filters;

        @NameInMap("HideHeaders")
        public ResponseHeaderConfigurationRuleHideHeaders hideHeaders;

        @NameInMap("Name")
        public String name;

        public static ResponseHeaderConfigurationRule build(java.util.Map<String, ?> map) throws Exception {
            ResponseHeaderConfigurationRule self = new ResponseHeaderConfigurationRule();
            return TeaModel.build(map, self);
        }

        public ResponseHeaderConfigurationRule setFilters(ResponseHeaderConfigurationRuleFilters filters) {
            this.filters = filters;
            return this;
        }
        public ResponseHeaderConfigurationRuleFilters getFilters() {
            return this.filters;
        }

        public ResponseHeaderConfigurationRule setHideHeaders(ResponseHeaderConfigurationRuleHideHeaders hideHeaders) {
            this.hideHeaders = hideHeaders;
            return this;
        }
        public ResponseHeaderConfigurationRuleHideHeaders getHideHeaders() {
            return this.hideHeaders;
        }

        public ResponseHeaderConfigurationRule setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

    }

}
