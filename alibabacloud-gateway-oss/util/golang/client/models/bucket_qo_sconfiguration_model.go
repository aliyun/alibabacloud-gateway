// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketQoSConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetExclusive(v bool) *BucketQoSConfiguration
	GetExclusive() *bool
	SetExtranetDownloadBandwidth(v int64) *BucketQoSConfiguration
	GetExtranetDownloadBandwidth() *int64
	SetExtranetQps(v int64) *BucketQoSConfiguration
	GetExtranetQps() *int64
	SetExtranetUploadBandwidth(v int64) *BucketQoSConfiguration
	GetExtranetUploadBandwidth() *int64
	SetIntranetDownloadBandwidth(v int64) *BucketQoSConfiguration
	GetIntranetDownloadBandwidth() *int64
	SetIntranetQps(v int64) *BucketQoSConfiguration
	GetIntranetQps() *int64
	SetIntranetUploadBandwidth(v int64) *BucketQoSConfiguration
	GetIntranetUploadBandwidth() *int64
	SetTotalDownloadBandwidth(v int64) *BucketQoSConfiguration
	GetTotalDownloadBandwidth() *int64
	SetTotalQps(v int64) *BucketQoSConfiguration
	GetTotalQps() *int64
	SetTotalUploadBandwidth(v int64) *BucketQoSConfiguration
	GetTotalUploadBandwidth() *int64
}

type BucketQoSConfiguration struct {
	// example:
	//
	// true
	Exclusive *bool `json:"Exclusive,omitempty" xml:"Exclusive,omitempty"`
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

func (s BucketQoSConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BucketQoSConfiguration) GoString() string {
	return s.String()
}

func (s *BucketQoSConfiguration) GetExclusive() *bool {
	return s.Exclusive
}

func (s *BucketQoSConfiguration) GetExtranetDownloadBandwidth() *int64 {
	return s.ExtranetDownloadBandwidth
}

func (s *BucketQoSConfiguration) GetExtranetQps() *int64 {
	return s.ExtranetQps
}

func (s *BucketQoSConfiguration) GetExtranetUploadBandwidth() *int64 {
	return s.ExtranetUploadBandwidth
}

func (s *BucketQoSConfiguration) GetIntranetDownloadBandwidth() *int64 {
	return s.IntranetDownloadBandwidth
}

func (s *BucketQoSConfiguration) GetIntranetQps() *int64 {
	return s.IntranetQps
}

func (s *BucketQoSConfiguration) GetIntranetUploadBandwidth() *int64 {
	return s.IntranetUploadBandwidth
}

func (s *BucketQoSConfiguration) GetTotalDownloadBandwidth() *int64 {
	return s.TotalDownloadBandwidth
}

func (s *BucketQoSConfiguration) GetTotalQps() *int64 {
	return s.TotalQps
}

func (s *BucketQoSConfiguration) GetTotalUploadBandwidth() *int64 {
	return s.TotalUploadBandwidth
}

func (s *BucketQoSConfiguration) SetExclusive(v bool) *BucketQoSConfiguration {
	s.Exclusive = &v
	return s
}

func (s *BucketQoSConfiguration) SetExtranetDownloadBandwidth(v int64) *BucketQoSConfiguration {
	s.ExtranetDownloadBandwidth = &v
	return s
}

func (s *BucketQoSConfiguration) SetExtranetQps(v int64) *BucketQoSConfiguration {
	s.ExtranetQps = &v
	return s
}

func (s *BucketQoSConfiguration) SetExtranetUploadBandwidth(v int64) *BucketQoSConfiguration {
	s.ExtranetUploadBandwidth = &v
	return s
}

func (s *BucketQoSConfiguration) SetIntranetDownloadBandwidth(v int64) *BucketQoSConfiguration {
	s.IntranetDownloadBandwidth = &v
	return s
}

func (s *BucketQoSConfiguration) SetIntranetQps(v int64) *BucketQoSConfiguration {
	s.IntranetQps = &v
	return s
}

func (s *BucketQoSConfiguration) SetIntranetUploadBandwidth(v int64) *BucketQoSConfiguration {
	s.IntranetUploadBandwidth = &v
	return s
}

func (s *BucketQoSConfiguration) SetTotalDownloadBandwidth(v int64) *BucketQoSConfiguration {
	s.TotalDownloadBandwidth = &v
	return s
}

func (s *BucketQoSConfiguration) SetTotalQps(v int64) *BucketQoSConfiguration {
	s.TotalQps = &v
	return s
}

func (s *BucketQoSConfiguration) SetTotalUploadBandwidth(v int64) *BucketQoSConfiguration {
	s.TotalUploadBandwidth = &v
	return s
}

func (s *BucketQoSConfiguration) Validate() error {
	return dara.Validate(s)
}
