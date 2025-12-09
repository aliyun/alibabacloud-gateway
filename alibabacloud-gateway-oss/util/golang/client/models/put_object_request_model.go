// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iPutObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v io.Reader) *PutObjectRequest
	GetBody() io.Reader
}

type PutObjectRequest struct {
	// The body of the request.
	//
	// example:
	//
	// Binary content
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s PutObjectRequest) GoString() string {
	return s.String()
}

func (s *PutObjectRequest) GetBody() io.Reader {
	return s.Body
}

func (s *PutObjectRequest) SetBody(v io.Reader) *PutObjectRequest {
	s.Body = v
	return s
}

func (s *PutObjectRequest) Validate() error {
	return dara.Validate(s)
}
