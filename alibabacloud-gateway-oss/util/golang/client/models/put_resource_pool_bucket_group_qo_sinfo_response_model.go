// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourcePoolBucketGroupQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutResourcePoolBucketGroupQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutResourcePoolBucketGroupQoSInfoResponse
	GetStatusCode() *int32
}

type PutResourcePoolBucketGroupQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutResourcePoolBucketGroupQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PutResourcePoolBucketGroupQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *PutResourcePoolBucketGroupQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutResourcePoolBucketGroupQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutResourcePoolBucketGroupQoSInfoResponse) SetHeaders(v map[string]*string) *PutResourcePoolBucketGroupQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *PutResourcePoolBucketGroupQoSInfoResponse) SetStatusCode(v int32) *PutResourcePoolBucketGroupQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PutResourcePoolBucketGroupQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
