// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectRetentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetention(v *Retention) *GetObjectRetentionResponseBody
	GetRetention() *Retention
}

type GetObjectRetentionResponseBody struct {
	Retention *Retention `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s GetObjectRetentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectRetentionResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectRetentionResponseBody) GetRetention() *Retention {
	return s.Retention
}

func (s *GetObjectRetentionResponseBody) SetRetention(v *Retention) *GetObjectRetentionResponseBody {
	s.Retention = v
	return s
}

func (s *GetObjectRetentionResponseBody) Validate() error {
	if s.Retention != nil {
		if err := s.Retention.Validate(); err != nil {
			return err
		}
	}
	return nil
}
