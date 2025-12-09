// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostAsyncFetchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostAsyncFetchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostAsyncFetchTaskResponse
	GetStatusCode() *int32
	SetBody(v *PostAsyncFetchTaskResponseBody) *PostAsyncFetchTaskResponse
	GetBody() *PostAsyncFetchTaskResponseBody
}

type PostAsyncFetchTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostAsyncFetchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostAsyncFetchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PostAsyncFetchTaskResponse) GoString() string {
	return s.String()
}

func (s *PostAsyncFetchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostAsyncFetchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostAsyncFetchTaskResponse) GetBody() *PostAsyncFetchTaskResponseBody {
	return s.Body
}

func (s *PostAsyncFetchTaskResponse) SetHeaders(v map[string]*string) *PostAsyncFetchTaskResponse {
	s.Headers = v
	return s
}

func (s *PostAsyncFetchTaskResponse) SetStatusCode(v int32) *PostAsyncFetchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PostAsyncFetchTaskResponse) SetBody(v *PostAsyncFetchTaskResponseBody) *PostAsyncFetchTaskResponse {
	s.Body = v
	return s
}

func (s *PostAsyncFetchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
