// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListResourcePoolsResult(v *ListResourcePoolsResult) *ListResourcePoolsResponseBody
	GetListResourcePoolsResult() *ListResourcePoolsResult
}

type ListResourcePoolsResponseBody struct {
	ListResourcePoolsResult *ListResourcePoolsResult `json:"ListResourcePoolsResult,omitempty" xml:"ListResourcePoolsResult,omitempty"`
}

func (s ListResourcePoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePoolsResponseBody) GetListResourcePoolsResult() *ListResourcePoolsResult {
	return s.ListResourcePoolsResult
}

func (s *ListResourcePoolsResponseBody) SetListResourcePoolsResult(v *ListResourcePoolsResult) *ListResourcePoolsResponseBody {
	s.ListResourcePoolsResult = v
	return s
}

func (s *ListResourcePoolsResponseBody) Validate() error {
	if s.ListResourcePoolsResult != nil {
		if err := s.ListResourcePoolsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
