// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVectorBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutVectorBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutVectorBucketResponse
	GetStatusCode() *int32
}

type PutVectorBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutVectorBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s PutVectorBucketResponse) GoString() string {
	return s.String()
}

func (s *PutVectorBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutVectorBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutVectorBucketResponse) SetHeaders(v map[string]*string) *PutVectorBucketResponse {
	s.Headers = v
	return s
}

func (s *PutVectorBucketResponse) SetStatusCode(v int32) *PutVectorBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *PutVectorBucketResponse) Validate() error {
	return dara.Validate(s)
}
