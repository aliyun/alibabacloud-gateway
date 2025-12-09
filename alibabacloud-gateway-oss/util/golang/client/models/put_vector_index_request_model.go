// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVectorIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *VectorIndexPut) *PutVectorIndexRequest
	GetBody() *VectorIndexPut
}

type PutVectorIndexRequest struct {
	Body *VectorIndexPut `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutVectorIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s PutVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *PutVectorIndexRequest) GetBody() *VectorIndexPut {
	return s.Body
}

func (s *PutVectorIndexRequest) SetBody(v *VectorIndexPut) *PutVectorIndexRequest {
	s.Body = v
	return s
}

func (s *PutVectorIndexRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
