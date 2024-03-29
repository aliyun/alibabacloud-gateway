// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class WebsiteConfiguration extends TeaModel {
    @NameInMap("ErrorDocument")
    public ErrorDocument errorDocument;

    @NameInMap("IndexDocument")
    public IndexDocument indexDocument;

    @NameInMap("RoutingRules")
    public RoutingRules routingRules;

    public static WebsiteConfiguration build(java.util.Map<String, ?> map) throws Exception {
        WebsiteConfiguration self = new WebsiteConfiguration();
        return TeaModel.build(map, self);
    }

    public WebsiteConfiguration setErrorDocument(ErrorDocument errorDocument) {
        this.errorDocument = errorDocument;
        return this;
    }
    public ErrorDocument getErrorDocument() {
        return this.errorDocument;
    }

    public WebsiteConfiguration setIndexDocument(IndexDocument indexDocument) {
        this.indexDocument = indexDocument;
        return this;
    }
    public IndexDocument getIndexDocument() {
        return this.indexDocument;
    }

    public WebsiteConfiguration setRoutingRules(RoutingRules routingRules) {
        this.routingRules = routingRules;
        return this;
    }
    public RoutingRules getRoutingRules() {
        return this.routingRules;
    }

    public static class RoutingRules extends TeaModel {
        @NameInMap("RoutingRule")
        public java.util.List<RoutingRule> routingRule;

        public static RoutingRules build(java.util.Map<String, ?> map) throws Exception {
            RoutingRules self = new RoutingRules();
            return TeaModel.build(map, self);
        }

        public RoutingRules setRoutingRule(java.util.List<RoutingRule> routingRule) {
            this.routingRule = routingRule;
            return this;
        }
        public java.util.List<RoutingRule> getRoutingRule() {
            return this.routingRule;
        }

    }

}
