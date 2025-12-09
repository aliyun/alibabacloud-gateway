// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListUserRegionsResult(v *ListUserRegionsResult) *ListUserRegionsResponseBody
	GetListUserRegionsResult() *ListUserRegionsResult
}

type ListUserRegionsResponseBody struct {
	ListUserRegionsResult *ListUserRegionsResult `json:"ListUserRegionsResult,omitempty" xml:"ListUserRegionsResult,omitempty"`
}

func (s ListUserRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserRegionsResponseBody) GetListUserRegionsResult() *ListUserRegionsResult {
	return s.ListUserRegionsResult
}

func (s *ListUserRegionsResponseBody) SetListUserRegionsResult(v *ListUserRegionsResult) *ListUserRegionsResponseBody {
	s.ListUserRegionsResult = v
	return s
}

func (s *ListUserRegionsResponseBody) Validate() error {
	if s.ListUserRegionsResult != nil {
		if err := s.ListUserRegionsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
