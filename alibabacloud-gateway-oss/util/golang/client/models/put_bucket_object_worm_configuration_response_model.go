// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketObjectWormConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketObjectWormConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketObjectWormConfigurationResponse
	GetStatusCode() *int32
}

type PutBucketObjectWormConfigurationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketObjectWormConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketObjectWormConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketObjectWormConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketObjectWormConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketObjectWormConfigurationResponse) SetHeaders(v map[string]*string) *PutBucketObjectWormConfigurationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketObjectWormConfigurationResponse) SetStatusCode(v int32) *PutBucketObjectWormConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketObjectWormConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
