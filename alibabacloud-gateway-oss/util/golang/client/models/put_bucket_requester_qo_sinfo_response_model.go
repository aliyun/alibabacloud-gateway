// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRequesterQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketRequesterQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketRequesterQoSInfoResponse
	GetStatusCode() *int32
}

type PutBucketRequesterQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketRequesterQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRequesterQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRequesterQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketRequesterQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketRequesterQoSInfoResponse) SetHeaders(v map[string]*string) *PutBucketRequesterQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRequesterQoSInfoResponse) SetStatusCode(v int32) *PutBucketRequesterQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketRequesterQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
