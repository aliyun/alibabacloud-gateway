// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CertificateConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>493****-cn-hangzhou</p>
     */
    @NameInMap("CertId")
    public String certId;

    /**
     * <strong>example:</strong>
     * <p>-----BEGIN CERTIFICATE----- MIIDhDCCAmwCCQCFs8ixARsyrDANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMC **** -----END CERTIFICATE-----</p>
     */
    @NameInMap("Certificate")
    public String certificate;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("DeleteCertificate")
    public Boolean deleteCertificate;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("Force")
    public Boolean force;

    /**
     * <strong>example:</strong>
     * <p>493****-cn-hangzhou</p>
     */
    @NameInMap("PreviousCertId")
    public String previousCertId;

    /**
     * <strong>example:</strong>
     * <p>-----BEGIN CERTIFICATE----- MIIDhDCCAmwCCQCFs8ixARsyrDANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMC **** -----END CERTIFICATE-----</p>
     */
    @NameInMap("PrivateKey")
    public String privateKey;

    public static CertificateConfiguration build(java.util.Map<String, ?> map) throws Exception {
        CertificateConfiguration self = new CertificateConfiguration();
        return TeaModel.build(map, self);
    }

    public CertificateConfiguration setCertId(String certId) {
        this.certId = certId;
        return this;
    }
    public String getCertId() {
        return this.certId;
    }

    public CertificateConfiguration setCertificate(String certificate) {
        this.certificate = certificate;
        return this;
    }
    public String getCertificate() {
        return this.certificate;
    }

    public CertificateConfiguration setDeleteCertificate(Boolean deleteCertificate) {
        this.deleteCertificate = deleteCertificate;
        return this;
    }
    public Boolean getDeleteCertificate() {
        return this.deleteCertificate;
    }

    public CertificateConfiguration setForce(Boolean force) {
        this.force = force;
        return this;
    }
    public Boolean getForce() {
        return this.force;
    }

    public CertificateConfiguration setPreviousCertId(String previousCertId) {
        this.previousCertId = previousCertId;
        return this;
    }
    public String getPreviousCertId() {
        return this.previousCertId;
    }

    public CertificateConfiguration setPrivateKey(String privateKey) {
        this.privateKey = privateKey;
        return this;
    }
    public String getPrivateKey() {
        return this.privateKey;
    }

}
