// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCache(v *CacheDetailInfo) *GetCacheResponseBody
	GetCache() *CacheDetailInfo
}

type GetCacheResponseBody struct {
	Cache *CacheDetailInfo `json:"Cache,omitempty" xml:"Cache,omitempty"`
}

func (s GetCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCacheResponseBody) GoString() string {
	return s.String()
}

func (s *GetCacheResponseBody) GetCache() *CacheDetailInfo {
	return s.Cache
}

func (s *GetCacheResponseBody) SetCache(v *CacheDetailInfo) *GetCacheResponseBody {
	s.Cache = v
	return s
}

func (s *GetCacheResponseBody) Validate() error {
	if s.Cache != nil {
		if err := s.Cache.Validate(); err != nil {
			return err
		}
	}
	return nil
}
