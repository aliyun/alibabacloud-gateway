// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DoMetaQueryResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about objects.</p>
     */
    @NameInMap("Files")
    public DoMetaQueryResponseBodyFiles files;

    /**
     * <p>The token that is used in the next request to retrieve a new page of results when the total number of objects exceeds the value of MaxResults. The value of NextToken is used to return the unreturned results in the next request.</p>
     * <br>
     * <p>A value is returned for this parameter only when not all objects are returned.</p>
     */
    @NameInMap("NextToken")
    public String nextToken;

    public static DoMetaQueryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DoMetaQueryResponseBody self = new DoMetaQueryResponseBody();
        return TeaModel.build(map, self);
    }

    public DoMetaQueryResponseBody setFiles(DoMetaQueryResponseBodyFiles files) {
        this.files = files;
        return this;
    }
    public DoMetaQueryResponseBodyFiles getFiles() {
        return this.files;
    }

    public DoMetaQueryResponseBody setNextToken(String nextToken) {
        this.nextToken = nextToken;
        return this;
    }
    public String getNextToken() {
        return this.nextToken;
    }

    public static class DoMetaQueryResponseBodyFiles extends TeaModel {
        /**
         * <p>The container that stores the information about a single object.</p>
         */
        @NameInMap("File")
        public java.util.List<MetaQueryFile> file;

        public static DoMetaQueryResponseBodyFiles build(java.util.Map<String, ?> map) throws Exception {
            DoMetaQueryResponseBodyFiles self = new DoMetaQueryResponseBodyFiles();
            return TeaModel.build(map, self);
        }

        public DoMetaQueryResponseBodyFiles setFile(java.util.List<MetaQueryFile> file) {
            this.file = file;
            return this;
        }
        public java.util.List<MetaQueryFile> getFile() {
            return this.file;
        }

    }

}
