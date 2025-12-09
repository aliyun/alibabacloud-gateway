# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CnameCertificate(DaraModel):
    def __init__(
        self,
        cert_id: str = None,
        creation_date: str = None,
        fingerprint: str = None,
        status: str = None,
        type: str = None,
        valid_end_date: str = None,
        valid_start_date: str = None,
    ):
        self.cert_id = cert_id
        self.creation_date = creation_date
        self.fingerprint = fingerprint
        self.status = status
        self.type = type
        self.valid_end_date = valid_end_date
        self.valid_start_date = valid_start_date

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cert_id is not None:
            result['CertId'] = self.cert_id

        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.fingerprint is not None:
            result['Fingerprint'] = self.fingerprint

        if self.status is not None:
            result['Status'] = self.status

        if self.type is not None:
            result['Type'] = self.type

        if self.valid_end_date is not None:
            result['ValidEndDate'] = self.valid_end_date

        if self.valid_start_date is not None:
            result['ValidStartDate'] = self.valid_start_date

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CertId') is not None:
            self.cert_id = m.get('CertId')

        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('Fingerprint') is not None:
            self.fingerprint = m.get('Fingerprint')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        if m.get('ValidEndDate') is not None:
            self.valid_end_date = m.get('ValidEndDate')

        if m.get('ValidStartDate') is not None:
            self.valid_start_date = m.get('ValidStartDate')

        return self

