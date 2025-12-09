// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataLakeCachePrefetchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDataLakeCachePrefetchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDataLakeCachePrefetchJobResponse
	GetStatusCode() *int32
	SetBody(v *PutDataLakeCachePrefetchJobResponseBody) *PutDataLakeCachePrefetchJobResponse
	GetBody() *PutDataLakeCachePrefetchJobResponseBody
}

type PutDataLakeCachePrefetchJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDataLakeCachePrefetchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDataLakeCachePrefetchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeCachePrefetchJobResponse) GoString() string {
	return s.String()
}

func (s *PutDataLakeCachePrefetchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDataLakeCachePrefetchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDataLakeCachePrefetchJobResponse) GetBody() *PutDataLakeCachePrefetchJobResponseBody {
	return s.Body
}

func (s *PutDataLakeCachePrefetchJobResponse) SetHeaders(v map[string]*string) *PutDataLakeCachePrefetchJobResponse {
	s.Headers = v
	return s
}

func (s *PutDataLakeCachePrefetchJobResponse) SetStatusCode(v int32) *PutDataLakeCachePrefetchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDataLakeCachePrefetchJobResponse) SetBody(v *PutDataLakeCachePrefetchJobResponseBody) *PutDataLakeCachePrefetchJobResponse {
	s.Body = v
	return s
}

func (s *PutDataLakeCachePrefetchJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
