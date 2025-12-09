// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataLakeStorageTransferJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateDataLakeStorageTransferJob(v *CreateDataLakeStorageTransferJob) *PutDataLakeStorageTransferJobRequest
	GetCreateDataLakeStorageTransferJob() *CreateDataLakeStorageTransferJob
	SetXOssDatalakeJobId(v string) *PutDataLakeStorageTransferJobRequest
	GetXOssDatalakeJobId() *string
}

type PutDataLakeStorageTransferJobRequest struct {
	// This parameter is required.
	CreateDataLakeStorageTransferJob *CreateDataLakeStorageTransferJob `json:"CreateDataLakeStorageTransferJob,omitempty" xml:"CreateDataLakeStorageTransferJob,omitempty"`
	XOssDatalakeJobId                *string                           `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s PutDataLakeStorageTransferJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeStorageTransferJobRequest) GoString() string {
	return s.String()
}

func (s *PutDataLakeStorageTransferJobRequest) GetCreateDataLakeStorageTransferJob() *CreateDataLakeStorageTransferJob {
	return s.CreateDataLakeStorageTransferJob
}

func (s *PutDataLakeStorageTransferJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *PutDataLakeStorageTransferJobRequest) SetCreateDataLakeStorageTransferJob(v *CreateDataLakeStorageTransferJob) *PutDataLakeStorageTransferJobRequest {
	s.CreateDataLakeStorageTransferJob = v
	return s
}

func (s *PutDataLakeStorageTransferJobRequest) SetXOssDatalakeJobId(v string) *PutDataLakeStorageTransferJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *PutDataLakeStorageTransferJobRequest) Validate() error {
	if s.CreateDataLakeStorageTransferJob != nil {
		if err := s.CreateDataLakeStorageTransferJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
