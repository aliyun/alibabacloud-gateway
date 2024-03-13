// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CompleteMultipartUploadResponseBody extends TeaModel {
    /**
     * <p>The container that stores the results of the CompleteMultipartUpload request.</p>
     */
    @NameInMap("CompleteMultipartUploadResult")
    public CompleteMultipartUploadResult completeMultipartUploadResult;

    public static CompleteMultipartUploadResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUploadResponseBody self = new CompleteMultipartUploadResponseBody();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUploadResponseBody setCompleteMultipartUploadResult(CompleteMultipartUploadResult completeMultipartUploadResult) {
        this.completeMultipartUploadResult = completeMultipartUploadResult;
        return this;
    }
    public CompleteMultipartUploadResult getCompleteMultipartUploadResult() {
        return this.completeMultipartUploadResult;
    }

    public static class CompleteMultipartUploadResult extends TeaModel {
        /**
         * <p>The name of the bucket that contains the object you want to restore.</p>
         */
        @NameInMap("Bucket")
        public String bucket;

        /**
         * <p>The ETag that is generated when an object is created. ETags are used to identify the content of objects.</p>
         * <br>
         * <p>If an object is created by calling the CompleteMultipartUpload operation, the ETag value is not the MD5 hash of the object content but a unique value calculated based on a specific rule.</p>
         * <br>
         * <p>> The ETag of an object can be used to check whether the object content is modified. However, we recommend that you use the MD5 hash of an object rather than the ETag value of the object to verify data integrity.</p>
         */
        @NameInMap("ETag")
        public String ETag;

        /**
         * <p>The encoding type of the object name in the response. If this parameter is specified in the request, the object name is encoded in the response.</p>
         */
        @NameInMap("EncodingType")
        public String encodingType;

        /**
         * <p>The name of the uploaded object.</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <p>The URL that is used to access the uploaded object.</p>
         */
        @NameInMap("Location")
        public String location;

        public static CompleteMultipartUploadResult build(java.util.Map<String, ?> map) throws Exception {
            CompleteMultipartUploadResult self = new CompleteMultipartUploadResult();
            return TeaModel.build(map, self);
        }

        public CompleteMultipartUploadResult setBucket(String bucket) {
            this.bucket = bucket;
            return this;
        }
        public String getBucket() {
            return this.bucket;
        }

        public CompleteMultipartUploadResult setETag(String ETag) {
            this.ETag = ETag;
            return this;
        }
        public String getETag() {
            return this.ETag;
        }

        public CompleteMultipartUploadResult setEncodingType(String encodingType) {
            this.encodingType = encodingType;
            return this;
        }
        public String getEncodingType() {
            return this.encodingType;
        }

        public CompleteMultipartUploadResult setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public CompleteMultipartUploadResult setLocation(String location) {
            this.location = location;
            return this;
        }
        public String getLocation() {
            return this.location;
        }

    }

}
