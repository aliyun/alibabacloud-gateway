// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectIdentifier interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ObjectIdentifier
	GetKey() *string
	SetVersionId(v string) *ObjectIdentifier
	GetVersionId() *string
}

type ObjectIdentifier struct {
	// This parameter is required.
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ObjectIdentifier) String() string {
	return dara.Prettify(s)
}

func (s ObjectIdentifier) GoString() string {
	return s.String()
}

func (s *ObjectIdentifier) GetKey() *string {
	return s.Key
}

func (s *ObjectIdentifier) GetVersionId() *string {
	return s.VersionId
}

func (s *ObjectIdentifier) SetKey(v string) *ObjectIdentifier {
	s.Key = &v
	return s
}

func (s *ObjectIdentifier) SetVersionId(v string) *ObjectIdentifier {
	s.VersionId = &v
	return s
}

func (s *ObjectIdentifier) Validate() error {
	return dara.Validate(s)
}
