// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpload interface {
	dara.Model
	String() string
	GoString() string
	SetInitiated(v string) *Upload
	GetInitiated() *string
	SetKey(v string) *Upload
	GetKey() *string
	SetUploadId(v string) *Upload
	GetUploadId() *string
}

type Upload struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	Initiated *string `json:"Initiated,omitempty" xml:"Initiated,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	UploadId  *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s Upload) String() string {
	return dara.Prettify(s)
}

func (s Upload) GoString() string {
	return s.String()
}

func (s *Upload) GetInitiated() *string {
	return s.Initiated
}

func (s *Upload) GetKey() *string {
	return s.Key
}

func (s *Upload) GetUploadId() *string {
	return s.UploadId
}

func (s *Upload) SetInitiated(v string) *Upload {
	s.Initiated = &v
	return s
}

func (s *Upload) SetKey(v string) *Upload {
	s.Key = &v
	return s
}

func (s *Upload) SetUploadId(v string) *Upload {
	s.UploadId = &v
	return s
}

func (s *Upload) Validate() error {
	return dara.Validate(s)
}
