// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVector interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *VectorData) *Vector
	GetData() *VectorData
	SetKey(v string) *Vector
	GetKey() *string
	SetMetadata(v interface{}) *Vector
	GetMetadata() interface{}
}

type Vector struct {
	Data *VectorData `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// test-key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// {"key1":"value1"}
	Metadata interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s Vector) String() string {
	return dara.Prettify(s)
}

func (s Vector) GoString() string {
	return s.String()
}

func (s *Vector) GetData() *VectorData {
	return s.Data
}

func (s *Vector) GetKey() *string {
	return s.Key
}

func (s *Vector) GetMetadata() interface{} {
	return s.Metadata
}

func (s *Vector) SetData(v *VectorData) *Vector {
	s.Data = v
	return s
}

func (s *Vector) SetKey(v string) *Vector {
	s.Key = &v
	return s
}

func (s *Vector) SetMetadata(v interface{}) *Vector {
	s.Metadata = v
	return s
}

func (s *Vector) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
