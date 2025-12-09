// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputSerialization(v *InputSerialization) *SelectMetaRequest
	GetInputSerialization() *InputSerialization
	SetOverwriteIfExists(v bool) *SelectMetaRequest
	GetOverwriteIfExists() *bool
}

type SelectMetaRequest struct {
	InputSerialization *InputSerialization `json:"InputSerialization,omitempty" xml:"InputSerialization,omitempty"`
	OverwriteIfExists  *bool               `json:"OverwriteIfExists,omitempty" xml:"OverwriteIfExists,omitempty"`
}

func (s SelectMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectMetaRequest) GoString() string {
	return s.String()
}

func (s *SelectMetaRequest) GetInputSerialization() *InputSerialization {
	return s.InputSerialization
}

func (s *SelectMetaRequest) GetOverwriteIfExists() *bool {
	return s.OverwriteIfExists
}

func (s *SelectMetaRequest) SetInputSerialization(v *InputSerialization) *SelectMetaRequest {
	s.InputSerialization = v
	return s
}

func (s *SelectMetaRequest) SetOverwriteIfExists(v bool) *SelectMetaRequest {
	s.OverwriteIfExists = &v
	return s
}

func (s *SelectMetaRequest) Validate() error {
	if s.InputSerialization != nil {
		if err := s.InputSerialization.Validate(); err != nil {
			return err
		}
	}
	return nil
}
