// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointResult interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointArn(v string) *CreateAccessPointResult
	GetAccessPointArn() *string
	SetAlias(v string) *CreateAccessPointResult
	GetAlias() *string
}

type CreateAccessPointResult struct {
	AccessPointArn *string `json:"AccessPointArn,omitempty" xml:"AccessPointArn,omitempty"`
	Alias          *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s CreateAccessPointResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointResult) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResult) GetAccessPointArn() *string {
	return s.AccessPointArn
}

func (s *CreateAccessPointResult) GetAlias() *string {
	return s.Alias
}

func (s *CreateAccessPointResult) SetAccessPointArn(v string) *CreateAccessPointResult {
	s.AccessPointArn = &v
	return s
}

func (s *CreateAccessPointResult) SetAlias(v string) *CreateAccessPointResult {
	s.Alias = &v
	return s
}

func (s *CreateAccessPointResult) Validate() error {
	return dara.Validate(s)
}
