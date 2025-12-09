// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheBaseInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZone(v string) *CacheBaseInfo
	GetAvailableZone() *string
	SetCreationDate(v string) *CacheBaseInfo
	GetCreationDate() *string
	SetName(v string) *CacheBaseInfo
	GetName() *string
	SetQuotaConfiguration(v *CacheQuotaConfiguration) *CacheBaseInfo
	GetQuotaConfiguration() *CacheQuotaConfiguration
}

type CacheBaseInfo struct {
	// example:
	//
	// cn-hangzhou-j
	AvailableZone *string `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty"`
	// example:
	//
	// 2023-09-12T15:26:29.000Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// example:
	//
	// cache1
	Name               *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	QuotaConfiguration *CacheQuotaConfiguration `json:"QuotaConfiguration,omitempty" xml:"QuotaConfiguration,omitempty"`
}

func (s CacheBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s CacheBaseInfo) GoString() string {
	return s.String()
}

func (s *CacheBaseInfo) GetAvailableZone() *string {
	return s.AvailableZone
}

func (s *CacheBaseInfo) GetCreationDate() *string {
	return s.CreationDate
}

func (s *CacheBaseInfo) GetName() *string {
	return s.Name
}

func (s *CacheBaseInfo) GetQuotaConfiguration() *CacheQuotaConfiguration {
	return s.QuotaConfiguration
}

func (s *CacheBaseInfo) SetAvailableZone(v string) *CacheBaseInfo {
	s.AvailableZone = &v
	return s
}

func (s *CacheBaseInfo) SetCreationDate(v string) *CacheBaseInfo {
	s.CreationDate = &v
	return s
}

func (s *CacheBaseInfo) SetName(v string) *CacheBaseInfo {
	s.Name = &v
	return s
}

func (s *CacheBaseInfo) SetQuotaConfiguration(v *CacheQuotaConfiguration) *CacheBaseInfo {
	s.QuotaConfiguration = v
	return s
}

func (s *CacheBaseInfo) Validate() error {
	if s.QuotaConfiguration != nil {
		if err := s.QuotaConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
