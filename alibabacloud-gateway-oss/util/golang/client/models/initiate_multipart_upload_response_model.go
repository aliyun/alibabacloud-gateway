// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateMultipartUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitiateMultipartUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitiateMultipartUploadResponse
	GetStatusCode() *int32
	SetBody(v *InitiateMultipartUploadResponseBody) *InitiateMultipartUploadResponse
	GetBody() *InitiateMultipartUploadResponseBody
}

type InitiateMultipartUploadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitiateMultipartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitiateMultipartUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s InitiateMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitiateMultipartUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitiateMultipartUploadResponse) GetBody() *InitiateMultipartUploadResponseBody {
	return s.Body
}

func (s *InitiateMultipartUploadResponse) SetHeaders(v map[string]*string) *InitiateMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *InitiateMultipartUploadResponse) SetStatusCode(v int32) *InitiateMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *InitiateMultipartUploadResponse) SetBody(v *InitiateMultipartUploadResponseBody) *InitiateMultipartUploadResponse {
	s.Body = v
	return s
}

func (s *InitiateMultipartUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
