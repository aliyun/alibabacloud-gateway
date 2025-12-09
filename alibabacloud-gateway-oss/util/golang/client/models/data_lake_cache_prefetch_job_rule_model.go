// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeCachePrefetchJobRule interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixFilter(v *DataLakeCachePrefetchJobRulePrefixFilter) *DataLakeCachePrefetchJobRule
	GetPrefixFilter() *DataLakeCachePrefetchJobRulePrefixFilter
	SetTag(v string) *DataLakeCachePrefetchJobRule
	GetTag() *string
}

type DataLakeCachePrefetchJobRule struct {
	PrefixFilter *DataLakeCachePrefetchJobRulePrefixFilter `json:"PrefixFilter,omitempty" xml:"PrefixFilter,omitempty" type:"Struct"`
	Tag          *string                                   `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DataLakeCachePrefetchJobRule) String() string {
	return dara.Prettify(s)
}

func (s DataLakeCachePrefetchJobRule) GoString() string {
	return s.String()
}

func (s *DataLakeCachePrefetchJobRule) GetPrefixFilter() *DataLakeCachePrefetchJobRulePrefixFilter {
	return s.PrefixFilter
}

func (s *DataLakeCachePrefetchJobRule) GetTag() *string {
	return s.Tag
}

func (s *DataLakeCachePrefetchJobRule) SetPrefixFilter(v *DataLakeCachePrefetchJobRulePrefixFilter) *DataLakeCachePrefetchJobRule {
	s.PrefixFilter = v
	return s
}

func (s *DataLakeCachePrefetchJobRule) SetTag(v string) *DataLakeCachePrefetchJobRule {
	s.Tag = &v
	return s
}

func (s *DataLakeCachePrefetchJobRule) Validate() error {
	if s.PrefixFilter != nil {
		if err := s.PrefixFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataLakeCachePrefetchJobRulePrefixFilter struct {
	Excludes *DataLakeCachePrefetchJobRulePrefixFilterExcludes `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Struct"`
	Includes *DataLakeCachePrefetchJobRulePrefixFilterIncludes `json:"Includes,omitempty" xml:"Includes,omitempty" type:"Struct"`
}

func (s DataLakeCachePrefetchJobRulePrefixFilter) String() string {
	return dara.Prettify(s)
}

func (s DataLakeCachePrefetchJobRulePrefixFilter) GoString() string {
	return s.String()
}

func (s *DataLakeCachePrefetchJobRulePrefixFilter) GetExcludes() *DataLakeCachePrefetchJobRulePrefixFilterExcludes {
	return s.Excludes
}

func (s *DataLakeCachePrefetchJobRulePrefixFilter) GetIncludes() *DataLakeCachePrefetchJobRulePrefixFilterIncludes {
	return s.Includes
}

func (s *DataLakeCachePrefetchJobRulePrefixFilter) SetExcludes(v *DataLakeCachePrefetchJobRulePrefixFilterExcludes) *DataLakeCachePrefetchJobRulePrefixFilter {
	s.Excludes = v
	return s
}

func (s *DataLakeCachePrefetchJobRulePrefixFilter) SetIncludes(v *DataLakeCachePrefetchJobRulePrefixFilterIncludes) *DataLakeCachePrefetchJobRulePrefixFilter {
	s.Includes = v
	return s
}

func (s *DataLakeCachePrefetchJobRulePrefixFilter) Validate() error {
	if s.Excludes != nil {
		if err := s.Excludes.Validate(); err != nil {
			return err
		}
	}
	if s.Includes != nil {
		if err := s.Includes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataLakeCachePrefetchJobRulePrefixFilterExcludes struct {
	Exclude []*string `json:"Exclude,omitempty" xml:"Exclude,omitempty" type:"Repeated"`
}

func (s DataLakeCachePrefetchJobRulePrefixFilterExcludes) String() string {
	return dara.Prettify(s)
}

func (s DataLakeCachePrefetchJobRulePrefixFilterExcludes) GoString() string {
	return s.String()
}

func (s *DataLakeCachePrefetchJobRulePrefixFilterExcludes) GetExclude() []*string {
	return s.Exclude
}

func (s *DataLakeCachePrefetchJobRulePrefixFilterExcludes) SetExclude(v []*string) *DataLakeCachePrefetchJobRulePrefixFilterExcludes {
	s.Exclude = v
	return s
}

func (s *DataLakeCachePrefetchJobRulePrefixFilterExcludes) Validate() error {
	return dara.Validate(s)
}

type DataLakeCachePrefetchJobRulePrefixFilterIncludes struct {
	Include []*string `json:"Include,omitempty" xml:"Include,omitempty" type:"Repeated"`
}

func (s DataLakeCachePrefetchJobRulePrefixFilterIncludes) String() string {
	return dara.Prettify(s)
}

func (s DataLakeCachePrefetchJobRulePrefixFilterIncludes) GoString() string {
	return s.String()
}

func (s *DataLakeCachePrefetchJobRulePrefixFilterIncludes) GetInclude() []*string {
	return s.Include
}

func (s *DataLakeCachePrefetchJobRulePrefixFilterIncludes) SetInclude(v []*string) *DataLakeCachePrefetchJobRulePrefixFilterIncludes {
	s.Include = v
	return s
}

func (s *DataLakeCachePrefetchJobRulePrefixFilterIncludes) Validate() error {
	return dara.Validate(s)
}
