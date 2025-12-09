// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInventoryEncryption interface {
	dara.Model
	String() string
	GoString() string
	SetSSEKMS(v *SSEKMS) *InventoryEncryption
	GetSSEKMS() *SSEKMS
	SetSSEOSS(v string) *InventoryEncryption
	GetSSEOSS() *string
}

type InventoryEncryption struct {
	SSEKMS *SSEKMS `json:"SSE-KMS,omitempty" xml:"SSE-KMS,omitempty"`
	SSEOSS *string `json:"SSE-OSS,omitempty" xml:"SSE-OSS,omitempty"`
}

func (s InventoryEncryption) String() string {
	return dara.Prettify(s)
}

func (s InventoryEncryption) GoString() string {
	return s.String()
}

func (s *InventoryEncryption) GetSSEKMS() *SSEKMS {
	return s.SSEKMS
}

func (s *InventoryEncryption) GetSSEOSS() *string {
	return s.SSEOSS
}

func (s *InventoryEncryption) SetSSEKMS(v *SSEKMS) *InventoryEncryption {
	s.SSEKMS = v
	return s
}

func (s *InventoryEncryption) SetSSEOSS(v string) *InventoryEncryption {
	s.SSEOSS = &v
	return s
}

func (s *InventoryEncryption) Validate() error {
	if s.SSEKMS != nil {
		if err := s.SSEKMS.Validate(); err != nil {
			return err
		}
	}
	return nil
}
