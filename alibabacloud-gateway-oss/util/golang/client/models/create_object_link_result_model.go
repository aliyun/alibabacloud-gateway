// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateObjectLinkResult interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *CreateObjectLinkResult
	GetBucket() *string
	SetETag(v string) *CreateObjectLinkResult
	GetETag() *string
	SetKey(v string) *CreateObjectLinkResult
	GetKey() *string
}

type CreateObjectLinkResult struct {
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// "5D3D4F3D1E6C690977E79E413C5F951D"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// ink-object
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s CreateObjectLinkResult) String() string {
	return dara.Prettify(s)
}

func (s CreateObjectLinkResult) GoString() string {
	return s.String()
}

func (s *CreateObjectLinkResult) GetBucket() *string {
	return s.Bucket
}

func (s *CreateObjectLinkResult) GetETag() *string {
	return s.ETag
}

func (s *CreateObjectLinkResult) GetKey() *string {
	return s.Key
}

func (s *CreateObjectLinkResult) SetBucket(v string) *CreateObjectLinkResult {
	s.Bucket = &v
	return s
}

func (s *CreateObjectLinkResult) SetETag(v string) *CreateObjectLinkResult {
	s.ETag = &v
	return s
}

func (s *CreateObjectLinkResult) SetKey(v string) *CreateObjectLinkResult {
	s.Key = &v
	return s
}

func (s *CreateObjectLinkResult) Validate() error {
	return dara.Validate(s)
}
