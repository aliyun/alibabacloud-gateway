# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BucketStat(DaraModel):
    def __init__(
        self,
        archive_multipart_part_count: int = None,
        archive_multipart_part_storage: int = None,
        archive_object_count: int = None,
        archive_real_storage: int = None,
        archive_storage: int = None,
        cold_archive_multipart_part_count: int = None,
        cold_archive_multipart_part_storage: int = None,
        cold_archive_object_count: int = None,
        cold_archive_real_storage: int = None,
        cold_archive_storage: int = None,
        deep_cold_archive_multipart_part_count: int = None,
        deep_cold_archive_multipart_part_storage: int = None,
        deep_cold_archive_object_count: int = None,
        deep_cold_archive_real_storage: int = None,
        deep_cold_archive_storage: int = None,
        delete_marker_count: int = None,
        infrequent_access_object_count: int = None,
        infrequent_access_real_storage: int = None,
        infrequent_access_storage: int = None,
        infrequent_multipart_part_count: int = None,
        infrequent_multipart_part_storage: int = None,
        last_modified_time: int = None,
        live_channel_count: int = None,
        multipart_part_count: int = None,
        multipart_part_storage: int = None,
        multipart_upload_count: int = None,
        object_count: int = None,
        standard_multipart_part_count: int = None,
        standard_multipart_part_storage: int = None,
        standard_object_count: int = None,
        standard_storage: int = None,
        storage: int = None,
    ):
        # The number of Archive Multipart parts in the bucket.
        self.archive_multipart_part_count = archive_multipart_part_count
        # The storage usage of Archive Multipart parts in the bucket. Unit: bytes.
        self.archive_multipart_part_storage = archive_multipart_part_storage
        # The number of Archive objects in the bucket.
        self.archive_object_count = archive_object_count
        # The actual storage usage of Archive objects in the bucket. Unit: bytes.
        self.archive_real_storage = archive_real_storage
        # The billed storage usage of Archive objects in the bucket. Unit: bytes.
        self.archive_storage = archive_storage
        # The number of Cold Archive Multipart parts in the bucket.
        self.cold_archive_multipart_part_count = cold_archive_multipart_part_count
        # The storage usage of Cold Archive Multipart parts in the bucket. Unit: bytes.
        self.cold_archive_multipart_part_storage = cold_archive_multipart_part_storage
        # The number of Cold Archive objects in the bucket.
        self.cold_archive_object_count = cold_archive_object_count
        # The actual storage usage of Cold Archive objects in the bucket. Unit: bytes.
        self.cold_archive_real_storage = cold_archive_real_storage
        # The billed storage usage of Cold Archive objects in the bucket. Unit: bytes.
        self.cold_archive_storage = cold_archive_storage
        # The number of Deep Cold Archive Multipart parts in the bucket.
        self.deep_cold_archive_multipart_part_count = deep_cold_archive_multipart_part_count
        # The storage usage of Deep Cold Archive Multipart parts in the bucket. Unit: bytes.
        self.deep_cold_archive_multipart_part_storage = deep_cold_archive_multipart_part_storage
        # The number of Deep Cold Archive objects in the bucket.
        self.deep_cold_archive_object_count = deep_cold_archive_object_count
        # The actual storage usage of Deep Cold Archive objects in the bucket. Unit: bytes.
        self.deep_cold_archive_real_storage = deep_cold_archive_real_storage
        # The billed storage usage of Deep Cold Archive objects in the bucket. Unit: bytes.
        self.deep_cold_archive_storage = deep_cold_archive_storage
        # The number of deletemarker in the bucket.
        self.delete_marker_count = delete_marker_count
        # The number of IA objects in the bucket.
        self.infrequent_access_object_count = infrequent_access_object_count
        # The actual storage usage of IA objects in the bucket. Unit: bytes.
        self.infrequent_access_real_storage = infrequent_access_real_storage
        # The billed storage usage of IA objects in the bucket. Unit: bytes.
        self.infrequent_access_storage = infrequent_access_storage
        # The number of IA Multipart parts in the bucket.
        self.infrequent_multipart_part_count = infrequent_multipart_part_count
        # The storage usage of IA Multipart parts in the bucket. Unit: bytes.
        self.infrequent_multipart_part_storage = infrequent_multipart_part_storage
        # The time when the obtained information was last modified. The value of this parameter is a UNIX timestamp. Unit: seconds.
        self.last_modified_time = last_modified_time
        # The number of LiveChannels in the bucket.
        self.live_channel_count = live_channel_count
        # The number of mulitpart parts in the bucket.
        self.multipart_part_count = multipart_part_count
        # The storage usage of Multipart parts in the bucket. Unit: bytes.
        self.multipart_part_storage = multipart_part_storage
        # The number of multipart upload tasks that have been initiated but are not completed or canceled.
        self.multipart_upload_count = multipart_upload_count
        # The total number of objects in the bucket.
        self.object_count = object_count
        # The number of Standard Multipart parts in the bucket.
        self.standard_multipart_part_count = standard_multipart_part_count
        # The storage usage of Standard Multipart parts in the bucket. Unit: bytes.
        self.standard_multipart_part_storage = standard_multipart_part_storage
        # The number of Standard objects in the bucket.
        self.standard_object_count = standard_object_count
        # The storage usage of Standard objects in the bucket. Unit: bytes.
        self.standard_storage = standard_storage
        # The storage usage of the bucket. Unit: bytes.
        self.storage = storage

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.archive_multipart_part_count is not None:
            result['ArchiveMultipartPartCount'] = self.archive_multipart_part_count

        if self.archive_multipart_part_storage is not None:
            result['ArchiveMultipartPartStorage'] = self.archive_multipart_part_storage

        if self.archive_object_count is not None:
            result['ArchiveObjectCount'] = self.archive_object_count

        if self.archive_real_storage is not None:
            result['ArchiveRealStorage'] = self.archive_real_storage

        if self.archive_storage is not None:
            result['ArchiveStorage'] = self.archive_storage

        if self.cold_archive_multipart_part_count is not None:
            result['ColdArchiveMultipartPartCount'] = self.cold_archive_multipart_part_count

        if self.cold_archive_multipart_part_storage is not None:
            result['ColdArchiveMultipartPartStorage'] = self.cold_archive_multipart_part_storage

        if self.cold_archive_object_count is not None:
            result['ColdArchiveObjectCount'] = self.cold_archive_object_count

        if self.cold_archive_real_storage is not None:
            result['ColdArchiveRealStorage'] = self.cold_archive_real_storage

        if self.cold_archive_storage is not None:
            result['ColdArchiveStorage'] = self.cold_archive_storage

        if self.deep_cold_archive_multipart_part_count is not None:
            result['DeepColdArchiveMultipartPartCount'] = self.deep_cold_archive_multipart_part_count

        if self.deep_cold_archive_multipart_part_storage is not None:
            result['DeepColdArchiveMultipartPartStorage'] = self.deep_cold_archive_multipart_part_storage

        if self.deep_cold_archive_object_count is not None:
            result['DeepColdArchiveObjectCount'] = self.deep_cold_archive_object_count

        if self.deep_cold_archive_real_storage is not None:
            result['DeepColdArchiveRealStorage'] = self.deep_cold_archive_real_storage

        if self.deep_cold_archive_storage is not None:
            result['DeepColdArchiveStorage'] = self.deep_cold_archive_storage

        if self.delete_marker_count is not None:
            result['DeleteMarkerCount'] = self.delete_marker_count

        if self.infrequent_access_object_count is not None:
            result['InfrequentAccessObjectCount'] = self.infrequent_access_object_count

        if self.infrequent_access_real_storage is not None:
            result['InfrequentAccessRealStorage'] = self.infrequent_access_real_storage

        if self.infrequent_access_storage is not None:
            result['InfrequentAccessStorage'] = self.infrequent_access_storage

        if self.infrequent_multipart_part_count is not None:
            result['InfrequentMultipartPartCount'] = self.infrequent_multipart_part_count

        if self.infrequent_multipart_part_storage is not None:
            result['InfrequentMultipartPartStorage'] = self.infrequent_multipart_part_storage

        if self.last_modified_time is not None:
            result['LastModifiedTime'] = self.last_modified_time

        if self.live_channel_count is not None:
            result['LiveChannelCount'] = self.live_channel_count

        if self.multipart_part_count is not None:
            result['MultipartPartCount'] = self.multipart_part_count

        if self.multipart_part_storage is not None:
            result['MultipartPartStorage'] = self.multipart_part_storage

        if self.multipart_upload_count is not None:
            result['MultipartUploadCount'] = self.multipart_upload_count

        if self.object_count is not None:
            result['ObjectCount'] = self.object_count

        if self.standard_multipart_part_count is not None:
            result['StandardMultipartPartCount'] = self.standard_multipart_part_count

        if self.standard_multipart_part_storage is not None:
            result['StandardMultipartPartStorage'] = self.standard_multipart_part_storage

        if self.standard_object_count is not None:
            result['StandardObjectCount'] = self.standard_object_count

        if self.standard_storage is not None:
            result['StandardStorage'] = self.standard_storage

        if self.storage is not None:
            result['Storage'] = self.storage

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ArchiveMultipartPartCount') is not None:
            self.archive_multipart_part_count = m.get('ArchiveMultipartPartCount')

        if m.get('ArchiveMultipartPartStorage') is not None:
            self.archive_multipart_part_storage = m.get('ArchiveMultipartPartStorage')

        if m.get('ArchiveObjectCount') is not None:
            self.archive_object_count = m.get('ArchiveObjectCount')

        if m.get('ArchiveRealStorage') is not None:
            self.archive_real_storage = m.get('ArchiveRealStorage')

        if m.get('ArchiveStorage') is not None:
            self.archive_storage = m.get('ArchiveStorage')

        if m.get('ColdArchiveMultipartPartCount') is not None:
            self.cold_archive_multipart_part_count = m.get('ColdArchiveMultipartPartCount')

        if m.get('ColdArchiveMultipartPartStorage') is not None:
            self.cold_archive_multipart_part_storage = m.get('ColdArchiveMultipartPartStorage')

        if m.get('ColdArchiveObjectCount') is not None:
            self.cold_archive_object_count = m.get('ColdArchiveObjectCount')

        if m.get('ColdArchiveRealStorage') is not None:
            self.cold_archive_real_storage = m.get('ColdArchiveRealStorage')

        if m.get('ColdArchiveStorage') is not None:
            self.cold_archive_storage = m.get('ColdArchiveStorage')

        if m.get('DeepColdArchiveMultipartPartCount') is not None:
            self.deep_cold_archive_multipart_part_count = m.get('DeepColdArchiveMultipartPartCount')

        if m.get('DeepColdArchiveMultipartPartStorage') is not None:
            self.deep_cold_archive_multipart_part_storage = m.get('DeepColdArchiveMultipartPartStorage')

        if m.get('DeepColdArchiveObjectCount') is not None:
            self.deep_cold_archive_object_count = m.get('DeepColdArchiveObjectCount')

        if m.get('DeepColdArchiveRealStorage') is not None:
            self.deep_cold_archive_real_storage = m.get('DeepColdArchiveRealStorage')

        if m.get('DeepColdArchiveStorage') is not None:
            self.deep_cold_archive_storage = m.get('DeepColdArchiveStorage')

        if m.get('DeleteMarkerCount') is not None:
            self.delete_marker_count = m.get('DeleteMarkerCount')

        if m.get('InfrequentAccessObjectCount') is not None:
            self.infrequent_access_object_count = m.get('InfrequentAccessObjectCount')

        if m.get('InfrequentAccessRealStorage') is not None:
            self.infrequent_access_real_storage = m.get('InfrequentAccessRealStorage')

        if m.get('InfrequentAccessStorage') is not None:
            self.infrequent_access_storage = m.get('InfrequentAccessStorage')

        if m.get('InfrequentMultipartPartCount') is not None:
            self.infrequent_multipart_part_count = m.get('InfrequentMultipartPartCount')

        if m.get('InfrequentMultipartPartStorage') is not None:
            self.infrequent_multipart_part_storage = m.get('InfrequentMultipartPartStorage')

        if m.get('LastModifiedTime') is not None:
            self.last_modified_time = m.get('LastModifiedTime')

        if m.get('LiveChannelCount') is not None:
            self.live_channel_count = m.get('LiveChannelCount')

        if m.get('MultipartPartCount') is not None:
            self.multipart_part_count = m.get('MultipartPartCount')

        if m.get('MultipartPartStorage') is not None:
            self.multipart_part_storage = m.get('MultipartPartStorage')

        if m.get('MultipartUploadCount') is not None:
            self.multipart_upload_count = m.get('MultipartUploadCount')

        if m.get('ObjectCount') is not None:
            self.object_count = m.get('ObjectCount')

        if m.get('StandardMultipartPartCount') is not None:
            self.standard_multipart_part_count = m.get('StandardMultipartPartCount')

        if m.get('StandardMultipartPartStorage') is not None:
            self.standard_multipart_part_storage = m.get('StandardMultipartPartStorage')

        if m.get('StandardObjectCount') is not None:
            self.standard_object_count = m.get('StandardObjectCount')

        if m.get('StandardStorage') is not None:
            self.standard_storage = m.get('StandardStorage')

        if m.get('Storage') is not None:
            self.storage = m.get('Storage')

        return self

