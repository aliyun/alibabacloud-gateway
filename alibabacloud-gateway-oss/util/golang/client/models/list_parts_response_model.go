// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPartsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPartsResponse
	GetStatusCode() *int32
	SetBody(v *ListPartsResponseBody) *ListPartsResponse
	GetBody() *ListPartsResponseBody
}

type ListPartsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPartsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPartsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPartsResponse) GoString() string {
	return s.String()
}

func (s *ListPartsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPartsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPartsResponse) GetBody() *ListPartsResponseBody {
	return s.Body
}

func (s *ListPartsResponse) SetHeaders(v map[string]*string) *ListPartsResponse {
	s.Headers = v
	return s
}

func (s *ListPartsResponse) SetStatusCode(v int32) *ListPartsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartsResponse) SetBody(v *ListPartsResponseBody) *ListPartsResponse {
	s.Body = v
	return s
}

func (s *ListPartsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
