// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationListJobs interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v int64) *BatchOperationListJobs
	GetCreationTime() *int64
	SetDescription(v string) *BatchOperationListJobs
	GetDescription() *string
	SetJobId(v string) *BatchOperationListJobs
	GetJobId() *string
	SetOperation(v string) *BatchOperationListJobs
	GetOperation() *string
	SetPriority(v int32) *BatchOperationListJobs
	GetPriority() *int32
	SetProgressSummary(v *BatchOperationJobProgressSummary) *BatchOperationListJobs
	GetProgressSummary() *BatchOperationJobProgressSummary
	SetStatus(v string) *BatchOperationListJobs
	GetStatus() *string
	SetTerminationDate(v int64) *BatchOperationListJobs
	GetTerminationDate() *int64
}

type BatchOperationListJobs struct {
	// example:
	//
	// 1730250699
	CreationTime *int64  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// IDBkODc3M2RlZjgyNjQ0NDViZGV****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// CopyObject
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 5
	Priority        *int32                            `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProgressSummary *BatchOperationJobProgressSummary `json:"ProgressSummary,omitempty" xml:"ProgressSummary,omitempty"`
	// example:
	//
	// Preparing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1730250699
	TerminationDate *int64 `json:"TerminationDate,omitempty" xml:"TerminationDate,omitempty"`
}

func (s BatchOperationListJobs) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationListJobs) GoString() string {
	return s.String()
}

func (s *BatchOperationListJobs) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *BatchOperationListJobs) GetDescription() *string {
	return s.Description
}

func (s *BatchOperationListJobs) GetJobId() *string {
	return s.JobId
}

func (s *BatchOperationListJobs) GetOperation() *string {
	return s.Operation
}

func (s *BatchOperationListJobs) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchOperationListJobs) GetProgressSummary() *BatchOperationJobProgressSummary {
	return s.ProgressSummary
}

func (s *BatchOperationListJobs) GetStatus() *string {
	return s.Status
}

func (s *BatchOperationListJobs) GetTerminationDate() *int64 {
	return s.TerminationDate
}

func (s *BatchOperationListJobs) SetCreationTime(v int64) *BatchOperationListJobs {
	s.CreationTime = &v
	return s
}

func (s *BatchOperationListJobs) SetDescription(v string) *BatchOperationListJobs {
	s.Description = &v
	return s
}

func (s *BatchOperationListJobs) SetJobId(v string) *BatchOperationListJobs {
	s.JobId = &v
	return s
}

func (s *BatchOperationListJobs) SetOperation(v string) *BatchOperationListJobs {
	s.Operation = &v
	return s
}

func (s *BatchOperationListJobs) SetPriority(v int32) *BatchOperationListJobs {
	s.Priority = &v
	return s
}

func (s *BatchOperationListJobs) SetProgressSummary(v *BatchOperationJobProgressSummary) *BatchOperationListJobs {
	s.ProgressSummary = v
	return s
}

func (s *BatchOperationListJobs) SetStatus(v string) *BatchOperationListJobs {
	s.Status = &v
	return s
}

func (s *BatchOperationListJobs) SetTerminationDate(v int64) *BatchOperationListJobs {
	s.TerminationDate = &v
	return s
}

func (s *BatchOperationListJobs) Validate() error {
	if s.ProgressSummary != nil {
		if err := s.ProgressSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}
