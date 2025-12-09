// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAntiDDosInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAntiDDOSListConfiguration(v *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) *GetUserAntiDDosInfoResponseBody
	GetAntiDDOSListConfiguration() *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration
}

type GetUserAntiDDosInfoResponseBody struct {
	// The container that stores the list of Anti-DDoS instances.
	AntiDDOSListConfiguration *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration `json:"AntiDDOSListConfiguration,omitempty" xml:"AntiDDOSListConfiguration,omitempty" type:"Struct"`
}

func (s GetUserAntiDDosInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAntiDDosInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAntiDDosInfoResponseBody) GetAntiDDOSListConfiguration() *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration {
	return s.AntiDDOSListConfiguration
}

func (s *GetUserAntiDDosInfoResponseBody) SetAntiDDOSListConfiguration(v *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) *GetUserAntiDDosInfoResponseBody {
	s.AntiDDOSListConfiguration = v
	return s
}

func (s *GetUserAntiDDosInfoResponseBody) Validate() error {
	if s.AntiDDOSListConfiguration != nil {
		if err := s.AntiDDOSListConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration struct {
	// The container that stores information about the Anti-DDoS instance.
	AntiDDOSConfiguration []*UserAntiDDOSInfo `json:"AntiDDOSConfiguration,omitempty" xml:"AntiDDOSConfiguration,omitempty" type:"Repeated"`
}

func (s GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) GoString() string {
	return s.String()
}

func (s *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) GetAntiDDOSConfiguration() []*UserAntiDDOSInfo {
	return s.AntiDDOSConfiguration
}

func (s *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) SetAntiDDOSConfiguration(v []*UserAntiDDOSInfo) *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration {
	s.AntiDDOSConfiguration = v
	return s
}

func (s *GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration) Validate() error {
	if s.AntiDDOSConfiguration != nil {
		for _, item := range s.AntiDDOSConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
