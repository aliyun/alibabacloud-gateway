// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketAntiDDOSConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCnames(v *BucketAntiDDOSConfigurationCnames) *BucketAntiDDOSConfiguration
	GetCnames() *BucketAntiDDOSConfigurationCnames
}

type BucketAntiDDOSConfiguration struct {
	Cnames *BucketAntiDDOSConfigurationCnames `json:"Cnames,omitempty" xml:"Cnames,omitempty" type:"Struct"`
}

func (s BucketAntiDDOSConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BucketAntiDDOSConfiguration) GoString() string {
	return s.String()
}

func (s *BucketAntiDDOSConfiguration) GetCnames() *BucketAntiDDOSConfigurationCnames {
	return s.Cnames
}

func (s *BucketAntiDDOSConfiguration) SetCnames(v *BucketAntiDDOSConfigurationCnames) *BucketAntiDDOSConfiguration {
	s.Cnames = v
	return s
}

func (s *BucketAntiDDOSConfiguration) Validate() error {
	if s.Cnames != nil {
		if err := s.Cnames.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BucketAntiDDOSConfigurationCnames struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s BucketAntiDDOSConfigurationCnames) String() string {
	return dara.Prettify(s)
}

func (s BucketAntiDDOSConfigurationCnames) GoString() string {
	return s.String()
}

func (s *BucketAntiDDOSConfigurationCnames) GetDomain() []*string {
	return s.Domain
}

func (s *BucketAntiDDOSConfigurationCnames) SetDomain(v []*string) *BucketAntiDDOSConfigurationCnames {
	s.Domain = v
	return s
}

func (s *BucketAntiDDOSConfigurationCnames) Validate() error {
	return dara.Validate(s)
}
