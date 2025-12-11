# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketReplicationLocationResponseBody(DaraModel):
    def __init__(
        self,
        replication_location: main_models.GetBucketReplicationLocationResponseBodyReplicationLocation = None,
    ):
        # The container that stores the region in which the destination bucket can be located.
        self.replication_location = replication_location

    def validate(self):
        if self.replication_location:
            self.replication_location.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.replication_location is not None:
            result['ReplicationLocation'] = self.replication_location.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationLocation') is not None:
            temp_model = main_models.GetBucketReplicationLocationResponseBodyReplicationLocation()
            self.replication_location = temp_model.from_map(m.get('ReplicationLocation'))

        return self

class GetBucketReplicationLocationResponseBodyReplicationLocation(DaraModel):
    def __init__(
        self,
        location: List[str] = None,
        location_rtcconstraint: main_models.GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint = None,
        location_transfer_type_constraint: main_models.GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint = None,
    ):
        # The regions in which the destination bucket can be located.
        self.location = location
        # The container that stores regions in which the RTC can be enabled.
        self.location_rtcconstraint = location_rtcconstraint
        # The container that stores regions in which the destination bucket can be located with TransferType specified.
        self.location_transfer_type_constraint = location_transfer_type_constraint

    def validate(self):
        if self.location_rtcconstraint:
            self.location_rtcconstraint.validate()
        if self.location_transfer_type_constraint:
            self.location_transfer_type_constraint.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.location is not None:
            result['Location'] = self.location

        if self.location_rtcconstraint is not None:
            result['LocationRTCConstraint'] = self.location_rtcconstraint.to_map()

        if self.location_transfer_type_constraint is not None:
            result['LocationTransferTypeConstraint'] = self.location_transfer_type_constraint.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            self.location = m.get('Location')

        if m.get('LocationRTCConstraint') is not None:
            temp_model = main_models.GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint()
            self.location_rtcconstraint = temp_model.from_map(m.get('LocationRTCConstraint'))

        if m.get('LocationTransferTypeConstraint') is not None:
            temp_model = main_models.GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint()
            self.location_transfer_type_constraint = temp_model.from_map(m.get('LocationTransferTypeConstraint'))

        return self

class GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint(DaraModel):
    def __init__(
        self,
        location_transfer_type: List[main_models.LocationTransferType] = None,
    ):
        # The container that stores regions in which the destination bucket can be located with the TransferType information.
        self.location_transfer_type = location_transfer_type

    def validate(self):
        if self.location_transfer_type:
            for v1 in self.location_transfer_type:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['LocationTransferType'] = []
        if self.location_transfer_type is not None:
            for k1 in self.location_transfer_type:
                result['LocationTransferType'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.location_transfer_type = []
        if m.get('LocationTransferType') is not None:
            for k1 in m.get('LocationTransferType'):
                temp_model = main_models.LocationTransferType()
                self.location_transfer_type.append(temp_model.from_map(k1))

        return self

class GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint(DaraModel):
    def __init__(
        self,
        location: List[str] = None,
    ):
        # The regions where RTC is supported.
        self.location = location

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.location is not None:
            result['Location'] = self.location

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            self.location = m.get('Location')

        return self

