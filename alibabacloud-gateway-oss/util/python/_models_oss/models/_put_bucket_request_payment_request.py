# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketRequestPaymentRequest(DaraModel):
    def __init__(
        self,
        request_payment_configuration: main_models.RequestPaymentConfiguration = None,
    ):
        # The container that stores pay-by-requester configurations.
        self.request_payment_configuration = request_payment_configuration

    def validate(self):
        if self.request_payment_configuration:
            self.request_payment_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.request_payment_configuration is not None:
            result['RequestPaymentConfiguration'] = self.request_payment_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RequestPaymentConfiguration') is not None:
            temp_model = main_models.RequestPaymentConfiguration()
            self.request_payment_configuration = temp_model.from_map(m.get('RequestPaymentConfiguration'))

        return self

