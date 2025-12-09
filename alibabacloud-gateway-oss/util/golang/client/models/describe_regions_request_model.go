// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v string) *DescribeRegionsRequest
	GetRegions() *string
}

type DescribeRegionsRequest struct {
	// The region ID of the request.
	//
	// example:
	//
	// oss-cn-hangzhou
	Regions *string `json:"regions,omitempty" xml:"regions,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetRegions() *string {
	return s.Regions
}

func (s *DescribeRegionsRequest) SetRegions(v string) *DescribeRegionsRequest {
	s.Regions = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
