// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLifecycleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutBucketLifecycleHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAllowSameActionOverlap(v string) *PutBucketLifecycleHeaders
	GetXOssAllowSameActionOverlap() *string
}

type PutBucketLifecycleHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// Specifies whether to allow overlapped prefixes. Valid values:
	//
	// true: Overlapped prefixes are allowed.
	//
	// false: Overlapped prefixes are not allowed.
	//
	// example:
	//
	// true
	XOssAllowSameActionOverlap *string `json:"x-oss-allow-same-action-overlap,omitempty" xml:"x-oss-allow-same-action-overlap,omitempty"`
}

func (s PutBucketLifecycleHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLifecycleHeaders) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutBucketLifecycleHeaders) GetXOssAllowSameActionOverlap() *string {
	return s.XOssAllowSameActionOverlap
}

func (s *PutBucketLifecycleHeaders) SetCommonHeaders(v map[string]*string) *PutBucketLifecycleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketLifecycleHeaders) SetXOssAllowSameActionOverlap(v string) *PutBucketLifecycleHeaders {
	s.XOssAllowSameActionOverlap = &v
	return s
}

func (s *PutBucketLifecycleHeaders) Validate() error {
	return dara.Validate(s)
}
