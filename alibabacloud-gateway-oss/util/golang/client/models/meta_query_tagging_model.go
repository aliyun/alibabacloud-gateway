// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryTagging interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetaQueryTagging
	GetKey() *string
	SetValue(v string) *MetaQueryTagging
	GetValue() *string
}

type MetaQueryTagging struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MetaQueryTagging) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryTagging) GoString() string {
	return s.String()
}

func (s *MetaQueryTagging) GetKey() *string {
	return s.Key
}

func (s *MetaQueryTagging) GetValue() *string {
	return s.Value
}

func (s *MetaQueryTagging) SetKey(v string) *MetaQueryTagging {
	s.Key = &v
	return s
}

func (s *MetaQueryTagging) SetValue(v string) *MetaQueryTagging {
	s.Value = &v
	return s
}

func (s *MetaQueryTagging) Validate() error {
	return dara.Validate(s)
}
