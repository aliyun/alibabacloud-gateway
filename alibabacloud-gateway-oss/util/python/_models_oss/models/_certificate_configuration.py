# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CertificateConfiguration(DaraModel):
    def __init__(
        self,
        cert_id: str = None,
        certificate: str = None,
        delete_certificate: bool = None,
        force: bool = None,
        previous_cert_id: str = None,
        private_key: str = None,
    ):
        # The ID of the certificate.
        self.cert_id = cert_id
        # The public key of the certificate.
        self.certificate = certificate
        # Specifies whether to delete the certificate. Valid values:
        # - true: deletes the certificate.
        # - false: does not delete the certificate.
        self.delete_certificate = delete_certificate
        # Specifies whether to overwrite the certificate. Valid values:
        # - true: overwrites the certificate.
        # - false: does not overwrite the certificate.
        self.force = force
        # The ID of the certificate. If the Force parameter is not set to true, the OSS server checks whether the value of the Force parameter matches the current certificate ID. If the value does not match the certificate ID, an error is returned.
        # <notice>If you do not specify the PreviousCertId parameter when you bind a certificate, you must set the Force parameter to true.</notice>
        self.previous_cert_id = previous_cert_id
        # The private key of the certificate.
        self.private_key = private_key

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cert_id is not None:
            result['CertId'] = self.cert_id

        if self.certificate is not None:
            result['Certificate'] = self.certificate

        if self.delete_certificate is not None:
            result['DeleteCertificate'] = self.delete_certificate

        if self.force is not None:
            result['Force'] = self.force

        if self.previous_cert_id is not None:
            result['PreviousCertId'] = self.previous_cert_id

        if self.private_key is not None:
            result['PrivateKey'] = self.private_key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CertId') is not None:
            self.cert_id = m.get('CertId')

        if m.get('Certificate') is not None:
            self.certificate = m.get('Certificate')

        if m.get('DeleteCertificate') is not None:
            self.delete_certificate = m.get('DeleteCertificate')

        if m.get('Force') is not None:
            self.force = m.get('Force')

        if m.get('PreviousCertId') is not None:
            self.previous_cert_id = m.get('PreviousCertId')

        if m.get('PrivateKey') is not None:
            self.private_key = m.get('PrivateKey')

        return self

