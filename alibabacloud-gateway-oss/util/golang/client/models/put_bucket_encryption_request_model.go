// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketEncryptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServerSideEncryptionRule(v *ServerSideEncryptionRule) *PutBucketEncryptionRequest
	GetServerSideEncryptionRule() *ServerSideEncryptionRule
}

type PutBucketEncryptionRequest struct {
	// The container that stores server-side encryption rules.
	ServerSideEncryptionRule *ServerSideEncryptionRule `json:"ServerSideEncryptionRule,omitempty" xml:"ServerSideEncryptionRule,omitempty"`
}

func (s PutBucketEncryptionRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketEncryptionRequest) GoString() string {
	return s.String()
}

func (s *PutBucketEncryptionRequest) GetServerSideEncryptionRule() *ServerSideEncryptionRule {
	return s.ServerSideEncryptionRule
}

func (s *PutBucketEncryptionRequest) SetServerSideEncryptionRule(v *ServerSideEncryptionRule) *PutBucketEncryptionRequest {
	s.ServerSideEncryptionRule = v
	return s
}

func (s *PutBucketEncryptionRequest) Validate() error {
	if s.ServerSideEncryptionRule != nil {
		if err := s.ServerSideEncryptionRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
