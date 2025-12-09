// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobPriorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJobId(v string) *UpdateJobPriorityRequest
	GetBatchJobId() *string
	SetTargetPriority(v int32) *UpdateJobPriorityRequest
	GetTargetPriority() *int32
}

type UpdateJobPriorityRequest struct {
	BatchJobId     *string `json:"batchJobId,omitempty" xml:"batchJobId,omitempty"`
	TargetPriority *int32  `json:"targetPriority,omitempty" xml:"targetPriority,omitempty"`
}

func (s UpdateJobPriorityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPriorityRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobPriorityRequest) GetBatchJobId() *string {
	return s.BatchJobId
}

func (s *UpdateJobPriorityRequest) GetTargetPriority() *int32 {
	return s.TargetPriority
}

func (s *UpdateJobPriorityRequest) SetBatchJobId(v string) *UpdateJobPriorityRequest {
	s.BatchJobId = &v
	return s
}

func (s *UpdateJobPriorityRequest) SetTargetPriority(v int32) *UpdateJobPriorityRequest {
	s.TargetPriority = &v
	return s
}

func (s *UpdateJobPriorityRequest) Validate() error {
	return dara.Validate(s)
}
