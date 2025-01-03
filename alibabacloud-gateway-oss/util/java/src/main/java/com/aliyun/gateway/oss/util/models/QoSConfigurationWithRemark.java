// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class QoSConfigurationWithRemark extends TeaModel {
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
     * <p>备注</p>
     */
    @NameInMap("Remark")
    public Long remark;

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

    public static QoSConfigurationWithRemark build(java.util.Map<String, ?> map) throws Exception {
        QoSConfigurationWithRemark self = new QoSConfigurationWithRemark();
        return TeaModel.build(map, self);
    }

    public QoSConfigurationWithRemark setExtranetDownloadBandwidth(Long extranetDownloadBandwidth) {
        this.extranetDownloadBandwidth = extranetDownloadBandwidth;
        return this;
    }
    public Long getExtranetDownloadBandwidth() {
        return this.extranetDownloadBandwidth;
    }

    public QoSConfigurationWithRemark setExtranetQps(Long extranetQps) {
        this.extranetQps = extranetQps;
        return this;
    }
    public Long getExtranetQps() {
        return this.extranetQps;
    }

    public QoSConfigurationWithRemark setExtranetUploadBandwidth(Long extranetUploadBandwidth) {
        this.extranetUploadBandwidth = extranetUploadBandwidth;
        return this;
    }
    public Long getExtranetUploadBandwidth() {
        return this.extranetUploadBandwidth;
    }

    public QoSConfigurationWithRemark setIntranetDownloadBandwidth(Long intranetDownloadBandwidth) {
        this.intranetDownloadBandwidth = intranetDownloadBandwidth;
        return this;
    }
    public Long getIntranetDownloadBandwidth() {
        return this.intranetDownloadBandwidth;
    }

    public QoSConfigurationWithRemark setIntranetQps(Long intranetQps) {
        this.intranetQps = intranetQps;
        return this;
    }
    public Long getIntranetQps() {
        return this.intranetQps;
    }

    public QoSConfigurationWithRemark setIntranetUploadBandwidth(Long intranetUploadBandwidth) {
        this.intranetUploadBandwidth = intranetUploadBandwidth;
        return this;
    }
    public Long getIntranetUploadBandwidth() {
        return this.intranetUploadBandwidth;
    }

    public QoSConfigurationWithRemark setRemark(Long remark) {
        this.remark = remark;
        return this;
    }
    public Long getRemark() {
        return this.remark;
    }

    public QoSConfigurationWithRemark setTotalDownloadBandwidth(Long totalDownloadBandwidth) {
        this.totalDownloadBandwidth = totalDownloadBandwidth;
        return this;
    }
    public Long getTotalDownloadBandwidth() {
        return this.totalDownloadBandwidth;
    }

    public QoSConfigurationWithRemark setTotalQps(Long totalQps) {
        this.totalQps = totalQps;
        return this;
    }
    public Long getTotalQps() {
        return this.totalQps;
    }

    public QoSConfigurationWithRemark setTotalUploadBandwidth(Long totalUploadBandwidth) {
        this.totalUploadBandwidth = totalUploadBandwidth;
        return this;
    }
    public Long getTotalUploadBandwidth() {
        return this.totalUploadBandwidth;
    }

}
