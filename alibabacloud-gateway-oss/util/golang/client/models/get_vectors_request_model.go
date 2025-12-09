// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *GetVectorsRequest
	GetIndexName() *string
	SetKeys(v []*string) *GetVectorsRequest
	GetKeys() []*string
	SetReturnData(v bool) *GetVectorsRequest
	GetReturnData() *bool
	SetReturnMetadata(v bool) *GetVectorsRequest
	GetReturnMetadata() *bool
}

type GetVectorsRequest struct {
	IndexName      *string   `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Keys           []*string `json:"keys,omitempty" xml:"keys,omitempty" type:"Repeated"`
	ReturnData     *bool     `json:"returnData,omitempty" xml:"returnData,omitempty"`
	ReturnMetadata *bool     `json:"returnMetadata,omitempty" xml:"returnMetadata,omitempty"`
}

func (s GetVectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVectorsRequest) GoString() string {
	return s.String()
}

func (s *GetVectorsRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *GetVectorsRequest) GetKeys() []*string {
	return s.Keys
}

func (s *GetVectorsRequest) GetReturnData() *bool {
	return s.ReturnData
}

func (s *GetVectorsRequest) GetReturnMetadata() *bool {
	return s.ReturnMetadata
}

func (s *GetVectorsRequest) SetIndexName(v string) *GetVectorsRequest {
	s.IndexName = &v
	return s
}

func (s *GetVectorsRequest) SetKeys(v []*string) *GetVectorsRequest {
	s.Keys = v
	return s
}

func (s *GetVectorsRequest) SetReturnData(v bool) *GetVectorsRequest {
	s.ReturnData = &v
	return s
}

func (s *GetVectorsRequest) SetReturnMetadata(v bool) *GetVectorsRequest {
	s.ReturnMetadata = &v
	return s
}

func (s *GetVectorsRequest) Validate() error {
	return dara.Validate(s)
}
