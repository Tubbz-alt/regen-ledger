// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v1alpha1/query.proto

package data

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// QueryByCidRequest is the Query/ByCid request type.
type QueryByCidRequest struct {
	// cid is a Content Identifier for the data corresponding to the IPFS CID
	// specification: https://github.com/multiformats/cid.
	Cid []byte `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (m *QueryByCidRequest) Reset()         { *m = QueryByCidRequest{} }
func (m *QueryByCidRequest) String() string { return proto.CompactTextString(m) }
func (*QueryByCidRequest) ProtoMessage()    {}
func (*QueryByCidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22891d65265dd29b, []int{0}
}
func (m *QueryByCidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryByCidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryByCidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryByCidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryByCidRequest.Merge(m, src)
}
func (m *QueryByCidRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryByCidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryByCidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryByCidRequest proto.InternalMessageInfo

func (m *QueryByCidRequest) GetCid() []byte {
	if m != nil {
		return m.Cid
	}
	return nil
}

// QueryByCidResponse is the Query/ByCid response type.
type QueryByCidResponse struct {
	// timestamp is the timestamp of the block at which the data was anchored.
	Timestamp *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// signers are the addresses of the accounts which have signed the data.
	Signers []string `protobuf:"bytes,2,rep,name=signers,proto3" json:"signers,omitempty"`
	// content is the content of the data, if it was stored on-chain.
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *QueryByCidResponse) Reset()         { *m = QueryByCidResponse{} }
func (m *QueryByCidResponse) String() string { return proto.CompactTextString(m) }
func (*QueryByCidResponse) ProtoMessage()    {}
func (*QueryByCidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22891d65265dd29b, []int{1}
}
func (m *QueryByCidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryByCidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryByCidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryByCidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryByCidResponse.Merge(m, src)
}
func (m *QueryByCidResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryByCidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryByCidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryByCidResponse proto.InternalMessageInfo

func (m *QueryByCidResponse) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *QueryByCidResponse) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *QueryByCidResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// QueryBySignerRequest is the Query/BySigner request type.
type QueryBySignerRequest struct {
	// signer is the address of the signer to query by.
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// pagination is the PageRequest to use for pagination.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryBySignerRequest) Reset()         { *m = QueryBySignerRequest{} }
func (m *QueryBySignerRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBySignerRequest) ProtoMessage()    {}
func (*QueryBySignerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22891d65265dd29b, []int{2}
}
func (m *QueryBySignerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBySignerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBySignerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBySignerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBySignerRequest.Merge(m, src)
}
func (m *QueryBySignerRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBySignerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBySignerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBySignerRequest proto.InternalMessageInfo

func (m *QueryBySignerRequest) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *QueryBySignerRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryBySignerResponse is the Query/BySigner response type.
type QueryBySignerResponse struct {
	// cids are in the CIDs returned in this page of the query.
	Cids [][]byte `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	// pagination is the pagination PageResponse.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryBySignerResponse) Reset()         { *m = QueryBySignerResponse{} }
func (m *QueryBySignerResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBySignerResponse) ProtoMessage()    {}
func (*QueryBySignerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22891d65265dd29b, []int{3}
}
func (m *QueryBySignerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBySignerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBySignerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBySignerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBySignerResponse.Merge(m, src)
}
func (m *QueryBySignerResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBySignerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBySignerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBySignerResponse proto.InternalMessageInfo

func (m *QueryBySignerResponse) GetCids() [][]byte {
	if m != nil {
		return m.Cids
	}
	return nil
}

func (m *QueryBySignerResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryByCidRequest)(nil), "regen.data.v1alpha1.QueryByCidRequest")
	proto.RegisterType((*QueryByCidResponse)(nil), "regen.data.v1alpha1.QueryByCidResponse")
	proto.RegisterType((*QueryBySignerRequest)(nil), "regen.data.v1alpha1.QueryBySignerRequest")
	proto.RegisterType((*QueryBySignerResponse)(nil), "regen.data.v1alpha1.QueryBySignerResponse")
}

func init() { proto.RegisterFile("regen/data/v1alpha1/query.proto", fileDescriptor_22891d65265dd29b) }

var fileDescriptor_22891d65265dd29b = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x96, 0x0d, 0x6a, 0x76, 0x00, 0xf3, 0x43, 0x51, 0x84, 0xd2, 0x12, 0x69, 0x6b,
	0x99, 0x98, 0xad, 0x8e, 0x0b, 0xe7, 0x21, 0x8d, 0x2b, 0x04, 0x4e, 0xdc, 0x9c, 0xe4, 0x91, 0x59,
	0xb4, 0x76, 0x16, 0xbb, 0x85, 0x69, 0xda, 0x81, 0xfd, 0x05, 0x93, 0xb8, 0xf3, 0xf7, 0x70, 0x9c,
	0xc4, 0x85, 0x23, 0x6a, 0xf9, 0x43, 0x50, 0x6c, 0x87, 0x15, 0xa8, 0x28, 0x37, 0x3f, 0xf9, 0xf3,
	0xde, 0xf7, 0xfb, 0x7e, 0xe0, 0x5e, 0x05, 0x05, 0x48, 0x96, 0x73, 0xc3, 0xd9, 0x6c, 0xc4, 0xc7,
	0xe5, 0x11, 0x1f, 0xb1, 0xe3, 0x29, 0x54, 0x27, 0xb4, 0xac, 0x94, 0x51, 0xe4, 0x8e, 0x05, 0x68,
	0x0d, 0xd0, 0x06, 0x08, 0x1f, 0x14, 0x4a, 0x15, 0x63, 0x60, 0xbc, 0x14, 0x8c, 0x4b, 0xa9, 0x0c,
	0x37, 0x42, 0x49, 0xed, 0x52, 0xc2, 0x9e, 0xff, 0xb5, 0x51, 0x3a, 0x7d, 0xcb, 0x8c, 0x98, 0x80,
	0x36, 0x7c, 0x52, 0x7a, 0x60, 0x37, 0x53, 0x7a, 0xa2, 0x34, 0x4b, 0xb9, 0x06, 0x27, 0xc6, 0x66,
	0xa3, 0x14, 0x0c, 0x1f, 0xb1, 0x92, 0x17, 0x42, 0xda, 0x6a, 0x8e, 0x8d, 0xb7, 0xf1, 0xed, 0x97,
	0x35, 0x71, 0x70, 0xf2, 0x4c, 0xe4, 0x09, 0x1c, 0x4f, 0x41, 0x1b, 0x72, 0x0b, 0x77, 0x32, 0x91,
	0x07, 0xa8, 0x8f, 0x86, 0x5b, 0x49, 0xfd, 0x8c, 0xcf, 0x11, 0x26, 0xcb, 0x9c, 0x2e, 0x95, 0xd4,
	0x40, 0x9e, 0xe2, 0xee, 0x2f, 0x71, 0x8b, 0xdf, 0xdc, 0x0f, 0xa9, 0xb3, 0x47, 0x1b, 0x7b, 0xf4,
	0x75, 0x43, 0x24, 0x57, 0x30, 0x09, 0xf0, 0x75, 0x2d, 0x0a, 0x09, 0x95, 0x0e, 0xda, 0xfd, 0xce,
	0xb0, 0x9b, 0x34, 0x61, 0xfd, 0x93, 0x29, 0x69, 0x40, 0x9a, 0xa0, 0x63, 0x0d, 0x34, 0x61, 0x3c,
	0xc3, 0x77, 0xbd, 0x87, 0x57, 0x96, 0x6d, 0xec, 0xde, 0xc7, 0x9b, 0x2e, 0xd9, 0x5a, 0xe8, 0x26,
	0x3e, 0x22, 0x87, 0x18, 0x5f, 0xf5, 0x1b, 0xb4, 0xad, 0xbd, 0x1d, 0xea, 0x86, 0x43, 0xeb, 0xe1,
	0x50, 0xb7, 0x09, 0x3f, 0x1c, 0xfa, 0x82, 0x17, 0xe0, 0x6b, 0x26, 0x4b, 0x99, 0xb1, 0xc1, 0xf7,
	0xfe, 0xd0, 0xf5, 0xed, 0x13, 0x7c, 0x2d, 0x13, 0xb9, 0x0e, 0x50, 0xbf, 0x33, 0xdc, 0x4a, 0xec,
	0x9b, 0x3c, 0x5f, 0x21, 0x3a, 0x58, 0x2b, 0xea, 0x0a, 0x2e, 0xab, 0xee, 0x7f, 0x6e, 0xe3, 0x0d,
	0x2b, 0x4b, 0x3e, 0x22, 0xbc, 0x61, 0xe7, 0x4e, 0x76, 0xe8, 0x8a, 0x73, 0xa1, 0x7f, 0x2d, 0x30,
	0x1c, 0xac, 0xe5, 0x9c, 0x60, 0x3c, 0x38, 0xff, 0xfa, 0xe3, 0x53, 0xfb, 0x21, 0xe9, 0xb1, 0x55,
	0x87, 0x5a, 0x37, 0xc4, 0x4e, 0x33, 0x91, 0x9f, 0x91, 0x0b, 0x84, 0x6f, 0x34, 0xfd, 0x93, 0x47,
	0xff, 0x2a, 0xff, 0xdb, 0x6e, 0xc2, 0xdd, 0xff, 0x41, 0xbd, 0x99, 0x3d, 0x6b, 0x66, 0x40, 0xb6,
	0x57, 0x9a, 0xf1, 0xf7, 0xc1, 0x4e, 0xdd, 0xe3, 0xec, 0xe0, 0xf0, 0xcb, 0x3c, 0x42, 0x97, 0xf3,
	0x08, 0x7d, 0x9f, 0x47, 0xe8, 0x62, 0x11, 0xb5, 0x2e, 0x17, 0x51, 0xeb, 0xdb, 0x22, 0x6a, 0xbd,
	0x79, 0x5c, 0x08, 0x73, 0x34, 0x4d, 0x69, 0xa6, 0x26, 0xae, 0xd4, 0x9e, 0x04, 0xf3, 0x5e, 0x55,
	0xef, 0x7c, 0x34, 0x86, 0xbc, 0x80, 0x8a, 0x7d, 0xb0, 0x0a, 0xe9, 0xa6, 0xbd, 0xd4, 0x27, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x91, 0xf8, 0x91, 0xac, 0x03, 0x00, 0x00,
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
	// ByCid queries data based on its CID.
	ByCid(ctx context.Context, in *QueryByCidRequest, opts ...grpc.CallOption) (*QueryByCidResponse, error)
	// BySigner queries data based on signers.
	BySigner(ctx context.Context, in *QueryBySignerRequest, opts ...grpc.CallOption) (*QueryBySignerResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ByCid(ctx context.Context, in *QueryByCidRequest, opts ...grpc.CallOption) (*QueryByCidResponse, error) {
	out := new(QueryByCidResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Query/ByCid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BySigner(ctx context.Context, in *QueryBySignerRequest, opts ...grpc.CallOption) (*QueryBySignerResponse, error) {
	out := new(QueryBySignerResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Query/BySigner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ByCid queries data based on its CID.
	ByCid(context.Context, *QueryByCidRequest) (*QueryByCidResponse, error)
	// BySigner queries data based on signers.
	BySigner(context.Context, *QueryBySignerRequest) (*QueryBySignerResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ByCid(ctx context.Context, req *QueryByCidRequest) (*QueryByCidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByCid not implemented")
}
func (*UnimplementedQueryServer) BySigner(ctx context.Context, req *QueryBySignerRequest) (*QueryBySignerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BySigner not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ByCid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByCidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ByCid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Query/ByCid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ByCid(ctx, req.(*QueryByCidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BySigner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBySignerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BySigner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Query/BySigner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BySigner(ctx, req.(*QueryBySignerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "regen.data.v1alpha1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByCid",
			Handler:    _Query_ByCid_Handler,
		},
		{
			MethodName: "BySigner",
			Handler:    _Query_BySigner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/data/v1alpha1/query.proto",
}

func (m *QueryByCidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryByCidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryByCidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryByCidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryByCidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryByCidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryBySignerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBySignerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBySignerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBySignerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBySignerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBySignerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Cids) > 0 {
		for iNdEx := len(m.Cids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Cids[iNdEx])
			copy(dAtA[i:], m.Cids[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Cids[iNdEx])))
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
func (m *QueryByCidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryByCidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBySignerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBySignerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Cids) > 0 {
		for _, b := range m.Cids {
			l = len(b)
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
func (m *QueryByCidRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryByCidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryByCidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = append(m.Cid[:0], dAtA[iNdEx:postIndex]...)
			if m.Cid == nil {
				m.Cid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryByCidResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryByCidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryByCidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
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
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryBySignerRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBySignerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBySignerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
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
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryBySignerResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBySignerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBySignerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cids", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cids = append(m.Cids, make([]byte, postIndex-iNdEx))
			copy(m.Cids[len(m.Cids)-1], dAtA[iNdEx:postIndex])
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
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
