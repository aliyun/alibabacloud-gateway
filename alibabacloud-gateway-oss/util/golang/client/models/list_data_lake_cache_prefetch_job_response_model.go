// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCachePrefetchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeCachePrefetchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeCachePrefetchJobResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeCachePrefetchJobResponseBody) *ListDataLakeCachePrefetchJobResponse
	GetBody() *ListDataLakeCachePrefetchJobResponseBody
}

type ListDataLakeCachePrefetchJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeCachePrefetchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeCachePrefetchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeCachePrefetchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeCachePrefetchJobResponse) GetBody() *ListDataLakeCachePrefetchJobResponseBody {
	return s.Body
}

func (s *ListDataLakeCachePrefetchJobResponse) SetHeaders(v map[string]*string) *ListDataLakeCachePrefetchJobResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponse) SetStatusCode(v int32) *ListDataLakeCachePrefetchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponse) SetBody(v *ListDataLakeCachePrefetchJobResponseBody) *ListDataLakeCachePrefetchJobResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
