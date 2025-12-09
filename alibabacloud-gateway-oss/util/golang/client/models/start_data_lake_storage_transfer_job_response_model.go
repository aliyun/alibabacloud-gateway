// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataLakeStorageTransferJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDataLakeStorageTransferJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDataLakeStorageTransferJobResponse
	GetStatusCode() *int32
	SetBody(v *StartDataLakeStorageTransferJobResponseBody) *StartDataLakeStorageTransferJobResponse
	GetBody() *StartDataLakeStorageTransferJobResponseBody
}

type StartDataLakeStorageTransferJobResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDataLakeStorageTransferJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDataLakeStorageTransferJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDataLakeStorageTransferJobResponse) GoString() string {
	return s.String()
}

func (s *StartDataLakeStorageTransferJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDataLakeStorageTransferJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDataLakeStorageTransferJobResponse) GetBody() *StartDataLakeStorageTransferJobResponseBody {
	return s.Body
}

func (s *StartDataLakeStorageTransferJobResponse) SetHeaders(v map[string]*string) *StartDataLakeStorageTransferJobResponse {
	s.Headers = v
	return s
}

func (s *StartDataLakeStorageTransferJobResponse) SetStatusCode(v int32) *StartDataLakeStorageTransferJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDataLakeStorageTransferJobResponse) SetBody(v *StartDataLakeStorageTransferJobResponseBody) *StartDataLakeStorageTransferJobResponse {
	s.Body = v
	return s
}

func (s *StartDataLakeStorageTransferJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
