// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelimiter(v string) *ListObjectVersionsRequest
	GetDelimiter() *string
	SetEncodingType(v string) *ListObjectVersionsRequest
	GetEncodingType() *string
	SetKeyMarker(v string) *ListObjectVersionsRequest
	GetKeyMarker() *string
	SetMaxKeys(v int64) *ListObjectVersionsRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListObjectVersionsRequest
	GetPrefix() *string
	SetVersionIdMarker(v string) *ListObjectVersionsRequest
	GetVersionIdMarker() *string
}

type ListObjectVersionsRequest struct {
	// The character that is used to group objects by name. If you specify prefix and delimiter in the request, the response contains CommonPrefixes. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes. If you specify prefix and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// /
	Delimiter *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	// The encoding type of the content in the response. By default, this parameter is left empty. Set the value to URL.
	//
	// >  The values of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the value of Delimiter, Marker, Prefix, NextMarker, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	// The name of the object after which the GetBucketVersions (ListObjectVersions) operation begins. If this parameter is specified, objects whose name is alphabetically after the value of key-marker are returned. Use key-marker and version-id-marker in combination. The value of key-marker must be less than 1,024 bytes in length.
	//
	// By default, this parameter is left empty.
	//
	// >  You must also specify key-marker if you specify version-id-marker.
	//
	// example:
	//
	// example
	KeyMarker *string `json:"key-marker,omitempty" xml:"key-marker,omitempty"`
	// The maximum number of objects to be returned. If the number of returned objects exceeds the value of max-keys, the response contains `NextKeyMarker` and `NextVersionIdMarker`. Specify the values of `NextKeyMarker` and `NextVersionIdMarker` as the markers for the next request. Valid values: 1 to 999. Default value: 100.
	//
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// The prefix that the names of returned objects must contain.
	//
	// 	- The value of prefix must be less than 1,024 bytes in length.
	//
	// 	- If you specify prefix, the names of the returned objects contain the prefix.
	//
	// If you set prefix to a directory name, the objects whose name starts with the prefix are listed. The returned objects consist of all objects and subdirectories in the directory.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// fun
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The version ID of the object specified in key-marker after which the GetBucketVersions (ListObjectVersions) operation begins. The versions are returned from the latest version to the earliest version. If version-id-marker is not specified, the GetBucketVersions (ListObjectVersions) operation starts from the latest version of the object whose name is alphabetically after the value of key-marker by default.
	//
	// By default, this parameter is left empty.
	//
	// Valid values: version IDs.
	//
	// example:
	//
	// CAEQMxiBgICbof2D0BYiIGRhZjgwMzJiMjA3MjQ0ODE5MWYxZDYwMzJlZjU1****
	VersionIdMarker *string `json:"version-id-marker,omitempty" xml:"version-id-marker,omitempty"`
}

func (s ListObjectVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListObjectVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsRequest) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListObjectVersionsRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectVersionsRequest) GetKeyMarker() *string {
	return s.KeyMarker
}

func (s *ListObjectVersionsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListObjectVersionsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectVersionsRequest) GetVersionIdMarker() *string {
	return s.VersionIdMarker
}

func (s *ListObjectVersionsRequest) SetDelimiter(v string) *ListObjectVersionsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListObjectVersionsRequest) SetEncodingType(v string) *ListObjectVersionsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListObjectVersionsRequest) SetKeyMarker(v string) *ListObjectVersionsRequest {
	s.KeyMarker = &v
	return s
}

func (s *ListObjectVersionsRequest) SetMaxKeys(v int64) *ListObjectVersionsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectVersionsRequest) SetPrefix(v string) *ListObjectVersionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListObjectVersionsRequest) SetVersionIdMarker(v string) *ListObjectVersionsRequest {
	s.VersionIdMarker = &v
	return s
}

func (s *ListObjectVersionsRequest) Validate() error {
	return dara.Validate(s)
}
