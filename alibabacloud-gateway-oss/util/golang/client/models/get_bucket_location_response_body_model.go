// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLocationConstraint(v string) *GetBucketLocationResponseBody
	GetLocationConstraint() *string
}

type GetBucketLocationResponseBody struct {
	// The region in which the bucket resides.\\
	//
	// Examples: oss-cn-hangzhou, oss-cn-shanghai, oss-cn-qingdao, oss-cn-beijing, oss-cn-zhangjiakou, oss-cn-hongkong, oss-cn-shenzhen, oss-us-west-1, oss-us-east-1, and oss-ap-southeast-1.
	//
	// \\
	//
	// For more information about the regions in which buckets reside, see [Regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// example:
	//
	// oss-cn-hangzhou
	LocationConstraint *string `json:"LocationConstraint,omitempty" xml:"LocationConstraint,omitempty"`
}

func (s GetBucketLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLocationResponseBody) GetLocationConstraint() *string {
	return s.LocationConstraint
}

func (s *GetBucketLocationResponseBody) SetLocationConstraint(v string) *GetBucketLocationResponseBody {
	s.LocationConstraint = &v
	return s
}

func (s *GetBucketLocationResponseBody) Validate() error {
	return dara.Validate(s)
}
