// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataLakeCachePrefetchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDataLakeCachePrefetchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDataLakeCachePrefetchJobResponse
	GetStatusCode() *int32
}

type StartDataLakeCachePrefetchJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StartDataLakeCachePrefetchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDataLakeCachePrefetchJobResponse) GoString() string {
	return s.String()
}

func (s *StartDataLakeCachePrefetchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDataLakeCachePrefetchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDataLakeCachePrefetchJobResponse) SetHeaders(v map[string]*string) *StartDataLakeCachePrefetchJobResponse {
	s.Headers = v
	return s
}

func (s *StartDataLakeCachePrefetchJobResponse) SetStatusCode(v int32) *StartDataLakeCachePrefetchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDataLakeCachePrefetchJobResponse) Validate() error {
	return dara.Validate(s)
}
