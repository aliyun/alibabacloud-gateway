// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameToken(v *CnameToken) *GetCnameTokenResponseBody
	GetCnameToken() *CnameToken
}

type GetCnameTokenResponseBody struct {
	// The container in which the CNAME token is stored.
	CnameToken *CnameToken `json:"CnameToken,omitempty" xml:"CnameToken,omitempty"`
}

func (s GetCnameTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCnameTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCnameTokenResponseBody) GetCnameToken() *CnameToken {
	return s.CnameToken
}

func (s *GetCnameTokenResponseBody) SetCnameToken(v *CnameToken) *GetCnameTokenResponseBody {
	s.CnameToken = v
	return s
}

func (s *GetCnameTokenResponseBody) Validate() error {
	if s.CnameToken != nil {
		if err := s.CnameToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}
