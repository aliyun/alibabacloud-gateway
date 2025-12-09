// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyObjectResponse
	GetStatusCode() *int32
	SetBody(v *CopyObjectResponseBody) *CopyObjectResponse
	GetBody() *CopyObjectResponseBody
}

type CopyObjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectResponse) GoString() string {
	return s.String()
}

func (s *CopyObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyObjectResponse) GetBody() *CopyObjectResponseBody {
	return s.Body
}

func (s *CopyObjectResponse) SetHeaders(v map[string]*string) *CopyObjectResponse {
	s.Headers = v
	return s
}

func (s *CopyObjectResponse) SetStatusCode(v int32) *CopyObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyObjectResponse) SetBody(v *CopyObjectResponseBody) *CopyObjectResponse {
	s.Body = v
	return s
}

func (s *CopyObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
