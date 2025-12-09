// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeStorageTransferJob interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorRoleId(v string) *CreateDataLakeStorageTransferJob
	GetExecutorRoleId() *string
	SetIncludes(v []*string) *CreateDataLakeStorageTransferJob
	GetIncludes() []*string
	SetLogBaseDir(v string) *CreateDataLakeStorageTransferJob
	GetLogBaseDir() *string
	SetNeedVerify(v bool) *CreateDataLakeStorageTransferJob
	GetNeedVerify() *bool
	SetTag(v string) *CreateDataLakeStorageTransferJob
	GetTag() *string
}

type CreateDataLakeStorageTransferJob struct {
	// example:
	//
	// AliyunOSSRole
	ExecutorRoleId *string   `json:"ExecutorRoleId,omitempty" xml:"ExecutorRoleId,omitempty"`
	Includes       []*string `json:"Includes,omitempty" xml:"Includes,omitempty" type:"Repeated"`
	// example:
	//
	// log
	LogBaseDir *string `json:"LogBaseDir,omitempty" xml:"LogBaseDir,omitempty"`
	// example:
	//
	// false
	NeedVerify *bool `json:"NeedVerify,omitempty" xml:"NeedVerify,omitempty"`
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateDataLakeStorageTransferJob) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeStorageTransferJob) GoString() string {
	return s.String()
}

func (s *CreateDataLakeStorageTransferJob) GetExecutorRoleId() *string {
	return s.ExecutorRoleId
}

func (s *CreateDataLakeStorageTransferJob) GetIncludes() []*string {
	return s.Includes
}

func (s *CreateDataLakeStorageTransferJob) GetLogBaseDir() *string {
	return s.LogBaseDir
}

func (s *CreateDataLakeStorageTransferJob) GetNeedVerify() *bool {
	return s.NeedVerify
}

func (s *CreateDataLakeStorageTransferJob) GetTag() *string {
	return s.Tag
}

func (s *CreateDataLakeStorageTransferJob) SetExecutorRoleId(v string) *CreateDataLakeStorageTransferJob {
	s.ExecutorRoleId = &v
	return s
}

func (s *CreateDataLakeStorageTransferJob) SetIncludes(v []*string) *CreateDataLakeStorageTransferJob {
	s.Includes = v
	return s
}

func (s *CreateDataLakeStorageTransferJob) SetLogBaseDir(v string) *CreateDataLakeStorageTransferJob {
	s.LogBaseDir = &v
	return s
}

func (s *CreateDataLakeStorageTransferJob) SetNeedVerify(v bool) *CreateDataLakeStorageTransferJob {
	s.NeedVerify = &v
	return s
}

func (s *CreateDataLakeStorageTransferJob) SetTag(v string) *CreateDataLakeStorageTransferJob {
	s.Tag = &v
	return s
}

func (s *CreateDataLakeStorageTransferJob) Validate() error {
	return dara.Validate(s)
}
