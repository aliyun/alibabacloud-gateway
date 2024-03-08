// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListPartsRequest extends TeaModel {
    /**
     * <p>The encoding type of the object name in the response. The object name can contain any characters encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse certain control characters, such as characters with an ASCII value from 0 to 10. You can set the Encoding-type parameter to encode the returned object name.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>The maximum number of parts that can be returned by OSS. </p>
     * <br>
     * <p>Default value: 1000.</p>
     * <br>
     * <p>Maximum value: 1000.</p>
     */
    @NameInMap("max-parts")
    public Long maxParts;

    /**
     * <p>The number of the part after which the list starts. All parts whose part numbers are greater than this parameter value are listed. </p>
     * <p>Default value: null.</p>
     */
    @NameInMap("part-number-marker")
    public Long partNumberMarker;

    /**
     * <p>The ID of the multipart upload task.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static ListPartsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListPartsRequest self = new ListPartsRequest();
        return TeaModel.build(map, self);
    }

    public ListPartsRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListPartsRequest setMaxParts(Long maxParts) {
        this.maxParts = maxParts;
        return this;
    }
    public Long getMaxParts() {
        return this.maxParts;
    }

    public ListPartsRequest setPartNumberMarker(Long partNumberMarker) {
        this.partNumberMarker = partNumberMarker;
        return this;
    }
    public Long getPartNumberMarker() {
        return this.partNumberMarker;
    }

    public ListPartsRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
