// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolRequesterQoSInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListResourcePoolRequesterQoSInfosResult(v *ListResourcePoolRequesterQoSInfosResult) *ListResourcePoolRequesterQoSInfosResponseBody
	GetListResourcePoolRequesterQoSInfosResult() *ListResourcePoolRequesterQoSInfosResult
}

type ListResourcePoolRequesterQoSInfosResponseBody struct {
	ListResourcePoolRequesterQoSInfosResult *ListResourcePoolRequesterQoSInfosResult `json:"ListResourcePoolRequesterQoSInfosResult,omitempty" xml:"ListResourcePoolRequesterQoSInfosResult,omitempty"`
}

func (s ListResourcePoolRequesterQoSInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolRequesterQoSInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePoolRequesterQoSInfosResponseBody) GetListResourcePoolRequesterQoSInfosResult() *ListResourcePoolRequesterQoSInfosResult {
	return s.ListResourcePoolRequesterQoSInfosResult
}

func (s *ListResourcePoolRequesterQoSInfosResponseBody) SetListResourcePoolRequesterQoSInfosResult(v *ListResourcePoolRequesterQoSInfosResult) *ListResourcePoolRequesterQoSInfosResponseBody {
	s.ListResourcePoolRequesterQoSInfosResult = v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResponseBody) Validate() error {
	if s.ListResourcePoolRequesterQoSInfosResult != nil {
		if err := s.ListResourcePoolRequesterQoSInfosResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
