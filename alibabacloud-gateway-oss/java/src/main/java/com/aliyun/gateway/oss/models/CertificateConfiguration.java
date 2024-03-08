// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class CertificateConfiguration extends TeaModel {
    @NameInMap("CertId")
    public String certId;

    @NameInMap("Certificate")
    public String certificate;

    @NameInMap("DeleteCertificate")
    public Boolean deleteCertificate;

    @NameInMap("Force")
    public Boolean force;

    @NameInMap("PreviousCertId")
    public String previousCertId;

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
