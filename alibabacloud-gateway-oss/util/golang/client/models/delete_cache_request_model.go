// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeCacheAvailableZone(v string) *DeleteCacheRequest
	GetXOssDatalakeCacheAvailableZone() *string
	SetXOssDatalakeCacheName(v string) *DeleteCacheRequest
	GetXOssDatalakeCacheName() *string
}

type DeleteCacheRequest struct {
	XOssDatalakeCacheAvailableZone *string `json:"x-oss-datalake-cache-available-zone,omitempty" xml:"x-oss-datalake-cache-available-zone,omitempty"`
	XOssDatalakeCacheName          *string `json:"x-oss-datalake-cache-name,omitempty" xml:"x-oss-datalake-cache-name,omitempty"`
}

func (s DeleteCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteCacheRequest) GetXOssDatalakeCacheAvailableZone() *string {
	return s.XOssDatalakeCacheAvailableZone
}

func (s *DeleteCacheRequest) GetXOssDatalakeCacheName() *string {
	return s.XOssDatalakeCacheName
}

func (s *DeleteCacheRequest) SetXOssDatalakeCacheAvailableZone(v string) *DeleteCacheRequest {
	s.XOssDatalakeCacheAvailableZone = &v
	return s
}

func (s *DeleteCacheRequest) SetXOssDatalakeCacheName(v string) *DeleteCacheRequest {
	s.XOssDatalakeCacheName = &v
	return s
}

func (s *DeleteCacheRequest) Validate() error {
	return dara.Validate(s)
}
