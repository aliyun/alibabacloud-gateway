# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CompleteMultipartUploadResponseBody(DaraModel):
    def __init__(
        self,
        complete_multipart_upload_result: main_models.CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult = None,
    ):
        # The container that stores the results of the CompleteMultipartUpload request.
        self.complete_multipart_upload_result = complete_multipart_upload_result

    def validate(self):
        if self.complete_multipart_upload_result:
            self.complete_multipart_upload_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.complete_multipart_upload_result is not None:
            result['CompleteMultipartUploadResult'] = self.complete_multipart_upload_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CompleteMultipartUploadResult') is not None:
            temp_model = main_models.CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult()
            self.complete_multipart_upload_result = temp_model.from_map(m.get('CompleteMultipartUploadResult'))

        return self

class CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        etag: str = None,
        encoding_type: str = None,
        key: str = None,
        location: str = None,
    ):
        # The name of the bucket that contains the object you want to restore.
        self.bucket = bucket
        # The ETag that is generated when an object is created. ETags are used to identify the content of objects.
        # 
        # If an object is created by calling the CompleteMultipartUpload operation, the ETag value is not the MD5 hash of the object content but a unique value calculated based on a specific rule.
        # 
        # > The ETag of an object can be used to check whether the object content is modified. However, we recommend that you use the MD5 hash of an object rather than the ETag value of the object to verify data integrity.
        self.etag = etag
        # The encoding type of the object name in the response. If this parameter is specified in the request, the object name is encoded in the response.
        self.encoding_type = encoding_type
        # The name of the uploaded object.
        self.key = key
        # The URL that is used to access the uploaded object.
        self.location = location

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.etag is not None:
            result['ETag'] = self.etag

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        if self.key is not None:
            result['Key'] = self.key

        if self.location is not None:
            result['Location'] = self.location

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('Location') is not None:
            self.location = m.get('Location')

        return self

