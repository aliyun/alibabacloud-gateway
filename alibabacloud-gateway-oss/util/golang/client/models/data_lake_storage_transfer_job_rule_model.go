// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeStorageTransferJobRule interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorRoleId(v string) *DataLakeStorageTransferJobRule
	GetExecutorRoleId() *string
	SetLogBaseDir(v string) *DataLakeStorageTransferJobRule
	GetLogBaseDir() *string
	SetNeedVerify(v bool) *DataLakeStorageTransferJobRule
	GetNeedVerify() *bool
	SetPrefixFilter(v *DataLakeStorageTransferJobRulePrefixFilter) *DataLakeStorageTransferJobRule
	GetPrefixFilter() *DataLakeStorageTransferJobRulePrefixFilter
	SetTag(v string) *DataLakeStorageTransferJobRule
	GetTag() *string
}

type DataLakeStorageTransferJobRule struct {
	// example:
	//
	// AliyunOSSRole
	ExecutorRoleId *string `json:"ExecutorRoleId,omitempty" xml:"ExecutorRoleId,omitempty"`
	// example:
	//
	// log/
	LogBaseDir *string `json:"LogBaseDir,omitempty" xml:"LogBaseDir,omitempty"`
	// example:
	//
	// false
	NeedVerify   *bool                                       `json:"NeedVerify,omitempty" xml:"NeedVerify,omitempty"`
	PrefixFilter *DataLakeStorageTransferJobRulePrefixFilter `json:"PrefixFilter,omitempty" xml:"PrefixFilter,omitempty" type:"Struct"`
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DataLakeStorageTransferJobRule) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobRule) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobRule) GetExecutorRoleId() *string {
	return s.ExecutorRoleId
}

func (s *DataLakeStorageTransferJobRule) GetLogBaseDir() *string {
	return s.LogBaseDir
}

func (s *DataLakeStorageTransferJobRule) GetNeedVerify() *bool {
	return s.NeedVerify
}

func (s *DataLakeStorageTransferJobRule) GetPrefixFilter() *DataLakeStorageTransferJobRulePrefixFilter {
	return s.PrefixFilter
}

func (s *DataLakeStorageTransferJobRule) GetTag() *string {
	return s.Tag
}

func (s *DataLakeStorageTransferJobRule) SetExecutorRoleId(v string) *DataLakeStorageTransferJobRule {
	s.ExecutorRoleId = &v
	return s
}

func (s *DataLakeStorageTransferJobRule) SetLogBaseDir(v string) *DataLakeStorageTransferJobRule {
	s.LogBaseDir = &v
	return s
}

func (s *DataLakeStorageTransferJobRule) SetNeedVerify(v bool) *DataLakeStorageTransferJobRule {
	s.NeedVerify = &v
	return s
}

func (s *DataLakeStorageTransferJobRule) SetPrefixFilter(v *DataLakeStorageTransferJobRulePrefixFilter) *DataLakeStorageTransferJobRule {
	s.PrefixFilter = v
	return s
}

func (s *DataLakeStorageTransferJobRule) SetTag(v string) *DataLakeStorageTransferJobRule {
	s.Tag = &v
	return s
}

func (s *DataLakeStorageTransferJobRule) Validate() error {
	if s.PrefixFilter != nil {
		if err := s.PrefixFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataLakeStorageTransferJobRulePrefixFilter struct {
	Includes *DataLakeStorageTransferJobRulePrefixFilterIncludes `json:"Includes,omitempty" xml:"Includes,omitempty" type:"Struct"`
}

func (s DataLakeStorageTransferJobRulePrefixFilter) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobRulePrefixFilter) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobRulePrefixFilter) GetIncludes() *DataLakeStorageTransferJobRulePrefixFilterIncludes {
	return s.Includes
}

func (s *DataLakeStorageTransferJobRulePrefixFilter) SetIncludes(v *DataLakeStorageTransferJobRulePrefixFilterIncludes) *DataLakeStorageTransferJobRulePrefixFilter {
	s.Includes = v
	return s
}

func (s *DataLakeStorageTransferJobRulePrefixFilter) Validate() error {
	if s.Includes != nil {
		if err := s.Includes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataLakeStorageTransferJobRulePrefixFilterIncludes struct {
	Include []*string `json:"Include,omitempty" xml:"Include,omitempty" type:"Repeated"`
}

func (s DataLakeStorageTransferJobRulePrefixFilterIncludes) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobRulePrefixFilterIncludes) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobRulePrefixFilterIncludes) GetInclude() []*string {
	return s.Include
}

func (s *DataLakeStorageTransferJobRulePrefixFilterIncludes) SetInclude(v []*string) *DataLakeStorageTransferJobRulePrefixFilterIncludes {
	s.Include = v
	return s
}

func (s *DataLakeStorageTransferJobRulePrefixFilterIncludes) Validate() error {
	return dara.Validate(s)
}
