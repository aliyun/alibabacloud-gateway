// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class FileGroupInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>test-bucket</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <strong>example:</strong>
     * <p>&quot;06A4DDABDD5F65868B8C5919E76487D6&quot;</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>284</p>
     */
    @NameInMap("FileLength")
    public Long fileLength;

    /**
     * <p>FileGroup类型文件的信息</p>
     */
    @NameInMap("FilePart")
    public FilePart filePart;

    /**
     * <strong>example:</strong>
     * <p>file-goup-example</p>
     */
    @NameInMap("Key")
    public String key;

    public static FileGroupInfo build(java.util.Map<String, ?> map) throws Exception {
        FileGroupInfo self = new FileGroupInfo();
        return TeaModel.build(map, self);
    }

    public FileGroupInfo setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public FileGroupInfo setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public FileGroupInfo setFileLength(Long fileLength) {
        this.fileLength = fileLength;
        return this;
    }
    public Long getFileLength() {
        return this.fileLength;
    }

    public FileGroupInfo setFilePart(FilePart filePart) {
        this.filePart = filePart;
        return this;
    }
    public FilePart getFilePart() {
        return this.filePart;
    }

    public FileGroupInfo setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public static class Part extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>&quot;EB327B57BB20D17C293A966115FE27BD&quot;</p>
         */
        @NameInMap("ETag")
        public String ETag;

        /**
         * <strong>example:</strong>
         * <p>test.txt</p>
         */
        @NameInMap("PartName")
        public String partName;

        /**
         * <strong>example:</strong>
         * <p>3</p>
         */
        @NameInMap("PartNumber")
        public Long partNumber;

        /**
         * <strong>example:</strong>
         * <p>38</p>
         */
        @NameInMap("PartSize")
        public Long partSize;

        public static Part build(java.util.Map<String, ?> map) throws Exception {
            Part self = new Part();
            return TeaModel.build(map, self);
        }

        public Part setETag(String ETag) {
            this.ETag = ETag;
            return this;
        }
        public String getETag() {
            return this.ETag;
        }

        public Part setPartName(String partName) {
            this.partName = partName;
            return this;
        }
        public String getPartName() {
            return this.partName;
        }

        public Part setPartNumber(Long partNumber) {
            this.partNumber = partNumber;
            return this;
        }
        public Long getPartNumber() {
            return this.partNumber;
        }

        public Part setPartSize(Long partSize) {
            this.partSize = partSize;
            return this;
        }
        public Long getPartSize() {
            return this.partSize;
        }

    }

    public static class FilePart extends TeaModel {
        @NameInMap("Part")
        public java.util.List<Part> part;

        public static FilePart build(java.util.Map<String, ?> map) throws Exception {
            FilePart self = new FilePart();
            return TeaModel.build(map, self);
        }

        public FilePart setPart(java.util.List<Part> part) {
            this.part = part;
            return this;
        }
        public java.util.List<Part> getPart() {
            return this.part;
        }

    }

}
