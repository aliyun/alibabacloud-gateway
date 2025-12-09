// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectsReq interface {
	dara.Model
	String() string
	GoString() string
	SetObject(v []*GetObjectsReqObject) *GetObjectsReq
	GetObject() []*GetObjectsReqObject
}

type GetObjectsReq struct {
	Object []*GetObjectsReqObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
}

func (s GetObjectsReq) String() string {
	return dara.Prettify(s)
}

func (s GetObjectsReq) GoString() string {
	return s.String()
}

func (s *GetObjectsReq) GetObject() []*GetObjectsReqObject {
	return s.Object
}

func (s *GetObjectsReq) SetObject(v []*GetObjectsReqObject) *GetObjectsReq {
	s.Object = v
	return s
}

func (s *GetObjectsReq) Validate() error {
	if s.Object != nil {
		for _, item := range s.Object {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetObjectsReqObject struct {
	// example:
	//
	// test.txt
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// example:
	//
	// bytes=20-60
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
	// example:
	//
	// 2
	RefId *int32 `json:"RefId,omitempty" xml:"RefId,omitempty"`
}

func (s GetObjectsReqObject) String() string {
	return dara.Prettify(s)
}

func (s GetObjectsReqObject) GoString() string {
	return s.String()
}

func (s *GetObjectsReqObject) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetObjectsReqObject) GetRange() *string {
	return s.Range
}

func (s *GetObjectsReqObject) GetRefId() *int32 {
	return s.RefId
}

func (s *GetObjectsReqObject) SetObjectName(v string) *GetObjectsReqObject {
	s.ObjectName = &v
	return s
}

func (s *GetObjectsReqObject) SetRange(v string) *GetObjectsReqObject {
	s.Range = &v
	return s
}

func (s *GetObjectsReqObject) SetRefId(v int32) *GetObjectsReqObject {
	s.RefId = &v
	return s
}

func (s *GetObjectsReqObject) Validate() error {
	return dara.Validate(s)
}
