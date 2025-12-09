// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLifecycleRule interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleAbortMultipartUpload(v *LifecycleRuleLifecycleAbortMultipartUpload) *LifecycleRule
	GetLifecycleAbortMultipartUpload() *LifecycleRuleLifecycleAbortMultipartUpload
	SetAtimeBase(v int64) *LifecycleRule
	GetAtimeBase() *int64
	SetLifecycleExpiration(v *LifecycleRuleLifecycleExpiration) *LifecycleRule
	GetLifecycleExpiration() *LifecycleRuleLifecycleExpiration
	SetFilter(v *LifecycleRuleFilter) *LifecycleRule
	GetFilter() *LifecycleRuleFilter
	SetID(v string) *LifecycleRule
	GetID() *string
	SetNoncurrentVersionExpiration(v *LifecycleRuleNoncurrentVersionExpiration) *LifecycleRule
	GetNoncurrentVersionExpiration() *LifecycleRuleNoncurrentVersionExpiration
	SetNoncurrentVersionTransition(v []*LifecycleRuleNoncurrentVersionTransition) *LifecycleRule
	GetNoncurrentVersionTransition() []*LifecycleRuleNoncurrentVersionTransition
	SetPrefix(v string) *LifecycleRule
	GetPrefix() *string
	SetStatus(v string) *LifecycleRule
	GetStatus() *string
	SetTag(v []*Tag) *LifecycleRule
	GetTag() []*Tag
	SetLifecycleTransition(v []*LifecycleRuleLifecycleTransition) *LifecycleRule
	GetLifecycleTransition() []*LifecycleRuleLifecycleTransition
}

type LifecycleRule struct {
	LifecycleAbortMultipartUpload *LifecycleRuleLifecycleAbortMultipartUpload `json:"AbortMultipartUpload,omitempty" xml:"AbortMultipartUpload,omitempty" type:"Struct"`
	// example:
	//
	// 1626158525
	AtimeBase           *int64                            `json:"AtimeBase,omitempty" xml:"AtimeBase,omitempty"`
	LifecycleExpiration *LifecycleRuleLifecycleExpiration `json:"Expiration,omitempty" xml:"Expiration,omitempty" type:"Struct"`
	Filter              *LifecycleRuleFilter              `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// rule1
	ID                          *string                                     `json:"ID,omitempty" xml:"ID,omitempty"`
	NoncurrentVersionExpiration *LifecycleRuleNoncurrentVersionExpiration   `json:"NoncurrentVersionExpiration,omitempty" xml:"NoncurrentVersionExpiration,omitempty" type:"Struct"`
	NoncurrentVersionTransition []*LifecycleRuleNoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty" xml:"NoncurrentVersionTransition,omitempty" type:"Repeated"`
	// example:
	//
	// logs/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status              *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag                 []*Tag                              `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	LifecycleTransition []*LifecycleRuleLifecycleTransition `json:"Transition,omitempty" xml:"Transition,omitempty" type:"Repeated"`
}

func (s LifecycleRule) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRule) GoString() string {
	return s.String()
}

func (s *LifecycleRule) GetLifecycleAbortMultipartUpload() *LifecycleRuleLifecycleAbortMultipartUpload {
	return s.LifecycleAbortMultipartUpload
}

func (s *LifecycleRule) GetAtimeBase() *int64 {
	return s.AtimeBase
}

func (s *LifecycleRule) GetLifecycleExpiration() *LifecycleRuleLifecycleExpiration {
	return s.LifecycleExpiration
}

func (s *LifecycleRule) GetFilter() *LifecycleRuleFilter {
	return s.Filter
}

func (s *LifecycleRule) GetID() *string {
	return s.ID
}

func (s *LifecycleRule) GetNoncurrentVersionExpiration() *LifecycleRuleNoncurrentVersionExpiration {
	return s.NoncurrentVersionExpiration
}

func (s *LifecycleRule) GetNoncurrentVersionTransition() []*LifecycleRuleNoncurrentVersionTransition {
	return s.NoncurrentVersionTransition
}

func (s *LifecycleRule) GetPrefix() *string {
	return s.Prefix
}

func (s *LifecycleRule) GetStatus() *string {
	return s.Status
}

func (s *LifecycleRule) GetTag() []*Tag {
	return s.Tag
}

func (s *LifecycleRule) GetLifecycleTransition() []*LifecycleRuleLifecycleTransition {
	return s.LifecycleTransition
}

func (s *LifecycleRule) SetLifecycleAbortMultipartUpload(v *LifecycleRuleLifecycleAbortMultipartUpload) *LifecycleRule {
	s.LifecycleAbortMultipartUpload = v
	return s
}

func (s *LifecycleRule) SetAtimeBase(v int64) *LifecycleRule {
	s.AtimeBase = &v
	return s
}

func (s *LifecycleRule) SetLifecycleExpiration(v *LifecycleRuleLifecycleExpiration) *LifecycleRule {
	s.LifecycleExpiration = v
	return s
}

func (s *LifecycleRule) SetFilter(v *LifecycleRuleFilter) *LifecycleRule {
	s.Filter = v
	return s
}

func (s *LifecycleRule) SetID(v string) *LifecycleRule {
	s.ID = &v
	return s
}

func (s *LifecycleRule) SetNoncurrentVersionExpiration(v *LifecycleRuleNoncurrentVersionExpiration) *LifecycleRule {
	s.NoncurrentVersionExpiration = v
	return s
}

func (s *LifecycleRule) SetNoncurrentVersionTransition(v []*LifecycleRuleNoncurrentVersionTransition) *LifecycleRule {
	s.NoncurrentVersionTransition = v
	return s
}

func (s *LifecycleRule) SetPrefix(v string) *LifecycleRule {
	s.Prefix = &v
	return s
}

func (s *LifecycleRule) SetStatus(v string) *LifecycleRule {
	s.Status = &v
	return s
}

func (s *LifecycleRule) SetTag(v []*Tag) *LifecycleRule {
	s.Tag = v
	return s
}

func (s *LifecycleRule) SetLifecycleTransition(v []*LifecycleRuleLifecycleTransition) *LifecycleRule {
	s.LifecycleTransition = v
	return s
}

func (s *LifecycleRule) Validate() error {
	if s.LifecycleAbortMultipartUpload != nil {
		if err := s.LifecycleAbortMultipartUpload.Validate(); err != nil {
			return err
		}
	}
	if s.LifecycleExpiration != nil {
		if err := s.LifecycleExpiration.Validate(); err != nil {
			return err
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.NoncurrentVersionExpiration != nil {
		if err := s.NoncurrentVersionExpiration.Validate(); err != nil {
			return err
		}
	}
	if s.NoncurrentVersionTransition != nil {
		for _, item := range s.NoncurrentVersionTransition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LifecycleTransition != nil {
		for _, item := range s.LifecycleTransition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LifecycleRuleLifecycleAbortMultipartUpload struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2002-10-11T00:00:00.000Z
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// example:
	//
	// 300
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s LifecycleRuleLifecycleAbortMultipartUpload) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleLifecycleAbortMultipartUpload) GoString() string {
	return s.String()
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) GetCreatedBeforeDate() *string {
	return s.CreatedBeforeDate
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) GetDays() *int32 {
	return s.Days
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) SetCreatedBeforeDate(v string) *LifecycleRuleLifecycleAbortMultipartUpload {
	s.CreatedBeforeDate = &v
	return s
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) SetDays(v int32) *LifecycleRuleLifecycleAbortMultipartUpload {
	s.Days = &v
	return s
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) Validate() error {
	return dara.Validate(s)
}

type LifecycleRuleLifecycleExpiration struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2002-10-11T00:00:00.000Z
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2002-10-11T00:00:00.000Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// example:
	//
	// true
	ExpiredObjectDeleteMarker *bool `json:"ExpiredObjectDeleteMarker,omitempty" xml:"ExpiredObjectDeleteMarker,omitempty"`
}

func (s LifecycleRuleLifecycleExpiration) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleLifecycleExpiration) GoString() string {
	return s.String()
}

func (s *LifecycleRuleLifecycleExpiration) GetCreatedBeforeDate() *string {
	return s.CreatedBeforeDate
}

func (s *LifecycleRuleLifecycleExpiration) GetDate() *string {
	return s.Date
}

func (s *LifecycleRuleLifecycleExpiration) GetDays() *int32 {
	return s.Days
}

func (s *LifecycleRuleLifecycleExpiration) GetExpiredObjectDeleteMarker() *bool {
	return s.ExpiredObjectDeleteMarker
}

func (s *LifecycleRuleLifecycleExpiration) SetCreatedBeforeDate(v string) *LifecycleRuleLifecycleExpiration {
	s.CreatedBeforeDate = &v
	return s
}

func (s *LifecycleRuleLifecycleExpiration) SetDate(v string) *LifecycleRuleLifecycleExpiration {
	s.Date = &v
	return s
}

func (s *LifecycleRuleLifecycleExpiration) SetDays(v int32) *LifecycleRuleLifecycleExpiration {
	s.Days = &v
	return s
}

func (s *LifecycleRuleLifecycleExpiration) SetExpiredObjectDeleteMarker(v bool) *LifecycleRuleLifecycleExpiration {
	s.ExpiredObjectDeleteMarker = &v
	return s
}

func (s *LifecycleRuleLifecycleExpiration) Validate() error {
	return dara.Validate(s)
}

type LifecycleRuleFilter struct {
	Not *LifecycleRuleFilterNot `json:"Not,omitempty" xml:"Not,omitempty" type:"Struct"`
	// example:
	//
	// 10240
	ObjectSizeGreaterThan *int64 `json:"ObjectSizeGreaterThan,omitempty" xml:"ObjectSizeGreaterThan,omitempty"`
	// example:
	//
	// 10240
	ObjectSizeLessThan *int64 `json:"ObjectSizeLessThan,omitempty" xml:"ObjectSizeLessThan,omitempty"`
}

func (s LifecycleRuleFilter) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleFilter) GoString() string {
	return s.String()
}

func (s *LifecycleRuleFilter) GetNot() *LifecycleRuleFilterNot {
	return s.Not
}

func (s *LifecycleRuleFilter) GetObjectSizeGreaterThan() *int64 {
	return s.ObjectSizeGreaterThan
}

func (s *LifecycleRuleFilter) GetObjectSizeLessThan() *int64 {
	return s.ObjectSizeLessThan
}

func (s *LifecycleRuleFilter) SetNot(v *LifecycleRuleFilterNot) *LifecycleRuleFilter {
	s.Not = v
	return s
}

func (s *LifecycleRuleFilter) SetObjectSizeGreaterThan(v int64) *LifecycleRuleFilter {
	s.ObjectSizeGreaterThan = &v
	return s
}

func (s *LifecycleRuleFilter) SetObjectSizeLessThan(v int64) *LifecycleRuleFilter {
	s.ObjectSizeLessThan = &v
	return s
}

func (s *LifecycleRuleFilter) Validate() error {
	if s.Not != nil {
		if err := s.Not.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LifecycleRuleFilterNot struct {
	// example:
	//
	// logs/keep/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Tag    *Tag    `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s LifecycleRuleFilterNot) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleFilterNot) GoString() string {
	return s.String()
}

func (s *LifecycleRuleFilterNot) GetPrefix() *string {
	return s.Prefix
}

func (s *LifecycleRuleFilterNot) GetTag() *Tag {
	return s.Tag
}

func (s *LifecycleRuleFilterNot) SetPrefix(v string) *LifecycleRuleFilterNot {
	s.Prefix = &v
	return s
}

func (s *LifecycleRuleFilterNot) SetTag(v *Tag) *LifecycleRuleFilterNot {
	s.Tag = v
	return s
}

func (s *LifecycleRuleFilterNot) Validate() error {
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LifecycleRuleNoncurrentVersionExpiration struct {
	// example:
	//
	// 350
	NoncurrentDays *int32 `json:"NoncurrentDays,omitempty" xml:"NoncurrentDays,omitempty"`
}

func (s LifecycleRuleNoncurrentVersionExpiration) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleNoncurrentVersionExpiration) GoString() string {
	return s.String()
}

func (s *LifecycleRuleNoncurrentVersionExpiration) GetNoncurrentDays() *int32 {
	return s.NoncurrentDays
}

func (s *LifecycleRuleNoncurrentVersionExpiration) SetNoncurrentDays(v int32) *LifecycleRuleNoncurrentVersionExpiration {
	s.NoncurrentDays = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionExpiration) Validate() error {
	return dara.Validate(s)
}

type LifecycleRuleNoncurrentVersionTransition struct {
	// example:
	//
	// true
	AllowSmallFile *bool `json:"AllowSmallFile,omitempty" xml:"AllowSmallFile,omitempty"`
	// example:
	//
	// true
	IsAccessTime *bool `json:"IsAccessTime,omitempty" xml:"IsAccessTime,omitempty"`
	// example:
	//
	// 200
	NoncurrentDays *int32 `json:"NoncurrentDays,omitempty" xml:"NoncurrentDays,omitempty"`
	// example:
	//
	// false
	ReturnToStdWhenVisit *bool   `json:"ReturnToStdWhenVisit,omitempty" xml:"ReturnToStdWhenVisit,omitempty"`
	StorageClass         *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s LifecycleRuleNoncurrentVersionTransition) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleNoncurrentVersionTransition) GoString() string {
	return s.String()
}

func (s *LifecycleRuleNoncurrentVersionTransition) GetAllowSmallFile() *bool {
	return s.AllowSmallFile
}

func (s *LifecycleRuleNoncurrentVersionTransition) GetIsAccessTime() *bool {
	return s.IsAccessTime
}

func (s *LifecycleRuleNoncurrentVersionTransition) GetNoncurrentDays() *int32 {
	return s.NoncurrentDays
}

func (s *LifecycleRuleNoncurrentVersionTransition) GetReturnToStdWhenVisit() *bool {
	return s.ReturnToStdWhenVisit
}

func (s *LifecycleRuleNoncurrentVersionTransition) GetStorageClass() *string {
	return s.StorageClass
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetAllowSmallFile(v bool) *LifecycleRuleNoncurrentVersionTransition {
	s.AllowSmallFile = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetIsAccessTime(v bool) *LifecycleRuleNoncurrentVersionTransition {
	s.IsAccessTime = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetNoncurrentDays(v int32) *LifecycleRuleNoncurrentVersionTransition {
	s.NoncurrentDays = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetReturnToStdWhenVisit(v bool) *LifecycleRuleNoncurrentVersionTransition {
	s.ReturnToStdWhenVisit = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetStorageClass(v string) *LifecycleRuleNoncurrentVersionTransition {
	s.StorageClass = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionTransition) Validate() error {
	return dara.Validate(s)
}

type LifecycleRuleLifecycleTransition struct {
	// example:
	//
	// true
	AllowSmallFile *bool `json:"AllowSmallFile,omitempty" xml:"AllowSmallFile,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2002-10-11T00:00:00.000Z
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	// example:
	//
	// 200
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// example:
	//
	// true
	IsAccessTime *bool `json:"IsAccessTime,omitempty" xml:"IsAccessTime,omitempty"`
	// example:
	//
	// false
	ReturnToStdWhenVisit *bool   `json:"ReturnToStdWhenVisit,omitempty" xml:"ReturnToStdWhenVisit,omitempty"`
	StorageClass         *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s LifecycleRuleLifecycleTransition) String() string {
	return dara.Prettify(s)
}

func (s LifecycleRuleLifecycleTransition) GoString() string {
	return s.String()
}

func (s *LifecycleRuleLifecycleTransition) GetAllowSmallFile() *bool {
	return s.AllowSmallFile
}

func (s *LifecycleRuleLifecycleTransition) GetCreatedBeforeDate() *string {
	return s.CreatedBeforeDate
}

func (s *LifecycleRuleLifecycleTransition) GetDays() *int32 {
	return s.Days
}

func (s *LifecycleRuleLifecycleTransition) GetIsAccessTime() *bool {
	return s.IsAccessTime
}

func (s *LifecycleRuleLifecycleTransition) GetReturnToStdWhenVisit() *bool {
	return s.ReturnToStdWhenVisit
}

func (s *LifecycleRuleLifecycleTransition) GetStorageClass() *string {
	return s.StorageClass
}

func (s *LifecycleRuleLifecycleTransition) SetAllowSmallFile(v bool) *LifecycleRuleLifecycleTransition {
	s.AllowSmallFile = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetCreatedBeforeDate(v string) *LifecycleRuleLifecycleTransition {
	s.CreatedBeforeDate = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetDays(v int32) *LifecycleRuleLifecycleTransition {
	s.Days = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetIsAccessTime(v bool) *LifecycleRuleLifecycleTransition {
	s.IsAccessTime = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetReturnToStdWhenVisit(v bool) *LifecycleRuleLifecycleTransition {
	s.ReturnToStdWhenVisit = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetStorageClass(v string) *LifecycleRuleLifecycleTransition {
	s.StorageClass = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) Validate() error {
	return dara.Validate(s)
}
