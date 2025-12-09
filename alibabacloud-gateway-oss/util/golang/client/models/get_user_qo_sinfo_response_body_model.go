// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserQoSInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQoSConfiguration(v *UserQosConfiguration) *GetUserQoSInfoResponseBody
	GetQoSConfiguration() *UserQosConfiguration
}

type GetUserQoSInfoResponseBody struct {
	QoSConfiguration *UserQosConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
}

func (s GetUserQoSInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserQoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserQoSInfoResponseBody) GetQoSConfiguration() *UserQosConfiguration {
	return s.QoSConfiguration
}

func (s *GetUserQoSInfoResponseBody) SetQoSConfiguration(v *UserQosConfiguration) *GetUserQoSInfoResponseBody {
	s.QoSConfiguration = v
	return s
}

func (s *GetUserQoSInfoResponseBody) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
