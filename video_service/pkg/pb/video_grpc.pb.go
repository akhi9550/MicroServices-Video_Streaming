// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/video.proto

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

const (
	VideoService_UploadVideo_FullMethodName  = "/pb.VideoService/UploadVideo"
	VideoService_StreamVideo_FullMethodName  = "/pb.VideoService/StreamVideo"
	VideoService_FindAllVideo_FullMethodName = "/pb.VideoService/FindAllVideo"
)

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	UploadVideo(ctx context.Context, opts ...grpc.CallOption) (VideoService_UploadVideoClient, error)
	StreamVideo(ctx context.Context, in *StreamVideoRequest, opts ...grpc.CallOption) (VideoService_StreamVideoClient, error)
	FindAllVideo(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (VideoService_UploadVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoService_ServiceDesc.Streams[0], VideoService_UploadVideo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &videoServiceUploadVideoClient{stream}
	return x, nil
}

type VideoService_UploadVideoClient interface {
	Send(*UploadVideoRequest) error
	CloseAndRecv() (*UploadVideoResponse, error)
	grpc.ClientStream
}

type videoServiceUploadVideoClient struct {
	grpc.ClientStream
}

func (x *videoServiceUploadVideoClient) Send(m *UploadVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoServiceUploadVideoClient) CloseAndRecv() (*UploadVideoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadVideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoServiceClient) StreamVideo(ctx context.Context, in *StreamVideoRequest, opts ...grpc.CallOption) (VideoService_StreamVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoService_ServiceDesc.Streams[1], VideoService_StreamVideo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &videoServiceStreamVideoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoService_StreamVideoClient interface {
	Recv() (*StreamVideoResponse, error)
	grpc.ClientStream
}

type videoServiceStreamVideoClient struct {
	grpc.ClientStream
}

func (x *videoServiceStreamVideoClient) Recv() (*StreamVideoResponse, error) {
	m := new(StreamVideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoServiceClient) FindAllVideo(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, VideoService_FindAllVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	UploadVideo(VideoService_UploadVideoServer) error
	StreamVideo(*StreamVideoRequest, VideoService_StreamVideoServer) error
	FindAllVideo(context.Context, *FindAllRequest) (*FindAllResponse, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) UploadVideo(VideoService_UploadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoServiceServer) StreamVideo(*StreamVideoRequest, VideoService_StreamVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamVideo not implemented")
}
func (UnimplementedVideoServiceServer) FindAllVideo(context.Context, *FindAllRequest) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllVideo not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_UploadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoServiceServer).UploadVideo(&videoServiceUploadVideoServer{stream})
}

type VideoService_UploadVideoServer interface {
	SendAndClose(*UploadVideoResponse) error
	Recv() (*UploadVideoRequest, error)
	grpc.ServerStream
}

type videoServiceUploadVideoServer struct {
	grpc.ServerStream
}

func (x *videoServiceUploadVideoServer) SendAndClose(m *UploadVideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoServiceUploadVideoServer) Recv() (*UploadVideoRequest, error) {
	m := new(UploadVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VideoService_StreamVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamVideoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoServiceServer).StreamVideo(m, &videoServiceStreamVideoServer{stream})
}

type VideoService_StreamVideoServer interface {
	Send(*StreamVideoResponse) error
	grpc.ServerStream
}

type videoServiceStreamVideoServer struct {
	grpc.ServerStream
}

func (x *videoServiceStreamVideoServer) Send(m *StreamVideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VideoService_FindAllVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).FindAllVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_FindAllVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).FindAllVideo(ctx, req.(*FindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllVideo",
			Handler:    _VideoService_FindAllVideo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadVideo",
			Handler:       _VideoService_UploadVideo_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamVideo",
			Handler:       _VideoService_StreamVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/pb/video.proto",
}
