// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: pkg/fluxion-grpc/fluxion.proto

package fluxion_grpc

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

// FluxionServiceClient is the client API for FluxionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FluxionServiceClient interface {
	// Sends a Match command
	Match(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (*MatchResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
}

type fluxionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFluxionServiceClient(cc grpc.ClientConnInterface) FluxionServiceClient {
	return &fluxionServiceClient{cc}
}

func (c *fluxionServiceClient) Match(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (*MatchResponse, error) {
	out := new(MatchResponse)
	err := c.cc.Invoke(ctx, "/fluxion.FluxionService/Match", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxionServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/fluxion.FluxionService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxionServiceClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/fluxion.FluxionService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FluxionServiceServer is the server API for FluxionService service.
// All implementations must embed UnimplementedFluxionServiceServer
// for forward compatibility
type FluxionServiceServer interface {
	// Sends a Match command
	Match(context.Context, *MatchRequest) (*MatchResponse, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
	Init(context.Context, *InitRequest) (*InitResponse, error)
	mustEmbedUnimplementedFluxionServiceServer()
}

// UnimplementedFluxionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFluxionServiceServer struct {
}

func (UnimplementedFluxionServiceServer) Match(context.Context, *MatchRequest) (*MatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (UnimplementedFluxionServiceServer) Cancel(context.Context, *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedFluxionServiceServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedFluxionServiceServer) mustEmbedUnimplementedFluxionServiceServer() {}

// UnsafeFluxionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FluxionServiceServer will
// result in compilation errors.
type UnsafeFluxionServiceServer interface {
	mustEmbedUnimplementedFluxionServiceServer()
}

func RegisterFluxionServiceServer(s grpc.ServiceRegistrar, srv FluxionServiceServer) {
	s.RegisterService(&FluxionService_ServiceDesc, srv)
}

func _FluxionService_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxionServiceServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fluxion.FluxionService/Match",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxionServiceServer).Match(ctx, req.(*MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxionService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxionServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fluxion.FluxionService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxionServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxionService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxionServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fluxion.FluxionService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxionServiceServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FluxionService_ServiceDesc is the grpc.ServiceDesc for FluxionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FluxionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fluxion.FluxionService",
	HandlerType: (*FluxionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Match",
			Handler:    _FluxionService_Match_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _FluxionService_Cancel_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _FluxionService_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/fluxion-grpc/fluxion.proto",
}
