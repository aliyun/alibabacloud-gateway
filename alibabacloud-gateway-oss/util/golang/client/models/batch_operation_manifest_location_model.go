// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationManifestLocation interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *BatchOperationManifestLocation
	GetBucket() *string
	SetObject(v string) *BatchOperationManifestLocation
	GetObject() *string
	SetVersionId(v string) *BatchOperationManifestLocation
	GetVersionId() *string
}

type BatchOperationManifestLocation struct {
	// This parameter is required.
	//
	// example:
	//
	// bucketname
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// objectname
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// ********gICx0puvmxkiIDBkODc3M2RlZjgyNjQ0NDViZGVlYmRmNzI2********
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s BatchOperationManifestLocation) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationManifestLocation) GoString() string {
	return s.String()
}

func (s *BatchOperationManifestLocation) GetBucket() *string {
	return s.Bucket
}

func (s *BatchOperationManifestLocation) GetObject() *string {
	return s.Object
}

func (s *BatchOperationManifestLocation) GetVersionId() *string {
	return s.VersionId
}

func (s *BatchOperationManifestLocation) SetBucket(v string) *BatchOperationManifestLocation {
	s.Bucket = &v
	return s
}

func (s *BatchOperationManifestLocation) SetObject(v string) *BatchOperationManifestLocation {
	s.Object = &v
	return s
}

func (s *BatchOperationManifestLocation) SetVersionId(v string) *BatchOperationManifestLocation {
	s.VersionId = &v
	return s
}

func (s *BatchOperationManifestLocation) Validate() error {
	return dara.Validate(s)
}
