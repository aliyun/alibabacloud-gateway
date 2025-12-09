// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketReplicationResponse
	GetStatusCode() *int32
}

type DeleteBucketReplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketReplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketReplicationResponse) SetHeaders(v map[string]*string) *DeleteBucketReplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketReplicationResponse) SetStatusCode(v int32) *DeleteBucketReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketReplicationResponse) Validate() error {
	return dara.Validate(s)
}
