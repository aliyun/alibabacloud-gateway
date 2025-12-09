// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCachePrefetchJobHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeCachePrefetchJobHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeCachePrefetchJobHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeCachePrefetchJobHistoryResponseBody) *ListDataLakeCachePrefetchJobHistoryResponse
	GetBody() *ListDataLakeCachePrefetchJobHistoryResponseBody
}

type ListDataLakeCachePrefetchJobHistoryResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeCachePrefetchJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeCachePrefetchJobHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) GetBody() *ListDataLakeCachePrefetchJobHistoryResponseBody {
	return s.Body
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) SetHeaders(v map[string]*string) *ListDataLakeCachePrefetchJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) SetStatusCode(v int32) *ListDataLakeCachePrefetchJobHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) SetBody(v *ListDataLakeCachePrefetchJobHistoryResponseBody) *ListDataLakeCachePrefetchJobHistoryResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeCachePrefetchJobHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
