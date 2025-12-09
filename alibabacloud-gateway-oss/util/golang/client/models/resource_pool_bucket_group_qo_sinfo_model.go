// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourcePoolBucketGroupQoSInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucketGroup(v string) *ResourcePoolBucketGroupQoSInfo
	GetBucketGroup() *string
	SetQoSConfiguration(v *QoSConfiguration) *ResourcePoolBucketGroupQoSInfo
	GetQoSConfiguration() *QoSConfiguration
}

type ResourcePoolBucketGroupQoSInfo struct {
	// example:
	//
	// test-group
	BucketGroup      *string           `json:"BucketGroup,omitempty" xml:"BucketGroup,omitempty"`
	QoSConfiguration *QoSConfiguration `json:"QoSConfiguration,omitempty" xml:"QoSConfiguration,omitempty"`
}

func (s ResourcePoolBucketGroupQoSInfo) String() string {
	return dara.Prettify(s)
}

func (s ResourcePoolBucketGroupQoSInfo) GoString() string {
	return s.String()
}

func (s *ResourcePoolBucketGroupQoSInfo) GetBucketGroup() *string {
	return s.BucketGroup
}

func (s *ResourcePoolBucketGroupQoSInfo) GetQoSConfiguration() *QoSConfiguration {
	return s.QoSConfiguration
}

func (s *ResourcePoolBucketGroupQoSInfo) SetBucketGroup(v string) *ResourcePoolBucketGroupQoSInfo {
	s.BucketGroup = &v
	return s
}

func (s *ResourcePoolBucketGroupQoSInfo) SetQoSConfiguration(v *QoSConfiguration) *ResourcePoolBucketGroupQoSInfo {
	s.QoSConfiguration = v
	return s
}

func (s *ResourcePoolBucketGroupQoSInfo) Validate() error {
	if s.QoSConfiguration != nil {
		if err := s.QoSConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
