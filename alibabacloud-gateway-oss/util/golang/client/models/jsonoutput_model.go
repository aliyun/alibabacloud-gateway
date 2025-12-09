// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJSONOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRecordDelimiter(v string) *JSONOutput
	GetRecordDelimiter() *string
}

type JSONOutput struct {
	RecordDelimiter *string `json:"RecordDelimiter,omitempty" xml:"RecordDelimiter,omitempty"`
}

func (s JSONOutput) String() string {
	return dara.Prettify(s)
}

func (s JSONOutput) GoString() string {
	return s.String()
}

func (s *JSONOutput) GetRecordDelimiter() *string {
	return s.RecordDelimiter
}

func (s *JSONOutput) SetRecordDelimiter(v string) *JSONOutput {
	s.RecordDelimiter = &v
	return s
}

func (s *JSONOutput) Validate() error {
	return dara.Validate(s)
}
