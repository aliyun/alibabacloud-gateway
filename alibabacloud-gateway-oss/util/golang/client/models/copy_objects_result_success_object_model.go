// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsResultSuccessObject interface {
	dara.Model
	String() string
	GoString() string
	SetETag(v string) *CopyObjectsResultSuccessObject
	GetETag() *string
	SetSourceKey(v string) *CopyObjectsResultSuccessObject
	GetSourceKey() *string
	SetTargetKey(v string) *CopyObjectsResultSuccessObject
	GetTargetKey() *string
}

type CopyObjectsResultSuccessObject struct {
	// example:
	//
	// 5EB63BBBE01EEED093CB22BB8F5ACDC3
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// abc/source.txt
	SourceKey *string `json:"SourceKey,omitempty" xml:"SourceKey,omitempty"`
	// example:
	//
	// abc/target.txt
	TargetKey *string `json:"TargetKey,omitempty" xml:"TargetKey,omitempty"`
}

func (s CopyObjectsResultSuccessObject) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResultSuccessObject) GoString() string {
	return s.String()
}

func (s *CopyObjectsResultSuccessObject) GetETag() *string {
	return s.ETag
}

func (s *CopyObjectsResultSuccessObject) GetSourceKey() *string {
	return s.SourceKey
}

func (s *CopyObjectsResultSuccessObject) GetTargetKey() *string {
	return s.TargetKey
}

func (s *CopyObjectsResultSuccessObject) SetETag(v string) *CopyObjectsResultSuccessObject {
	s.ETag = &v
	return s
}

func (s *CopyObjectsResultSuccessObject) SetSourceKey(v string) *CopyObjectsResultSuccessObject {
	s.SourceKey = &v
	return s
}

func (s *CopyObjectsResultSuccessObject) SetTargetKey(v string) *CopyObjectsResultSuccessObject {
	s.TargetKey = &v
	return s
}

func (s *CopyObjectsResultSuccessObject) Validate() error {
	return dara.Validate(s)
}
