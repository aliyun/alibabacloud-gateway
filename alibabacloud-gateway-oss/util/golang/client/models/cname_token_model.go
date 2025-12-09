// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCnameToken interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *CnameToken
	GetBucket() *string
	SetCname(v string) *CnameToken
	GetCname() *string
	SetExpireTime(v string) *CnameToken
	GetExpireTime() *string
	SetToken(v string) *CnameToken
	GetToken() *string
}

type CnameToken struct {
	Bucket     *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Cname      *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CnameToken) String() string {
	return dara.Prettify(s)
}

func (s CnameToken) GoString() string {
	return s.String()
}

func (s *CnameToken) GetBucket() *string {
	return s.Bucket
}

func (s *CnameToken) GetCname() *string {
	return s.Cname
}

func (s *CnameToken) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CnameToken) GetToken() *string {
	return s.Token
}

func (s *CnameToken) SetBucket(v string) *CnameToken {
	s.Bucket = &v
	return s
}

func (s *CnameToken) SetCname(v string) *CnameToken {
	s.Cname = &v
	return s
}

func (s *CnameToken) SetExpireTime(v string) *CnameToken {
	s.ExpireTime = &v
	return s
}

func (s *CnameToken) SetToken(v string) *CnameToken {
	s.Token = &v
	return s
}

func (s *CnameToken) Validate() error {
	return dara.Validate(s)
}
