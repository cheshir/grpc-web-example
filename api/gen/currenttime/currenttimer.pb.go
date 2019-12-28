// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currenttimer.proto

package currenttime

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

type CurrentTimeResponse struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CurrentTimeResponse) Reset()         { *m = CurrentTimeResponse{} }
func (m *CurrentTimeResponse) String() string { return proto.CompactTextString(m) }
func (*CurrentTimeResponse) ProtoMessage()    {}
func (*CurrentTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69f15d075c93b3f3, []int{0}
}

func (m *CurrentTimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentTimeResponse.Unmarshal(m, b)
}
func (m *CurrentTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentTimeResponse.Marshal(b, m, deterministic)
}
func (m *CurrentTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentTimeResponse.Merge(m, src)
}
func (m *CurrentTimeResponse) XXX_Size() int {
	return xxx_messageInfo_CurrentTimeResponse.Size(m)
}
func (m *CurrentTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentTimeResponse proto.InternalMessageInfo

func (m *CurrentTimeResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*CurrentTimeResponse)(nil), "currenttime.CurrentTimeResponse")
}

func init() { proto.RegisterFile("currenttimer.proto", fileDescriptor_69f15d075c93b3f3) }

var fileDescriptor_69f15d075c93b3f3 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2e, 0x2d, 0x2a,
	0x4a, 0xcd, 0x2b, 0x29, 0xc9, 0xcc, 0x4d, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x46, 0x12, 0x93, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x25, 0x95, 0xa6,
	0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x54, 0x4a, 0xc9, 0xa3, 0x4b, 0x82, 0xb4, 0x14, 0x97,
	0x24, 0xe6, 0x16, 0x40, 0x14, 0x28, 0xf9, 0x73, 0x09, 0x3b, 0x43, 0x0c, 0x0b, 0xc9, 0xcc, 0x4d,
	0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe0, 0xe2, 0x84, 0xab, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0x98, 0xa5, 0x07, 0x33, 0x4b, 0x2f, 0x04, 0xa6,
	0x22, 0x08, 0xa1, 0xd8, 0x28, 0x86, 0x8b, 0x07, 0xc9, 0xc0, 0x22, 0x21, 0x1f, 0x2e, 0x3e, 0xf7,
	0xd4, 0x12, 0x24, 0x21, 0x21, 0x31, 0x0c, 0x83, 0x5c, 0x41, 0x2e, 0x96, 0x52, 0xd0, 0x43, 0xf2,
	0x96, 0x1e, 0x16, 0x57, 0x29, 0x31, 0x24, 0xb1, 0x81, 0xf5, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x22, 0x4a, 0x44, 0x51, 0x16, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CurrentTimerClient is the client API for CurrentTimer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrentTimerClient interface {
	GetCurrentTime(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CurrentTimeResponse, error)
}

type currentTimerClient struct {
	cc *grpc.ClientConn
}

func NewCurrentTimerClient(cc *grpc.ClientConn) CurrentTimerClient {
	return &currentTimerClient{cc}
}

func (c *currentTimerClient) GetCurrentTime(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CurrentTimeResponse, error) {
	out := new(CurrentTimeResponse)
	err := c.cc.Invoke(ctx, "/currenttime.CurrentTimer/GetCurrentTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentTimerServer is the server API for CurrentTimer service.
type CurrentTimerServer interface {
	GetCurrentTime(context.Context, *empty.Empty) (*CurrentTimeResponse, error)
}

func RegisterCurrentTimerServer(s *grpc.Server, srv CurrentTimerServer) {
	s.RegisterService(&_CurrentTimer_serviceDesc, srv)
}

func _CurrentTimer_GetCurrentTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTimerServer).GetCurrentTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currenttime.CurrentTimer/GetCurrentTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTimerServer).GetCurrentTime(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurrentTimer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "currenttime.CurrentTimer",
	HandlerType: (*CurrentTimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentTime",
			Handler:    _CurrentTimer_GetCurrentTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currenttimer.proto",
}
