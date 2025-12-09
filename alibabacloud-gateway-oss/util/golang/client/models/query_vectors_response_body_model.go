// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVectors(v []*QueriedVector) *QueryVectorsResponseBody
	GetVectors() []*QueriedVector
}

type QueryVectorsResponseBody struct {
	Vectors []*QueriedVector `json:"vectors,omitempty" xml:"vectors,omitempty" type:"Repeated"`
}

func (s QueryVectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVectorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVectorsResponseBody) GetVectors() []*QueriedVector {
	return s.Vectors
}

func (s *QueryVectorsResponseBody) SetVectors(v []*QueriedVector) *QueryVectorsResponseBody {
	s.Vectors = v
	return s
}

func (s *QueryVectorsResponseBody) Validate() error {
	if s.Vectors != nil {
		for _, item := range s.Vectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
