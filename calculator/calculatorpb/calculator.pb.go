// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

/*
Package calculatorpb is a generated protocol buffer package.

It is generated from these files:
	calculator/calculatorpb/calculator.proto

It has these top-level messages:
	Nums
	SumRequest
	SumResponse
	PrimeNumberDecompositionRequest
	PrimeNumberDecompositionResponse
	ComputeAverageRequest
	ComputeAverageResponse
	FindMaximumRequest
	FindMaximumResponse
	SquareRootRequest
	SquareRootResponse
*/
package calculatorpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Nums struct {
	FirstNum  int32 `protobuf:"varint,1,opt,name=firstNum" json:"firstNum,omitempty"`
	SecondNum int32 `protobuf:"varint,2,opt,name=secondNum" json:"secondNum,omitempty"`
}

func (m *Nums) Reset()                    { *m = Nums{} }
func (m *Nums) String() string            { return proto.CompactTextString(m) }
func (*Nums) ProtoMessage()               {}
func (*Nums) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Nums) GetFirstNum() int32 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

func (m *Nums) GetSecondNum() int32 {
	if m != nil {
		return m.SecondNum
	}
	return 0
}

type SumRequest struct {
	Nums *Nums `protobuf:"bytes,1,opt,name=nums" json:"nums,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SumRequest) GetNums() *Nums {
	if m != nil {
		return m.Nums
	}
	return nil
}

type SumResponse struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *SumResponse) Reset()                    { *m = SumResponse{} }
func (m *SumResponse) String() string            { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()               {}
func (*SumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *PrimeNumberDecompositionRequest) Reset()                    { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string            { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()               {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PrimeNumberDecompositionRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

func (m *PrimeNumberDecompositionResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAverageRequest struct {
	Num float32 `protobuf:"fixed32,1,opt,name=num" json:"num,omitempty"`
}

func (m *ComputeAverageRequest) Reset()                    { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string            { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()               {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ComputeAverageRequest) GetNum() float32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ComputeAverageResponse struct {
	Result float32 `protobuf:"fixed32,1,opt,name=result" json:"result,omitempty"`
}

func (m *ComputeAverageResponse) Reset()                    { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string            { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()               {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ComputeAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaximumRequest struct {
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *FindMaximumRequest) Reset()                    { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string            { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()               {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FindMaximumRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FindMaximumResponse struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *FindMaximumResponse) Reset()                    { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string            { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()               {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FindMaximumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type SquareRootRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *SquareRootRequest) Reset()                    { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string            { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()               {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot float64 `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot" json:"number_root,omitempty"`
}

func (m *SquareRootResponse) Reset()                    { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string            { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()               {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*Nums)(nil), "calculator.Nums")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calculator.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calculator.PrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SumService service

type SumServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (SumService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (SumService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (SumService_FindMaximumClient, error)
	//    error handling
	//    this RPC will throw an exception if the sent number is negative
	//    The error being sent is of type INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := grpc.Invoke(ctx, "/calculator.SumService/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (SumService_PrimeNumberDecompositionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SumService_serviceDesc.Streams[0], c.cc, "/calculator.SumService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type sumServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *sumServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (SumService_ComputeAverageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SumService_serviceDesc.Streams[1], c.cc, "/calculator.SumService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceComputeAverageClient{stream}
	return x, nil
}

type SumService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type sumServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *sumServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (SumService_FindMaximumClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SumService_serviceDesc.Streams[2], c.cc, "/calculator.SumService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceFindMaximumClient{stream}
	return x, nil
}

type SumService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type sumServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *sumServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := grpc.Invoke(ctx, "/calculator.SumService/SquareRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SumService service

type SumServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, SumService_PrimeNumberDecompositionServer) error
	ComputeAverage(SumService_ComputeAverageServer) error
	FindMaximum(SumService_FindMaximumServer) error
	//    error handling
	//    this RPC will throw an exception if the sent number is negative
	//    The error being sent is of type INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).PrimeNumberDecomposition(m, &sumServicePrimeNumberDecompositionServer{stream})
}

type SumService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type sumServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *sumServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SumService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).ComputeAverage(&sumServiceComputeAverageServer{stream})
}

type SumService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type sumServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *sumServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SumService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).FindMaximum(&sumServiceFindMaximumServer{stream})
}

type SumService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type sumServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *sumServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SumService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.SumService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _SumService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _SumService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _SumService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _SumService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}

func init() { proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0xb5, 0x80, 0x44, 0x6f, 0x0d, 0xc1, 0x6b, 0xac, 0xa4, 0x51, 0xc0, 0x89, 0x1a, 0x0c, 0x8a,
	0x04, 0x62, 0x62, 0x7c, 0xf2, 0x2b, 0xbe, 0x41, 0x4c, 0xeb, 0x93, 0x3e, 0x98, 0x52, 0xc6, 0x4d,
	0x13, 0xa6, 0x53, 0xe6, 0x83, 0xec, 0xfe, 0xbe, 0xfd, 0x63, 0x9b, 0xb6, 0x94, 0x0e, 0x1f, 0x5d,
	0xf6, 0x6d, 0xee, 0x3d, 0xe7, 0x9e, 0x33, 0x33, 0x67, 0x32, 0x30, 0x08, 0x83, 0x55, 0xa8, 0x57,
	0x81, 0xe2, 0xe2, 0x43, 0xb9, 0x4c, 0x16, 0x46, 0x31, 0x4a, 0x04, 0x57, 0x1c, 0xa1, 0xec, 0x90,
	0x2f, 0xd0, 0x98, 0x6b, 0x26, 0xd1, 0x85, 0x07, 0xff, 0x23, 0x21, 0xd5, 0x5c, 0xb3, 0x8e, 0xd5,
	0xb7, 0x06, 0xf7, 0xbd, 0x5d, 0x8d, 0xcf, 0xe1, 0xa1, 0xa4, 0x21, 0x8f, 0x97, 0x29, 0x58, 0xcb,
	0xc0, 0xb2, 0x41, 0x26, 0x00, 0xbe, 0x66, 0x1e, 0x5d, 0x6b, 0x2a, 0x15, 0xbe, 0x82, 0x46, 0xac,
	0x99, 0xcc, 0x34, 0xec, 0x49, 0x7b, 0x64, 0x98, 0xa7, 0x3e, 0x5e, 0x86, 0x92, 0xd7, 0x60, 0x67,
	0x33, 0x32, 0xe1, 0xb1, 0xa4, 0xe8, 0x40, 0x53, 0x50, 0xa9, 0x57, 0x6a, 0x6b, 0xbd, 0xad, 0xc8,
	0x14, 0x7a, 0xbf, 0x44, 0xc4, 0xe8, 0x5c, 0xb3, 0x05, 0x15, 0x3f, 0x68, 0xc8, 0x59, 0xc2, 0x65,
	0xa4, 0x22, 0x1e, 0x17, 0x7e, 0x6d, 0xa8, 0xc7, 0xbb, 0x2d, 0xa7, 0x4b, 0xf2, 0x19, 0xfa, 0xd5,
	0x43, 0x67, 0x0c, 0xdf, 0xc2, 0xd3, 0xef, 0x9c, 0x25, 0x5a, 0xd1, 0xaf, 0x1b, 0x2a, 0x82, 0x0b,
	0x7a, 0xc2, 0xa6, 0x96, 0xdb, 0x8c, 0xc1, 0x39, 0xa4, 0x9e, 0x14, 0xaf, 0xed, 0xc4, 0xdf, 0x00,
	0xfe, 0x8c, 0xe2, 0xe5, 0x2c, 0xb8, 0x8c, 0x58, 0x79, 0x61, 0xc7, 0x07, 0x78, 0x0f, 0x4f, 0xf6,
	0x78, 0x67, 0xf6, 0x3c, 0x84, 0xc7, 0xfe, 0x5a, 0x07, 0x82, 0x7a, 0x9c, 0xab, 0x42, 0xd5, 0x81,
	0x66, 0x9c, 0x9d, 0xbf, 0x20, 0xe7, 0x15, 0xf9, 0x08, 0x68, 0x92, 0xb7, 0xd2, 0x3d, 0xb0, 0x73,
	0xfc, 0x9f, 0xe0, 0x3c, 0xd7, 0xb7, 0x3c, 0xc8, 0x5b, 0x29, 0x71, 0x72, 0x5d, 0xcf, 0x42, 0xf6,
	0xa9, 0xd8, 0x44, 0x21, 0xc5, 0x4f, 0x50, 0xf7, 0x35, 0x43, 0xc7, 0x4c, 0xb7, 0x7c, 0x03, 0xee,
	0xb3, 0xa3, 0x7e, 0xee, 0x43, 0xee, 0xe1, 0x15, 0x74, 0xaa, 0xc2, 0xc1, 0xa1, 0x39, 0x76, 0x26,
	0x77, 0xf7, 0xdd, 0xdd, 0xc8, 0x85, 0xf1, 0xd8, 0xc2, 0xbf, 0xd0, 0xda, 0x0f, 0x0c, 0x5f, 0x9a,
	0x1a, 0x27, 0x73, 0x77, 0xc9, 0x6d, 0x94, 0x42, 0x7c, 0x60, 0xe1, 0x6f, 0xb0, 0x8d, 0xcc, 0xb0,
	0x6b, 0x8e, 0x1d, 0x87, 0xee, 0xf6, 0x2a, 0xf1, 0x52, 0x73, 0x6c, 0xe1, 0x0c, 0xa0, 0x4c, 0x0b,
	0x5f, 0xec, 0x5d, 0xeb, 0x61, 0xe4, 0x6e, 0xb7, 0x0a, 0x2e, 0x24, 0xbf, 0xb5, 0xfe, 0x3c, 0x32,
	0x3f, 0x86, 0x45, 0x33, 0xfb, 0x0e, 0xa6, 0x37, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xea, 0xb0,
	0x50, 0x3a, 0x04, 0x00, 0x00,
}
