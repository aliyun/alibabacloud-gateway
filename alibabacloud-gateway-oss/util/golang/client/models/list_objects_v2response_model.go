// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListObjectsV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListObjectsV2Response
	GetStatusCode() *int32
	SetBody(v *ListObjectsV2ResponseBody) *ListObjectsV2Response
	GetBody() *ListObjectsV2ResponseBody
}

type ListObjectsV2Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListObjectsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListObjectsV2Response) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsV2Response) GoString() string {
	return s.String()
}

func (s *ListObjectsV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListObjectsV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListObjectsV2Response) GetBody() *ListObjectsV2ResponseBody {
	return s.Body
}

func (s *ListObjectsV2Response) SetHeaders(v map[string]*string) *ListObjectsV2Response {
	s.Headers = v
	return s
}

func (s *ListObjectsV2Response) SetStatusCode(v int32) *ListObjectsV2Response {
	s.StatusCode = &v
	return s
}

func (s *ListObjectsV2Response) SetBody(v *ListObjectsV2ResponseBody) *ListObjectsV2Response {
	s.Body = v
	return s
}

func (s *ListObjectsV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
