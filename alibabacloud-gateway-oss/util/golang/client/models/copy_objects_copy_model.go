// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsCopy interface {
	dara.Model
	String() string
	GoString() string
	SetObject(v []*CopyObjectsCopyObject) *CopyObjectsCopy
	GetObject() []*CopyObjectsCopyObject
}

type CopyObjectsCopy struct {
	Object []*CopyObjectsCopyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
}

func (s CopyObjectsCopy) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsCopy) GoString() string {
	return s.String()
}

func (s *CopyObjectsCopy) GetObject() []*CopyObjectsCopyObject {
	return s.Object
}

func (s *CopyObjectsCopy) SetObject(v []*CopyObjectsCopyObject) *CopyObjectsCopy {
	s.Object = v
	return s
}

func (s *CopyObjectsCopy) Validate() error {
	if s.Object != nil {
		for _, item := range s.Object {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CopyObjectsCopyObject struct {
	// example:
	//
	// abc/source.txt
	SourceKey *string `json:"SourceKey,omitempty" xml:"SourceKey,omitempty"`
	// example:
	//
	// def/target.txt
	TargetKey *string `json:"TargetKey,omitempty" xml:"TargetKey,omitempty"`
}

func (s CopyObjectsCopyObject) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsCopyObject) GoString() string {
	return s.String()
}

func (s *CopyObjectsCopyObject) GetSourceKey() *string {
	return s.SourceKey
}

func (s *CopyObjectsCopyObject) GetTargetKey() *string {
	return s.TargetKey
}

func (s *CopyObjectsCopyObject) SetSourceKey(v string) *CopyObjectsCopyObject {
	s.SourceKey = &v
	return s
}

func (s *CopyObjectsCopyObject) SetTargetKey(v string) *CopyObjectsCopyObject {
	s.TargetKey = &v
	return s
}

func (s *CopyObjectsCopyObject) Validate() error {
	return dara.Validate(s)
}
