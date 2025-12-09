// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeStorageTransferJobId interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DataLakeStorageTransferJobId
	GetId() *string
}

type DataLakeStorageTransferJobId struct {
	// example:
	//
	// abcde787e3c04af2af290ec52d123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DataLakeStorageTransferJobId) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobId) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobId) GetId() *string {
	return s.Id
}

func (s *DataLakeStorageTransferJobId) SetId(v string) *DataLakeStorageTransferJobId {
	s.Id = &v
	return s
}

func (s *DataLakeStorageTransferJobId) Validate() error {
	return dara.Validate(s)
}
