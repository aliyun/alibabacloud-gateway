// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeStorageTransferJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLakeStorageTransferJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLakeStorageTransferJobResponse
	GetStatusCode() *int32
}

type DeleteDataLakeStorageTransferJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDataLakeStorageTransferJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeStorageTransferJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeStorageTransferJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLakeStorageTransferJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLakeStorageTransferJobResponse) SetHeaders(v map[string]*string) *DeleteDataLakeStorageTransferJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLakeStorageTransferJobResponse) SetStatusCode(v int32) *DeleteDataLakeStorageTransferJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLakeStorageTransferJobResponse) Validate() error {
	return dara.Validate(s)
}
