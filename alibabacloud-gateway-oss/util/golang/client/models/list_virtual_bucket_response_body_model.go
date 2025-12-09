// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListVirtualBucketResult(v *ListVirtualBucketResult) *ListVirtualBucketResponseBody
	GetListVirtualBucketResult() *ListVirtualBucketResult
}

type ListVirtualBucketResponseBody struct {
	ListVirtualBucketResult *ListVirtualBucketResult `json:"ListVirtualBucketResult,omitempty" xml:"ListVirtualBucketResult,omitempty"`
}

func (s ListVirtualBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBucketResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualBucketResponseBody) GetListVirtualBucketResult() *ListVirtualBucketResult {
	return s.ListVirtualBucketResult
}

func (s *ListVirtualBucketResponseBody) SetListVirtualBucketResult(v *ListVirtualBucketResult) *ListVirtualBucketResponseBody {
	s.ListVirtualBucketResult = v
	return s
}

func (s *ListVirtualBucketResponseBody) Validate() error {
	if s.ListVirtualBucketResult != nil {
		if err := s.ListVirtualBucketResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
