// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataAcceleratorConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratePaths(v *AcceleratePaths) *DataAcceleratorConfiguration
	GetAcceleratePaths() *AcceleratePaths
	SetAvailableZone(v string) *DataAcceleratorConfiguration
	GetAvailableZone() *string
	SetQuota(v string) *DataAcceleratorConfiguration
	GetQuota() *string
}

type DataAcceleratorConfiguration struct {
	AcceleratePaths *AcceleratePaths `json:"AcceleratePaths,omitempty" xml:"AcceleratePaths,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	AvailableZone *string `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty"`
	// example:
	//
	// 102400
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s DataAcceleratorConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DataAcceleratorConfiguration) GoString() string {
	return s.String()
}

func (s *DataAcceleratorConfiguration) GetAcceleratePaths() *AcceleratePaths {
	return s.AcceleratePaths
}

func (s *DataAcceleratorConfiguration) GetAvailableZone() *string {
	return s.AvailableZone
}

func (s *DataAcceleratorConfiguration) GetQuota() *string {
	return s.Quota
}

func (s *DataAcceleratorConfiguration) SetAcceleratePaths(v *AcceleratePaths) *DataAcceleratorConfiguration {
	s.AcceleratePaths = v
	return s
}

func (s *DataAcceleratorConfiguration) SetAvailableZone(v string) *DataAcceleratorConfiguration {
	s.AvailableZone = &v
	return s
}

func (s *DataAcceleratorConfiguration) SetQuota(v string) *DataAcceleratorConfiguration {
	s.Quota = &v
	return s
}

func (s *DataAcceleratorConfiguration) Validate() error {
	if s.AcceleratePaths != nil {
		if err := s.AcceleratePaths.Validate(); err != nil {
			return err
		}
	}
	return nil
}
