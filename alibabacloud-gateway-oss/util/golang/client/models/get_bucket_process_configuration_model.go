// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketProcessConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetBucketChannelConfig(v *BucketChannelConfig) *GetBucketProcessConfiguration
	GetBucketChannelConfig() *BucketChannelConfig
	SetCompliedHost(v string) *GetBucketProcessConfiguration
	GetCompliedHost() *string
	SetLastModified(v string) *GetBucketProcessConfiguration
	GetLastModified() *string
	SetSourceFileProtect(v string) *GetBucketProcessConfiguration
	GetSourceFileProtect() *string
	SetSourceFileProtectSuffix(v string) *GetBucketProcessConfiguration
	GetSourceFileProtectSuffix() *string
	SetStyleDelimiters(v string) *GetBucketProcessConfiguration
	GetStyleDelimiters() *string
	SetVersion(v int32) *GetBucketProcessConfiguration
	GetVersion() *int32
}

type GetBucketProcessConfiguration struct {
	BucketChannelConfig *BucketChannelConfig `json:"BucketChannelConfig,omitempty" xml:"BucketChannelConfig,omitempty"`
	// example:
	//
	// Img
	CompliedHost *string `json:"CompliedHost,omitempty" xml:"CompliedHost,omitempty"`
	// example:
	//
	// 2024-10-18T09:51:54.000Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// example:
	//
	// Disabled
	SourceFileProtect *string `json:"SourceFileProtect,omitempty" xml:"SourceFileProtect,omitempty"`
	// example:
	//
	// .jpg
	SourceFileProtectSuffix *string `json:"SourceFileProtectSuffix,omitempty" xml:"SourceFileProtectSuffix,omitempty"`
	// example:
	//
	// -,_,/,!
	StyleDelimiters *string `json:"StyleDelimiters,omitempty" xml:"StyleDelimiters,omitempty"`
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetBucketProcessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketProcessConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketProcessConfiguration) GetBucketChannelConfig() *BucketChannelConfig {
	return s.BucketChannelConfig
}

func (s *GetBucketProcessConfiguration) GetCompliedHost() *string {
	return s.CompliedHost
}

func (s *GetBucketProcessConfiguration) GetLastModified() *string {
	return s.LastModified
}

func (s *GetBucketProcessConfiguration) GetSourceFileProtect() *string {
	return s.SourceFileProtect
}

func (s *GetBucketProcessConfiguration) GetSourceFileProtectSuffix() *string {
	return s.SourceFileProtectSuffix
}

func (s *GetBucketProcessConfiguration) GetStyleDelimiters() *string {
	return s.StyleDelimiters
}

func (s *GetBucketProcessConfiguration) GetVersion() *int32 {
	return s.Version
}

func (s *GetBucketProcessConfiguration) SetBucketChannelConfig(v *BucketChannelConfig) *GetBucketProcessConfiguration {
	s.BucketChannelConfig = v
	return s
}

func (s *GetBucketProcessConfiguration) SetCompliedHost(v string) *GetBucketProcessConfiguration {
	s.CompliedHost = &v
	return s
}

func (s *GetBucketProcessConfiguration) SetLastModified(v string) *GetBucketProcessConfiguration {
	s.LastModified = &v
	return s
}

func (s *GetBucketProcessConfiguration) SetSourceFileProtect(v string) *GetBucketProcessConfiguration {
	s.SourceFileProtect = &v
	return s
}

func (s *GetBucketProcessConfiguration) SetSourceFileProtectSuffix(v string) *GetBucketProcessConfiguration {
	s.SourceFileProtectSuffix = &v
	return s
}

func (s *GetBucketProcessConfiguration) SetStyleDelimiters(v string) *GetBucketProcessConfiguration {
	s.StyleDelimiters = &v
	return s
}

func (s *GetBucketProcessConfiguration) SetVersion(v int32) *GetBucketProcessConfiguration {
	s.Version = &v
	return s
}

func (s *GetBucketProcessConfiguration) Validate() error {
	if s.BucketChannelConfig != nil {
		if err := s.BucketChannelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
