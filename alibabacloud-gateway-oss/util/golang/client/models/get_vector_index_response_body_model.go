// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIndex(v *VectorIndex) *GetVectorIndexResponseBody
	GetIndex() *VectorIndex
}

type GetVectorIndexResponseBody struct {
	Index *VectorIndex `json:"index,omitempty" xml:"index,omitempty"`
}

func (s GetVectorIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVectorIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetVectorIndexResponseBody) GetIndex() *VectorIndex {
	return s.Index
}

func (s *GetVectorIndexResponseBody) SetIndex(v *VectorIndex) *GetVectorIndexResponseBody {
	s.Index = v
	return s
}

func (s *GetVectorIndexResponseBody) Validate() error {
	if s.Index != nil {
		if err := s.Index.Validate(); err != nil {
			return err
		}
	}
	return nil
}
