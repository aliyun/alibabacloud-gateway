// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteMultipartUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompleteMultipartUpload(v *CompleteMultipartUpload) *CompleteMultipartUploadRequest
	GetCompleteMultipartUpload() *CompleteMultipartUpload
	SetEncodingType(v string) *CompleteMultipartUploadRequest
	GetEncodingType() *string
	SetUploadId(v string) *CompleteMultipartUploadRequest
	GetUploadId() *string
}

type CompleteMultipartUploadRequest struct {
	// The container that stores the content of the CompleteMultipartUpload request.
	CompleteMultipartUpload *CompleteMultipartUpload `json:"CompleteMultipartUpload,omitempty" xml:"CompleteMultipartUpload,omitempty"`
	// The encodingtype of the object name in the response. Only URL encoding is supported.
	//
	// The object name can contain characters that are encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters, such as characters with an ASCII value from 0 to 10. You can configure this parameter to encode the object name in the response.
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	// The identifier of the multipart upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0004B9895DBBB6E****
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s CompleteMultipartUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadRequest) GetCompleteMultipartUpload() *CompleteMultipartUpload {
	return s.CompleteMultipartUpload
}

func (s *CompleteMultipartUploadRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *CompleteMultipartUploadRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CompleteMultipartUploadRequest) SetCompleteMultipartUpload(v *CompleteMultipartUpload) *CompleteMultipartUploadRequest {
	s.CompleteMultipartUpload = v
	return s
}

func (s *CompleteMultipartUploadRequest) SetEncodingType(v string) *CompleteMultipartUploadRequest {
	s.EncodingType = &v
	return s
}

func (s *CompleteMultipartUploadRequest) SetUploadId(v string) *CompleteMultipartUploadRequest {
	s.UploadId = &v
	return s
}

func (s *CompleteMultipartUploadRequest) Validate() error {
	if s.CompleteMultipartUpload != nil {
		if err := s.CompleteMultipartUpload.Validate(); err != nil {
			return err
		}
	}
	return nil
}
