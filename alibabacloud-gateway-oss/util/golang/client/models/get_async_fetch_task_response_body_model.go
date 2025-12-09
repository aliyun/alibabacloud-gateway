// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncFetchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncFetchTaskInfo(v *AsyncFetchTaskInfo) *GetAsyncFetchTaskResponseBody
	GetAsyncFetchTaskInfo() *AsyncFetchTaskInfo
}

type GetAsyncFetchTaskResponseBody struct {
	AsyncFetchTaskInfo *AsyncFetchTaskInfo `json:"AsyncFetchTaskInfo,omitempty" xml:"AsyncFetchTaskInfo,omitempty"`
}

func (s GetAsyncFetchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncFetchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncFetchTaskResponseBody) GetAsyncFetchTaskInfo() *AsyncFetchTaskInfo {
	return s.AsyncFetchTaskInfo
}

func (s *GetAsyncFetchTaskResponseBody) SetAsyncFetchTaskInfo(v *AsyncFetchTaskInfo) *GetAsyncFetchTaskResponseBody {
	s.AsyncFetchTaskInfo = v
	return s
}

func (s *GetAsyncFetchTaskResponseBody) Validate() error {
	if s.AsyncFetchTaskInfo != nil {
		if err := s.AsyncFetchTaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
