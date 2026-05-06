# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class Bucket(DaraModel):
    def __init__(
        self,
        creation_date: str = None,
        extranet_endpoint: str = None,
        intranet_endpoint: str = None,
        location: str = None,
        name: str = None,
        region: str = None,
        storage_class: str = None,
    ):
        # The time when the bucket was created. Format: `yyyy-mm-ddThh:mm:ss.timezone`.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.creation_date = creation_date
        # The public endpoint of the region in which the bucket resides.
        self.extranet_endpoint = extranet_endpoint
        # The internal endpoint of the region in which the bucket you access from ECS instances resides. The bucket and ECS instances are in the same region.
        self.intranet_endpoint = intranet_endpoint
        # The data center in which the bucket is located.
        self.location = location
        # The name of the bucket.
        self.name = name
        # The region in which the bucket is located.
        self.region = region
        # The storage class of the bucket. Valid values: Standard, IA, Archive, and ColdArchive.
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.extranet_endpoint is not None:
            result['ExtranetEndpoint'] = self.extranet_endpoint

        if self.intranet_endpoint is not None:
            result['IntranetEndpoint'] = self.intranet_endpoint

        if self.location is not None:
            result['Location'] = self.location

        if self.name is not None:
            result['Name'] = self.name

        if self.region is not None:
            result['Region'] = self.region

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('ExtranetEndpoint') is not None:
            self.extranet_endpoint = m.get('ExtranetEndpoint')

        if m.get('IntranetEndpoint') is not None:
            self.intranet_endpoint = m.get('IntranetEndpoint')

        if m.get('Location') is not None:
            self.location = m.get('Location')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Region') is not None:
            self.region = m.get('Region')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        return self

