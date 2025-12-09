// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataLakeStorageTransferJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *StopDataLakeStorageTransferJobRequest
	GetXOssDatalakeJobId() *string
}

type StopDataLakeStorageTransferJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s StopDataLakeStorageTransferJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDataLakeStorageTransferJobRequest) GoString() string {
	return s.String()
}

func (s *StopDataLakeStorageTransferJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *StopDataLakeStorageTransferJobRequest) SetXOssDatalakeJobId(v string) *StopDataLakeStorageTransferJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *StopDataLakeStorageTransferJobRequest) Validate() error {
	return dara.Validate(s)
}
