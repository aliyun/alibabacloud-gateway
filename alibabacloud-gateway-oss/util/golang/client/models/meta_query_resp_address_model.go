// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryRespAddress interface {
	dara.Model
	String() string
	GoString() string
	SetAddressLine(v string) *MetaQueryRespAddress
	GetAddressLine() *string
	SetCity(v string) *MetaQueryRespAddress
	GetCity() *string
	SetDistrict(v string) *MetaQueryRespAddress
	GetDistrict() *string
	SetLanguage(v string) *MetaQueryRespAddress
	GetLanguage() *string
	SetProvince(v string) *MetaQueryRespAddress
	GetProvince() *string
	SetTownship(v string) *MetaQueryRespAddress
	GetTownship() *string
}

type MetaQueryRespAddress struct {
	// example:
	//
	// 中国浙江省杭州市余杭区文一西路969号
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	// example:
	//
	// 杭州市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 余杭区
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	// example:
	//
	// zh-Hans
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// 文一西路
	Township *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s MetaQueryRespAddress) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryRespAddress) GoString() string {
	return s.String()
}

func (s *MetaQueryRespAddress) GetAddressLine() *string {
	return s.AddressLine
}

func (s *MetaQueryRespAddress) GetCity() *string {
	return s.City
}

func (s *MetaQueryRespAddress) GetDistrict() *string {
	return s.District
}

func (s *MetaQueryRespAddress) GetLanguage() *string {
	return s.Language
}

func (s *MetaQueryRespAddress) GetProvince() *string {
	return s.Province
}

func (s *MetaQueryRespAddress) GetTownship() *string {
	return s.Township
}

func (s *MetaQueryRespAddress) SetAddressLine(v string) *MetaQueryRespAddress {
	s.AddressLine = &v
	return s
}

func (s *MetaQueryRespAddress) SetCity(v string) *MetaQueryRespAddress {
	s.City = &v
	return s
}

func (s *MetaQueryRespAddress) SetDistrict(v string) *MetaQueryRespAddress {
	s.District = &v
	return s
}

func (s *MetaQueryRespAddress) SetLanguage(v string) *MetaQueryRespAddress {
	s.Language = &v
	return s
}

func (s *MetaQueryRespAddress) SetProvince(v string) *MetaQueryRespAddress {
	s.Province = &v
	return s
}

func (s *MetaQueryRespAddress) SetTownship(v string) *MetaQueryRespAddress {
	s.Township = &v
	return s
}

func (s *MetaQueryRespAddress) Validate() error {
	return dara.Validate(s)
}
