// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BucketQoSConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("Exclusive")
    public Boolean exclusive;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("ExtranetDownloadBandwidth")
    public Long extranetDownloadBandwidth;

    /**
     * <strong>example:</strong>
     * <p>20000</p>
     */
    @NameInMap("ExtranetQps")
    public Long extranetQps;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("ExtranetUploadBandwidth")
    public Long extranetUploadBandwidth;

    /**
     * <strong>example:</strong>
     * <p>3</p>
     */
    @NameInMap("IntranetDownloadBandwidth")
    public Long intranetDownloadBandwidth;

    /**
     * <strong>example:</strong>
     * <p>20000</p>
     */
    @NameInMap("IntranetQps")
    public Long intranetQps;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("IntranetUploadBandwidth")
    public Long intranetUploadBandwidth;

    /**
     * <strong>example:</strong>
     * <p>5</p>
     */
    @NameInMap("TotalDownloadBandwidth")
    public Long totalDownloadBandwidth;

    /**
     * <strong>example:</strong>
     * <p>20000</p>
     */
    @NameInMap("TotalQps")
    public Long totalQps;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("TotalUploadBandwidth")
    public Long totalUploadBandwidth;

    public static BucketQoSConfiguration build(java.util.Map<String, ?> map) throws Exception {
        BucketQoSConfiguration self = new BucketQoSConfiguration();
        return TeaModel.build(map, self);
    }

    public BucketQoSConfiguration setExclusive(Boolean exclusive) {
        this.exclusive = exclusive;
        return this;
    }
    public Boolean getExclusive() {
        return this.exclusive;
    }

    public BucketQoSConfiguration setExtranetDownloadBandwidth(Long extranetDownloadBandwidth) {
        this.extranetDownloadBandwidth = extranetDownloadBandwidth;
        return this;
    }
    public Long getExtranetDownloadBandwidth() {
        return this.extranetDownloadBandwidth;
    }

    public BucketQoSConfiguration setExtranetQps(Long extranetQps) {
        this.extranetQps = extranetQps;
        return this;
    }
    public Long getExtranetQps() {
        return this.extranetQps;
    }

    public BucketQoSConfiguration setExtranetUploadBandwidth(Long extranetUploadBandwidth) {
        this.extranetUploadBandwidth = extranetUploadBandwidth;
        return this;
    }
    public Long getExtranetUploadBandwidth() {
        return this.extranetUploadBandwidth;
    }

    public BucketQoSConfiguration setIntranetDownloadBandwidth(Long intranetDownloadBandwidth) {
        this.intranetDownloadBandwidth = intranetDownloadBandwidth;
        return this;
    }
    public Long getIntranetDownloadBandwidth() {
        return this.intranetDownloadBandwidth;
    }

    public BucketQoSConfiguration setIntranetQps(Long intranetQps) {
        this.intranetQps = intranetQps;
        return this;
    }
    public Long getIntranetQps() {
        return this.intranetQps;
    }

    public BucketQoSConfiguration setIntranetUploadBandwidth(Long intranetUploadBandwidth) {
        this.intranetUploadBandwidth = intranetUploadBandwidth;
        return this;
    }
    public Long getIntranetUploadBandwidth() {
        return this.intranetUploadBandwidth;
    }

    public BucketQoSConfiguration setTotalDownloadBandwidth(Long totalDownloadBandwidth) {
        this.totalDownloadBandwidth = totalDownloadBandwidth;
        return this;
    }
    public Long getTotalDownloadBandwidth() {
        return this.totalDownloadBandwidth;
    }

    public BucketQoSConfiguration setTotalQps(Long totalQps) {
        this.totalQps = totalQps;
        return this;
    }
    public Long getTotalQps() {
        return this.totalQps;
    }

    public BucketQoSConfiguration setTotalUploadBandwidth(Long totalUploadBandwidth) {
        this.totalUploadBandwidth = totalUploadBandwidth;
        return this;
    }
    public Long getTotalUploadBandwidth() {
        return this.totalUploadBandwidth;
    }

}
