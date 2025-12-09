// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketResourceGroupResponse
	GetStatusCode() *int32
}

type PutBucketResourceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *PutBucketResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketResourceGroupResponse) SetHeaders(v map[string]*string) *PutBucketResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *PutBucketResourceGroupResponse) SetStatusCode(v int32) *PutBucketResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
