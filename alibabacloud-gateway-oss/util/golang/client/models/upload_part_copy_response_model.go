// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartCopyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadPartCopyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadPartCopyResponse
	GetStatusCode() *int32
	SetBody(v *UploadPartCopyResponseBody) *UploadPartCopyResponse
	GetBody() *UploadPartCopyResponseBody
}

type UploadPartCopyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadPartCopyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadPartCopyResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadPartCopyResponse) GoString() string {
	return s.String()
}

func (s *UploadPartCopyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadPartCopyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadPartCopyResponse) GetBody() *UploadPartCopyResponseBody {
	return s.Body
}

func (s *UploadPartCopyResponse) SetHeaders(v map[string]*string) *UploadPartCopyResponse {
	s.Headers = v
	return s
}

func (s *UploadPartCopyResponse) SetStatusCode(v int32) *UploadPartCopyResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPartCopyResponse) SetBody(v *UploadPartCopyResponseBody) *UploadPartCopyResponse {
	s.Body = v
	return s
}

func (s *UploadPartCopyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
