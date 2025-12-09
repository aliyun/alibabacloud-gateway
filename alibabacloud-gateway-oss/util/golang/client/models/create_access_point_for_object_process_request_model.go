// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointForObjectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAccessPointForObjectProcessConfiguration(v *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) *CreateAccessPointForObjectProcessRequest
	GetCreateAccessPointForObjectProcessConfiguration() *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration
}

type CreateAccessPointForObjectProcessRequest struct {
	// The container that stores information about the Object FC Access Point.
	CreateAccessPointForObjectProcessConfiguration *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration `json:"CreateAccessPointForObjectProcessConfiguration,omitempty" xml:"CreateAccessPointForObjectProcessConfiguration,omitempty" type:"Struct"`
}

func (s CreateAccessPointForObjectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointForObjectProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessPointForObjectProcessRequest) GetCreateAccessPointForObjectProcessConfiguration() *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration {
	return s.CreateAccessPointForObjectProcessConfiguration
}

func (s *CreateAccessPointForObjectProcessRequest) SetCreateAccessPointForObjectProcessConfiguration(v *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) *CreateAccessPointForObjectProcessRequest {
	s.CreateAccessPointForObjectProcessConfiguration = v
	return s
}

func (s *CreateAccessPointForObjectProcessRequest) Validate() error {
	if s.CreateAccessPointForObjectProcessConfiguration != nil {
		if err := s.CreateAccessPointForObjectProcessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration struct {
	// The name of the access point.
	//
	// example:
	//
	// ap-01
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// Whether allow anonymous user to access this FC Access Point.
	//
	// example:
	//
	// disable
	AllowAnonymousAccessForObjectProcess *string `json:"AllowAnonymousAccessForObjectProcess,omitempty" xml:"AllowAnonymousAccessForObjectProcess,omitempty"`
	// The container that stores the processing information about the Object FC Access Point.
	ObjectProcessConfiguration *ObjectProcessConfiguration `json:"ObjectProcessConfiguration,omitempty" xml:"ObjectProcessConfiguration,omitempty"`
}

func (s CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) GetAllowAnonymousAccessForObjectProcess() *string {
	return s.AllowAnonymousAccessForObjectProcess
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) GetObjectProcessConfiguration() *ObjectProcessConfiguration {
	return s.ObjectProcessConfiguration
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) SetAccessPointName(v string) *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration {
	s.AccessPointName = &v
	return s
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) SetAllowAnonymousAccessForObjectProcess(v string) *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration {
	s.AllowAnonymousAccessForObjectProcess = &v
	return s
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) SetObjectProcessConfiguration(v *ObjectProcessConfiguration) *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration {
	s.ObjectProcessConfiguration = v
	return s
}

func (s *CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration) Validate() error {
	if s.ObjectProcessConfiguration != nil {
		if err := s.ObjectProcessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
