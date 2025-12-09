// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointConfigForObjectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetAccessPointConfigForObjectProcessResult(v *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) *GetAccessPointConfigForObjectProcessResponseBody
	GetGetAccessPointConfigForObjectProcessResult() *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult
}

type GetAccessPointConfigForObjectProcessResponseBody struct {
	// The container that stores the configurations of the Object FC Access Point.
	GetAccessPointConfigForObjectProcessResult *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult `json:"GetAccessPointConfigForObjectProcessResult,omitempty" xml:"GetAccessPointConfigForObjectProcessResult,omitempty" type:"Struct"`
}

func (s GetAccessPointConfigForObjectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointConfigForObjectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPointConfigForObjectProcessResponseBody) GetGetAccessPointConfigForObjectProcessResult() *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult {
	return s.GetAccessPointConfigForObjectProcessResult
}

func (s *GetAccessPointConfigForObjectProcessResponseBody) SetGetAccessPointConfigForObjectProcessResult(v *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) *GetAccessPointConfigForObjectProcessResponseBody {
	s.GetAccessPointConfigForObjectProcessResult = v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponseBody) Validate() error {
	if s.GetAccessPointConfigForObjectProcessResult != nil {
		if err := s.GetAccessPointConfigForObjectProcessResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult struct {
	// Whether allow anonymous user to access this FC Access Points.
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

func (s GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) GoString() string {
	return s.String()
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) GetAllowAnonymousAccessForObjectProcess() *string {
	return s.AllowAnonymousAccessForObjectProcess
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) GetObjectProcessConfiguration() *ObjectProcessConfiguration {
	return s.ObjectProcessConfiguration
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) SetAllowAnonymousAccessForObjectProcess(v string) *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult {
	s.AllowAnonymousAccessForObjectProcess = &v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) SetObjectProcessConfiguration(v *ObjectProcessConfiguration) *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult {
	s.ObjectProcessConfiguration = v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult) Validate() error {
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
