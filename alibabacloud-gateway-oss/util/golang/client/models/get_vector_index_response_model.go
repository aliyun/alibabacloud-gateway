// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVectorIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVectorIndexResponse
	GetStatusCode() *int32
	SetBody(v *GetVectorIndexResponseBody) *GetVectorIndexResponse
	GetBody() *GetVectorIndexResponseBody
}

type GetVectorIndexResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVectorIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVectorIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVectorIndexResponse) GoString() string {
	return s.String()
}

func (s *GetVectorIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVectorIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVectorIndexResponse) GetBody() *GetVectorIndexResponseBody {
	return s.Body
}

func (s *GetVectorIndexResponse) SetHeaders(v map[string]*string) *GetVectorIndexResponse {
	s.Headers = v
	return s
}

func (s *GetVectorIndexResponse) SetStatusCode(v int32) *GetVectorIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVectorIndexResponse) SetBody(v *GetVectorIndexResponseBody) *GetVectorIndexResponse {
	s.Body = v
	return s
}

func (s *GetVectorIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
