// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointForObjectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAccessPointForObjectProcessResult(v *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) *CreateAccessPointForObjectProcessResponseBody
	GetCreateAccessPointForObjectProcessResult() *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult
}

type CreateAccessPointForObjectProcessResponseBody struct {
	// The container that stores information about the Object FC Access Point.
	CreateAccessPointForObjectProcessResult *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult `json:"CreateAccessPointForObjectProcessResult,omitempty" xml:"CreateAccessPointForObjectProcessResult,omitempty" type:"Struct"`
}

func (s CreateAccessPointForObjectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointForObjectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessPointForObjectProcessResponseBody) GetCreateAccessPointForObjectProcessResult() *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult {
	return s.CreateAccessPointForObjectProcessResult
}

func (s *CreateAccessPointForObjectProcessResponseBody) SetCreateAccessPointForObjectProcessResult(v *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) *CreateAccessPointForObjectProcessResponseBody {
	s.CreateAccessPointForObjectProcessResult = v
	return s
}

func (s *CreateAccessPointForObjectProcessResponseBody) Validate() error {
	if s.CreateAccessPointForObjectProcessResult != nil {
		if err := s.CreateAccessPointForObjectProcessResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult struct {
	// The alias of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01-3b00521f653d2b3223680ec39dbbe2****-opapalias
	AccessPointForObjectProcessAlias *string `json:"AccessPointForObjectProcessAlias,omitempty" xml:"AccessPointForObjectProcessAlias,omitempty"`
	// The ARN of the Object FC Access Point.
	//
	// example:
	//
	// acs:oss:cn-qingdao:119335441657143:accesspointforobjectprocess/fc-ap-01
	AccessPointForObjectProcessArn *string `json:"AccessPointForObjectProcessArn,omitempty" xml:"AccessPointForObjectProcessArn,omitempty"`
}

func (s CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) GoString() string {
	return s.String()
}

func (s *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) GetAccessPointForObjectProcessAlias() *string {
	return s.AccessPointForObjectProcessAlias
}

func (s *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) GetAccessPointForObjectProcessArn() *string {
	return s.AccessPointForObjectProcessArn
}

func (s *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) SetAccessPointForObjectProcessAlias(v string) *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult {
	s.AccessPointForObjectProcessAlias = &v
	return s
}

func (s *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) SetAccessPointForObjectProcessArn(v string) *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult {
	s.AccessPointForObjectProcessArn = &v
	return s
}

func (s *CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult) Validate() error {
	return dara.Validate(s)
}
