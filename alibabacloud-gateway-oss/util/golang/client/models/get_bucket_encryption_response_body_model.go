// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketEncryptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetServerSideEncryptionRule(v *GetBucketEncryptionResponseBodyServerSideEncryptionRule) *GetBucketEncryptionResponseBody
	GetServerSideEncryptionRule() *GetBucketEncryptionResponseBodyServerSideEncryptionRule
}

type GetBucketEncryptionResponseBody struct {
	// The container that stores server-side encryption rules.
	ServerSideEncryptionRule *GetBucketEncryptionResponseBodyServerSideEncryptionRule `json:"ServerSideEncryptionRule,omitempty" xml:"ServerSideEncryptionRule,omitempty" type:"Struct"`
}

func (s GetBucketEncryptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketEncryptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketEncryptionResponseBody) GetServerSideEncryptionRule() *GetBucketEncryptionResponseBodyServerSideEncryptionRule {
	return s.ServerSideEncryptionRule
}

func (s *GetBucketEncryptionResponseBody) SetServerSideEncryptionRule(v *GetBucketEncryptionResponseBodyServerSideEncryptionRule) *GetBucketEncryptionResponseBody {
	s.ServerSideEncryptionRule = v
	return s
}

func (s *GetBucketEncryptionResponseBody) Validate() error {
	if s.ServerSideEncryptionRule != nil {
		if err := s.ServerSideEncryptionRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketEncryptionResponseBodyServerSideEncryptionRule struct {
	// The container that stores the default server-side encryption method.
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefault `json:"ApplyServerSideEncryptionByDefault,omitempty" xml:"ApplyServerSideEncryptionByDefault,omitempty"`
}

func (s GetBucketEncryptionResponseBodyServerSideEncryptionRule) String() string {
	return dara.Prettify(s)
}

func (s GetBucketEncryptionResponseBodyServerSideEncryptionRule) GoString() string {
	return s.String()
}

func (s *GetBucketEncryptionResponseBodyServerSideEncryptionRule) GetApplyServerSideEncryptionByDefault() *ApplyServerSideEncryptionByDefault {
	return s.ApplyServerSideEncryptionByDefault
}

func (s *GetBucketEncryptionResponseBodyServerSideEncryptionRule) SetApplyServerSideEncryptionByDefault(v *ApplyServerSideEncryptionByDefault) *GetBucketEncryptionResponseBodyServerSideEncryptionRule {
	s.ApplyServerSideEncryptionByDefault = v
	return s
}

func (s *GetBucketEncryptionResponseBodyServerSideEncryptionRule) Validate() error {
	if s.ApplyServerSideEncryptionByDefault != nil {
		if err := s.ApplyServerSideEncryptionByDefault.Validate(); err != nil {
			return err
		}
	}
	return nil
}
