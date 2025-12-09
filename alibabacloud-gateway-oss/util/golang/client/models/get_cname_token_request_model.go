// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v string) *GetCnameTokenRequest
	GetCname() *string
}

type GetCnameTokenRequest struct {
	// The name of the CNAME record that is mapped to the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty"`
}

func (s GetCnameTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCnameTokenRequest) GoString() string {
	return s.String()
}

func (s *GetCnameTokenRequest) GetCname() *string {
	return s.Cname
}

func (s *GetCnameTokenRequest) SetCname(v string) *GetCnameTokenRequest {
	s.Cname = &v
	return s
}

func (s *GetCnameTokenRequest) Validate() error {
	return dara.Validate(s)
}
