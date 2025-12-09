// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetVectors(v []*Vector) *GetVectorsResponseBody
	GetVectors() []*Vector
}

type GetVectorsResponseBody struct {
	Vectors []*Vector `json:"vectors,omitempty" xml:"vectors,omitempty" type:"Repeated"`
}

func (s GetVectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVectorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetVectorsResponseBody) GetVectors() []*Vector {
	return s.Vectors
}

func (s *GetVectorsResponseBody) SetVectors(v []*Vector) *GetVectorsResponseBody {
	s.Vectors = v
	return s
}

func (s *GetVectorsResponseBody) Validate() error {
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
