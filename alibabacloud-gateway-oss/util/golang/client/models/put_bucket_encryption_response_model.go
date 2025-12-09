// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketEncryptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketEncryptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketEncryptionResponse
	GetStatusCode() *int32
}

type PutBucketEncryptionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketEncryptionResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketEncryptionResponse) GoString() string {
	return s.String()
}

func (s *PutBucketEncryptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketEncryptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketEncryptionResponse) SetHeaders(v map[string]*string) *PutBucketEncryptionResponse {
	s.Headers = v
	return s
}

func (s *PutBucketEncryptionResponse) SetStatusCode(v int32) *PutBucketEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketEncryptionResponse) Validate() error {
	return dara.Validate(s)
}
