// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCopyObject interface {
	dara.Model
	String() string
	GoString() string
	SetMetadataDirective(v string) *BatchCopyObject
	GetMetadataDirective() *string
	SetNewObjectMetadata(v string) *BatchCopyObject
	GetNewObjectMetadata() *string
	SetNewObjectTagging(v string) *BatchCopyObject
	GetNewObjectTagging() *string
	SetObjectAcl(v string) *BatchCopyObject
	GetObjectAcl() *string
	SetServerSideDataEncryption(v string) *BatchCopyObject
	GetServerSideDataEncryption() *string
	SetServerSideEncryption(v string) *BatchCopyObject
	GetServerSideEncryption() *string
	SetStorageClass(v string) *BatchCopyObject
	GetStorageClass() *string
	SetTaggingDirective(v string) *BatchCopyObject
	GetTaggingDirective() *string
	SetTargetKeyPrefix(v string) *BatchCopyObject
	GetTargetKeyPrefix() *string
	SetTargetResource(v string) *BatchCopyObject
	GetTargetResource() *string
}

type BatchCopyObject struct {
	// example:
	//
	// REPLACE
	MetadataDirective *string `json:"MetadataDirective,omitempty" xml:"MetadataDirective,omitempty"`
	NewObjectMetadata *string `json:"NewObjectMetadata,omitempty" xml:"NewObjectMetadata,omitempty"`
	// example:
	//
	// TagA=A&TagB=B
	NewObjectTagging *string `json:"NewObjectTagging,omitempty" xml:"NewObjectTagging,omitempty"`
	ObjectAcl        *string `json:"ObjectAcl,omitempty" xml:"ObjectAcl,omitempty"`
	// example:
	//
	// SM4
	ServerSideDataEncryption *string `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	StorageClass         *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// example:
	//
	// REPLACE
	TaggingDirective *string `json:"TaggingDirective,omitempty" xml:"TaggingDirective,omitempty"`
	// example:
	//
	// /xxx
	TargetKeyPrefix *string `json:"TargetKeyPrefix,omitempty" xml:"TargetKeyPrefix,omitempty"`
	// example:
	//
	// bucketname
	TargetResource *string `json:"TargetResource,omitempty" xml:"TargetResource,omitempty"`
}

func (s BatchCopyObject) String() string {
	return dara.Prettify(s)
}

func (s BatchCopyObject) GoString() string {
	return s.String()
}

func (s *BatchCopyObject) GetMetadataDirective() *string {
	return s.MetadataDirective
}

func (s *BatchCopyObject) GetNewObjectMetadata() *string {
	return s.NewObjectMetadata
}

func (s *BatchCopyObject) GetNewObjectTagging() *string {
	return s.NewObjectTagging
}

func (s *BatchCopyObject) GetObjectAcl() *string {
	return s.ObjectAcl
}

func (s *BatchCopyObject) GetServerSideDataEncryption() *string {
	return s.ServerSideDataEncryption
}

func (s *BatchCopyObject) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *BatchCopyObject) GetStorageClass() *string {
	return s.StorageClass
}

func (s *BatchCopyObject) GetTaggingDirective() *string {
	return s.TaggingDirective
}

func (s *BatchCopyObject) GetTargetKeyPrefix() *string {
	return s.TargetKeyPrefix
}

func (s *BatchCopyObject) GetTargetResource() *string {
	return s.TargetResource
}

func (s *BatchCopyObject) SetMetadataDirective(v string) *BatchCopyObject {
	s.MetadataDirective = &v
	return s
}

func (s *BatchCopyObject) SetNewObjectMetadata(v string) *BatchCopyObject {
	s.NewObjectMetadata = &v
	return s
}

func (s *BatchCopyObject) SetNewObjectTagging(v string) *BatchCopyObject {
	s.NewObjectTagging = &v
	return s
}

func (s *BatchCopyObject) SetObjectAcl(v string) *BatchCopyObject {
	s.ObjectAcl = &v
	return s
}

func (s *BatchCopyObject) SetServerSideDataEncryption(v string) *BatchCopyObject {
	s.ServerSideDataEncryption = &v
	return s
}

func (s *BatchCopyObject) SetServerSideEncryption(v string) *BatchCopyObject {
	s.ServerSideEncryption = &v
	return s
}

func (s *BatchCopyObject) SetStorageClass(v string) *BatchCopyObject {
	s.StorageClass = &v
	return s
}

func (s *BatchCopyObject) SetTaggingDirective(v string) *BatchCopyObject {
	s.TaggingDirective = &v
	return s
}

func (s *BatchCopyObject) SetTargetKeyPrefix(v string) *BatchCopyObject {
	s.TargetKeyPrefix = &v
	return s
}

func (s *BatchCopyObject) SetTargetResource(v string) *BatchCopyObject {
	s.TargetResource = &v
	return s
}

func (s *BatchCopyObject) Validate() error {
	return dara.Validate(s)
}
