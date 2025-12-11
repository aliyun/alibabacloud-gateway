# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class RegionInfo(DaraModel):
    def __init__(
        self,
        accelerate_endpoint: str = None,
        internal_endpoint: str = None,
        internet_endpoint: str = None,
        region: str = None,
    ):
        self.accelerate_endpoint = accelerate_endpoint
        self.internal_endpoint = internal_endpoint
        self.internet_endpoint = internet_endpoint
        self.region = region

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.accelerate_endpoint is not None:
            result['AccelerateEndpoint'] = self.accelerate_endpoint

        if self.internal_endpoint is not None:
            result['InternalEndpoint'] = self.internal_endpoint

        if self.internet_endpoint is not None:
            result['InternetEndpoint'] = self.internet_endpoint

        if self.region is not None:
            result['Region'] = self.region

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccelerateEndpoint') is not None:
            self.accelerate_endpoint = m.get('AccelerateEndpoint')

        if m.get('InternalEndpoint') is not None:
            self.internal_endpoint = m.get('InternalEndpoint')

        if m.get('InternetEndpoint') is not None:
            self.internet_endpoint = m.get('InternetEndpoint')

        if m.get('Region') is not None:
            self.region = m.get('Region')

        return self

