// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCnameSummary interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v *CnameCertificate) *CnameSummary
	GetCertificate() *CnameCertificate
	SetDomain(v string) *CnameSummary
	GetDomain() *string
	SetLastModified(v string) *CnameSummary
	GetLastModified() *string
	SetStatus(v string) *CnameSummary
	GetStatus() *string
}

type CnameSummary struct {
	Certificate  *CnameCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Domain       *string           `json:"Domain,omitempty" xml:"Domain,omitempty"`
	LastModified *string           `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Status       *string           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CnameSummary) String() string {
	return dara.Prettify(s)
}

func (s CnameSummary) GoString() string {
	return s.String()
}

func (s *CnameSummary) GetCertificate() *CnameCertificate {
	return s.Certificate
}

func (s *CnameSummary) GetDomain() *string {
	return s.Domain
}

func (s *CnameSummary) GetLastModified() *string {
	return s.LastModified
}

func (s *CnameSummary) GetStatus() *string {
	return s.Status
}

func (s *CnameSummary) SetCertificate(v *CnameCertificate) *CnameSummary {
	s.Certificate = v
	return s
}

func (s *CnameSummary) SetDomain(v string) *CnameSummary {
	s.Domain = &v
	return s
}

func (s *CnameSummary) SetLastModified(v string) *CnameSummary {
	s.LastModified = &v
	return s
}

func (s *CnameSummary) SetStatus(v string) *CnameSummary {
	s.Status = &v
	return s
}

func (s *CnameSummary) Validate() error {
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	return nil
}
