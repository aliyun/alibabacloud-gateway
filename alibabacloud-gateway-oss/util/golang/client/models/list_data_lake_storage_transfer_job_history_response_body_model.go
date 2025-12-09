// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeStorageTransferJobHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListDataLakeStorageTransferJobHistory(v *ListDataLakeStorageTransferJobHistory) *ListDataLakeStorageTransferJobHistoryResponseBody
	GetListDataLakeStorageTransferJobHistory() *ListDataLakeStorageTransferJobHistory
}

type ListDataLakeStorageTransferJobHistoryResponseBody struct {
	ListDataLakeStorageTransferJobHistory *ListDataLakeStorageTransferJobHistory `json:"ListDataLakeStorageTransferJobHistory,omitempty" xml:"ListDataLakeStorageTransferJobHistory,omitempty"`
}

func (s ListDataLakeStorageTransferJobHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeStorageTransferJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeStorageTransferJobHistoryResponseBody) GetListDataLakeStorageTransferJobHistory() *ListDataLakeStorageTransferJobHistory {
	return s.ListDataLakeStorageTransferJobHistory
}

func (s *ListDataLakeStorageTransferJobHistoryResponseBody) SetListDataLakeStorageTransferJobHistory(v *ListDataLakeStorageTransferJobHistory) *ListDataLakeStorageTransferJobHistoryResponseBody {
	s.ListDataLakeStorageTransferJobHistory = v
	return s
}

func (s *ListDataLakeStorageTransferJobHistoryResponseBody) Validate() error {
	if s.ListDataLakeStorageTransferJobHistory != nil {
		if err := s.ListDataLakeStorageTransferJobHistory.Validate(); err != nil {
			return err
		}
	}
	return nil
}
