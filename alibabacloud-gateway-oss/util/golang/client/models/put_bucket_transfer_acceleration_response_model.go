// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketTransferAccelerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketTransferAccelerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketTransferAccelerationResponse
	GetStatusCode() *int32
}

type PutBucketTransferAccelerationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketTransferAccelerationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketTransferAccelerationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketTransferAccelerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketTransferAccelerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketTransferAccelerationResponse) SetHeaders(v map[string]*string) *PutBucketTransferAccelerationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketTransferAccelerationResponse) SetStatusCode(v int32) *PutBucketTransferAccelerationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketTransferAccelerationResponse) Validate() error {
	return dara.Validate(s)
}
