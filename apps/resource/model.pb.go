// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: apps/resource/pb/model.proto

package resource

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vendor int32

const (
	Vendor_VENDOR_UNSPECIFIED Vendor = 0
	Vendor_VENDOR_ALI         Vendor = 1
	Vendor_VENDOR_TENCENT     Vendor = 2
)

// Enum value maps for Vendor.
var (
	Vendor_name = map[int32]string{
		0: "VENDOR_UNSPECIFIED",
		1: "VENDOR_ALI",
		2: "VENDOR_TENCENT",
	}
	Vendor_value = map[string]int32{
		"VENDOR_UNSPECIFIED": 0,
		"VENDOR_ALI":         1,
		"VENDOR_TENCENT":     2,
	}
)

func (x Vendor) Enum() *Vendor {
	p := new(Vendor)
	*p = x
	return p
}

func (x Vendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vendor) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_resource_pb_model_proto_enumTypes[0].Descriptor()
}

func (Vendor) Type() protoreflect.EnumType {
	return &file_apps_resource_pb_model_proto_enumTypes[0]
}

func (x Vendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vendor.Descriptor instead.
func (Vendor) EnumDescriptor() ([]byte, []int) {
	return file_apps_resource_pb_model_proto_rawDescGZIP(), []int{0}
}

// Type 资源类型
type Type int32

const (
	Type_TYPE_UNSPECIFIED Type = 0
	Type_TYPE_ECS         Type = 1
	Type_TYPE_RDS         Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_ECS",
		2: "TYPE_RDS",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_ECS":         1,
		"TYPE_RDS":         2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_resource_pb_model_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_apps_resource_pb_model_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_apps_resource_pb_model_proto_rawDescGZIP(), []int{1}
}

type ResourceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Resource `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ResourceSet) Reset() {
	*x = ResourceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSet) ProtoMessage() {}

func (x *ResourceSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSet.ProtoReflect.Descriptor instead.
func (*ResourceSet) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_model_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResourceSet) GetItems() []*Resource {
	if x != nil {
		return x.Items
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"embedded"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty" gorm:"embedded"`
	// @gotags: gorm:"embedded"
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty" gorm:"embedded"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Resource) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"created_at" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"created_at" gorm:"primaryKey"`
	// @gotags: json:"created_at" gorm:"autoCreateTime"
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at" gorm:"autoCreateTime"`
	// @gotags: json:"updated_at" gorm:"autoUpdateTime"
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at" gorm:"autoUpdateTime"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_model_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Meta) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Meta) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// --------------------------------------------- 通用属性 ---------------------------------------------
	Vendor Vendor `protobuf:"varint,1,opt,name=vendor,proto3,enum=cmdb.resource.Vendor" json:"vendor,omitempty"`
	Type   Type   `protobuf:"varint,2,opt,name=type,proto3,enum=cmdb.resource.Type" json:"type,omitempty"`
	// @gotags: gorm:"unique"
	ResourceId string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty" gorm:"unique"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Region     string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Zone       string `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone,omitempty"`
	Status     string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// --------------------------------------------- 运维属性 ---------------------------------------------
	// @gotags: gorm:"json"
	PrivateAddress []string `protobuf:"bytes,8,rep,name=private_address,json=privateAddress,proto3" json:"private_address,omitempty" gorm:"json"`
	// @gotags: gorm:"json"
	PublicAddress []string `protobuf:"bytes,9,rep,name=public_address,json=publicAddress,proto3" json:"public_address,omitempty" gorm:"json"`
	Cpu           int32    `protobuf:"varint,10,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory        int64    `protobuf:"varint,11,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage       int64    `protobuf:"varint,12,opt,name=storage,proto3" json:"storage,omitempty"`
	// --------------------------------------------- 其他属性 ---------------------------------------------
	// @gotags: gorm:"json"
	Extra map[string]string `protobuf:"bytes,13,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" gorm:"json"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_model_proto_rawDescGZIP(), []int{3}
}

func (x *Spec) GetVendor() Vendor {
	if x != nil {
		return x.Vendor
	}
	return Vendor_VENDOR_UNSPECIFIED
}

func (x *Spec) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *Spec) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Spec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Spec) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Spec) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Spec) GetPrivateAddress() []string {
	if x != nil {
		return x.PrivateAddress
	}
	return nil
}

func (x *Spec) GetPublicAddress() []string {
	if x != nil {
		return x.PublicAddress
	}
	return nil
}

func (x *Spec) GetCpu() int32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Spec) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Spec) GetStorage() int64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

func (x *Spec) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_apps_resource_pb_model_proto protoreflect.FileDescriptor

var file_apps_resource_pb_model_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x52, 0x0a,
	0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x5c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22,
	0x54, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdb, 0x03, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2d,
	0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x27, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x70,
	0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x44, 0x0a, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x12, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f,
	0x41, 0x4c, 0x49, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f,
	0x54, 0x45, 0x4e, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0x38, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x43, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x44,
	0x53, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_resource_pb_model_proto_rawDescOnce sync.Once
	file_apps_resource_pb_model_proto_rawDescData = file_apps_resource_pb_model_proto_rawDesc
)

func file_apps_resource_pb_model_proto_rawDescGZIP() []byte {
	file_apps_resource_pb_model_proto_rawDescOnce.Do(func() {
		file_apps_resource_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_resource_pb_model_proto_rawDescData)
	})
	return file_apps_resource_pb_model_proto_rawDescData
}

var file_apps_resource_pb_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_resource_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_resource_pb_model_proto_goTypes = []any{
	(Vendor)(0),         // 0: cmdb.resource.Vendor
	(Type)(0),           // 1: cmdb.resource.Type
	(*ResourceSet)(nil), // 2: cmdb.resource.ResourceSet
	(*Resource)(nil),    // 3: cmdb.resource.Resource
	(*Meta)(nil),        // 4: cmdb.resource.Meta
	(*Spec)(nil),        // 5: cmdb.resource.Spec
	nil,                 // 6: cmdb.resource.Spec.ExtraEntry
}
var file_apps_resource_pb_model_proto_depIdxs = []int32{
	3, // 0: cmdb.resource.ResourceSet.items:type_name -> cmdb.resource.Resource
	4, // 1: cmdb.resource.Resource.meta:type_name -> cmdb.resource.Meta
	5, // 2: cmdb.resource.Resource.spec:type_name -> cmdb.resource.Spec
	0, // 3: cmdb.resource.Spec.vendor:type_name -> cmdb.resource.Vendor
	1, // 4: cmdb.resource.Spec.type:type_name -> cmdb.resource.Type
	6, // 5: cmdb.resource.Spec.extra:type_name -> cmdb.resource.Spec.ExtraEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_resource_pb_model_proto_init() }
func file_apps_resource_pb_model_proto_init() {
	if File_apps_resource_pb_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_resource_pb_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_resource_pb_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_resource_pb_model_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_resource_pb_model_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Spec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_resource_pb_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_resource_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_resource_pb_model_proto_depIdxs,
		EnumInfos:         file_apps_resource_pb_model_proto_enumTypes,
		MessageInfos:      file_apps_resource_pb_model_proto_msgTypes,
	}.Build()
	File_apps_resource_pb_model_proto = out.File
	file_apps_resource_pb_model_proto_rawDesc = nil
	file_apps_resource_pb_model_proto_goTypes = nil
	file_apps_resource_pb_model_proto_depIdxs = nil
}
