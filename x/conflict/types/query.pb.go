// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lava/conflict/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9a5ee7cc69cc24, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9a5ee7cc69cc24, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetConflictVoteRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetConflictVoteRequest) Reset()         { *m = QueryGetConflictVoteRequest{} }
func (m *QueryGetConflictVoteRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetConflictVoteRequest) ProtoMessage()    {}
func (*QueryGetConflictVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9a5ee7cc69cc24, []int{2}
}
func (m *QueryGetConflictVoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetConflictVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetConflictVoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetConflictVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetConflictVoteRequest.Merge(m, src)
}
func (m *QueryGetConflictVoteRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetConflictVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetConflictVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetConflictVoteRequest proto.InternalMessageInfo

func (m *QueryGetConflictVoteRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetConflictVoteResponse struct {
	ConflictVote ConflictVote `protobuf:"bytes,1,opt,name=conflictVote,proto3" json:"conflictVote"`
}

func (m *QueryGetConflictVoteResponse) Reset()         { *m = QueryGetConflictVoteResponse{} }
func (m *QueryGetConflictVoteResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetConflictVoteResponse) ProtoMessage()    {}
func (*QueryGetConflictVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9a5ee7cc69cc24, []int{3}
}
func (m *QueryGetConflictVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetConflictVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetConflictVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetConflictVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetConflictVoteResponse.Merge(m, src)
}
func (m *QueryGetConflictVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetConflictVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetConflictVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetConflictVoteResponse proto.InternalMessageInfo

func (m *QueryGetConflictVoteResponse) GetConflictVote() ConflictVote {
	if m != nil {
		return m.ConflictVote
	}
	return ConflictVote{}
}

type QueryAllConflictVoteRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllConflictVoteRequest) Reset()         { *m = QueryAllConflictVoteRequest{} }
func (m *QueryAllConflictVoteRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllConflictVoteRequest) ProtoMessage()    {}
func (*QueryAllConflictVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9a5ee7cc69cc24, []int{4}
}
func (m *QueryAllConflictVoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllConflictVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllConflictVoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllConflictVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllConflictVoteRequest.Merge(m, src)
}
func (m *QueryAllConflictVoteRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllConflictVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllConflictVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllConflictVoteRequest proto.InternalMessageInfo

func (m *QueryAllConflictVoteRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllConflictVoteResponse struct {
	ConflictVote []ConflictVote      `protobuf:"bytes,1,rep,name=conflictVote,proto3" json:"conflictVote"`
	Pagination   *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllConflictVoteResponse) Reset()         { *m = QueryAllConflictVoteResponse{} }
func (m *QueryAllConflictVoteResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllConflictVoteResponse) ProtoMessage()    {}
func (*QueryAllConflictVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9a5ee7cc69cc24, []int{5}
}
func (m *QueryAllConflictVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllConflictVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllConflictVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllConflictVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllConflictVoteResponse.Merge(m, src)
}
func (m *QueryAllConflictVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllConflictVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllConflictVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllConflictVoteResponse proto.InternalMessageInfo

func (m *QueryAllConflictVoteResponse) GetConflictVote() []ConflictVote {
	if m != nil {
		return m.ConflictVote
	}
	return nil
}

func (m *QueryAllConflictVoteResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "lavanet.lava.conflict.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "lavanet.lava.conflict.QueryParamsResponse")
	proto.RegisterType((*QueryGetConflictVoteRequest)(nil), "lavanet.lava.conflict.QueryGetConflictVoteRequest")
	proto.RegisterType((*QueryGetConflictVoteResponse)(nil), "lavanet.lava.conflict.QueryGetConflictVoteResponse")
	proto.RegisterType((*QueryAllConflictVoteRequest)(nil), "lavanet.lava.conflict.QueryAllConflictVoteRequest")
	proto.RegisterType((*QueryAllConflictVoteResponse)(nil), "lavanet.lava.conflict.QueryAllConflictVoteResponse")
}

func init() { proto.RegisterFile("lava/conflict/query.proto", fileDescriptor_ea9a5ee7cc69cc24) }

var fileDescriptor_ea9a5ee7cc69cc24 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x8d, 0x55, 0xc2, 0x4c, 0x42, 0x32, 0x45, 0x82, 0xb0, 0x65, 0x60, 0x60, 0x8c,
	0x69, 0xb2, 0xb5, 0x96, 0x1b, 0xa7, 0x15, 0x89, 0x9d, 0x90, 0x46, 0x0e, 0x1c, 0xb8, 0x20, 0x27,
	0x98, 0x10, 0x29, 0xb1, 0xb3, 0xda, 0xad, 0x36, 0x21, 0x2e, 0x1c, 0x38, 0x23, 0xf1, 0x25, 0xb8,
	0x70, 0xe5, 0x33, 0xec, 0x38, 0x89, 0x0b, 0x27, 0x84, 0x5a, 0x6e, 0x7c, 0x09, 0x14, 0xdb, 0xd1,
	0x12, 0x2d, 0x29, 0x45, 0x3b, 0xd5, 0x75, 0xde, 0xff, 0xff, 0xff, 0xbd, 0xbe, 0xd7, 0xc0, 0x9b,
	0x29, 0x9b, 0x30, 0x1a, 0x49, 0xf1, 0x26, 0x4d, 0x22, 0x4d, 0x0f, 0xc7, 0x7c, 0x74, 0x4c, 0xf2,
	0x91, 0xd4, 0x12, 0x5d, 0x2f, 0x1e, 0x09, 0xae, 0x49, 0xf1, 0x49, 0xca, 0x12, 0x6f, 0x2d, 0x96,
	0x32, 0x4e, 0x39, 0x65, 0x79, 0x42, 0x99, 0x10, 0x52, 0x33, 0x9d, 0x48, 0xa1, 0xac, 0xc8, 0xdb,
	0x8e, 0xa4, 0xca, 0xa4, 0xa2, 0x21, 0x53, 0xdc, 0xba, 0xd1, 0xc9, 0x6e, 0xc8, 0x35, 0xdb, 0xa5,
	0x39, 0x8b, 0x13, 0x61, 0x8a, 0x5d, 0xad, 0x57, 0xcf, 0xce, 0xd9, 0x88, 0x65, 0xa5, 0xcf, 0x9d,
	0xfa, 0xb3, 0xf2, 0xf0, 0x6a, 0x22, 0x35, 0x77, 0x25, 0xbd, 0x58, 0xc6, 0xd2, 0x1c, 0x69, 0x71,
	0xb2, 0xb7, 0xb8, 0x07, 0xd1, 0xf3, 0x22, 0xf6, 0xc0, 0xb8, 0x05, 0xfc, 0x70, 0xcc, 0x95, 0xc6,
	0x01, 0xbc, 0x56, 0xbb, 0x55, 0xb9, 0x14, 0x8a, 0xa3, 0xc7, 0xb0, 0x6b, 0x53, 0x6f, 0x80, 0xdb,
	0x60, 0xeb, 0x4a, 0x7f, 0x9d, 0x34, 0xf6, 0x4c, 0xac, 0x6c, 0x78, 0xe9, 0xe4, 0xe7, 0x46, 0x27,
	0x70, 0x12, 0x3c, 0x80, 0xb7, 0x8c, 0xe7, 0x3e, 0xd7, 0x4f, 0x5c, 0xe1, 0x0b, 0xa9, 0xb9, 0x8b,
	0x44, 0x3d, 0xb8, 0x92, 0x88, 0xd7, 0xfc, 0xc8, 0x58, 0x5f, 0x0e, 0xec, 0x17, 0x9c, 0xc1, 0xb5,
	0x66, 0x91, 0x23, 0x7a, 0x06, 0x57, 0xa3, 0xca, 0xbd, 0xe3, 0xba, 0xdb, 0xc2, 0x55, 0xb5, 0x70,
	0x74, 0x35, 0x39, 0xe6, 0x8e, 0x71, 0x2f, 0x4d, 0x9b, 0x18, 0x9f, 0x42, 0x78, 0x36, 0x15, 0x97,
	0xb5, 0x49, 0xec, 0x08, 0x49, 0x31, 0x42, 0x62, 0x17, 0xc2, 0x8d, 0x90, 0x1c, 0xb0, 0xb8, 0xd4,
	0x06, 0x15, 0x25, 0xfe, 0x06, 0x5c, 0x5b, 0xe7, 0x72, 0x5a, 0xdb, 0x5a, 0xbe, 0x40, 0x5b, 0x68,
	0xbf, 0xc6, 0xbd, 0x64, 0xb8, 0x1f, 0xfc, 0x93, 0xdb, 0xb2, 0x54, 0xc1, 0xfb, 0x7f, 0x96, 0xe1,
	0x8a, 0x01, 0x47, 0x1f, 0x01, 0xec, 0xda, 0x31, 0xa3, 0x87, 0x2d, 0x58, 0xe7, 0xf7, 0xca, 0xdb,
	0x5e, 0xa4, 0xd4, 0xe6, 0xe2, 0xfb, 0x1f, 0xbe, 0xff, 0xfe, 0xbc, 0xb4, 0x81, 0xd6, 0xa9, 0xd3,
	0xd0, 0xa6, 0xfd, 0x47, 0x5f, 0x01, 0x5c, 0xad, 0xfe, 0x00, 0xa8, 0x3f, 0x2f, 0xa3, 0x79, 0xf9,
	0xbc, 0xc1, 0x7f, 0x69, 0x1c, 0xe0, 0x23, 0x03, 0x48, 0xd0, 0x4e, 0x0b, 0x60, 0xed, 0x4f, 0x48,
	0xdf, 0x99, 0x85, 0x7e, 0x8f, 0xbe, 0x00, 0x78, 0xb5, 0x6a, 0xb7, 0x97, 0xa6, 0xf3, 0x91, 0x9b,
	0x77, 0x71, 0x3e, 0x72, 0xcb, 0x5e, 0xe1, 0x1d, 0x83, 0xbc, 0x89, 0xee, 0x2d, 0x82, 0x3c, 0x1c,
	0x9e, 0x4c, 0x7d, 0x70, 0x3a, 0xf5, 0xc1, 0xaf, 0xa9, 0x0f, 0x3e, 0xcd, 0xfc, 0xce, 0xe9, 0xcc,
	0xef, 0xfc, 0x98, 0xf9, 0x9d, 0x97, 0x5b, 0x71, 0xa2, 0xdf, 0x8e, 0x43, 0x12, 0xc9, 0xac, 0xee,
	0x74, 0x74, 0xe6, 0xa5, 0x8f, 0x73, 0xae, 0xc2, 0xae, 0x79, 0xcd, 0x0c, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x49, 0xf0, 0x6a, 0xd6, 0x39, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a ConflictVote by index.
	ConflictVote(ctx context.Context, in *QueryGetConflictVoteRequest, opts ...grpc.CallOption) (*QueryGetConflictVoteResponse, error)
	// Queries a list of ConflictVote items.
	ConflictVoteAll(ctx context.Context, in *QueryAllConflictVoteRequest, opts ...grpc.CallOption) (*QueryAllConflictVoteResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.conflict.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConflictVote(ctx context.Context, in *QueryGetConflictVoteRequest, opts ...grpc.CallOption) (*QueryGetConflictVoteResponse, error) {
	out := new(QueryGetConflictVoteResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.conflict.Query/ConflictVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConflictVoteAll(ctx context.Context, in *QueryAllConflictVoteRequest, opts ...grpc.CallOption) (*QueryAllConflictVoteResponse, error) {
	out := new(QueryAllConflictVoteResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.conflict.Query/ConflictVoteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a ConflictVote by index.
	ConflictVote(context.Context, *QueryGetConflictVoteRequest) (*QueryGetConflictVoteResponse, error)
	// Queries a list of ConflictVote items.
	ConflictVoteAll(context.Context, *QueryAllConflictVoteRequest) (*QueryAllConflictVoteResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ConflictVote(ctx context.Context, req *QueryGetConflictVoteRequest) (*QueryGetConflictVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConflictVote not implemented")
}
func (*UnimplementedQueryServer) ConflictVoteAll(ctx context.Context, req *QueryAllConflictVoteRequest) (*QueryAllConflictVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConflictVoteAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.conflict.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConflictVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetConflictVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConflictVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.conflict.Query/ConflictVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConflictVote(ctx, req.(*QueryGetConflictVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConflictVoteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllConflictVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConflictVoteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.conflict.Query/ConflictVoteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConflictVoteAll(ctx, req.(*QueryAllConflictVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lavanet.lava.conflict.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ConflictVote",
			Handler:    _Query_ConflictVote_Handler,
		},
		{
			MethodName: "ConflictVoteAll",
			Handler:    _Query_ConflictVoteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lava/conflict/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetConflictVoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetConflictVoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetConflictVoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetConflictVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetConflictVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetConflictVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ConflictVote.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllConflictVoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllConflictVoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllConflictVoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllConflictVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllConflictVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllConflictVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConflictVote) > 0 {
		for iNdEx := len(m.ConflictVote) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConflictVote[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetConflictVoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetConflictVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ConflictVote.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllConflictVoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllConflictVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ConflictVote) > 0 {
		for _, e := range m.ConflictVote {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetConflictVoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetConflictVoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetConflictVoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetConflictVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetConflictVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetConflictVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictVote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConflictVote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllConflictVoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllConflictVoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllConflictVoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllConflictVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllConflictVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllConflictVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictVote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConflictVote = append(m.ConflictVote, ConflictVote{})
			if err := m.ConflictVote[len(m.ConflictVote)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
