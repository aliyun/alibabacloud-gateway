// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipleObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelete(v *Delete) *DeleteMultipleObjectsRequest
	GetDelete() *Delete
	SetEncodingType(v string) *DeleteMultipleObjectsRequest
	GetEncodingType() *string
}

type DeleteMultipleObjectsRequest struct {
	Delete *Delete `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// The encoding type of the object name in the response. The value of the Key parameter is UTF-8 encoded. If the Key parameter includes control characters that are not supported by the XML 1.0 standard, you can specify this header to encode the value of the Key parameter in the response.
	//
	// example:
	//
	// url
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
}

func (s DeleteMultipleObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipleObjectsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsRequest) GetDelete() *Delete {
	return s.Delete
}

func (s *DeleteMultipleObjectsRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *DeleteMultipleObjectsRequest) SetDelete(v *Delete) *DeleteMultipleObjectsRequest {
	s.Delete = v
	return s
}

func (s *DeleteMultipleObjectsRequest) SetEncodingType(v string) *DeleteMultipleObjectsRequest {
	s.EncodingType = &v
	return s
}

func (s *DeleteMultipleObjectsRequest) Validate() error {
	if s.Delete != nil {
		if err := s.Delete.Validate(); err != nil {
			return err
		}
	}
	return nil
}
