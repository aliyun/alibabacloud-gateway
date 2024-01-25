// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class MetaQuery extends TeaModel {
    @NameInMap("Aggregations")
    public MetaQueryAggregations aggregations;

    @NameInMap("MaxResults")
    public Long maxResults;

    @NameInMap("NextToken")
    public String nextToken;

    @NameInMap("Order")
    public String order;

    @NameInMap("Query")
    public String query;

    @NameInMap("Sort")
    public String sort;

    public static MetaQuery build(java.util.Map<String, ?> map) throws Exception {
        MetaQuery self = new MetaQuery();
        return TeaModel.build(map, self);
    }

    public MetaQuery setAggregations(MetaQueryAggregations aggregations) {
        this.aggregations = aggregations;
        return this;
    }
    public MetaQueryAggregations getAggregations() {
        return this.aggregations;
    }

    public MetaQuery setMaxResults(Long maxResults) {
        this.maxResults = maxResults;
        return this;
    }
    public Long getMaxResults() {
        return this.maxResults;
    }

    public MetaQuery setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

    public MetaQuery setOrder(String order) {
        this.order = order;
        return this;
    }
    public String getOrder() {
        return this.order;
    }

    public MetaQuery setQuery(String query) {
        this.query = query;
        return this;
    }
    public String getQuery() {
        return this.query;
    }

    public MetaQuery setSort(String sort) {
        this.sort = sort;
        return this;
    }
    public String getSort() {
        return this.sort;
    }

    public static class MetaQueryAggregations extends TeaModel {
        @NameInMap("Aggregation")
        public java.util.List<MetaQueryAggregation> aggregation;

        public static MetaQueryAggregations build(java.util.Map<String, ?> map) throws Exception {
            MetaQueryAggregations self = new MetaQueryAggregations();
            return TeaModel.build(map, self);
        }

        public MetaQueryAggregations setAggregation(java.util.List<MetaQueryAggregation> aggregation) {
            this.aggregation = aggregation;
            return this;
        }
        public java.util.List<MetaQueryAggregation> getAggregation() {
            return this.aggregation;
        }

    }

}
