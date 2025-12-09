// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncFetchTaskResult interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *AsyncFetchTaskResult
	GetTaskId() *string
}

type AsyncFetchTaskResult struct {
	// example:
	//
	// abc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsyncFetchTaskResult) String() string {
	return dara.Prettify(s)
}

func (s AsyncFetchTaskResult) GoString() string {
	return s.String()
}

func (s *AsyncFetchTaskResult) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncFetchTaskResult) SetTaskId(v string) *AsyncFetchTaskResult {
	s.TaskId = &v
	return s
}

func (s *AsyncFetchTaskResult) Validate() error {
	return dara.Validate(s)
}
