// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeStorageTransferJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeStorageTransferJob(v *DataLakeStorageTransferJob) *GetDataLakeStorageTransferJobResponseBody
	GetDataLakeStorageTransferJob() *DataLakeStorageTransferJob
}

type GetDataLakeStorageTransferJobResponseBody struct {
	DataLakeStorageTransferJob *DataLakeStorageTransferJob `json:"DataLakeStorageTransferJob,omitempty" xml:"DataLakeStorageTransferJob,omitempty"`
}

func (s GetDataLakeStorageTransferJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeStorageTransferJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakeStorageTransferJobResponseBody) GetDataLakeStorageTransferJob() *DataLakeStorageTransferJob {
	return s.DataLakeStorageTransferJob
}

func (s *GetDataLakeStorageTransferJobResponseBody) SetDataLakeStorageTransferJob(v *DataLakeStorageTransferJob) *GetDataLakeStorageTransferJobResponseBody {
	s.DataLakeStorageTransferJob = v
	return s
}

func (s *GetDataLakeStorageTransferJobResponseBody) Validate() error {
	if s.DataLakeStorageTransferJob != nil {
		if err := s.DataLakeStorageTransferJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
