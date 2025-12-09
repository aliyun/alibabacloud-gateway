// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortMultipartUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortMultipartUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortMultipartUploadResponse
	GetStatusCode() *int32
}

type AbortMultipartUploadResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AbortMultipartUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *AbortMultipartUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortMultipartUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortMultipartUploadResponse) SetHeaders(v map[string]*string) *AbortMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *AbortMultipartUploadResponse) SetStatusCode(v int32) *AbortMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortMultipartUploadResponse) Validate() error {
	return dara.Validate(s)
}
