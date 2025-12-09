// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVectorsResponse
	GetStatusCode() *int32
	SetBody(v *QueryVectorsResponseBody) *QueryVectorsResponse
	GetBody() *QueryVectorsResponseBody
}

type QueryVectorsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVectorsResponse) GoString() string {
	return s.String()
}

func (s *QueryVectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVectorsResponse) GetBody() *QueryVectorsResponseBody {
	return s.Body
}

func (s *QueryVectorsResponse) SetHeaders(v map[string]*string) *QueryVectorsResponse {
	s.Headers = v
	return s
}

func (s *QueryVectorsResponse) SetStatusCode(v int32) *QueryVectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVectorsResponse) SetBody(v *QueryVectorsResponseBody) *QueryVectorsResponse {
	s.Body = v
	return s
}

func (s *QueryVectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
