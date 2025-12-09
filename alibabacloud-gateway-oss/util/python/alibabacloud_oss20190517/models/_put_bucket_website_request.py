# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketWebsiteRequest(DaraModel):
    def __init__(
        self,
        website_configuration: main_models.WebsiteConfiguration = None,
    ):
        # The container that stores the website configuration.
        self.website_configuration = website_configuration

    def validate(self):
        if self.website_configuration:
            self.website_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.website_configuration is not None:
            result['WebsiteConfiguration'] = self.website_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('WebsiteConfiguration') is not None:
            temp_model = main_models.WebsiteConfiguration()
            self.website_configuration = temp_model.from_map(m.get('WebsiteConfiguration'))

        return self

