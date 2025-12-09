// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResourcePoolBucketGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketResourcePoolBucketGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketResourcePoolBucketGroupResponse
	GetStatusCode() *int32
}

type PutBucketResourcePoolBucketGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketResourcePoolBucketGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResourcePoolBucketGroupResponse) GoString() string {
	return s.String()
}

func (s *PutBucketResourcePoolBucketGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketResourcePoolBucketGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketResourcePoolBucketGroupResponse) SetHeaders(v map[string]*string) *PutBucketResourcePoolBucketGroupResponse {
	s.Headers = v
	return s
}

func (s *PutBucketResourcePoolBucketGroupResponse) SetStatusCode(v int32) *PutBucketResourcePoolBucketGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketResourcePoolBucketGroupResponse) Validate() error {
	return dara.Validate(s)
}
