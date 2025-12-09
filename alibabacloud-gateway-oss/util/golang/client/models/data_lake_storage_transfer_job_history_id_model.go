// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeStorageTransferJobHistoryId interface {
	dara.Model
	String() string
	GoString() string
	SetHistoryId(v string) *DataLakeStorageTransferJobHistoryId
	GetHistoryId() *string
}

type DataLakeStorageTransferJobHistoryId struct {
	// example:
	//
	// abcdef03370a4b32ac6bc69eb1123456
	HistoryId *string `json:"HistoryId,omitempty" xml:"HistoryId,omitempty"`
}

func (s DataLakeStorageTransferJobHistoryId) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobHistoryId) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobHistoryId) GetHistoryId() *string {
	return s.HistoryId
}

func (s *DataLakeStorageTransferJobHistoryId) SetHistoryId(v string) *DataLakeStorageTransferJobHistoryId {
	s.HistoryId = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryId) Validate() error {
	return dara.Validate(s)
}
