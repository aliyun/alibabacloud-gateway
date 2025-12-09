// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadPartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadPartResponse
	GetStatusCode() *int32
}

type UploadPartResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UploadPartResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadPartResponse) GoString() string {
	return s.String()
}

func (s *UploadPartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadPartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadPartResponse) SetHeaders(v map[string]*string) *UploadPartResponse {
	s.Headers = v
	return s
}

func (s *UploadPartResponse) SetStatusCode(v int32) *UploadPartResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPartResponse) Validate() error {
	return dara.Validate(s)
}
