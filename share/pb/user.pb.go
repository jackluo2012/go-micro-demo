// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	user.proto
	userStatus.proto

It has these top-level messages:
	User
	ModifyUserReq
	ModifyUserRep
	SelectUserReq
	SelectUserRep
	DeletetUserReq
	DeletetUserRep
	InsertUserReq
	InsertUserRep
	Session
	GetSessionByUIDReq
	GetSessionByUIDRep
	GetSessionByTokenReq
	UpdateConnectorAddrReq
	GetSessionByTokenRep
	GetConnectorAddrReq
	GetConnectorAddrRep
	UpdateConnectorAddrRep
	NewSessionReq
	NewSessionRep
	RemoveSessionReq
	RemoveSessionRep
	RefreshSessionReq
	RefreshSessionRep
	UserConnectedReq
	UserConnectedRep
	UserDisonnectedReq
	UserDisonnectedRep
	RemoveSessionByUIDReq
	RemoveSessionByUIDRep
	GetUserIDByTokenReq
	GetUserIDByTokenRep
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ModifyUserReq struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
}

func (m *ModifyUserReq) Reset()                    { *m = ModifyUserReq{} }
func (m *ModifyUserReq) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserReq) ProtoMessage()               {}
func (*ModifyUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ModifyUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModifyUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModifyUserReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ModifyUserReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ModifyUserRep struct {
}

func (m *ModifyUserRep) Reset()                    { *m = ModifyUserRep{} }
func (m *ModifyUserRep) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserRep) ProtoMessage()               {}
func (*ModifyUserRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SelectUserReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *SelectUserReq) Reset()                    { *m = SelectUserReq{} }
func (m *SelectUserReq) String() string            { return proto.CompactTextString(m) }
func (*SelectUserReq) ProtoMessage()               {}
func (*SelectUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SelectUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SelectUserRep struct {
	Users *User `protobuf:"bytes,1,opt,name=users" json:"users,omitempty"`
}

func (m *SelectUserRep) Reset()                    { *m = SelectUserRep{} }
func (m *SelectUserRep) String() string            { return proto.CompactTextString(m) }
func (*SelectUserRep) ProtoMessage()               {}
func (*SelectUserRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SelectUserRep) GetUsers() *User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeletetUserReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeletetUserReq) Reset()                    { *m = DeletetUserReq{} }
func (m *DeletetUserReq) String() string            { return proto.CompactTextString(m) }
func (*DeletetUserReq) ProtoMessage()               {}
func (*DeletetUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeletetUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeletetUserRep struct {
}

func (m *DeletetUserRep) Reset()                    { *m = DeletetUserRep{} }
func (m *DeletetUserRep) String() string            { return proto.CompactTextString(m) }
func (*DeletetUserRep) ProtoMessage()               {}
func (*DeletetUserRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type InsertUserReq struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
}

func (m *InsertUserReq) Reset()                    { *m = InsertUserReq{} }
func (m *InsertUserReq) String() string            { return proto.CompactTextString(m) }
func (*InsertUserReq) ProtoMessage()               {}
func (*InsertUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *InsertUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertUserReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InsertUserReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type InsertUserRep struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
}

func (m *InsertUserRep) Reset()                    { *m = InsertUserRep{} }
func (m *InsertUserRep) String() string            { return proto.CompactTextString(m) }
func (*InsertUserRep) ProtoMessage()               {}
func (*InsertUserRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *InsertUserRep) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertUserRep) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertUserRep) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InsertUserRep) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*ModifyUserReq)(nil), "pb.ModifyUserReq")
	proto.RegisterType((*ModifyUserRep)(nil), "pb.ModifyUserRep")
	proto.RegisterType((*SelectUserReq)(nil), "pb.SelectUserReq")
	proto.RegisterType((*SelectUserRep)(nil), "pb.SelectUserRep")
	proto.RegisterType((*DeletetUserReq)(nil), "pb.DeletetUserReq")
	proto.RegisterType((*DeletetUserRep)(nil), "pb.DeletetUserRep")
	proto.RegisterType((*InsertUserReq)(nil), "pb.InsertUserReq")
	proto.RegisterType((*InsertUserRep)(nil), "pb.InsertUserRep")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	// 增
	InsertUser(ctx context.Context, in *InsertUserReq, opts ...client.CallOption) (*InsertUserRep, error)
	// 删
	DeletetUser(ctx context.Context, in *DeletetUserReq, opts ...client.CallOption) (*DeletetUserRep, error)
	// 查
	SelectUser(ctx context.Context, in *SelectUserReq, opts ...client.CallOption) (*SelectUserRep, error)
	// 改
	ModifyUser(ctx context.Context, in *ModifyUserReq, opts ...client.CallOption) (*ModifyUserRep, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) InsertUser(ctx context.Context, in *InsertUserReq, opts ...client.CallOption) (*InsertUserRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.InsertUser", in)
	out := new(InsertUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeletetUser(ctx context.Context, in *DeletetUserReq, opts ...client.CallOption) (*DeletetUserRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.DeletetUser", in)
	out := new(DeletetUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SelectUser(ctx context.Context, in *SelectUserReq, opts ...client.CallOption) (*SelectUserRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.SelectUser", in)
	out := new(SelectUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ModifyUser(ctx context.Context, in *ModifyUserReq, opts ...client.CallOption) (*ModifyUserRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ModifyUser", in)
	out := new(ModifyUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// 增
	InsertUser(context.Context, *InsertUserReq, *InsertUserRep) error
	// 删
	DeletetUser(context.Context, *DeletetUserReq, *DeletetUserRep) error
	// 查
	SelectUser(context.Context, *SelectUserReq, *SelectUserRep) error
	// 改
	ModifyUser(context.Context, *ModifyUserReq, *ModifyUserRep) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) InsertUser(ctx context.Context, in *InsertUserReq, out *InsertUserRep) error {
	return h.UserServiceHandler.InsertUser(ctx, in, out)
}

func (h *UserService) DeletetUser(ctx context.Context, in *DeletetUserReq, out *DeletetUserRep) error {
	return h.UserServiceHandler.DeletetUser(ctx, in, out)
}

func (h *UserService) SelectUser(ctx context.Context, in *SelectUserReq, out *SelectUserRep) error {
	return h.UserServiceHandler.SelectUser(ctx, in, out)
}

func (h *UserService) ModifyUser(ctx context.Context, in *ModifyUserReq, out *ModifyUserRep) error {
	return h.UserServiceHandler.ModifyUser(ctx, in, out)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x2f, 0xf1, 0xe2, 0x9f, 0x09, 0x77, 0xea, 0x60, 0xb1, 0xa4, 0xd0, 0xb0, 0xd5, 0x55,
	0x11, 0xce, 0x03, 0x6b, 0xc1, 0xc6, 0x42, 0x90, 0x1c, 0x36, 0x76, 0x26, 0x3b, 0x62, 0xe0, 0x4c,
	0xc6, 0xdd, 0x28, 0xf8, 0x8d, 0xfd, 0x18, 0xb2, 0x1b, 0x25, 0xee, 0x85, 0x2b, 0xd3, 0x65, 0x5f,
	0xde, 0xec, 0x8f, 0x79, 0x33, 0x0b, 0xf0, 0x61, 0x48, 0x67, 0xac, 0x9b, 0xb6, 0xc1, 0x90, 0x0b,
	0xf9, 0x04, 0xd3, 0x47, 0x43, 0x1a, 0xe7, 0x10, 0x56, 0x4a, 0x04, 0x69, 0xb0, 0x88, 0xf2, 0xb0,
	0x52, 0x88, 0x30, 0xad, 0x9f, 0xdf, 0x48, 0x84, 0x69, 0xb0, 0x38, 0xca, 0xdd, 0x37, 0x0a, 0x38,
	0xb8, 0x51, 0x4a, 0x93, 0x31, 0x62, 0xcf, 0xc9, 0x7f, 0x47, 0x3c, 0x83, 0xe8, 0xe1, 0xb5, 0xa9,
	0x49, 0x4c, 0x9d, 0xde, 0x1d, 0x64, 0x09, 0xb3, 0xfb, 0x46, 0x55, 0x2f, 0x5f, 0x96, 0x90, 0xd3,
	0xfb, 0x28, 0x90, 0x63, 0x1f, 0xc2, 0xf2, 0x02, 0x66, 0x6b, 0xda, 0x50, 0xd9, 0xee, 0xa0, 0xca,
	0x4b, 0xdf, 0xc0, 0x78, 0x0e, 0x91, 0x4d, 0xc5, 0x38, 0x4f, 0xbc, 0x3c, 0xcc, 0xb8, 0xc8, 0xdc,
	0xbf, 0x4e, 0x96, 0x29, 0xcc, 0x6f, 0x69, 0x43, 0x2d, 0xed, 0xbc, 0xf2, 0x64, 0xcb, 0xc1, 0xb6,
	0xf7, 0xbb, 0xda, 0x90, 0x6e, 0xc7, 0xec, 0x7d, 0x0b, 0xc2, 0x63, 0x40, 0x96, 0xdf, 0x01, 0xc4,
	0xf6, 0xfe, 0x35, 0xe9, 0xcf, 0xaa, 0x24, 0x5c, 0x01, 0xf4, 0x50, 0x3c, 0xb5, 0x61, 0x79, 0x9d,
	0x26, 0x03, 0x89, 0xe5, 0x04, 0xaf, 0x21, 0xfe, 0x97, 0x10, 0xa2, 0xf5, 0xf8, 0xa1, 0x26, 0x43,
	0xcd, 0x16, 0xae, 0x00, 0xfa, 0x69, 0x75, 0x38, 0x6f, 0xbc, 0xc9, 0x40, 0xfa, 0xad, 0xea, 0xb7,
	0xa2, 0xab, 0xf2, 0x56, 0x31, 0x19, 0x48, 0x2c, 0x27, 0xc5, 0xbe, 0x7b, 0x17, 0x57, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xda, 0x06, 0x39, 0x51, 0x25, 0x03, 0x00, 0x00,
}
