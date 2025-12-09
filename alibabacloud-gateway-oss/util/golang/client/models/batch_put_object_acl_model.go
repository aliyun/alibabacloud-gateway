// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutObjectAcl interface {
	dara.Model
	String() string
	GoString() string
	SetObjectAcl(v string) *BatchPutObjectAcl
	GetObjectAcl() *string
}

type BatchPutObjectAcl struct {
	ObjectAcl *string `json:"ObjectAcl,omitempty" xml:"ObjectAcl,omitempty"`
}

func (s BatchPutObjectAcl) String() string {
	return dara.Prettify(s)
}

func (s BatchPutObjectAcl) GoString() string {
	return s.String()
}

func (s *BatchPutObjectAcl) GetObjectAcl() *string {
	return s.ObjectAcl
}

func (s *BatchPutObjectAcl) SetObjectAcl(v string) *BatchPutObjectAcl {
	s.ObjectAcl = &v
	return s
}

func (s *BatchPutObjectAcl) Validate() error {
	return dara.Validate(s)
}
