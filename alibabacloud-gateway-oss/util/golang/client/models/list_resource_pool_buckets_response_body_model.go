// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListResourcePoolBucketsResult(v *ListResourcePoolBucketsResult) *ListResourcePoolBucketsResponseBody
	GetListResourcePoolBucketsResult() *ListResourcePoolBucketsResult
}

type ListResourcePoolBucketsResponseBody struct {
	ListResourcePoolBucketsResult *ListResourcePoolBucketsResult `json:"ListResourcePoolBucketsResult,omitempty" xml:"ListResourcePoolBucketsResult,omitempty"`
}

func (s ListResourcePoolBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketsResponseBody) GetListResourcePoolBucketsResult() *ListResourcePoolBucketsResult {
	return s.ListResourcePoolBucketsResult
}

func (s *ListResourcePoolBucketsResponseBody) SetListResourcePoolBucketsResult(v *ListResourcePoolBucketsResult) *ListResourcePoolBucketsResponseBody {
	s.ListResourcePoolBucketsResult = v
	return s
}

func (s *ListResourcePoolBucketsResponseBody) Validate() error {
	if s.ListResourcePoolBucketsResult != nil {
		if err := s.ListResourcePoolBucketsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
