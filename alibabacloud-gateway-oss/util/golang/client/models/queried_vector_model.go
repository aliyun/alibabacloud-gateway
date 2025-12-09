// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueriedVector interface {
	dara.Model
	String() string
	GoString() string
	SetDistance(v int64) *QueriedVector
	GetDistance() *int64
	SetKey(v string) *QueriedVector
	GetKey() *string
	SetMetadata(v interface{}) *QueriedVector
	GetMetadata() interface{}
}

type QueriedVector struct {
	// example:
	//
	// 32
	Distance *int64 `json:"distance,omitempty" xml:"distance,omitempty"`
	// example:
	//
	// test-vetor
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// {"key1":"value"}
	Metadata interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s QueriedVector) String() string {
	return dara.Prettify(s)
}

func (s QueriedVector) GoString() string {
	return s.String()
}

func (s *QueriedVector) GetDistance() *int64 {
	return s.Distance
}

func (s *QueriedVector) GetKey() *string {
	return s.Key
}

func (s *QueriedVector) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueriedVector) SetDistance(v int64) *QueriedVector {
	s.Distance = &v
	return s
}

func (s *QueriedVector) SetKey(v string) *QueriedVector {
	s.Key = &v
	return s
}

func (s *QueriedVector) SetMetadata(v interface{}) *QueriedVector {
	s.Metadata = v
	return s
}

func (s *QueriedVector) Validate() error {
	return dara.Validate(s)
}
