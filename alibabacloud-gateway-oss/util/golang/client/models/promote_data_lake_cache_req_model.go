// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteDataLakeCacheReq interface {
	dara.Model
	String() string
	GoString() string
	SetObject(v *PromoteDataLakeCacheReqObject) *PromoteDataLakeCacheReq
	GetObject() *PromoteDataLakeCacheReqObject
}

type PromoteDataLakeCacheReq struct {
	Object *PromoteDataLakeCacheReqObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
}

func (s PromoteDataLakeCacheReq) String() string {
	return dara.Prettify(s)
}

func (s PromoteDataLakeCacheReq) GoString() string {
	return s.String()
}

func (s *PromoteDataLakeCacheReq) GetObject() *PromoteDataLakeCacheReqObject {
	return s.Object
}

func (s *PromoteDataLakeCacheReq) SetObject(v *PromoteDataLakeCacheReqObject) *PromoteDataLakeCacheReq {
	s.Object = v
	return s
}

func (s *PromoteDataLakeCacheReq) Validate() error {
	if s.Object != nil {
		if err := s.Object.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PromoteDataLakeCacheReqObject struct {
	// example:
	//
	// test.data
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// example:
	//
	// 1024-2048
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s PromoteDataLakeCacheReqObject) String() string {
	return dara.Prettify(s)
}

func (s PromoteDataLakeCacheReqObject) GoString() string {
	return s.String()
}

func (s *PromoteDataLakeCacheReqObject) GetObjectName() *string {
	return s.ObjectName
}

func (s *PromoteDataLakeCacheReqObject) GetRange() *string {
	return s.Range
}

func (s *PromoteDataLakeCacheReqObject) SetObjectName(v string) *PromoteDataLakeCacheReqObject {
	s.ObjectName = &v
	return s
}

func (s *PromoteDataLakeCacheReqObject) SetRange(v string) *PromoteDataLakeCacheReqObject {
	s.Range = &v
	return s
}

func (s *PromoteDataLakeCacheReqObject) Validate() error {
	return dara.Validate(s)
}
