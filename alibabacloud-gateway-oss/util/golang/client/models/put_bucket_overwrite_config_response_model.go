// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketOverwriteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketOverwriteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketOverwriteConfigResponse
	GetStatusCode() *int32
}

type PutBucketOverwriteConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketOverwriteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketOverwriteConfigResponse) GoString() string {
	return s.String()
}

func (s *PutBucketOverwriteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketOverwriteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketOverwriteConfigResponse) SetHeaders(v map[string]*string) *PutBucketOverwriteConfigResponse {
	s.Headers = v
	return s
}

func (s *PutBucketOverwriteConfigResponse) SetStatusCode(v int32) *PutBucketOverwriteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketOverwriteConfigResponse) Validate() error {
	return dara.Validate(s)
}
