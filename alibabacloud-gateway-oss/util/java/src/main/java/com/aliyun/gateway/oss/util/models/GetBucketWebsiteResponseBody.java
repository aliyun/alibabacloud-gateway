// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketWebsiteResponseBody extends TeaModel {
    /**
     * <p>The container that stores the error page.</p>
     */
    @NameInMap("ErrorDocument")
    public ErrorDocument errorDocument;

    /**
     * <p>The container that stores the default homepage.</p>
     */
    @NameInMap("IndexDocument")
    public IndexDocument indexDocument;

    /**
     * <p>The container that stores redirection rules or mirroring-based back-to-origin rules.</p>
     */
    @NameInMap("RoutingRules")
    public GetBucketWebsiteResponseBodyRoutingRules routingRules;

    public static GetBucketWebsiteResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketWebsiteResponseBody self = new GetBucketWebsiteResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketWebsiteResponseBody setErrorDocument(ErrorDocument errorDocument) {
        this.errorDocument = errorDocument;
        return this;
    }
    public ErrorDocument getErrorDocument() {
        return this.errorDocument;
    }

    public GetBucketWebsiteResponseBody setIndexDocument(IndexDocument indexDocument) {
        this.indexDocument = indexDocument;
        return this;
    }
    public IndexDocument getIndexDocument() {
        return this.indexDocument;
    }

    public GetBucketWebsiteResponseBody setRoutingRules(GetBucketWebsiteResponseBodyRoutingRules routingRules) {
        this.routingRules = routingRules;
        return this;
    }
    public GetBucketWebsiteResponseBodyRoutingRules getRoutingRules() {
        return this.routingRules;
    }

    public static class GetBucketWebsiteResponseBodyRoutingRules extends TeaModel {
        /**
         * <p>The redirection rules or mirroring-based back-to-origin rules.</p>
         */
        @NameInMap("RoutingRule")
        public java.util.List<RoutingRule> routingRule;

        public static GetBucketWebsiteResponseBodyRoutingRules build(java.util.Map<String, ?> map) throws Exception {
            GetBucketWebsiteResponseBodyRoutingRules self = new GetBucketWebsiteResponseBodyRoutingRules();
            return TeaModel.build(map, self);
        }

        public GetBucketWebsiteResponseBodyRoutingRules setRoutingRule(java.util.List<RoutingRule> routingRule) {
            this.routingRule = routingRule;
            return this;
        }
        public java.util.List<RoutingRule> getRoutingRule() {
            return this.routingRule;
        }

    }

}
