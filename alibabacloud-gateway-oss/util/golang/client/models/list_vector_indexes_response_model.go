// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorIndexesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVectorIndexesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVectorIndexesResponse
	GetStatusCode() *int32
	SetBody(v *ListVectorIndexesResponseBody) *ListVectorIndexesResponse
	GetBody() *ListVectorIndexesResponseBody
}

type ListVectorIndexesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVectorIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVectorIndexesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVectorIndexesResponse) GoString() string {
	return s.String()
}

func (s *ListVectorIndexesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVectorIndexesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVectorIndexesResponse) GetBody() *ListVectorIndexesResponseBody {
	return s.Body
}

func (s *ListVectorIndexesResponse) SetHeaders(v map[string]*string) *ListVectorIndexesResponse {
	s.Headers = v
	return s
}

func (s *ListVectorIndexesResponse) SetStatusCode(v int32) *ListVectorIndexesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVectorIndexesResponse) SetBody(v *ListVectorIndexesResponseBody) *ListVectorIndexesResponse {
	s.Body = v
	return s
}

func (s *ListVectorIndexesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
