// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v *BucketInfoBucket) *BucketInfo
	GetBucket() *BucketInfoBucket
}

type BucketInfo struct {
	Bucket *BucketInfoBucket `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Struct"`
}

func (s BucketInfo) String() string {
	return dara.Prettify(s)
}

func (s BucketInfo) GoString() string {
	return s.String()
}

func (s *BucketInfo) GetBucket() *BucketInfoBucket {
	return s.Bucket
}

func (s *BucketInfo) SetBucket(v *BucketInfoBucket) *BucketInfo {
	s.Bucket = v
	return s
}

func (s *BucketInfo) Validate() error {
	if s.Bucket != nil {
		if err := s.Bucket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BucketInfoBucket struct {
	AccessControlList *AccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty"`
	// example:
	//
	// Disabled
	AccessMonitor *string `json:"AccessMonitor,omitempty" xml:"AccessMonitor,omitempty"`
	// example:
	//
	// false
	BlockPublicAccess *bool                         `json:"BlockPublicAccess,omitempty" xml:"BlockPublicAccess,omitempty"`
	BucketPolicy      *BucketInfoBucketBucketPolicy `json:"BucketPolicy,omitempty" xml:"BucketPolicy,omitempty" type:"Struct"`
	// example:
	//
	// An example bucket.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-01-06T08:20:09.000Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// example:
	//
	// Disabled
	CrossRegionReplication *string `json:"CrossRegionReplication,omitempty" xml:"CrossRegionReplication,omitempty"`
	DataRedundancyType     *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	ExtranetEndpoint       *string `json:"ExtranetEndpoint,omitempty" xml:"ExtranetEndpoint,omitempty"`
	IntranetEndpoint       *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// test-bucket
	Name                     *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner                    *Owner                                    `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ResourceGroupId          *string                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerSideEncryptionRule *BucketInfoBucketServerSideEncryptionRule `json:"ServerSideEncryptionRule,omitempty" xml:"ServerSideEncryptionRule,omitempty" type:"Struct"`
	StorageClass             *string                                   `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// example:
	//
	// Disabled
	TransferAcceleration *string `json:"TransferAcceleration,omitempty" xml:"TransferAcceleration,omitempty"`
	Versioning           *string `json:"Versioning,omitempty" xml:"Versioning,omitempty"`
}

func (s BucketInfoBucket) String() string {
	return dara.Prettify(s)
}

func (s BucketInfoBucket) GoString() string {
	return s.String()
}

func (s *BucketInfoBucket) GetAccessControlList() *AccessControlList {
	return s.AccessControlList
}

func (s *BucketInfoBucket) GetAccessMonitor() *string {
	return s.AccessMonitor
}

func (s *BucketInfoBucket) GetBlockPublicAccess() *bool {
	return s.BlockPublicAccess
}

func (s *BucketInfoBucket) GetBucketPolicy() *BucketInfoBucketBucketPolicy {
	return s.BucketPolicy
}

func (s *BucketInfoBucket) GetComment() *string {
	return s.Comment
}

func (s *BucketInfoBucket) GetCreationDate() *string {
	return s.CreationDate
}

func (s *BucketInfoBucket) GetCrossRegionReplication() *string {
	return s.CrossRegionReplication
}

func (s *BucketInfoBucket) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *BucketInfoBucket) GetExtranetEndpoint() *string {
	return s.ExtranetEndpoint
}

func (s *BucketInfoBucket) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *BucketInfoBucket) GetLocation() *string {
	return s.Location
}

func (s *BucketInfoBucket) GetName() *string {
	return s.Name
}

func (s *BucketInfoBucket) GetOwner() *Owner {
	return s.Owner
}

func (s *BucketInfoBucket) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BucketInfoBucket) GetServerSideEncryptionRule() *BucketInfoBucketServerSideEncryptionRule {
	return s.ServerSideEncryptionRule
}

func (s *BucketInfoBucket) GetStorageClass() *string {
	return s.StorageClass
}

func (s *BucketInfoBucket) GetTransferAcceleration() *string {
	return s.TransferAcceleration
}

func (s *BucketInfoBucket) GetVersioning() *string {
	return s.Versioning
}

func (s *BucketInfoBucket) SetAccessControlList(v *AccessControlList) *BucketInfoBucket {
	s.AccessControlList = v
	return s
}

func (s *BucketInfoBucket) SetAccessMonitor(v string) *BucketInfoBucket {
	s.AccessMonitor = &v
	return s
}

func (s *BucketInfoBucket) SetBlockPublicAccess(v bool) *BucketInfoBucket {
	s.BlockPublicAccess = &v
	return s
}

func (s *BucketInfoBucket) SetBucketPolicy(v *BucketInfoBucketBucketPolicy) *BucketInfoBucket {
	s.BucketPolicy = v
	return s
}

func (s *BucketInfoBucket) SetComment(v string) *BucketInfoBucket {
	s.Comment = &v
	return s
}

func (s *BucketInfoBucket) SetCreationDate(v string) *BucketInfoBucket {
	s.CreationDate = &v
	return s
}

func (s *BucketInfoBucket) SetCrossRegionReplication(v string) *BucketInfoBucket {
	s.CrossRegionReplication = &v
	return s
}

func (s *BucketInfoBucket) SetDataRedundancyType(v string) *BucketInfoBucket {
	s.DataRedundancyType = &v
	return s
}

func (s *BucketInfoBucket) SetExtranetEndpoint(v string) *BucketInfoBucket {
	s.ExtranetEndpoint = &v
	return s
}

func (s *BucketInfoBucket) SetIntranetEndpoint(v string) *BucketInfoBucket {
	s.IntranetEndpoint = &v
	return s
}

func (s *BucketInfoBucket) SetLocation(v string) *BucketInfoBucket {
	s.Location = &v
	return s
}

func (s *BucketInfoBucket) SetName(v string) *BucketInfoBucket {
	s.Name = &v
	return s
}

func (s *BucketInfoBucket) SetOwner(v *Owner) *BucketInfoBucket {
	s.Owner = v
	return s
}

func (s *BucketInfoBucket) SetResourceGroupId(v string) *BucketInfoBucket {
	s.ResourceGroupId = &v
	return s
}

func (s *BucketInfoBucket) SetServerSideEncryptionRule(v *BucketInfoBucketServerSideEncryptionRule) *BucketInfoBucket {
	s.ServerSideEncryptionRule = v
	return s
}

func (s *BucketInfoBucket) SetStorageClass(v string) *BucketInfoBucket {
	s.StorageClass = &v
	return s
}

func (s *BucketInfoBucket) SetTransferAcceleration(v string) *BucketInfoBucket {
	s.TransferAcceleration = &v
	return s
}

func (s *BucketInfoBucket) SetVersioning(v string) *BucketInfoBucket {
	s.Versioning = &v
	return s
}

func (s *BucketInfoBucket) Validate() error {
	if s.AccessControlList != nil {
		if err := s.AccessControlList.Validate(); err != nil {
			return err
		}
	}
	if s.BucketPolicy != nil {
		if err := s.BucketPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	if s.ServerSideEncryptionRule != nil {
		if err := s.ServerSideEncryptionRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BucketInfoBucketBucketPolicy struct {
	// example:
	//
	// example-bucket
	LogBucket *string `json:"LogBucket,omitempty" xml:"LogBucket,omitempty"`
	// example:
	//
	// log/
	LogPrefix *string `json:"LogPrefix,omitempty" xml:"LogPrefix,omitempty"`
}

func (s BucketInfoBucketBucketPolicy) String() string {
	return dara.Prettify(s)
}

func (s BucketInfoBucketBucketPolicy) GoString() string {
	return s.String()
}

func (s *BucketInfoBucketBucketPolicy) GetLogBucket() *string {
	return s.LogBucket
}

func (s *BucketInfoBucketBucketPolicy) GetLogPrefix() *string {
	return s.LogPrefix
}

func (s *BucketInfoBucketBucketPolicy) SetLogBucket(v string) *BucketInfoBucketBucketPolicy {
	s.LogBucket = &v
	return s
}

func (s *BucketInfoBucketBucketPolicy) SetLogPrefix(v string) *BucketInfoBucketBucketPolicy {
	s.LogPrefix = &v
	return s
}

func (s *BucketInfoBucketBucketPolicy) Validate() error {
	return dara.Validate(s)
}

type BucketInfoBucketServerSideEncryptionRule struct {
	// example:
	//
	// SM4
	KMSDataEncryption *string `json:"KMSDataEncryption,omitempty" xml:"KMSDataEncryption,omitempty"`
	// example:
	//
	// ****
	KMSMasterKeyID *string `json:"KMSMasterKeyID,omitempty" xml:"KMSMasterKeyID,omitempty"`
	// example:
	//
	// None
	SSEAlgorithm *string `json:"SSEAlgorithm,omitempty" xml:"SSEAlgorithm,omitempty"`
}

func (s BucketInfoBucketServerSideEncryptionRule) String() string {
	return dara.Prettify(s)
}

func (s BucketInfoBucketServerSideEncryptionRule) GoString() string {
	return s.String()
}

func (s *BucketInfoBucketServerSideEncryptionRule) GetKMSDataEncryption() *string {
	return s.KMSDataEncryption
}

func (s *BucketInfoBucketServerSideEncryptionRule) GetKMSMasterKeyID() *string {
	return s.KMSMasterKeyID
}

func (s *BucketInfoBucketServerSideEncryptionRule) GetSSEAlgorithm() *string {
	return s.SSEAlgorithm
}

func (s *BucketInfoBucketServerSideEncryptionRule) SetKMSDataEncryption(v string) *BucketInfoBucketServerSideEncryptionRule {
	s.KMSDataEncryption = &v
	return s
}

func (s *BucketInfoBucketServerSideEncryptionRule) SetKMSMasterKeyID(v string) *BucketInfoBucketServerSideEncryptionRule {
	s.KMSMasterKeyID = &v
	return s
}

func (s *BucketInfoBucketServerSideEncryptionRule) SetSSEAlgorithm(v string) *BucketInfoBucketServerSideEncryptionRule {
	s.SSEAlgorithm = &v
	return s
}

func (s *BucketInfoBucketServerSideEncryptionRule) Validate() error {
	return dara.Validate(s)
}
