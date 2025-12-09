// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeStorageTransferJobHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *ListDataLakeStorageTransferJobHistoryRequest
	GetXOssDatalakeJobId() *string
}

type ListDataLakeStorageTransferJobHistoryRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s ListDataLakeStorageTransferJobHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeStorageTransferJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeStorageTransferJobHistoryRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *ListDataLakeStorageTransferJobHistoryRequest) SetXOssDatalakeJobId(v string) *ListDataLakeStorageTransferJobHistoryRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *ListDataLakeStorageTransferJobHistoryRequest) Validate() error {
	return dara.Validate(s)
}
