// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPartUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPartUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPartUploadResponse
	GetStatusCode() *int32
	SetBody(v *StartPartUploadResponseBody) *StartPartUploadResponse
	GetBody() *StartPartUploadResponseBody
}

type StartPartUploadResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPartUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPartUploadResponse) GoString() string {
	return s.String()
}

func (s *StartPartUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPartUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPartUploadResponse) GetBody() *StartPartUploadResponseBody {
	return s.Body
}

func (s *StartPartUploadResponse) SetHeaders(v map[string]*string) *StartPartUploadResponse {
	s.Headers = v
	return s
}

func (s *StartPartUploadResponse) SetStatusCode(v int32) *StartPartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPartUploadResponse) SetBody(v *StartPartUploadResponseBody) *StartPartUploadResponse {
	s.Body = v
	return s
}

func (s *StartPartUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
