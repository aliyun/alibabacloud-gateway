// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserAntiDDOSInfo interface {
	dara.Model
	String() string
	GoString() string
	SetActiveTime(v int64) *UserAntiDDOSInfo
	GetActiveTime() *int64
	SetCtime(v int64) *UserAntiDDOSInfo
	GetCtime() *int64
	SetInstanceId(v string) *UserAntiDDOSInfo
	GetInstanceId() *string
	SetMtime(v int64) *UserAntiDDOSInfo
	GetMtime() *int64
	SetOwner(v string) *UserAntiDDOSInfo
	GetOwner() *string
	SetStatus(v string) *UserAntiDDOSInfo
	GetStatus() *string
}

type UserAntiDDOSInfo struct {
	ActiveTime *int64  `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	Ctime      *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mtime      *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UserAntiDDOSInfo) String() string {
	return dara.Prettify(s)
}

func (s UserAntiDDOSInfo) GoString() string {
	return s.String()
}

func (s *UserAntiDDOSInfo) GetActiveTime() *int64 {
	return s.ActiveTime
}

func (s *UserAntiDDOSInfo) GetCtime() *int64 {
	return s.Ctime
}

func (s *UserAntiDDOSInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UserAntiDDOSInfo) GetMtime() *int64 {
	return s.Mtime
}

func (s *UserAntiDDOSInfo) GetOwner() *string {
	return s.Owner
}

func (s *UserAntiDDOSInfo) GetStatus() *string {
	return s.Status
}

func (s *UserAntiDDOSInfo) SetActiveTime(v int64) *UserAntiDDOSInfo {
	s.ActiveTime = &v
	return s
}

func (s *UserAntiDDOSInfo) SetCtime(v int64) *UserAntiDDOSInfo {
	s.Ctime = &v
	return s
}

func (s *UserAntiDDOSInfo) SetInstanceId(v string) *UserAntiDDOSInfo {
	s.InstanceId = &v
	return s
}

func (s *UserAntiDDOSInfo) SetMtime(v int64) *UserAntiDDOSInfo {
	s.Mtime = &v
	return s
}

func (s *UserAntiDDOSInfo) SetOwner(v string) *UserAntiDDOSInfo {
	s.Owner = &v
	return s
}

func (s *UserAntiDDOSInfo) SetStatus(v string) *UserAntiDDOSInfo {
	s.Status = &v
	return s
}

func (s *UserAntiDDOSInfo) Validate() error {
	return dara.Validate(s)
}
