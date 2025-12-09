// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationManifestGenerator interface {
	dara.Model
	String() string
	GoString() string
	SetPrefix(v string) *BatchOperationManifestGenerator
	GetPrefix() *string
	SetSourceBucket(v string) *BatchOperationManifestGenerator
	GetSourceBucket() *string
}

type BatchOperationManifestGenerator struct {
	// example:
	//
	// /
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bucketname
	SourceBucket *string `json:"SourceBucket,omitempty" xml:"SourceBucket,omitempty"`
}

func (s BatchOperationManifestGenerator) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationManifestGenerator) GoString() string {
	return s.String()
}

func (s *BatchOperationManifestGenerator) GetPrefix() *string {
	return s.Prefix
}

func (s *BatchOperationManifestGenerator) GetSourceBucket() *string {
	return s.SourceBucket
}

func (s *BatchOperationManifestGenerator) SetPrefix(v string) *BatchOperationManifestGenerator {
	s.Prefix = &v
	return s
}

func (s *BatchOperationManifestGenerator) SetSourceBucket(v string) *BatchOperationManifestGenerator {
	s.SourceBucket = &v
	return s
}

func (s *BatchOperationManifestGenerator) Validate() error {
	return dara.Validate(s)
}
