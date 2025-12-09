// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCopyObjectsResult(v *CopyObjectsResult) *CopyObjectsResponseBody
	GetCopyObjectsResult() *CopyObjectsResult
}

type CopyObjectsResponseBody struct {
	CopyObjectsResult *CopyObjectsResult `json:"CopyObjectsResult,omitempty" xml:"CopyObjectsResult,omitempty"`
}

func (s CopyObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *CopyObjectsResponseBody) GetCopyObjectsResult() *CopyObjectsResult {
	return s.CopyObjectsResult
}

func (s *CopyObjectsResponseBody) SetCopyObjectsResult(v *CopyObjectsResult) *CopyObjectsResponseBody {
	s.CopyObjectsResult = v
	return s
}

func (s *CopyObjectsResponseBody) Validate() error {
	if s.CopyObjectsResult != nil {
		if err := s.CopyObjectsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
