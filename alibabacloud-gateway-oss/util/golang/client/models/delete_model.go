// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelete interface {
	dara.Model
	String() string
	GoString() string
	SetObjects(v []*ObjectIdentifier) *Delete
	GetObjects() []*ObjectIdentifier
	SetQuiet(v bool) *Delete
	GetQuiet() *bool
}

type Delete struct {
	Objects []*ObjectIdentifier `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	Quiet   *bool               `json:"Quiet,omitempty" xml:"Quiet,omitempty"`
}

func (s Delete) String() string {
	return dara.Prettify(s)
}

func (s Delete) GoString() string {
	return s.String()
}

func (s *Delete) GetObjects() []*ObjectIdentifier {
	return s.Objects
}

func (s *Delete) GetQuiet() *bool {
	return s.Quiet
}

func (s *Delete) SetObjects(v []*ObjectIdentifier) *Delete {
	s.Objects = v
	return s
}

func (s *Delete) SetQuiet(v bool) *Delete {
	s.Quiet = &v
	return s
}

func (s *Delete) Validate() error {
	if s.Objects != nil {
		for _, item := range s.Objects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
