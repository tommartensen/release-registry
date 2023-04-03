// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	// Create creates a new Release.
	Create(ctx context.Context, in *ReleaseServiceCreateRequest, opts ...grpc.CallOption) (*ReleaseServiceCreateResponse, error)
	// List returns all tracked Releases.
	List(ctx context.Context, in *ReleaseServiceListRequest, opts ...grpc.CallOption) (*ReleaseServiceListResponse, error)
	// Get returns a Release identified by its tag.
	Get(ctx context.Context, in *ReleaseServiceGetRequest, opts ...grpc.CallOption) (*ReleaseServiceGetResponse, error)
	// Approve approves a Release identified by its tag for a QualityMilestone.
	Approve(ctx context.Context, in *ReleaseServiceApproveRequest, opts ...grpc.CallOption) (*ReleaseServiceApproveResponse, error)
	// Reject marks a Release identified by its tag as rejected.
	Reject(ctx context.Context, in *ReleaseServiceRejectRequest, opts ...grpc.CallOption) (*ReleaseServiceRejectResponse, error)
	// FindLatest returns the latest release for a given query.
	FindLatest(ctx context.Context, in *ReleaseServiceFindLatestRequest, opts ...grpc.CallOption) (*ReleaseServiceFindLatestResponse, error)
}

type releaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseServiceClient(cc grpc.ClientConnInterface) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) Create(ctx context.Context, in *ReleaseServiceCreateRequest, opts ...grpc.CallOption) (*ReleaseServiceCreateResponse, error) {
	out := new(ReleaseServiceCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.ReleaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) List(ctx context.Context, in *ReleaseServiceListRequest, opts ...grpc.CallOption) (*ReleaseServiceListResponse, error) {
	out := new(ReleaseServiceListResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.ReleaseService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Get(ctx context.Context, in *ReleaseServiceGetRequest, opts ...grpc.CallOption) (*ReleaseServiceGetResponse, error) {
	out := new(ReleaseServiceGetResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.ReleaseService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Approve(ctx context.Context, in *ReleaseServiceApproveRequest, opts ...grpc.CallOption) (*ReleaseServiceApproveResponse, error) {
	out := new(ReleaseServiceApproveResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.ReleaseService/Approve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Reject(ctx context.Context, in *ReleaseServiceRejectRequest, opts ...grpc.CallOption) (*ReleaseServiceRejectResponse, error) {
	out := new(ReleaseServiceRejectResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.ReleaseService/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) FindLatest(ctx context.Context, in *ReleaseServiceFindLatestRequest, opts ...grpc.CallOption) (*ReleaseServiceFindLatestResponse, error) {
	out := new(ReleaseServiceFindLatestResponse)
	err := c.cc.Invoke(ctx, "/proto.api.v1.ReleaseService/FindLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
// All implementations must embed UnimplementedReleaseServiceServer
// for forward compatibility
type ReleaseServiceServer interface {
	// Create creates a new Release.
	Create(context.Context, *ReleaseServiceCreateRequest) (*ReleaseServiceCreateResponse, error)
	// List returns all tracked Releases.
	List(context.Context, *ReleaseServiceListRequest) (*ReleaseServiceListResponse, error)
	// Get returns a Release identified by its tag.
	Get(context.Context, *ReleaseServiceGetRequest) (*ReleaseServiceGetResponse, error)
	// Approve approves a Release identified by its tag for a QualityMilestone.
	Approve(context.Context, *ReleaseServiceApproveRequest) (*ReleaseServiceApproveResponse, error)
	// Reject marks a Release identified by its tag as rejected.
	Reject(context.Context, *ReleaseServiceRejectRequest) (*ReleaseServiceRejectResponse, error)
	// FindLatest returns the latest release for a given query.
	FindLatest(context.Context, *ReleaseServiceFindLatestRequest) (*ReleaseServiceFindLatestResponse, error)
	mustEmbedUnimplementedReleaseServiceServer()
}

// UnimplementedReleaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReleaseServiceServer struct {
}

func (UnimplementedReleaseServiceServer) Create(context.Context, *ReleaseServiceCreateRequest) (*ReleaseServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReleaseServiceServer) List(context.Context, *ReleaseServiceListRequest) (*ReleaseServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedReleaseServiceServer) Get(context.Context, *ReleaseServiceGetRequest) (*ReleaseServiceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedReleaseServiceServer) Approve(context.Context, *ReleaseServiceApproveRequest) (*ReleaseServiceApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedReleaseServiceServer) Reject(context.Context, *ReleaseServiceRejectRequest) (*ReleaseServiceRejectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reject not implemented")
}
func (UnimplementedReleaseServiceServer) FindLatest(context.Context, *ReleaseServiceFindLatestRequest) (*ReleaseServiceFindLatestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatest not implemented")
}
func (UnimplementedReleaseServiceServer) mustEmbedUnimplementedReleaseServiceServer() {}

// UnsafeReleaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReleaseServiceServer will
// result in compilation errors.
type UnsafeReleaseServiceServer interface {
	mustEmbedUnimplementedReleaseServiceServer()
}

func RegisterReleaseServiceServer(s grpc.ServiceRegistrar, srv ReleaseServiceServer) {
	s.RegisterService(&ReleaseService_ServiceDesc, srv)
}

func _ReleaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.ReleaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Create(ctx, req.(*ReleaseServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.ReleaseService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).List(ctx, req.(*ReleaseServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseServiceGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.ReleaseService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Get(ctx, req.(*ReleaseServiceGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseServiceApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.ReleaseService/Approve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Approve(ctx, req.(*ReleaseServiceApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseServiceRejectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.ReleaseService/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Reject(ctx, req.(*ReleaseServiceRejectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_FindLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseServiceFindLatestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).FindLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.v1.ReleaseService/FindLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).FindLatest(ctx, req.(*ReleaseServiceFindLatestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReleaseService_ServiceDesc is the grpc.ServiceDesc for ReleaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReleaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.v1.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReleaseService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ReleaseService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ReleaseService_Get_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _ReleaseService_Approve_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _ReleaseService_Reject_Handler,
		},
		{
			MethodName: "FindLatest",
			Handler:    _ReleaseService_FindLatest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/v1/release.proto",
}
