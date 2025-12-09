// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListAccessPointsResult(v *ListAccessPointsResult) *ListAccessPointsResponseBody
	GetListAccessPointsResult() *ListAccessPointsResult
}

type ListAccessPointsResponseBody struct {
	// The container that stores the information about access points.
	ListAccessPointsResult *ListAccessPointsResult `json:"ListAccessPointsResult,omitempty" xml:"ListAccessPointsResult,omitempty"`
}

func (s ListAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBody) GetListAccessPointsResult() *ListAccessPointsResult {
	return s.ListAccessPointsResult
}

func (s *ListAccessPointsResponseBody) SetListAccessPointsResult(v *ListAccessPointsResult) *ListAccessPointsResponseBody {
	s.ListAccessPointsResult = v
	return s
}

func (s *ListAccessPointsResponseBody) Validate() error {
	if s.ListAccessPointsResult != nil {
		if err := s.ListAccessPointsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
