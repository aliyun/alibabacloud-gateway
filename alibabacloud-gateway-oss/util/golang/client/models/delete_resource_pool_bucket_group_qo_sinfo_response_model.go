// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourcePoolBucketGroupQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourcePoolBucketGroupQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourcePoolBucketGroupQoSInfoResponse
	GetStatusCode() *int32
}

type DeleteResourcePoolBucketGroupQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteResourcePoolBucketGroupQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourcePoolBucketGroupQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourcePoolBucketGroupQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourcePoolBucketGroupQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourcePoolBucketGroupQoSInfoResponse) SetHeaders(v map[string]*string) *DeleteResourcePoolBucketGroupQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourcePoolBucketGroupQoSInfoResponse) SetStatusCode(v int32) *DeleteResourcePoolBucketGroupQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourcePoolBucketGroupQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
