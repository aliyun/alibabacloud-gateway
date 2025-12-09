// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataLakeStorageTransferJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDataLakeStorageTransferJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDataLakeStorageTransferJobResponse
	GetStatusCode() *int32
}

type StopDataLakeStorageTransferJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopDataLakeStorageTransferJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDataLakeStorageTransferJobResponse) GoString() string {
	return s.String()
}

func (s *StopDataLakeStorageTransferJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDataLakeStorageTransferJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDataLakeStorageTransferJobResponse) SetHeaders(v map[string]*string) *StopDataLakeStorageTransferJobResponse {
	s.Headers = v
	return s
}

func (s *StopDataLakeStorageTransferJobResponse) SetStatusCode(v int32) *StopDataLakeStorageTransferJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataLakeStorageTransferJobResponse) Validate() error {
	return dara.Validate(s)
}
