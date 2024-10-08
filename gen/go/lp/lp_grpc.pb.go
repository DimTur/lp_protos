// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: lp.proto

package lp1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LearningPlatform_CreateChannel_FullMethodName = "/lp.v1.LearningPlatform/CreateChannel"
	LearningPlatform_CreatePlan_FullMethodName    = "/lp.v1.LearningPlatform/CreatePlan"
	LearningPlatform_CreateLesson_FullMethodName  = "/lp.v1.LearningPlatform/CreateLesson"
	LearningPlatform_GetChannel_FullMethodName    = "/lp.v1.LearningPlatform/GetChannel"
	LearningPlatform_GetPlan_FullMethodName       = "/lp.v1.LearningPlatform/GetPlan"
	LearningPlatform_GetLesson_FullMethodName     = "/lp.v1.LearningPlatform/GetLesson"
)

// LearningPlatformClient is the client API for LearningPlatform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearningPlatformClient interface {
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error)
	CreateLesson(ctx context.Context, in *CreateLessonRequest, opts ...grpc.CallOption) (*CreateLessonResponse, error)
	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error)
	GetPlan(ctx context.Context, in *GetPlanRequest, opts ...grpc.CallOption) (*GetPlanResponse, error)
	GetLesson(ctx context.Context, in *GetLessonRequest, opts ...grpc.CallOption) (*GetLessonResponse, error)
}

type learningPlatformClient struct {
	cc grpc.ClientConnInterface
}

func NewLearningPlatformClient(cc grpc.ClientConnInterface) LearningPlatformClient {
	return &learningPlatformClient{cc}
}

func (c *learningPlatformClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, LearningPlatform_CreateChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlatformClient) CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePlanResponse)
	err := c.cc.Invoke(ctx, LearningPlatform_CreatePlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlatformClient) CreateLesson(ctx context.Context, in *CreateLessonRequest, opts ...grpc.CallOption) (*CreateLessonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLessonResponse)
	err := c.cc.Invoke(ctx, LearningPlatform_CreateLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlatformClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChannelResponse)
	err := c.cc.Invoke(ctx, LearningPlatform_GetChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlatformClient) GetPlan(ctx context.Context, in *GetPlanRequest, opts ...grpc.CallOption) (*GetPlanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlanResponse)
	err := c.cc.Invoke(ctx, LearningPlatform_GetPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningPlatformClient) GetLesson(ctx context.Context, in *GetLessonRequest, opts ...grpc.CallOption) (*GetLessonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLessonResponse)
	err := c.cc.Invoke(ctx, LearningPlatform_GetLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearningPlatformServer is the server API for LearningPlatform service.
// All implementations must embed UnimplementedLearningPlatformServer
// for forward compatibility.
type LearningPlatformServer interface {
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error)
	CreateLesson(context.Context, *CreateLessonRequest) (*CreateLessonResponse, error)
	GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error)
	GetPlan(context.Context, *GetPlanRequest) (*GetPlanResponse, error)
	GetLesson(context.Context, *GetLessonRequest) (*GetLessonResponse, error)
	mustEmbedUnimplementedLearningPlatformServer()
}

// UnimplementedLearningPlatformServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLearningPlatformServer struct{}

func (UnimplementedLearningPlatformServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedLearningPlatformServer) CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedLearningPlatformServer) CreateLesson(context.Context, *CreateLessonRequest) (*CreateLessonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLesson not implemented")
}
func (UnimplementedLearningPlatformServer) GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}
func (UnimplementedLearningPlatformServer) GetPlan(context.Context, *GetPlanRequest) (*GetPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlan not implemented")
}
func (UnimplementedLearningPlatformServer) GetLesson(context.Context, *GetLessonRequest) (*GetLessonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLesson not implemented")
}
func (UnimplementedLearningPlatformServer) mustEmbedUnimplementedLearningPlatformServer() {}
func (UnimplementedLearningPlatformServer) testEmbeddedByValue()                          {}

// UnsafeLearningPlatformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearningPlatformServer will
// result in compilation errors.
type UnsafeLearningPlatformServer interface {
	mustEmbedUnimplementedLearningPlatformServer()
}

func RegisterLearningPlatformServer(s grpc.ServiceRegistrar, srv LearningPlatformServer) {
	// If the following call pancis, it indicates UnimplementedLearningPlatformServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LearningPlatform_ServiceDesc, srv)
}

func _LearningPlatform_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlatformServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningPlatform_CreateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlatformServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlatform_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlatformServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningPlatform_CreatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlatformServer).CreatePlan(ctx, req.(*CreatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlatform_CreateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlatformServer).CreateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningPlatform_CreateLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlatformServer).CreateLesson(ctx, req.(*CreateLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlatform_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlatformServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningPlatform_GetChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlatformServer).GetChannel(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlatform_GetPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlatformServer).GetPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningPlatform_GetPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlatformServer).GetPlan(ctx, req.(*GetPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningPlatform_GetLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningPlatformServer).GetLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningPlatform_GetLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningPlatformServer).GetLesson(ctx, req.(*GetLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LearningPlatform_ServiceDesc is the grpc.ServiceDesc for LearningPlatform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearningPlatform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lp.v1.LearningPlatform",
	HandlerType: (*LearningPlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _LearningPlatform_CreateChannel_Handler,
		},
		{
			MethodName: "CreatePlan",
			Handler:    _LearningPlatform_CreatePlan_Handler,
		},
		{
			MethodName: "CreateLesson",
			Handler:    _LearningPlatform_CreateLesson_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _LearningPlatform_GetChannel_Handler,
		},
		{
			MethodName: "GetPlan",
			Handler:    _LearningPlatform_GetPlan_Handler,
		},
		{
			MethodName: "GetLesson",
			Handler:    _LearningPlatform_GetLesson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lp.proto",
}
