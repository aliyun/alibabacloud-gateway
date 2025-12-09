// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupQoSInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListResourcePoolBucketGroupQoSInfosResult(v *ListResourcePoolBucketGroupQoSInfosResult) *ListResourcePoolBucketGroupQoSInfosResponseBody
	GetListResourcePoolBucketGroupQoSInfosResult() *ListResourcePoolBucketGroupQoSInfosResult
}

type ListResourcePoolBucketGroupQoSInfosResponseBody struct {
	ListResourcePoolBucketGroupQoSInfosResult *ListResourcePoolBucketGroupQoSInfosResult `json:"ListResourcePoolBucketGroupQoSInfosResult,omitempty" xml:"ListResourcePoolBucketGroupQoSInfosResult,omitempty"`
}

func (s ListResourcePoolBucketGroupQoSInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupQoSInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupQoSInfosResponseBody) GetListResourcePoolBucketGroupQoSInfosResult() *ListResourcePoolBucketGroupQoSInfosResult {
	return s.ListResourcePoolBucketGroupQoSInfosResult
}

func (s *ListResourcePoolBucketGroupQoSInfosResponseBody) SetListResourcePoolBucketGroupQoSInfosResult(v *ListResourcePoolBucketGroupQoSInfosResult) *ListResourcePoolBucketGroupQoSInfosResponseBody {
	s.ListResourcePoolBucketGroupQoSInfosResult = v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResponseBody) Validate() error {
	if s.ListResourcePoolBucketGroupQoSInfosResult != nil {
		if err := s.ListResourcePoolBucketGroupQoSInfosResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
