// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketCnameConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v *BucketCnameConfigurationCname) *BucketCnameConfiguration
	GetCname() *BucketCnameConfigurationCname
}

type BucketCnameConfiguration struct {
	Cname *BucketCnameConfigurationCname `json:"Cname,omitempty" xml:"Cname,omitempty" type:"Struct"`
}

func (s BucketCnameConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BucketCnameConfiguration) GoString() string {
	return s.String()
}

func (s *BucketCnameConfiguration) GetCname() *BucketCnameConfigurationCname {
	return s.Cname
}

func (s *BucketCnameConfiguration) SetCname(v *BucketCnameConfigurationCname) *BucketCnameConfiguration {
	s.Cname = v
	return s
}

func (s *BucketCnameConfiguration) Validate() error {
	if s.Cname != nil {
		if err := s.Cname.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BucketCnameConfigurationCname struct {
	CertificateConfiguration *CertificateConfiguration `json:"CertificateConfiguration,omitempty" xml:"CertificateConfiguration,omitempty"`
	Domain                   *string                   `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s BucketCnameConfigurationCname) String() string {
	return dara.Prettify(s)
}

func (s BucketCnameConfigurationCname) GoString() string {
	return s.String()
}

func (s *BucketCnameConfigurationCname) GetCertificateConfiguration() *CertificateConfiguration {
	return s.CertificateConfiguration
}

func (s *BucketCnameConfigurationCname) GetDomain() *string {
	return s.Domain
}

func (s *BucketCnameConfigurationCname) SetCertificateConfiguration(v *CertificateConfiguration) *BucketCnameConfigurationCname {
	s.CertificateConfiguration = v
	return s
}

func (s *BucketCnameConfigurationCname) SetDomain(v string) *BucketCnameConfigurationCname {
	s.Domain = &v
	return s
}

func (s *BucketCnameConfigurationCname) Validate() error {
	if s.CertificateConfiguration != nil {
		if err := s.CertificateConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
