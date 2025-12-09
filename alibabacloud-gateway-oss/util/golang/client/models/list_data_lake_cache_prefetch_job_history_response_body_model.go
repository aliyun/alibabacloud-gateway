// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCachePrefetchJobHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListDataLakeCachePrefetchJobHistory(v *ListDataLakeCachePrefetchJobHistory) *ListDataLakeCachePrefetchJobHistoryResponseBody
	GetListDataLakeCachePrefetchJobHistory() *ListDataLakeCachePrefetchJobHistory
}

type ListDataLakeCachePrefetchJobHistoryResponseBody struct {
	ListDataLakeCachePrefetchJobHistory *ListDataLakeCachePrefetchJobHistory `json:"ListDataLakeCachePrefetchJobHistory,omitempty" xml:"ListDataLakeCachePrefetchJobHistory,omitempty"`
}

func (s ListDataLakeCachePrefetchJobHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobHistoryResponseBody) GetListDataLakeCachePrefetchJobHistory() *ListDataLakeCachePrefetchJobHistory {
	return s.ListDataLakeCachePrefetchJobHistory
}

func (s *ListDataLakeCachePrefetchJobHistoryResponseBody) SetListDataLakeCachePrefetchJobHistory(v *ListDataLakeCachePrefetchJobHistory) *ListDataLakeCachePrefetchJobHistoryResponseBody {
	s.ListDataLakeCachePrefetchJobHistory = v
	return s
}

func (s *ListDataLakeCachePrefetchJobHistoryResponseBody) Validate() error {
	if s.ListDataLakeCachePrefetchJobHistory != nil {
		if err := s.ListDataLakeCachePrefetchJobHistory.Validate(); err != nil {
			return err
		}
	}
	return nil
}
