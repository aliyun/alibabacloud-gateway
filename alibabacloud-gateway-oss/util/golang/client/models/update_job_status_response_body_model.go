// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateJobStatusResult(v *BatchOperationUpdateJobStatusResult) *UpdateJobStatusResponseBody
	GetUpdateJobStatusResult() *BatchOperationUpdateJobStatusResult
}

type UpdateJobStatusResponseBody struct {
	UpdateJobStatusResult *BatchOperationUpdateJobStatusResult `json:"UpdateJobStatusResult,omitempty" xml:"UpdateJobStatusResult,omitempty"`
}

func (s UpdateJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobStatusResponseBody) GetUpdateJobStatusResult() *BatchOperationUpdateJobStatusResult {
	return s.UpdateJobStatusResult
}

func (s *UpdateJobStatusResponseBody) SetUpdateJobStatusResult(v *BatchOperationUpdateJobStatusResult) *UpdateJobStatusResponseBody {
	s.UpdateJobStatusResult = v
	return s
}

func (s *UpdateJobStatusResponseBody) Validate() error {
	if s.UpdateJobStatusResult != nil {
		if err := s.UpdateJobStatusResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
