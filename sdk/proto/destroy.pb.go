// Code generated by protoc-gen-go. DO NOT EDIT.
// source: destroy.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DestroyClient is the client API for Destroy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DestroyClient interface {
	DestroySpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Destroy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
}

type destroyClient struct {
	cc *grpc.ClientConn
}

func NewDestroyClient(cc *grpc.ClientConn) DestroyClient {
	return &destroyClient{cc}
}

func (c *destroyClient) DestroySpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Destroy/DestroySpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destroyClient) Destroy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Destroy/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DestroyServer is the server API for Destroy service.
type DestroyServer interface {
	DestroySpec(context.Context, *empty.Empty) (*FuncSpec, error)
	Destroy(context.Context, *FuncSpec_Args) (*empty.Empty, error)
}

func RegisterDestroyServer(s *grpc.Server, srv DestroyServer) {
	s.RegisterService(&_Destroy_serviceDesc, srv)
}

func _Destroy_DestroySpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestroyServer).DestroySpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destroy/DestroySpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestroyServer).DestroySpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destroy_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestroyServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destroy/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestroyServer).Destroy(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _Destroy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Destroy",
	HandlerType: (*DestroyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DestroySpec",
			Handler:    _Destroy_DestroySpec_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Destroy_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "destroy.proto",
}

func init() { proto.RegisterFile("destroy.proto", fileDescriptor_destroy_966cd8d70e071941) }

var fileDescriptor_destroy_966cd8d70e071941 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2d, 0x2e,
	0x29, 0xca, 0xaf, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xd2, 0xe9,
	0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x09,
	0x54, 0x8d, 0x14, 0x4f, 0x41, 0x4e, 0x69, 0x7a, 0x66, 0x1e, 0x84, 0x67, 0x54, 0xc5, 0xc5, 0xee,
	0x02, 0x31, 0x42, 0xc8, 0x8c, 0x8b, 0x1b, 0xca, 0x0c, 0x2e, 0x48, 0x4d, 0x16, 0x12, 0xd3, 0x83,
	0x98, 0xa2, 0x07, 0x33, 0x45, 0xcf, 0x15, 0x64, 0x8a, 0x14, 0x3f, 0x44, 0x40, 0xcf, 0xad, 0x34,
	0x2f, 0x19, 0xac, 0xd0, 0x1c, 0x61, 0x84, 0x08, 0x9a, 0x9c, 0x9e, 0x63, 0x51, 0x7a, 0xb1, 0x14,
	0x0e, 0x93, 0x92, 0xd8, 0xc0, 0x7c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x07, 0x6a,
	0x43, 0xc5, 0x00, 0x00, 0x00,
}