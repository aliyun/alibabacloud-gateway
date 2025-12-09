# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BucketCnameConfiguration(DaraModel):
    def __init__(
        self,
        cname: main_models.BucketCnameConfigurationCname = None,
    ):
        self.cname = cname

    def validate(self):
        if self.cname:
            self.cname.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cname is not None:
            result['Cname'] = self.cname.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cname') is not None:
            temp_model = main_models.BucketCnameConfigurationCname()
            self.cname = temp_model.from_map(m.get('Cname'))

        return self

class BucketCnameConfigurationCname(DaraModel):
    def __init__(
        self,
        certificate_configuration: main_models.CertificateConfiguration = None,
        domain: str = None,
    ):
        self.certificate_configuration = certificate_configuration
        self.domain = domain

    def validate(self):
        if self.certificate_configuration:
            self.certificate_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.certificate_configuration is not None:
            result['CertificateConfiguration'] = self.certificate_configuration.to_map()

        if self.domain is not None:
            result['Domain'] = self.domain

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CertificateConfiguration') is not None:
            temp_model = main_models.CertificateConfiguration()
            self.certificate_configuration = temp_model.from_map(m.get('CertificateConfiguration'))

        if m.get('Domain') is not None:
            self.domain = m.get('Domain')

        return self

