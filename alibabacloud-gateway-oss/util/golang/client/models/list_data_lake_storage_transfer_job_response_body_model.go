// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeStorageTransferJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeStorageTransferJobs(v *DataLakeStorageTransferJobs) *ListDataLakeStorageTransferJobResponseBody
	GetDataLakeStorageTransferJobs() *DataLakeStorageTransferJobs
}

type ListDataLakeStorageTransferJobResponseBody struct {
	DataLakeStorageTransferJobs *DataLakeStorageTransferJobs `json:"DataLakeStorageTransferJobs,omitempty" xml:"DataLakeStorageTransferJobs,omitempty"`
}

func (s ListDataLakeStorageTransferJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeStorageTransferJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeStorageTransferJobResponseBody) GetDataLakeStorageTransferJobs() *DataLakeStorageTransferJobs {
	return s.DataLakeStorageTransferJobs
}

func (s *ListDataLakeStorageTransferJobResponseBody) SetDataLakeStorageTransferJobs(v *DataLakeStorageTransferJobs) *ListDataLakeStorageTransferJobResponseBody {
	s.DataLakeStorageTransferJobs = v
	return s
}

func (s *ListDataLakeStorageTransferJobResponseBody) Validate() error {
	if s.DataLakeStorageTransferJobs != nil {
		if err := s.DataLakeStorageTransferJobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}
