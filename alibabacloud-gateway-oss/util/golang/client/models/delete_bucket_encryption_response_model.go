// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketEncryptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketEncryptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketEncryptionResponse
	GetStatusCode() *int32
}

type DeleteBucketEncryptionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketEncryptionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketEncryptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketEncryptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketEncryptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketEncryptionResponse) SetHeaders(v map[string]*string) *DeleteBucketEncryptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketEncryptionResponse) SetStatusCode(v int32) *DeleteBucketEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketEncryptionResponse) Validate() error {
	return dara.Validate(s)
}
