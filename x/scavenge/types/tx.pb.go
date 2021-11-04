// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSubmitScavenge struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	SolutionHash string `protobuf:"bytes,2,opt,name=solutionHash,proto3" json:"solutionHash,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reward       string `protobuf:"bytes,4,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (m *MsgSubmitScavenge) Reset()         { *m = MsgSubmitScavenge{} }
func (m *MsgSubmitScavenge) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitScavenge) ProtoMessage()    {}
func (*MsgSubmitScavenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ec6f1544538068, []int{0}
}
func (m *MsgSubmitScavenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitScavenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitScavenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitScavenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitScavenge.Merge(m, src)
}
func (m *MsgSubmitScavenge) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitScavenge) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitScavenge.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitScavenge proto.InternalMessageInfo

func (m *MsgSubmitScavenge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitScavenge) GetSolutionHash() string {
	if m != nil {
		return m.SolutionHash
	}
	return ""
}

func (m *MsgSubmitScavenge) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgSubmitScavenge) GetReward() string {
	if m != nil {
		return m.Reward
	}
	return ""
}

type MsgSubmitScavengeResponse struct {
}

func (m *MsgSubmitScavengeResponse) Reset()         { *m = MsgSubmitScavengeResponse{} }
func (m *MsgSubmitScavengeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitScavengeResponse) ProtoMessage()    {}
func (*MsgSubmitScavengeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ec6f1544538068, []int{1}
}
func (m *MsgSubmitScavengeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitScavengeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitScavengeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitScavengeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitScavengeResponse.Merge(m, src)
}
func (m *MsgSubmitScavengeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitScavengeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitScavengeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitScavengeResponse proto.InternalMessageInfo

type MsgCommitSolution struct {
	Creator               string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	SolutionHash          string `protobuf:"bytes,2,opt,name=solutionHash,proto3" json:"solutionHash,omitempty"`
	SolutionScavengerHash string `protobuf:"bytes,3,opt,name=solutionScavengerHash,proto3" json:"solutionScavengerHash,omitempty"`
}

func (m *MsgCommitSolution) Reset()         { *m = MsgCommitSolution{} }
func (m *MsgCommitSolution) String() string { return proto.CompactTextString(m) }
func (*MsgCommitSolution) ProtoMessage()    {}
func (*MsgCommitSolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ec6f1544538068, []int{2}
}
func (m *MsgCommitSolution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitSolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitSolution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitSolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitSolution.Merge(m, src)
}
func (m *MsgCommitSolution) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitSolution) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitSolution.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitSolution proto.InternalMessageInfo

func (m *MsgCommitSolution) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCommitSolution) GetSolutionHash() string {
	if m != nil {
		return m.SolutionHash
	}
	return ""
}

func (m *MsgCommitSolution) GetSolutionScavengerHash() string {
	if m != nil {
		return m.SolutionScavengerHash
	}
	return ""
}

type MsgCommitSolutionResponse struct {
}

func (m *MsgCommitSolutionResponse) Reset()         { *m = MsgCommitSolutionResponse{} }
func (m *MsgCommitSolutionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCommitSolutionResponse) ProtoMessage()    {}
func (*MsgCommitSolutionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ec6f1544538068, []int{3}
}
func (m *MsgCommitSolutionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitSolutionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitSolutionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitSolutionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitSolutionResponse.Merge(m, src)
}
func (m *MsgCommitSolutionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitSolutionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitSolutionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitSolutionResponse proto.InternalMessageInfo

type MsgRevealSolution struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Solution string `protobuf:"bytes,2,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (m *MsgRevealSolution) Reset()         { *m = MsgRevealSolution{} }
func (m *MsgRevealSolution) String() string { return proto.CompactTextString(m) }
func (*MsgRevealSolution) ProtoMessage()    {}
func (*MsgRevealSolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ec6f1544538068, []int{4}
}
func (m *MsgRevealSolution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevealSolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevealSolution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevealSolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevealSolution.Merge(m, src)
}
func (m *MsgRevealSolution) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevealSolution) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevealSolution.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevealSolution proto.InternalMessageInfo

func (m *MsgRevealSolution) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRevealSolution) GetSolution() string {
	if m != nil {
		return m.Solution
	}
	return ""
}

type MsgRevealSolutionResponse struct {
}

func (m *MsgRevealSolutionResponse) Reset()         { *m = MsgRevealSolutionResponse{} }
func (m *MsgRevealSolutionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevealSolutionResponse) ProtoMessage()    {}
func (*MsgRevealSolutionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ec6f1544538068, []int{5}
}
func (m *MsgRevealSolutionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevealSolutionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevealSolutionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevealSolutionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevealSolutionResponse.Merge(m, src)
}
func (m *MsgRevealSolutionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevealSolutionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevealSolutionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevealSolutionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitScavenge)(nil), "coreators.scavenge.scavenge.MsgSubmitScavenge")
	proto.RegisterType((*MsgSubmitScavengeResponse)(nil), "coreators.scavenge.scavenge.MsgSubmitScavengeResponse")
	proto.RegisterType((*MsgCommitSolution)(nil), "coreators.scavenge.scavenge.MsgCommitSolution")
	proto.RegisterType((*MsgCommitSolutionResponse)(nil), "coreators.scavenge.scavenge.MsgCommitSolutionResponse")
	proto.RegisterType((*MsgRevealSolution)(nil), "coreators.scavenge.scavenge.MsgRevealSolution")
	proto.RegisterType((*MsgRevealSolutionResponse)(nil), "coreators.scavenge.scavenge.MsgRevealSolutionResponse")
}

func init() { proto.RegisterFile("scavenge/tx.proto", fileDescriptor_62ec6f1544538068) }

var fileDescriptor_62ec6f1544538068 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4e, 0x4e, 0x2c,
	0x4b, 0xcd, 0x4b, 0x4f, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e,
	0xce, 0x2f, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x83, 0x49, 0xc2, 0x19, 0x4a, 0xdd, 0x8c,
	0x5c, 0x82, 0xbe, 0xc5, 0xe9, 0xc1, 0xa5, 0x49, 0xb9, 0x99, 0x25, 0xc1, 0x50, 0x51, 0x21, 0x09,
	0x2e, 0xf6, 0x64, 0x88, 0x1e, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x89,
	0x8b, 0xa7, 0x38, 0x3f, 0xa7, 0xb4, 0x24, 0x33, 0x3f, 0xcf, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x09,
	0x2c, 0x8d, 0x22, 0x26, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x00, 0x12,
	0x92, 0x60, 0x06, 0x2b, 0x41, 0x16, 0x12, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a,
	0x91, 0x60, 0x01, 0x4b, 0x42, 0x79, 0x4a, 0xd2, 0x5c, 0x92, 0x18, 0x8e, 0x09, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x6a, 0x87, 0x38, 0xd5, 0x39, 0x3f, 0x17, 0x24, 0x0b, 0xb5, 0x90,
	0x42, 0xa7, 0x9a, 0x70, 0x89, 0xc2, 0xf8, 0x30, 0xfb, 0x8a, 0xc0, 0x8a, 0x21, 0x8e, 0xc6, 0x2e,
	0x09, 0x75, 0x26, 0xaa, 0x43, 0xe0, 0xce, 0xf4, 0x04, 0xbb, 0x32, 0x28, 0xb5, 0x2c, 0x35, 0x31,
	0x87, 0x08, 0x57, 0x4a, 0x71, 0x71, 0xc0, 0x2c, 0x81, 0xba, 0x10, 0xce, 0x87, 0xda, 0x83, 0x6a,
	0x14, 0xcc, 0x1e, 0xa3, 0xcf, 0x4c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x15, 0x5c, 0x7c, 0x68,
	0xb1, 0xa7, 0xa7, 0x87, 0x27, 0xc6, 0xf5, 0x30, 0x02, 0x58, 0xca, 0x8c, 0x34, 0xf5, 0x30, 0x17,
	0x80, 0x6c, 0x46, 0x8b, 0x0c, 0x82, 0x36, 0xa3, 0xaa, 0x27, 0x6c, 0x33, 0xf6, 0x30, 0x06, 0xd9,
	0x8c, 0x16, 0xc0, 0x04, 0x6d, 0x46, 0x55, 0x4f, 0xd8, 0x66, 0xec, 0xa1, 0xee, 0xe4, 0x71, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70,
	0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x7a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x70, 0xb3, 0xf5, 0xe1, 0xd9, 0xb1, 0x02, 0xc1, 0x2c, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0xe7, 0x4e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0xf9,
	0x9b, 0x54, 0xb2, 0x03, 0x00, 0x00,
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
	SubmitScavenge(ctx context.Context, in *MsgSubmitScavenge, opts ...grpc.CallOption) (*MsgSubmitScavengeResponse, error)
	CommitSolution(ctx context.Context, in *MsgCommitSolution, opts ...grpc.CallOption) (*MsgCommitSolutionResponse, error)
	RevealSolution(ctx context.Context, in *MsgRevealSolution, opts ...grpc.CallOption) (*MsgRevealSolutionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitScavenge(ctx context.Context, in *MsgSubmitScavenge, opts ...grpc.CallOption) (*MsgSubmitScavengeResponse, error) {
	out := new(MsgSubmitScavengeResponse)
	err := c.cc.Invoke(ctx, "/coreators.scavenge.scavenge.Msg/SubmitScavenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommitSolution(ctx context.Context, in *MsgCommitSolution, opts ...grpc.CallOption) (*MsgCommitSolutionResponse, error) {
	out := new(MsgCommitSolutionResponse)
	err := c.cc.Invoke(ctx, "/coreators.scavenge.scavenge.Msg/CommitSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevealSolution(ctx context.Context, in *MsgRevealSolution, opts ...grpc.CallOption) (*MsgRevealSolutionResponse, error) {
	out := new(MsgRevealSolutionResponse)
	err := c.cc.Invoke(ctx, "/coreators.scavenge.scavenge.Msg/RevealSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitScavenge(context.Context, *MsgSubmitScavenge) (*MsgSubmitScavengeResponse, error)
	CommitSolution(context.Context, *MsgCommitSolution) (*MsgCommitSolutionResponse, error)
	RevealSolution(context.Context, *MsgRevealSolution) (*MsgRevealSolutionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitScavenge(ctx context.Context, req *MsgSubmitScavenge) (*MsgSubmitScavengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScavenge not implemented")
}
func (*UnimplementedMsgServer) CommitSolution(ctx context.Context, req *MsgCommitSolution) (*MsgCommitSolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitSolution not implemented")
}
func (*UnimplementedMsgServer) RevealSolution(ctx context.Context, req *MsgRevealSolution) (*MsgRevealSolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevealSolution not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitScavenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitScavenge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitScavenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreators.scavenge.scavenge.Msg/SubmitScavenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitScavenge(ctx, req.(*MsgSubmitScavenge))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommitSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommitSolution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommitSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreators.scavenge.scavenge.Msg/CommitSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommitSolution(ctx, req.(*MsgCommitSolution))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevealSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevealSolution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevealSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coreators.scavenge.scavenge.Msg/RevealSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevealSolution(ctx, req.(*MsgRevealSolution))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coreators.scavenge.scavenge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitScavenge",
			Handler:    _Msg_SubmitScavenge_Handler,
		},
		{
			MethodName: "CommitSolution",
			Handler:    _Msg_CommitSolution_Handler,
		},
		{
			MethodName: "RevealSolution",
			Handler:    _Msg_RevealSolution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scavenge/tx.proto",
}

func (m *MsgSubmitScavenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitScavenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitScavenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reward) > 0 {
		i -= len(m.Reward)
		copy(dAtA[i:], m.Reward)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Reward)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SolutionHash) > 0 {
		i -= len(m.SolutionHash)
		copy(dAtA[i:], m.SolutionHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SolutionHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitScavengeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitScavengeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitScavengeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCommitSolution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitSolution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitSolution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SolutionScavengerHash) > 0 {
		i -= len(m.SolutionScavengerHash)
		copy(dAtA[i:], m.SolutionScavengerHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SolutionScavengerHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SolutionHash) > 0 {
		i -= len(m.SolutionHash)
		copy(dAtA[i:], m.SolutionHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SolutionHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCommitSolutionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitSolutionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitSolutionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevealSolution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevealSolution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevealSolution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Solution) > 0 {
		i -= len(m.Solution)
		copy(dAtA[i:], m.Solution)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Solution)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevealSolutionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevealSolutionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevealSolutionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSubmitScavenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SolutionHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Reward)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitScavengeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCommitSolution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SolutionHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SolutionScavengerHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCommitSolutionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevealSolution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Solution)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevealSolutionResponse) Size() (n int) {
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
func (m *MsgSubmitScavenge) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitScavenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitScavenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionHash", wireType)
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
			m.SolutionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
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
			m.Reward = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgSubmitScavengeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitScavengeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitScavengeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgCommitSolution) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCommitSolution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitSolution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionHash", wireType)
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
			m.SolutionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionScavengerHash", wireType)
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
			m.SolutionScavengerHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgCommitSolutionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCommitSolutionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitSolutionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgRevealSolution) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevealSolution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevealSolution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Solution", wireType)
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
			m.Solution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgRevealSolutionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevealSolutionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevealSolutionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
