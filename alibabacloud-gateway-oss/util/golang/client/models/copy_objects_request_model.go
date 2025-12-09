// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCopy(v *CopyObjectsCopy) *CopyObjectsRequest
	GetCopy() *CopyObjectsCopy
}

type CopyObjectsRequest struct {
	Copy *CopyObjectsCopy `json:"Copy,omitempty" xml:"Copy,omitempty"`
}

func (s CopyObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsRequest) GoString() string {
	return s.String()
}

func (s *CopyObjectsRequest) GetCopy() *CopyObjectsCopy {
	return s.Copy
}

func (s *CopyObjectsRequest) SetCopy(v *CopyObjectsCopy) *CopyObjectsRequest {
	s.Copy = v
	return s
}

func (s *CopyObjectsRequest) Validate() error {
	if s.Copy != nil {
		if err := s.Copy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
