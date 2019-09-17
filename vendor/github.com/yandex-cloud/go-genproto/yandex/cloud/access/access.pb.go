// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/access/access.proto

package access

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccessBindingAction int32

const (
	AccessBindingAction_ACCESS_BINDING_ACTION_UNSPECIFIED AccessBindingAction = 0
	// Addition of an access binding.
	AccessBindingAction_ADD AccessBindingAction = 1
	// Removal of an access binding.
	AccessBindingAction_REMOVE AccessBindingAction = 2
)

var AccessBindingAction_name = map[int32]string{
	0: "ACCESS_BINDING_ACTION_UNSPECIFIED",
	1: "ADD",
	2: "REMOVE",
}

var AccessBindingAction_value = map[string]int32{
	"ACCESS_BINDING_ACTION_UNSPECIFIED": 0,
	"ADD":                               1,
	"REMOVE":                            2,
}

func (x AccessBindingAction) String() string {
	return proto.EnumName(AccessBindingAction_name, int32(x))
}

func (AccessBindingAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{0}
}

type Subject struct {
	// ID of the subject.
	//
	// It can contain one of the following values:
	// * `allAuthenticatedUsers`: A special system identifier that represents anyone
	//    who is authenticated. It can be used only if the [type] is `system`.
	//
	// * `<cloud generated id>`: An identifier that represents a user account.
	//    It can be used only if the [type] is `userAccount` or `serviceAccount`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type of the subject.
	//
	// It can contain one of the following values:
	// * `system`: System group. This type represents several accounts with a common system identifier.
	// * `userAccount`: An user account (for example, "alice.the.girl@yandex.ru"). This type represents the [yandex.cloud.iam.v1.UserAccount] resource.
	// * `serviceAccount`: A service account. This type represents the [yandex.cloud.iam.v1.ServiceAccount] resource.
	//
	// For more information, see [Subject to which the role is assigned](/docs/iam/concepts/access-control/#subject).
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subject) Reset()         { *m = Subject{} }
func (m *Subject) String() string { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()    {}
func (*Subject) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{0}
}

func (m *Subject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subject.Unmarshal(m, b)
}
func (m *Subject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subject.Marshal(b, m, deterministic)
}
func (m *Subject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subject.Merge(m, src)
}
func (m *Subject) XXX_Size() int {
	return xxx_messageInfo_Subject.Size(m)
}
func (m *Subject) XXX_DiscardUnknown() {
	xxx_messageInfo_Subject.DiscardUnknown(m)
}

var xxx_messageInfo_Subject proto.InternalMessageInfo

func (m *Subject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subject) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type AccessBinding struct {
	// ID of the [yandex.cloud.iam.v1.Role] that is assigned to the [subject].
	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// Identity for which access binding is being created.
	// It can represent an account with a unique ID or several accounts with a system identifier.
	Subject              *Subject `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessBinding) Reset()         { *m = AccessBinding{} }
func (m *AccessBinding) String() string { return proto.CompactTextString(m) }
func (*AccessBinding) ProtoMessage()    {}
func (*AccessBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{1}
}

func (m *AccessBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessBinding.Unmarshal(m, b)
}
func (m *AccessBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessBinding.Marshal(b, m, deterministic)
}
func (m *AccessBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessBinding.Merge(m, src)
}
func (m *AccessBinding) XXX_Size() int {
	return xxx_messageInfo_AccessBinding.Size(m)
}
func (m *AccessBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessBinding.DiscardUnknown(m)
}

var xxx_messageInfo_AccessBinding proto.InternalMessageInfo

func (m *AccessBinding) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *AccessBinding) GetSubject() *Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

type ListAccessBindingsRequest struct {
	// ID of the resource to list access bindings for.
	//
	// To get the resource ID, use a corresponding List request.
	// For example, use the [yandex.cloud.resourcemanager.v1.CloudService.List] request to get the Cloud resource ID.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListAccessBindingsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set [page_token]
	// to the [ListAccessBindingsResponse.next_page_token]
	// returned by a previous list request to get the next page of results.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessBindingsRequest) Reset()         { *m = ListAccessBindingsRequest{} }
func (m *ListAccessBindingsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccessBindingsRequest) ProtoMessage()    {}
func (*ListAccessBindingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{2}
}

func (m *ListAccessBindingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessBindingsRequest.Unmarshal(m, b)
}
func (m *ListAccessBindingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessBindingsRequest.Marshal(b, m, deterministic)
}
func (m *ListAccessBindingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessBindingsRequest.Merge(m, src)
}
func (m *ListAccessBindingsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccessBindingsRequest.Size(m)
}
func (m *ListAccessBindingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessBindingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessBindingsRequest proto.InternalMessageInfo

func (m *ListAccessBindingsRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ListAccessBindingsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAccessBindingsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListAccessBindingsResponse struct {
	// List of access bindings for the specified resource.
	AccessBindings []*AccessBinding `protobuf:"bytes,1,rep,name=access_bindings,json=accessBindings,proto3" json:"access_bindings,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListAccessBindingsRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListAccessBindingsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessBindingsResponse) Reset()         { *m = ListAccessBindingsResponse{} }
func (m *ListAccessBindingsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccessBindingsResponse) ProtoMessage()    {}
func (*ListAccessBindingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{3}
}

func (m *ListAccessBindingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessBindingsResponse.Unmarshal(m, b)
}
func (m *ListAccessBindingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessBindingsResponse.Marshal(b, m, deterministic)
}
func (m *ListAccessBindingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessBindingsResponse.Merge(m, src)
}
func (m *ListAccessBindingsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccessBindingsResponse.Size(m)
}
func (m *ListAccessBindingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessBindingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessBindingsResponse proto.InternalMessageInfo

func (m *ListAccessBindingsResponse) GetAccessBindings() []*AccessBinding {
	if m != nil {
		return m.AccessBindings
	}
	return nil
}

func (m *ListAccessBindingsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type SetAccessBindingsRequest struct {
	// ID of the resource for which access bindings are being set.
	//
	// To get the resource ID, use a corresponding List request.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Access bindings to be set. For more information, see [Access Bindings](/docs/iam/concepts/access-control/#access-bindings).
	AccessBindings       []*AccessBinding `protobuf:"bytes,2,rep,name=access_bindings,json=accessBindings,proto3" json:"access_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SetAccessBindingsRequest) Reset()         { *m = SetAccessBindingsRequest{} }
func (m *SetAccessBindingsRequest) String() string { return proto.CompactTextString(m) }
func (*SetAccessBindingsRequest) ProtoMessage()    {}
func (*SetAccessBindingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{4}
}

func (m *SetAccessBindingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAccessBindingsRequest.Unmarshal(m, b)
}
func (m *SetAccessBindingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAccessBindingsRequest.Marshal(b, m, deterministic)
}
func (m *SetAccessBindingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAccessBindingsRequest.Merge(m, src)
}
func (m *SetAccessBindingsRequest) XXX_Size() int {
	return xxx_messageInfo_SetAccessBindingsRequest.Size(m)
}
func (m *SetAccessBindingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAccessBindingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetAccessBindingsRequest proto.InternalMessageInfo

func (m *SetAccessBindingsRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *SetAccessBindingsRequest) GetAccessBindings() []*AccessBinding {
	if m != nil {
		return m.AccessBindings
	}
	return nil
}

type SetAccessBindingsMetadata struct {
	// ID of the resource for which access bindings are being set.
	ResourceId           string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetAccessBindingsMetadata) Reset()         { *m = SetAccessBindingsMetadata{} }
func (m *SetAccessBindingsMetadata) String() string { return proto.CompactTextString(m) }
func (*SetAccessBindingsMetadata) ProtoMessage()    {}
func (*SetAccessBindingsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{5}
}

func (m *SetAccessBindingsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAccessBindingsMetadata.Unmarshal(m, b)
}
func (m *SetAccessBindingsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAccessBindingsMetadata.Marshal(b, m, deterministic)
}
func (m *SetAccessBindingsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAccessBindingsMetadata.Merge(m, src)
}
func (m *SetAccessBindingsMetadata) XXX_Size() int {
	return xxx_messageInfo_SetAccessBindingsMetadata.Size(m)
}
func (m *SetAccessBindingsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAccessBindingsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SetAccessBindingsMetadata proto.InternalMessageInfo

func (m *SetAccessBindingsMetadata) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

type UpdateAccessBindingsRequest struct {
	// ID of the resource for which access bindings are being updated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Updates to access bindings.
	AccessBindingDeltas  []*AccessBindingDelta `protobuf:"bytes,2,rep,name=access_binding_deltas,json=accessBindingDeltas,proto3" json:"access_binding_deltas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateAccessBindingsRequest) Reset()         { *m = UpdateAccessBindingsRequest{} }
func (m *UpdateAccessBindingsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAccessBindingsRequest) ProtoMessage()    {}
func (*UpdateAccessBindingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{6}
}

func (m *UpdateAccessBindingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccessBindingsRequest.Unmarshal(m, b)
}
func (m *UpdateAccessBindingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccessBindingsRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAccessBindingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccessBindingsRequest.Merge(m, src)
}
func (m *UpdateAccessBindingsRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAccessBindingsRequest.Size(m)
}
func (m *UpdateAccessBindingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccessBindingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccessBindingsRequest proto.InternalMessageInfo

func (m *UpdateAccessBindingsRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *UpdateAccessBindingsRequest) GetAccessBindingDeltas() []*AccessBindingDelta {
	if m != nil {
		return m.AccessBindingDeltas
	}
	return nil
}

type UpdateAccessBindingsMetadata struct {
	// ID of the resource for which access bindings are being updated.
	ResourceId           string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAccessBindingsMetadata) Reset()         { *m = UpdateAccessBindingsMetadata{} }
func (m *UpdateAccessBindingsMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateAccessBindingsMetadata) ProtoMessage()    {}
func (*UpdateAccessBindingsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{7}
}

func (m *UpdateAccessBindingsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccessBindingsMetadata.Unmarshal(m, b)
}
func (m *UpdateAccessBindingsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccessBindingsMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateAccessBindingsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccessBindingsMetadata.Merge(m, src)
}
func (m *UpdateAccessBindingsMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateAccessBindingsMetadata.Size(m)
}
func (m *UpdateAccessBindingsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccessBindingsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccessBindingsMetadata proto.InternalMessageInfo

func (m *UpdateAccessBindingsMetadata) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

type AccessBindingDelta struct {
	// The action that is being performed on an access binding.
	Action AccessBindingAction `protobuf:"varint,1,opt,name=action,proto3,enum=yandex.cloud.access.AccessBindingAction" json:"action,omitempty"`
	// Access binding. For more information, see [Access Bindings](/docs/iam/concepts/access-control/#access-bindings).
	AccessBinding        *AccessBinding `protobuf:"bytes,2,opt,name=access_binding,json=accessBinding,proto3" json:"access_binding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AccessBindingDelta) Reset()         { *m = AccessBindingDelta{} }
func (m *AccessBindingDelta) String() string { return proto.CompactTextString(m) }
func (*AccessBindingDelta) ProtoMessage()    {}
func (*AccessBindingDelta) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72fa1116199e82e, []int{8}
}

func (m *AccessBindingDelta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessBindingDelta.Unmarshal(m, b)
}
func (m *AccessBindingDelta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessBindingDelta.Marshal(b, m, deterministic)
}
func (m *AccessBindingDelta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessBindingDelta.Merge(m, src)
}
func (m *AccessBindingDelta) XXX_Size() int {
	return xxx_messageInfo_AccessBindingDelta.Size(m)
}
func (m *AccessBindingDelta) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessBindingDelta.DiscardUnknown(m)
}

var xxx_messageInfo_AccessBindingDelta proto.InternalMessageInfo

func (m *AccessBindingDelta) GetAction() AccessBindingAction {
	if m != nil {
		return m.Action
	}
	return AccessBindingAction_ACCESS_BINDING_ACTION_UNSPECIFIED
}

func (m *AccessBindingDelta) GetAccessBinding() *AccessBinding {
	if m != nil {
		return m.AccessBinding
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.access.AccessBindingAction", AccessBindingAction_name, AccessBindingAction_value)
	proto.RegisterType((*Subject)(nil), "yandex.cloud.access.Subject")
	proto.RegisterType((*AccessBinding)(nil), "yandex.cloud.access.AccessBinding")
	proto.RegisterType((*ListAccessBindingsRequest)(nil), "yandex.cloud.access.ListAccessBindingsRequest")
	proto.RegisterType((*ListAccessBindingsResponse)(nil), "yandex.cloud.access.ListAccessBindingsResponse")
	proto.RegisterType((*SetAccessBindingsRequest)(nil), "yandex.cloud.access.SetAccessBindingsRequest")
	proto.RegisterType((*SetAccessBindingsMetadata)(nil), "yandex.cloud.access.SetAccessBindingsMetadata")
	proto.RegisterType((*UpdateAccessBindingsRequest)(nil), "yandex.cloud.access.UpdateAccessBindingsRequest")
	proto.RegisterType((*UpdateAccessBindingsMetadata)(nil), "yandex.cloud.access.UpdateAccessBindingsMetadata")
	proto.RegisterType((*AccessBindingDelta)(nil), "yandex.cloud.access.AccessBindingDelta")
}

func init() { proto.RegisterFile("yandex/cloud/access/access.proto", fileDescriptor_f72fa1116199e82e) }

var fileDescriptor_f72fa1116199e82e = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x3f, 0x27, 0xfd, 0x92, 0xe6, 0x86, 0xa4, 0xd1, 0x44, 0x48, 0x6e, 0x29, 0x22, 0xb5,
	0x04, 0x8d, 0x90, 0xe2, 0xfc, 0x41, 0x88, 0x05, 0x29, 0x10, 0x27, 0x29, 0xb2, 0xa0, 0x49, 0x6b,
	0xb7, 0x2c, 0xd8, 0x58, 0x13, 0xcf, 0x28, 0x18, 0x82, 0x6d, 0x32, 0x13, 0xd4, 0xf6, 0x11, 0xba,
	0x63, 0x0f, 0x8f, 0x80, 0x78, 0x8c, 0xf6, 0x51, 0x78, 0x06, 0x56, 0xc8, 0x63, 0xa7, 0x8a, 0x89,
	0xa5, 0x66, 0xd1, 0xd5, 0x58, 0xbe, 0xe7, 0x9e, 0xfb, 0x3b, 0x33, 0x9a, 0x81, 0xca, 0x19, 0x76,
	0x09, 0x3d, 0xad, 0xdb, 0x13, 0x6f, 0x46, 0xea, 0xd8, 0xb6, 0x29, 0x63, 0xd1, 0xa2, 0xfa, 0x53,
	0x8f, 0x7b, 0xa8, 0x1c, 0x2a, 0x54, 0xa1, 0x50, 0xc3, 0xd2, 0xd6, 0xfd, 0x58, 0xdb, 0x57, 0x3c,
	0x71, 0x08, 0xe6, 0x8e, 0xe7, 0x86, 0x3d, 0xca, 0x33, 0xc8, 0x9a, 0xb3, 0xd1, 0x47, 0x6a, 0x73,
	0x24, 0x43, 0xca, 0x21, 0xb2, 0x54, 0x91, 0xaa, 0x39, 0x6d, 0xfd, 0xe2, 0xaa, 0xb9, 0xd6, 0xde,
	0x7b, 0xda, 0x30, 0x52, 0x0e, 0x41, 0x08, 0xd6, 0xf8, 0x99, 0x4f, 0xe5, 0x54, 0x50, 0x33, 0xc4,
	0xb7, 0xe2, 0x43, 0xa1, 0x23, 0x26, 0x68, 0x8e, 0x4b, 0x1c, 0x77, 0x8c, 0x76, 0x20, 0x3b, 0xf5,
	0x26, 0xd4, 0x4a, 0xf0, 0xc8, 0x04, 0x05, 0x9d, 0xa0, 0x36, 0x64, 0x59, 0x38, 0x4c, 0x58, 0xe5,
	0x5b, 0xdb, 0x6a, 0x02, 0xb2, 0x1a, 0x01, 0x69, 0x6b, 0xbf, 0x2f, 0x9b, 0x92, 0x31, 0x6f, 0x51,
	0x7e, 0x48, 0xb0, 0xf9, 0xd6, 0x61, 0x3c, 0x36, 0x96, 0x19, 0xf4, 0xcb, 0x8c, 0x32, 0x8e, 0x6a,
	0x90, 0x9f, 0x52, 0xe6, 0xcd, 0xa6, 0xf6, 0x02, 0xc2, 0x9d, 0xc0, 0xe1, 0x1a, 0x03, 0xe6, 0x02,
	0x9d, 0xa0, 0x5d, 0xc8, 0xf9, 0x78, 0x4c, 0x2d, 0xe6, 0x9c, 0x87, 0xb9, 0xd2, 0x1a, 0xfc, 0xb9,
	0x6c, 0x66, 0xda, 0x7b, 0xcd, 0x46, 0xa3, 0x61, 0xac, 0x07, 0x45, 0xd3, 0x39, 0xa7, 0xa8, 0x0a,
	0x20, 0x84, 0xdc, 0xfb, 0x44, 0x5d, 0x39, 0x2d, 0x6c, 0x73, 0x17, 0x57, 0xcd, 0xff, 0x85, 0xd2,
	0x10, 0x2e, 0xc7, 0x41, 0x4d, 0xf9, 0x26, 0xc1, 0x56, 0x12, 0x1f, 0xf3, 0x3d, 0x97, 0x51, 0xf4,
	0x06, 0x36, 0xc2, 0x7c, 0xd6, 0x28, 0x2a, 0xc9, 0x52, 0x25, 0x5d, 0xcd, 0xb7, 0x94, 0xc4, 0x4d,
	0x88, 0xb9, 0x18, 0x45, 0x1c, 0x33, 0x45, 0x8f, 0x60, 0xc3, 0xa5, 0xa7, 0xdc, 0x5a, 0x40, 0x0b,
	0x0f, 0xa7, 0x10, 0xfc, 0x3e, 0xbc, 0x66, 0xfa, 0x2e, 0x81, 0x6c, 0xd2, 0xdb, 0xd9, 0xb2, 0xa3,
	0xe5, 0x00, 0xa9, 0x55, 0x03, 0x44, 0x67, 0xf9, 0x4f, 0x0c, 0xa5, 0x0d, 0x9b, 0x4b, 0x74, 0x07,
	0x94, 0x63, 0x82, 0x39, 0x46, 0x0f, 0x12, 0xf0, 0x16, 0x81, 0x94, 0x5f, 0x12, 0xdc, 0x3b, 0xf1,
	0x09, 0xe6, 0xf4, 0x56, 0xf2, 0x61, 0xb8, 0x1b, 0xcf, 0x67, 0x11, 0x3a, 0xe1, 0x78, 0x9e, 0x72,
	0xf7, 0xe6, 0x94, 0xbd, 0x40, 0x1f, 0x45, 0x2d, 0xe3, 0xa5, 0x0a, 0x53, 0x5e, 0xc2, 0x76, 0x12,
	0xf0, 0xea, 0x91, 0x7f, 0x4a, 0x80, 0x96, 0x47, 0xa2, 0x7d, 0xc8, 0x60, 0x3b, 0xb8, 0xd5, 0xa2,
	0xa5, 0xd8, 0xaa, 0xde, 0xcc, 0xda, 0x11, 0xfa, 0x08, 0x36, 0xea, 0x46, 0x43, 0x28, 0xc6, 0xb7,
	0x20, 0xba, 0xa7, 0xab, 0x9f, 0x70, 0x21, 0x16, 0xfb, 0xf1, 0x11, 0x94, 0x13, 0xa6, 0xa2, 0x87,
	0xb0, 0xd3, 0xe9, 0x76, 0xfb, 0xa6, 0x69, 0x69, 0xfa, 0xa0, 0xa7, 0x0f, 0x5e, 0x5b, 0x9d, 0xee,
	0xb1, 0x3e, 0x1c, 0x58, 0x27, 0x03, 0xf3, 0xb0, 0xdf, 0xd5, 0xf7, 0xf5, 0x7e, 0xaf, 0xf4, 0x1f,
	0xca, 0x42, 0xba, 0xd3, 0xeb, 0x95, 0x24, 0x04, 0x90, 0x31, 0xfa, 0x07, 0xc3, 0x77, 0xfd, 0x52,
	0x4a, 0x7b, 0xf5, 0xfe, 0xc5, 0xd8, 0xe1, 0x1f, 0x66, 0x23, 0xd5, 0xf6, 0x3e, 0xd7, 0x43, 0xae,
	0x5a, 0xf8, 0xba, 0x8d, 0xbd, 0xda, 0x98, 0xba, 0xe2, 0x61, 0xab, 0x27, 0xbc, 0x96, 0xcf, 0xc3,
	0x65, 0x94, 0x11, 0x8a, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xce, 0x12, 0xcf, 0x52,
	0x05, 0x00, 0x00,
}
