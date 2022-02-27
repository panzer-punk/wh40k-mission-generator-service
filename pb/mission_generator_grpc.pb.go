// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// MissionGeneratorClient is the client API for MissionGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MissionGeneratorClient interface {
	Generate(ctx context.Context, in *GenerateMissionRequest, opts ...grpc.CallOption) (*Mission, error)
	List(ctx context.Context, in *ListMissionRequest, opts ...grpc.CallOption) (MissionGenerator_ListClient, error)
}

type missionGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewMissionGeneratorClient(cc grpc.ClientConnInterface) MissionGeneratorClient {
	return &missionGeneratorClient{cc}
}

func (c *missionGeneratorClient) Generate(ctx context.Context, in *GenerateMissionRequest, opts ...grpc.CallOption) (*Mission, error) {
	out := new(Mission)
	err := c.cc.Invoke(ctx, "/pb.MissionGenerator/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionGeneratorClient) List(ctx context.Context, in *ListMissionRequest, opts ...grpc.CallOption) (MissionGenerator_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &MissionGenerator_ServiceDesc.Streams[0], "/pb.MissionGenerator/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &missionGeneratorListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MissionGenerator_ListClient interface {
	Recv() (*Mission, error)
	grpc.ClientStream
}

type missionGeneratorListClient struct {
	grpc.ClientStream
}

func (x *missionGeneratorListClient) Recv() (*Mission, error) {
	m := new(Mission)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MissionGeneratorServer is the server API for MissionGenerator service.
// All implementations must embed UnimplementedMissionGeneratorServer
// for forward compatibility
type MissionGeneratorServer interface {
	Generate(context.Context, *GenerateMissionRequest) (*Mission, error)
	List(*ListMissionRequest, MissionGenerator_ListServer) error
	mustEmbedUnimplementedMissionGeneratorServer()
}

// UnimplementedMissionGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedMissionGeneratorServer struct {
}

func (UnimplementedMissionGeneratorServer) Generate(context.Context, *GenerateMissionRequest) (*Mission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedMissionGeneratorServer) List(*ListMissionRequest, MissionGenerator_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMissionGeneratorServer) mustEmbedUnimplementedMissionGeneratorServer() {}

// UnsafeMissionGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MissionGeneratorServer will
// result in compilation errors.
type UnsafeMissionGeneratorServer interface {
	mustEmbedUnimplementedMissionGeneratorServer()
}

func RegisterMissionGeneratorServer(s grpc.ServiceRegistrar, srv MissionGeneratorServer) {
	s.RegisterService(&MissionGenerator_ServiceDesc, srv)
}

func _MissionGenerator_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionGeneratorServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MissionGenerator/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionGeneratorServer).Generate(ctx, req.(*GenerateMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionGenerator_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMissionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MissionGeneratorServer).List(m, &missionGeneratorListServer{stream})
}

type MissionGenerator_ListServer interface {
	Send(*Mission) error
	grpc.ServerStream
}

type missionGeneratorListServer struct {
	grpc.ServerStream
}

func (x *missionGeneratorListServer) Send(m *Mission) error {
	return x.ServerStream.SendMsg(m)
}

// MissionGenerator_ServiceDesc is the grpc.ServiceDesc for MissionGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MissionGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MissionGenerator",
	HandlerType: (*MissionGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _MissionGenerator_Generate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _MissionGenerator_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mission_generator.proto",
}