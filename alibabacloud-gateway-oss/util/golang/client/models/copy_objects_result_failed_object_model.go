// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsResultFailedObject interface {
	dara.Model
	String() string
	GoString() string
	SetErrorStatus(v string) *CopyObjectsResultFailedObject
	GetErrorStatus() *string
	SetSourceKey(v string) *CopyObjectsResultFailedObject
	GetSourceKey() *string
	SetTargetKey(v string) *CopyObjectsResultFailedObject
	GetTargetKey() *string
}

type CopyObjectsResultFailedObject struct {
	// example:
	//
	// NoSuchKey
	ErrorStatus *string `json:"ErrorStatus,omitempty" xml:"ErrorStatus,omitempty"`
	// example:
	//
	// abc/source.txt
	SourceKey *string `json:"SourceKey,omitempty" xml:"SourceKey,omitempty"`
	// example:
	//
	// abc/target.txt
	TargetKey *string `json:"TargetKey,omitempty" xml:"TargetKey,omitempty"`
}

func (s CopyObjectsResultFailedObject) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResultFailedObject) GoString() string {
	return s.String()
}

func (s *CopyObjectsResultFailedObject) GetErrorStatus() *string {
	return s.ErrorStatus
}

func (s *CopyObjectsResultFailedObject) GetSourceKey() *string {
	return s.SourceKey
}

func (s *CopyObjectsResultFailedObject) GetTargetKey() *string {
	return s.TargetKey
}

func (s *CopyObjectsResultFailedObject) SetErrorStatus(v string) *CopyObjectsResultFailedObject {
	s.ErrorStatus = &v
	return s
}

func (s *CopyObjectsResultFailedObject) SetSourceKey(v string) *CopyObjectsResultFailedObject {
	s.SourceKey = &v
	return s
}

func (s *CopyObjectsResultFailedObject) SetTargetKey(v string) *CopyObjectsResultFailedObject {
	s.TargetKey = &v
	return s
}

func (s *CopyObjectsResultFailedObject) Validate() error {
	return dara.Validate(s)
}
