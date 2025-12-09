// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketDataAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketDataAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketDataAcceleratorResponse
	GetStatusCode() *int32
}

type DeleteBucketDataAcceleratorResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketDataAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketDataAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketDataAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketDataAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketDataAcceleratorResponse) SetHeaders(v map[string]*string) *DeleteBucketDataAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketDataAcceleratorResponse) SetStatusCode(v int32) *DeleteBucketDataAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketDataAcceleratorResponse) Validate() error {
	return dara.Validate(s)
}
