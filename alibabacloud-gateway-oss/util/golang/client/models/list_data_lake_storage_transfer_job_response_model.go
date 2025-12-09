// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeStorageTransferJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeStorageTransferJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeStorageTransferJobResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeStorageTransferJobResponseBody) *ListDataLakeStorageTransferJobResponse
	GetBody() *ListDataLakeStorageTransferJobResponseBody
}

type ListDataLakeStorageTransferJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeStorageTransferJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeStorageTransferJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeStorageTransferJobResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeStorageTransferJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeStorageTransferJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeStorageTransferJobResponse) GetBody() *ListDataLakeStorageTransferJobResponseBody {
	return s.Body
}

func (s *ListDataLakeStorageTransferJobResponse) SetHeaders(v map[string]*string) *ListDataLakeStorageTransferJobResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeStorageTransferJobResponse) SetStatusCode(v int32) *ListDataLakeStorageTransferJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeStorageTransferJobResponse) SetBody(v *ListDataLakeStorageTransferJobResponseBody) *ListDataLakeStorageTransferJobResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeStorageTransferJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
