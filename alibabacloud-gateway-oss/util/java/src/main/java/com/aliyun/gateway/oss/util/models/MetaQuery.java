// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQuery extends TeaModel {
    @NameInMap("Aggregations")
    public Aggregations aggregations;

    /**
     * <strong>example:</strong>
     * <p>100</p>
     */
    @NameInMap("MaxResults")
    public Long maxResults;

    @NameInMap("MediaTypes")
    public MediaTypes mediaTypes;

    /**
     * <strong>example:</strong>
     * <p>MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****</p>
     */
    @NameInMap("NextToken")
    public String nextToken;

    @NameInMap("Order")
    public String order;

    /**
     * <strong>example:</strong>
     * <p>{&quot;Field&quot;: &quot;Size&quot;,&quot;Value&quot;: &quot;1048576&quot;,&quot;Operation&quot;: &quot;gt&quot;}</p>
     */
    @NameInMap("Query")
    public String query;

    /**
     * <strong>example:</strong>
     * <p>{&quot;Operation&quot;:&quot;gt&quot;, &quot;Field&quot;: &quot;Size&quot;, &quot;Value&quot;: &quot;30&quot;}</p>
     */
    @NameInMap("SimpleQuery")
    public String simpleQuery;

    /**
     * <strong>example:</strong>
     * <p>Size</p>
     */
    @NameInMap("Sort")
    public String sort;

    public static MetaQuery build(java.util.Map<String, ?> map) throws Exception {
        MetaQuery self = new MetaQuery();
        return TeaModel.build(map, self);
    }

    public MetaQuery setAggregations(Aggregations aggregations) {
        this.aggregations = aggregations;
        return this;
    }
    public Aggregations getAggregations() {
        return this.aggregations;
    }

    public MetaQuery setMaxResults(Long maxResults) {
        this.maxResults = maxResults;
        return this;
    }
    public Long getMaxResults() {
        return this.maxResults;
    }

    public MetaQuery setMediaTypes(MediaTypes mediaTypes) {
        this.mediaTypes = mediaTypes;
        return this;
    }
    public MediaTypes getMediaTypes() {
        return this.mediaTypes;
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

    public MetaQuery setSimpleQuery(String simpleQuery) {
        this.simpleQuery = simpleQuery;
        return this;
    }
    public String getSimpleQuery() {
        return this.simpleQuery;
    }

    public MetaQuery setSort(String sort) {
        this.sort = sort;
        return this;
    }
    public String getSort() {
        return this.sort;
    }

    public static class Aggregations extends TeaModel {
        @NameInMap("Aggregation")
        public java.util.List<MetaQueryAggregation> aggregation;

        public static Aggregations build(java.util.Map<String, ?> map) throws Exception {
            Aggregations self = new Aggregations();
            return TeaModel.build(map, self);
        }

        public Aggregations setAggregation(java.util.List<MetaQueryAggregation> aggregation) {
            this.aggregation = aggregation;
            return this;
        }
        public java.util.List<MetaQueryAggregation> getAggregation() {
            return this.aggregation;
        }

    }

    public static class MediaTypes extends TeaModel {
        @NameInMap("MediaType")
        public java.util.List<String> mediaType;

        public static MediaTypes build(java.util.Map<String, ?> map) throws Exception {
            MediaTypes self = new MediaTypes();
            return TeaModel.build(map, self);
        }

        public MediaTypes setMediaType(java.util.List<String> mediaType) {
            this.mediaType = mediaType;
            return this;
        }
        public java.util.List<String> getMediaType() {
            return this.mediaType;
        }

    }

}
