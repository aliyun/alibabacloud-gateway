// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServerSideEncryptionRule interface {
	dara.Model
	String() string
	GoString() string
	SetApplyServerSideEncryptionByDefault(v *ApplyServerSideEncryptionByDefault) *ServerSideEncryptionRule
	GetApplyServerSideEncryptionByDefault() *ApplyServerSideEncryptionByDefault
}

type ServerSideEncryptionRule struct {
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefault `json:"ApplyServerSideEncryptionByDefault,omitempty" xml:"ApplyServerSideEncryptionByDefault,omitempty"`
}

func (s ServerSideEncryptionRule) String() string {
	return dara.Prettify(s)
}

func (s ServerSideEncryptionRule) GoString() string {
	return s.String()
}

func (s *ServerSideEncryptionRule) GetApplyServerSideEncryptionByDefault() *ApplyServerSideEncryptionByDefault {
	return s.ApplyServerSideEncryptionByDefault
}

func (s *ServerSideEncryptionRule) SetApplyServerSideEncryptionByDefault(v *ApplyServerSideEncryptionByDefault) *ServerSideEncryptionRule {
	s.ApplyServerSideEncryptionByDefault = v
	return s
}

func (s *ServerSideEncryptionRule) Validate() error {
	if s.ApplyServerSideEncryptionByDefault != nil {
		if err := s.ApplyServerSideEncryptionByDefault.Validate(); err != nil {
			return err
		}
	}
	return nil
}
