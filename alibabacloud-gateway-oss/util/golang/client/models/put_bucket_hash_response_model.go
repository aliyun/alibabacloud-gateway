// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketHashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketHashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketHashResponse
	GetStatusCode() *int32
}

type PutBucketHashResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketHashResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketHashResponse) GoString() string {
	return s.String()
}

func (s *PutBucketHashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketHashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketHashResponse) SetHeaders(v map[string]*string) *PutBucketHashResponse {
	s.Headers = v
	return s
}

func (s *PutBucketHashResponse) SetStatusCode(v int32) *PutBucketHashResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketHashResponse) Validate() error {
	return dara.Validate(s)
}
