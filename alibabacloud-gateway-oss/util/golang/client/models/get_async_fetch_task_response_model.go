// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncFetchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncFetchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncFetchTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncFetchTaskResponseBody) *GetAsyncFetchTaskResponse
	GetBody() *GetAsyncFetchTaskResponseBody
}

type GetAsyncFetchTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncFetchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncFetchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncFetchTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncFetchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncFetchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncFetchTaskResponse) GetBody() *GetAsyncFetchTaskResponseBody {
	return s.Body
}

func (s *GetAsyncFetchTaskResponse) SetHeaders(v map[string]*string) *GetAsyncFetchTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncFetchTaskResponse) SetStatusCode(v int32) *GetAsyncFetchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncFetchTaskResponse) SetBody(v *GetAsyncFetchTaskResponseBody) *GetAsyncFetchTaskResponse {
	s.Body = v
	return s
}

func (s *GetAsyncFetchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
