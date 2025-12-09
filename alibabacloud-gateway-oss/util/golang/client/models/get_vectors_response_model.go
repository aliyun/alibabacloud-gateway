// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVectorsResponse
	GetStatusCode() *int32
	SetBody(v *GetVectorsResponseBody) *GetVectorsResponse
	GetBody() *GetVectorsResponseBody
}

type GetVectorsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVectorsResponse) GoString() string {
	return s.String()
}

func (s *GetVectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVectorsResponse) GetBody() *GetVectorsResponseBody {
	return s.Body
}

func (s *GetVectorsResponse) SetHeaders(v map[string]*string) *GetVectorsResponse {
	s.Headers = v
	return s
}

func (s *GetVectorsResponse) SetStatusCode(v int32) *GetVectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVectorsResponse) SetBody(v *GetVectorsResponseBody) *GetVectorsResponse {
	s.Body = v
	return s
}

func (s *GetVectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
