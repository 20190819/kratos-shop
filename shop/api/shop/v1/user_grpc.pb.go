// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: shop/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShopClient is the client API for Shop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*RegisterReply, error)
	Captcha(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CaptchaReply, error)
	Detail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DetailReply, error)
}

type shopClient struct {
	cc grpc.ClientConnInterface
}

func NewShopClient(cc grpc.ClientConnInterface) ShopClient {
	return &shopClient{cc}
}

func (c *shopClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.shop.v1.Shop/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.shop.v1.Shop/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Captcha(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CaptchaReply, error) {
	out := new(CaptchaReply)
	err := c.cc.Invoke(ctx, "/api.shop.v1.Shop/Captcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Detail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DetailReply, error) {
	out := new(DetailReply)
	err := c.cc.Invoke(ctx, "/api.shop.v1.Shop/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServer is the server API for Shop service.
// All implementations must embed UnimplementedShopServer
// for forward compatibility
type ShopServer interface {
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	Login(context.Context, *LoginReq) (*RegisterReply, error)
	Captcha(context.Context, *emptypb.Empty) (*CaptchaReply, error)
	Detail(context.Context, *emptypb.Empty) (*DetailReply, error)
	mustEmbedUnimplementedShopServer()
}

// UnimplementedShopServer must be embedded to have forward compatible implementations.
type UnimplementedShopServer struct {
}

func (UnimplementedShopServer) Register(context.Context, *RegisterReq) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedShopServer) Login(context.Context, *LoginReq) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedShopServer) Captcha(context.Context, *emptypb.Empty) (*CaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Captcha not implemented")
}
func (UnimplementedShopServer) Detail(context.Context, *emptypb.Empty) (*DetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedShopServer) mustEmbedUnimplementedShopServer() {}

// UnsafeShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServer will
// result in compilation errors.
type UnsafeShopServer interface {
	mustEmbedUnimplementedShopServer()
}

func RegisterShopServer(s grpc.ServiceRegistrar, srv ShopServer) {
	s.RegisterService(&Shop_ServiceDesc, srv)
}

func _Shop_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.v1.Shop/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.v1.Shop/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_Captcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Captcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.v1.Shop/Captcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Captcha(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.v1.Shop/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Detail(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Shop_ServiceDesc is the grpc.ServiceDesc for Shop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.shop.v1.Shop",
	HandlerType: (*ShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Shop_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Shop_Login_Handler,
		},
		{
			MethodName: "Captcha",
			Handler:    _Shop_Captcha_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Shop_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/v1/user.proto",
}
