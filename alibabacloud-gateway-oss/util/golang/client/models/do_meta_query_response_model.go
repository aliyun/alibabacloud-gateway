// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoMetaQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DoMetaQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DoMetaQueryResponse
	GetStatusCode() *int32
	SetBody(v *DoMetaQueryResponseBody) *DoMetaQueryResponse
	GetBody() *DoMetaQueryResponseBody
}

type DoMetaQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DoMetaQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DoMetaQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s DoMetaQueryResponse) GoString() string {
	return s.String()
}

func (s *DoMetaQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DoMetaQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DoMetaQueryResponse) GetBody() *DoMetaQueryResponseBody {
	return s.Body
}

func (s *DoMetaQueryResponse) SetHeaders(v map[string]*string) *DoMetaQueryResponse {
	s.Headers = v
	return s
}

func (s *DoMetaQueryResponse) SetStatusCode(v int32) *DoMetaQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DoMetaQueryResponse) SetBody(v *DoMetaQueryResponseBody) *DoMetaQueryResponse {
	s.Body = v
	return s
}

func (s *DoMetaQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
