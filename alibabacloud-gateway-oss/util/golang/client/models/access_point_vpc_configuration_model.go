// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessPointVpcConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *AccessPointVpcConfiguration
	GetVpcId() *string
}

type AccessPointVpcConfiguration struct {
	// example:
	//
	// vpc-t4nlw426y44rd3iq4****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AccessPointVpcConfiguration) String() string {
	return dara.Prettify(s)
}

func (s AccessPointVpcConfiguration) GoString() string {
	return s.String()
}

func (s *AccessPointVpcConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *AccessPointVpcConfiguration) SetVpcId(v string) *AccessPointVpcConfiguration {
	s.VpcId = &v
	return s
}

func (s *AccessPointVpcConfiguration) Validate() error {
	return dara.Validate(s)
}
