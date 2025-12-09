// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationManifest interface {
	dara.Model
	String() string
	GoString() string
	SetLocation(v *BatchOperationManifestLocation) *BatchOperationManifest
	GetLocation() *BatchOperationManifestLocation
	SetSpec(v *BatchOperationJobSpec) *BatchOperationManifest
	GetSpec() *BatchOperationJobSpec
}

type BatchOperationManifest struct {
	// This parameter is required.
	Location *BatchOperationManifestLocation `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	Spec *BatchOperationJobSpec `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s BatchOperationManifest) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationManifest) GoString() string {
	return s.String()
}

func (s *BatchOperationManifest) GetLocation() *BatchOperationManifestLocation {
	return s.Location
}

func (s *BatchOperationManifest) GetSpec() *BatchOperationJobSpec {
	return s.Spec
}

func (s *BatchOperationManifest) SetLocation(v *BatchOperationManifestLocation) *BatchOperationManifest {
	s.Location = v
	return s
}

func (s *BatchOperationManifest) SetSpec(v *BatchOperationJobSpec) *BatchOperationManifest {
	s.Spec = v
	return s
}

func (s *BatchOperationManifest) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
