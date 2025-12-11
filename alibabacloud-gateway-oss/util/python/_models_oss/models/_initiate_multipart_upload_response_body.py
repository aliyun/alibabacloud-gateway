# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class InitiateMultipartUploadResponseBody(DaraModel):
    def __init__(
        self,
        initiate_multipart_upload_result: main_models.InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult = None,
    ):
        # The container that stores the results of the InitiateMultipartUpload request.
        self.initiate_multipart_upload_result = initiate_multipart_upload_result

    def validate(self):
        if self.initiate_multipart_upload_result:
            self.initiate_multipart_upload_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.initiate_multipart_upload_result is not None:
            result['InitiateMultipartUploadResult'] = self.initiate_multipart_upload_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InitiateMultipartUploadResult') is not None:
            temp_model = main_models.InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult()
            self.initiate_multipart_upload_result = temp_model.from_map(m.get('InitiateMultipartUploadResult'))

        return self

class InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        encoding_type: str = None,
        key: str = None,
        upload_id: str = None,
    ):
        # The name of the bucket to which the object is uploaded by the multipart upload task.
        self.bucket = bucket
        # The encoding type of the object name in the response. If the encoding-type parameter is specified in the request, the object name in the response is encoded.
        self.encoding_type = encoding_type
        # The name of the object that is uploaded by the multipart upload task.
        self.key = key
        # The upload ID that uniquely identifies the multipart upload task. The upload ID is used to call UploadPart and CompleteMultipartUpload later.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        if self.key is not None:
            result['Key'] = self.key

        if self.upload_id is not None:
            result['UploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')

        return self

