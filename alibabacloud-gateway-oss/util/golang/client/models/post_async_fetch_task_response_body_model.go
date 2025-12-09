// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostAsyncFetchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncFetchTaskResult(v *AsyncFetchTaskResult) *PostAsyncFetchTaskResponseBody
	GetAsyncFetchTaskResult() *AsyncFetchTaskResult
}

type PostAsyncFetchTaskResponseBody struct {
	AsyncFetchTaskResult *AsyncFetchTaskResult `json:"AsyncFetchTaskResult,omitempty" xml:"AsyncFetchTaskResult,omitempty"`
}

func (s PostAsyncFetchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostAsyncFetchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PostAsyncFetchTaskResponseBody) GetAsyncFetchTaskResult() *AsyncFetchTaskResult {
	return s.AsyncFetchTaskResult
}

func (s *PostAsyncFetchTaskResponseBody) SetAsyncFetchTaskResult(v *AsyncFetchTaskResult) *PostAsyncFetchTaskResponseBody {
	s.AsyncFetchTaskResult = v
	return s
}

func (s *PostAsyncFetchTaskResponseBody) Validate() error {
	if s.AsyncFetchTaskResult != nil {
		if err := s.AsyncFetchTaskResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
