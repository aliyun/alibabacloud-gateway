// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketDataAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketDataAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketDataAcceleratorResponse
	GetStatusCode() *int32
}

type PutBucketDataAcceleratorResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketDataAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketDataAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *PutBucketDataAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketDataAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketDataAcceleratorResponse) SetHeaders(v map[string]*string) *PutBucketDataAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *PutBucketDataAcceleratorResponse) SetStatusCode(v int32) *PutBucketDataAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketDataAcceleratorResponse) Validate() error {
	return dara.Validate(s)
}
