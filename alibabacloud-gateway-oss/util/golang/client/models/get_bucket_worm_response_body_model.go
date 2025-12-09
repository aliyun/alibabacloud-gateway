// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketWormResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetWormConfiguration(v *GetBucketWormResponseBodyWormConfiguration) *GetBucketWormResponseBody
	GetWormConfiguration() *GetBucketWormResponseBodyWormConfiguration
}

type GetBucketWormResponseBody struct {
	// The container that stores the information about retention policies of the bucket.
	WormConfiguration *GetBucketWormResponseBodyWormConfiguration `json:"WormConfiguration,omitempty" xml:"WormConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketWormResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketWormResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketWormResponseBody) GetWormConfiguration() *GetBucketWormResponseBodyWormConfiguration {
	return s.WormConfiguration
}

func (s *GetBucketWormResponseBody) SetWormConfiguration(v *GetBucketWormResponseBodyWormConfiguration) *GetBucketWormResponseBody {
	s.WormConfiguration = v
	return s
}

func (s *GetBucketWormResponseBody) Validate() error {
	if s.WormConfiguration != nil {
		if err := s.WormConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketWormResponseBodyWormConfiguration struct {
	// The time at which the retention policy was created.
	//
	// example:
	//
	// 2020-10-15T15:50:32
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The time at which the retention policy will be expired.
	//
	// example:
	//
	// 2020-10-16T15:50:32
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The number of days for which objects can be retained.
	//
	// example:
	//
	// 20
	RetentionPeriodInDays *int32 `json:"RetentionPeriodInDays,omitempty" xml:"RetentionPeriodInDays,omitempty"`
	// The status of the retention policy. Valid values:
	//
	// - InProgress: indicates that the retention policy is in the InProgress state. By default, a retention policy is in the InProgress state after it is created. The policy remains in this state for 24 hours.
	//
	// - Locked: indicates that the retention policy is in the Locked state.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the retention policy.
	//
	// >Note If the specified retention policy ID that is used to query the retention policy configurations of the bucket does not exist, OSS returns the 404 error code.
	//
	// example:
	//
	// 1666E2CFB2B3418****
	WormId *string `json:"WormId,omitempty" xml:"WormId,omitempty"`
}

func (s GetBucketWormResponseBodyWormConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketWormResponseBodyWormConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketWormResponseBodyWormConfiguration) GetCreationDate() *string {
	return s.CreationDate
}

func (s *GetBucketWormResponseBodyWormConfiguration) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *GetBucketWormResponseBodyWormConfiguration) GetRetentionPeriodInDays() *int32 {
	return s.RetentionPeriodInDays
}

func (s *GetBucketWormResponseBodyWormConfiguration) GetState() *string {
	return s.State
}

func (s *GetBucketWormResponseBodyWormConfiguration) GetWormId() *string {
	return s.WormId
}

func (s *GetBucketWormResponseBodyWormConfiguration) SetCreationDate(v string) *GetBucketWormResponseBodyWormConfiguration {
	s.CreationDate = &v
	return s
}

func (s *GetBucketWormResponseBodyWormConfiguration) SetExpirationDate(v string) *GetBucketWormResponseBodyWormConfiguration {
	s.ExpirationDate = &v
	return s
}

func (s *GetBucketWormResponseBodyWormConfiguration) SetRetentionPeriodInDays(v int32) *GetBucketWormResponseBodyWormConfiguration {
	s.RetentionPeriodInDays = &v
	return s
}

func (s *GetBucketWormResponseBodyWormConfiguration) SetState(v string) *GetBucketWormResponseBodyWormConfiguration {
	s.State = &v
	return s
}

func (s *GetBucketWormResponseBodyWormConfiguration) SetWormId(v string) *GetBucketWormResponseBodyWormConfiguration {
	s.WormId = &v
	return s
}

func (s *GetBucketWormResponseBodyWormConfiguration) Validate() error {
	return dara.Validate(s)
}
