// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: upload/upload.proto

package upload

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UploadAlbumsService_CreateAlbum_FullMethodName      = "/proto.UploadAlbumsService/CreateAlbum"
	UploadAlbumsService_UpdateAlbum_FullMethodName      = "/proto.UploadAlbumsService/UpdateAlbum"
	UploadAlbumsService_DeleteAlbum_FullMethodName      = "/proto.UploadAlbumsService/DeleteAlbum"
	UploadAlbumsService_AddToAlbumArtist_FullMethodName = "/proto.UploadAlbumsService/AddToAlbumArtist"
	UploadAlbumsService_AddToAlbumGenre_FullMethodName  = "/proto.UploadAlbumsService/AddToAlbumGenre"
)

// UploadAlbumsServiceClient is the client API for UploadAlbumsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadAlbumsServiceClient interface {
	CreateAlbum(ctx context.Context, in *CreateAlbumReq, opts ...grpc.CallOption) (*CreateAlbumRes, error)
	UpdateAlbum(ctx context.Context, in *UpdateAlbumReq, opts ...grpc.CallOption) (*UpdateAlbumRes, error)
	DeleteAlbum(ctx context.Context, in *DeleteAlbumReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddToAlbumArtist(ctx context.Context, in *AlbumsArtistsReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddToAlbumGenre(ctx context.Context, in *AlbumsGenresReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type uploadAlbumsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadAlbumsServiceClient(cc grpc.ClientConnInterface) UploadAlbumsServiceClient {
	return &uploadAlbumsServiceClient{cc}
}

func (c *uploadAlbumsServiceClient) CreateAlbum(ctx context.Context, in *CreateAlbumReq, opts ...grpc.CallOption) (*CreateAlbumRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAlbumRes)
	err := c.cc.Invoke(ctx, UploadAlbumsService_CreateAlbum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadAlbumsServiceClient) UpdateAlbum(ctx context.Context, in *UpdateAlbumReq, opts ...grpc.CallOption) (*UpdateAlbumRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAlbumRes)
	err := c.cc.Invoke(ctx, UploadAlbumsService_UpdateAlbum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadAlbumsServiceClient) DeleteAlbum(ctx context.Context, in *DeleteAlbumReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadAlbumsService_DeleteAlbum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadAlbumsServiceClient) AddToAlbumArtist(ctx context.Context, in *AlbumsArtistsReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadAlbumsService_AddToAlbumArtist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadAlbumsServiceClient) AddToAlbumGenre(ctx context.Context, in *AlbumsGenresReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadAlbumsService_AddToAlbumGenre_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadAlbumsServiceServer is the server API for UploadAlbumsService service.
// All implementations must embed UnimplementedUploadAlbumsServiceServer
// for forward compatibility.
type UploadAlbumsServiceServer interface {
	CreateAlbum(context.Context, *CreateAlbumReq) (*CreateAlbumRes, error)
	UpdateAlbum(context.Context, *UpdateAlbumReq) (*UpdateAlbumRes, error)
	DeleteAlbum(context.Context, *DeleteAlbumReq) (*emptypb.Empty, error)
	AddToAlbumArtist(context.Context, *AlbumsArtistsReq) (*emptypb.Empty, error)
	AddToAlbumGenre(context.Context, *AlbumsGenresReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedUploadAlbumsServiceServer()
}

// UnimplementedUploadAlbumsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadAlbumsServiceServer struct{}

func (UnimplementedUploadAlbumsServiceServer) CreateAlbum(context.Context, *CreateAlbumReq) (*CreateAlbumRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlbum not implemented")
}
func (UnimplementedUploadAlbumsServiceServer) UpdateAlbum(context.Context, *UpdateAlbumReq) (*UpdateAlbumRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlbum not implemented")
}
func (UnimplementedUploadAlbumsServiceServer) DeleteAlbum(context.Context, *DeleteAlbumReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlbum not implemented")
}
func (UnimplementedUploadAlbumsServiceServer) AddToAlbumArtist(context.Context, *AlbumsArtistsReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToAlbumArtist not implemented")
}
func (UnimplementedUploadAlbumsServiceServer) AddToAlbumGenre(context.Context, *AlbumsGenresReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToAlbumGenre not implemented")
}
func (UnimplementedUploadAlbumsServiceServer) mustEmbedUnimplementedUploadAlbumsServiceServer() {}
func (UnimplementedUploadAlbumsServiceServer) testEmbeddedByValue()                             {}

// UnsafeUploadAlbumsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadAlbumsServiceServer will
// result in compilation errors.
type UnsafeUploadAlbumsServiceServer interface {
	mustEmbedUnimplementedUploadAlbumsServiceServer()
}

func RegisterUploadAlbumsServiceServer(s grpc.ServiceRegistrar, srv UploadAlbumsServiceServer) {
	// If the following call pancis, it indicates UnimplementedUploadAlbumsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UploadAlbumsService_ServiceDesc, srv)
}

func _UploadAlbumsService_CreateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlbumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadAlbumsServiceServer).CreateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadAlbumsService_CreateAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadAlbumsServiceServer).CreateAlbum(ctx, req.(*CreateAlbumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadAlbumsService_UpdateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlbumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadAlbumsServiceServer).UpdateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadAlbumsService_UpdateAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadAlbumsServiceServer).UpdateAlbum(ctx, req.(*UpdateAlbumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadAlbumsService_DeleteAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlbumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadAlbumsServiceServer).DeleteAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadAlbumsService_DeleteAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadAlbumsServiceServer).DeleteAlbum(ctx, req.(*DeleteAlbumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadAlbumsService_AddToAlbumArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumsArtistsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadAlbumsServiceServer).AddToAlbumArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadAlbumsService_AddToAlbumArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadAlbumsServiceServer).AddToAlbumArtist(ctx, req.(*AlbumsArtistsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadAlbumsService_AddToAlbumGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumsGenresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadAlbumsServiceServer).AddToAlbumGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadAlbumsService_AddToAlbumGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadAlbumsServiceServer).AddToAlbumGenre(ctx, req.(*AlbumsGenresReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadAlbumsService_ServiceDesc is the grpc.ServiceDesc for UploadAlbumsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadAlbumsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploadAlbumsService",
	HandlerType: (*UploadAlbumsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlbum",
			Handler:    _UploadAlbumsService_CreateAlbum_Handler,
		},
		{
			MethodName: "UpdateAlbum",
			Handler:    _UploadAlbumsService_UpdateAlbum_Handler,
		},
		{
			MethodName: "DeleteAlbum",
			Handler:    _UploadAlbumsService_DeleteAlbum_Handler,
		},
		{
			MethodName: "AddToAlbumArtist",
			Handler:    _UploadAlbumsService_AddToAlbumArtist_Handler,
		},
		{
			MethodName: "AddToAlbumGenre",
			Handler:    _UploadAlbumsService_AddToAlbumGenre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload/upload.proto",
}

const (
	UploadArtistService_CreateArtist_FullMethodName     = "/proto.UploadArtistService/CreateArtist"
	UploadArtistService_UpdateArtist_FullMethodName     = "/proto.UploadArtistService/UpdateArtist"
	UploadArtistService_DeleteArtist_FullMethodName     = "/proto.UploadArtistService/DeleteArtist"
	UploadArtistService_AddToArtistTrack_FullMethodName = "/proto.UploadArtistService/AddToArtistTrack"
	UploadArtistService_AddToArtistGenre_FullMethodName = "/proto.UploadArtistService/AddToArtistGenre"
)

// UploadArtistServiceClient is the client API for UploadArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadArtistServiceClient interface {
	CreateArtist(ctx context.Context, in *CreateArtistReq, opts ...grpc.CallOption) (*CreateArtistRes, error)
	UpdateArtist(ctx context.Context, in *UpdateArtistReq, opts ...grpc.CallOption) (*UpdateArtistRes, error)
	DeleteArtist(ctx context.Context, in *DeleteArtistReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddToArtistTrack(ctx context.Context, in *ArtistsTracksReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddToArtistGenre(ctx context.Context, in *ArtistsGenresReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type uploadArtistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadArtistServiceClient(cc grpc.ClientConnInterface) UploadArtistServiceClient {
	return &uploadArtistServiceClient{cc}
}

func (c *uploadArtistServiceClient) CreateArtist(ctx context.Context, in *CreateArtistReq, opts ...grpc.CallOption) (*CreateArtistRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateArtistRes)
	err := c.cc.Invoke(ctx, UploadArtistService_CreateArtist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadArtistServiceClient) UpdateArtist(ctx context.Context, in *UpdateArtistReq, opts ...grpc.CallOption) (*UpdateArtistRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateArtistRes)
	err := c.cc.Invoke(ctx, UploadArtistService_UpdateArtist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadArtistServiceClient) DeleteArtist(ctx context.Context, in *DeleteArtistReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadArtistService_DeleteArtist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadArtistServiceClient) AddToArtistTrack(ctx context.Context, in *ArtistsTracksReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadArtistService_AddToArtistTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadArtistServiceClient) AddToArtistGenre(ctx context.Context, in *ArtistsGenresReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadArtistService_AddToArtistGenre_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadArtistServiceServer is the server API for UploadArtistService service.
// All implementations must embed UnimplementedUploadArtistServiceServer
// for forward compatibility.
type UploadArtistServiceServer interface {
	CreateArtist(context.Context, *CreateArtistReq) (*CreateArtistRes, error)
	UpdateArtist(context.Context, *UpdateArtistReq) (*UpdateArtistRes, error)
	DeleteArtist(context.Context, *DeleteArtistReq) (*emptypb.Empty, error)
	AddToArtistTrack(context.Context, *ArtistsTracksReq) (*emptypb.Empty, error)
	AddToArtistGenre(context.Context, *ArtistsGenresReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedUploadArtistServiceServer()
}

// UnimplementedUploadArtistServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadArtistServiceServer struct{}

func (UnimplementedUploadArtistServiceServer) CreateArtist(context.Context, *CreateArtistReq) (*CreateArtistRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArtist not implemented")
}
func (UnimplementedUploadArtistServiceServer) UpdateArtist(context.Context, *UpdateArtistReq) (*UpdateArtistRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtist not implemented")
}
func (UnimplementedUploadArtistServiceServer) DeleteArtist(context.Context, *DeleteArtistReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtist not implemented")
}
func (UnimplementedUploadArtistServiceServer) AddToArtistTrack(context.Context, *ArtistsTracksReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToArtistTrack not implemented")
}
func (UnimplementedUploadArtistServiceServer) AddToArtistGenre(context.Context, *ArtistsGenresReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToArtistGenre not implemented")
}
func (UnimplementedUploadArtistServiceServer) mustEmbedUnimplementedUploadArtistServiceServer() {}
func (UnimplementedUploadArtistServiceServer) testEmbeddedByValue()                             {}

// UnsafeUploadArtistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadArtistServiceServer will
// result in compilation errors.
type UnsafeUploadArtistServiceServer interface {
	mustEmbedUnimplementedUploadArtistServiceServer()
}

func RegisterUploadArtistServiceServer(s grpc.ServiceRegistrar, srv UploadArtistServiceServer) {
	// If the following call pancis, it indicates UnimplementedUploadArtistServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UploadArtistService_ServiceDesc, srv)
}

func _UploadArtistService_CreateArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadArtistServiceServer).CreateArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadArtistService_CreateArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadArtistServiceServer).CreateArtist(ctx, req.(*CreateArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadArtistService_UpdateArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadArtistServiceServer).UpdateArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadArtistService_UpdateArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadArtistServiceServer).UpdateArtist(ctx, req.(*UpdateArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadArtistService_DeleteArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadArtistServiceServer).DeleteArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadArtistService_DeleteArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadArtistServiceServer).DeleteArtist(ctx, req.(*DeleteArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadArtistService_AddToArtistTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistsTracksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadArtistServiceServer).AddToArtistTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadArtistService_AddToArtistTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadArtistServiceServer).AddToArtistTrack(ctx, req.(*ArtistsTracksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadArtistService_AddToArtistGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistsGenresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadArtistServiceServer).AddToArtistGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadArtistService_AddToArtistGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadArtistServiceServer).AddToArtistGenre(ctx, req.(*ArtistsGenresReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadArtistService_ServiceDesc is the grpc.ServiceDesc for UploadArtistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadArtistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploadArtistService",
	HandlerType: (*UploadArtistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArtist",
			Handler:    _UploadArtistService_CreateArtist_Handler,
		},
		{
			MethodName: "UpdateArtist",
			Handler:    _UploadArtistService_UpdateArtist_Handler,
		},
		{
			MethodName: "DeleteArtist",
			Handler:    _UploadArtistService_DeleteArtist_Handler,
		},
		{
			MethodName: "AddToArtistTrack",
			Handler:    _UploadArtistService_AddToArtistTrack_Handler,
		},
		{
			MethodName: "AddToArtistGenre",
			Handler:    _UploadArtistService_AddToArtistGenre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload/upload.proto",
}

const (
	UploadGenresService_CreateGenre_FullMethodName = "/proto.UploadGenresService/CreateGenre"
)

// UploadGenresServiceClient is the client API for UploadGenresService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadGenresServiceClient interface {
	CreateGenre(ctx context.Context, in *CreateGenreReq, opts ...grpc.CallOption) (*CreateGenreRes, error)
}

type uploadGenresServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadGenresServiceClient(cc grpc.ClientConnInterface) UploadGenresServiceClient {
	return &uploadGenresServiceClient{cc}
}

func (c *uploadGenresServiceClient) CreateGenre(ctx context.Context, in *CreateGenreReq, opts ...grpc.CallOption) (*CreateGenreRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGenreRes)
	err := c.cc.Invoke(ctx, UploadGenresService_CreateGenre_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadGenresServiceServer is the server API for UploadGenresService service.
// All implementations must embed UnimplementedUploadGenresServiceServer
// for forward compatibility.
type UploadGenresServiceServer interface {
	CreateGenre(context.Context, *CreateGenreReq) (*CreateGenreRes, error)
	mustEmbedUnimplementedUploadGenresServiceServer()
}

// UnimplementedUploadGenresServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadGenresServiceServer struct{}

func (UnimplementedUploadGenresServiceServer) CreateGenre(context.Context, *CreateGenreReq) (*CreateGenreRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenre not implemented")
}
func (UnimplementedUploadGenresServiceServer) mustEmbedUnimplementedUploadGenresServiceServer() {}
func (UnimplementedUploadGenresServiceServer) testEmbeddedByValue()                             {}

// UnsafeUploadGenresServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadGenresServiceServer will
// result in compilation errors.
type UnsafeUploadGenresServiceServer interface {
	mustEmbedUnimplementedUploadGenresServiceServer()
}

func RegisterUploadGenresServiceServer(s grpc.ServiceRegistrar, srv UploadGenresServiceServer) {
	// If the following call pancis, it indicates UnimplementedUploadGenresServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UploadGenresService_ServiceDesc, srv)
}

func _UploadGenresService_CreateGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGenreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadGenresServiceServer).CreateGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadGenresService_CreateGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadGenresServiceServer).CreateGenre(ctx, req.(*CreateGenreReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadGenresService_ServiceDesc is the grpc.ServiceDesc for UploadGenresService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadGenresService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploadGenresService",
	HandlerType: (*UploadGenresServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGenre",
			Handler:    _UploadGenresService_CreateGenre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload/upload.proto",
}

const (
	UploadTrackService_CreateTrack_FullMethodName     = "/proto.UploadTrackService/CreateTrack"
	UploadTrackService_UpdateTrack_FullMethodName     = "/proto.UploadTrackService/UpdateTrack"
	UploadTrackService_DeleteTrack_FullMethodName     = "/proto.UploadTrackService/DeleteTrack"
	UploadTrackService_AddToTrackGenre_FullMethodName = "/proto.UploadTrackService/AddToTrackGenre"
)

// UploadTrackServiceClient is the client API for UploadTrackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadTrackServiceClient interface {
	CreateTrack(ctx context.Context, in *CreateTrackReq, opts ...grpc.CallOption) (*CreateTrackRes, error)
	UpdateTrack(ctx context.Context, in *UpdateTrackReq, opts ...grpc.CallOption) (*UpdateTrackRes, error)
	DeleteTrack(ctx context.Context, in *DeleteTrackReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddToTrackGenre(ctx context.Context, in *TracksGenresReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type uploadTrackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadTrackServiceClient(cc grpc.ClientConnInterface) UploadTrackServiceClient {
	return &uploadTrackServiceClient{cc}
}

func (c *uploadTrackServiceClient) CreateTrack(ctx context.Context, in *CreateTrackReq, opts ...grpc.CallOption) (*CreateTrackRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTrackRes)
	err := c.cc.Invoke(ctx, UploadTrackService_CreateTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadTrackServiceClient) UpdateTrack(ctx context.Context, in *UpdateTrackReq, opts ...grpc.CallOption) (*UpdateTrackRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTrackRes)
	err := c.cc.Invoke(ctx, UploadTrackService_UpdateTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadTrackServiceClient) DeleteTrack(ctx context.Context, in *DeleteTrackReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadTrackService_DeleteTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadTrackServiceClient) AddToTrackGenre(ctx context.Context, in *TracksGenresReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UploadTrackService_AddToTrackGenre_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadTrackServiceServer is the server API for UploadTrackService service.
// All implementations must embed UnimplementedUploadTrackServiceServer
// for forward compatibility.
type UploadTrackServiceServer interface {
	CreateTrack(context.Context, *CreateTrackReq) (*CreateTrackRes, error)
	UpdateTrack(context.Context, *UpdateTrackReq) (*UpdateTrackRes, error)
	DeleteTrack(context.Context, *DeleteTrackReq) (*emptypb.Empty, error)
	AddToTrackGenre(context.Context, *TracksGenresReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedUploadTrackServiceServer()
}

// UnimplementedUploadTrackServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadTrackServiceServer struct{}

func (UnimplementedUploadTrackServiceServer) CreateTrack(context.Context, *CreateTrackReq) (*CreateTrackRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrack not implemented")
}
func (UnimplementedUploadTrackServiceServer) UpdateTrack(context.Context, *UpdateTrackReq) (*UpdateTrackRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrack not implemented")
}
func (UnimplementedUploadTrackServiceServer) DeleteTrack(context.Context, *DeleteTrackReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrack not implemented")
}
func (UnimplementedUploadTrackServiceServer) AddToTrackGenre(context.Context, *TracksGenresReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToTrackGenre not implemented")
}
func (UnimplementedUploadTrackServiceServer) mustEmbedUnimplementedUploadTrackServiceServer() {}
func (UnimplementedUploadTrackServiceServer) testEmbeddedByValue()                            {}

// UnsafeUploadTrackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadTrackServiceServer will
// result in compilation errors.
type UnsafeUploadTrackServiceServer interface {
	mustEmbedUnimplementedUploadTrackServiceServer()
}

func RegisterUploadTrackServiceServer(s grpc.ServiceRegistrar, srv UploadTrackServiceServer) {
	// If the following call pancis, it indicates UnimplementedUploadTrackServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UploadTrackService_ServiceDesc, srv)
}

func _UploadTrackService_CreateTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadTrackServiceServer).CreateTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadTrackService_CreateTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadTrackServiceServer).CreateTrack(ctx, req.(*CreateTrackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadTrackService_UpdateTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadTrackServiceServer).UpdateTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadTrackService_UpdateTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadTrackServiceServer).UpdateTrack(ctx, req.(*UpdateTrackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadTrackService_DeleteTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadTrackServiceServer).DeleteTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadTrackService_DeleteTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadTrackServiceServer).DeleteTrack(ctx, req.(*DeleteTrackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadTrackService_AddToTrackGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TracksGenresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadTrackServiceServer).AddToTrackGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadTrackService_AddToTrackGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadTrackServiceServer).AddToTrackGenre(ctx, req.(*TracksGenresReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadTrackService_ServiceDesc is the grpc.ServiceDesc for UploadTrackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadTrackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploadTrackService",
	HandlerType: (*UploadTrackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrack",
			Handler:    _UploadTrackService_CreateTrack_Handler,
		},
		{
			MethodName: "UpdateTrack",
			Handler:    _UploadTrackService_UpdateTrack_Handler,
		},
		{
			MethodName: "DeleteTrack",
			Handler:    _UploadTrackService_DeleteTrack_Handler,
		},
		{
			MethodName: "AddToTrackGenre",
			Handler:    _UploadTrackService_AddToTrackGenre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload/upload.proto",
}
