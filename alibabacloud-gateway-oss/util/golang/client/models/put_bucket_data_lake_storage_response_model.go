// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketDataLakeStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketDataLakeStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketDataLakeStorageResponse
	GetStatusCode() *int32
}

type PutBucketDataLakeStorageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketDataLakeStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketDataLakeStorageResponse) GoString() string {
	return s.String()
}

func (s *PutBucketDataLakeStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketDataLakeStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketDataLakeStorageResponse) SetHeaders(v map[string]*string) *PutBucketDataLakeStorageResponse {
	s.Headers = v
	return s
}

func (s *PutBucketDataLakeStorageResponse) SetStatusCode(v int32) *PutBucketDataLakeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketDataLakeStorageResponse) Validate() error {
	return dara.Validate(s)
}
