// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart.proto

package cart

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CartRegisterReq struct {
	CartName             string   `protobuf:"bytes,1,opt,name=CartName,proto3" json:"CartName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartRegisterReq) Reset()         { *m = CartRegisterReq{} }
func (m *CartRegisterReq) String() string { return proto.CompactTextString(m) }
func (*CartRegisterReq) ProtoMessage()    {}
func (*CartRegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{1}
}

func (m *CartRegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartRegisterReq.Unmarshal(m, b)
}
func (m *CartRegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartRegisterReq.Marshal(b, m, deterministic)
}
func (m *CartRegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartRegisterReq.Merge(m, src)
}
func (m *CartRegisterReq) XXX_Size() int {
	return xxx_messageInfo_CartRegisterReq.Size(m)
}
func (m *CartRegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CartRegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_CartRegisterReq proto.InternalMessageInfo

func (m *CartRegisterReq) GetCartName() string {
	if m != nil {
		return m.CartName
	}
	return ""
}

type CartRegisterResp struct {
	CartId               int64    `protobuf:"varint,1,opt,name=CartId,proto3" json:"CartId,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartRegisterResp) Reset()         { *m = CartRegisterResp{} }
func (m *CartRegisterResp) String() string { return proto.CompactTextString(m) }
func (*CartRegisterResp) ProtoMessage()    {}
func (*CartRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{2}
}

func (m *CartRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartRegisterResp.Unmarshal(m, b)
}
func (m *CartRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartRegisterResp.Marshal(b, m, deterministic)
}
func (m *CartRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartRegisterResp.Merge(m, src)
}
func (m *CartRegisterResp) XXX_Size() int {
	return xxx_messageInfo_CartRegisterResp.Size(m)
}
func (m *CartRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CartRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_CartRegisterResp proto.InternalMessageInfo

func (m *CartRegisterResp) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *CartRegisterResp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type CartUpdateReq struct {
	CartId               int64    `protobuf:"varint,1,opt,name=CartId,proto3" json:"CartId,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartUpdateReq) Reset()         { *m = CartUpdateReq{} }
func (m *CartUpdateReq) String() string { return proto.CompactTextString(m) }
func (*CartUpdateReq) ProtoMessage()    {}
func (*CartUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{3}
}

func (m *CartUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartUpdateReq.Unmarshal(m, b)
}
func (m *CartUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartUpdateReq.Marshal(b, m, deterministic)
}
func (m *CartUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartUpdateReq.Merge(m, src)
}
func (m *CartUpdateReq) XXX_Size() int {
	return xxx_messageInfo_CartUpdateReq.Size(m)
}
func (m *CartUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CartUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CartUpdateReq proto.InternalMessageInfo

func (m *CartUpdateReq) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *CartUpdateReq) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type CartUpdateResp struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartUpdateResp) Reset()         { *m = CartUpdateResp{} }
func (m *CartUpdateResp) String() string { return proto.CompactTextString(m) }
func (*CartUpdateResp) ProtoMessage()    {}
func (*CartUpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{4}
}

func (m *CartUpdateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartUpdateResp.Unmarshal(m, b)
}
func (m *CartUpdateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartUpdateResp.Marshal(b, m, deterministic)
}
func (m *CartUpdateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartUpdateResp.Merge(m, src)
}
func (m *CartUpdateResp) XXX_Size() int {
	return xxx_messageInfo_CartUpdateResp.Size(m)
}
func (m *CartUpdateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CartUpdateResp.DiscardUnknown(m)
}

var xxx_messageInfo_CartUpdateResp proto.InternalMessageInfo

func (m *CartUpdateResp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type AlarmReq struct {
	CartId               int64    `protobuf:"varint,1,opt,name=CartId,proto3" json:"CartId,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmReq) Reset()         { *m = AlarmReq{} }
func (m *AlarmReq) String() string { return proto.CompactTextString(m) }
func (*AlarmReq) ProtoMessage()    {}
func (*AlarmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{5}
}

func (m *AlarmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmReq.Unmarshal(m, b)
}
func (m *AlarmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmReq.Marshal(b, m, deterministic)
}
func (m *AlarmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmReq.Merge(m, src)
}
func (m *AlarmReq) XXX_Size() int {
	return xxx_messageInfo_AlarmReq.Size(m)
}
func (m *AlarmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmReq.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmReq proto.InternalMessageInfo

func (m *AlarmReq) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *AlarmReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AlarmResp struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmResp) Reset()         { *m = AlarmResp{} }
func (m *AlarmResp) String() string { return proto.CompactTextString(m) }
func (*AlarmResp) ProtoMessage()    {}
func (*AlarmResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{6}
}

func (m *AlarmResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmResp.Unmarshal(m, b)
}
func (m *AlarmResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmResp.Marshal(b, m, deterministic)
}
func (m *AlarmResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmResp.Merge(m, src)
}
func (m *AlarmResp) XXX_Size() int {
	return xxx_messageInfo_AlarmResp.Size(m)
}
func (m *AlarmResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmResp.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmResp proto.InternalMessageInfo

func (m *AlarmResp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type CartInfo struct {
	CartId               int64    `protobuf:"varint,1,opt,name=CartId,proto3" json:"CartId,omitempty"`
	CartName             string   `protobuf:"bytes,2,opt,name=CartName,proto3" json:"CartName,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartInfo) Reset()         { *m = CartInfo{} }
func (m *CartInfo) String() string { return proto.CompactTextString(m) }
func (*CartInfo) ProtoMessage()    {}
func (*CartInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{7}
}

func (m *CartInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartInfo.Unmarshal(m, b)
}
func (m *CartInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartInfo.Marshal(b, m, deterministic)
}
func (m *CartInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartInfo.Merge(m, src)
}
func (m *CartInfo) XXX_Size() int {
	return xxx_messageInfo_CartInfo.Size(m)
}
func (m *CartInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CartInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CartInfo proto.InternalMessageInfo

func (m *CartInfo) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *CartInfo) GetCartName() string {
	if m != nil {
		return m.CartName
	}
	return ""
}

func (m *CartInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type GetInfoResp struct {
	CartList             []*CartInfo `protobuf:"bytes,1,rep,name=CartList,proto3" json:"CartList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetInfoResp) Reset()         { *m = GetInfoResp{} }
func (m *GetInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetInfoResp) ProtoMessage()    {}
func (*GetInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{8}
}

func (m *GetInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResp.Unmarshal(m, b)
}
func (m *GetInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResp.Marshal(b, m, deterministic)
}
func (m *GetInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResp.Merge(m, src)
}
func (m *GetInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetInfoResp.Size(m)
}
func (m *GetInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResp proto.InternalMessageInfo

func (m *GetInfoResp) GetCartList() []*CartInfo {
	if m != nil {
		return m.CartList
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "cart.Empty")
	proto.RegisterType((*CartRegisterReq)(nil), "cart.CartRegisterReq")
	proto.RegisterType((*CartRegisterResp)(nil), "cart.CartRegisterResp")
	proto.RegisterType((*CartUpdateReq)(nil), "cart.CartUpdateReq")
	proto.RegisterType((*CartUpdateResp)(nil), "cart.CartUpdateResp")
	proto.RegisterType((*AlarmReq)(nil), "cart.AlarmReq")
	proto.RegisterType((*AlarmResp)(nil), "cart.AlarmResp")
	proto.RegisterType((*CartInfo)(nil), "cart.CartInfo")
	proto.RegisterType((*GetInfoResp)(nil), "cart.GetInfoResp")
}

func init() { proto.RegisterFile("cart.proto", fileDescriptor_bf731a5c8f9a516f) }

var fileDescriptor_bf731a5c8f9a516f = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x4b, 0x3b, 0x31,
	0x10, 0x25, 0xdb, 0xff, 0xd3, 0xdf, 0xaf, 0xad, 0x63, 0x2d, 0x65, 0xbd, 0x2c, 0x39, 0xad, 0x82,
	0x3d, 0x54, 0x41, 0x14, 0x7a, 0x10, 0x11, 0x29, 0xa8, 0x85, 0xa8, 0x1f, 0x20, 0xda, 0x58, 0x4a,
	0xad, 0x1b, 0x37, 0xb9, 0xf8, 0x39, 0xfd, 0x42, 0x92, 0x6c, 0xba, 0x9b, 0xc5, 0xd6, 0xdb, 0xbe,
	0x99, 0x79, 0xef, 0xed, 0xbc, 0x09, 0xc0, 0x2b, 0x4f, 0xf5, 0x48, 0xa6, 0x89, 0x4e, 0xb0, 0x6a,
	0xbe, 0x69, 0x03, 0x6a, 0x37, 0x6b, 0xa9, 0xbf, 0xe8, 0x09, 0x74, 0xaf, 0x79, 0xaa, 0x99, 0x58,
	0x2c, 0x95, 0x16, 0x29, 0x13, 0x9f, 0x18, 0x42, 0xd3, 0x94, 0x1e, 0xf8, 0x5a, 0x0c, 0x49, 0x44,
	0xe2, 0x16, 0xcb, 0x31, 0xbd, 0x84, 0x5e, 0x79, 0x5c, 0x49, 0x1c, 0x40, 0xdd, 0xd4, 0xa6, 0x73,
	0x3b, 0x5d, 0x61, 0x0e, 0x61, 0x07, 0x82, 0xd9, 0x6a, 0x18, 0x44, 0x24, 0x6e, 0xb2, 0x60, 0xb6,
	0xa2, 0x13, 0xf8, 0x6f, 0x3a, 0xcf, 0x72, 0xce, 0xb5, 0x30, 0x46, 0xbb, 0x88, 0x7d, 0xa8, 0x3d,
	0x6a, 0xae, 0x85, 0xe5, 0xb6, 0x58, 0x06, 0x68, 0x04, 0x1d, 0x9f, 0xae, 0xa4, 0x33, 0x20, 0xb9,
	0xc1, 0x19, 0x34, 0xaf, 0xde, 0x79, 0xba, 0xfe, 0x4b, 0xbb, 0x07, 0x95, 0x7b, 0xb5, 0x70, 0xca,
	0xe6, 0x93, 0x1e, 0x42, 0xcb, 0xb1, 0xb6, 0x48, 0x3e, 0x65, 0x59, 0x4c, 0x3f, 0xde, 0x92, 0x9d,
	0x92, 0x7e, 0x5e, 0x41, 0x39, 0xaf, 0x62, 0x95, 0xaa, 0xbf, 0xca, 0x05, 0xb4, 0x6f, 0x85, 0x15,
	0xb5, 0xa6, 0xc7, 0x99, 0xc0, 0xdd, 0x52, 0xe9, 0x21, 0x89, 0x2a, 0x71, 0x7b, 0xdc, 0x19, 0xd9,
	0x8b, 0x6d, 0xac, 0x59, 0xde, 0x1f, 0x7f, 0x13, 0xa8, 0x1a, 0x80, 0x13, 0xf8, 0xe7, 0x5f, 0x02,
	0x0f, 0x0a, 0x8a, 0x77, 0xcc, 0x70, 0xb0, 0xad, 0xac, 0x24, 0x9e, 0x03, 0x14, 0x69, 0xe2, 0x7e,
	0x31, 0x95, 0x9f, 0x27, 0xec, 0xff, 0x2e, 0x2a, 0x89, 0x31, 0xd4, 0x6c, 0x5c, 0xe8, 0xfe, 0x71,
	0x93, 0x78, 0xd8, 0x2d, 0x61, 0x25, 0xf1, 0x08, 0x1a, 0x6e, 0x4b, 0x6c, 0x67, 0x3d, 0xfb, 0xe4,
	0xc2, 0xbd, 0x0c, 0x78, 0x09, 0xbc, 0xd4, 0xed, 0xdb, 0x3c, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0x61, 0x3e, 0xd5, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartClient interface {
	CartRegister(ctx context.Context, in *CartRegisterReq, opts ...grpc.CallOption) (*CartRegisterResp, error)
	CartUpdate(ctx context.Context, in *CartUpdateReq, opts ...grpc.CallOption) (*CartUpdateResp, error)
	Alarm(ctx context.Context, in *AlarmReq, opts ...grpc.CallOption) (*AlarmResp, error)
	GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetInfoResp, error)
}

type cartClient struct {
	cc *grpc.ClientConn
}

func NewCartClient(cc *grpc.ClientConn) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) CartRegister(ctx context.Context, in *CartRegisterReq, opts ...grpc.CallOption) (*CartRegisterResp, error) {
	out := new(CartRegisterResp)
	err := c.cc.Invoke(ctx, "/cart.Cart/CartRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) CartUpdate(ctx context.Context, in *CartUpdateReq, opts ...grpc.CallOption) (*CartUpdateResp, error) {
	out := new(CartUpdateResp)
	err := c.cc.Invoke(ctx, "/cart.Cart/CartUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) Alarm(ctx context.Context, in *AlarmReq, opts ...grpc.CallOption) (*AlarmResp, error) {
	out := new(AlarmResp)
	err := c.cc.Invoke(ctx, "/cart.Cart/Alarm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetInfoResp, error) {
	out := new(GetInfoResp)
	err := c.cc.Invoke(ctx, "/cart.Cart/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
type CartServer interface {
	CartRegister(context.Context, *CartRegisterReq) (*CartRegisterResp, error)
	CartUpdate(context.Context, *CartUpdateReq) (*CartUpdateResp, error)
	Alarm(context.Context, *AlarmReq) (*AlarmResp, error)
	GetInfo(context.Context, *Empty) (*GetInfoResp, error)
}

// UnimplementedCartServer can be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (*UnimplementedCartServer) CartRegister(ctx context.Context, req *CartRegisterReq) (*CartRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartRegister not implemented")
}
func (*UnimplementedCartServer) CartUpdate(ctx context.Context, req *CartUpdateReq) (*CartUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartUpdate not implemented")
}
func (*UnimplementedCartServer) Alarm(ctx context.Context, req *AlarmReq) (*AlarmResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alarm not implemented")
}
func (*UnimplementedCartServer) GetInfo(ctx context.Context, req *Empty) (*GetInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func RegisterCartServer(s *grpc.Server, srv CartServer) {
	s.RegisterService(&_Cart_serviceDesc, srv)
}

func _Cart_CartRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).CartRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.Cart/CartRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).CartRegister(ctx, req.(*CartRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_CartUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).CartUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.Cart/CartUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).CartUpdate(ctx, req.(*CartUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_Alarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).Alarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.Cart/Alarm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).Alarm(ctx, req.(*AlarmReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.Cart/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cart_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CartRegister",
			Handler:    _Cart_CartRegister_Handler,
		},
		{
			MethodName: "CartUpdate",
			Handler:    _Cart_CartUpdate_Handler,
		},
		{
			MethodName: "Alarm",
			Handler:    _Cart_Alarm_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Cart_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
