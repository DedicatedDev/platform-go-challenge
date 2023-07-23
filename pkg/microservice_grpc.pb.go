// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MicroserviceClient is the client API for Microservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
	// Assets
	AddAsset(ctx context.Context, in *AddAssetRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
	GetAssets(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
	DeleteAsset(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
	// Favorites
	AddFavorites(ctx context.Context, in *AddFavoriteRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
	GetFavorites(ctx context.Context, in *GetFavoritesRequest, opts ...grpc.CallOption) (*NetWorkResponse, error)
}

type microserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceClient(cc grpc.ClientConnInterface) MicroserviceClient {
	return &microserviceClient{cc}
}

func (c *microserviceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) AddAsset(ctx context.Context, in *AddAssetRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/AddAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) GetAssets(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/GetAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) DeleteAsset(ctx context.Context, in *DeleteAssetRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/DeleteAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) AddFavorites(ctx context.Context, in *AddFavoriteRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/AddFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceClient) GetFavorites(ctx context.Context, in *GetFavoritesRequest, opts ...grpc.CallOption) (*NetWorkResponse, error) {
	out := new(NetWorkResponse)
	err := c.cc.Invoke(ctx, "/gwi_api.v1.Microservice/GetFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceServer is the server API for Microservice service.
// All implementations must embed UnimplementedMicroserviceServer
// for forward compatibility
type MicroserviceServer interface {
	SignUp(context.Context, *SignUpRequest) (*NetWorkResponse, error)
	SignIn(context.Context, *SignInRequest) (*NetWorkResponse, error)
	// Assets
	AddAsset(context.Context, *AddAssetRequest) (*NetWorkResponse, error)
	GetAssets(context.Context, *GetAssetRequest) (*NetWorkResponse, error)
	DeleteAsset(context.Context, *DeleteAssetRequest) (*NetWorkResponse, error)
	// Favorites
	AddFavorites(context.Context, *AddFavoriteRequest) (*NetWorkResponse, error)
	GetFavorites(context.Context, *GetFavoritesRequest) (*NetWorkResponse, error)
	mustEmbedUnimplementedMicroserviceServer()
}

// UnimplementedMicroserviceServer must be embedded to have forward compatible implementations.
type UnimplementedMicroserviceServer struct {
}

func (UnimplementedMicroserviceServer) SignUp(context.Context, *SignUpRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedMicroserviceServer) SignIn(context.Context, *SignInRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedMicroserviceServer) AddAsset(context.Context, *AddAssetRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAsset not implemented")
}
func (UnimplementedMicroserviceServer) GetAssets(context.Context, *GetAssetRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssets not implemented")
}
func (UnimplementedMicroserviceServer) DeleteAsset(context.Context, *DeleteAssetRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAsset not implemented")
}
func (UnimplementedMicroserviceServer) AddFavorites(context.Context, *AddFavoriteRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFavorites not implemented")
}
func (UnimplementedMicroserviceServer) GetFavorites(context.Context, *GetFavoritesRequest) (*NetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorites not implemented")
}
func (UnimplementedMicroserviceServer) mustEmbedUnimplementedMicroserviceServer() {}

// UnsafeMicroserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceServer will
// result in compilation errors.
type UnsafeMicroserviceServer interface {
	mustEmbedUnimplementedMicroserviceServer()
}

func RegisterMicroserviceServer(s grpc.ServiceRegistrar, srv MicroserviceServer) {
	s.RegisterService(&Microservice_ServiceDesc, srv)
}

func _Microservice_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_AddAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).AddAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/AddAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).AddAsset(ctx, req.(*AddAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_GetAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).GetAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/GetAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).GetAssets(ctx, req.(*GetAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_DeleteAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).DeleteAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/DeleteAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).DeleteAsset(ctx, req.(*DeleteAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_AddFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).AddFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/AddFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).AddFavorites(ctx, req.(*AddFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microservice_GetFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).GetFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwi_api.v1.Microservice/GetFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).GetFavorites(ctx, req.(*GetFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Microservice_ServiceDesc is the grpc.ServiceDesc for Microservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Microservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gwi_api.v1.Microservice",
	HandlerType: (*MicroserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Microservice_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Microservice_SignIn_Handler,
		},
		{
			MethodName: "AddAsset",
			Handler:    _Microservice_AddAsset_Handler,
		},
		{
			MethodName: "GetAssets",
			Handler:    _Microservice_GetAssets_Handler,
		},
		{
			MethodName: "DeleteAsset",
			Handler:    _Microservice_DeleteAsset_Handler,
		},
		{
			MethodName: "AddFavorites",
			Handler:    _Microservice_AddFavorites_Handler,
		},
		{
			MethodName: "GetFavorites",
			Handler:    _Microservice_GetFavorites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservice.proto",
}
