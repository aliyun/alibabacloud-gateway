// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCommonHeaderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v *CommonHeaders) *GetBucketCommonHeaderResponseBody
	GetCommonHeaders() *CommonHeaders
}

type GetBucketCommonHeaderResponseBody struct {
	CommonHeaders *CommonHeaders `json:"CommonHeaders,omitempty" xml:"CommonHeaders,omitempty"`
}

func (s GetBucketCommonHeaderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCommonHeaderResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketCommonHeaderResponseBody) GetCommonHeaders() *CommonHeaders {
	return s.CommonHeaders
}

func (s *GetBucketCommonHeaderResponseBody) SetCommonHeaders(v *CommonHeaders) *GetBucketCommonHeaderResponseBody {
	s.CommonHeaders = v
	return s
}

func (s *GetBucketCommonHeaderResponseBody) Validate() error {
	if s.CommonHeaders != nil {
		if err := s.CommonHeaders.Validate(); err != nil {
			return err
		}
	}
	return nil
}
