// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexDocument interface {
	dara.Model
	String() string
	GoString() string
	SetSuffix(v string) *IndexDocument
	GetSuffix() *string
	SetSupportSubDir(v bool) *IndexDocument
	GetSupportSubDir() *bool
	SetType(v int64) *IndexDocument
	GetType() *int64
}

type IndexDocument struct {
	// example:
	//
	// index.html
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	// example:
	//
	// true
	SupportSubDir *bool `json:"SupportSubDir,omitempty" xml:"SupportSubDir,omitempty"`
	// example:
	//
	// 0
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IndexDocument) String() string {
	return dara.Prettify(s)
}

func (s IndexDocument) GoString() string {
	return s.String()
}

func (s *IndexDocument) GetSuffix() *string {
	return s.Suffix
}

func (s *IndexDocument) GetSupportSubDir() *bool {
	return s.SupportSubDir
}

func (s *IndexDocument) GetType() *int64 {
	return s.Type
}

func (s *IndexDocument) SetSuffix(v string) *IndexDocument {
	s.Suffix = &v
	return s
}

func (s *IndexDocument) SetSupportSubDir(v bool) *IndexDocument {
	s.SupportSubDir = &v
	return s
}

func (s *IndexDocument) SetType(v int64) *IndexDocument {
	s.Type = &v
	return s
}

func (s *IndexDocument) Validate() error {
	return dara.Validate(s)
}
