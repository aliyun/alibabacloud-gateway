# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CORSConfiguration(DaraModel):
    def __init__(
        self,
        corsrule: List[main_models.CORSRule] = None,
        response_vary: bool = None,
    ):
        # The container that stores the CORS rules.
        # 
        # Up to 10 CORS rules can be configured for a bucket. The XML message body in a request can be up to 16 KB in size.
        self.corsrule = corsrule
        # Indicates whether the Vary: Origin header is returned. Valid values:
        # 
        # *   true: The Vary: Origin header is returned regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds.
        # *   false (default): The Vary: Origin header is not returned.
        # 
        # >  This parameter is valid only when at least one CORS rule is configured.
        self.response_vary = response_vary

    def validate(self):
        if self.corsrule:
            for v1 in self.corsrule:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['CORSRule'] = []
        if self.corsrule is not None:
            for k1 in self.corsrule:
                result['CORSRule'].append(k1.to_map() if k1 else None)

        if self.response_vary is not None:
            result['ResponseVary'] = self.response_vary

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.corsrule = []
        if m.get('CORSRule') is not None:
            for k1 in m.get('CORSRule'):
                temp_model = main_models.CORSRule()
                self.corsrule.append(temp_model.from_map(k1))

        if m.get('ResponseVary') is not None:
            self.response_vary = m.get('ResponseVary')

        return self

