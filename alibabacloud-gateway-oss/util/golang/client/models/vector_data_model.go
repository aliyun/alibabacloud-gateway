// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVectorData interface {
	dara.Model
	String() string
	GoString() string
	SetFloat32(v []*float32) *VectorData
	GetFloat32() []*float32
}

type VectorData struct {
	Float32 []*float32 `json:"float32,omitempty" xml:"float32,omitempty" type:"Repeated"`
}

func (s VectorData) String() string {
	return dara.Prettify(s)
}

func (s VectorData) GoString() string {
	return s.String()
}

func (s *VectorData) GetFloat32() []*float32 {
	return s.Float32
}

func (s *VectorData) SetFloat32(v []*float32) *VectorData {
	s.Float32 = v
	return s
}

func (s *VectorData) Validate() error {
	return dara.Validate(s)
}
