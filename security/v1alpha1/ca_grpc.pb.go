// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: security/v1alpha1/ca.proto

// Keep this package for backward compatibility.

package v1alpha1

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
	IstioCertificateService_CreateCertificate_FullMethodName = "/istio.v1.auth.IstioCertificateService/CreateCertificate"
)

// IstioCertificateServiceClient is the client API for IstioCertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for managing certificates issued by the CA.
type IstioCertificateServiceClient interface {
	// Using provided CSR, returns a signed certificate.
	CreateCertificate(ctx context.Context, in *IstioCertificateRequest, opts ...grpc.CallOption) (*IstioCertificateResponse, error)
}

type istioCertificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIstioCertificateServiceClient(cc grpc.ClientConnInterface) IstioCertificateServiceClient {
	return &istioCertificateServiceClient{cc}
}

func (c *istioCertificateServiceClient) CreateCertificate(ctx context.Context, in *IstioCertificateRequest, opts ...grpc.CallOption) (*IstioCertificateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IstioCertificateResponse)
	err := c.cc.Invoke(ctx, IstioCertificateService_CreateCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IstioCertificateServiceServer is the server API for IstioCertificateService service.
// All implementations must embed UnimplementedIstioCertificateServiceServer
// for forward compatibility.
//
// Service for managing certificates issued by the CA.
type IstioCertificateServiceServer interface {
	// Using provided CSR, returns a signed certificate.
	CreateCertificate(context.Context, *IstioCertificateRequest) (*IstioCertificateResponse, error)
	mustEmbedUnimplementedIstioCertificateServiceServer()
}

// UnimplementedIstioCertificateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIstioCertificateServiceServer struct{}

func (UnimplementedIstioCertificateServiceServer) CreateCertificate(context.Context, *IstioCertificateRequest) (*IstioCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedIstioCertificateServiceServer) mustEmbedUnimplementedIstioCertificateServiceServer() {
}
func (UnimplementedIstioCertificateServiceServer) testEmbeddedByValue() {}

// UnsafeIstioCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IstioCertificateServiceServer will
// result in compilation errors.
type UnsafeIstioCertificateServiceServer interface {
	mustEmbedUnimplementedIstioCertificateServiceServer()
}

func RegisterIstioCertificateServiceServer(s grpc.ServiceRegistrar, srv IstioCertificateServiceServer) {
	// If the following call pancis, it indicates UnimplementedIstioCertificateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IstioCertificateService_ServiceDesc, srv)
}

func _IstioCertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IstioCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioCertificateService_CreateCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCertificateServiceServer).CreateCertificate(ctx, req.(*IstioCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IstioCertificateService_ServiceDesc is the grpc.ServiceDesc for IstioCertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IstioCertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "istio.v1.auth.IstioCertificateService",
	HandlerType: (*IstioCertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _IstioCertificateService_CreateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/v1alpha1/ca.proto",
}
