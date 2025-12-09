// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperation interface {
	dara.Model
	String() string
	GoString() string
	SetCopyObject(v *BatchCopyObject) *BatchOperation
	GetCopyObject() *BatchCopyObject
	SetDeleteObjectTagging(v interface{}) *BatchOperation
	GetDeleteObjectTagging() interface{}
	SetPutObjectAcl(v *BatchPutObjectAcl) *BatchOperation
	GetPutObjectAcl() *BatchPutObjectAcl
	SetPutObjectTagging(v *BatchPutObjectTagging) *BatchOperation
	GetPutObjectTagging() *BatchPutObjectTagging
	SetRestoreObject(v *BatchRestoreObject) *BatchOperation
	GetRestoreObject() *BatchRestoreObject
}

type BatchOperation struct {
	CopyObject          *BatchCopyObject       `json:"CopyObject,omitempty" xml:"CopyObject,omitempty"`
	DeleteObjectTagging interface{}            `json:"DeleteObjectTagging,omitempty" xml:"DeleteObjectTagging,omitempty"`
	PutObjectAcl        *BatchPutObjectAcl     `json:"PutObjectAcl,omitempty" xml:"PutObjectAcl,omitempty"`
	PutObjectTagging    *BatchPutObjectTagging `json:"PutObjectTagging,omitempty" xml:"PutObjectTagging,omitempty"`
	RestoreObject       *BatchRestoreObject    `json:"RestoreObject,omitempty" xml:"RestoreObject,omitempty"`
}

func (s BatchOperation) String() string {
	return dara.Prettify(s)
}

func (s BatchOperation) GoString() string {
	return s.String()
}

func (s *BatchOperation) GetCopyObject() *BatchCopyObject {
	return s.CopyObject
}

func (s *BatchOperation) GetDeleteObjectTagging() interface{} {
	return s.DeleteObjectTagging
}

func (s *BatchOperation) GetPutObjectAcl() *BatchPutObjectAcl {
	return s.PutObjectAcl
}

func (s *BatchOperation) GetPutObjectTagging() *BatchPutObjectTagging {
	return s.PutObjectTagging
}

func (s *BatchOperation) GetRestoreObject() *BatchRestoreObject {
	return s.RestoreObject
}

func (s *BatchOperation) SetCopyObject(v *BatchCopyObject) *BatchOperation {
	s.CopyObject = v
	return s
}

func (s *BatchOperation) SetDeleteObjectTagging(v interface{}) *BatchOperation {
	s.DeleteObjectTagging = v
	return s
}

func (s *BatchOperation) SetPutObjectAcl(v *BatchPutObjectAcl) *BatchOperation {
	s.PutObjectAcl = v
	return s
}

func (s *BatchOperation) SetPutObjectTagging(v *BatchPutObjectTagging) *BatchOperation {
	s.PutObjectTagging = v
	return s
}

func (s *BatchOperation) SetRestoreObject(v *BatchRestoreObject) *BatchOperation {
	s.RestoreObject = v
	return s
}

func (s *BatchOperation) Validate() error {
	if s.CopyObject != nil {
		if err := s.CopyObject.Validate(); err != nil {
			return err
		}
	}
	if s.PutObjectAcl != nil {
		if err := s.PutObjectAcl.Validate(); err != nil {
			return err
		}
	}
	if s.PutObjectTagging != nil {
		if err := s.PutObjectTagging.Validate(); err != nil {
			return err
		}
	}
	if s.RestoreObject != nil {
		if err := s.RestoreObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}
