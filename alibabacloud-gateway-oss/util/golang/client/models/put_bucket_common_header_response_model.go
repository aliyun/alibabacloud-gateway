// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCommonHeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketCommonHeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketCommonHeaderResponse
	GetStatusCode() *int32
}

type PutBucketCommonHeaderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketCommonHeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCommonHeaderResponse) GoString() string {
	return s.String()
}

func (s *PutBucketCommonHeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketCommonHeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketCommonHeaderResponse) SetHeaders(v map[string]*string) *PutBucketCommonHeaderResponse {
	s.Headers = v
	return s
}

func (s *PutBucketCommonHeaderResponse) SetStatusCode(v int32) *PutBucketCommonHeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketCommonHeaderResponse) Validate() error {
	return dara.Validate(s)
}
