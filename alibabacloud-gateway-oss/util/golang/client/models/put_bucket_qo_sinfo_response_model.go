// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketQoSInfoResponse
	GetStatusCode() *int32
}

type PutBucketQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *PutBucketQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketQoSInfoResponse) SetHeaders(v map[string]*string) *PutBucketQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *PutBucketQoSInfoResponse) SetStatusCode(v int32) *PutBucketQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
