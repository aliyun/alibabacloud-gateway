// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryResp interface {
	dara.Model
	String() string
	GoString() string
	SetAggregations(v *MetaQueryRespAggregations) *MetaQueryResp
	GetAggregations() *MetaQueryRespAggregations
	SetFiles(v *MetaQueryRespFiles) *MetaQueryResp
	GetFiles() *MetaQueryRespFiles
	SetNextToken(v string) *MetaQueryResp
	GetNextToken() *string
}

type MetaQueryResp struct {
	Aggregations *MetaQueryRespAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Struct"`
	Files        *MetaQueryRespFiles        `json:"Files,omitempty" xml:"Files,omitempty" type:"Struct"`
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s MetaQueryResp) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryResp) GoString() string {
	return s.String()
}

func (s *MetaQueryResp) GetAggregations() *MetaQueryRespAggregations {
	return s.Aggregations
}

func (s *MetaQueryResp) GetFiles() *MetaQueryRespFiles {
	return s.Files
}

func (s *MetaQueryResp) GetNextToken() *string {
	return s.NextToken
}

func (s *MetaQueryResp) SetAggregations(v *MetaQueryRespAggregations) *MetaQueryResp {
	s.Aggregations = v
	return s
}

func (s *MetaQueryResp) SetFiles(v *MetaQueryRespFiles) *MetaQueryResp {
	s.Files = v
	return s
}

func (s *MetaQueryResp) SetNextToken(v string) *MetaQueryResp {
	s.NextToken = &v
	return s
}

func (s *MetaQueryResp) Validate() error {
	if s.Aggregations != nil {
		if err := s.Aggregations.Validate(); err != nil {
			return err
		}
	}
	if s.Files != nil {
		if err := s.Files.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryRespAggregations struct {
	Aggregation []*MetaQueryAggregationsResult `json:"Aggregation,omitempty" xml:"Aggregation,omitempty" type:"Repeated"`
}

func (s MetaQueryRespAggregations) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryRespAggregations) GoString() string {
	return s.String()
}

func (s *MetaQueryRespAggregations) GetAggregation() []*MetaQueryAggregationsResult {
	return s.Aggregation
}

func (s *MetaQueryRespAggregations) SetAggregation(v []*MetaQueryAggregationsResult) *MetaQueryRespAggregations {
	s.Aggregation = v
	return s
}

func (s *MetaQueryRespAggregations) Validate() error {
	if s.Aggregation != nil {
		for _, item := range s.Aggregation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetaQueryRespFiles struct {
	File []*MetaQueryFile `json:"File,omitempty" xml:"File,omitempty" type:"Repeated"`
}

func (s MetaQueryRespFiles) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryRespFiles) GoString() string {
	return s.String()
}

func (s *MetaQueryRespFiles) GetFile() []*MetaQueryFile {
	return s.File
}

func (s *MetaQueryRespFiles) SetFile(v []*MetaQueryFile) *MetaQueryRespFiles {
	s.File = v
	return s
}

func (s *MetaQueryRespFiles) Validate() error {
	if s.File != nil {
		for _, item := range s.File {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
