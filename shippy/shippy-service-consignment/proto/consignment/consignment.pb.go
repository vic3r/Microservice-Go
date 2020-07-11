// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package consignment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "consignment.Response")
	proto.RegisterType((*Container)(nil), "consignment.Container")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x24, 0x7d, 0x67, 0x53, 0x51, 0x61, 0x89, 0xd6, 0x82, 0x03, 0x51, 0xb8, 0xf4, 0x54, 0xa4,
	0x22, 0x71, 0x40, 0xdc, 0x22, 0x51, 0xf5, 0xea, 0x7e, 0x00, 0x2a, 0xc9, 0x2a, 0x5d, 0x89, 0xda,
	0xc1, 0x76, 0xcb, 0xdf, 0xf4, 0x0b, 0xf8, 0x48, 0x54, 0xa7, 0xa1, 0x2e, 0xa8, 0x37, 0xcf, 0xce,
	0x38, 0x99, 0x99, 0x35, 0xdc, 0x97, 0x5a, 0x59, 0xf5, 0x90, 0x29, 0x69, 0xa8, 0x90, 0x6b, 0x94,
	0xd6, 0x3f, 0x4f, 0x1c, 0xcb, 0x22, 0x6f, 0x94, 0x7c, 0x07, 0x10, 0xa5, 0x47, 0xcc, 0x2e, 0xa1,
	0x41, 0x39, 0x0f, 0xe2, 0x60, 0x1c, 0x8a, 0x06, 0xe5, 0x2c, 0x86, 0x28, 0x47, 0x93, 0x69, 0x2a,
	0x2d, 0x29, 0xc9, 0x1b, 0x8e, 0xf0, 0x47, 0x6c, 0x08, 0x9d, 0x2f, 0xa4, 0x62, 0x65, 0x79, 0x33,
	0x0e, 0xc6, 0x6d, 0x71, 0x40, 0xec, 0x09, 0x20, 0x53, 0xd2, 0x2e, 0x49, 0xa2, 0x36, 0xbc, 0x15,
	0x37, 0xc7, 0xd1, 0x74, 0x38, 0xf1, 0xed, 0xa4, 0x35, 0x2d, 0x3c, 0x25, 0xbb, 0x85, 0x70, 0x8b,
	0xc6, 0xe0, 0xc7, 0x1b, 0xe5, 0xbc, 0xed, 0xfe, 0xd7, 0xab, 0x06, 0xf3, 0x3c, 0xe9, 0x03, 0xcc,
	0xd0, 0x0a, 0xfc, 0xdc, 0xa0, 0xb1, 0xc9, 0x2e, 0x80, 0x9e, 0x40, 0x53, 0x2a, 0x69, 0x90, 0x71,
	0xe8, 0x66, 0x1a, 0x97, 0x16, 0x2b, 0xfb, 0x3d, 0x51, 0x43, 0xf6, 0x0c, 0x7e, 0x64, 0x97, 0x21,
	0x9a, 0xf2, 0xbf, 0x56, 0xea, 0xb3, 0xf0, 0xc5, 0xec, 0x05, 0xfa, 0x1e, 0x34, 0xbc, 0xe9, 0x72,
	0x9c, 0xbf, 0x7c, 0xa2, 0x4e, 0xd6, 0x10, 0xfe, 0x86, 0xfc, 0x57, 0xed, 0x1d, 0x44, 0xd9, 0xc6,
	0x58, 0xb5, 0x46, 0xbd, 0x8f, 0x5a, 0x55, 0x0b, 0xf5, 0x68, 0x9e, 0xef, 0x9b, 0x55, 0x9a, 0x0a,
	0x92, 0xae, 0xd9, 0x50, 0x1c, 0x10, 0x1b, 0x41, 0x77, 0x63, 0xaa, 0x4b, 0xad, 0x8a, 0xd8, 0xc3,
	0x79, 0x3e, 0xdd, 0x05, 0x30, 0x58, 0xac, 0xa8, 0x2c, 0x49, 0x16, 0x0b, 0xd4, 0x5b, 0xca, 0x90,
	0xbd, 0xc2, 0x55, 0xea, 0x7a, 0xf0, 0xb7, 0x7c, 0xd6, 0xff, 0xcd, 0xf5, 0x09, 0x53, 0x97, 0x9b,
	0x5c, 0xb0, 0x14, 0x06, 0x33, 0xb4, 0x9e, 0xd4, 0xb0, 0xd1, 0x89, 0xf6, 0xb8, 0x97, 0xb3, 0x1f,
	0x79, 0xef, 0xb8, 0x17, 0xf8, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x51, 0xbe, 0xda, 0xa5, 0xa8,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/GetConsignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
	GetConsignments(context.Context, *GetRequest) (*Response, error)
}

// UnimplementedShippingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShippingServiceServer struct {
}

func (*UnimplementedShippingServiceServer) CreateConsignment(ctx context.Context, req *Consignment) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsignment not implemented")
}
func (*UnimplementedShippingServiceServer) GetConsignments(ctx context.Context, req *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsignments not implemented")
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetConsignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetConsignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/GetConsignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetConsignments(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
		{
			MethodName: "GetConsignments",
			Handler:    _ShippingService_GetConsignments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}
