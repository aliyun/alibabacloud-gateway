// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectLegalHoldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectLegalHoldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectLegalHoldResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectLegalHoldResponseBody) *GetObjectLegalHoldResponse
	GetBody() *GetObjectLegalHoldResponseBody
}

type GetObjectLegalHoldResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectLegalHoldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectLegalHoldResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectLegalHoldResponse) GoString() string {
	return s.String()
}

func (s *GetObjectLegalHoldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectLegalHoldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectLegalHoldResponse) GetBody() *GetObjectLegalHoldResponseBody {
	return s.Body
}

func (s *GetObjectLegalHoldResponse) SetHeaders(v map[string]*string) *GetObjectLegalHoldResponse {
	s.Headers = v
	return s
}

func (s *GetObjectLegalHoldResponse) SetStatusCode(v int32) *GetObjectLegalHoldResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectLegalHoldResponse) SetBody(v *GetObjectLegalHoldResponseBody) *GetObjectLegalHoldResponse {
	s.Body = v
	return s
}

func (s *GetObjectLegalHoldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
