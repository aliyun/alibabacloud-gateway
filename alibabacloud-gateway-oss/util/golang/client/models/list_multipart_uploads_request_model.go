// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultipartUploadsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelimiter(v string) *ListMultipartUploadsRequest
	GetDelimiter() *string
	SetEncodingType(v string) *ListMultipartUploadsRequest
	GetEncodingType() *string
	SetKeyMarker(v string) *ListMultipartUploadsRequest
	GetKeyMarker() *string
	SetMaxUploads(v int64) *ListMultipartUploadsRequest
	GetMaxUploads() *int64
	SetPrefix(v string) *ListMultipartUploadsRequest
	GetPrefix() *string
	SetUploadIdMarker(v string) *ListMultipartUploadsRequest
	GetUploadIdMarker() *string
}

type ListMultipartUploadsRequest struct {
	// The character used to group objects by name. Objects whose names contain the same string that ranges from the specified prefix to the delimiter that appears for the first time are grouped as a CommonPrefixes element.
	//
	// example:
	//
	// /
	Delimiter *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	// The encoding type of the object name in the response. Values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key can be encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters such as characters with an American Standard Code for Information Interchange (ASCII) value from 0 to 10. You can set the encoding-type parameter to encode values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key in the response.
	//
	// Default value: null
	//
	// example:
	//
	// The upload ID of the multipart upload task after which the list operation starts. This parameter is used together with the key-marker parameter.
	//
	//   - If the key-marker parameter is not specified, OSS ignores the upload-id-marker parameter.
	//
	//   - If the key-marker parameter is specified, the query result includes the following tasks:
	//
	//     - Multipart upload tasks in which object names are alphabetically after the value of the key-marker parameter.
	//
	//     - Multipart upload tasks in which object names are the same as the value of the key-marker parameter but whose upload IDs are greater than the value of the upload-id-marker parameter.
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	// This parameter is used together with the upload-id-marker parameter to specify the position from which the next list begins.
	//
	// - If the upload-id-marker parameter is not set, Object Storage Service (OSS) returns all multipart upload tasks in which object names are alphabetically after the key-marker value.
	//
	// - If the upload-id-marker parameter is set, the response includes the following tasks:
	//
	//   - Multipart upload tasks in which object names are alphabetically after the key-marker value in alphabetical order
	//
	//   - Multipart upload tasks in which object names are the same as the key-marker parameter value but whose upload IDs are greater than the upload-id-marker parameter value
	//
	// example:
	//
	// test1.avi
	KeyMarker *string `json:"key-marker,omitempty" xml:"key-marker,omitempty"`
	// The maximumnumber of multipart upload tasks that can be returned for the current request. Default value: 1000. Maximum value: 1000.
	//
	// example:
	//
	// 1000
	MaxUploads *int64 `json:"max-uploads,omitempty" xml:"max-uploads,omitempty"`
	// The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.
	//
	// >You can use prefixes to group and manage objects in buckets in the same way you manage a folder in a file system.
	//
	// example:
	//
	// fun/
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The upload ID of the multipart upload task after which the list begins. This parameter is used together with the key-marker parameter.
	//
	// - If the key-marker parameter is not set, OSS ignores the upload-id-marker parameter.
	//
	// - If the key-marker parameter is configured, the query result includes:
	//
	//   - Multipart upload tasks in which object names are alphabetically after the key-marker value in alphabetical order
	//
	//   - Multipart upload tasks in which object names are the same as the key-marker parameter value but whose upload IDs are greater than the upload-id-marker parameter value
	//
	// example:
	//
	// 0004B99B8E707874FC2D692FA5D7****
	UploadIdMarker *string `json:"upload-id-marker,omitempty" xml:"upload-id-marker,omitempty"`
}

func (s ListMultipartUploadsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultipartUploadsRequest) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsRequest) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListMultipartUploadsRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListMultipartUploadsRequest) GetKeyMarker() *string {
	return s.KeyMarker
}

func (s *ListMultipartUploadsRequest) GetMaxUploads() *int64 {
	return s.MaxUploads
}

func (s *ListMultipartUploadsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListMultipartUploadsRequest) GetUploadIdMarker() *string {
	return s.UploadIdMarker
}

func (s *ListMultipartUploadsRequest) SetDelimiter(v string) *ListMultipartUploadsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetEncodingType(v string) *ListMultipartUploadsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetKeyMarker(v string) *ListMultipartUploadsRequest {
	s.KeyMarker = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetMaxUploads(v int64) *ListMultipartUploadsRequest {
	s.MaxUploads = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetPrefix(v string) *ListMultipartUploadsRequest {
	s.Prefix = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetUploadIdMarker(v string) *ListMultipartUploadsRequest {
	s.UploadIdMarker = &v
	return s
}

func (s *ListMultipartUploadsRequest) Validate() error {
	return dara.Validate(s)
}
