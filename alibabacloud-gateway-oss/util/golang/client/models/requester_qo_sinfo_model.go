// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequesterQoSInfo interface {
	dara.Model
	String() string
	GoString() string
	SetQoSConfiguration(v *QoSConfiguration) *RequesterQoSInfo
	GetQoSConfiguration() *QoSConfiguration
	SetRequester(v string) *RequesterQoSInfo
	GetRequester() *string
}

type RequesterQoSInfo struct {
	QoSConfiguration *QoSConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
	// example:
	//
	// 21234567890090
	Requester *string `json:"Requester,omitempty" xml:"Requester,omitempty"`
}

func (s RequesterQoSInfo) String() string {
	return dara.Prettify(s)
}

func (s RequesterQoSInfo) GoString() string {
	return s.String()
}

func (s *RequesterQoSInfo) GetQoSConfiguration() *QoSConfiguration {
	return s.QoSConfiguration
}

func (s *RequesterQoSInfo) GetRequester() *string {
	return s.Requester
}

func (s *RequesterQoSInfo) SetQoSConfiguration(v *QoSConfiguration) *RequesterQoSInfo {
	s.QoSConfiguration = v
	return s
}

func (s *RequesterQoSInfo) SetRequester(v string) *RequesterQoSInfo {
	s.Requester = &v
	return s
}

func (s *RequesterQoSInfo) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
