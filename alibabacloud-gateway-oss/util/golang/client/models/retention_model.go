// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetention interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *Retention
	GetMode() *string
	SetRetainUtilDateInMs(v string) *Retention
	GetRetainUtilDateInMs() *string
}

type Retention struct {
	Mode               *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RetainUtilDateInMs *string `json:"RetainUtilDateInMs,omitempty" xml:"RetainUtilDateInMs,omitempty"`
}

func (s Retention) String() string {
	return dara.Prettify(s)
}

func (s Retention) GoString() string {
	return s.String()
}

func (s *Retention) GetMode() *string {
	return s.Mode
}

func (s *Retention) GetRetainUtilDateInMs() *string {
	return s.RetainUtilDateInMs
}

func (s *Retention) SetMode(v string) *Retention {
	s.Mode = &v
	return s
}

func (s *Retention) SetRetainUtilDateInMs(v string) *Retention {
	s.RetainUtilDateInMs = &v
	return s
}

func (s *Retention) Validate() error {
	return dara.Validate(s)
}
