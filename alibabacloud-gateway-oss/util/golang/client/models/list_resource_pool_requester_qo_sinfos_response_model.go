// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolRequesterQoSInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePoolRequesterQoSInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePoolRequesterQoSInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePoolRequesterQoSInfosResponseBody) *ListResourcePoolRequesterQoSInfosResponse
	GetBody() *ListResourcePoolRequesterQoSInfosResponseBody
}

type ListResourcePoolRequesterQoSInfosResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePoolRequesterQoSInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePoolRequesterQoSInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolRequesterQoSInfosResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePoolRequesterQoSInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePoolRequesterQoSInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePoolRequesterQoSInfosResponse) GetBody() *ListResourcePoolRequesterQoSInfosResponseBody {
	return s.Body
}

func (s *ListResourcePoolRequesterQoSInfosResponse) SetHeaders(v map[string]*string) *ListResourcePoolRequesterQoSInfosResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResponse) SetStatusCode(v int32) *ListResourcePoolRequesterQoSInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResponse) SetBody(v *ListResourcePoolRequesterQoSInfosResponseBody) *ListResourcePoolRequesterQoSInfosResponse {
	s.Body = v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
