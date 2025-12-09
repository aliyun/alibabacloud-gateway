// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyServerSideEncryptionByDefault interface {
	dara.Model
	String() string
	GoString() string
	SetKMSDataEncryption(v string) *ApplyServerSideEncryptionByDefault
	GetKMSDataEncryption() *string
	SetKMSMasterKeyID(v string) *ApplyServerSideEncryptionByDefault
	GetKMSMasterKeyID() *string
	SetSSEAlgorithm(v string) *ApplyServerSideEncryptionByDefault
	GetSSEAlgorithm() *string
}

type ApplyServerSideEncryptionByDefault struct {
	KMSDataEncryption *string `json:"KMSDataEncryption,omitempty" xml:"KMSDataEncryption,omitempty"`
	KMSMasterKeyID    *string `json:"KMSMasterKeyID,omitempty" xml:"KMSMasterKeyID,omitempty"`
	SSEAlgorithm      *string `json:"SSEAlgorithm,omitempty" xml:"SSEAlgorithm,omitempty"`
}

func (s ApplyServerSideEncryptionByDefault) String() string {
	return dara.Prettify(s)
}

func (s ApplyServerSideEncryptionByDefault) GoString() string {
	return s.String()
}

func (s *ApplyServerSideEncryptionByDefault) GetKMSDataEncryption() *string {
	return s.KMSDataEncryption
}

func (s *ApplyServerSideEncryptionByDefault) GetKMSMasterKeyID() *string {
	return s.KMSMasterKeyID
}

func (s *ApplyServerSideEncryptionByDefault) GetSSEAlgorithm() *string {
	return s.SSEAlgorithm
}

func (s *ApplyServerSideEncryptionByDefault) SetKMSDataEncryption(v string) *ApplyServerSideEncryptionByDefault {
	s.KMSDataEncryption = &v
	return s
}

func (s *ApplyServerSideEncryptionByDefault) SetKMSMasterKeyID(v string) *ApplyServerSideEncryptionByDefault {
	s.KMSMasterKeyID = &v
	return s
}

func (s *ApplyServerSideEncryptionByDefault) SetSSEAlgorithm(v string) *ApplyServerSideEncryptionByDefault {
	s.SSEAlgorithm = &v
	return s
}

func (s *ApplyServerSideEncryptionByDefault) Validate() error {
	return dara.Validate(s)
}
