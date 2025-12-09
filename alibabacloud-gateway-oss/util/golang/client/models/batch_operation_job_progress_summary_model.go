// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationJobProgressSummary interface {
	dara.Model
	String() string
	GoString() string
	SetElapsedTimeInActiveSeconds(v int64) *BatchOperationJobProgressSummary
	GetElapsedTimeInActiveSeconds() *int64
	SetNumberOfTasksFailed(v int64) *BatchOperationJobProgressSummary
	GetNumberOfTasksFailed() *int64
	SetNumberOfTasksSucceeded(v int64) *BatchOperationJobProgressSummary
	GetNumberOfTasksSucceeded() *int64
	SetTotalNumberOfTasks(v int64) *BatchOperationJobProgressSummary
	GetTotalNumberOfTasks() *int64
}

type BatchOperationJobProgressSummary struct {
	// example:
	//
	// 3600
	ElapsedTimeInActiveSeconds *int64 `json:"ElapsedTimeInActiveSeconds,omitempty" xml:"ElapsedTimeInActiveSeconds,omitempty"`
	// example:
	//
	// 0
	NumberOfTasksFailed *int64 `json:"NumberOfTasksFailed,omitempty" xml:"NumberOfTasksFailed,omitempty"`
	// example:
	//
	// 100
	NumberOfTasksSucceeded *int64 `json:"NumberOfTasksSucceeded,omitempty" xml:"NumberOfTasksSucceeded,omitempty"`
	// example:
	//
	// 100
	TotalNumberOfTasks *int64 `json:"TotalNumberOfTasks,omitempty" xml:"TotalNumberOfTasks,omitempty"`
}

func (s BatchOperationJobProgressSummary) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationJobProgressSummary) GoString() string {
	return s.String()
}

func (s *BatchOperationJobProgressSummary) GetElapsedTimeInActiveSeconds() *int64 {
	return s.ElapsedTimeInActiveSeconds
}

func (s *BatchOperationJobProgressSummary) GetNumberOfTasksFailed() *int64 {
	return s.NumberOfTasksFailed
}

func (s *BatchOperationJobProgressSummary) GetNumberOfTasksSucceeded() *int64 {
	return s.NumberOfTasksSucceeded
}

func (s *BatchOperationJobProgressSummary) GetTotalNumberOfTasks() *int64 {
	return s.TotalNumberOfTasks
}

func (s *BatchOperationJobProgressSummary) SetElapsedTimeInActiveSeconds(v int64) *BatchOperationJobProgressSummary {
	s.ElapsedTimeInActiveSeconds = &v
	return s
}

func (s *BatchOperationJobProgressSummary) SetNumberOfTasksFailed(v int64) *BatchOperationJobProgressSummary {
	s.NumberOfTasksFailed = &v
	return s
}

func (s *BatchOperationJobProgressSummary) SetNumberOfTasksSucceeded(v int64) *BatchOperationJobProgressSummary {
	s.NumberOfTasksSucceeded = &v
	return s
}

func (s *BatchOperationJobProgressSummary) SetTotalNumberOfTasks(v int64) *BatchOperationJobProgressSummary {
	s.TotalNumberOfTasks = &v
	return s
}

func (s *BatchOperationJobProgressSummary) Validate() error {
	return dara.Validate(s)
}
