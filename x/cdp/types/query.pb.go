// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cdp/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// this line is used by starport scaffolding # 3
type QueryGetCdpRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetCdpRequest) Reset()         { *m = QueryGetCdpRequest{} }
func (m *QueryGetCdpRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCdpRequest) ProtoMessage()    {}
func (*QueryGetCdpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d15b26084609686f, []int{0}
}
func (m *QueryGetCdpRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCdpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCdpRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCdpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCdpRequest.Merge(m, src)
}
func (m *QueryGetCdpRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCdpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCdpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCdpRequest proto.InternalMessageInfo

func (m *QueryGetCdpRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryGetCdpResponse struct {
	Cdp *Cdp `protobuf:"bytes,1,opt,name=Cdp,proto3" json:"Cdp,omitempty"`
}

func (m *QueryGetCdpResponse) Reset()         { *m = QueryGetCdpResponse{} }
func (m *QueryGetCdpResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCdpResponse) ProtoMessage()    {}
func (*QueryGetCdpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d15b26084609686f, []int{1}
}
func (m *QueryGetCdpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCdpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCdpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCdpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCdpResponse.Merge(m, src)
}
func (m *QueryGetCdpResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCdpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCdpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCdpResponse proto.InternalMessageInfo

func (m *QueryGetCdpResponse) GetCdp() *Cdp {
	if m != nil {
		return m.Cdp
	}
	return nil
}

type QueryAllCdpRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCdpRequest) Reset()         { *m = QueryAllCdpRequest{} }
func (m *QueryAllCdpRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCdpRequest) ProtoMessage()    {}
func (*QueryAllCdpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d15b26084609686f, []int{2}
}
func (m *QueryAllCdpRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCdpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCdpRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCdpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCdpRequest.Merge(m, src)
}
func (m *QueryAllCdpRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCdpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCdpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCdpRequest proto.InternalMessageInfo

func (m *QueryAllCdpRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCdpResponse struct {
	Cdp        []*Cdp              `protobuf:"bytes,1,rep,name=Cdp,proto3" json:"Cdp,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCdpResponse) Reset()         { *m = QueryAllCdpResponse{} }
func (m *QueryAllCdpResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCdpResponse) ProtoMessage()    {}
func (*QueryAllCdpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d15b26084609686f, []int{3}
}
func (m *QueryAllCdpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCdpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCdpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCdpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCdpResponse.Merge(m, src)
}
func (m *QueryAllCdpResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCdpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCdpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCdpResponse proto.InternalMessageInfo

func (m *QueryAllCdpResponse) GetCdp() []*Cdp {
	if m != nil {
		return m.Cdp
	}
	return nil
}

func (m *QueryAllCdpResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetCdpRequest)(nil), "lcnem.jpyx.cdp.QueryGetCdpRequest")
	proto.RegisterType((*QueryGetCdpResponse)(nil), "lcnem.jpyx.cdp.QueryGetCdpResponse")
	proto.RegisterType((*QueryAllCdpRequest)(nil), "lcnem.jpyx.cdp.QueryAllCdpRequest")
	proto.RegisterType((*QueryAllCdpResponse)(nil), "lcnem.jpyx.cdp.QueryAllCdpResponse")
}

func init() { proto.RegisterFile("cdp/query.proto", fileDescriptor_d15b26084609686f) }

var fileDescriptor_d15b26084609686f = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0xc6, 0x9b, 0x94, 0x2d, 0xec, 0x2c, 0xdb, 0x85, 0x29, 0x0b, 0x25, 0x4a, 0xa8, 0xf1, 0x2f,
	0x1e, 0x66, 0x68, 0x3d, 0xea, 0xa5, 0x16, 0xec, 0x55, 0x7b, 0x14, 0x2f, 0x49, 0x66, 0x88, 0x91,
	0x34, 0x33, 0xed, 0x4c, 0xc5, 0x22, 0x5e, 0x04, 0xef, 0x82, 0x5f, 0xca, 0x63, 0xc1, 0x8b, 0x47,
	0x69, 0xbd, 0xf8, 0x2d, 0x24, 0x33, 0x23, 0x26, 0xd5, 0xd6, 0x63, 0xc8, 0xf3, 0xbc, 0xbf, 0xdf,
	0xfb, 0x26, 0xe0, 0x5f, 0x48, 0x38, 0x1e, 0x8c, 0xe8, 0x70, 0x8c, 0xf8, 0x90, 0x49, 0x06, 0xab,
	0x49, 0x98, 0xd2, 0x3e, 0xba, 0xe0, 0xe3, 0x2b, 0x14, 0x12, 0xee, 0xac, 0x46, 0x8c, 0x45, 0x09,
	0xc5, 0x3e, 0x8f, 0xb1, 0x9f, 0xa6, 0x4c, 0xfa, 0x32, 0x66, 0xa9, 0xd0, 0x69, 0x67, 0x37, 0x64,
	0xa2, 0xcf, 0x04, 0x0e, 0x7c, 0x41, 0xf5, 0x18, 0x7c, 0xd9, 0x0c, 0xa8, 0xf4, 0x9b, 0x98, 0xfb,
	0x51, 0x9c, 0xaa, 0xb0, 0xc9, 0xfe, 0xcd, 0x50, 0x21, 0xe1, 0xfa, 0xd1, 0xdb, 0x00, 0xf0, 0x24,
	0x2b, 0x74, 0xa9, 0xec, 0x10, 0xde, 0xa3, 0x83, 0x11, 0x15, 0x12, 0x56, 0x81, 0x1d, 0x93, 0xba,
	0xd5, 0xb0, 0x76, 0x7e, 0xf7, 0xec, 0x98, 0x78, 0x07, 0xa0, 0x56, 0x48, 0x09, 0xce, 0x52, 0x41,
	0xe1, 0x26, 0x28, 0x77, 0x08, 0x57, 0xb9, 0x3f, 0xad, 0x1a, 0x2a, 0x3a, 0xa3, 0x2c, 0x99, 0xbd,
	0xf7, 0xce, 0x0c, 0xa3, 0x9d, 0x24, 0x39, 0xc6, 0x11, 0x00, 0x9f, 0x72, 0x66, 0xc6, 0x16, 0xd2,
	0x9b, 0xa0, 0x6c, 0x13, 0xa4, 0x0f, 0x62, 0x36, 0x41, 0xc7, 0x7e, 0x44, 0x4d, 0xb7, 0x97, 0x6b,
	0x7a, 0x77, 0x96, 0x91, 0xfb, 0x18, 0x3f, 0x2f, 0x57, 0x5e, 0x26, 0x07, 0xbb, 0x05, 0x0d, 0x5b,
	0x69, 0x6c, 0xff, 0xa8, 0xa1, 0x19, 0x79, 0x8f, 0xd6, 0x9b, 0x05, 0x7e, 0x29, 0x0f, 0xc8, 0x15,
	0x19, 0x7a, 0xf3, 0xcc, 0xaf, 0x87, 0x76, 0xd6, 0x97, 0x66, 0x34, 0xc5, 0x6b, 0xdc, 0x3e, 0xbd,
	0x3e, 0xd8, 0x0e, 0xac, 0x63, 0x15, 0xc6, 0x59, 0x18, 0x9b, 0xcf, 0x88, 0xaf, 0x63, 0x72, 0x03,
	0x53, 0x50, 0xe9, 0x10, 0xde, 0x4e, 0x92, 0x05, 0xd0, 0xc2, 0xe5, 0x17, 0x40, 0x8b, 0xe7, 0xf3,
	0x56, 0x14, 0xf4, 0x3f, 0xac, 0x7d, 0x03, 0x3d, 0xdc, 0x7f, 0x9c, 0xba, 0xd6, 0x64, 0xea, 0x5a,
	0x2f, 0x53, 0xd7, 0xba, 0x9f, 0xb9, 0xa5, 0xc9, 0xcc, 0x2d, 0x3d, 0xcf, 0xdc, 0xd2, 0xe9, 0x5a,
	0x14, 0xcb, 0xf3, 0x51, 0x80, 0x42, 0xd6, 0xcf, 0x17, 0x75, 0x55, 0x8e, 0x39, 0x15, 0x41, 0x45,
	0xfd, 0x79, 0x7b, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0x68, 0xe5, 0x94, 0xf5, 0x02, 0x00,
	0x00,
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
	// this line is used by starport scaffolding # 2
	Cdp(ctx context.Context, in *QueryGetCdpRequest, opts ...grpc.CallOption) (*QueryGetCdpResponse, error)
	CdpAll(ctx context.Context, in *QueryAllCdpRequest, opts ...grpc.CallOption) (*QueryAllCdpResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Cdp(ctx context.Context, in *QueryGetCdpRequest, opts ...grpc.CallOption) (*QueryGetCdpResponse, error) {
	out := new(QueryGetCdpResponse)
	err := c.cc.Invoke(ctx, "/lcnem.jpyx.cdp.Query/Cdp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CdpAll(ctx context.Context, in *QueryAllCdpRequest, opts ...grpc.CallOption) (*QueryAllCdpResponse, error) {
	out := new(QueryAllCdpResponse)
	err := c.cc.Invoke(ctx, "/lcnem.jpyx.cdp.Query/CdpAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Cdp(context.Context, *QueryGetCdpRequest) (*QueryGetCdpResponse, error)
	CdpAll(context.Context, *QueryAllCdpRequest) (*QueryAllCdpResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Cdp(ctx context.Context, req *QueryGetCdpRequest) (*QueryGetCdpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cdp not implemented")
}
func (*UnimplementedQueryServer) CdpAll(ctx context.Context, req *QueryAllCdpRequest) (*QueryAllCdpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CdpAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Cdp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCdpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Cdp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcnem.jpyx.cdp.Query/Cdp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Cdp(ctx, req.(*QueryGetCdpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CdpAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCdpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CdpAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcnem.jpyx.cdp.Query/CdpAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CdpAll(ctx, req.(*QueryAllCdpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lcnem.jpyx.cdp.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Cdp",
			Handler:    _Query_Cdp_Handler,
		},
		{
			MethodName: "CdpAll",
			Handler:    _Query_CdpAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cdp/query.proto",
}

func (m *QueryGetCdpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCdpRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCdpRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCdpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCdpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCdpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cdp != nil {
		{
			size, err := m.Cdp.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllCdpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCdpRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCdpRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllCdpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCdpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCdpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Cdp) > 0 {
		for iNdEx := len(m.Cdp) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cdp[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetCdpRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCdpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cdp != nil {
		l = m.Cdp.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCdpRequest) Size() (n int) {
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

func (m *QueryAllCdpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Cdp) > 0 {
		for _, e := range m.Cdp {
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
func (m *QueryGetCdpRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetCdpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCdpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetCdpResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetCdpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCdpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cdp", wireType)
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
			if m.Cdp == nil {
				m.Cdp = &Cdp{}
			}
			if err := m.Cdp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllCdpRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllCdpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCdpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllCdpResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllCdpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCdpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cdp", wireType)
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
			m.Cdp = append(m.Cdp, &Cdp{})
			if err := m.Cdp[len(m.Cdp)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
