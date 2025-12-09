// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserQosConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultQoSConfiguration(v *QoSConfigurationWithRemark) *UserQosConfiguration
	GetDefaultQoSConfiguration() *QoSConfigurationWithRemark
	SetExtranetDownloadBandwidth(v int64) *UserQosConfiguration
	GetExtranetDownloadBandwidth() *int64
	SetExtranetQps(v int64) *UserQosConfiguration
	GetExtranetQps() *int64
	SetExtranetUploadBandwidth(v int64) *UserQosConfiguration
	GetExtranetUploadBandwidth() *int64
	SetIntranetDownloadBandwidth(v int64) *UserQosConfiguration
	GetIntranetDownloadBandwidth() *int64
	SetIntranetQps(v int64) *UserQosConfiguration
	GetIntranetQps() *int64
	SetIntranetUploadBandwidth(v int64) *UserQosConfiguration
	GetIntranetUploadBandwidth() *int64
	SetRegion(v string) *UserQosConfiguration
	GetRegion() *string
	SetRemark(v int64) *UserQosConfiguration
	GetRemark() *int64
	SetTotalDownloadBandwidth(v int64) *UserQosConfiguration
	GetTotalDownloadBandwidth() *int64
	SetTotalQps(v int64) *UserQosConfiguration
	GetTotalQps() *int64
	SetTotalUploadBandwidth(v int64) *UserQosConfiguration
	GetTotalUploadBandwidth() *int64
}

type UserQosConfiguration struct {
	DefaultQoSConfiguration *QoSConfigurationWithRemark `json:"DefaultQoSConfiguration,omitempty" xml:"DefaultQoSConfiguration,omitempty"`
	// example:
	//
	// 2
	ExtranetDownloadBandwidth *int64 `json:"ExtranetDownloadBandwidth,omitempty" xml:"ExtranetDownloadBandwidth,omitempty"`
	// example:
	//
	// 20000
	ExtranetQps *int64 `json:"ExtranetQps,omitempty" xml:"ExtranetQps,omitempty"`
	// example:
	//
	// 2
	ExtranetUploadBandwidth *int64 `json:"ExtranetUploadBandwidth,omitempty" xml:"ExtranetUploadBandwidth,omitempty"`
	// example:
	//
	// 3
	IntranetDownloadBandwidth *int64 `json:"IntranetDownloadBandwidth,omitempty" xml:"IntranetDownloadBandwidth,omitempty"`
	// example:
	//
	// 20000
	IntranetQps *int64 `json:"IntranetQps,omitempty" xml:"IntranetQps,omitempty"`
	// example:
	//
	// 1
	IntranetUploadBandwidth *int64 `json:"IntranetUploadBandwidth,omitempty" xml:"IntranetUploadBandwidth,omitempty"`
	// example:
	//
	// oss-cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 备注
	Remark *int64 `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 5
	TotalDownloadBandwidth *int64 `json:"TotalDownloadBandwidth,omitempty" xml:"TotalDownloadBandwidth,omitempty"`
	// example:
	//
	// 20000
	TotalQps *int64 `json:"TotalQps,omitempty" xml:"TotalQps,omitempty"`
	// example:
	//
	// 2
	TotalUploadBandwidth *int64 `json:"TotalUploadBandwidth,omitempty" xml:"TotalUploadBandwidth,omitempty"`
}

func (s UserQosConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UserQosConfiguration) GoString() string {
	return s.String()
}

func (s *UserQosConfiguration) GetDefaultQoSConfiguration() *QoSConfigurationWithRemark {
	return s.DefaultQoSConfiguration
}

func (s *UserQosConfiguration) GetExtranetDownloadBandwidth() *int64 {
	return s.ExtranetDownloadBandwidth
}

func (s *UserQosConfiguration) GetExtranetQps() *int64 {
	return s.ExtranetQps
}

func (s *UserQosConfiguration) GetExtranetUploadBandwidth() *int64 {
	return s.ExtranetUploadBandwidth
}

func (s *UserQosConfiguration) GetIntranetDownloadBandwidth() *int64 {
	return s.IntranetDownloadBandwidth
}

func (s *UserQosConfiguration) GetIntranetQps() *int64 {
	return s.IntranetQps
}

func (s *UserQosConfiguration) GetIntranetUploadBandwidth() *int64 {
	return s.IntranetUploadBandwidth
}

func (s *UserQosConfiguration) GetRegion() *string {
	return s.Region
}

func (s *UserQosConfiguration) GetRemark() *int64 {
	return s.Remark
}

func (s *UserQosConfiguration) GetTotalDownloadBandwidth() *int64 {
	return s.TotalDownloadBandwidth
}

func (s *UserQosConfiguration) GetTotalQps() *int64 {
	return s.TotalQps
}

func (s *UserQosConfiguration) GetTotalUploadBandwidth() *int64 {
	return s.TotalUploadBandwidth
}

func (s *UserQosConfiguration) SetDefaultQoSConfiguration(v *QoSConfigurationWithRemark) *UserQosConfiguration {
	s.DefaultQoSConfiguration = v
	return s
}

func (s *UserQosConfiguration) SetExtranetDownloadBandwidth(v int64) *UserQosConfiguration {
	s.ExtranetDownloadBandwidth = &v
	return s
}

func (s *UserQosConfiguration) SetExtranetQps(v int64) *UserQosConfiguration {
	s.ExtranetQps = &v
	return s
}

func (s *UserQosConfiguration) SetExtranetUploadBandwidth(v int64) *UserQosConfiguration {
	s.ExtranetUploadBandwidth = &v
	return s
}

func (s *UserQosConfiguration) SetIntranetDownloadBandwidth(v int64) *UserQosConfiguration {
	s.IntranetDownloadBandwidth = &v
	return s
}

func (s *UserQosConfiguration) SetIntranetQps(v int64) *UserQosConfiguration {
	s.IntranetQps = &v
	return s
}

func (s *UserQosConfiguration) SetIntranetUploadBandwidth(v int64) *UserQosConfiguration {
	s.IntranetUploadBandwidth = &v
	return s
}

func (s *UserQosConfiguration) SetRegion(v string) *UserQosConfiguration {
	s.Region = &v
	return s
}

func (s *UserQosConfiguration) SetRemark(v int64) *UserQosConfiguration {
	s.Remark = &v
	return s
}

func (s *UserQosConfiguration) SetTotalDownloadBandwidth(v int64) *UserQosConfiguration {
	s.TotalDownloadBandwidth = &v
	return s
}

func (s *UserQosConfiguration) SetTotalQps(v int64) *UserQosConfiguration {
	s.TotalQps = &v
	return s
}

func (s *UserQosConfiguration) SetTotalUploadBandwidth(v int64) *UserQosConfiguration {
	s.TotalUploadBandwidth = &v
	return s
}

func (s *UserQosConfiguration) Validate() error {
	if s.DefaultQoSConfiguration != nil {
		if err := s.DefaultQoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
