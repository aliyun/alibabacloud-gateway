// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipleObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultipleObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultipleObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultipleObjectsResponseBody) *DeleteMultipleObjectsResponse
	GetBody() *DeleteMultipleObjectsResponseBody
}

type DeleteMultipleObjectsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultipleObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultipleObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipleObjectsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultipleObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultipleObjectsResponse) GetBody() *DeleteMultipleObjectsResponseBody {
	return s.Body
}

func (s *DeleteMultipleObjectsResponse) SetHeaders(v map[string]*string) *DeleteMultipleObjectsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultipleObjectsResponse) SetStatusCode(v int32) *DeleteMultipleObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultipleObjectsResponse) SetBody(v *DeleteMultipleObjectsResponseBody) *DeleteMultipleObjectsResponse {
	s.Body = v
	return s
}

func (s *DeleteMultipleObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
