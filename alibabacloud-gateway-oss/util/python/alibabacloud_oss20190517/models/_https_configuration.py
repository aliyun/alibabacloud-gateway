# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class HttpsConfiguration(DaraModel):
    def __init__(
        self,
        cipher_suite: main_models.HttpsConfigurationCipherSuite = None,
        tls: main_models.HttpsConfigurationTLS = None,
    ):
        self.cipher_suite = cipher_suite
        self.tls = tls

    def validate(self):
        if self.cipher_suite:
            self.cipher_suite.validate()
        if self.tls:
            self.tls.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cipher_suite is not None:
            result['CipherSuite'] = self.cipher_suite.to_map()

        if self.tls is not None:
            result['TLS'] = self.tls.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CipherSuite') is not None:
            temp_model = main_models.HttpsConfigurationCipherSuite()
            self.cipher_suite = temp_model.from_map(m.get('CipherSuite'))

        if m.get('TLS') is not None:
            temp_model = main_models.HttpsConfigurationTLS()
            self.tls = temp_model.from_map(m.get('TLS'))

        return self

class HttpsConfigurationTLS(DaraModel):
    def __init__(
        self,
        enable: bool = None,
        tlsversion: List[str] = None,
    ):
        # This parameter is required.
        self.enable = enable
        self.tlsversion = tlsversion

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.enable is not None:
            result['Enable'] = self.enable

        if self.tlsversion is not None:
            result['TLSVersion'] = self.tlsversion

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Enable') is not None:
            self.enable = m.get('Enable')

        if m.get('TLSVersion') is not None:
            self.tlsversion = m.get('TLSVersion')

        return self

class HttpsConfigurationCipherSuite(DaraModel):
    def __init__(
        self,
        custom_cipher_suite: List[str] = None,
        enable: bool = None,
        strong_cipher_suite: bool = None,
        tls13custom_cipher_suite: List[str] = None,
    ):
        self.custom_cipher_suite = custom_cipher_suite
        self.enable = enable
        self.strong_cipher_suite = strong_cipher_suite
        self.tls13custom_cipher_suite = tls13custom_cipher_suite

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.custom_cipher_suite is not None:
            result['CustomCipherSuite'] = self.custom_cipher_suite

        if self.enable is not None:
            result['Enable'] = self.enable

        if self.strong_cipher_suite is not None:
            result['StrongCipherSuite'] = self.strong_cipher_suite

        if self.tls13custom_cipher_suite is not None:
            result['TLS13CustomCipherSuite'] = self.tls13custom_cipher_suite

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CustomCipherSuite') is not None:
            self.custom_cipher_suite = m.get('CustomCipherSuite')

        if m.get('Enable') is not None:
            self.enable = m.get('Enable')

        if m.get('StrongCipherSuite') is not None:
            self.strong_cipher_suite = m.get('StrongCipherSuite')

        if m.get('TLS13CustomCipherSuite') is not None:
            self.tls13custom_cipher_suite = m.get('TLS13CustomCipherSuite')

        return self

