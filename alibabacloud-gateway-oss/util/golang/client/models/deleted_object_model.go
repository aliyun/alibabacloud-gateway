// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletedObject interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMarker(v bool) *DeletedObject
	GetDeleteMarker() *bool
	SetDeleteMarkerVersionId(v string) *DeletedObject
	GetDeleteMarkerVersionId() *string
	SetKey(v string) *DeletedObject
	GetKey() *string
	SetVersionId(v string) *DeletedObject
	GetVersionId() *string
}

type DeletedObject struct {
	DeleteMarker          *bool   `json:"DeleteMarker,omitempty" xml:"DeleteMarker,omitempty"`
	DeleteMarkerVersionId *string `json:"DeleteMarkerVersionId,omitempty" xml:"DeleteMarkerVersionId,omitempty"`
	Key                   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	VersionId             *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeletedObject) String() string {
	return dara.Prettify(s)
}

func (s DeletedObject) GoString() string {
	return s.String()
}

func (s *DeletedObject) GetDeleteMarker() *bool {
	return s.DeleteMarker
}

func (s *DeletedObject) GetDeleteMarkerVersionId() *string {
	return s.DeleteMarkerVersionId
}

func (s *DeletedObject) GetKey() *string {
	return s.Key
}

func (s *DeletedObject) GetVersionId() *string {
	return s.VersionId
}

func (s *DeletedObject) SetDeleteMarker(v bool) *DeletedObject {
	s.DeleteMarker = &v
	return s
}

func (s *DeletedObject) SetDeleteMarkerVersionId(v string) *DeletedObject {
	s.DeleteMarkerVersionId = &v
	return s
}

func (s *DeletedObject) SetKey(v string) *DeletedObject {
	s.Key = &v
	return s
}

func (s *DeletedObject) SetVersionId(v string) *DeletedObject {
	s.VersionId = &v
	return s
}

func (s *DeletedObject) Validate() error {
	return dara.Validate(s)
}
