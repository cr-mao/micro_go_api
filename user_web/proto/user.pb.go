// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PasswrodCheckInfo struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	EncryptedPassword    string   `protobuf:"bytes,2,opt,name=encryptedPassword,proto3" json:"encryptedPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswrodCheckInfo) Reset()         { *m = PasswrodCheckInfo{} }
func (m *PasswrodCheckInfo) String() string { return proto.CompactTextString(m) }
func (*PasswrodCheckInfo) ProtoMessage()    {}
func (*PasswrodCheckInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *PasswrodCheckInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswrodCheckInfo.Unmarshal(m, b)
}
func (m *PasswrodCheckInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswrodCheckInfo.Marshal(b, m, deterministic)
}
func (m *PasswrodCheckInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswrodCheckInfo.Merge(m, src)
}
func (m *PasswrodCheckInfo) XXX_Size() int {
	return xxx_messageInfo_PasswrodCheckInfo.Size(m)
}
func (m *PasswrodCheckInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswrodCheckInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PasswrodCheckInfo proto.InternalMessageInfo

func (m *PasswrodCheckInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PasswrodCheckInfo) GetEncryptedPassword() string {
	if m != nil {
		return m.EncryptedPassword
	}
	return ""
}

type CheckPasswordResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPasswordResponse) Reset()         { *m = CheckPasswordResponse{} }
func (m *CheckPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*CheckPasswordResponse) ProtoMessage()    {}
func (*CheckPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CheckPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPasswordResponse.Unmarshal(m, b)
}
func (m *CheckPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPasswordResponse.Marshal(b, m, deterministic)
}
func (m *CheckPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPasswordResponse.Merge(m, src)
}
func (m *CheckPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_CheckPasswordResponse.Size(m)
}
func (m *CheckPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPasswordResponse proto.InternalMessageInfo

func (m *CheckPasswordResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PageInfo struct {
	Pn                   uint32   `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	PSize                uint32   `protobuf:"varint,2,opt,name=pSize,proto3" json:"pSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *PageInfo) GetPSize() uint32 {
	if m != nil {
		return m.PSize
	}
	return 0
}

type UpdateUserInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	BirthDay             uint64   `protobuf:"varint,3,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfo) Reset()         { *m = UpdateUserInfo{} }
func (m *UpdateUserInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfo) ProtoMessage()    {}
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UpdateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfo.Unmarshal(m, b)
}
func (m *UpdateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfo.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfo.Merge(m, src)
}
func (m *UpdateUserInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfo.Size(m)
}
func (m *UpdateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfo proto.InternalMessageInfo

func (m *UpdateUserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UpdateUserInfo) GetBirthDay() uint64 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

type MobileRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileRequest) Reset()         { *m = MobileRequest{} }
func (m *MobileRequest) String() string { return proto.CompactTextString(m) }
func (*MobileRequest) ProtoMessage()    {}
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *MobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileRequest.Unmarshal(m, b)
}
func (m *MobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileRequest.Marshal(b, m, deterministic)
}
func (m *MobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileRequest.Merge(m, src)
}
func (m *MobileRequest) XXX_Size() int {
	return xxx_messageInfo_MobileRequest.Size(m)
}
func (m *MobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MobileRequest proto.InternalMessageInfo

func (m *MobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type IdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserListResponse struct {
	Total                int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *UserListResponse) GetData() []*UserInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserInfoResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PassWord             string   `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName             string   `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	BirthDay             uint64   `protobuf:"varint,5,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Role                 string   `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResponse) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *UserInfoResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfoResponse) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserInfoResponse) GetBirthDay() uint64 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

func (m *UserInfoResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type CreateUserInfo struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	PassWord             string   `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInfo) Reset()         { *m = CreateUserInfo{} }
func (m *CreateUserInfo) String() string { return proto.CompactTextString(m) }
func (*CreateUserInfo) ProtoMessage()    {}
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *CreateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInfo.Unmarshal(m, b)
}
func (m *CreateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInfo.Marshal(b, m, deterministic)
}
func (m *CreateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInfo.Merge(m, src)
}
func (m *CreateUserInfo) XXX_Size() int {
	return xxx_messageInfo_CreateUserInfo.Size(m)
}
func (m *CreateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInfo proto.InternalMessageInfo

func (m *CreateUserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CreateUserInfo) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *CreateUserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func init() {
	proto.RegisterType((*PasswrodCheckInfo)(nil), "PasswrodCheckInfo")
	proto.RegisterType((*CheckPasswordResponse)(nil), "CheckPasswordResponse")
	proto.RegisterType((*PageInfo)(nil), "PageInfo")
	proto.RegisterType((*UpdateUserInfo)(nil), "UpdateUserInfo")
	proto.RegisterType((*MobileRequest)(nil), "MobileRequest")
	proto.RegisterType((*IdRequest)(nil), "IdRequest")
	proto.RegisterType((*UserListResponse)(nil), "UserListResponse")
	proto.RegisterType((*UserInfoResponse)(nil), "UserInfoResponse")
	proto.RegisterType((*CreateUserInfo)(nil), "CreateUserInfo")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0x26, 0x69, 0xda, 0x6b, 0xe7, 0x68, 0x6b, 0x87, 0xb3, 0x84, 0xdc, 0x4b, 0x09, 0x88, 0x05,
	0x8f, 0xbd, 0xf3, 0xd4, 0xa7, 0x7b, 0xeb, 0x29, 0x52, 0xf0, 0x47, 0x89, 0x1c, 0x8a, 0x20, 0x98,
	0x26, 0x73, 0xbd, 0x70, 0x6d, 0x36, 0x66, 0xb7, 0x48, 0xfc, 0x73, 0xfc, 0x3f, 0xfc, 0xdf, 0x24,
	0x9b, 0x6c, 0x9a, 0xd8, 0x22, 0xdc, 0x53, 0xf2, 0xed, 0x7c, 0x3b, 0x33, 0xdf, 0x37, 0x3b, 0x00,
	0x5b, 0x41, 0x29, 0x4b, 0x52, 0x2e, 0xb9, 0x73, 0xba, 0xe2, 0x7c, 0xb5, 0xa6, 0x73, 0x85, 0x96,
	0xdb, 0xdb, 0x73, 0xda, 0x24, 0x32, 0x2b, 0x82, 0xee, 0x37, 0x18, 0x2d, 0x7c, 0x21, 0x7e, 0xa6,
	0x3c, 0xbc, 0xbe, 0xa3, 0xe0, 0x7e, 0x1e, 0xdf, 0x72, 0x74, 0xa0, 0x9b, 0xe4, 0x87, 0x3c, 0x0d,
	0x6d, 0x63, 0x62, 0x4c, 0x7b, 0x5e, 0x85, 0xf1, 0x0c, 0x46, 0x14, 0x07, 0x69, 0x96, 0x48, 0x0a,
	0x17, 0x9a, 0x64, 0x2a, 0xd2, 0x7e, 0xc0, 0x7d, 0x0e, 0x8f, 0x55, 0x5a, 0x7d, 0xe0, 0x91, 0x48,
	0x78, 0x2c, 0x08, 0x6d, 0x38, 0x12, 0xdb, 0x20, 0x20, 0x21, 0x54, 0x85, 0xae, 0xa7, 0xa1, 0x7b,
	0x01, 0xdd, 0x85, 0xbf, 0x22, 0xd5, 0xc8, 0x00, 0xcc, 0x24, 0x56, 0x84, 0xbe, 0x67, 0x26, 0x31,
	0x9e, 0x40, 0x3b, 0xf9, 0x14, 0xfd, 0x22, 0x55, 0xb0, 0xef, 0x15, 0xc0, 0xfd, 0x02, 0x83, 0x9b,
	0x24, 0xf4, 0x25, 0xdd, 0x08, 0x4a, 0xf5, 0xbd, 0xa8, 0x68, 0xbd, 0xed, 0x99, 0x51, 0x98, 0x0b,
	0x8a, 0xa3, 0xe0, 0xfe, 0x83, 0xbf, 0xa1, 0xb2, 0xd7, 0x0a, 0xe7, 0xb1, 0x65, 0x94, 0xca, 0xbb,
	0xd7, 0x7e, 0x66, 0xb7, 0x26, 0xc6, 0xd4, 0xf2, 0x2a, 0xec, 0x3e, 0x85, 0xfe, 0x7b, 0xbe, 0x8c,
	0xd6, 0xe4, 0xd1, 0x8f, 0x2d, 0x09, 0x89, 0x63, 0xe8, 0x6c, 0xd4, 0x41, 0xe9, 0x4b, 0x89, 0xdc,
	0x53, 0xe8, 0xcd, 0x43, 0x4d, 0xfa, 0xa7, 0xba, 0xfb, 0x11, 0x1e, 0xe5, 0x9d, 0xbd, 0x8b, 0x84,
	0xac, 0xf4, 0x9f, 0x40, 0x5b, 0x72, 0xe9, 0xaf, 0x4b, 0x5a, 0x01, 0xf0, 0x09, 0x58, 0xa1, 0x2f,
	0x7d, 0xdb, 0x9c, 0xb4, 0xa6, 0xc7, 0x97, 0x23, 0xa6, 0x05, 0xe9, 0x6b, 0x9e, 0x0a, 0xbb, 0xbf,
	0x8d, 0x22, 0x63, 0x3d, 0x74, 0x48, 0x73, 0x3e, 0xb4, 0xcf, 0xbb, 0xf9, 0x54, 0xb8, 0x26, 0xa3,
	0x55, 0x97, 0xd1, 0xf0, 0xc9, 0xfa, 0x8f, 0x4f, 0xed, 0xa6, 0x4f, 0x88, 0x60, 0xa5, 0x7c, 0x4d,
	0x76, 0x47, 0xdd, 0x51, 0xff, 0xee, 0x77, 0x18, 0x5c, 0xa7, 0x54, 0x9f, 0x4a, 0x3d, 0xbb, 0xb1,
	0x9f, 0xfd, 0xa1, 0xdd, 0x5e, 0xfe, 0x31, 0xc1, 0xca, 0x93, 0xe3, 0x33, 0x38, 0x7e, 0x4b, 0x52,
	0x7b, 0x8c, 0x3d, 0xa6, 0x1f, 0x90, 0x53, 0x58, 0xd8, 0x70, 0xfe, 0x25, 0x0c, 0x4b, 0xf2, 0x2c,
	0x2b, 0x86, 0x8b, 0x03, 0xd6, 0x98, 0xb2, 0xb3, 0x6f, 0x3c, 0x9e, 0x55, 0x25, 0x66, 0xd9, 0x3c,
	0x44, 0x60, 0xd5, 0xb8, 0x0f, 0xb1, 0x2f, 0x00, 0x76, 0xda, 0x71, 0xc8, 0x9a, 0x46, 0x1c, 0xba,
	0xf1, 0x0a, 0x60, 0xf7, 0x86, 0x71, 0xc8, 0x9a, 0x0f, 0xda, 0x19, 0xb3, 0x62, 0x89, 0x99, 0x5e,
	0x62, 0xf6, 0x26, 0x5f, 0x62, 0xbc, 0x82, 0x7e, 0x63, 0xbf, 0x10, 0xd9, 0xde, 0x3a, 0x3b, 0x63,
	0x76, 0x70, 0x07, 0x67, 0xbd, 0xaf, 0x47, 0xec, 0xaa, 0x48, 0xd8, 0x51, 0x9f, 0x17, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe3, 0x59, 0x42, 0x61, 0x38, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
	GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckPassword(ctx context.Context, in *PasswrodCheckInfo, opts ...grpc.CallOption) (*CheckPasswordResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassword(ctx context.Context, in *PasswrodCheckInfo, opts ...grpc.CallOption) (*CheckPasswordResponse, error) {
	out := new(CheckPasswordResponse)
	err := c.cc.Invoke(ctx, "/User/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*empty.Empty, error)
	CheckPassword(context.Context, *PasswrodCheckInfo) (*CheckPasswordResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserList(ctx context.Context, req *PageInfo) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedUserServer) GetUserByMobile(ctx context.Context, req *MobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (*UnimplementedUserServer) GetUserById(ctx context.Context, req *IdRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateUserInfo) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UpdateUserInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServer) CheckPassword(ctx context.Context, req *PasswrodCheckInfo) (*CheckPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswrodCheckInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassword(ctx, req.(*PasswrodCheckInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _User_GetUserByMobile_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _User_CheckPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
