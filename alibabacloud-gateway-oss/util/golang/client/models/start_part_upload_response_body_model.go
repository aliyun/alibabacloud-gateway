// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPartUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetStartPartUploadResult(v *StartPartUploadResult) *StartPartUploadResponseBody
	GetStartPartUploadResult() *StartPartUploadResult
}

type StartPartUploadResponseBody struct {
	StartPartUploadResult *StartPartUploadResult `json:"StartPartUploadResult,omitempty" xml:"StartPartUploadResult,omitempty"`
}

func (s StartPartUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *StartPartUploadResponseBody) GetStartPartUploadResult() *StartPartUploadResult {
	return s.StartPartUploadResult
}

func (s *StartPartUploadResponseBody) SetStartPartUploadResult(v *StartPartUploadResult) *StartPartUploadResponseBody {
	s.StartPartUploadResult = v
	return s
}

func (s *StartPartUploadResponseBody) Validate() error {
	if s.StartPartUploadResult != nil {
		if err := s.StartPartUploadResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
