// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryOpenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v *MetaQueryOpenRequestFilters) *MetaQueryOpenRequest
	GetFilters() *MetaQueryOpenRequestFilters
}

type MetaQueryOpenRequest struct {
	Filters *MetaQueryOpenRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
}

func (s MetaQueryOpenRequest) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryOpenRequest) GoString() string {
	return s.String()
}

func (s *MetaQueryOpenRequest) GetFilters() *MetaQueryOpenRequestFilters {
	return s.Filters
}

func (s *MetaQueryOpenRequest) SetFilters(v *MetaQueryOpenRequestFilters) *MetaQueryOpenRequest {
	s.Filters = v
	return s
}

func (s *MetaQueryOpenRequest) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryOpenRequestFilters struct {
	Filter []*string `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s MetaQueryOpenRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryOpenRequestFilters) GoString() string {
	return s.String()
}

func (s *MetaQueryOpenRequestFilters) GetFilter() []*string {
	return s.Filter
}

func (s *MetaQueryOpenRequestFilters) SetFilter(v []*string) *MetaQueryOpenRequestFilters {
	s.Filter = v
	return s
}

func (s *MetaQueryOpenRequestFilters) Validate() error {
	return dara.Validate(s)
}
