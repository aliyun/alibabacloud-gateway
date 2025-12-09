// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBucketConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDataRedundancyType(v string) *CreateBucketConfiguration
	GetDataRedundancyType() *string
	SetStorageClass(v string) *CreateBucketConfiguration
	GetStorageClass() *string
}

type CreateBucketConfiguration struct {
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	StorageClass       *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s CreateBucketConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateBucketConfiguration) GoString() string {
	return s.String()
}

func (s *CreateBucketConfiguration) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *CreateBucketConfiguration) GetStorageClass() *string {
	return s.StorageClass
}

func (s *CreateBucketConfiguration) SetDataRedundancyType(v string) *CreateBucketConfiguration {
	s.DataRedundancyType = &v
	return s
}

func (s *CreateBucketConfiguration) SetStorageClass(v string) *CreateBucketConfiguration {
	s.StorageClass = &v
	return s
}

func (s *CreateBucketConfiguration) Validate() error {
	return dara.Validate(s)
}
