// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketAclResponse
	GetStatusCode() *int32
}

type PutBucketAclResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketAclResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketAclResponse) GoString() string {
	return s.String()
}

func (s *PutBucketAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketAclResponse) SetHeaders(v map[string]*string) *PutBucketAclResponse {
	s.Headers = v
	return s
}

func (s *PutBucketAclResponse) SetStatusCode(v int32) *PutBucketAclResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketAclResponse) Validate() error {
	return dara.Validate(s)
}
