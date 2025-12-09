// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourcePoolBucket interface {
	dara.Model
	String() string
	GoString() string
	SetJoinTime(v string) *ResourcePoolBucket
	GetJoinTime() *string
	SetName(v string) *ResourcePoolBucket
	GetName() *string
}

type ResourcePoolBucket struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-07-24T08:42:32.000Z
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// example:
	//
	// test-bucket
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ResourcePoolBucket) String() string {
	return dara.Prettify(s)
}

func (s ResourcePoolBucket) GoString() string {
	return s.String()
}

func (s *ResourcePoolBucket) GetJoinTime() *string {
	return s.JoinTime
}

func (s *ResourcePoolBucket) GetName() *string {
	return s.Name
}

func (s *ResourcePoolBucket) SetJoinTime(v string) *ResourcePoolBucket {
	s.JoinTime = &v
	return s
}

func (s *ResourcePoolBucket) SetName(v string) *ResourcePoolBucket {
	s.Name = &v
	return s
}

func (s *ResourcePoolBucket) Validate() error {
	return dara.Validate(s)
}
