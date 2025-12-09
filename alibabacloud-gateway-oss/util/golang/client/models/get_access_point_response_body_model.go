// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetAccessPointResult(v *GetAccessPointResult) *GetAccessPointResponseBody
	GetGetAccessPointResult() *GetAccessPointResult
}

type GetAccessPointResponseBody struct {
	// The container that stores the information about the access point.
	GetAccessPointResult *GetAccessPointResult `json:"GetAccessPointResult,omitempty" xml:"GetAccessPointResult,omitempty"`
}

func (s GetAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPointResponseBody) GetGetAccessPointResult() *GetAccessPointResult {
	return s.GetAccessPointResult
}

func (s *GetAccessPointResponseBody) SetGetAccessPointResult(v *GetAccessPointResult) *GetAccessPointResponseBody {
	s.GetAccessPointResult = v
	return s
}

func (s *GetAccessPointResponseBody) Validate() error {
	if s.GetAccessPointResult != nil {
		if err := s.GetAccessPointResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
