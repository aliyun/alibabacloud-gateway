// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupBucketInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *GroupBucketInfo
	GetBucketName() *string
}

type GroupBucketInfo struct {
	// example:
	//
	// test-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s GroupBucketInfo) String() string {
	return dara.Prettify(s)
}

func (s GroupBucketInfo) GoString() string {
	return s.String()
}

func (s *GroupBucketInfo) GetBucketName() *string {
	return s.BucketName
}

func (s *GroupBucketInfo) SetBucketName(v string) *GroupBucketInfo {
	s.BucketName = &v
	return s
}

func (s *GroupBucketInfo) Validate() error {
	return dara.Validate(s)
}
