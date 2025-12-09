// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeStorageTransferJobHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeStorageTransferJobHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeStorageTransferJobHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeStorageTransferJobHistoryResponseBody) *ListDataLakeStorageTransferJobHistoryResponse
	GetBody() *ListDataLakeStorageTransferJobHistoryResponseBody
}

type ListDataLakeStorageTransferJobHistoryResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeStorageTransferJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeStorageTransferJobHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeStorageTransferJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) GetBody() *ListDataLakeStorageTransferJobHistoryResponseBody {
	return s.Body
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) SetHeaders(v map[string]*string) *ListDataLakeStorageTransferJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) SetStatusCode(v int32) *ListDataLakeStorageTransferJobHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) SetBody(v *ListDataLakeStorageTransferJobHistoryResponseBody) *ListDataLakeStorageTransferJobHistoryResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeStorageTransferJobHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
