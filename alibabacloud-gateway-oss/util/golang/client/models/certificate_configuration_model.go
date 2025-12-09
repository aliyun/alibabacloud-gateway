// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificateConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *CertificateConfiguration
	GetCertId() *string
	SetCertificate(v string) *CertificateConfiguration
	GetCertificate() *string
	SetDeleteCertificate(v bool) *CertificateConfiguration
	GetDeleteCertificate() *bool
	SetForce(v bool) *CertificateConfiguration
	GetForce() *bool
	SetPreviousCertId(v string) *CertificateConfiguration
	GetPreviousCertId() *string
	SetPrivateKey(v string) *CertificateConfiguration
	GetPrivateKey() *string
}

type CertificateConfiguration struct {
	// example:
	//
	// 493****-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIDhDCCAmwCCQCFs8ixARsyrDANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMC ***	- -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// false
	DeleteCertificate *bool `json:"DeleteCertificate,omitempty" xml:"DeleteCertificate,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// 493****-cn-hangzhou
	PreviousCertId *string `json:"PreviousCertId,omitempty" xml:"PreviousCertId,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIDhDCCAmwCCQCFs8ixARsyrDANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMC ***	- -----END CERTIFICATE-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s CertificateConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CertificateConfiguration) GoString() string {
	return s.String()
}

func (s *CertificateConfiguration) GetCertId() *string {
	return s.CertId
}

func (s *CertificateConfiguration) GetCertificate() *string {
	return s.Certificate
}

func (s *CertificateConfiguration) GetDeleteCertificate() *bool {
	return s.DeleteCertificate
}

func (s *CertificateConfiguration) GetForce() *bool {
	return s.Force
}

func (s *CertificateConfiguration) GetPreviousCertId() *string {
	return s.PreviousCertId
}

func (s *CertificateConfiguration) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CertificateConfiguration) SetCertId(v string) *CertificateConfiguration {
	s.CertId = &v
	return s
}

func (s *CertificateConfiguration) SetCertificate(v string) *CertificateConfiguration {
	s.Certificate = &v
	return s
}

func (s *CertificateConfiguration) SetDeleteCertificate(v bool) *CertificateConfiguration {
	s.DeleteCertificate = &v
	return s
}

func (s *CertificateConfiguration) SetForce(v bool) *CertificateConfiguration {
	s.Force = &v
	return s
}

func (s *CertificateConfiguration) SetPreviousCertId(v string) *CertificateConfiguration {
	s.PreviousCertId = &v
	return s
}

func (s *CertificateConfiguration) SetPrivateKey(v string) *CertificateConfiguration {
	s.PrivateKey = &v
	return s
}

func (s *CertificateConfiguration) Validate() error {
	return dara.Validate(s)
}
