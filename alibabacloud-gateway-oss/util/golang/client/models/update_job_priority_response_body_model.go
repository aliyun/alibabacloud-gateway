// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobPriorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateJobPriorityResult(v *BatchOperationUpdateJobPriorityResult) *UpdateJobPriorityResponseBody
	GetUpdateJobPriorityResult() *BatchOperationUpdateJobPriorityResult
}

type UpdateJobPriorityResponseBody struct {
	UpdateJobPriorityResult *BatchOperationUpdateJobPriorityResult `json:"UpdateJobPriorityResult,omitempty" xml:"UpdateJobPriorityResult,omitempty"`
}

func (s UpdateJobPriorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPriorityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobPriorityResponseBody) GetUpdateJobPriorityResult() *BatchOperationUpdateJobPriorityResult {
	return s.UpdateJobPriorityResult
}

func (s *UpdateJobPriorityResponseBody) SetUpdateJobPriorityResult(v *BatchOperationUpdateJobPriorityResult) *UpdateJobPriorityResponseBody {
	s.UpdateJobPriorityResult = v
	return s
}

func (s *UpdateJobPriorityResponseBody) Validate() error {
	if s.UpdateJobPriorityResult != nil {
		if err := s.UpdateJobPriorityResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
