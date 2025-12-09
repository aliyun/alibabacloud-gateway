// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataLakeCachePrefetchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDataLakeCachePrefetchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDataLakeCachePrefetchJobResponse
	GetStatusCode() *int32
}

type StopDataLakeCachePrefetchJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopDataLakeCachePrefetchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDataLakeCachePrefetchJobResponse) GoString() string {
	return s.String()
}

func (s *StopDataLakeCachePrefetchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDataLakeCachePrefetchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDataLakeCachePrefetchJobResponse) SetHeaders(v map[string]*string) *StopDataLakeCachePrefetchJobResponse {
	s.Headers = v
	return s
}

func (s *StopDataLakeCachePrefetchJobResponse) SetStatusCode(v int32) *StopDataLakeCachePrefetchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataLakeCachePrefetchJobResponse) Validate() error {
	return dara.Validate(s)
}
