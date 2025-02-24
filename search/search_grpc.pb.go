// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: search/search.proto

package search

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
	SearchAlbumsService_SearchAlbumsWithName_FullMethodName   = "/proto.SearchAlbumsService/SearchAlbumsWithName"
	SearchAlbumsService_SearchAlbumsWithGenres_FullMethodName = "/proto.SearchAlbumsService/SearchAlbumsWithGenres"
)

// SearchAlbumsServiceClient is the client API for SearchAlbumsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchAlbumsServiceClient interface {
	SearchAlbumsWithName(ctx context.Context, in *SearchWithNameReq, opts ...grpc.CallOption) (*SearchAlbumsRes, error)
	SearchAlbumsWithGenres(ctx context.Context, in *SearchWithGenresReq, opts ...grpc.CallOption) (*SearchAlbumsRes, error)
}

type searchAlbumsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchAlbumsServiceClient(cc grpc.ClientConnInterface) SearchAlbumsServiceClient {
	return &searchAlbumsServiceClient{cc}
}

func (c *searchAlbumsServiceClient) SearchAlbumsWithName(ctx context.Context, in *SearchWithNameReq, opts ...grpc.CallOption) (*SearchAlbumsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchAlbumsRes)
	err := c.cc.Invoke(ctx, SearchAlbumsService_SearchAlbumsWithName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchAlbumsServiceClient) SearchAlbumsWithGenres(ctx context.Context, in *SearchWithGenresReq, opts ...grpc.CallOption) (*SearchAlbumsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchAlbumsRes)
	err := c.cc.Invoke(ctx, SearchAlbumsService_SearchAlbumsWithGenres_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchAlbumsServiceServer is the server API for SearchAlbumsService service.
// All implementations must embed UnimplementedSearchAlbumsServiceServer
// for forward compatibility.
type SearchAlbumsServiceServer interface {
	SearchAlbumsWithName(context.Context, *SearchWithNameReq) (*SearchAlbumsRes, error)
	SearchAlbumsWithGenres(context.Context, *SearchWithGenresReq) (*SearchAlbumsRes, error)
	mustEmbedUnimplementedSearchAlbumsServiceServer()
}

// UnimplementedSearchAlbumsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchAlbumsServiceServer struct{}

func (UnimplementedSearchAlbumsServiceServer) SearchAlbumsWithName(context.Context, *SearchWithNameReq) (*SearchAlbumsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAlbumsWithName not implemented")
}
func (UnimplementedSearchAlbumsServiceServer) SearchAlbumsWithGenres(context.Context, *SearchWithGenresReq) (*SearchAlbumsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAlbumsWithGenres not implemented")
}
func (UnimplementedSearchAlbumsServiceServer) mustEmbedUnimplementedSearchAlbumsServiceServer() {}
func (UnimplementedSearchAlbumsServiceServer) testEmbeddedByValue()                             {}

// UnsafeSearchAlbumsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchAlbumsServiceServer will
// result in compilation errors.
type UnsafeSearchAlbumsServiceServer interface {
	mustEmbedUnimplementedSearchAlbumsServiceServer()
}

func RegisterSearchAlbumsServiceServer(s grpc.ServiceRegistrar, srv SearchAlbumsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSearchAlbumsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchAlbumsService_ServiceDesc, srv)
}

func _SearchAlbumsService_SearchAlbumsWithName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWithNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchAlbumsServiceServer).SearchAlbumsWithName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchAlbumsService_SearchAlbumsWithName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchAlbumsServiceServer).SearchAlbumsWithName(ctx, req.(*SearchWithNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchAlbumsService_SearchAlbumsWithGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWithGenresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchAlbumsServiceServer).SearchAlbumsWithGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchAlbumsService_SearchAlbumsWithGenres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchAlbumsServiceServer).SearchAlbumsWithGenres(ctx, req.(*SearchWithGenresReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchAlbumsService_ServiceDesc is the grpc.ServiceDesc for SearchAlbumsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchAlbumsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchAlbumsService",
	HandlerType: (*SearchAlbumsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchAlbumsWithName",
			Handler:    _SearchAlbumsService_SearchAlbumsWithName_Handler,
		},
		{
			MethodName: "SearchAlbumsWithGenres",
			Handler:    _SearchAlbumsService_SearchAlbumsWithGenres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search/search.proto",
}

const (
	SearchArtistsService_SearchArtistsWithName_FullMethodName   = "/proto.SearchArtistsService/SearchArtistsWithName"
	SearchArtistsService_SearchArtistsWithGenres_FullMethodName = "/proto.SearchArtistsService/SearchArtistsWithGenres"
)

// SearchArtistsServiceClient is the client API for SearchArtistsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchArtistsServiceClient interface {
	SearchArtistsWithName(ctx context.Context, in *SearchWithNameReq, opts ...grpc.CallOption) (*SearchArtistsRes, error)
	SearchArtistsWithGenres(ctx context.Context, in *SearchWithGenresReq, opts ...grpc.CallOption) (*SearchArtistsRes, error)
}

type searchArtistsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchArtistsServiceClient(cc grpc.ClientConnInterface) SearchArtistsServiceClient {
	return &searchArtistsServiceClient{cc}
}

func (c *searchArtistsServiceClient) SearchArtistsWithName(ctx context.Context, in *SearchWithNameReq, opts ...grpc.CallOption) (*SearchArtistsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchArtistsRes)
	err := c.cc.Invoke(ctx, SearchArtistsService_SearchArtistsWithName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchArtistsServiceClient) SearchArtistsWithGenres(ctx context.Context, in *SearchWithGenresReq, opts ...grpc.CallOption) (*SearchArtistsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchArtistsRes)
	err := c.cc.Invoke(ctx, SearchArtistsService_SearchArtistsWithGenres_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchArtistsServiceServer is the server API for SearchArtistsService service.
// All implementations must embed UnimplementedSearchArtistsServiceServer
// for forward compatibility.
type SearchArtistsServiceServer interface {
	SearchArtistsWithName(context.Context, *SearchWithNameReq) (*SearchArtistsRes, error)
	SearchArtistsWithGenres(context.Context, *SearchWithGenresReq) (*SearchArtistsRes, error)
	mustEmbedUnimplementedSearchArtistsServiceServer()
}

// UnimplementedSearchArtistsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchArtistsServiceServer struct{}

func (UnimplementedSearchArtistsServiceServer) SearchArtistsWithName(context.Context, *SearchWithNameReq) (*SearchArtistsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArtistsWithName not implemented")
}
func (UnimplementedSearchArtistsServiceServer) SearchArtistsWithGenres(context.Context, *SearchWithGenresReq) (*SearchArtistsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArtistsWithGenres not implemented")
}
func (UnimplementedSearchArtistsServiceServer) mustEmbedUnimplementedSearchArtistsServiceServer() {}
func (UnimplementedSearchArtistsServiceServer) testEmbeddedByValue()                              {}

// UnsafeSearchArtistsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchArtistsServiceServer will
// result in compilation errors.
type UnsafeSearchArtistsServiceServer interface {
	mustEmbedUnimplementedSearchArtistsServiceServer()
}

func RegisterSearchArtistsServiceServer(s grpc.ServiceRegistrar, srv SearchArtistsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSearchArtistsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchArtistsService_ServiceDesc, srv)
}

func _SearchArtistsService_SearchArtistsWithName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWithNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchArtistsServiceServer).SearchArtistsWithName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchArtistsService_SearchArtistsWithName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchArtistsServiceServer).SearchArtistsWithName(ctx, req.(*SearchWithNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchArtistsService_SearchArtistsWithGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWithGenresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchArtistsServiceServer).SearchArtistsWithGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchArtistsService_SearchArtistsWithGenres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchArtistsServiceServer).SearchArtistsWithGenres(ctx, req.(*SearchWithGenresReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchArtistsService_ServiceDesc is the grpc.ServiceDesc for SearchArtistsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchArtistsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchArtistsService",
	HandlerType: (*SearchArtistsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchArtistsWithName",
			Handler:    _SearchArtistsService_SearchArtistsWithName_Handler,
		},
		{
			MethodName: "SearchArtistsWithGenres",
			Handler:    _SearchArtistsService_SearchArtistsWithGenres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search/search.proto",
}

const (
	SearchTracksService_SearchTracksWithName_FullMethodName   = "/proto.SearchTracksService/SearchTracksWithName"
	SearchTracksService_SearchTracksWithGenres_FullMethodName = "/proto.SearchTracksService/SearchTracksWithGenres"
	SearchTracksService_ListenersChart_FullMethodName         = "/proto.SearchTracksService/ListenersChart"
	SearchTracksService_PlaycountsChart_FullMethodName        = "/proto.SearchTracksService/PlaycountsChart"
)

// SearchTracksServiceClient is the client API for SearchTracksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchTracksServiceClient interface {
	SearchTracksWithName(ctx context.Context, in *SearchWithNameReq, opts ...grpc.CallOption) (*SearchTracksRes, error)
	SearchTracksWithGenres(ctx context.Context, in *SearchWithGenresReq, opts ...grpc.CallOption) (*SearchTracksRes, error)
	ListenersChart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchTracksRes, error)
	PlaycountsChart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchTracksRes, error)
}

type searchTracksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchTracksServiceClient(cc grpc.ClientConnInterface) SearchTracksServiceClient {
	return &searchTracksServiceClient{cc}
}

func (c *searchTracksServiceClient) SearchTracksWithName(ctx context.Context, in *SearchWithNameReq, opts ...grpc.CallOption) (*SearchTracksRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTracksRes)
	err := c.cc.Invoke(ctx, SearchTracksService_SearchTracksWithName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchTracksServiceClient) SearchTracksWithGenres(ctx context.Context, in *SearchWithGenresReq, opts ...grpc.CallOption) (*SearchTracksRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTracksRes)
	err := c.cc.Invoke(ctx, SearchTracksService_SearchTracksWithGenres_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchTracksServiceClient) ListenersChart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchTracksRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTracksRes)
	err := c.cc.Invoke(ctx, SearchTracksService_ListenersChart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchTracksServiceClient) PlaycountsChart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchTracksRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTracksRes)
	err := c.cc.Invoke(ctx, SearchTracksService_PlaycountsChart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchTracksServiceServer is the server API for SearchTracksService service.
// All implementations must embed UnimplementedSearchTracksServiceServer
// for forward compatibility.
type SearchTracksServiceServer interface {
	SearchTracksWithName(context.Context, *SearchWithNameReq) (*SearchTracksRes, error)
	SearchTracksWithGenres(context.Context, *SearchWithGenresReq) (*SearchTracksRes, error)
	ListenersChart(context.Context, *emptypb.Empty) (*SearchTracksRes, error)
	PlaycountsChart(context.Context, *emptypb.Empty) (*SearchTracksRes, error)
	mustEmbedUnimplementedSearchTracksServiceServer()
}

// UnimplementedSearchTracksServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchTracksServiceServer struct{}

func (UnimplementedSearchTracksServiceServer) SearchTracksWithName(context.Context, *SearchWithNameReq) (*SearchTracksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTracksWithName not implemented")
}
func (UnimplementedSearchTracksServiceServer) SearchTracksWithGenres(context.Context, *SearchWithGenresReq) (*SearchTracksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTracksWithGenres not implemented")
}
func (UnimplementedSearchTracksServiceServer) ListenersChart(context.Context, *emptypb.Empty) (*SearchTracksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenersChart not implemented")
}
func (UnimplementedSearchTracksServiceServer) PlaycountsChart(context.Context, *emptypb.Empty) (*SearchTracksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaycountsChart not implemented")
}
func (UnimplementedSearchTracksServiceServer) mustEmbedUnimplementedSearchTracksServiceServer() {}
func (UnimplementedSearchTracksServiceServer) testEmbeddedByValue()                             {}

// UnsafeSearchTracksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchTracksServiceServer will
// result in compilation errors.
type UnsafeSearchTracksServiceServer interface {
	mustEmbedUnimplementedSearchTracksServiceServer()
}

func RegisterSearchTracksServiceServer(s grpc.ServiceRegistrar, srv SearchTracksServiceServer) {
	// If the following call pancis, it indicates UnimplementedSearchTracksServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchTracksService_ServiceDesc, srv)
}

func _SearchTracksService_SearchTracksWithName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWithNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTracksServiceServer).SearchTracksWithName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchTracksService_SearchTracksWithName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTracksServiceServer).SearchTracksWithName(ctx, req.(*SearchWithNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchTracksService_SearchTracksWithGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWithGenresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTracksServiceServer).SearchTracksWithGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchTracksService_SearchTracksWithGenres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTracksServiceServer).SearchTracksWithGenres(ctx, req.(*SearchWithGenresReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchTracksService_ListenersChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTracksServiceServer).ListenersChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchTracksService_ListenersChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTracksServiceServer).ListenersChart(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchTracksService_PlaycountsChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTracksServiceServer).PlaycountsChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchTracksService_PlaycountsChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTracksServiceServer).PlaycountsChart(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchTracksService_ServiceDesc is the grpc.ServiceDesc for SearchTracksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchTracksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchTracksService",
	HandlerType: (*SearchTracksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchTracksWithName",
			Handler:    _SearchTracksService_SearchTracksWithName_Handler,
		},
		{
			MethodName: "SearchTracksWithGenres",
			Handler:    _SearchTracksService_SearchTracksWithGenres_Handler,
		},
		{
			MethodName: "ListenersChart",
			Handler:    _SearchTracksService_ListenersChart_Handler,
		},
		{
			MethodName: "PlaycountsChart",
			Handler:    _SearchTracksService_PlaycountsChart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search/search.proto",
}
