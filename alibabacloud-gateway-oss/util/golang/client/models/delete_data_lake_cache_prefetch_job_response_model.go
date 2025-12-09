// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeCachePrefetchJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLakeCachePrefetchJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLakeCachePrefetchJobResponse
	GetStatusCode() *int32
}

type DeleteDataLakeCachePrefetchJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDataLakeCachePrefetchJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeCachePrefetchJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeCachePrefetchJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLakeCachePrefetchJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLakeCachePrefetchJobResponse) SetHeaders(v map[string]*string) *DeleteDataLakeCachePrefetchJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLakeCachePrefetchJobResponse) SetStatusCode(v int32) *DeleteDataLakeCachePrefetchJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLakeCachePrefetchJobResponse) Validate() error {
	return dara.Validate(s)
}
