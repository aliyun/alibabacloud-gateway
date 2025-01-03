// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UserQosConfiguration extends TeaModel {
    @NameInMap("DefaultQoSConfiguration")
    public QoSConfigurationWithRemark defaultQoSConfiguration;

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
     * <p>oss-cn-shanghai</p>
     */
    @NameInMap("Region")
    public String region;

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

    public static UserQosConfiguration build(java.util.Map<String, ?> map) throws Exception {
        UserQosConfiguration self = new UserQosConfiguration();
        return TeaModel.build(map, self);
    }

    public UserQosConfiguration setDefaultQoSConfiguration(QoSConfigurationWithRemark defaultQoSConfiguration) {
        this.defaultQoSConfiguration = defaultQoSConfiguration;
        return this;
    }
    public QoSConfigurationWithRemark getDefaultQoSConfiguration() {
        return this.defaultQoSConfiguration;
    }

    public UserQosConfiguration setExtranetDownloadBandwidth(Long extranetDownloadBandwidth) {
        this.extranetDownloadBandwidth = extranetDownloadBandwidth;
        return this;
    }
    public Long getExtranetDownloadBandwidth() {
        return this.extranetDownloadBandwidth;
    }

    public UserQosConfiguration setExtranetQps(Long extranetQps) {
        this.extranetQps = extranetQps;
        return this;
    }
    public Long getExtranetQps() {
        return this.extranetQps;
    }

    public UserQosConfiguration setExtranetUploadBandwidth(Long extranetUploadBandwidth) {
        this.extranetUploadBandwidth = extranetUploadBandwidth;
        return this;
    }
    public Long getExtranetUploadBandwidth() {
        return this.extranetUploadBandwidth;
    }

    public UserQosConfiguration setIntranetDownloadBandwidth(Long intranetDownloadBandwidth) {
        this.intranetDownloadBandwidth = intranetDownloadBandwidth;
        return this;
    }
    public Long getIntranetDownloadBandwidth() {
        return this.intranetDownloadBandwidth;
    }

    public UserQosConfiguration setIntranetQps(Long intranetQps) {
        this.intranetQps = intranetQps;
        return this;
    }
    public Long getIntranetQps() {
        return this.intranetQps;
    }

    public UserQosConfiguration setIntranetUploadBandwidth(Long intranetUploadBandwidth) {
        this.intranetUploadBandwidth = intranetUploadBandwidth;
        return this;
    }
    public Long getIntranetUploadBandwidth() {
        return this.intranetUploadBandwidth;
    }

    public UserQosConfiguration setRegion(String region) {
        this.region = region;
        return this;
    }
    public String getRegion() {
        return this.region;
    }

    public UserQosConfiguration setRemark(Long remark) {
        this.remark = remark;
        return this;
    }
    public Long getRemark() {
        return this.remark;
    }

    public UserQosConfiguration setTotalDownloadBandwidth(Long totalDownloadBandwidth) {
        this.totalDownloadBandwidth = totalDownloadBandwidth;
        return this;
    }
    public Long getTotalDownloadBandwidth() {
        return this.totalDownloadBandwidth;
    }

    public UserQosConfiguration setTotalQps(Long totalQps) {
        this.totalQps = totalQps;
        return this;
    }
    public Long getTotalQps() {
        return this.totalQps;
    }

    public UserQosConfiguration setTotalUploadBandwidth(Long totalUploadBandwidth) {
        this.totalUploadBandwidth = totalUploadBandwidth;
        return this;
    }
    public Long getTotalUploadBandwidth() {
        return this.totalUploadBandwidth;
    }

}
