// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketReplicationBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketReplicationBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketReplicationBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *PutBucketReplicationBandwidthResponseBody) *PutBucketReplicationBandwidthResponse
	GetBody() *PutBucketReplicationBandwidthResponseBody
}

type PutBucketReplicationBandwidthResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutBucketReplicationBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutBucketReplicationBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketReplicationBandwidthResponse) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketReplicationBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketReplicationBandwidthResponse) GetBody() *PutBucketReplicationBandwidthResponseBody {
	return s.Body
}

func (s *PutBucketReplicationBandwidthResponse) SetHeaders(v map[string]*string) *PutBucketReplicationBandwidthResponse {
	s.Headers = v
	return s
}

func (s *PutBucketReplicationBandwidthResponse) SetStatusCode(v int32) *PutBucketReplicationBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketReplicationBandwidthResponse) SetBody(v *PutBucketReplicationBandwidthResponseBody) *PutBucketReplicationBandwidthResponse {
	s.Body = v
	return s
}

func (s *PutBucketReplicationBandwidthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
