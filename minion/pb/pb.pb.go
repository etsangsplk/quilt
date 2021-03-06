// Code generated by protoc-gen-go.
// source: minion/pb/pb.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	minion/pb/pb.proto

It has these top-level messages:
	MinionConfig
	Reply
	Request
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type MinionConfig_Role int32

const (
	MinionConfig_NONE   MinionConfig_Role = 0
	MinionConfig_WORKER MinionConfig_Role = 1
	MinionConfig_MASTER MinionConfig_Role = 2
)

var MinionConfig_Role_name = map[int32]string{
	0: "NONE",
	1: "WORKER",
	2: "MASTER",
}
var MinionConfig_Role_value = map[string]int32{
	"NONE":   0,
	"WORKER": 1,
	"MASTER": 2,
}

func (x MinionConfig_Role) String() string {
	return proto.EnumName(MinionConfig_Role_name, int32(x))
}
func (MinionConfig_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type MinionConfig struct {
	ID             string            `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Role           MinionConfig_Role `protobuf:"varint,2,opt,name=role,enum=MinionConfig_Role" json:"role,omitempty"`
	PrivateIP      string            `protobuf:"bytes,3,opt,name=PrivateIP,json=privateIP" json:"PrivateIP,omitempty"`
	Spec           string            `protobuf:"bytes,4,opt,name=Spec,json=spec" json:"Spec,omitempty"`
	Provider       string            `protobuf:"bytes,5,opt,name=Provider,json=provider" json:"Provider,omitempty"`
	Size           string            `protobuf:"bytes,6,opt,name=Size,json=size" json:"Size,omitempty"`
	Region         string            `protobuf:"bytes,7,opt,name=Region,json=region" json:"Region,omitempty"`
	FloatingIP     string            `protobuf:"bytes,8,opt,name=FloatingIP,json=floatingIP" json:"FloatingIP,omitempty"`
	EtcdMembers    []string          `protobuf:"bytes,9,rep,name=EtcdMembers,json=etcdMembers" json:"EtcdMembers,omitempty"`
	AuthorizedKeys []string          `protobuf:"bytes,10,rep,name=AuthorizedKeys,json=authorizedKeys" json:"AuthorizedKeys,omitempty"`
}

func (m *MinionConfig) Reset()                    { *m = MinionConfig{} }
func (m *MinionConfig) String() string            { return proto.CompactTextString(m) }
func (*MinionConfig) ProtoMessage()               {}
func (*MinionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MinionConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MinionConfig) GetRole() MinionConfig_Role {
	if m != nil {
		return m.Role
	}
	return MinionConfig_NONE
}

func (m *MinionConfig) GetPrivateIP() string {
	if m != nil {
		return m.PrivateIP
	}
	return ""
}

func (m *MinionConfig) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *MinionConfig) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *MinionConfig) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *MinionConfig) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *MinionConfig) GetFloatingIP() string {
	if m != nil {
		return m.FloatingIP
	}
	return ""
}

func (m *MinionConfig) GetEtcdMembers() []string {
	if m != nil {
		return m.EtcdMembers
	}
	return nil
}

func (m *MinionConfig) GetAuthorizedKeys() []string {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

type Reply struct {
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*MinionConfig)(nil), "MinionConfig")
	proto.RegisterType((*Reply)(nil), "Reply")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterEnum("MinionConfig_Role", MinionConfig_Role_name, MinionConfig_Role_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Minion service

type MinionClient interface {
	SetMinionConfig(ctx context.Context, in *MinionConfig, opts ...grpc.CallOption) (*Reply, error)
	GetMinionConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*MinionConfig, error)
}

type minionClient struct {
	cc *grpc.ClientConn
}

func NewMinionClient(cc *grpc.ClientConn) MinionClient {
	return &minionClient{cc}
}

func (c *minionClient) SetMinionConfig(ctx context.Context, in *MinionConfig, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/Minion/SetMinionConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) GetMinionConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*MinionConfig, error) {
	out := new(MinionConfig)
	err := grpc.Invoke(ctx, "/Minion/GetMinionConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Minion service

type MinionServer interface {
	SetMinionConfig(context.Context, *MinionConfig) (*Reply, error)
	GetMinionConfig(context.Context, *Request) (*MinionConfig, error)
}

func RegisterMinionServer(s *grpc.Server, srv MinionServer) {
	s.RegisterService(&_Minion_serviceDesc, srv)
}

func _Minion_SetMinionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).SetMinionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Minion/SetMinionConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).SetMinionConfig(ctx, req.(*MinionConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minion_GetMinionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).GetMinionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Minion/GetMinionConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).GetMinionConfig(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Minion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Minion",
	HandlerType: (*MinionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMinionConfig",
			Handler:    _Minion_SetMinionConfig_Handler,
		},
		{
			MethodName: "GetMinionConfig",
			Handler:    _Minion_GetMinionConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minion/pb/pb.proto",
}

func init() { proto.RegisterFile("minion/pb/pb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0x51, 0xcb, 0xda, 0x30,
	0x14, 0x86, 0x6d, 0xad, 0xb5, 0x3d, 0x6e, 0x55, 0xce, 0xc5, 0x08, 0x32, 0x46, 0xe9, 0x85, 0x94,
	0x31, 0x2a, 0xb8, 0x5f, 0x20, 0xb3, 0x1b, 0x22, 0x6a, 0x89, 0x83, 0x5d, 0x5b, 0x3d, 0xba, 0x40,
	0x6d, 0xb2, 0xb4, 0x0a, 0xfa, 0x3b, 0xf7, 0x83, 0x86, 0xd1, 0x6d, 0x5f, 0xbf, 0xbb, 0x9c, 0xe7,
	0x3c, 0x6f, 0x20, 0x6f, 0x00, 0x4f, 0xa2, 0x14, 0xb2, 0x1c, 0xab, 0x7c, 0xac, 0xf2, 0x44, 0x69,
	0x59, 0xcb, 0xe8, 0xb7, 0x0d, 0x6f, 0x96, 0x06, 0x7f, 0x91, 0xe5, 0x41, 0x1c, 0x31, 0x00, 0x7b,
	0x3e, 0x63, 0x56, 0x68, 0xc5, 0x3e, 0xb7, 0xc5, 0x0c, 0x47, 0xe0, 0x68, 0x59, 0x10, 0xb3, 0x43,
	0x2b, 0x0e, 0x26, 0x98, 0xbc, 0x94, 0x13, 0x2e, 0x0b, 0xe2, 0x66, 0x8f, 0xef, 0xc1, 0xcf, 0xb4,
	0xb8, 0x6c, 0x6b, 0x9a, 0x67, 0xac, 0x6d, 0xe2, 0xbe, 0xfa, 0x0b, 0x10, 0xc1, 0xd9, 0x28, 0xda,
	0x31, 0xc7, 0x2c, 0x9c, 0x4a, 0xd1, 0x0e, 0x87, 0xe0, 0x65, 0x5a, 0x5e, 0xc4, 0x9e, 0x34, 0xeb,
	0x18, 0xee, 0xa9, 0xe7, 0x6c, 0x7c, 0x71, 0x23, 0xe6, 0x3e, 0x7d, 0x71, 0x23, 0x7c, 0x07, 0x2e,
	0xa7, 0xa3, 0x90, 0x25, 0xeb, 0x1a, 0xea, 0x6a, 0x33, 0xe1, 0x07, 0x80, 0xaf, 0x85, 0xdc, 0xd6,
	0xa2, 0x3c, 0xce, 0x33, 0xe6, 0x99, 0x1d, 0x1c, 0xfe, 0x11, 0x0c, 0xa1, 0x97, 0xd6, 0xbb, 0xfd,
	0x92, 0x4e, 0x39, 0xe9, 0x8a, 0xf9, 0x61, 0x3b, 0xf6, 0x79, 0x8f, 0xfe, 0x23, 0x1c, 0x41, 0x30,
	0x3d, 0xd7, 0x3f, 0xa5, 0x16, 0x37, 0xda, 0x2f, 0xe8, 0x5a, 0x31, 0x30, 0x52, 0xb0, 0x6d, 0xd0,
	0x28, 0x06, 0xe7, 0xfe, 0x62, 0xf4, 0xc0, 0x59, 0xad, 0x57, 0xe9, 0xa0, 0x85, 0x00, 0xee, 0x8f,
	0x35, 0x5f, 0xa4, 0x7c, 0x60, 0xdd, 0xcf, 0xcb, 0xe9, 0xe6, 0x7b, 0xca, 0x07, 0x76, 0xd4, 0x85,
	0x0e, 0x27, 0x55, 0x5c, 0x23, 0x1f, 0xba, 0x9c, 0x7e, 0x9d, 0xa9, 0xaa, 0x27, 0x39, 0xb8, 0x8f,
	0xf2, 0xf0, 0x23, 0xf4, 0x37, 0x54, 0x37, 0x6a, 0x7f, 0xdb, 0x28, 0x76, 0xe8, 0x26, 0x8f, 0x78,
	0x0b, 0x3f, 0x41, 0xff, 0xdb, 0x2b, 0xd7, 0x4b, 0x9e, 0x57, 0x0e, 0x9b, 0xa9, 0xa8, 0x95, 0xbb,
	0xe6, 0x57, 0x3f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xab, 0x2c, 0xe5, 0x31, 0xeb, 0x01, 0x00,
	0x00,
}
