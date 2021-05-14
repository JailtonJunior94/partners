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

// CepServiceClient is the client API for CepService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CepServiceClient interface {
	GetCep(ctx context.Context, in *CepRequest, opts ...grpc.CallOption) (*CepResponse, error)
}

type cepServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCepServiceClient(cc grpc.ClientConnInterface) CepServiceClient {
	return &cepServiceClient{cc}
}

func (c *cepServiceClient) GetCep(ctx context.Context, in *CepRequest, opts ...grpc.CallOption) (*CepResponse, error) {
	out := new(CepResponse)
	err := c.cc.Invoke(ctx, "/github.com.jailtonjunior94.go.challenge.CepService/GetCep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CepServiceServer is the server API for CepService service.
// All implementations must embed UnimplementedCepServiceServer
// for forward compatibility
type CepServiceServer interface {
	GetCep(context.Context, *CepRequest) (*CepResponse, error)
	mustEmbedUnimplementedCepServiceServer()
}

// UnimplementedCepServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCepServiceServer struct {
}

func (UnimplementedCepServiceServer) GetCep(context.Context, *CepRequest) (*CepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCep not implemented")
}
func (UnimplementedCepServiceServer) mustEmbedUnimplementedCepServiceServer() {}

// UnsafeCepServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CepServiceServer will
// result in compilation errors.
type UnsafeCepServiceServer interface {
	mustEmbedUnimplementedCepServiceServer()
}

func RegisterCepServiceServer(s grpc.ServiceRegistrar, srv CepServiceServer) {
	s.RegisterService(&CepService_ServiceDesc, srv)
}

func _CepService_GetCep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CepServiceServer).GetCep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.jailtonjunior94.go.challenge.CepService/GetCep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CepServiceServer).GetCep(ctx, req.(*CepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CepService_ServiceDesc is the grpc.ServiceDesc for CepService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CepService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.jailtonjunior94.go.challenge.CepService",
	HandlerType: (*CepServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCep",
			Handler:    _CepService_GetCep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cep.proto",
}
