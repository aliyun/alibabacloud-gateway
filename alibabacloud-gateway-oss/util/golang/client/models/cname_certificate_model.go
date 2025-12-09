// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCnameCertificate interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *CnameCertificate
	GetCertId() *string
	SetCreationDate(v string) *CnameCertificate
	GetCreationDate() *string
	SetFingerprint(v string) *CnameCertificate
	GetFingerprint() *string
	SetStatus(v string) *CnameCertificate
	GetStatus() *string
	SetType(v string) *CnameCertificate
	GetType() *string
	SetValidEndDate(v string) *CnameCertificate
	GetValidEndDate() *string
	SetValidStartDate(v string) *CnameCertificate
	GetValidStartDate() *string
}

type CnameCertificate struct {
	CertId         *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CreationDate   *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	Fingerprint    *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ValidEndDate   *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s CnameCertificate) String() string {
	return dara.Prettify(s)
}

func (s CnameCertificate) GoString() string {
	return s.String()
}

func (s *CnameCertificate) GetCertId() *string {
	return s.CertId
}

func (s *CnameCertificate) GetCreationDate() *string {
	return s.CreationDate
}

func (s *CnameCertificate) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *CnameCertificate) GetStatus() *string {
	return s.Status
}

func (s *CnameCertificate) GetType() *string {
	return s.Type
}

func (s *CnameCertificate) GetValidEndDate() *string {
	return s.ValidEndDate
}

func (s *CnameCertificate) GetValidStartDate() *string {
	return s.ValidStartDate
}

func (s *CnameCertificate) SetCertId(v string) *CnameCertificate {
	s.CertId = &v
	return s
}

func (s *CnameCertificate) SetCreationDate(v string) *CnameCertificate {
	s.CreationDate = &v
	return s
}

func (s *CnameCertificate) SetFingerprint(v string) *CnameCertificate {
	s.Fingerprint = &v
	return s
}

func (s *CnameCertificate) SetStatus(v string) *CnameCertificate {
	s.Status = &v
	return s
}

func (s *CnameCertificate) SetType(v string) *CnameCertificate {
	s.Type = &v
	return s
}

func (s *CnameCertificate) SetValidEndDate(v string) *CnameCertificate {
	s.ValidEndDate = &v
	return s
}

func (s *CnameCertificate) SetValidStartDate(v string) *CnameCertificate {
	s.ValidStartDate = &v
	return s
}

func (s *CnameCertificate) Validate() error {
	return dara.Validate(s)
}
