// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInventoryOSSBucketDestination interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *InventoryOSSBucketDestination
	GetAccountId() *string
	SetBucket(v string) *InventoryOSSBucketDestination
	GetBucket() *string
	SetEncryption(v *InventoryEncryption) *InventoryOSSBucketDestination
	GetEncryption() *InventoryEncryption
	SetFormat(v string) *InventoryOSSBucketDestination
	GetFormat() *string
	SetPrefix(v string) *InventoryOSSBucketDestination
	GetPrefix() *string
	SetRoleArn(v string) *InventoryOSSBucketDestination
	GetRoleArn() *string
}

type InventoryOSSBucketDestination struct {
	// example:
	//
	// 100000000000000
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// acs:oss:::bucket_0001
	Bucket     *string              `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Encryption *InventoryEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	Format     *string              `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// prefix1/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// acs:ram::100000000000000:role/AliyunOSSRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s InventoryOSSBucketDestination) String() string {
	return dara.Prettify(s)
}

func (s InventoryOSSBucketDestination) GoString() string {
	return s.String()
}

func (s *InventoryOSSBucketDestination) GetAccountId() *string {
	return s.AccountId
}

func (s *InventoryOSSBucketDestination) GetBucket() *string {
	return s.Bucket
}

func (s *InventoryOSSBucketDestination) GetEncryption() *InventoryEncryption {
	return s.Encryption
}

func (s *InventoryOSSBucketDestination) GetFormat() *string {
	return s.Format
}

func (s *InventoryOSSBucketDestination) GetPrefix() *string {
	return s.Prefix
}

func (s *InventoryOSSBucketDestination) GetRoleArn() *string {
	return s.RoleArn
}

func (s *InventoryOSSBucketDestination) SetAccountId(v string) *InventoryOSSBucketDestination {
	s.AccountId = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetBucket(v string) *InventoryOSSBucketDestination {
	s.Bucket = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetEncryption(v *InventoryEncryption) *InventoryOSSBucketDestination {
	s.Encryption = v
	return s
}

func (s *InventoryOSSBucketDestination) SetFormat(v string) *InventoryOSSBucketDestination {
	s.Format = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetPrefix(v string) *InventoryOSSBucketDestination {
	s.Prefix = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetRoleArn(v string) *InventoryOSSBucketDestination {
	s.RoleArn = &v
	return s
}

func (s *InventoryOSSBucketDestination) Validate() error {
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	return nil
}
