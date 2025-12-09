// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationReport interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *BatchOperationReport
	GetBucket() *string
	SetEnabled(v bool) *BatchOperationReport
	GetEnabled() *bool
	SetFormat(v string) *BatchOperationReport
	GetFormat() *string
	SetPrefix(v string) *BatchOperationReport
	GetPrefix() *string
	SetReportScope(v string) *BatchOperationReport
	GetReportScope() *string
}

type BatchOperationReport struct {
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
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// csv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// /
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// AllTasks
	ReportScope *string `json:"ReportScope,omitempty" xml:"ReportScope,omitempty"`
}

func (s BatchOperationReport) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationReport) GoString() string {
	return s.String()
}

func (s *BatchOperationReport) GetBucket() *string {
	return s.Bucket
}

func (s *BatchOperationReport) GetEnabled() *bool {
	return s.Enabled
}

func (s *BatchOperationReport) GetFormat() *string {
	return s.Format
}

func (s *BatchOperationReport) GetPrefix() *string {
	return s.Prefix
}

func (s *BatchOperationReport) GetReportScope() *string {
	return s.ReportScope
}

func (s *BatchOperationReport) SetBucket(v string) *BatchOperationReport {
	s.Bucket = &v
	return s
}

func (s *BatchOperationReport) SetEnabled(v bool) *BatchOperationReport {
	s.Enabled = &v
	return s
}

func (s *BatchOperationReport) SetFormat(v string) *BatchOperationReport {
	s.Format = &v
	return s
}

func (s *BatchOperationReport) SetPrefix(v string) *BatchOperationReport {
	s.Prefix = &v
	return s
}

func (s *BatchOperationReport) SetReportScope(v string) *BatchOperationReport {
	s.ReportScope = &v
	return s
}

func (s *BatchOperationReport) Validate() error {
	return dara.Validate(s)
}
