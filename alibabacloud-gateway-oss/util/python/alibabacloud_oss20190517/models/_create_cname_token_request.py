# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CreateCnameTokenRequest(DaraModel):
    def __init__(
        self,
        bucket_cname_configuration: main_models.CreateCnameTokenRequestBucketCnameConfiguration = None,
    ):
        # The container that stores the CNAME record.
        self.bucket_cname_configuration = bucket_cname_configuration

    def validate(self):
        if self.bucket_cname_configuration:
            self.bucket_cname_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_cname_configuration is not None:
            result['BucketCnameConfiguration'] = self.bucket_cname_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCnameConfiguration') is not None:
            temp_model = main_models.CreateCnameTokenRequestBucketCnameConfiguration()
            self.bucket_cname_configuration = temp_model.from_map(m.get('BucketCnameConfiguration'))

        return self

class CreateCnameTokenRequestBucketCnameConfiguration(DaraModel):
    def __init__(
        self,
        cname: main_models.CreateCnameTokenRequestBucketCnameConfigurationCname = None,
    ):
        # The container that stores the CNAME information.
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
            temp_model = main_models.CreateCnameTokenRequestBucketCnameConfigurationCname()
            self.cname = temp_model.from_map(m.get('Cname'))

        return self

class CreateCnameTokenRequestBucketCnameConfigurationCname(DaraModel):
    def __init__(
        self,
        domain: str = None,
    ):
        # The custom domain name.
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.domain is not None:
            result['Domain'] = self.domain

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')

        return self

