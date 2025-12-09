// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketReplicationResponse
	GetStatusCode() *int32
}

type PutBucketReplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketReplicationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketReplicationResponse) SetHeaders(v map[string]*string) *PutBucketReplicationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketReplicationResponse) SetStatusCode(v int32) *PutBucketReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketReplicationResponse) Validate() error {
	return dara.Validate(s)
}
