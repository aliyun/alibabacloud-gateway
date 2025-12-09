// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCnameTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCnameConfiguration(v *CreateCnameTokenRequestBucketCnameConfiguration) *CreateCnameTokenRequest
	GetBucketCnameConfiguration() *CreateCnameTokenRequestBucketCnameConfiguration
}

type CreateCnameTokenRequest struct {
	// The container that stores the CNAME record.
	BucketCnameConfiguration *CreateCnameTokenRequestBucketCnameConfiguration `json:"BucketCnameConfiguration,omitempty" xml:"BucketCnameConfiguration,omitempty" type:"Struct"`
}

func (s CreateCnameTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCnameTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateCnameTokenRequest) GetBucketCnameConfiguration() *CreateCnameTokenRequestBucketCnameConfiguration {
	return s.BucketCnameConfiguration
}

func (s *CreateCnameTokenRequest) SetBucketCnameConfiguration(v *CreateCnameTokenRequestBucketCnameConfiguration) *CreateCnameTokenRequest {
	s.BucketCnameConfiguration = v
	return s
}

func (s *CreateCnameTokenRequest) Validate() error {
	if s.BucketCnameConfiguration != nil {
		if err := s.BucketCnameConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCnameTokenRequestBucketCnameConfiguration struct {
	// The container that stores the CNAME information.
	Cname *CreateCnameTokenRequestBucketCnameConfigurationCname `json:"Cname,omitempty" xml:"Cname,omitempty" type:"Struct"`
}

func (s CreateCnameTokenRequestBucketCnameConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateCnameTokenRequestBucketCnameConfiguration) GoString() string {
	return s.String()
}

func (s *CreateCnameTokenRequestBucketCnameConfiguration) GetCname() *CreateCnameTokenRequestBucketCnameConfigurationCname {
	return s.Cname
}

func (s *CreateCnameTokenRequestBucketCnameConfiguration) SetCname(v *CreateCnameTokenRequestBucketCnameConfigurationCname) *CreateCnameTokenRequestBucketCnameConfiguration {
	s.Cname = v
	return s
}

func (s *CreateCnameTokenRequestBucketCnameConfiguration) Validate() error {
	if s.Cname != nil {
		if err := s.Cname.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCnameTokenRequestBucketCnameConfigurationCname struct {
	// The custom domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CreateCnameTokenRequestBucketCnameConfigurationCname) String() string {
	return dara.Prettify(s)
}

func (s CreateCnameTokenRequestBucketCnameConfigurationCname) GoString() string {
	return s.String()
}

func (s *CreateCnameTokenRequestBucketCnameConfigurationCname) GetDomain() *string {
	return s.Domain
}

func (s *CreateCnameTokenRequestBucketCnameConfigurationCname) SetDomain(v string) *CreateCnameTokenRequestBucketCnameConfigurationCname {
	s.Domain = &v
	return s
}

func (s *CreateCnameTokenRequestBucketCnameConfigurationCname) Validate() error {
	return dara.Validate(s)
}
