// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class DoMetaQueryResponseBody extends TeaModel {
    /**
     * <p>The container for the query result.</p>
     */
    @NameInMap("MetaQuery")
    public MetaQuery metaQuery;

    public static DoMetaQueryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DoMetaQueryResponseBody self = new DoMetaQueryResponseBody();
        return TeaModel.build(map, self);
    }

    public DoMetaQueryResponseBody setMetaQuery(MetaQuery metaQuery) {
        this.metaQuery = metaQuery;
        return this;
    }
    public MetaQuery getMetaQuery() {
        return this.metaQuery;
    }

    public static class Files extends TeaModel {
        /**
         * <p>The list of file information.</p>
         */
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

    public static class MetaQuery extends TeaModel {
        /**
         * <p>The container for the information about objects.</p>
         */
        @NameInMap("Files")
        public Files files;

        /**
         * <p>The token that is used for the next query when the total number of objects exceeds the value of MaxResults.</p>
         * <p>The value of NextToken is used to return the unreturned results in the next query.</p>
         * <br>
         * <p>This parameter has a value only when not all objects are returned.</p>
         */
        @NameInMap("NextToken")
        public String nextToken;

        public static MetaQuery build(java.util.Map<String, ?> map) throws Exception {
            MetaQuery self = new MetaQuery();
            return TeaModel.build(map, self);
        }

        public MetaQuery setFiles(Files files) {
            this.files = files;
            return this;
        }
        public Files getFiles() {
            return this.files;
        }

        public MetaQuery setNextToken(String nextToken) {
            this.nextToken = nextToken;
            return this;
        }
        public String getNextToken() {
            return this.nextToken;
        }

    }

}
