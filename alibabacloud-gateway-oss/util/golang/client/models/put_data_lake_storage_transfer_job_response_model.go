// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataLakeStorageTransferJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDataLakeStorageTransferJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDataLakeStorageTransferJobResponse
	GetStatusCode() *int32
	SetBody(v *PutDataLakeStorageTransferJobResponseBody) *PutDataLakeStorageTransferJobResponse
	GetBody() *PutDataLakeStorageTransferJobResponseBody
}

type PutDataLakeStorageTransferJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDataLakeStorageTransferJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDataLakeStorageTransferJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeStorageTransferJobResponse) GoString() string {
	return s.String()
}

func (s *PutDataLakeStorageTransferJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDataLakeStorageTransferJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDataLakeStorageTransferJobResponse) GetBody() *PutDataLakeStorageTransferJobResponseBody {
	return s.Body
}

func (s *PutDataLakeStorageTransferJobResponse) SetHeaders(v map[string]*string) *PutDataLakeStorageTransferJobResponse {
	s.Headers = v
	return s
}

func (s *PutDataLakeStorageTransferJobResponse) SetStatusCode(v int32) *PutDataLakeStorageTransferJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDataLakeStorageTransferJobResponse) SetBody(v *PutDataLakeStorageTransferJobResponseBody) *PutDataLakeStorageTransferJobResponse {
	s.Body = v
	return s
}

func (s *PutDataLakeStorageTransferJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
