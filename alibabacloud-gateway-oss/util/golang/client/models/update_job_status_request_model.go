// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJobId(v string) *UpdateJobStatusRequest
	GetBatchJobId() *string
	SetRequestedJobStatus(v string) *UpdateJobStatusRequest
	GetRequestedJobStatus() *string
	SetStatusUpdateReason(v string) *UpdateJobStatusRequest
	GetStatusUpdateReason() *string
}

type UpdateJobStatusRequest struct {
	BatchJobId         *string `json:"batchJobId,omitempty" xml:"batchJobId,omitempty"`
	RequestedJobStatus *string `json:"requestedJobStatus,omitempty" xml:"requestedJobStatus,omitempty"`
	StatusUpdateReason *string `json:"statusUpdateReason,omitempty" xml:"statusUpdateReason,omitempty"`
}

func (s UpdateJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobStatusRequest) GetBatchJobId() *string {
	return s.BatchJobId
}

func (s *UpdateJobStatusRequest) GetRequestedJobStatus() *string {
	return s.RequestedJobStatus
}

func (s *UpdateJobStatusRequest) GetStatusUpdateReason() *string {
	return s.StatusUpdateReason
}

func (s *UpdateJobStatusRequest) SetBatchJobId(v string) *UpdateJobStatusRequest {
	s.BatchJobId = &v
	return s
}

func (s *UpdateJobStatusRequest) SetRequestedJobStatus(v string) *UpdateJobStatusRequest {
	s.RequestedJobStatus = &v
	return s
}

func (s *UpdateJobStatusRequest) SetStatusUpdateReason(v string) *UpdateJobStatusRequest {
	s.StatusUpdateReason = &v
	return s
}

func (s *UpdateJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
