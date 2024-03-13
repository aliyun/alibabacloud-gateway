// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UploadPartCopyResponseBody extends TeaModel {
    /**
     * <p>The container that stores the copy result.</p>
     */
    @NameInMap("CopyPartResult")
    public CopyPartResult copyPartResult;

    public static UploadPartCopyResponseBody build(java.util.Map<String, ?> map) throws Exception {
        UploadPartCopyResponseBody self = new UploadPartCopyResponseBody();
        return TeaModel.build(map, self);
    }

    public UploadPartCopyResponseBody setCopyPartResult(CopyPartResult copyPartResult) {
        this.copyPartResult = copyPartResult;
        return this;
    }
    public CopyPartResult getCopyPartResult() {
        return this.copyPartResult;
    }

    public static class CopyPartResult extends TeaModel {
        /**
         * <p>The ETag of the copied part.</p>
         */
        @NameInMap("ETag")
        public String ETag;

        /**
         * <p>The last modified time of copy source.</p>
         */
        @NameInMap("LastModified")
        public String lastModified;

        public static CopyPartResult build(java.util.Map<String, ?> map) throws Exception {
            CopyPartResult self = new CopyPartResult();
            return TeaModel.build(map, self);
        }

        public CopyPartResult setETag(String ETag) {
            this.ETag = ETag;
            return this;
        }
        public String getETag() {
            return this.ETag;
        }

        public CopyPartResult setLastModified(String lastModified) {
            this.lastModified = lastModified;
            return this;
        }
        public String getLastModified() {
            return this.lastModified;
        }

    }

}
