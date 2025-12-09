// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVirtualBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutVirtualBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutVirtualBucketResponse
	GetStatusCode() *int32
}

type PutVirtualBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutVirtualBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s PutVirtualBucketResponse) GoString() string {
	return s.String()
}

func (s *PutVirtualBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutVirtualBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutVirtualBucketResponse) SetHeaders(v map[string]*string) *PutVirtualBucketResponse {
	s.Headers = v
	return s
}

func (s *PutVirtualBucketResponse) SetStatusCode(v int32) *PutVirtualBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *PutVirtualBucketResponse) Validate() error {
	return dara.Validate(s)
}
