// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBucketAntiDDosInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBucketAntiDDosInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBucketAntiDDosInfoResponse
	GetStatusCode() *int32
}

type UpdateBucketAntiDDosInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateBucketAntiDDosInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBucketAntiDDosInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateBucketAntiDDosInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBucketAntiDDosInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBucketAntiDDosInfoResponse) SetHeaders(v map[string]*string) *UpdateBucketAntiDDosInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateBucketAntiDDosInfoResponse) SetStatusCode(v int32) *UpdateBucketAntiDDosInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBucketAntiDDosInfoResponse) Validate() error {
	return dara.Validate(s)
}
