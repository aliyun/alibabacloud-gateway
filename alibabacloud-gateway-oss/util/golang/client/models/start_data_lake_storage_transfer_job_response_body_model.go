// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataLakeStorageTransferJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeStorageTransferJobHistoryId(v *DataLakeStorageTransferJobHistoryId) *StartDataLakeStorageTransferJobResponseBody
	GetDataLakeStorageTransferJobHistoryId() *DataLakeStorageTransferJobHistoryId
}

type StartDataLakeStorageTransferJobResponseBody struct {
	DataLakeStorageTransferJobHistoryId *DataLakeStorageTransferJobHistoryId `json:"DataLakeStorageTransferJobHistoryId,omitempty" xml:"DataLakeStorageTransferJobHistoryId,omitempty"`
}

func (s StartDataLakeStorageTransferJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDataLakeStorageTransferJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartDataLakeStorageTransferJobResponseBody) GetDataLakeStorageTransferJobHistoryId() *DataLakeStorageTransferJobHistoryId {
	return s.DataLakeStorageTransferJobHistoryId
}

func (s *StartDataLakeStorageTransferJobResponseBody) SetDataLakeStorageTransferJobHistoryId(v *DataLakeStorageTransferJobHistoryId) *StartDataLakeStorageTransferJobResponseBody {
	s.DataLakeStorageTransferJobHistoryId = v
	return s
}

func (s *StartDataLakeStorageTransferJobResponseBody) Validate() error {
	if s.DataLakeStorageTransferJobHistoryId != nil {
		if err := s.DataLakeStorageTransferJobHistoryId.Validate(); err != nil {
			return err
		}
	}
	return nil
}
