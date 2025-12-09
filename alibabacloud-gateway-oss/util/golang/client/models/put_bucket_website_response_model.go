// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketWebsiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketWebsiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketWebsiteResponse
	GetStatusCode() *int32
}

type PutBucketWebsiteResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketWebsiteResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketWebsiteResponse) GoString() string {
	return s.String()
}

func (s *PutBucketWebsiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketWebsiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketWebsiteResponse) SetHeaders(v map[string]*string) *PutBucketWebsiteResponse {
	s.Headers = v
	return s
}

func (s *PutBucketWebsiteResponse) SetStatusCode(v int32) *PutBucketWebsiteResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketWebsiteResponse) Validate() error {
	return dara.Validate(s)
}
