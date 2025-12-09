// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketHttpsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketHttpsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketHttpsConfigResponse
	GetStatusCode() *int32
}

type PutBucketHttpsConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketHttpsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketHttpsConfigResponse) GoString() string {
	return s.String()
}

func (s *PutBucketHttpsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketHttpsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketHttpsConfigResponse) SetHeaders(v map[string]*string) *PutBucketHttpsConfigResponse {
	s.Headers = v
	return s
}

func (s *PutBucketHttpsConfigResponse) SetStatusCode(v int32) *PutBucketHttpsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketHttpsConfigResponse) Validate() error {
	return dara.Validate(s)
}
