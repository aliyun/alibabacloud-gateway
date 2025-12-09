// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheQuotaConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaDesc(v *CacheQuotaConfigurationQuotaDesc) *CacheQuotaConfiguration
	GetQuotaDesc() *CacheQuotaConfigurationQuotaDesc
}

type CacheQuotaConfiguration struct {
	QuotaDesc *CacheQuotaConfigurationQuotaDesc `json:"QuotaDesc,omitempty" xml:"QuotaDesc,omitempty" type:"Struct"`
}

func (s CacheQuotaConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CacheQuotaConfiguration) GoString() string {
	return s.String()
}

func (s *CacheQuotaConfiguration) GetQuotaDesc() *CacheQuotaConfigurationQuotaDesc {
	return s.QuotaDesc
}

func (s *CacheQuotaConfiguration) SetQuotaDesc(v *CacheQuotaConfigurationQuotaDesc) *CacheQuotaConfiguration {
	s.QuotaDesc = v
	return s
}

func (s *CacheQuotaConfiguration) Validate() error {
	if s.QuotaDesc != nil {
		if err := s.QuotaDesc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CacheQuotaConfigurationQuotaDesc struct {
	// example:
	//
	// 30
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s CacheQuotaConfigurationQuotaDesc) String() string {
	return dara.Prettify(s)
}

func (s CacheQuotaConfigurationQuotaDesc) GoString() string {
	return s.String()
}

func (s *CacheQuotaConfigurationQuotaDesc) GetQuota() *int64 {
	return s.Quota
}

func (s *CacheQuotaConfigurationQuotaDesc) SetQuota(v int64) *CacheQuotaConfigurationQuotaDesc {
	s.Quota = &v
	return s
}

func (s *CacheQuotaConfigurationQuotaDesc) Validate() error {
	return dara.Validate(s)
}
