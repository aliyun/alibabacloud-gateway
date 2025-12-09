// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeCachePrefetchJob interface {
	dara.Model
	String() string
	GoString() string
	SetExcludes(v []*string) *CreateDataLakeCachePrefetchJob
	GetExcludes() []*string
	SetIncludes(v []*string) *CreateDataLakeCachePrefetchJob
	GetIncludes() []*string
	SetTag(v string) *CreateDataLakeCachePrefetchJob
	GetTag() *string
}

type CreateDataLakeCachePrefetchJob struct {
	Excludes []*string `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Repeated"`
	Includes []*string `json:"Includes,omitempty" xml:"Includes,omitempty" type:"Repeated"`
	// example:
	//
	// Jobxxx
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateDataLakeCachePrefetchJob) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeCachePrefetchJob) GoString() string {
	return s.String()
}

func (s *CreateDataLakeCachePrefetchJob) GetExcludes() []*string {
	return s.Excludes
}

func (s *CreateDataLakeCachePrefetchJob) GetIncludes() []*string {
	return s.Includes
}

func (s *CreateDataLakeCachePrefetchJob) GetTag() *string {
	return s.Tag
}

func (s *CreateDataLakeCachePrefetchJob) SetExcludes(v []*string) *CreateDataLakeCachePrefetchJob {
	s.Excludes = v
	return s
}

func (s *CreateDataLakeCachePrefetchJob) SetIncludes(v []*string) *CreateDataLakeCachePrefetchJob {
	s.Includes = v
	return s
}

func (s *CreateDataLakeCachePrefetchJob) SetTag(v string) *CreateDataLakeCachePrefetchJob {
	s.Tag = &v
	return s
}

func (s *CreateDataLakeCachePrefetchJob) Validate() error {
	return dara.Validate(s)
}
