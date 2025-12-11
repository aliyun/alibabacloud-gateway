# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListMultipartUploadsResponseBody(DaraModel):
    def __init__(
        self,
        list_multipart_uploads_result: main_models.ListMultipartUploadsResponseBodyListMultipartUploadsResult = None,
    ):
        # The container that stores the response to the ListMultipartUpload request.
        self.list_multipart_uploads_result = list_multipart_uploads_result

    def validate(self):
        if self.list_multipart_uploads_result:
            self.list_multipart_uploads_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_multipart_uploads_result is not None:
            result['ListMultipartUploadsResult'] = self.list_multipart_uploads_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListMultipartUploadsResult') is not None:
            temp_model = main_models.ListMultipartUploadsResponseBodyListMultipartUploadsResult()
            self.list_multipart_uploads_result = temp_model.from_map(m.get('ListMultipartUploadsResult'))

        return self

class ListMultipartUploadsResponseBodyListMultipartUploadsResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        common_prefixes: List[main_models.CommonPrefix] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        key_marker: str = None,
        max_uploads: int = None,
        next_key_marker: str = None,
        next_upload_id_marker: str = None,
        prefix: str = None,
        uploads: List[main_models.Upload] = None,
        upload_id_marker: str = None,
    ):
        # The name of the bucket.
        self.bucket = bucket
        # If the delimiter parameter is specified in the request, the response contains the CommonPrefixes parameter. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.
        self.common_prefixes = common_prefixes
        # The character used to group objects by name. If you specify the Delimiter parameter in the request, the response contains the CommonPrefixes element. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in
        self.delimiter = delimiter
        # The method used to encode the object name in the response. If encoding-type is specified in the request, values of those elements including Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key are encoded in the returned result.
        self.encoding_type = encoding_type
        # Indicates whether the list of multipart upload tasks returned in the response is truncated. Default value: false. Valid values:
        # 
        # - true: Only part of the results are returned this time.
        # 
        # - false: All results are returned.
        self.is_truncated = is_truncated
        # The name of the object that corresponds to the multipart upload task after which the list begins.
        self.key_marker = key_marker
        # The maximum number of multipart upload tasks returned by OSS.
        self.max_uploads = max_uploads
        # The object name marker in the response for the next request to return the remaining results.
        self.next_key_marker = next_key_marker
        # The NextUploadMarker value that is used for the UploadMarker value in the next request if the response does not contain all required results.
        self.next_upload_id_marker = next_upload_id_marker
        # The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.
        self.prefix = prefix
        # The ID list of the multipart upload tasks.
        self.uploads = uploads
        # The upload ID of the multipart upload task after which the list begins.
        self.upload_id_marker = upload_id_marker

    def validate(self):
        if self.common_prefixes:
            for v1 in self.common_prefixes:
                 if v1:
                    v1.validate()
        if self.uploads:
            for v1 in self.uploads:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k1 in self.common_prefixes:
                result['CommonPrefixes'].append(k1.to_map() if k1 else None)

        if self.delimiter is not None:
            result['Delimiter'] = self.delimiter

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.key_marker is not None:
            result['KeyMarker'] = self.key_marker

        if self.max_uploads is not None:
            result['MaxUploads'] = self.max_uploads

        if self.next_key_marker is not None:
            result['NextKeyMarker'] = self.next_key_marker

        if self.next_upload_id_marker is not None:
            result['NextUploadIdMarker'] = self.next_upload_id_marker

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        result['Upload'] = []
        if self.uploads is not None:
            for k1 in self.uploads:
                result['Upload'].append(k1.to_map() if k1 else None)

        if self.upload_id_marker is not None:
            result['UploadIdMarker'] = self.upload_id_marker

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k1 in m.get('CommonPrefixes'):
                temp_model = main_models.CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k1))

        if m.get('Delimiter') is not None:
            self.delimiter = m.get('Delimiter')

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('KeyMarker') is not None:
            self.key_marker = m.get('KeyMarker')

        if m.get('MaxUploads') is not None:
            self.max_uploads = m.get('MaxUploads')

        if m.get('NextKeyMarker') is not None:
            self.next_key_marker = m.get('NextKeyMarker')

        if m.get('NextUploadIdMarker') is not None:
            self.next_upload_id_marker = m.get('NextUploadIdMarker')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        self.uploads = []
        if m.get('Upload') is not None:
            for k1 in m.get('Upload'):
                temp_model = main_models.Upload()
                self.uploads.append(temp_model.from_map(k1))

        if m.get('UploadIdMarker') is not None:
            self.upload_id_marker = m.get('UploadIdMarker')

        return self

