// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paymentstream/query.proto

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

type QueryAllPaymentStreamRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPaymentStreamRequest) Reset()         { *m = QueryAllPaymentStreamRequest{} }
func (m *QueryAllPaymentStreamRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPaymentStreamRequest) ProtoMessage()    {}
func (*QueryAllPaymentStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa2c1661fb88ff7, []int{0}
}
func (m *QueryAllPaymentStreamRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPaymentStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPaymentStreamRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPaymentStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPaymentStreamRequest.Merge(m, src)
}
func (m *QueryAllPaymentStreamRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPaymentStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPaymentStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPaymentStreamRequest proto.InternalMessageInfo

func (m *QueryAllPaymentStreamRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPaymentStreamResponse struct {
	PaymentStreams []PaymentStream     `protobuf:"bytes,1,rep,name=payment_streams,json=paymentStreams,proto3" json:"payment_streams" yaml:"payment_streams"`
	Pagination     *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPaymentStreamResponse) Reset()         { *m = QueryAllPaymentStreamResponse{} }
func (m *QueryAllPaymentStreamResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPaymentStreamResponse) ProtoMessage()    {}
func (*QueryAllPaymentStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fa2c1661fb88ff7, []int{1}
}
func (m *QueryAllPaymentStreamResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPaymentStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPaymentStreamResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPaymentStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPaymentStreamResponse.Merge(m, src)
}
func (m *QueryAllPaymentStreamResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPaymentStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPaymentStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPaymentStreamResponse proto.InternalMessageInfo

func (m *QueryAllPaymentStreamResponse) GetPaymentStreams() []PaymentStream {
	if m != nil {
		return m.PaymentStreams
	}
	return nil
}

func (m *QueryAllPaymentStreamResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllPaymentStreamRequest)(nil), "OmniFlix.paymentstream.paymentstream.QueryAllPaymentStreamRequest")
	proto.RegisterType((*QueryAllPaymentStreamResponse)(nil), "OmniFlix.paymentstream.paymentstream.QueryAllPaymentStreamResponse")
}

func init() { proto.RegisterFile("paymentstream/query.proto", fileDescriptor_5fa2c1661fb88ff7) }

var fileDescriptor_5fa2c1661fb88ff7 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0x29, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xf1, 0xcf, 0xcd, 0xcb, 0x74, 0xcb, 0xc9, 0xac, 0xd0,
	0x43, 0x51, 0x83, 0xca, 0x93, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8,
	0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x98, 0x21, 0xa5,
	0x95, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x94, 0x58, 0x9c, 0x0a, 0x31, 0x5c, 0xbf, 0xcc,
	0x30, 0x29, 0xb5, 0x24, 0xd1, 0x50, 0xbf, 0x20, 0x31, 0x3d, 0x33, 0x0f, 0xac, 0x18, 0xaa, 0x56,
	0x11, 0xd5, 0x29, 0x68, 0x96, 0x82, 0x95, 0x88, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20,
	0x16, 0x44, 0x54, 0x29, 0x8d, 0x4b, 0x26, 0x10, 0x64, 0xb4, 0x63, 0x4e, 0x4e, 0x00, 0x44, 0x53,
	0x30, 0x58, 0x53, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x1b, 0x17, 0x17, 0xc2, 0x32,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x35, 0x3d, 0x88, 0xcb, 0xf4, 0x40, 0x2e, 0xd3, 0x83,
	0x78, 0x1b, 0xea, 0x32, 0xbd, 0x80, 0xc4, 0xf4, 0x54, 0xa8, 0xde, 0x20, 0x24, 0x9d, 0x4a, 0xcf,
	0x18, 0xb9, 0x64, 0x71, 0x58, 0x54, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x54, 0xc3, 0xc5, 0x0f,
	0x75, 0x76, 0x3c, 0xc4, 0xdd, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xc6, 0x7a, 0xc4,
	0x04, 0xa6, 0x1e, 0x8a, 0xa9, 0x4e, 0x72, 0x27, 0xee, 0xc9, 0x33, 0x7c, 0xba, 0x27, 0x2f, 0x56,
	0x99, 0x98, 0x9b, 0x63, 0xa5, 0x84, 0x66, 0xb2, 0x52, 0x10, 0x5f, 0x01, 0xb2, 0xf2, 0x62, 0x21,
	0x77, 0x14, 0x7f, 0x32, 0x81, 0xfd, 0xa9, 0x4e, 0xd0, 0x9f, 0x10, 0xa7, 0x23, 0x7b, 0xd4, 0xe8,
	0x30, 0x23, 0x17, 0x2b, 0xd8, 0xa3, 0x42, 0x3b, 0x19, 0xb9, 0x04, 0x51, 0x1c, 0x55, 0xec, 0x98,
	0x93, 0x23, 0xe4, 0x44, 0x9c, 0x6f, 0xf0, 0x45, 0x8a, 0x94, 0x33, 0x45, 0x66, 0x40, 0x1c, 0xad,
	0x24, 0xdd, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x51, 0x21, 0x61, 0x7d, 0x88, 0xf2, 0xcc, 0xbc, 0x74,
	0x5d, 0x98, 0x7e, 0xa7, 0xa0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2,
	0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xb9, 0x02, 0x96, 0xde,
	0x74, 0xa1, 0xc9, 0xaf, 0x02, 0x35, 0x01, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81,
	0x53, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x83, 0xdb, 0x3f, 0xeb, 0x37, 0x03, 0x00, 0x00,
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
	PaymentStreamsAll(ctx context.Context, in *QueryAllPaymentStreamRequest, opts ...grpc.CallOption) (*QueryAllPaymentStreamResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) PaymentStreamsAll(ctx context.Context, in *QueryAllPaymentStreamRequest, opts ...grpc.CallOption) (*QueryAllPaymentStreamResponse, error) {
	out := new(QueryAllPaymentStreamResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.paymentstream.paymentstream.Query/PaymentStreamsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	PaymentStreamsAll(context.Context, *QueryAllPaymentStreamRequest) (*QueryAllPaymentStreamResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) PaymentStreamsAll(ctx context.Context, req *QueryAllPaymentStreamRequest) (*QueryAllPaymentStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentStreamsAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_PaymentStreamsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPaymentStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PaymentStreamsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.paymentstream.paymentstream.Query/PaymentStreamsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PaymentStreamsAll(ctx, req.(*QueryAllPaymentStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OmniFlix.paymentstream.paymentstream.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PaymentStreamsAll",
			Handler:    _Query_PaymentStreamsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paymentstream/query.proto",
}

func (m *QueryAllPaymentStreamRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPaymentStreamRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPaymentStreamRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllPaymentStreamResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPaymentStreamResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPaymentStreamResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.PaymentStreams) > 0 {
		for iNdEx := len(m.PaymentStreams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PaymentStreams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryAllPaymentStreamRequest) Size() (n int) {
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

func (m *QueryAllPaymentStreamResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PaymentStreams) > 0 {
		for _, e := range m.PaymentStreams {
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
func (m *QueryAllPaymentStreamRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPaymentStreamRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPaymentStreamRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllPaymentStreamResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPaymentStreamResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPaymentStreamResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentStreams", wireType)
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
			m.PaymentStreams = append(m.PaymentStreams, PaymentStream{})
			if err := m.PaymentStreams[len(m.PaymentStreams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
