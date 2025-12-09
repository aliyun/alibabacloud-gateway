// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCommonHeaderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v *CommonHeaders) *PutBucketCommonHeaderRequest
	GetCommonHeaders() *CommonHeaders
}

type PutBucketCommonHeaderRequest struct {
	// User-defined response headers configuration
	CommonHeaders *CommonHeaders `json:"CommonHeaders,omitempty" xml:"CommonHeaders,omitempty"`
}

func (s PutBucketCommonHeaderRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCommonHeaderRequest) GoString() string {
	return s.String()
}

func (s *PutBucketCommonHeaderRequest) GetCommonHeaders() *CommonHeaders {
	return s.CommonHeaders
}

func (s *PutBucketCommonHeaderRequest) SetCommonHeaders(v *CommonHeaders) *PutBucketCommonHeaderRequest {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketCommonHeaderRequest) Validate() error {
	if s.CommonHeaders != nil {
		if err := s.CommonHeaders.Validate(); err != nil {
			return err
		}
	}
	return nil
}
