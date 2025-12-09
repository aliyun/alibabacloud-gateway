// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationUpdateJobStatusResult interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *BatchOperationUpdateJobStatusResult
	GetJobId() *string
	SetStatus(v string) *BatchOperationUpdateJobStatusResult
	GetStatus() *string
	SetStatusUpdateReason(v string) *BatchOperationUpdateJobStatusResult
	GetStatusUpdateReason() *string
}

type BatchOperationUpdateJobStatusResult struct {
	// example:
	//
	// IDBkODc3M2RlZjgyNjQ0NDViZGV****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Preparing
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusUpdateReason *string `json:"StatusUpdateReason,omitempty" xml:"StatusUpdateReason,omitempty"`
}

func (s BatchOperationUpdateJobStatusResult) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationUpdateJobStatusResult) GoString() string {
	return s.String()
}

func (s *BatchOperationUpdateJobStatusResult) GetJobId() *string {
	return s.JobId
}

func (s *BatchOperationUpdateJobStatusResult) GetStatus() *string {
	return s.Status
}

func (s *BatchOperationUpdateJobStatusResult) GetStatusUpdateReason() *string {
	return s.StatusUpdateReason
}

func (s *BatchOperationUpdateJobStatusResult) SetJobId(v string) *BatchOperationUpdateJobStatusResult {
	s.JobId = &v
	return s
}

func (s *BatchOperationUpdateJobStatusResult) SetStatus(v string) *BatchOperationUpdateJobStatusResult {
	s.Status = &v
	return s
}

func (s *BatchOperationUpdateJobStatusResult) SetStatusUpdateReason(v string) *BatchOperationUpdateJobStatusResult {
	s.StatusUpdateReason = &v
	return s
}

func (s *BatchOperationUpdateJobStatusResult) Validate() error {
	return dara.Validate(s)
}
