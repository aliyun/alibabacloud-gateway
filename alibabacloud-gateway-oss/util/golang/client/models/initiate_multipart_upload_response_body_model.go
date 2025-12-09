// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateMultipartUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInitiateMultipartUploadResult(v *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) *InitiateMultipartUploadResponseBody
	GetInitiateMultipartUploadResult() *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult
}

type InitiateMultipartUploadResponseBody struct {
	// The container that stores the results of the InitiateMultipartUpload request.
	InitiateMultipartUploadResult *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult `json:"InitiateMultipartUploadResult,omitempty" xml:"InitiateMultipartUploadResult,omitempty" type:"Struct"`
}

func (s InitiateMultipartUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitiateMultipartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadResponseBody) GetInitiateMultipartUploadResult() *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult {
	return s.InitiateMultipartUploadResult
}

func (s *InitiateMultipartUploadResponseBody) SetInitiateMultipartUploadResult(v *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) *InitiateMultipartUploadResponseBody {
	s.InitiateMultipartUploadResult = v
	return s
}

func (s *InitiateMultipartUploadResponseBody) Validate() error {
	if s.InitiateMultipartUploadResult != nil {
		if err := s.InitiateMultipartUploadResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult struct {
	// The name of the bucket to which the object is uploaded by the multipart upload task.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The encoding type of the object name in the response. If the encoding-type parameter is specified in the request, the object name in the response is encoded.
	//
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// The name of the object that is uploaded by the multipart upload task.
	//
	// example:
	//
	// test.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The upload ID that uniquely identifies the multipart upload task. The upload ID is used to call UploadPart and CompleteMultipartUpload later.
	//
	// example:
	//
	// 0004B9894A22E5B1888A1E29F823****
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) String() string {
	return dara.Prettify(s)
}

func (s InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) GetBucket() *string {
	return s.Bucket
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) GetKey() *string {
	return s.Key
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) GetUploadId() *string {
	return s.UploadId
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) SetBucket(v string) *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult {
	s.Bucket = &v
	return s
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) SetEncodingType(v string) *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult {
	s.EncodingType = &v
	return s
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) SetKey(v string) *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult {
	s.Key = &v
	return s
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) SetUploadId(v string) *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult {
	s.UploadId = &v
	return s
}

func (s *InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult) Validate() error {
	return dara.Validate(s)
}
