// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRegionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *ListUserRegionsResultRegions) *ListUserRegionsResult
	GetRegions() *ListUserRegionsResultRegions
}

type ListUserRegionsResult struct {
	Regions *ListUserRegionsResultRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s ListUserRegionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserRegionsResult) GoString() string {
	return s.String()
}

func (s *ListUserRegionsResult) GetRegions() *ListUserRegionsResultRegions {
	return s.Regions
}

func (s *ListUserRegionsResult) SetRegions(v *ListUserRegionsResultRegions) *ListUserRegionsResult {
	s.Regions = v
	return s
}

func (s *ListUserRegionsResult) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserRegionsResultRegions struct {
	Region []*string `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s ListUserRegionsResultRegions) String() string {
	return dara.Prettify(s)
}

func (s ListUserRegionsResultRegions) GoString() string {
	return s.String()
}

func (s *ListUserRegionsResultRegions) GetRegion() []*string {
	return s.Region
}

func (s *ListUserRegionsResultRegions) SetRegion(v []*string) *ListUserRegionsResultRegions {
	s.Region = v
	return s
}

func (s *ListUserRegionsResultRegions) Validate() error {
	return dara.Validate(s)
}
