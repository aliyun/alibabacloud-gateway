// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryResp extends TeaModel {
    @NameInMap("Aggregations")
    public Aggregations aggregations;

    @NameInMap("Files")
    public Files files;

    /**
     * <strong>example:</strong>
     * <p>MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****</p>
     */
    @NameInMap("NextToken")
    public String nextToken;

    public static MetaQueryResp build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryResp self = new MetaQueryResp();
        return TeaModel.build(map, self);
    }

    public MetaQueryResp setAggregations(Aggregations aggregations) {
        this.aggregations = aggregations;
        return this;
    }
    public Aggregations getAggregations() {
        return this.aggregations;
    }

    public MetaQueryResp setFiles(Files files) {
        this.files = files;
        return this;
    }
    public Files getFiles() {
        return this.files;
    }

    public MetaQueryResp setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

    public static class Aggregations extends TeaModel {
        @NameInMap("Aggregation")
        public java.util.List<MetaQueryAggregationsResult> aggregation;

        public static Aggregations build(java.util.Map<String, ?> map) throws Exception {
            Aggregations self = new Aggregations();
            return TeaModel.build(map, self);
        }

        public Aggregations setAggregation(java.util.List<MetaQueryAggregationsResult> aggregation) {
            this.aggregation = aggregation;
            return this;
        }
        public java.util.List<MetaQueryAggregationsResult> getAggregation() {
            return this.aggregation;
        }

    }

    public static class Files extends TeaModel {
        @NameInMap("File")
        public java.util.List<MetaQueryFile> file;

        public static Files build(java.util.Map<String, ?> map) throws Exception {
            Files self = new Files();
            return TeaModel.build(map, self);
        }

        public Files setFile(java.util.List<MetaQueryFile> file) {
            this.file = file;
            return this;
        }
        public java.util.List<MetaQueryFile> getFile() {
            return this.file;
        }

    }

}
