// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListResourcePoolBucketGroupsResult(v *ListResourcePoolBucketGroupsResult) *ListResourcePoolBucketGroupsResponseBody
	GetListResourcePoolBucketGroupsResult() *ListResourcePoolBucketGroupsResult
}

type ListResourcePoolBucketGroupsResponseBody struct {
	ListResourcePoolBucketGroupsResult *ListResourcePoolBucketGroupsResult `json:"ListResourcePoolBucketGroupsResult,omitempty" xml:"ListResourcePoolBucketGroupsResult,omitempty"`
}

func (s ListResourcePoolBucketGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupsResponseBody) GetListResourcePoolBucketGroupsResult() *ListResourcePoolBucketGroupsResult {
	return s.ListResourcePoolBucketGroupsResult
}

func (s *ListResourcePoolBucketGroupsResponseBody) SetListResourcePoolBucketGroupsResult(v *ListResourcePoolBucketGroupsResult) *ListResourcePoolBucketGroupsResponseBody {
	s.ListResourcePoolBucketGroupsResult = v
	return s
}

func (s *ListResourcePoolBucketGroupsResponseBody) Validate() error {
	if s.ListResourcePoolBucketGroupsResult != nil {
		if err := s.ListResourcePoolBucketGroupsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
