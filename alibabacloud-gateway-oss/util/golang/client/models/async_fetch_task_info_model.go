// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncFetchTaskInfo interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *AsyncFetchTaskInfo
	GetErrorMsg() *string
	SetState(v string) *AsyncFetchTaskInfo
	GetState() *string
	SetTaskId(v string) *AsyncFetchTaskInfo
	GetTaskId() *string
	SetTaskInfo(v *AsyncFetchTaskConfiguration) *AsyncFetchTaskInfo
	GetTaskInfo() *AsyncFetchTaskConfiguration
}

type AsyncFetchTaskInfo struct {
	// example:
	//
	// FileNotFound
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// abc
	TaskId   *string                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInfo *AsyncFetchTaskConfiguration `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty"`
}

func (s AsyncFetchTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s AsyncFetchTaskInfo) GoString() string {
	return s.String()
}

func (s *AsyncFetchTaskInfo) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *AsyncFetchTaskInfo) GetState() *string {
	return s.State
}

func (s *AsyncFetchTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncFetchTaskInfo) GetTaskInfo() *AsyncFetchTaskConfiguration {
	return s.TaskInfo
}

func (s *AsyncFetchTaskInfo) SetErrorMsg(v string) *AsyncFetchTaskInfo {
	s.ErrorMsg = &v
	return s
}

func (s *AsyncFetchTaskInfo) SetState(v string) *AsyncFetchTaskInfo {
	s.State = &v
	return s
}

func (s *AsyncFetchTaskInfo) SetTaskId(v string) *AsyncFetchTaskInfo {
	s.TaskId = &v
	return s
}

func (s *AsyncFetchTaskInfo) SetTaskInfo(v *AsyncFetchTaskConfiguration) *AsyncFetchTaskInfo {
	s.TaskInfo = v
	return s
}

func (s *AsyncFetchTaskInfo) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
