// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointConfigForObjectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPutAccessPointConfigForObjectProcessConfiguration(v *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) *PutAccessPointConfigForObjectProcessRequest
	GetPutAccessPointConfigForObjectProcessConfiguration() *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration
}

type PutAccessPointConfigForObjectProcessRequest struct {
	// The container that stores information about the Object FC Access Point.
	PutAccessPointConfigForObjectProcessConfiguration *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration `json:"PutAccessPointConfigForObjectProcessConfiguration,omitempty" xml:"PutAccessPointConfigForObjectProcessConfiguration,omitempty" type:"Struct"`
}

func (s PutAccessPointConfigForObjectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointConfigForObjectProcessRequest) GoString() string {
	return s.String()
}

func (s *PutAccessPointConfigForObjectProcessRequest) GetPutAccessPointConfigForObjectProcessConfiguration() *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration {
	return s.PutAccessPointConfigForObjectProcessConfiguration
}

func (s *PutAccessPointConfigForObjectProcessRequest) SetPutAccessPointConfigForObjectProcessConfiguration(v *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) *PutAccessPointConfigForObjectProcessRequest {
	s.PutAccessPointConfigForObjectProcessConfiguration = v
	return s
}

func (s *PutAccessPointConfigForObjectProcessRequest) Validate() error {
	if s.PutAccessPointConfigForObjectProcessConfiguration != nil {
		if err := s.PutAccessPointConfigForObjectProcessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration struct {
	// Whether allow anonymous user to access this FC Access Point.
	//
	// example:
	//
	// false
	AllowAnonymousAccessForObjectProcess *string `json:"AllowAnonymousAccessForObjectProcess,omitempty" xml:"AllowAnonymousAccessForObjectProcess,omitempty"`
	// The container that stores the processing information about the Object FC Access Point.
	ObjectProcessConfiguration *ObjectProcessConfiguration `json:"ObjectProcessConfiguration,omitempty" xml:"ObjectProcessConfiguration,omitempty"`
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
}

func (s PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) GoString() string {
	return s.String()
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) GetAllowAnonymousAccessForObjectProcess() *string {
	return s.AllowAnonymousAccessForObjectProcess
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) GetObjectProcessConfiguration() *ObjectProcessConfiguration {
	return s.ObjectProcessConfiguration
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) SetAllowAnonymousAccessForObjectProcess(v string) *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration {
	s.AllowAnonymousAccessForObjectProcess = &v
	return s
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) SetObjectProcessConfiguration(v *ObjectProcessConfiguration) *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration {
	s.ObjectProcessConfiguration = v
	return s
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration) Validate() error {
	if s.ObjectProcessConfiguration != nil {
		if err := s.ObjectProcessConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
