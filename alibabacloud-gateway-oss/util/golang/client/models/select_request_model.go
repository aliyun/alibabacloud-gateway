// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *SelectRequest
	GetExpression() *string
	SetInputSerialization(v *InputSerialization) *SelectRequest
	GetInputSerialization() *InputSerialization
	SetOptions(v *SelectRequestOptions) *SelectRequest
	GetOptions() *SelectRequestOptions
	SetOutputSerialization(v *OutputSerialization) *SelectRequest
	GetOutputSerialization() *OutputSerialization
}

type SelectRequest struct {
	Expression          *string               `json:"Expression,omitempty" xml:"Expression,omitempty"`
	InputSerialization  *InputSerialization   `json:"InputSerialization,omitempty" xml:"InputSerialization,omitempty"`
	Options             *SelectRequestOptions `json:"Options,omitempty" xml:"Options,omitempty"`
	OutputSerialization *OutputSerialization  `json:"OutputSerialization,omitempty" xml:"OutputSerialization,omitempty"`
}

func (s SelectRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectRequest) GoString() string {
	return s.String()
}

func (s *SelectRequest) GetExpression() *string {
	return s.Expression
}

func (s *SelectRequest) GetInputSerialization() *InputSerialization {
	return s.InputSerialization
}

func (s *SelectRequest) GetOptions() *SelectRequestOptions {
	return s.Options
}

func (s *SelectRequest) GetOutputSerialization() *OutputSerialization {
	return s.OutputSerialization
}

func (s *SelectRequest) SetExpression(v string) *SelectRequest {
	s.Expression = &v
	return s
}

func (s *SelectRequest) SetInputSerialization(v *InputSerialization) *SelectRequest {
	s.InputSerialization = v
	return s
}

func (s *SelectRequest) SetOptions(v *SelectRequestOptions) *SelectRequest {
	s.Options = v
	return s
}

func (s *SelectRequest) SetOutputSerialization(v *OutputSerialization) *SelectRequest {
	s.OutputSerialization = v
	return s
}

func (s *SelectRequest) Validate() error {
	if s.InputSerialization != nil {
		if err := s.InputSerialization.Validate(); err != nil {
			return err
		}
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	if s.OutputSerialization != nil {
		if err := s.OutputSerialization.Validate(); err != nil {
			return err
		}
	}
	return nil
}
