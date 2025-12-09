// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationUpdateJobPriorityResult interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *BatchOperationUpdateJobPriorityResult
	GetJobId() *string
	SetPriority(v int32) *BatchOperationUpdateJobPriorityResult
	GetPriority() *int32
}

type BatchOperationUpdateJobPriorityResult struct {
	// example:
	//
	// IDBkODc3M2RlZjgyNjQ0NDViZGV****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s BatchOperationUpdateJobPriorityResult) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationUpdateJobPriorityResult) GoString() string {
	return s.String()
}

func (s *BatchOperationUpdateJobPriorityResult) GetJobId() *string {
	return s.JobId
}

func (s *BatchOperationUpdateJobPriorityResult) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchOperationUpdateJobPriorityResult) SetJobId(v string) *BatchOperationUpdateJobPriorityResult {
	s.JobId = &v
	return s
}

func (s *BatchOperationUpdateJobPriorityResult) SetPriority(v int32) *BatchOperationUpdateJobPriorityResult {
	s.Priority = &v
	return s
}

func (s *BatchOperationUpdateJobPriorityResult) Validate() error {
	return dara.Validate(s)
}
