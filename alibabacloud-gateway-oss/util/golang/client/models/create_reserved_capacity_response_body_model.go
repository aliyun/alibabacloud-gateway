// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReservedCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateLargeReservedCapacityResult(v *CreateLargeReservedCapacityResult) *CreateReservedCapacityResponseBody
	GetCreateLargeReservedCapacityResult() *CreateLargeReservedCapacityResult
}

type CreateReservedCapacityResponseBody struct {
	CreateLargeReservedCapacityResult *CreateLargeReservedCapacityResult `json:"CreateLargeReservedCapacityResult,omitempty" xml:"CreateLargeReservedCapacityResult,omitempty"`
}

func (s CreateReservedCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReservedCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReservedCapacityResponseBody) GetCreateLargeReservedCapacityResult() *CreateLargeReservedCapacityResult {
	return s.CreateLargeReservedCapacityResult
}

func (s *CreateReservedCapacityResponseBody) SetCreateLargeReservedCapacityResult(v *CreateLargeReservedCapacityResult) *CreateReservedCapacityResponseBody {
	s.CreateLargeReservedCapacityResult = v
	return s
}

func (s *CreateReservedCapacityResponseBody) Validate() error {
	if s.CreateLargeReservedCapacityResult != nil {
		if err := s.CreateLargeReservedCapacityResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
