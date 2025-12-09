// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResult interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v int64) *DescribeJobResult
	GetCreationTime() *int64
	SetDescription(v string) *DescribeJobResult
	GetDescription() *string
	SetFailureReasons(v *BatchOperationFailureReasons) *DescribeJobResult
	GetFailureReasons() *BatchOperationFailureReasons
	SetJobId(v string) *DescribeJobResult
	GetJobId() *string
	SetKeyPrefixManifestGenerator(v *BatchOperationManifestGenerator) *DescribeJobResult
	GetKeyPrefixManifestGenerator() *BatchOperationManifestGenerator
	SetManifest(v *BatchOperationManifest) *DescribeJobResult
	GetManifest() *BatchOperationManifest
	SetOperation(v *BatchOperation) *DescribeJobResult
	GetOperation() *BatchOperation
	SetPriority(v int32) *DescribeJobResult
	GetPriority() *int32
	SetProgressSummary(v *BatchOperationJobProgressSummary) *DescribeJobResult
	GetProgressSummary() *BatchOperationJobProgressSummary
	SetReport(v *BatchOperationReport) *DescribeJobResult
	GetReport() *BatchOperationReport
	SetRoleArn(v string) *DescribeJobResult
	GetRoleArn() *string
	SetStatus(v string) *DescribeJobResult
	GetStatus() *string
	SetStatusUpdateReason(v string) *DescribeJobResult
	GetStatusUpdateReason() *string
	SetTerminationDate(v int64) *DescribeJobResult
	GetTerminationDate() *int64
}

type DescribeJobResult struct {
	// example:
	//
	// 1730250699
	CreationTime   *int64                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description    *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	FailureReasons *BatchOperationFailureReasons `json:"FailureReasons,omitempty" xml:"FailureReasons,omitempty"`
	// example:
	//
	// IDBkODc3M2RlZjgyNjQ0NDViZGV****
	JobId                      *string                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	KeyPrefixManifestGenerator *BatchOperationManifestGenerator `json:"KeyPrefixManifestGenerator,omitempty" xml:"KeyPrefixManifestGenerator,omitempty"`
	Manifest                   *BatchOperationManifest          `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
	Operation                  *BatchOperation                  `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 5
	Priority        *int32                            `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProgressSummary *BatchOperationJobProgressSummary `json:"ProgressSummary,omitempty" xml:"ProgressSummary,omitempty"`
	Report          *BatchOperationReport             `json:"Report,omitempty" xml:"Report,omitempty"`
	// example:
	//
	// acs:ram::174649585760****:role/AliyunOSSRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// Preparing
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusUpdateReason *string `json:"StatusUpdateReason,omitempty" xml:"StatusUpdateReason,omitempty"`
	// example:
	//
	// 1730250699
	TerminationDate *int64 `json:"TerminationDate,omitempty" xml:"TerminationDate,omitempty"`
}

func (s DescribeJobResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResult) GoString() string {
	return s.String()
}

func (s *DescribeJobResult) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeJobResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeJobResult) GetFailureReasons() *BatchOperationFailureReasons {
	return s.FailureReasons
}

func (s *DescribeJobResult) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResult) GetKeyPrefixManifestGenerator() *BatchOperationManifestGenerator {
	return s.KeyPrefixManifestGenerator
}

func (s *DescribeJobResult) GetManifest() *BatchOperationManifest {
	return s.Manifest
}

func (s *DescribeJobResult) GetOperation() *BatchOperation {
	return s.Operation
}

func (s *DescribeJobResult) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeJobResult) GetProgressSummary() *BatchOperationJobProgressSummary {
	return s.ProgressSummary
}

func (s *DescribeJobResult) GetReport() *BatchOperationReport {
	return s.Report
}

func (s *DescribeJobResult) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeJobResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobResult) GetStatusUpdateReason() *string {
	return s.StatusUpdateReason
}

func (s *DescribeJobResult) GetTerminationDate() *int64 {
	return s.TerminationDate
}

func (s *DescribeJobResult) SetCreationTime(v int64) *DescribeJobResult {
	s.CreationTime = &v
	return s
}

func (s *DescribeJobResult) SetDescription(v string) *DescribeJobResult {
	s.Description = &v
	return s
}

func (s *DescribeJobResult) SetFailureReasons(v *BatchOperationFailureReasons) *DescribeJobResult {
	s.FailureReasons = v
	return s
}

func (s *DescribeJobResult) SetJobId(v string) *DescribeJobResult {
	s.JobId = &v
	return s
}

func (s *DescribeJobResult) SetKeyPrefixManifestGenerator(v *BatchOperationManifestGenerator) *DescribeJobResult {
	s.KeyPrefixManifestGenerator = v
	return s
}

func (s *DescribeJobResult) SetManifest(v *BatchOperationManifest) *DescribeJobResult {
	s.Manifest = v
	return s
}

func (s *DescribeJobResult) SetOperation(v *BatchOperation) *DescribeJobResult {
	s.Operation = v
	return s
}

func (s *DescribeJobResult) SetPriority(v int32) *DescribeJobResult {
	s.Priority = &v
	return s
}

func (s *DescribeJobResult) SetProgressSummary(v *BatchOperationJobProgressSummary) *DescribeJobResult {
	s.ProgressSummary = v
	return s
}

func (s *DescribeJobResult) SetReport(v *BatchOperationReport) *DescribeJobResult {
	s.Report = v
	return s
}

func (s *DescribeJobResult) SetRoleArn(v string) *DescribeJobResult {
	s.RoleArn = &v
	return s
}

func (s *DescribeJobResult) SetStatus(v string) *DescribeJobResult {
	s.Status = &v
	return s
}

func (s *DescribeJobResult) SetStatusUpdateReason(v string) *DescribeJobResult {
	s.StatusUpdateReason = &v
	return s
}

func (s *DescribeJobResult) SetTerminationDate(v int64) *DescribeJobResult {
	s.TerminationDate = &v
	return s
}

func (s *DescribeJobResult) Validate() error {
	if s.FailureReasons != nil {
		if err := s.FailureReasons.Validate(); err != nil {
			return err
		}
	}
	if s.KeyPrefixManifestGenerator != nil {
		if err := s.KeyPrefixManifestGenerator.Validate(); err != nil {
			return err
		}
	}
	if s.Manifest != nil {
		if err := s.Manifest.Validate(); err != nil {
			return err
		}
	}
	if s.Operation != nil {
		if err := s.Operation.Validate(); err != nil {
			return err
		}
	}
	if s.ProgressSummary != nil {
		if err := s.ProgressSummary.Validate(); err != nil {
			return err
		}
	}
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	return nil
}
