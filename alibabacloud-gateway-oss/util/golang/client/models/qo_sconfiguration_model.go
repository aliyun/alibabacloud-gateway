// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQoSConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetExtranetDownloadBandwidth(v int64) *QoSConfiguration
	GetExtranetDownloadBandwidth() *int64
	SetExtranetQps(v int64) *QoSConfiguration
	GetExtranetQps() *int64
	SetExtranetUploadBandwidth(v int64) *QoSConfiguration
	GetExtranetUploadBandwidth() *int64
	SetIntranetDownloadBandwidth(v int64) *QoSConfiguration
	GetIntranetDownloadBandwidth() *int64
	SetIntranetQps(v int64) *QoSConfiguration
	GetIntranetQps() *int64
	SetIntranetUploadBandwidth(v int64) *QoSConfiguration
	GetIntranetUploadBandwidth() *int64
	SetTotalDownloadBandwidth(v int64) *QoSConfiguration
	GetTotalDownloadBandwidth() *int64
	SetTotalQps(v int64) *QoSConfiguration
	GetTotalQps() *int64
	SetTotalUploadBandwidth(v int64) *QoSConfiguration
	GetTotalUploadBandwidth() *int64
}

type QoSConfiguration struct {
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

func (s QoSConfiguration) String() string {
	return dara.Prettify(s)
}

func (s QoSConfiguration) GoString() string {
	return s.String()
}

func (s *QoSConfiguration) GetExtranetDownloadBandwidth() *int64 {
	return s.ExtranetDownloadBandwidth
}

func (s *QoSConfiguration) GetExtranetQps() *int64 {
	return s.ExtranetQps
}

func (s *QoSConfiguration) GetExtranetUploadBandwidth() *int64 {
	return s.ExtranetUploadBandwidth
}

func (s *QoSConfiguration) GetIntranetDownloadBandwidth() *int64 {
	return s.IntranetDownloadBandwidth
}

func (s *QoSConfiguration) GetIntranetQps() *int64 {
	return s.IntranetQps
}

func (s *QoSConfiguration) GetIntranetUploadBandwidth() *int64 {
	return s.IntranetUploadBandwidth
}

func (s *QoSConfiguration) GetTotalDownloadBandwidth() *int64 {
	return s.TotalDownloadBandwidth
}

func (s *QoSConfiguration) GetTotalQps() *int64 {
	return s.TotalQps
}

func (s *QoSConfiguration) GetTotalUploadBandwidth() *int64 {
	return s.TotalUploadBandwidth
}

func (s *QoSConfiguration) SetExtranetDownloadBandwidth(v int64) *QoSConfiguration {
	s.ExtranetDownloadBandwidth = &v
	return s
}

func (s *QoSConfiguration) SetExtranetQps(v int64) *QoSConfiguration {
	s.ExtranetQps = &v
	return s
}

func (s *QoSConfiguration) SetExtranetUploadBandwidth(v int64) *QoSConfiguration {
	s.ExtranetUploadBandwidth = &v
	return s
}

func (s *QoSConfiguration) SetIntranetDownloadBandwidth(v int64) *QoSConfiguration {
	s.IntranetDownloadBandwidth = &v
	return s
}

func (s *QoSConfiguration) SetIntranetQps(v int64) *QoSConfiguration {
	s.IntranetQps = &v
	return s
}

func (s *QoSConfiguration) SetIntranetUploadBandwidth(v int64) *QoSConfiguration {
	s.IntranetUploadBandwidth = &v
	return s
}

func (s *QoSConfiguration) SetTotalDownloadBandwidth(v int64) *QoSConfiguration {
	s.TotalDownloadBandwidth = &v
	return s
}

func (s *QoSConfiguration) SetTotalQps(v int64) *QoSConfiguration {
	s.TotalQps = &v
	return s
}

func (s *QoSConfiguration) SetTotalUploadBandwidth(v int64) *QoSConfiguration {
	s.TotalUploadBandwidth = &v
	return s
}

func (s *QoSConfiguration) Validate() error {
	return dara.Validate(s)
}
