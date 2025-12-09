// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketResourceGroupResponseBody) *GetBucketResourceGroupResponse
	GetBody() *GetBucketResourceGroupResponseBody
}

type GetBucketResourceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetBucketResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketResourceGroupResponse) GetBody() *GetBucketResourceGroupResponseBody {
	return s.Body
}

func (s *GetBucketResourceGroupResponse) SetHeaders(v map[string]*string) *GetBucketResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetBucketResourceGroupResponse) SetStatusCode(v int32) *GetBucketResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketResourceGroupResponse) SetBody(v *GetBucketResourceGroupResponseBody) *GetBucketResourceGroupResponse {
	s.Body = v
	return s
}

func (s *GetBucketResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
