// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InitiateMultipartUploadResponseBody extends TeaModel {
    /**
     * <p>The container that stores the results of the InitiateMultipartUpload request.</p>
     */
    @NameInMap("InitiateMultipartUploadResult")
    public InitiateMultipartUploadResult initiateMultipartUploadResult;

    public static InitiateMultipartUploadResponseBody build(java.util.Map<String, ?> map) throws Exception {
        InitiateMultipartUploadResponseBody self = new InitiateMultipartUploadResponseBody();
        return TeaModel.build(map, self);
    }

    public InitiateMultipartUploadResponseBody setInitiateMultipartUploadResult(InitiateMultipartUploadResult initiateMultipartUploadResult) {
        this.initiateMultipartUploadResult = initiateMultipartUploadResult;
        return this;
    }
    public InitiateMultipartUploadResult getInitiateMultipartUploadResult() {
        return this.initiateMultipartUploadResult;
    }

    public static class InitiateMultipartUploadResult extends TeaModel {
        /**
         * <p>The name of the bucket to which the object is uploaded by the multipart upload task.</p>
         * 
         * <strong>example:</strong>
         * <p>example-bucket</p>
         */
        @NameInMap("Bucket")
        public String bucket;

        /**
         * <p>The encoding type of the object name in the response. If the encoding-type parameter is specified in the request, the object name in the response is encoded.</p>
         * 
         * <strong>example:</strong>
         * <p>url</p>
         */
        @NameInMap("EncodingType")
        public String encodingType;

        /**
         * <p>The name of the object that is uploaded by the multipart upload task.</p>
         * 
         * <strong>example:</strong>
         * <p>test.txt</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <p>The upload ID that uniquely identifies the multipart upload task. The upload ID is used to call UploadPart and CompleteMultipartUpload later.</p>
         * 
         * <strong>example:</strong>
         * <p>0004B9894A22E5B1888A1E29F823****</p>
         */
        @NameInMap("UploadId")
        public String uploadId;

        public static InitiateMultipartUploadResult build(java.util.Map<String, ?> map) throws Exception {
            InitiateMultipartUploadResult self = new InitiateMultipartUploadResult();
            return TeaModel.build(map, self);
        }

        public InitiateMultipartUploadResult setBucket(String bucket) {
            this.bucket = bucket;
            return this;
        }
        public String getBucket() {
            return this.bucket;
        }

        public InitiateMultipartUploadResult setEncodingType(String encodingType) {
            this.encodingType = encodingType;
            return this;
        }
        public String getEncodingType() {
            return this.encodingType;
        }

        public InitiateMultipartUploadResult setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public InitiateMultipartUploadResult setUploadId(String uploadId) {
            this.uploadId = uploadId;
            return this;
        }
        public String getUploadId() {
            return this.uploadId;
        }

    }

}
