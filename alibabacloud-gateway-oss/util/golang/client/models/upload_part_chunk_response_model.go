// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadPartChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadPartChunkResponse
	GetStatusCode() *int32
}

type UploadPartChunkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UploadPartChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadPartChunkResponse) GoString() string {
	return s.String()
}

func (s *UploadPartChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadPartChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadPartChunkResponse) SetHeaders(v map[string]*string) *UploadPartChunkResponse {
	s.Headers = v
	return s
}

func (s *UploadPartChunkResponse) SetStatusCode(v int32) *UploadPartChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPartChunkResponse) Validate() error {
	return dara.Validate(s)
}
