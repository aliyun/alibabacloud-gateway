// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RoutingRuleCondition extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>404</p>
     */
    @NameInMap("HttpErrorCodeReturnedEquals")
    public Long httpErrorCodeReturnedEquals;

    @NameInMap("IncludeHeader")
    public java.util.List<IncludeHeader> includeHeader;

    /**
     * <strong>example:</strong>
     * <p>abc/</p>
     */
    @NameInMap("KeyPrefixEquals")
    public String keyPrefixEquals;

    /**
     * <strong>example:</strong>
     * <p>.txt</p>
     */
    @NameInMap("KeySuffixEquals")
    public String keySuffixEquals;

    public static RoutingRuleCondition build(java.util.Map<String, ?> map) throws Exception {
        RoutingRuleCondition self = new RoutingRuleCondition();
        return TeaModel.build(map, self);
    }

    public RoutingRuleCondition setHttpErrorCodeReturnedEquals(Long httpErrorCodeReturnedEquals) {
        this.httpErrorCodeReturnedEquals = httpErrorCodeReturnedEquals;
        return this;
    }
    public Long getHttpErrorCodeReturnedEquals() {
        return this.httpErrorCodeReturnedEquals;
    }

    public RoutingRuleCondition setIncludeHeader(java.util.List<IncludeHeader> includeHeader) {
        this.includeHeader = includeHeader;
        return this;
    }
    public java.util.List<IncludeHeader> getIncludeHeader() {
        return this.includeHeader;
    }

    public RoutingRuleCondition setKeyPrefixEquals(String keyPrefixEquals) {
        this.keyPrefixEquals = keyPrefixEquals;
        return this;
    }
    public String getKeyPrefixEquals() {
        return this.keyPrefixEquals;
    }

    public RoutingRuleCondition setKeySuffixEquals(String keySuffixEquals) {
        this.keySuffixEquals = keySuffixEquals;
        return this;
    }
    public String getKeySuffixEquals() {
        return this.keySuffixEquals;
    }

    public static class IncludeHeader extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>-test-suffix</p>
         */
        @NameInMap("EndsWith")
        public String endsWith;

        /**
         * <strong>example:</strong>
         * <p>test-value</p>
         */
        @NameInMap("Equals")
        public String equals;

        /**
         * <strong>example:</strong>
         * <p>test-header</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <strong>example:</strong>
         * <p>test-prefix-</p>
         */
        @NameInMap("StartsWith")
        public String startsWith;

        public static IncludeHeader build(java.util.Map<String, ?> map) throws Exception {
            IncludeHeader self = new IncludeHeader();
            return TeaModel.build(map, self);
        }

        public IncludeHeader setEndsWith(String endsWith) {
            this.endsWith = endsWith;
            return this;
        }
        public String getEndsWith() {
            return this.endsWith;
        }

        public IncludeHeader setEquals(String equals) {
            this.equals = equals;
            return this;
        }
        public String getEquals() {
            return this.equals;
        }

        public IncludeHeader setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public IncludeHeader setStartsWith(String startsWith) {
            this.startsWith = startsWith;
            return this;
        }
        public String getStartsWith() {
            return this.startsWith;
        }

    }

}
