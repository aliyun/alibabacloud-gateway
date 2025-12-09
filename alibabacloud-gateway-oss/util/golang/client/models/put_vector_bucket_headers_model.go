// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVectorBucketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutVectorBucketHeaders
	GetCommonHeaders() map[string]*string
	SetXOssResourceGroupId(v string) *PutVectorBucketHeaders
	GetXOssResourceGroupId() *string
}

type PutVectorBucketHeaders struct {
	CommonHeaders       map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XOssResourceGroupId *string            `json:"x-oss-resource-group-id,omitempty" xml:"x-oss-resource-group-id,omitempty"`
}

func (s PutVectorBucketHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutVectorBucketHeaders) GoString() string {
	return s.String()
}

func (s *PutVectorBucketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutVectorBucketHeaders) GetXOssResourceGroupId() *string {
	return s.XOssResourceGroupId
}

func (s *PutVectorBucketHeaders) SetCommonHeaders(v map[string]*string) *PutVectorBucketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutVectorBucketHeaders) SetXOssResourceGroupId(v string) *PutVectorBucketHeaders {
	s.XOssResourceGroupId = &v
	return s
}

func (s *PutVectorBucketHeaders) Validate() error {
	return dara.Validate(s)
}
