// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryUserMeta interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetaQueryUserMeta
	GetKey() *string
	SetValue(v string) *MetaQueryUserMeta
	GetValue() *string
}

type MetaQueryUserMeta struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MetaQueryUserMeta) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryUserMeta) GoString() string {
	return s.String()
}

func (s *MetaQueryUserMeta) GetKey() *string {
	return s.Key
}

func (s *MetaQueryUserMeta) GetValue() *string {
	return s.Value
}

func (s *MetaQueryUserMeta) SetKey(v string) *MetaQueryUserMeta {
	s.Key = &v
	return s
}

func (s *MetaQueryUserMeta) SetValue(v string) *MetaQueryUserMeta {
	s.Value = &v
	return s
}

func (s *MetaQueryUserMeta) Validate() error {
	return dara.Validate(s)
}
