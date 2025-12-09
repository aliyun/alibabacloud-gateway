// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketProcessConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetBucketChannelConfig(v *BucketChannelConfig) *BucketProcessConfiguration
	GetBucketChannelConfig() *BucketChannelConfig
	SetCompliedHost(v string) *BucketProcessConfiguration
	GetCompliedHost() *string
	SetOssDomainSupportAtProcess(v string) *BucketProcessConfiguration
	GetOssDomainSupportAtProcess() *string
	SetSourceFileProtect(v string) *BucketProcessConfiguration
	GetSourceFileProtect() *string
	SetSourceFileProtectSuffix(v string) *BucketProcessConfiguration
	GetSourceFileProtectSuffix() *string
	SetStyleDelimiters(v string) *BucketProcessConfiguration
	GetStyleDelimiters() *string
}

type BucketProcessConfiguration struct {
	BucketChannelConfig *BucketChannelConfig `json:"BucketChannelConfig,omitempty" xml:"BucketChannelConfig,omitempty"`
	// example:
	//
	// Img
	CompliedHost *string `json:"CompliedHost,omitempty" xml:"CompliedHost,omitempty"`
	// example:
	//
	// disabled
	OssDomainSupportAtProcess *string `json:"OssDomainSupportAtProcess,omitempty" xml:"OssDomainSupportAtProcess,omitempty"`
	// example:
	//
	// disabled
	SourceFileProtect *string `json:"SourceFileProtect,omitempty" xml:"SourceFileProtect,omitempty"`
	// example:
	//
	// gif
	SourceFileProtectSuffix *string `json:"SourceFileProtectSuffix,omitempty" xml:"SourceFileProtectSuffix,omitempty"`
	// example:
	//
	// -,|
	StyleDelimiters *string `json:"StyleDelimiters,omitempty" xml:"StyleDelimiters,omitempty"`
}

func (s BucketProcessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BucketProcessConfiguration) GoString() string {
	return s.String()
}

func (s *BucketProcessConfiguration) GetBucketChannelConfig() *BucketChannelConfig {
	return s.BucketChannelConfig
}

func (s *BucketProcessConfiguration) GetCompliedHost() *string {
	return s.CompliedHost
}

func (s *BucketProcessConfiguration) GetOssDomainSupportAtProcess() *string {
	return s.OssDomainSupportAtProcess
}

func (s *BucketProcessConfiguration) GetSourceFileProtect() *string {
	return s.SourceFileProtect
}

func (s *BucketProcessConfiguration) GetSourceFileProtectSuffix() *string {
	return s.SourceFileProtectSuffix
}

func (s *BucketProcessConfiguration) GetStyleDelimiters() *string {
	return s.StyleDelimiters
}

func (s *BucketProcessConfiguration) SetBucketChannelConfig(v *BucketChannelConfig) *BucketProcessConfiguration {
	s.BucketChannelConfig = v
	return s
}

func (s *BucketProcessConfiguration) SetCompliedHost(v string) *BucketProcessConfiguration {
	s.CompliedHost = &v
	return s
}

func (s *BucketProcessConfiguration) SetOssDomainSupportAtProcess(v string) *BucketProcessConfiguration {
	s.OssDomainSupportAtProcess = &v
	return s
}

func (s *BucketProcessConfiguration) SetSourceFileProtect(v string) *BucketProcessConfiguration {
	s.SourceFileProtect = &v
	return s
}

func (s *BucketProcessConfiguration) SetSourceFileProtectSuffix(v string) *BucketProcessConfiguration {
	s.SourceFileProtectSuffix = &v
	return s
}

func (s *BucketProcessConfiguration) SetStyleDelimiters(v string) *BucketProcessConfiguration {
	s.StyleDelimiters = &v
	return s
}

func (s *BucketProcessConfiguration) Validate() error {
	if s.BucketChannelConfig != nil {
		if err := s.BucketChannelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
