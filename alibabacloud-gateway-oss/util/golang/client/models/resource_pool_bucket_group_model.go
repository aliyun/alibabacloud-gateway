// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourcePoolBucketGroup interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBucketInfo(v []*GroupBucketInfo) *ResourcePoolBucketGroup
	GetGroupBucketInfo() []*GroupBucketInfo
	SetName(v string) *ResourcePoolBucketGroup
	GetName() *string
}

type ResourcePoolBucketGroup struct {
	GroupBucketInfo []*GroupBucketInfo `json:"GroupBucketInfo,omitempty" xml:"GroupBucketInfo,omitempty" type:"Repeated"`
	// example:
	//
	// test-group-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ResourcePoolBucketGroup) String() string {
	return dara.Prettify(s)
}

func (s ResourcePoolBucketGroup) GoString() string {
	return s.String()
}

func (s *ResourcePoolBucketGroup) GetGroupBucketInfo() []*GroupBucketInfo {
	return s.GroupBucketInfo
}

func (s *ResourcePoolBucketGroup) GetName() *string {
	return s.Name
}

func (s *ResourcePoolBucketGroup) SetGroupBucketInfo(v []*GroupBucketInfo) *ResourcePoolBucketGroup {
	s.GroupBucketInfo = v
	return s
}

func (s *ResourcePoolBucketGroup) SetName(v string) *ResourcePoolBucketGroup {
	s.Name = &v
	return s
}

func (s *ResourcePoolBucketGroup) Validate() error {
	if s.GroupBucketInfo != nil {
		for _, item := range s.GroupBucketInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
