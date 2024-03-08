// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class WebsiteConfiguration extends TeaModel {
    @NameInMap("ErrorDocument")
    public ErrorDocument errorDocument;

    @NameInMap("IndexDocument")
    public IndexDocument indexDocument;

    @NameInMap("RoutingRules")
    public WebsiteConfigurationRoutingRules routingRules;

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

    public WebsiteConfiguration setRoutingRules(WebsiteConfigurationRoutingRules routingRules) {
        this.routingRules = routingRules;
        return this;
    }
    public WebsiteConfigurationRoutingRules getRoutingRules() {
        return this.routingRules;
    }

    public static class WebsiteConfigurationRoutingRules extends TeaModel {
        @NameInMap("RoutingRule")
        public java.util.List<RoutingRule> routingRule;

        public static WebsiteConfigurationRoutingRules build(java.util.Map<String, ?> map) throws Exception {
            WebsiteConfigurationRoutingRules self = new WebsiteConfigurationRoutingRules();
            return TeaModel.build(map, self);
        }

        public WebsiteConfigurationRoutingRules setRoutingRule(java.util.List<RoutingRule> routingRule) {
            this.routingRule = routingRule;
            return this;
        }
        public java.util.List<RoutingRule> getRoutingRule() {
            return this.routingRule;
        }

    }

}
