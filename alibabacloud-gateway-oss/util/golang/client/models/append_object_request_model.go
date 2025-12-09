// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAppendObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v io.Reader) *AppendObjectRequest
	GetBody() io.Reader
	SetPosition(v int64) *AppendObjectRequest
	GetPosition() *int64
}

type AppendObjectRequest struct {
	// The request body.
	//
	// example:
	//
	// Binary content.
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	// The position from which the AppendObject operation starts.  Each time an AppendObject operation succeeds, the x-oss-next-append-position header is included in the response to specify the position from which the next AppendObject operation starts. The value of position in the first AppendObject operation performed on an object must be 0. The value of position in subsequent AppendObject operations performed on the object is the current length of the object. For example, if the value of position specified in the first AppendObject request is 0 and the value of content-length is 65536, the value of position in the second AppendObject request must be 65536.
	//
	//
	//
	// - If the value of position in the AppendObject request is 0 and the name of the object that you want to append is unique, you can set headers such as x-oss-server-side-encryption in an AppendObject request in the same way as you set in a PutObject request. If you add the x-oss-server-side-encryption header to an AppendObject request, the x-oss-server-side-encryption header is included in the response to the request. If you want to modify metadata, you can call the CopyObject operation.
	//
	// - If you call an AppendObject operation to append a 0 KB object whose position value is valid to an Appendable object, the status of the Appendable object is not changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
}

func (s AppendObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s AppendObjectRequest) GoString() string {
	return s.String()
}

func (s *AppendObjectRequest) GetBody() io.Reader {
	return s.Body
}

func (s *AppendObjectRequest) GetPosition() *int64 {
	return s.Position
}

func (s *AppendObjectRequest) SetBody(v io.Reader) *AppendObjectRequest {
	s.Body = v
	return s
}

func (s *AppendObjectRequest) SetPosition(v int64) *AppendObjectRequest {
	s.Position = &v
	return s
}

func (s *AppendObjectRequest) Validate() error {
	return dara.Validate(s)
}
