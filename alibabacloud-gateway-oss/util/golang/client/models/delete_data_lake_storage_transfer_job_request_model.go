// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeStorageTransferJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *DeleteDataLakeStorageTransferJobRequest
	GetXOssDatalakeJobId() *string
}

type DeleteDataLakeStorageTransferJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s DeleteDataLakeStorageTransferJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeStorageTransferJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeStorageTransferJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *DeleteDataLakeStorageTransferJobRequest) SetXOssDatalakeJobId(v string) *DeleteDataLakeStorageTransferJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *DeleteDataLakeStorageTransferJobRequest) Validate() error {
	return dara.Validate(s)
}
