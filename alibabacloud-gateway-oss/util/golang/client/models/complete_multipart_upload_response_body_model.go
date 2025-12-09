// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteMultipartUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompleteMultipartUploadResult(v *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) *CompleteMultipartUploadResponseBody
	GetCompleteMultipartUploadResult() *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult
}

type CompleteMultipartUploadResponseBody struct {
	// The container that stores the results of the CompleteMultipartUpload request.
	CompleteMultipartUploadResult *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult `json:"CompleteMultipartUploadResult,omitempty" xml:"CompleteMultipartUploadResult,omitempty" type:"Struct"`
}

func (s CompleteMultipartUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadResponseBody) GetCompleteMultipartUploadResult() *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult {
	return s.CompleteMultipartUploadResult
}

func (s *CompleteMultipartUploadResponseBody) SetCompleteMultipartUploadResult(v *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) *CompleteMultipartUploadResponseBody {
	s.CompleteMultipartUploadResult = v
	return s
}

func (s *CompleteMultipartUploadResponseBody) Validate() error {
	if s.CompleteMultipartUploadResult != nil {
		if err := s.CompleteMultipartUploadResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult struct {
	// The name of the bucket that contains the object you want to restore.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ETag that is generated when an object is created. ETags are used to identify the content of objects.
	//
	// If an object is created by calling the CompleteMultipartUpload operation, the ETag value is not the MD5 hash of the object content but a unique value calculated based on a specific rule.
	//
	// > The ETag of an object can be used to check whether the object content is modified. However, we recommend that you use the MD5 hash of an object rather than the ETag value of the object to verify data integrity.
	//
	// example:
	//
	// "B864DB6A936D376F9F8D3ED3BBE540****"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The encoding type of the object name in the response. If this parameter is specified in the request, the object name is encoded in the response.
	//
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// The name of the uploaded object.
	//
	// example:
	//
	// multipart.data
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The URL that is used to access the uploaded object.
	//
	// example:
	//
	// http://oss-example.oss-cn-hangzhou.aliyuncs.com/multipart.data
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) GetBucket() *string {
	return s.Bucket
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) GetETag() *string {
	return s.ETag
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) GetKey() *string {
	return s.Key
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) GetLocation() *string {
	return s.Location
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) SetBucket(v string) *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult {
	s.Bucket = &v
	return s
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) SetETag(v string) *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult {
	s.ETag = &v
	return s
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) SetEncodingType(v string) *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult {
	s.EncodingType = &v
	return s
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) SetKey(v string) *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult {
	s.Key = &v
	return s
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) SetLocation(v string) *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult {
	s.Location = &v
	return s
}

func (s *CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult) Validate() error {
	return dara.Validate(s)
}
