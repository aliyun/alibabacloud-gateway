// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientRequestToken(v string) *BatchCreateJobRequest
	GetClientRequestToken() *string
	SetConfirmationRequired(v bool) *BatchCreateJobRequest
	GetConfirmationRequired() *bool
	SetDescription(v string) *BatchCreateJobRequest
	GetDescription() *string
	SetKeyPrefixManifestGenerator(v *BatchOperationManifestGenerator) *BatchCreateJobRequest
	GetKeyPrefixManifestGenerator() *BatchOperationManifestGenerator
	SetManifest(v *BatchOperationManifest) *BatchCreateJobRequest
	GetManifest() *BatchOperationManifest
	SetOperation(v *BatchOperation) *BatchCreateJobRequest
	GetOperation() *BatchOperation
	SetPriority(v int32) *BatchCreateJobRequest
	GetPriority() *int32
	SetReport(v *BatchOperationReport) *BatchCreateJobRequest
	GetReport() *BatchOperationReport
	SetRoleArn(v string) *BatchCreateJobRequest
	GetRoleArn() *string
}

type BatchCreateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456789012****
	ClientRequestToken *string `json:"ClientRequestToken,omitempty" xml:"ClientRequestToken,omitempty"`
	// example:
	//
	// true
	ConfirmationRequired       *bool                            `json:"ConfirmationRequired,omitempty" xml:"ConfirmationRequired,omitempty"`
	Description                *string                          `json:"Description,omitempty" xml:"Description,omitempty"`
	KeyPrefixManifestGenerator *BatchOperationManifestGenerator `json:"KeyPrefixManifestGenerator,omitempty" xml:"KeyPrefixManifestGenerator,omitempty"`
	Manifest                   *BatchOperationManifest          `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
	// This parameter is required.
	Operation *BatchOperation `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	Report *BatchOperationReport `json:"Report,omitempty" xml:"Report,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789012****:root
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s BatchCreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateJobRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateJobRequest) GetClientRequestToken() *string {
	return s.ClientRequestToken
}

func (s *BatchCreateJobRequest) GetConfirmationRequired() *bool {
	return s.ConfirmationRequired
}

func (s *BatchCreateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateJobRequest) GetKeyPrefixManifestGenerator() *BatchOperationManifestGenerator {
	return s.KeyPrefixManifestGenerator
}

func (s *BatchCreateJobRequest) GetManifest() *BatchOperationManifest {
	return s.Manifest
}

func (s *BatchCreateJobRequest) GetOperation() *BatchOperation {
	return s.Operation
}

func (s *BatchCreateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchCreateJobRequest) GetReport() *BatchOperationReport {
	return s.Report
}

func (s *BatchCreateJobRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *BatchCreateJobRequest) SetClientRequestToken(v string) *BatchCreateJobRequest {
	s.ClientRequestToken = &v
	return s
}

func (s *BatchCreateJobRequest) SetConfirmationRequired(v bool) *BatchCreateJobRequest {
	s.ConfirmationRequired = &v
	return s
}

func (s *BatchCreateJobRequest) SetDescription(v string) *BatchCreateJobRequest {
	s.Description = &v
	return s
}

func (s *BatchCreateJobRequest) SetKeyPrefixManifestGenerator(v *BatchOperationManifestGenerator) *BatchCreateJobRequest {
	s.KeyPrefixManifestGenerator = v
	return s
}

func (s *BatchCreateJobRequest) SetManifest(v *BatchOperationManifest) *BatchCreateJobRequest {
	s.Manifest = v
	return s
}

func (s *BatchCreateJobRequest) SetOperation(v *BatchOperation) *BatchCreateJobRequest {
	s.Operation = v
	return s
}

func (s *BatchCreateJobRequest) SetPriority(v int32) *BatchCreateJobRequest {
	s.Priority = &v
	return s
}

func (s *BatchCreateJobRequest) SetReport(v *BatchOperationReport) *BatchCreateJobRequest {
	s.Report = v
	return s
}

func (s *BatchCreateJobRequest) SetRoleArn(v string) *BatchCreateJobRequest {
	s.RoleArn = &v
	return s
}

func (s *BatchCreateJobRequest) Validate() error {
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
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	return nil
}
