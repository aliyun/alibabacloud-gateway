// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectSummary interface {
	dara.Model
	String() string
	GoString() string
	SetETag(v string) *ObjectSummary
	GetETag() *string
	SetKey(v string) *ObjectSummary
	GetKey() *string
	SetLastModified(v string) *ObjectSummary
	GetLastModified() *string
	SetOwner(v *Owner) *ObjectSummary
	GetOwner() *Owner
	SetRestoreInfo(v string) *ObjectSummary
	GetRestoreInfo() *string
	SetSize(v int64) *ObjectSummary
	GetSize() *int64
	SetStorageClass(v string) *ObjectSummary
	GetStorageClass() *string
	SetTransitionTime(v string) *ObjectSummary
	GetTransitionTime() *string
	SetType(v string) *ObjectSummary
	GetType() *string
}

type ObjectSummary struct {
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5A0FE1****
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// fun/test.jpg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2012-02-24T08:42:32.000Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Owner        *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// ongoing-request="true”
	RestoreInfo *string `json:"RestoreInfo,omitempty" xml:"RestoreInfo,omitempty"`
	// example:
	//
	// 344606
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2012-02-24T08:42:32.000Z
	TransitionTime *string `json:"TransitionTime,omitempty" xml:"TransitionTime,omitempty"`
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ObjectSummary) String() string {
	return dara.Prettify(s)
}

func (s ObjectSummary) GoString() string {
	return s.String()
}

func (s *ObjectSummary) GetETag() *string {
	return s.ETag
}

func (s *ObjectSummary) GetKey() *string {
	return s.Key
}

func (s *ObjectSummary) GetLastModified() *string {
	return s.LastModified
}

func (s *ObjectSummary) GetOwner() *Owner {
	return s.Owner
}

func (s *ObjectSummary) GetRestoreInfo() *string {
	return s.RestoreInfo
}

func (s *ObjectSummary) GetSize() *int64 {
	return s.Size
}

func (s *ObjectSummary) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ObjectSummary) GetTransitionTime() *string {
	return s.TransitionTime
}

func (s *ObjectSummary) GetType() *string {
	return s.Type
}

func (s *ObjectSummary) SetETag(v string) *ObjectSummary {
	s.ETag = &v
	return s
}

func (s *ObjectSummary) SetKey(v string) *ObjectSummary {
	s.Key = &v
	return s
}

func (s *ObjectSummary) SetLastModified(v string) *ObjectSummary {
	s.LastModified = &v
	return s
}

func (s *ObjectSummary) SetOwner(v *Owner) *ObjectSummary {
	s.Owner = v
	return s
}

func (s *ObjectSummary) SetRestoreInfo(v string) *ObjectSummary {
	s.RestoreInfo = &v
	return s
}

func (s *ObjectSummary) SetSize(v int64) *ObjectSummary {
	s.Size = &v
	return s
}

func (s *ObjectSummary) SetStorageClass(v string) *ObjectSummary {
	s.StorageClass = &v
	return s
}

func (s *ObjectSummary) SetTransitionTime(v string) *ObjectSummary {
	s.TransitionTime = &v
	return s
}

func (s *ObjectSummary) SetType(v string) *ObjectSummary {
	s.Type = &v
	return s
}

func (s *ObjectSummary) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}
