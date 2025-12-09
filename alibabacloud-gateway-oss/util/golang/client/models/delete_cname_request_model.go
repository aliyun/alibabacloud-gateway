// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCnameConfiguration(v *DeleteCnameRequestBucketCnameConfiguration) *DeleteCnameRequest
	GetBucketCnameConfiguration() *DeleteCnameRequestBucketCnameConfiguration
}

type DeleteCnameRequest struct {
	// The container that stores the CNAME record.
	BucketCnameConfiguration *DeleteCnameRequestBucketCnameConfiguration `json:"BucketCnameConfiguration,omitempty" xml:"BucketCnameConfiguration,omitempty" type:"Struct"`
}

func (s DeleteCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCnameRequest) GoString() string {
	return s.String()
}

func (s *DeleteCnameRequest) GetBucketCnameConfiguration() *DeleteCnameRequestBucketCnameConfiguration {
	return s.BucketCnameConfiguration
}

func (s *DeleteCnameRequest) SetBucketCnameConfiguration(v *DeleteCnameRequestBucketCnameConfiguration) *DeleteCnameRequest {
	s.BucketCnameConfiguration = v
	return s
}

func (s *DeleteCnameRequest) Validate() error {
	if s.BucketCnameConfiguration != nil {
		if err := s.BucketCnameConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCnameRequestBucketCnameConfiguration struct {
	Cname *DeleteCnameRequestBucketCnameConfigurationCname `json:"Cname,omitempty" xml:"Cname,omitempty" type:"Struct"`
}

func (s DeleteCnameRequestBucketCnameConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DeleteCnameRequestBucketCnameConfiguration) GoString() string {
	return s.String()
}

func (s *DeleteCnameRequestBucketCnameConfiguration) GetCname() *DeleteCnameRequestBucketCnameConfigurationCname {
	return s.Cname
}

func (s *DeleteCnameRequestBucketCnameConfiguration) SetCname(v *DeleteCnameRequestBucketCnameConfigurationCname) *DeleteCnameRequestBucketCnameConfiguration {
	s.Cname = v
	return s
}

func (s *DeleteCnameRequestBucketCnameConfiguration) Validate() error {
	if s.Cname != nil {
		if err := s.Cname.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCnameRequestBucketCnameConfigurationCname struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteCnameRequestBucketCnameConfigurationCname) String() string {
	return dara.Prettify(s)
}

func (s DeleteCnameRequestBucketCnameConfigurationCname) GoString() string {
	return s.String()
}

func (s *DeleteCnameRequestBucketCnameConfigurationCname) GetDomain() *string {
	return s.Domain
}

func (s *DeleteCnameRequestBucketCnameConfigurationCname) SetDomain(v string) *DeleteCnameRequestBucketCnameConfigurationCname {
	s.Domain = &v
	return s
}

func (s *DeleteCnameRequestBucketCnameConfigurationCname) Validate() error {
	return dara.Validate(s)
}
