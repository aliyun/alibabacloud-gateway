// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListAllMyCacheResult(v *ListAllMyCacheResult) *ListCacheResponseBody
	GetListAllMyCacheResult() *ListAllMyCacheResult
}

type ListCacheResponseBody struct {
	ListAllMyCacheResult *ListAllMyCacheResult `json:"ListAllMyCacheResult,omitempty" xml:"ListAllMyCacheResult,omitempty"`
}

func (s ListCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCacheResponseBody) GoString() string {
	return s.String()
}

func (s *ListCacheResponseBody) GetListAllMyCacheResult() *ListAllMyCacheResult {
	return s.ListAllMyCacheResult
}

func (s *ListCacheResponseBody) SetListAllMyCacheResult(v *ListAllMyCacheResult) *ListCacheResponseBody {
	s.ListAllMyCacheResult = v
	return s
}

func (s *ListCacheResponseBody) Validate() error {
	if s.ListAllMyCacheResult != nil {
		if err := s.ListAllMyCacheResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
