// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class CopyObjectResponseBody extends TeaModel {
    /**
     * <p>The results of the CopyObject operation.</p>
     */
    @NameInMap("CopyObjectResult")
    public CopyObjectResult copyObjectResult;

    public static CopyObjectResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectResponseBody self = new CopyObjectResponseBody();
        return TeaModel.build(map, self);
    }

    public CopyObjectResponseBody setCopyObjectResult(CopyObjectResult copyObjectResult) {
        this.copyObjectResult = copyObjectResult;
        return this;
    }
    public CopyObjectResult getCopyObjectResult() {
        return this.copyObjectResult;
    }

    public static class CopyObjectResult extends TeaModel {
        /**
         * <p>The ETag value of the destination object.</p>
         */
        @NameInMap("ETag")
        public String ETag;

        /**
         * <p>The time when the destination object was last modified.</p>
         */
        @NameInMap("LastModified")
        public String lastModified;

        public static CopyObjectResult build(java.util.Map<String, ?> map) throws Exception {
            CopyObjectResult self = new CopyObjectResult();
            return TeaModel.build(map, self);
        }

        public CopyObjectResult setETag(String ETag) {
            this.ETag = ETag;
            return this;
        }
        public String getETag() {
            return this.ETag;
        }

        public CopyObjectResult setLastModified(String lastModified) {
            this.lastModified = lastModified;
            return this;
        }
        public String getLastModified() {
            return this.lastModified;
        }

    }

}
