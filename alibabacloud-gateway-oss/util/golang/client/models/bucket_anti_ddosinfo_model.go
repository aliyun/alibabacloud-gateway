// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketAntiDDOSInfo interface {
	dara.Model
	String() string
	GoString() string
	SetActiveTime(v int64) *BucketAntiDDOSInfo
	GetActiveTime() *int64
	SetBucket(v string) *BucketAntiDDOSInfo
	GetBucket() *string
	SetCnames(v *BucketAntiDDOSInfoCnames) *BucketAntiDDOSInfo
	GetCnames() *BucketAntiDDOSInfoCnames
	SetCtime(v int64) *BucketAntiDDOSInfo
	GetCtime() *int64
	SetInstanceId(v string) *BucketAntiDDOSInfo
	GetInstanceId() *string
	SetMtime(v int64) *BucketAntiDDOSInfo
	GetMtime() *int64
	SetOwner(v string) *BucketAntiDDOSInfo
	GetOwner() *string
	SetStatus(v string) *BucketAntiDDOSInfo
	GetStatus() *string
	SetType(v string) *BucketAntiDDOSInfo
	GetType() *string
}

type BucketAntiDDOSInfo struct {
	ActiveTime *int64                    `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	Bucket     *string                   `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Cnames     *BucketAntiDDOSInfoCnames `json:"Cnames,omitempty" xml:"Cnames,omitempty" type:"Struct"`
	Ctime      *int64                    `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	InstanceId *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mtime      *int64                    `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Owner      *string                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Status     *string                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BucketAntiDDOSInfo) String() string {
	return dara.Prettify(s)
}

func (s BucketAntiDDOSInfo) GoString() string {
	return s.String()
}

func (s *BucketAntiDDOSInfo) GetActiveTime() *int64 {
	return s.ActiveTime
}

func (s *BucketAntiDDOSInfo) GetBucket() *string {
	return s.Bucket
}

func (s *BucketAntiDDOSInfo) GetCnames() *BucketAntiDDOSInfoCnames {
	return s.Cnames
}

func (s *BucketAntiDDOSInfo) GetCtime() *int64 {
	return s.Ctime
}

func (s *BucketAntiDDOSInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BucketAntiDDOSInfo) GetMtime() *int64 {
	return s.Mtime
}

func (s *BucketAntiDDOSInfo) GetOwner() *string {
	return s.Owner
}

func (s *BucketAntiDDOSInfo) GetStatus() *string {
	return s.Status
}

func (s *BucketAntiDDOSInfo) GetType() *string {
	return s.Type
}

func (s *BucketAntiDDOSInfo) SetActiveTime(v int64) *BucketAntiDDOSInfo {
	s.ActiveTime = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetBucket(v string) *BucketAntiDDOSInfo {
	s.Bucket = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetCnames(v *BucketAntiDDOSInfoCnames) *BucketAntiDDOSInfo {
	s.Cnames = v
	return s
}

func (s *BucketAntiDDOSInfo) SetCtime(v int64) *BucketAntiDDOSInfo {
	s.Ctime = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetInstanceId(v string) *BucketAntiDDOSInfo {
	s.InstanceId = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetMtime(v int64) *BucketAntiDDOSInfo {
	s.Mtime = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetOwner(v string) *BucketAntiDDOSInfo {
	s.Owner = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetStatus(v string) *BucketAntiDDOSInfo {
	s.Status = &v
	return s
}

func (s *BucketAntiDDOSInfo) SetType(v string) *BucketAntiDDOSInfo {
	s.Type = &v
	return s
}

func (s *BucketAntiDDOSInfo) Validate() error {
	if s.Cnames != nil {
		if err := s.Cnames.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BucketAntiDDOSInfoCnames struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s BucketAntiDDOSInfoCnames) String() string {
	return dara.Prettify(s)
}

func (s BucketAntiDDOSInfoCnames) GoString() string {
	return s.String()
}

func (s *BucketAntiDDOSInfoCnames) GetDomain() []*string {
	return s.Domain
}

func (s *BucketAntiDDOSInfoCnames) SetDomain(v []*string) *BucketAntiDDOSInfoCnames {
	s.Domain = v
	return s
}

func (s *BucketAntiDDOSInfoCnames) Validate() error {
	return dara.Validate(s)
}
