// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpsConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCipherSuite(v *HttpsConfigurationCipherSuite) *HttpsConfiguration
	GetCipherSuite() *HttpsConfigurationCipherSuite
	SetTLS(v *HttpsConfigurationTLS) *HttpsConfiguration
	GetTLS() *HttpsConfigurationTLS
}

type HttpsConfiguration struct {
	CipherSuite *HttpsConfigurationCipherSuite `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty" type:"Struct"`
	TLS         *HttpsConfigurationTLS         `json:"TLS,omitempty" xml:"TLS,omitempty" type:"Struct"`
}

func (s HttpsConfiguration) String() string {
	return dara.Prettify(s)
}

func (s HttpsConfiguration) GoString() string {
	return s.String()
}

func (s *HttpsConfiguration) GetCipherSuite() *HttpsConfigurationCipherSuite {
	return s.CipherSuite
}

func (s *HttpsConfiguration) GetTLS() *HttpsConfigurationTLS {
	return s.TLS
}

func (s *HttpsConfiguration) SetCipherSuite(v *HttpsConfigurationCipherSuite) *HttpsConfiguration {
	s.CipherSuite = v
	return s
}

func (s *HttpsConfiguration) SetTLS(v *HttpsConfigurationTLS) *HttpsConfiguration {
	s.TLS = v
	return s
}

func (s *HttpsConfiguration) Validate() error {
	if s.CipherSuite != nil {
		if err := s.CipherSuite.Validate(); err != nil {
			return err
		}
	}
	if s.TLS != nil {
		if err := s.TLS.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpsConfigurationCipherSuite struct {
	CustomCipherSuite []*string `json:"CustomCipherSuite,omitempty" xml:"CustomCipherSuite,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// false
	StrongCipherSuite      *bool     `json:"StrongCipherSuite,omitempty" xml:"StrongCipherSuite,omitempty"`
	TLS13CustomCipherSuite []*string `json:"TLS13CustomCipherSuite,omitempty" xml:"TLS13CustomCipherSuite,omitempty" type:"Repeated"`
}

func (s HttpsConfigurationCipherSuite) String() string {
	return dara.Prettify(s)
}

func (s HttpsConfigurationCipherSuite) GoString() string {
	return s.String()
}

func (s *HttpsConfigurationCipherSuite) GetCustomCipherSuite() []*string {
	return s.CustomCipherSuite
}

func (s *HttpsConfigurationCipherSuite) GetEnable() *bool {
	return s.Enable
}

func (s *HttpsConfigurationCipherSuite) GetStrongCipherSuite() *bool {
	return s.StrongCipherSuite
}

func (s *HttpsConfigurationCipherSuite) GetTLS13CustomCipherSuite() []*string {
	return s.TLS13CustomCipherSuite
}

func (s *HttpsConfigurationCipherSuite) SetCustomCipherSuite(v []*string) *HttpsConfigurationCipherSuite {
	s.CustomCipherSuite = v
	return s
}

func (s *HttpsConfigurationCipherSuite) SetEnable(v bool) *HttpsConfigurationCipherSuite {
	s.Enable = &v
	return s
}

func (s *HttpsConfigurationCipherSuite) SetStrongCipherSuite(v bool) *HttpsConfigurationCipherSuite {
	s.StrongCipherSuite = &v
	return s
}

func (s *HttpsConfigurationCipherSuite) SetTLS13CustomCipherSuite(v []*string) *HttpsConfigurationCipherSuite {
	s.TLS13CustomCipherSuite = v
	return s
}

func (s *HttpsConfigurationCipherSuite) Validate() error {
	return dara.Validate(s)
}

type HttpsConfigurationTLS struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable     *bool     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	TLSVersion []*string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty" type:"Repeated"`
}

func (s HttpsConfigurationTLS) String() string {
	return dara.Prettify(s)
}

func (s HttpsConfigurationTLS) GoString() string {
	return s.String()
}

func (s *HttpsConfigurationTLS) GetEnable() *bool {
	return s.Enable
}

func (s *HttpsConfigurationTLS) GetTLSVersion() []*string {
	return s.TLSVersion
}

func (s *HttpsConfigurationTLS) SetEnable(v bool) *HttpsConfigurationTLS {
	s.Enable = &v
	return s
}

func (s *HttpsConfigurationTLS) SetTLSVersion(v []*string) *HttpsConfigurationTLS {
	s.TLSVersion = v
	return s
}

func (s *HttpsConfigurationTLS) Validate() error {
	return dara.Validate(s)
}
