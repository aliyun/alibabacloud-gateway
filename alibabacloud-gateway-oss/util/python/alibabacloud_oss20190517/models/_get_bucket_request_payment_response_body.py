# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketRequestPaymentResponseBody(DaraModel):
    def __init__(
        self,
        request_payment_configuration: main_models.GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration = None,
    ):
        # Indicates the container for the payer.
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
            temp_model = main_models.GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration()
            self.request_payment_configuration = temp_model.from_map(m.get('RequestPaymentConfiguration'))

        return self

class GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration(DaraModel):
    def __init__(
        self,
        payer: str = None,
    ):
        # Indicates who pays the download and request fees.
        self.payer = payer

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.payer is not None:
            result['Payer'] = self.payer

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Payer') is not None:
            self.payer = m.get('Payer')

        return self

