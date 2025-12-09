// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteMultipartUpload interface {
	dara.Model
	String() string
	GoString() string
	SetPart(v []*CompleteMultipartUploadPart) *CompleteMultipartUpload
	GetPart() []*CompleteMultipartUploadPart
}

type CompleteMultipartUpload struct {
	Part []*CompleteMultipartUploadPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s CompleteMultipartUpload) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUpload) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUpload) GetPart() []*CompleteMultipartUploadPart {
	return s.Part
}

func (s *CompleteMultipartUpload) SetPart(v []*CompleteMultipartUploadPart) *CompleteMultipartUpload {
	s.Part = v
	return s
}

func (s *CompleteMultipartUpload) Validate() error {
	if s.Part != nil {
		for _, item := range s.Part {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CompleteMultipartUploadPart struct {
	ETag       *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	PartNumber *int64  `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
}

func (s CompleteMultipartUploadPart) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUploadPart) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadPart) GetETag() *string {
	return s.ETag
}

func (s *CompleteMultipartUploadPart) GetPartNumber() *int64 {
	return s.PartNumber
}

func (s *CompleteMultipartUploadPart) SetETag(v string) *CompleteMultipartUploadPart {
	s.ETag = &v
	return s
}

func (s *CompleteMultipartUploadPart) SetPartNumber(v int64) *CompleteMultipartUploadPart {
	s.PartNumber = &v
	return s
}

func (s *CompleteMultipartUploadPart) Validate() error {
	return dara.Validate(s)
}
