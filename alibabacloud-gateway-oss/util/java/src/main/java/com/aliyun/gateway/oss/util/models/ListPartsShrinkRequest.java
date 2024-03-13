// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListPartsShrinkRequest extends TeaModel {
    /**
     * <p>The maximum number of parts that can be returned by OSS. </p>
     * <br>
     * <p>Default value: 1000.</p>
     * <br>
     * <p>Maximum value: 1000.</p>
     */
    @NameInMap("encoding-type")
    public String encodingTypeShrink;

    /**
     * <p>The maximum number of parts that can be returned by OSS.</p>
     * <br>
     * <p>Default value: 1000.</p>
     * <br>
     * <p>Maximum value: 1000.</p>
     */
    @NameInMap("max-parts")
    public Long maxParts;

    /**
     * <p>The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.</p>
     * <br>
     * <p>By default, this parameter is left empty.</p>
     */
    @NameInMap("part-number-marker")
    public Long partNumberMarker;

    /**
     * <p>The ID of the multipart upload task.</p>
     * <br>
     * <p>By default, this parameter is left empty.</p>
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
