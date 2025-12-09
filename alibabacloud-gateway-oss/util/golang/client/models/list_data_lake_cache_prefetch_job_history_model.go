// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCachePrefetchJobHistory interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeCachePrefetchJobHistory(v []*DataLakeCachePrefetchJobHistory) *ListDataLakeCachePrefetchJobHistory
	GetDataLakeCachePrefetchJobHistory() []*DataLakeCachePrefetchJobHistory
}

type ListDataLakeCachePrefetchJobHistory struct {
	DataLakeCachePrefetchJobHistory []*DataLakeCachePrefetchJobHistory `json:"DataLakeCachePrefetchJobHistory,omitempty" xml:"DataLakeCachePrefetchJobHistory,omitempty" type:"Repeated"`
}

func (s ListDataLakeCachePrefetchJobHistory) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobHistory) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobHistory) GetDataLakeCachePrefetchJobHistory() []*DataLakeCachePrefetchJobHistory {
	return s.DataLakeCachePrefetchJobHistory
}

func (s *ListDataLakeCachePrefetchJobHistory) SetDataLakeCachePrefetchJobHistory(v []*DataLakeCachePrefetchJobHistory) *ListDataLakeCachePrefetchJobHistory {
	s.DataLakeCachePrefetchJobHistory = v
	return s
}

func (s *ListDataLakeCachePrefetchJobHistory) Validate() error {
	if s.DataLakeCachePrefetchJobHistory != nil {
		for _, item := range s.DataLakeCachePrefetchJobHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
