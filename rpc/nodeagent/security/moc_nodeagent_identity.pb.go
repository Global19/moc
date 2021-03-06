// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_identity.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type IdentityRequest struct {
	Identitys            []*Identity      `protobuf:"bytes,1,rep,name=Identitys,proto3" json:"Identitys,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_495ce77adac601d4, []int{0}
}

func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (m *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(m, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

func (m *IdentityRequest) GetIdentitys() []*Identity {
	if m != nil {
		return m.Identitys
	}
	return nil
}

func (m *IdentityRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type IdentityResponse struct {
	Identitys            []*Identity         `protobuf:"bytes,1,rep,name=Identitys,proto3" json:"Identitys,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IdentityResponse) Reset()         { *m = IdentityResponse{} }
func (m *IdentityResponse) String() string { return proto.CompactTextString(m) }
func (*IdentityResponse) ProtoMessage()    {}
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_495ce77adac601d4, []int{1}
}

func (m *IdentityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityResponse.Unmarshal(m, b)
}
func (m *IdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityResponse.Marshal(b, m, deterministic)
}
func (m *IdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityResponse.Merge(m, src)
}
func (m *IdentityResponse) XXX_Size() int {
	return xxx_messageInfo_IdentityResponse.Size(m)
}
func (m *IdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityResponse proto.InternalMessageInfo

func (m *IdentityResponse) GetIdentitys() []*Identity {
	if m != nil {
		return m.Identitys
	}
	return nil
}

func (m *IdentityResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *IdentityResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Identity struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ResourceGroup        string         `protobuf:"bytes,3,opt,name=resourceGroup,proto3" json:"resourceGroup,omitempty"`
	Password             string         `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Token                string         `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Certificate          []byte         `protobuf:"bytes,6,opt,name=certificate,proto3" json:"certificate,omitempty"` // Deprecated: Do not use.
	Status               *common.Status `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	NewCertificate       string         `protobuf:"bytes,10,opt,name=newCertificate,proto3" json:"newCertificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_495ce77adac601d4, []int{2}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Identity) GetResourceGroup() string {
	if m != nil {
		return m.ResourceGroup
	}
	return ""
}

func (m *Identity) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Identity) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Deprecated: Do not use.
func (m *Identity) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Identity) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Identity) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Identity) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Identity) GetNewCertificate() string {
	if m != nil {
		return m.NewCertificate
	}
	return ""
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "moc.nodeagent.security.IdentityRequest")
	proto.RegisterType((*IdentityResponse)(nil), "moc.nodeagent.security.IdentityResponse")
	proto.RegisterType((*Identity)(nil), "moc.nodeagent.security.Identity")
}

func init() { proto.RegisterFile("moc_nodeagent_identity.proto", fileDescriptor_495ce77adac601d4) }

var fileDescriptor_495ce77adac601d4 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0xb5, 0xdd, 0xdd, 0x3a, 0xcd, 0x38, 0xa3, 0x04, 0xd1, 0x50, 0x54, 0xca, 0xb8, 0x68, 0x9f,
	0x5a, 0xac, 0x3e, 0x0b, 0x8e, 0x2c, 0xb2, 0x4f, 0x42, 0x5d, 0x7c, 0x10, 0x64, 0xc9, 0xa4, 0x77,
	0x6a, 0xd8, 0x69, 0x6e, 0x4d, 0x52, 0x87, 0xf9, 0x05, 0xfe, 0x09, 0xff, 0x8a, 0xff, 0x4d, 0x26,
	0xfd, 0xd8, 0x59, 0x11, 0xf6, 0xc1, 0xa7, 0x36, 0xe7, 0x9e, 0x7b, 0xee, 0xc9, 0xc9, 0x25, 0x4f,
	0x6a, 0x14, 0x97, 0x0a, 0x4b, 0xe0, 0x15, 0x28, 0x7b, 0x29, 0x4b, 0x50, 0x56, 0xda, 0x5d, 0xda,
	0x68, 0xb4, 0x48, 0x1f, 0xd5, 0x28, 0xd2, 0xb1, 0x9a, 0x1a, 0x10, 0xad, 0x96, 0x76, 0x17, 0x3d,
	0xab, 0x10, 0xab, 0x0d, 0x64, 0x8e, 0xb5, 0x6a, 0xd7, 0xd9, 0x56, 0xf3, 0xa6, 0x01, 0x6d, 0xba,
	0xbe, 0xe8, 0xf1, 0x5e, 0x55, 0x60, 0x5d, 0xa3, 0xea, 0x3f, 0x5d, 0x61, 0xf1, 0xd3, 0x23, 0xf7,
	0xcf, 0xfb, 0x19, 0x05, 0x7c, 0x6f, 0xc1, 0x58, 0xfa, 0x96, 0x84, 0x03, 0x64, 0x98, 0x17, 0x1f,
	0x25, 0xd3, 0x3c, 0x4e, 0xff, 0x3d, 0x38, 0x1d, 0x7b, 0xaf, 0x5b, 0xe8, 0x1b, 0x32, 0xfb, 0xd8,
	0x80, 0xe6, 0x56, 0xa2, 0xba, 0xd8, 0x35, 0xc0, 0xfc, 0xd8, 0x4b, 0xe6, 0xf9, 0xdc, 0x69, 0x8c,
	0x95, 0xe2, 0x26, 0x69, 0xf1, 0xcb, 0x23, 0x0f, 0xae, 0x9d, 0x98, 0x06, 0x95, 0x81, 0xff, 0xb6,
	0x92, 0x93, 0xa0, 0x00, 0xd3, 0x6e, 0xac, 0xf3, 0x30, 0xcd, 0xa3, 0xb4, 0x0b, 0x2a, 0x1d, 0x82,
	0x4a, 0x97, 0x88, 0x9b, 0xcf, 0x7c, 0xd3, 0x42, 0xd1, 0x33, 0xe9, 0x43, 0x72, 0x72, 0xa6, 0x35,
	0x6a, 0x76, 0x14, 0x7b, 0x49, 0x58, 0x74, 0x87, 0xc5, 0x6f, 0x9f, 0x4c, 0x06, 0x5d, 0x4a, 0xc9,
	0xb1, 0xe2, 0x35, 0x30, 0xcf, 0x31, 0xdc, 0x3f, 0x9d, 0x13, 0x5f, 0x96, 0x6e, 0x4c, 0x58, 0xf8,
	0xb2, 0xa4, 0xa7, 0x64, 0xa6, 0xc1, 0x60, 0xab, 0x05, 0x7c, 0xd0, 0xd8, 0x36, 0xbd, 0xdc, 0x4d,
	0x90, 0x46, 0x64, 0xd2, 0x70, 0x63, 0xb6, 0xa8, 0x4b, 0x76, 0xec, 0x08, 0xe3, 0x79, 0x6f, 0xc4,
	0xe2, 0x15, 0x28, 0x76, 0xd2, 0x19, 0x71, 0x07, 0x7a, 0x4a, 0xa6, 0x02, 0xb4, 0x95, 0x6b, 0x29,
	0xb8, 0x05, 0x16, 0xc4, 0x5e, 0x72, 0x6f, 0xe9, 0x33, 0xaf, 0x38, 0x84, 0xe9, 0x73, 0x12, 0x18,
	0xcb, 0x6d, 0x6b, 0xd8, 0x5d, 0x77, 0xf1, 0xa9, 0x4b, 0xed, 0x93, 0x83, 0x8a, 0xbe, 0xb4, 0x27,
	0x75, 0x17, 0x62, 0x93, 0x03, 0xd2, 0x59, 0x97, 0x62, 0x5f, 0xa2, 0x4f, 0xc9, 0xb1, 0xe5, 0x95,
	0x61, 0xa1, 0xa3, 0x84, 0x8e, 0x72, 0xc1, 0x2b, 0x53, 0x38, 0x98, 0xbe, 0x20, 0x73, 0x05, 0xdb,
	0xf7, 0x07, 0x8e, 0x88, 0x73, 0xfb, 0x17, 0x9a, 0x2b, 0x32, 0x1b, 0xe2, 0x7b, 0xb7, 0x7f, 0x37,
	0xfa, 0x95, 0x04, 0xe7, 0xea, 0x07, 0x5e, 0x01, 0x7d, 0x79, 0xeb, 0x8b, 0x76, 0x8b, 0x19, 0x25,
	0xb7, 0x13, 0xbb, 0xbd, 0x59, 0xdc, 0x59, 0xbe, 0xfa, 0x92, 0x55, 0xd2, 0x7e, 0x6b, 0x57, 0xa9,
	0xc0, 0x3a, 0xab, 0xa5, 0xd0, 0x68, 0x70, 0x6d, 0xb3, 0x1a, 0x45, 0xa6, 0x1b, 0x91, 0x8d, 0x2a,
	0xd9, 0xa0, 0xb2, 0x0a, 0xdc, 0x52, 0xbc, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xfe, 0xd0,
	0x7e, 0x83, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityAgentClient is the client API for IdentityAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityAgentClient interface {
	Invoke(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error)
}

type identityAgentClient struct {
	cc *grpc.ClientConn
}

func NewIdentityAgentClient(cc *grpc.ClientConn) IdentityAgentClient {
	return &identityAgentClient{cc}
}

func (c *identityAgentClient) Invoke(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error) {
	out := new(IdentityResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.IdentityAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityAgentServer is the server API for IdentityAgent service.
type IdentityAgentServer interface {
	Invoke(context.Context, *IdentityRequest) (*IdentityResponse, error)
}

// UnimplementedIdentityAgentServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityAgentServer struct {
}

func (*UnimplementedIdentityAgentServer) Invoke(ctx context.Context, req *IdentityRequest) (*IdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterIdentityAgentServer(s *grpc.Server, srv IdentityAgentServer) {
	s.RegisterService(&_IdentityAgent_serviceDesc, srv)
}

func _IdentityAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.IdentityAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityAgentServer).Invoke(ctx, req.(*IdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.security.IdentityAgent",
	HandlerType: (*IdentityAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _IdentityAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_identity.proto",
}
