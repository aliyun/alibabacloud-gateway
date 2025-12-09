// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMarkerEntry interface {
	dara.Model
	String() string
	GoString() string
	SetIsLatest(v bool) *DeleteMarkerEntry
	GetIsLatest() *bool
	SetKey(v string) *DeleteMarkerEntry
	GetKey() *string
	SetLastModified(v string) *DeleteMarkerEntry
	GetLastModified() *string
	SetOwner(v *Owner) *DeleteMarkerEntry
	GetOwner() *Owner
	SetVersionId(v string) *DeleteMarkerEntry
	GetVersionId() *string
}

type DeleteMarkerEntry struct {
	IsLatest *bool   `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Owner        *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteMarkerEntry) String() string {
	return dara.Prettify(s)
}

func (s DeleteMarkerEntry) GoString() string {
	return s.String()
}

func (s *DeleteMarkerEntry) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *DeleteMarkerEntry) GetKey() *string {
	return s.Key
}

func (s *DeleteMarkerEntry) GetLastModified() *string {
	return s.LastModified
}

func (s *DeleteMarkerEntry) GetOwner() *Owner {
	return s.Owner
}

func (s *DeleteMarkerEntry) GetVersionId() *string {
	return s.VersionId
}

func (s *DeleteMarkerEntry) SetIsLatest(v bool) *DeleteMarkerEntry {
	s.IsLatest = &v
	return s
}

func (s *DeleteMarkerEntry) SetKey(v string) *DeleteMarkerEntry {
	s.Key = &v
	return s
}

func (s *DeleteMarkerEntry) SetLastModified(v string) *DeleteMarkerEntry {
	s.LastModified = &v
	return s
}

func (s *DeleteMarkerEntry) SetOwner(v *Owner) *DeleteMarkerEntry {
	s.Owner = v
	return s
}

func (s *DeleteMarkerEntry) SetVersionId(v string) *DeleteMarkerEntry {
	s.VersionId = &v
	return s
}

func (s *DeleteMarkerEntry) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}
