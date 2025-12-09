// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZone(v string) *CreateCacheConfiguration
	GetAvailableZone() *string
	SetName(v string) *CreateCacheConfiguration
	GetName() *string
	SetQuotaConfiguration(v *CacheQuotaConfiguration) *CreateCacheConfiguration
	GetQuotaConfiguration() *CacheQuotaConfiguration
}

type CreateCacheConfiguration struct {
	// example:
	//
	// cn-hangzhou-j
	AvailableZone *string `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty"`
	// example:
	//
	// cache2
	Name               *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	QuotaConfiguration *CacheQuotaConfiguration `json:"QuotaConfiguration,omitempty" xml:"QuotaConfiguration,omitempty"`
}

func (s CreateCacheConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheConfiguration) GoString() string {
	return s.String()
}

func (s *CreateCacheConfiguration) GetAvailableZone() *string {
	return s.AvailableZone
}

func (s *CreateCacheConfiguration) GetName() *string {
	return s.Name
}

func (s *CreateCacheConfiguration) GetQuotaConfiguration() *CacheQuotaConfiguration {
	return s.QuotaConfiguration
}

func (s *CreateCacheConfiguration) SetAvailableZone(v string) *CreateCacheConfiguration {
	s.AvailableZone = &v
	return s
}

func (s *CreateCacheConfiguration) SetName(v string) *CreateCacheConfiguration {
	s.Name = &v
	return s
}

func (s *CreateCacheConfiguration) SetQuotaConfiguration(v *CacheQuotaConfiguration) *CreateCacheConfiguration {
	s.QuotaConfiguration = v
	return s
}

func (s *CreateCacheConfiguration) Validate() error {
	if s.QuotaConfiguration != nil {
		if err := s.QuotaConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
