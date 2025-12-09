// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketOverwriteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketOverwriteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketOverwriteConfigResponse
	GetStatusCode() *int32
}

type DeleteBucketOverwriteConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketOverwriteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketOverwriteConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketOverwriteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketOverwriteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketOverwriteConfigResponse) SetHeaders(v map[string]*string) *DeleteBucketOverwriteConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketOverwriteConfigResponse) SetStatusCode(v int32) *DeleteBucketOverwriteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketOverwriteConfigResponse) Validate() error {
	return dara.Validate(s)
}
