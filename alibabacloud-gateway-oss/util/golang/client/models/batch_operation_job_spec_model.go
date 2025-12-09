// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationJobSpec interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v string) *BatchOperationJobSpec
	GetFields() *string
	SetFormat(v string) *BatchOperationJobSpec
	GetFormat() *string
}

type BatchOperationJobSpec struct {
	// example:
	//
	// Bucket,Key,VersionId
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// csv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s BatchOperationJobSpec) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationJobSpec) GoString() string {
	return s.String()
}

func (s *BatchOperationJobSpec) GetFields() *string {
	return s.Fields
}

func (s *BatchOperationJobSpec) GetFormat() *string {
	return s.Format
}

func (s *BatchOperationJobSpec) SetFields(v string) *BatchOperationJobSpec {
	s.Fields = &v
	return s
}

func (s *BatchOperationJobSpec) SetFormat(v string) *BatchOperationJobSpec {
	s.Format = &v
	return s
}

func (s *BatchOperationJobSpec) Validate() error {
	return dara.Validate(s)
}
