// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketTransferAccelerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketTransferAccelerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketTransferAccelerationResponse
	GetStatusCode() *int32
}

type DeleteBucketTransferAccelerationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketTransferAccelerationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketTransferAccelerationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketTransferAccelerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketTransferAccelerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketTransferAccelerationResponse) SetHeaders(v map[string]*string) *DeleteBucketTransferAccelerationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketTransferAccelerationResponse) SetStatusCode(v int32) *DeleteBucketTransferAccelerationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketTransferAccelerationResponse) Validate() error {
	return dara.Validate(s)
}
