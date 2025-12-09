// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCnameTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameToken(v *CnameToken) *CreateCnameTokenResponseBody
	GetCnameToken() *CnameToken
}

type CreateCnameTokenResponseBody struct {
	// The container in which the CNAME token is stored.
	CnameToken *CnameToken `json:"CnameToken,omitempty" xml:"CnameToken,omitempty"`
}

func (s CreateCnameTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCnameTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCnameTokenResponseBody) GetCnameToken() *CnameToken {
	return s.CnameToken
}

func (s *CreateCnameTokenResponseBody) SetCnameToken(v *CnameToken) *CreateCnameTokenResponseBody {
	s.CnameToken = v
	return s
}

func (s *CreateCnameTokenResponseBody) Validate() error {
	if s.CnameToken != nil {
		if err := s.CnameToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}
