# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CompleteMultipartUploadRequest(DaraModel):
    def __init__(
        self,
        complete_multipart_upload: main_models.CompleteMultipartUpload = None,
        encoding_type: str = None,
        upload_id: str = None,
    ):
        # The container that stores the content of the CompleteMultipartUpload request.
        self.complete_multipart_upload = complete_multipart_upload
        # The encodingtype of the object name in the response. Only URL encoding is supported.
        # The object name can contain characters that are encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters, such as characters with an ASCII value from 0 to 10. You can configure this parameter to encode the object name in the response.
        self.encoding_type = encoding_type
        # The identifier of the multipart upload task.
        # 
        # This parameter is required.
        self.upload_id = upload_id

    def validate(self):
        if self.complete_multipart_upload:
            self.complete_multipart_upload.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.complete_multipart_upload is not None:
            result['CompleteMultipartUpload'] = self.complete_multipart_upload.to_map()

        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type

        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CompleteMultipartUpload') is not None:
            temp_model = main_models.CompleteMultipartUpload()
            self.complete_multipart_upload = temp_model.from_map(m.get('CompleteMultipartUpload'))

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self

