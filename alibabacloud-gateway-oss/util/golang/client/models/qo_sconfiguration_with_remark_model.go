// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQoSConfigurationWithRemark interface {
	dara.Model
	String() string
	GoString() string
	SetExtranetDownloadBandwidth(v int64) *QoSConfigurationWithRemark
	GetExtranetDownloadBandwidth() *int64
	SetExtranetQps(v int64) *QoSConfigurationWithRemark
	GetExtranetQps() *int64
	SetExtranetUploadBandwidth(v int64) *QoSConfigurationWithRemark
	GetExtranetUploadBandwidth() *int64
	SetIntranetDownloadBandwidth(v int64) *QoSConfigurationWithRemark
	GetIntranetDownloadBandwidth() *int64
	SetIntranetQps(v int64) *QoSConfigurationWithRemark
	GetIntranetQps() *int64
	SetIntranetUploadBandwidth(v int64) *QoSConfigurationWithRemark
	GetIntranetUploadBandwidth() *int64
	SetRemark(v int64) *QoSConfigurationWithRemark
	GetRemark() *int64
	SetTotalDownloadBandwidth(v int64) *QoSConfigurationWithRemark
	GetTotalDownloadBandwidth() *int64
	SetTotalQps(v int64) *QoSConfigurationWithRemark
	GetTotalQps() *int64
	SetTotalUploadBandwidth(v int64) *QoSConfigurationWithRemark
	GetTotalUploadBandwidth() *int64
}

type QoSConfigurationWithRemark struct {
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

func (s QoSConfigurationWithRemark) String() string {
	return dara.Prettify(s)
}

func (s QoSConfigurationWithRemark) GoString() string {
	return s.String()
}

func (s *QoSConfigurationWithRemark) GetExtranetDownloadBandwidth() *int64 {
	return s.ExtranetDownloadBandwidth
}

func (s *QoSConfigurationWithRemark) GetExtranetQps() *int64 {
	return s.ExtranetQps
}

func (s *QoSConfigurationWithRemark) GetExtranetUploadBandwidth() *int64 {
	return s.ExtranetUploadBandwidth
}

func (s *QoSConfigurationWithRemark) GetIntranetDownloadBandwidth() *int64 {
	return s.IntranetDownloadBandwidth
}

func (s *QoSConfigurationWithRemark) GetIntranetQps() *int64 {
	return s.IntranetQps
}

func (s *QoSConfigurationWithRemark) GetIntranetUploadBandwidth() *int64 {
	return s.IntranetUploadBandwidth
}

func (s *QoSConfigurationWithRemark) GetRemark() *int64 {
	return s.Remark
}

func (s *QoSConfigurationWithRemark) GetTotalDownloadBandwidth() *int64 {
	return s.TotalDownloadBandwidth
}

func (s *QoSConfigurationWithRemark) GetTotalQps() *int64 {
	return s.TotalQps
}

func (s *QoSConfigurationWithRemark) GetTotalUploadBandwidth() *int64 {
	return s.TotalUploadBandwidth
}

func (s *QoSConfigurationWithRemark) SetExtranetDownloadBandwidth(v int64) *QoSConfigurationWithRemark {
	s.ExtranetDownloadBandwidth = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetExtranetQps(v int64) *QoSConfigurationWithRemark {
	s.ExtranetQps = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetExtranetUploadBandwidth(v int64) *QoSConfigurationWithRemark {
	s.ExtranetUploadBandwidth = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetIntranetDownloadBandwidth(v int64) *QoSConfigurationWithRemark {
	s.IntranetDownloadBandwidth = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetIntranetQps(v int64) *QoSConfigurationWithRemark {
	s.IntranetQps = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetIntranetUploadBandwidth(v int64) *QoSConfigurationWithRemark {
	s.IntranetUploadBandwidth = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetRemark(v int64) *QoSConfigurationWithRemark {
	s.Remark = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetTotalDownloadBandwidth(v int64) *QoSConfigurationWithRemark {
	s.TotalDownloadBandwidth = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetTotalQps(v int64) *QoSConfigurationWithRemark {
	s.TotalQps = &v
	return s
}

func (s *QoSConfigurationWithRemark) SetTotalUploadBandwidth(v int64) *QoSConfigurationWithRemark {
	s.TotalUploadBandwidth = &v
	return s
}

func (s *QoSConfigurationWithRemark) Validate() error {
	return dara.Validate(s)
}
