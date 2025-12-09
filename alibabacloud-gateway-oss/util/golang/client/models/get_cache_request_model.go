// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeCacheAvailableZone(v string) *GetCacheRequest
	GetXOssDatalakeCacheAvailableZone() *string
	SetXOssDatalakeCacheName(v string) *GetCacheRequest
	GetXOssDatalakeCacheName() *string
	SetXOssDatalakeCacheVerbose(v bool) *GetCacheRequest
	GetXOssDatalakeCacheVerbose() *bool
}

type GetCacheRequest struct {
	// This parameter is required.
	XOssDatalakeCacheAvailableZone *string `json:"x-oss-datalake-cache-available-zone,omitempty" xml:"x-oss-datalake-cache-available-zone,omitempty"`
	// This parameter is required.
	XOssDatalakeCacheName    *string `json:"x-oss-datalake-cache-name,omitempty" xml:"x-oss-datalake-cache-name,omitempty"`
	XOssDatalakeCacheVerbose *bool   `json:"x-oss-datalake-cache-verbose,omitempty" xml:"x-oss-datalake-cache-verbose,omitempty"`
}

func (s GetCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCacheRequest) GoString() string {
	return s.String()
}

func (s *GetCacheRequest) GetXOssDatalakeCacheAvailableZone() *string {
	return s.XOssDatalakeCacheAvailableZone
}

func (s *GetCacheRequest) GetXOssDatalakeCacheName() *string {
	return s.XOssDatalakeCacheName
}

func (s *GetCacheRequest) GetXOssDatalakeCacheVerbose() *bool {
	return s.XOssDatalakeCacheVerbose
}

func (s *GetCacheRequest) SetXOssDatalakeCacheAvailableZone(v string) *GetCacheRequest {
	s.XOssDatalakeCacheAvailableZone = &v
	return s
}

func (s *GetCacheRequest) SetXOssDatalakeCacheName(v string) *GetCacheRequest {
	s.XOssDatalakeCacheName = &v
	return s
}

func (s *GetCacheRequest) SetXOssDatalakeCacheVerbose(v bool) *GetCacheRequest {
	s.XOssDatalakeCacheVerbose = &v
	return s
}

func (s *GetCacheRequest) Validate() error {
	return dara.Validate(s)
}
