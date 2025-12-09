// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataLakeStorageTransferJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeStorageTransferJobId(v *DataLakeStorageTransferJobId) *PutDataLakeStorageTransferJobResponseBody
	GetDataLakeStorageTransferJobId() *DataLakeStorageTransferJobId
}

type PutDataLakeStorageTransferJobResponseBody struct {
	DataLakeStorageTransferJobId *DataLakeStorageTransferJobId `json:"DataLakeStorageTransferJobId,omitempty" xml:"DataLakeStorageTransferJobId,omitempty"`
}

func (s PutDataLakeStorageTransferJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeStorageTransferJobResponseBody) GoString() string {
	return s.String()
}

func (s *PutDataLakeStorageTransferJobResponseBody) GetDataLakeStorageTransferJobId() *DataLakeStorageTransferJobId {
	return s.DataLakeStorageTransferJobId
}

func (s *PutDataLakeStorageTransferJobResponseBody) SetDataLakeStorageTransferJobId(v *DataLakeStorageTransferJobId) *PutDataLakeStorageTransferJobResponseBody {
	s.DataLakeStorageTransferJobId = v
	return s
}

func (s *PutDataLakeStorageTransferJobResponseBody) Validate() error {
	if s.DataLakeStorageTransferJobId != nil {
		if err := s.DataLakeStorageTransferJobId.Validate(); err != nil {
			return err
		}
	}
	return nil
}
