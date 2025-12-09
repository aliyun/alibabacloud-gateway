// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataLakeCachePrefetchJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeCachePrefetchJobID(v *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) *PutDataLakeCachePrefetchJobResponseBody
	GetDataLakeCachePrefetchJobID() *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID
}

type PutDataLakeCachePrefetchJobResponseBody struct {
	DataLakeCachePrefetchJobID *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID `json:"DataLakeCachePrefetchJobID,omitempty" xml:"DataLakeCachePrefetchJobID,omitempty" type:"Struct"`
}

func (s PutDataLakeCachePrefetchJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeCachePrefetchJobResponseBody) GoString() string {
	return s.String()
}

func (s *PutDataLakeCachePrefetchJobResponseBody) GetDataLakeCachePrefetchJobID() *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID {
	return s.DataLakeCachePrefetchJobID
}

func (s *PutDataLakeCachePrefetchJobResponseBody) SetDataLakeCachePrefetchJobID(v *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) *PutDataLakeCachePrefetchJobResponseBody {
	s.DataLakeCachePrefetchJobID = v
	return s
}

func (s *PutDataLakeCachePrefetchJobResponseBody) Validate() error {
	if s.DataLakeCachePrefetchJobID != nil {
		if err := s.DataLakeCachePrefetchJobID.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID struct {
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) GoString() string {
	return s.String()
}

func (s *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) GetID() *string {
	return s.ID
}

func (s *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) SetID(v string) *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID {
	s.ID = &v
	return s
}

func (s *PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID) Validate() error {
	return dara.Validate(s)
}
