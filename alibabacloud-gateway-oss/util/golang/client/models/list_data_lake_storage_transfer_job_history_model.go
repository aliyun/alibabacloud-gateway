// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeStorageTransferJobHistory interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeStorageTransferJobHistory(v []*DataLakeStorageTransferJobHistory) *ListDataLakeStorageTransferJobHistory
	GetDataLakeStorageTransferJobHistory() []*DataLakeStorageTransferJobHistory
}

type ListDataLakeStorageTransferJobHistory struct {
	DataLakeStorageTransferJobHistory []*DataLakeStorageTransferJobHistory `json:"DataLakeStorageTransferJobHistory,omitempty" xml:"DataLakeStorageTransferJobHistory,omitempty" type:"Repeated"`
}

func (s ListDataLakeStorageTransferJobHistory) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeStorageTransferJobHistory) GoString() string {
	return s.String()
}

func (s *ListDataLakeStorageTransferJobHistory) GetDataLakeStorageTransferJobHistory() []*DataLakeStorageTransferJobHistory {
	return s.DataLakeStorageTransferJobHistory
}

func (s *ListDataLakeStorageTransferJobHistory) SetDataLakeStorageTransferJobHistory(v []*DataLakeStorageTransferJobHistory) *ListDataLakeStorageTransferJobHistory {
	s.DataLakeStorageTransferJobHistory = v
	return s
}

func (s *ListDataLakeStorageTransferJobHistory) Validate() error {
	if s.DataLakeStorageTransferJobHistory != nil {
		for _, item := range s.DataLakeStorageTransferJobHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
