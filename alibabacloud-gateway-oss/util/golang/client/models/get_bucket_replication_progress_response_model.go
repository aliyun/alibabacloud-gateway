// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketReplicationProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketReplicationProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketReplicationProgressResponseBody) *GetBucketReplicationProgressResponse
	GetBody() *GetBucketReplicationProgressResponseBody
}

type GetBucketReplicationProgressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketReplicationProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketReplicationProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationProgressResponse) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketReplicationProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketReplicationProgressResponse) GetBody() *GetBucketReplicationProgressResponseBody {
	return s.Body
}

func (s *GetBucketReplicationProgressResponse) SetHeaders(v map[string]*string) *GetBucketReplicationProgressResponse {
	s.Headers = v
	return s
}

func (s *GetBucketReplicationProgressResponse) SetStatusCode(v int32) *GetBucketReplicationProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketReplicationProgressResponse) SetBody(v *GetBucketReplicationProgressResponseBody) *GetBucketReplicationProgressResponse {
	s.Body = v
	return s
}

func (s *GetBucketReplicationProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
