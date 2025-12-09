// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsResult interface {
	dara.Model
	String() string
	GoString() string
	SetFailed(v *CopyObjectsResultFailed) *CopyObjectsResult
	GetFailed() *CopyObjectsResultFailed
	SetSuccess(v *CopyObjectsResultSuccess) *CopyObjectsResult
	GetSuccess() *CopyObjectsResultSuccess
}

type CopyObjectsResult struct {
	Failed  *CopyObjectsResultFailed  `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Struct"`
	Success *CopyObjectsResultSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Struct"`
}

func (s CopyObjectsResult) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResult) GoString() string {
	return s.String()
}

func (s *CopyObjectsResult) GetFailed() *CopyObjectsResultFailed {
	return s.Failed
}

func (s *CopyObjectsResult) GetSuccess() *CopyObjectsResultSuccess {
	return s.Success
}

func (s *CopyObjectsResult) SetFailed(v *CopyObjectsResultFailed) *CopyObjectsResult {
	s.Failed = v
	return s
}

func (s *CopyObjectsResult) SetSuccess(v *CopyObjectsResultSuccess) *CopyObjectsResult {
	s.Success = v
	return s
}

func (s *CopyObjectsResult) Validate() error {
	if s.Failed != nil {
		if err := s.Failed.Validate(); err != nil {
			return err
		}
	}
	if s.Success != nil {
		if err := s.Success.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyObjectsResultFailed struct {
	Object []*CopyObjectsResultFailedObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
}

func (s CopyObjectsResultFailed) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResultFailed) GoString() string {
	return s.String()
}

func (s *CopyObjectsResultFailed) GetObject() []*CopyObjectsResultFailedObject {
	return s.Object
}

func (s *CopyObjectsResultFailed) SetObject(v []*CopyObjectsResultFailedObject) *CopyObjectsResultFailed {
	s.Object = v
	return s
}

func (s *CopyObjectsResultFailed) Validate() error {
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

type CopyObjectsResultSuccess struct {
	Object []*CopyObjectsResultSuccessObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
}

func (s CopyObjectsResultSuccess) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResultSuccess) GoString() string {
	return s.String()
}

func (s *CopyObjectsResultSuccess) GetObject() []*CopyObjectsResultSuccessObject {
	return s.Object
}

func (s *CopyObjectsResultSuccess) SetObject(v []*CopyObjectsResultSuccessObject) *CopyObjectsResultSuccess {
	s.Object = v
	return s
}

func (s *CopyObjectsResultSuccess) Validate() error {
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
