// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeCachePrefetchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakeCachePrefetchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakeCachePrefetchJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakeCachePrefetchJobResponseBody) *GetDataLakeCachePrefetchJobResponse
	GetBody() *GetDataLakeCachePrefetchJobResponseBody
}

type GetDataLakeCachePrefetchJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakeCachePrefetchJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakeCachePrefetchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeCachePrefetchJobResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakeCachePrefetchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakeCachePrefetchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakeCachePrefetchJobResponse) GetBody() *GetDataLakeCachePrefetchJobResponseBody {
	return s.Body
}

func (s *GetDataLakeCachePrefetchJobResponse) SetHeaders(v map[string]*string) *GetDataLakeCachePrefetchJobResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakeCachePrefetchJobResponse) SetStatusCode(v int32) *GetDataLakeCachePrefetchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakeCachePrefetchJobResponse) SetBody(v *GetDataLakeCachePrefetchJobResponseBody) *GetDataLakeCachePrefetchJobResponse {
	s.Body = v
	return s
}

func (s *GetDataLakeCachePrefetchJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
