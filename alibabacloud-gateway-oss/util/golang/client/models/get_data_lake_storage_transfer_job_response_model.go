// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeStorageTransferJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakeStorageTransferJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakeStorageTransferJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakeStorageTransferJobResponseBody) *GetDataLakeStorageTransferJobResponse
	GetBody() *GetDataLakeStorageTransferJobResponseBody
}

type GetDataLakeStorageTransferJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakeStorageTransferJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakeStorageTransferJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeStorageTransferJobResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakeStorageTransferJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakeStorageTransferJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakeStorageTransferJobResponse) GetBody() *GetDataLakeStorageTransferJobResponseBody {
	return s.Body
}

func (s *GetDataLakeStorageTransferJobResponse) SetHeaders(v map[string]*string) *GetDataLakeStorageTransferJobResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakeStorageTransferJobResponse) SetStatusCode(v int32) *GetDataLakeStorageTransferJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakeStorageTransferJobResponse) SetBody(v *GetDataLakeStorageTransferJobResponseBody) *GetDataLakeStorageTransferJobResponse {
	s.Body = v
	return s
}

func (s *GetDataLakeStorageTransferJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
