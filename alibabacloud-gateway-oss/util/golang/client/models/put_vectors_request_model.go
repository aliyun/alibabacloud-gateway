// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *PutVectorsRequest
	GetIndexName() *string
	SetVectors(v []*Vector) *PutVectorsRequest
	GetVectors() []*Vector
}

type PutVectorsRequest struct {
	IndexName *string   `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Vectors   []*Vector `json:"vectors,omitempty" xml:"vectors,omitempty" type:"Repeated"`
}

func (s PutVectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutVectorsRequest) GoString() string {
	return s.String()
}

func (s *PutVectorsRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *PutVectorsRequest) GetVectors() []*Vector {
	return s.Vectors
}

func (s *PutVectorsRequest) SetIndexName(v string) *PutVectorsRequest {
	s.IndexName = &v
	return s
}

func (s *PutVectorsRequest) SetVectors(v []*Vector) *PutVectorsRequest {
	s.Vectors = v
	return s
}

func (s *PutVectorsRequest) Validate() error {
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
