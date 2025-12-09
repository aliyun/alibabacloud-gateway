// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectVersion interface {
	dara.Model
	String() string
	GoString() string
	SetETag(v string) *ObjectVersion
	GetETag() *string
	SetIsLatest(v bool) *ObjectVersion
	GetIsLatest() *bool
	SetKey(v string) *ObjectVersion
	GetKey() *string
	SetLastModified(v string) *ObjectVersion
	GetLastModified() *string
	SetOwner(v *Owner) *ObjectVersion
	GetOwner() *Owner
	SetRestoreInfo(v string) *ObjectVersion
	GetRestoreInfo() *string
	SetSize(v int64) *ObjectVersion
	GetSize() *int64
	SetStorageClass(v string) *ObjectVersion
	GetStorageClass() *string
	SetTransitionTime(v string) *ObjectVersion
	GetTransitionTime() *string
	SetVersionId(v string) *ObjectVersion
	GetVersionId() *string
}

type ObjectVersion struct {
	// example:
	//
	// "250F8A0AE989679A22926A875F0A2****"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// false
	IsLatest *bool `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	// example:
	//
	// dic/test.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-04-09T07:27:28.000Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Owner        *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// ongoing-request="true"
	RestoreInfo *string `json:"RestoreInfo,omitempty" xml:"RestoreInfo,omitempty"`
	// example:
	//
	// 93731
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-04-09T07:27:28.000Z
	TransitionTime *string `json:"TransitionTime,omitempty" xml:"TransitionTime,omitempty"`
	// example:
	//
	// CAEQMxiBgMDNoP2D0BYiIDE3MWUxNzgxZDQxNTRiODI5OGYwZGMwNGY3MzZjN****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ObjectVersion) String() string {
	return dara.Prettify(s)
}

func (s ObjectVersion) GoString() string {
	return s.String()
}

func (s *ObjectVersion) GetETag() *string {
	return s.ETag
}

func (s *ObjectVersion) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *ObjectVersion) GetKey() *string {
	return s.Key
}

func (s *ObjectVersion) GetLastModified() *string {
	return s.LastModified
}

func (s *ObjectVersion) GetOwner() *Owner {
	return s.Owner
}

func (s *ObjectVersion) GetRestoreInfo() *string {
	return s.RestoreInfo
}

func (s *ObjectVersion) GetSize() *int64 {
	return s.Size
}

func (s *ObjectVersion) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ObjectVersion) GetTransitionTime() *string {
	return s.TransitionTime
}

func (s *ObjectVersion) GetVersionId() *string {
	return s.VersionId
}

func (s *ObjectVersion) SetETag(v string) *ObjectVersion {
	s.ETag = &v
	return s
}

func (s *ObjectVersion) SetIsLatest(v bool) *ObjectVersion {
	s.IsLatest = &v
	return s
}

func (s *ObjectVersion) SetKey(v string) *ObjectVersion {
	s.Key = &v
	return s
}

func (s *ObjectVersion) SetLastModified(v string) *ObjectVersion {
	s.LastModified = &v
	return s
}

func (s *ObjectVersion) SetOwner(v *Owner) *ObjectVersion {
	s.Owner = v
	return s
}

func (s *ObjectVersion) SetRestoreInfo(v string) *ObjectVersion {
	s.RestoreInfo = &v
	return s
}

func (s *ObjectVersion) SetSize(v int64) *ObjectVersion {
	s.Size = &v
	return s
}

func (s *ObjectVersion) SetStorageClass(v string) *ObjectVersion {
	s.StorageClass = &v
	return s
}

func (s *ObjectVersion) SetTransitionTime(v string) *ObjectVersion {
	s.TransitionTime = &v
	return s
}

func (s *ObjectVersion) SetVersionId(v string) *ObjectVersion {
	s.VersionId = &v
	return s
}

func (s *ObjectVersion) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}
