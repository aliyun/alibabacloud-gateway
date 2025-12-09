// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSSEKMS interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *SSEKMS
	GetKeyId() *string
}

type SSEKMS struct {
	// if can be null:
	// true
	//
	// example:
	//
	// abcd
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s SSEKMS) String() string {
	return dara.Prettify(s)
}

func (s SSEKMS) GoString() string {
	return s.String()
}

func (s *SSEKMS) GetKeyId() *string {
	return s.KeyId
}

func (s *SSEKMS) SetKeyId(v string) *SSEKMS {
	s.KeyId = &v
	return s
}

func (s *SSEKMS) Validate() error {
	return dara.Validate(s)
}
