// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paymentstream/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgStreamSend struct {
	Sender     string      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient  string      `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount     types.Coin  `protobuf:"bytes,3,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	EndTime    time.Time   `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	StreamType PaymentType `protobuf:"varint,5,opt,name=stream_type,json=streamType,proto3,enum=OmniFlix.paymentstream.paymentstream.PaymentType" json:"stream_type,omitempty" yaml:"stream_type"`
}

func (m *MsgStreamSend) Reset()         { *m = MsgStreamSend{} }
func (m *MsgStreamSend) String() string { return proto.CompactTextString(m) }
func (*MsgStreamSend) ProtoMessage()    {}
func (*MsgStreamSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2f80a4e0daf9187, []int{0}
}
func (m *MsgStreamSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStreamSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStreamSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStreamSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStreamSend.Merge(m, src)
}
func (m *MsgStreamSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgStreamSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStreamSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStreamSend proto.InternalMessageInfo

func (m *MsgStreamSend) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgStreamSend) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *MsgStreamSend) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *MsgStreamSend) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *MsgStreamSend) GetStreamType() PaymentType {
	if m != nil {
		return m.StreamType
	}
	return TypeDelayed
}

type MsgStreamSendResponse struct {
}

func (m *MsgStreamSendResponse) Reset()         { *m = MsgStreamSendResponse{} }
func (m *MsgStreamSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStreamSendResponse) ProtoMessage()    {}
func (*MsgStreamSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2f80a4e0daf9187, []int{1}
}
func (m *MsgStreamSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStreamSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStreamSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStreamSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStreamSendResponse.Merge(m, src)
}
func (m *MsgStreamSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStreamSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStreamSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStreamSendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStreamSend)(nil), "OmniFlix.paymentstream.paymentstream.MsgStreamSend")
	proto.RegisterType((*MsgStreamSendResponse)(nil), "OmniFlix.paymentstream.paymentstream.MsgStreamSendResponse")
}

func init() { proto.RegisterFile("paymentstream/tx.proto", fileDescriptor_a2f80a4e0daf9187) }

var fileDescriptor_a2f80a4e0daf9187 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0xb5, 0x10, 0xe8, 0x55, 0x80, 0x74, 0x82, 0x60, 0x02, 0xb2, 0x83, 0x85, 0x44, 0x96,
	0xde, 0x29, 0xe9, 0x82, 0x60, 0x0b, 0x12, 0x5b, 0x05, 0x72, 0x3b, 0xb1, 0x54, 0xfe, 0xf1, 0x30,
	0x07, 0xb9, 0x3b, 0xcb, 0x77, 0x41, 0xf1, 0xc2, 0xc6, 0xde, 0xbf, 0x83, 0xbf, 0xa4, 0x63, 0x47,
	0xa6, 0x14, 0x25, 0x13, 0x6b, 0xff, 0x02, 0x74, 0x3e, 0x5b, 0xc4, 0x9d, 0x2a, 0x26, 0xdf, 0x7b,
	0x7e, 0xef, 0x7d, 0xef, 0xfb, 0xbe, 0x3b, 0x3c, 0x28, 0xe2, 0x4a, 0x80, 0x34, 0xda, 0x94, 0x10,
	0x0b, 0x66, 0x96, 0xb4, 0x28, 0x95, 0x51, 0xe4, 0xc5, 0x7b, 0x21, 0xf9, 0xbb, 0x39, 0x5f, 0xd2,
	0x4e, 0x41, 0x37, 0x1a, 0xfa, 0xa9, 0xd2, 0x42, 0x69, 0x96, 0xc4, 0x1a, 0xd8, 0xb7, 0x49, 0x02,
	0x26, 0x9e, 0xb0, 0x54, 0x71, 0xe9, 0xa6, 0x0c, 0x1f, 0xe6, 0x2a, 0x57, 0xf5, 0x91, 0xd9, 0x53,
	0x93, 0x0d, 0x72, 0xa5, 0xf2, 0x39, 0xb0, 0x3a, 0x4a, 0x16, 0x9f, 0x98, 0xe1, 0x02, 0xb4, 0x89,
	0x45, 0xd1, 0x14, 0x3c, 0xef, 0x2e, 0x75, 0x6d, 0x03, 0x5b, 0x12, 0xfe, 0xd9, 0xc1, 0xf7, 0x8e,
	0x74, 0x7e, 0x5c, 0xe7, 0x8e, 0x41, 0x66, 0x64, 0x80, 0xfb, 0x1a, 0x64, 0x06, 0xa5, 0x87, 0x46,
	0x68, 0xbc, 0x17, 0x35, 0x11, 0x79, 0x86, 0xf7, 0x4a, 0x48, 0x79, 0xc1, 0x41, 0x1a, 0x6f, 0xa7,
	0xfe, 0xf5, 0x2f, 0x41, 0x12, 0xdc, 0x8f, 0x85, 0x5a, 0x48, 0xe3, 0xed, 0x8e, 0xd0, 0x78, 0x7f,
	0xfa, 0x84, 0x3a, 0x4a, 0xd4, 0x52, 0xa2, 0x0d, 0x25, 0xfa, 0x56, 0x71, 0x39, 0x63, 0xe7, 0xab,
	0xa0, 0xf7, 0xf3, 0x32, 0x78, 0x99, 0x73, 0xf3, 0x79, 0x91, 0xd0, 0x54, 0x09, 0xd6, 0xf0, 0x77,
	0x9f, 0x03, 0x9d, 0x7d, 0x65, 0xa6, 0x2a, 0x40, 0xd7, 0x0d, 0x51, 0x33, 0x99, 0x44, 0xf8, 0x2e,
	0xc8, 0xec, 0xd4, 0xb2, 0xf4, 0x6e, 0xd5, 0x28, 0x43, 0xea, 0x24, 0xa0, 0xad, 0x04, 0xf4, 0xa4,
	0x95, 0x60, 0xf6, 0xd4, 0xc2, 0x5c, 0xad, 0x82, 0x07, 0x55, 0x2c, 0xe6, 0xaf, 0xc3, 0xb6, 0x33,
	0x3c, 0xbb, 0x0c, 0x50, 0x74, 0x07, 0x64, 0x66, 0x4b, 0xc9, 0x17, 0xbc, 0xef, 0xf4, 0x38, 0xb5,
	0x80, 0xde, 0xed, 0x11, 0x1a, 0xdf, 0x9f, 0x4e, 0xe8, 0x4d, 0x5c, 0xa3, 0x1f, 0x5c, 0x74, 0x52,
	0x15, 0x30, 0x1b, 0x5c, 0xad, 0x02, 0xe2, 0x90, 0xb6, 0xe6, 0x85, 0x11, 0x76, 0x91, 0xad, 0x09,
	0x1f, 0xe3, 0x47, 0x1d, 0xa9, 0x23, 0xd0, 0x85, 0x92, 0x1a, 0xa6, 0x3f, 0x10, 0xde, 0x3d, 0xd2,
	0x39, 0xf9, 0x8e, 0xf1, 0x96, 0x11, 0x87, 0x37, 0xdb, 0xa2, 0x33, 0x72, 0xf8, 0xe6, 0x3f, 0x9a,
	0xda, 0x3d, 0x66, 0xd1, 0xf9, 0xda, 0x47, 0x17, 0x6b, 0x1f, 0xfd, 0x5e, 0xfb, 0xe8, 0x6c, 0xe3,
	0xf7, 0x2e, 0x36, 0x7e, 0xef, 0xd7, 0xc6, 0xef, 0x7d, 0x7c, 0xb5, 0xe5, 0x55, 0x0b, 0xd0, 0xde,
	0xa7, 0x83, 0xe6, 0x7a, 0x2d, 0xd9, 0xb5, 0x37, 0x60, 0x1d, 0x4c, 0xfa, 0xb5, 0x35, 0x87, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x59, 0x6c, 0x38, 0x1a, 0x21, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	StreamSend(ctx context.Context, in *MsgStreamSend, opts ...grpc.CallOption) (*MsgStreamSendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StreamSend(ctx context.Context, in *MsgStreamSend, opts ...grpc.CallOption) (*MsgStreamSendResponse, error) {
	out := new(MsgStreamSendResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.paymentstream.paymentstream.Msg/StreamSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	StreamSend(context.Context, *MsgStreamSend) (*MsgStreamSendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StreamSend(ctx context.Context, req *MsgStreamSend) (*MsgStreamSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamSend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StreamSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStreamSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StreamSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.paymentstream.paymentstream.Msg/StreamSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StreamSend(ctx, req.(*MsgStreamSend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OmniFlix.paymentstream.paymentstream.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamSend",
			Handler:    _Msg_StreamSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paymentstream/tx.proto",
}

func (m *MsgStreamSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStreamSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStreamSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StreamType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StreamType))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStreamSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStreamSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStreamSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStreamSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovTx(uint64(l))
	if m.StreamType != 0 {
		n += 1 + sovTx(uint64(m.StreamType))
	}
	return n
}

func (m *MsgStreamSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStreamSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStreamSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStreamSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamType", wireType)
			}
			m.StreamType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamType |= PaymentType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStreamSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStreamSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStreamSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
