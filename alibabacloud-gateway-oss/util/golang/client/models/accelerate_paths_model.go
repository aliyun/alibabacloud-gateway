// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceleratePaths interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCachePolicy(v string) *AcceleratePaths
	GetDefaultCachePolicy() *string
	SetPath(v []*AcceleratePathsPath) *AcceleratePaths
	GetPath() []*AcceleratePathsPath
}

type AcceleratePaths struct {
	// example:
	//
	// sync-warmup
	DefaultCachePolicy *string                `json:"DefaultCachePolicy,omitempty" xml:"DefaultCachePolicy,omitempty"`
	Path               []*AcceleratePathsPath `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s AcceleratePaths) String() string {
	return dara.Prettify(s)
}

func (s AcceleratePaths) GoString() string {
	return s.String()
}

func (s *AcceleratePaths) GetDefaultCachePolicy() *string {
	return s.DefaultCachePolicy
}

func (s *AcceleratePaths) GetPath() []*AcceleratePathsPath {
	return s.Path
}

func (s *AcceleratePaths) SetDefaultCachePolicy(v string) *AcceleratePaths {
	s.DefaultCachePolicy = &v
	return s
}

func (s *AcceleratePaths) SetPath(v []*AcceleratePathsPath) *AcceleratePaths {
	s.Path = v
	return s
}

func (s *AcceleratePaths) Validate() error {
	if s.Path != nil {
		for _, item := range s.Path {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AcceleratePathsPath struct {
	// example:
	//
	// sync-warmup
	CachePolicy *string `json:"CachePolicy,omitempty" xml:"CachePolicy,omitempty"`
	// example:
	//
	// acc-path/
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AcceleratePathsPath) String() string {
	return dara.Prettify(s)
}

func (s AcceleratePathsPath) GoString() string {
	return s.String()
}

func (s *AcceleratePathsPath) GetCachePolicy() *string {
	return s.CachePolicy
}

func (s *AcceleratePathsPath) GetName() *string {
	return s.Name
}

func (s *AcceleratePathsPath) SetCachePolicy(v string) *AcceleratePathsPath {
	s.CachePolicy = &v
	return s
}

func (s *AcceleratePathsPath) SetName(v string) *AcceleratePathsPath {
	s.Name = &v
	return s
}

func (s *AcceleratePathsPath) Validate() error {
	return dara.Validate(s)
}
