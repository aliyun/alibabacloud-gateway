// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryAggregation interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *MetaQueryAggregation
	GetField() *string
	SetOperation(v string) *MetaQueryAggregation
	GetOperation() *string
}

type MetaQueryAggregation struct {
	Field     *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s MetaQueryAggregation) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryAggregation) GoString() string {
	return s.String()
}

func (s *MetaQueryAggregation) GetField() *string {
	return s.Field
}

func (s *MetaQueryAggregation) GetOperation() *string {
	return s.Operation
}

func (s *MetaQueryAggregation) SetField(v string) *MetaQueryAggregation {
	s.Field = &v
	return s
}

func (s *MetaQueryAggregation) SetOperation(v string) *MetaQueryAggregation {
	s.Operation = &v
	return s
}

func (s *MetaQueryAggregation) Validate() error {
	return dara.Validate(s)
}
