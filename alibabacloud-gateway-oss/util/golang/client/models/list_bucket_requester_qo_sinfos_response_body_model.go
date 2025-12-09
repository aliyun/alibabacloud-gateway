// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRequesterQoSInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListBucketRequesterQoSInfosResult(v *ListBucketRequesterQoSInfosResult) *ListBucketRequesterQoSInfosResponseBody
	GetListBucketRequesterQoSInfosResult() *ListBucketRequesterQoSInfosResult
}

type ListBucketRequesterQoSInfosResponseBody struct {
	ListBucketRequesterQoSInfosResult *ListBucketRequesterQoSInfosResult `json:"ListBucketRequesterQoSInfosResult,omitempty" xml:"ListBucketRequesterQoSInfosResult,omitempty"`
}

func (s ListBucketRequesterQoSInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRequesterQoSInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketRequesterQoSInfosResponseBody) GetListBucketRequesterQoSInfosResult() *ListBucketRequesterQoSInfosResult {
	return s.ListBucketRequesterQoSInfosResult
}

func (s *ListBucketRequesterQoSInfosResponseBody) SetListBucketRequesterQoSInfosResult(v *ListBucketRequesterQoSInfosResult) *ListBucketRequesterQoSInfosResponseBody {
	s.ListBucketRequesterQoSInfosResult = v
	return s
}

func (s *ListBucketRequesterQoSInfosResponseBody) Validate() error {
	if s.ListBucketRequesterQoSInfosResult != nil {
		if err := s.ListBucketRequesterQoSInfosResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
