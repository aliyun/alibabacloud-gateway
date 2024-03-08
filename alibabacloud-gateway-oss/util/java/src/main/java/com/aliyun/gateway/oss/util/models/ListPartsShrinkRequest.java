// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListPartsShrinkRequest extends TeaModel {
    /**
     * <p>The encoding type of the object name in the response. The object name can contain any characters encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse certain control characters, such as characters with an ASCII value from 0 to 10. You can set the Encoding-type parameter to encode the returned object name.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("encoding-type")
    public String encodingTypeShrink;

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

    public static ListPartsShrinkRequest build(java.util.Map<String, ?> map) throws Exception {
        ListPartsShrinkRequest self = new ListPartsShrinkRequest();
        return TeaModel.build(map, self);
    }

    public ListPartsShrinkRequest setEncodingTypeShrink(String encodingTypeShrink) {
        this.encodingTypeShrink = encodingTypeShrink;
        return this;
    }
    public String getEncodingTypeShrink() {
        return this.encodingTypeShrink;
    }

    public ListPartsShrinkRequest setMaxParts(Long maxParts) {
        this.maxParts = maxParts;
        return this;
    }
    public Long getMaxParts() {
        return this.maxParts;
    }

    public ListPartsShrinkRequest setPartNumberMarker(Long partNumberMarker) {
        this.partNumberMarker = partNumberMarker;
        return this;
    }
    public Long getPartNumberMarker() {
        return this.partNumberMarker;
    }

    public ListPartsShrinkRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
