// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutSymlinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutSymlinkHeaders
	GetCommonHeaders() map[string]*string
	SetForbidOverwrite(v string) *PutSymlinkHeaders
	GetForbidOverwrite() *string
	SetAcl(v string) *PutSymlinkHeaders
	GetAcl() *string
	SetStorageClass(v string) *PutSymlinkHeaders
	GetStorageClass() *string
	SetSymlinkTargetKey(v string) *PutSymlinkHeaders
	GetSymlinkTargetKey() *string
}

type PutSymlinkHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// Specifies whether the PutSymlink operation overwrites the object that has the same name as that of the symbolic link you want to create.
	//
	//   - If the value of **x-oss-forbid-overwrite*	- is not specified or set to **false**, existing objects can be overwritten by objects that have the same names.
	//
	//   - If the value of **x-oss-forbid-overwrite*	- is set to **true**, existing objects cannot be overwritten by objects that have the same names.
	//
	// If you specify the **x-oss-forbid-overwrite*	- request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite*	- request header to perform a large number of operations (QPS greater than 1,000), contact technical support.
	//
	// > The **x-oss-forbid-overwrite*	- request header is invalid when versioning is enabled or suspended for the destination bucket. In this case, the object with the same name can be overwritten.
	//
	// example:
	//
	// true
	ForbidOverwrite *string `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	// The access control list (ACL) of the object. Default value: default.
	//
	// Valid values:
	//
	//
	//
	// - default: The ACL of the object is the same as that of the bucket in which the object is stored.
	//
	// - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object.
	//
	// - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value.
	//
	// - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value.
	//
	//
	//
	// For more information about the ACL, see **[ACL](https://help.aliyun.com/document_detail/100676.html)**.
	Acl *string `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	// The storage class of the bucket. Default value: Standard.  Valid values:
	//
	//
	//
	// - Standard
	//
	// - IA
	//
	// - Archive
	//
	// - ColdArchive
	StorageClass *string `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	// The target object to which the symbolic link points.
	//
	// The naming conventions for target objects are the same as those for objects.
	//
	//   - Similar to ObjectName, TargetObjectName must be URL-encoded.
	//
	//   - The target object to which a symbolic link points cannot be a symbolic link.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss.jpg
	SymlinkTargetKey *string `json:"x-oss-symlink-target,omitempty" xml:"x-oss-symlink-target,omitempty"`
}

func (s PutSymlinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutSymlinkHeaders) GoString() string {
	return s.String()
}

func (s *PutSymlinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutSymlinkHeaders) GetForbidOverwrite() *string {
	return s.ForbidOverwrite
}

func (s *PutSymlinkHeaders) GetAcl() *string {
	return s.Acl
}

func (s *PutSymlinkHeaders) GetStorageClass() *string {
	return s.StorageClass
}

func (s *PutSymlinkHeaders) GetSymlinkTargetKey() *string {
	return s.SymlinkTargetKey
}

func (s *PutSymlinkHeaders) SetCommonHeaders(v map[string]*string) *PutSymlinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutSymlinkHeaders) SetForbidOverwrite(v string) *PutSymlinkHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *PutSymlinkHeaders) SetAcl(v string) *PutSymlinkHeaders {
	s.Acl = &v
	return s
}

func (s *PutSymlinkHeaders) SetStorageClass(v string) *PutSymlinkHeaders {
	s.StorageClass = &v
	return s
}

func (s *PutSymlinkHeaders) SetSymlinkTargetKey(v string) *PutSymlinkHeaders {
	s.SymlinkTargetKey = &v
	return s
}

func (s *PutSymlinkHeaders) Validate() error {
	return dara.Validate(s)
}
