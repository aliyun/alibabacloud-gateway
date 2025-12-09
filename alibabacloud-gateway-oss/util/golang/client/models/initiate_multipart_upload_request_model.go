// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateMultipartUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodingType(v string) *InitiateMultipartUploadRequest
	GetEncodingType() *string
}

type InitiateMultipartUploadRequest struct {
	// The method used to encode the object name in the response. Only URL encoding is supported. The object name can contain characters encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse specific control characters, such as characters whose ASCII values range from 0 to 10. You can configure the encoding-type parameter to encode object names that include characters that cannot be parsed by XML 1.0 in the response.
	//
	// <br>Default value: null
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
}

func (s InitiateMultipartUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiateMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *InitiateMultipartUploadRequest) SetEncodingType(v string) *InitiateMultipartUploadRequest {
	s.EncodingType = &v
	return s
}

func (s *InitiateMultipartUploadRequest) Validate() error {
	return dara.Validate(s)
}
