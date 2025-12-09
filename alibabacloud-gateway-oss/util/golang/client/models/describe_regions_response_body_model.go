// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionInfoList(v *DescribeRegionsResponseBodyRegionInfoList) *DescribeRegionsResponseBody
	GetRegionInfoList() *DescribeRegionsResponseBodyRegionInfoList
}

type DescribeRegionsResponseBody struct {
	// The information about the regions.
	RegionInfoList *DescribeRegionsResponseBodyRegionInfoList `json:"RegionInfoList,omitempty" xml:"RegionInfoList,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegionInfoList() *DescribeRegionsResponseBodyRegionInfoList {
	return s.RegionInfoList
}

func (s *DescribeRegionsResponseBody) SetRegionInfoList(v *DescribeRegionsResponseBodyRegionInfoList) *DescribeRegionsResponseBody {
	s.RegionInfoList = v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.RegionInfoList != nil {
		if err := s.RegionInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionInfoList struct {
	// The information about the regions.
	RegionInfos []*RegionInfo `json:"RegionInfo,omitempty" xml:"RegionInfo,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionInfoList) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionInfoList) GetRegionInfos() []*RegionInfo {
	return s.RegionInfos
}

func (s *DescribeRegionsResponseBodyRegionInfoList) SetRegionInfos(v []*RegionInfo) *DescribeRegionsResponseBodyRegionInfoList {
	s.RegionInfos = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionInfoList) Validate() error {
	if s.RegionInfos != nil {
		for _, item := range s.RegionInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
