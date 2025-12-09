// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAccessPointResult(v *CreateAccessPointResult) *CreateAccessPointResponseBody
	GetCreateAccessPointResult() *CreateAccessPointResult
}

type CreateAccessPointResponseBody struct {
	// The container that stores the information about the access point.
	CreateAccessPointResult *CreateAccessPointResult `json:"CreateAccessPointResult,omitempty" xml:"CreateAccessPointResult,omitempty"`
}

func (s CreateAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponseBody) GetCreateAccessPointResult() *CreateAccessPointResult {
	return s.CreateAccessPointResult
}

func (s *CreateAccessPointResponseBody) SetCreateAccessPointResult(v *CreateAccessPointResult) *CreateAccessPointResponseBody {
	s.CreateAccessPointResult = v
	return s
}

func (s *CreateAccessPointResponseBody) Validate() error {
	if s.CreateAccessPointResult != nil {
		if err := s.CreateAccessPointResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
