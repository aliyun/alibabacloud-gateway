// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeStorageTransferJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *GetDataLakeStorageTransferJobRequest
	GetXOssDatalakeJobId() *string
	SetXOssDatalakeJobProgress(v string) *GetDataLakeStorageTransferJobRequest
	GetXOssDatalakeJobProgress() *string
}

type GetDataLakeStorageTransferJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId       *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
	XOssDatalakeJobProgress *string `json:"x-oss-datalake-job-progress,omitempty" xml:"x-oss-datalake-job-progress,omitempty"`
}

func (s GetDataLakeStorageTransferJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeStorageTransferJobRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakeStorageTransferJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *GetDataLakeStorageTransferJobRequest) GetXOssDatalakeJobProgress() *string {
	return s.XOssDatalakeJobProgress
}

func (s *GetDataLakeStorageTransferJobRequest) SetXOssDatalakeJobId(v string) *GetDataLakeStorageTransferJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *GetDataLakeStorageTransferJobRequest) SetXOssDatalakeJobProgress(v string) *GetDataLakeStorageTransferJobRequest {
	s.XOssDatalakeJobProgress = &v
	return s
}

func (s *GetDataLakeStorageTransferJobRequest) Validate() error {
	return dara.Validate(s)
}
