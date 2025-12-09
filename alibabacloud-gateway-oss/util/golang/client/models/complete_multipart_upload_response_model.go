// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteMultipartUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteMultipartUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteMultipartUploadResponse
	GetStatusCode() *int32
	SetBody(v *CompleteMultipartUploadResponseBody) *CompleteMultipartUploadResponse
	GetBody() *CompleteMultipartUploadResponseBody
}

type CompleteMultipartUploadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteMultipartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteMultipartUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteMultipartUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteMultipartUploadResponse) GetBody() *CompleteMultipartUploadResponseBody {
	return s.Body
}

func (s *CompleteMultipartUploadResponse) SetHeaders(v map[string]*string) *CompleteMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *CompleteMultipartUploadResponse) SetStatusCode(v int32) *CompleteMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteMultipartUploadResponse) SetBody(v *CompleteMultipartUploadResponseBody) *CompleteMultipartUploadResponse {
	s.Body = v
	return s
}

func (s *CompleteMultipartUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
