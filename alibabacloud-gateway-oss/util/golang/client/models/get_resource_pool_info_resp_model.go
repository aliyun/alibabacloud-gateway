// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolInfoResp interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetResourcePoolInfoResp
	GetCreateTime() *string
	SetName(v string) *GetResourcePoolInfoResp
	GetName() *string
	SetOwner(v string) *GetResourcePoolInfoResp
	GetOwner() *string
	SetQosConfiguration(v *QoSConfiguration) *GetResourcePoolInfoResp
	GetQosConfiguration() *QoSConfiguration
	SetRegion(v string) *GetResourcePoolInfoResp
	GetRegion() *string
}

type GetResourcePoolInfoResp struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-07-24T08:42:32.000Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// rp-for-ai
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1234567890987
	Owner            *string           `json:"Owner,omitempty" xml:"Owner,omitempty"`
	QosConfiguration *QoSConfiguration `json:"QosConfiguration,omitempty" xml:"QosConfiguration,omitempty"`
	// example:
	//
	// oss-cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetResourcePoolInfoResp) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolInfoResp) GoString() string {
	return s.String()
}

func (s *GetResourcePoolInfoResp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetResourcePoolInfoResp) GetName() *string {
	return s.Name
}

func (s *GetResourcePoolInfoResp) GetOwner() *string {
	return s.Owner
}

func (s *GetResourcePoolInfoResp) GetQosConfiguration() *QoSConfiguration {
	return s.QosConfiguration
}

func (s *GetResourcePoolInfoResp) GetRegion() *string {
	return s.Region
}

func (s *GetResourcePoolInfoResp) SetCreateTime(v string) *GetResourcePoolInfoResp {
	s.CreateTime = &v
	return s
}

func (s *GetResourcePoolInfoResp) SetName(v string) *GetResourcePoolInfoResp {
	s.Name = &v
	return s
}

func (s *GetResourcePoolInfoResp) SetOwner(v string) *GetResourcePoolInfoResp {
	s.Owner = &v
	return s
}

func (s *GetResourcePoolInfoResp) SetQosConfiguration(v *QoSConfiguration) *GetResourcePoolInfoResp {
	s.QosConfiguration = v
	return s
}

func (s *GetResourcePoolInfoResp) SetRegion(v string) *GetResourcePoolInfoResp {
	s.Region = &v
	return s
}

func (s *GetResourcePoolInfoResp) Validate() error {
	if s.QosConfiguration != nil {
		if err := s.QosConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
