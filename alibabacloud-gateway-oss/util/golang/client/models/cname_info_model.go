// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCnameInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v *CnameCertificate) *CnameInfo
	GetCertificate() *CnameCertificate
	SetDomain(v string) *CnameInfo
	GetDomain() *string
	SetLastModified(v string) *CnameInfo
	GetLastModified() *string
	SetStatus(v string) *CnameInfo
	GetStatus() *string
}

type CnameInfo struct {
	Certificate  *CnameCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Domain       *string           `json:"Domain,omitempty" xml:"Domain,omitempty"`
	LastModified *string           `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Status       *string           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CnameInfo) String() string {
	return dara.Prettify(s)
}

func (s CnameInfo) GoString() string {
	return s.String()
}

func (s *CnameInfo) GetCertificate() *CnameCertificate {
	return s.Certificate
}

func (s *CnameInfo) GetDomain() *string {
	return s.Domain
}

func (s *CnameInfo) GetLastModified() *string {
	return s.LastModified
}

func (s *CnameInfo) GetStatus() *string {
	return s.Status
}

func (s *CnameInfo) SetCertificate(v *CnameCertificate) *CnameInfo {
	s.Certificate = v
	return s
}

func (s *CnameInfo) SetDomain(v string) *CnameInfo {
	s.Domain = &v
	return s
}

func (s *CnameInfo) SetLastModified(v string) *CnameInfo {
	s.LastModified = &v
	return s
}

func (s *CnameInfo) SetStatus(v string) *CnameInfo {
	s.Status = &v
	return s
}

func (s *CnameInfo) Validate() error {
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	return nil
}
