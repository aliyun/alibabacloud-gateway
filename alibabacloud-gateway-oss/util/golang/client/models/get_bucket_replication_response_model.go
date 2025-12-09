// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketReplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketReplicationResponseBody) *GetBucketReplicationResponse
	GetBody() *GetBucketReplicationResponseBody
}

type GetBucketReplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketReplicationResponse) GetBody() *GetBucketReplicationResponseBody {
	return s.Body
}

func (s *GetBucketReplicationResponse) SetHeaders(v map[string]*string) *GetBucketReplicationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketReplicationResponse) SetStatusCode(v int32) *GetBucketReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketReplicationResponse) SetBody(v *GetBucketReplicationResponseBody) *GetBucketReplicationResponse {
	s.Body = v
	return s
}

func (s *GetBucketReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
