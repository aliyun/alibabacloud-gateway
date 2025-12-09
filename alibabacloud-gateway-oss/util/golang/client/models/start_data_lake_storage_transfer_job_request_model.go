// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataLakeStorageTransferJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *StartDataLakeStorageTransferJobRequest
	GetXOssDatalakeJobId() *string
}

type StartDataLakeStorageTransferJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s StartDataLakeStorageTransferJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDataLakeStorageTransferJobRequest) GoString() string {
	return s.String()
}

func (s *StartDataLakeStorageTransferJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *StartDataLakeStorageTransferJobRequest) SetXOssDatalakeJobId(v string) *StartDataLakeStorageTransferJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *StartDataLakeStorageTransferJobRequest) Validate() error {
	return dara.Validate(s)
}
