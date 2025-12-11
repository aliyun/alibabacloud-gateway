# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CnameSummary(DaraModel):
    def __init__(
        self,
        certificate: main_models.CnameCertificate = None,
        domain: str = None,
        last_modified: str = None,
        status: str = None,
    ):
        self.certificate = certificate
        self.domain = domain
        self.last_modified = last_modified
        self.status = status

    def validate(self):
        if self.certificate:
            self.certificate.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.certificate is not None:
            result['Certificate'] = self.certificate.to_map()

        if self.domain is not None:
            result['Domain'] = self.domain

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Certificate') is not None:
            temp_model = main_models.CnameCertificate()
            self.certificate = temp_model.from_map(m.get('Certificate'))

        if m.get('Domain') is not None:
            self.domain = m.get('Domain')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self

