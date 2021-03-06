// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package client

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// FetchAccountStateRequest is the request to fetch an account's balance and nonce.
type FetchAccountStateRequest struct {
	// The account address
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAccountStateRequest) Reset()         { *m = FetchAccountStateRequest{} }
func (m *FetchAccountStateRequest) String() string { return proto.CompactTextString(m) }
func (*FetchAccountStateRequest) ProtoMessage()    {}
func (*FetchAccountStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *FetchAccountStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAccountStateRequest.Unmarshal(m, b)
}
func (m *FetchAccountStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAccountStateRequest.Marshal(b, m, deterministic)
}
func (m *FetchAccountStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAccountStateRequest.Merge(m, src)
}
func (m *FetchAccountStateRequest) XXX_Size() int {
	return xxx_messageInfo_FetchAccountStateRequest.Size(m)
}
func (m *FetchAccountStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAccountStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAccountStateRequest proto.InternalMessageInfo

func (m *FetchAccountStateRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// FetchAccountStateResponse is the response of FetchAccountStateRequest.
type FetchAccountStateResponse struct {
	// The balance of the account (big.Int)
	Balance []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	// The nonce of the account
	Nonce                uint64   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAccountStateResponse) Reset()         { *m = FetchAccountStateResponse{} }
func (m *FetchAccountStateResponse) String() string { return proto.CompactTextString(m) }
func (*FetchAccountStateResponse) ProtoMessage()    {}
func (*FetchAccountStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func (m *FetchAccountStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAccountStateResponse.Unmarshal(m, b)
}
func (m *FetchAccountStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAccountStateResponse.Marshal(b, m, deterministic)
}
func (m *FetchAccountStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAccountStateResponse.Merge(m, src)
}
func (m *FetchAccountStateResponse) XXX_Size() int {
	return xxx_messageInfo_FetchAccountStateResponse.Size(m)
}
func (m *FetchAccountStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAccountStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAccountStateResponse proto.InternalMessageInfo

func (m *FetchAccountStateResponse) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *FetchAccountStateResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

// GetFreeTokenRequest is the request to get free token from the faucet smart contract.
type GetFreeTokenRequest struct {
	// The account address to receive the free token
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFreeTokenRequest) Reset()         { *m = GetFreeTokenRequest{} }
func (m *GetFreeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetFreeTokenRequest) ProtoMessage()    {}
func (*GetFreeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{2}
}

func (m *GetFreeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFreeTokenRequest.Unmarshal(m, b)
}
func (m *GetFreeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFreeTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetFreeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFreeTokenRequest.Merge(m, src)
}
func (m *GetFreeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetFreeTokenRequest.Size(m)
}
func (m *GetFreeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFreeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFreeTokenRequest proto.InternalMessageInfo

func (m *GetFreeTokenRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// GetFreeTokenResponse is the response of GetFreeTokenRequest.
type GetFreeTokenResponse struct {
	// The transaction Id that requests free token from the faucet.
	TxId                 []byte   `protobuf:"bytes,1,opt,name=txId,proto3" json:"txId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFreeTokenResponse) Reset()         { *m = GetFreeTokenResponse{} }
func (m *GetFreeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetFreeTokenResponse) ProtoMessage()    {}
func (*GetFreeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3}
}

func (m *GetFreeTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFreeTokenResponse.Unmarshal(m, b)
}
func (m *GetFreeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFreeTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetFreeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFreeTokenResponse.Merge(m, src)
}
func (m *GetFreeTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetFreeTokenResponse.Size(m)
}
func (m *GetFreeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFreeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFreeTokenResponse proto.InternalMessageInfo

func (m *GetFreeTokenResponse) GetTxId() []byte {
	if m != nil {
		return m.TxId
	}
	return nil
}

// StakingContractInfoRequest is the request to necessary info for stkaing.
type StakingContractInfoRequest struct {
	// The account address
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakingContractInfoRequest) Reset()         { *m = StakingContractInfoRequest{} }
func (m *StakingContractInfoRequest) String() string { return proto.CompactTextString(m) }
func (*StakingContractInfoRequest) ProtoMessage()    {}
func (*StakingContractInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{4}
}

func (m *StakingContractInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakingContractInfoRequest.Unmarshal(m, b)
}
func (m *StakingContractInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakingContractInfoRequest.Marshal(b, m, deterministic)
}
func (m *StakingContractInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingContractInfoRequest.Merge(m, src)
}
func (m *StakingContractInfoRequest) XXX_Size() int {
	return xxx_messageInfo_StakingContractInfoRequest.Size(m)
}
func (m *StakingContractInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingContractInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakingContractInfoRequest proto.InternalMessageInfo

func (m *StakingContractInfoRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// StakingContractInfoResponse is the response of GetStakingContractInfo.
type StakingContractInfoResponse struct {
	// Contract address.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// The balance of the staking account.
	Balance []byte `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	// The nonce of the staking account.
	Nonce                uint64   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakingContractInfoResponse) Reset()         { *m = StakingContractInfoResponse{} }
func (m *StakingContractInfoResponse) String() string { return proto.CompactTextString(m) }
func (*StakingContractInfoResponse) ProtoMessage()    {}
func (*StakingContractInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{5}
}

func (m *StakingContractInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakingContractInfoResponse.Unmarshal(m, b)
}
func (m *StakingContractInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakingContractInfoResponse.Marshal(b, m, deterministic)
}
func (m *StakingContractInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingContractInfoResponse.Merge(m, src)
}
func (m *StakingContractInfoResponse) XXX_Size() int {
	return xxx_messageInfo_StakingContractInfoResponse.Size(m)
}
func (m *StakingContractInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingContractInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakingContractInfoResponse proto.InternalMessageInfo

func (m *StakingContractInfoResponse) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *StakingContractInfoResponse) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *StakingContractInfoResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchAccountStateRequest)(nil), "client.FetchAccountStateRequest")
	proto.RegisterType((*FetchAccountStateResponse)(nil), "client.FetchAccountStateResponse")
	proto.RegisterType((*GetFreeTokenRequest)(nil), "client.GetFreeTokenRequest")
	proto.RegisterType((*GetFreeTokenResponse)(nil), "client.GetFreeTokenResponse")
	proto.RegisterType((*StakingContractInfoRequest)(nil), "client.StakingContractInfoRequest")
	proto.RegisterType((*StakingContractInfoResponse)(nil), "client.StakingContractInfoResponse")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x14, 0xfc, 0xda, 0x0f, 0x31, 0xbe, 0xd4, 0xa8, 0x2b, 0x31, 0xb5, 0x78, 0xa8, 0xeb, 0x05, 0x3d,
	0x60, 0xa2, 0xc6, 0x3b, 0x21, 0x81, 0x10, 0x6e, 0xad, 0x27, 0x2f, 0x66, 0xd9, 0x3e, 0xb5, 0x81,
	0xec, 0x62, 0xfb, 0x30, 0xfc, 0x19, 0xff, 0xab, 0xa1, 0xdd, 0x35, 0x35, 0xb6, 0x70, 0xdb, 0x99,
	0x7d, 0xf3, 0x26, 0x3b, 0xb3, 0xe0, 0xc9, 0x45, 0x8a, 0x8a, 0xfa, 0xcb, 0x4c, 0x93, 0x66, 0xed,
	0x12, 0xf1, 0x07, 0xf0, 0x47, 0x48, 0xf2, 0x7d, 0x20, 0xa5, 0x5e, 0x29, 0x8a, 0x49, 0x10, 0x46,
	0xf8, 0xb1, 0xc2, 0x9c, 0x98, 0x0f, 0xfb, 0x22, 0x49, 0x32, 0xcc, 0x73, 0xdf, 0x09, 0x9d, 0x9e,
	0x17, 0x59, 0xc8, 0xa7, 0x70, 0x5e, 0xa3, 0xca, 0x97, 0x5a, 0xe5, 0xb8, 0x91, 0xcd, 0xc4, 0x42,
	0x28, 0x89, 0x56, 0x66, 0x20, 0xeb, 0xc0, 0x9e, 0xd2, 0x1b, 0xde, 0x0d, 0x9d, 0x5e, 0x2b, 0x2a,
	0x01, 0xbf, 0x85, 0xd3, 0x31, 0xd2, 0x28, 0x43, 0x7c, 0xd2, 0x73, 0x54, 0xbb, 0xdd, 0x6f, 0xa0,
	0xf3, 0x5b, 0x60, 0x8c, 0x19, 0xb4, 0x68, 0x3d, 0x49, 0xcc, 0x78, 0x71, 0xe6, 0x8f, 0x10, 0xc4,
	0x24, 0xe6, 0xa9, 0x7a, 0x1b, 0x6a, 0x45, 0x99, 0x90, 0x34, 0x51, 0xaf, 0x7a, 0xb7, 0xc7, 0x1a,
	0xba, 0xb5, 0x3a, 0x63, 0x75, 0x0d, 0xc7, 0xd2, 0xf0, 0x2f, 0xd5, 0x0d, 0x07, 0xd1, 0x91, 0xe5,
	0x07, 0x25, 0x5d, 0x8d, 0xc3, 0x6d, 0x88, 0xe3, 0x7f, 0x25, 0x8e, 0xbb, 0x2f, 0x17, 0x0e, 0x87,
	0x45, 0x39, 0x31, 0x66, 0x9f, 0xa9, 0x44, 0xf6, 0x0c, 0x27, 0x7f, 0xd2, 0x66, 0x61, 0xdf, 0xf4,
	0xd9, 0x54, 0x5f, 0x70, 0xb9, 0x65, 0xa2, 0x7c, 0x06, 0xff, 0xc7, 0xa6, 0xe0, 0x55, 0xb3, 0x64,
	0x5d, 0x2b, 0xaa, 0xa9, 0x24, 0xb8, 0xa8, 0xbf, 0xfc, 0x59, 0x26, 0xe1, 0x6c, 0x8c, 0x54, 0x93,
	0x1b, 0xe3, 0x56, 0xd9, 0x5c, 0x46, 0x70, 0xb5, 0x75, 0xc6, 0x9a, 0xcc, 0xda, 0xc5, 0x07, 0xbe,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x3e, 0x0b, 0x3e, 0xd0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	FetchAccountState(ctx context.Context, in *FetchAccountStateRequest, opts ...grpc.CallOption) (*FetchAccountStateResponse, error)
	GetFreeToken(ctx context.Context, in *GetFreeTokenRequest, opts ...grpc.CallOption) (*GetFreeTokenResponse, error)
	GetStakingContractInfo(ctx context.Context, in *StakingContractInfoRequest, opts ...grpc.CallOption) (*StakingContractInfoResponse, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) FetchAccountState(ctx context.Context, in *FetchAccountStateRequest, opts ...grpc.CallOption) (*FetchAccountStateResponse, error) {
	out := new(FetchAccountStateResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/FetchAccountState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetFreeToken(ctx context.Context, in *GetFreeTokenRequest, opts ...grpc.CallOption) (*GetFreeTokenResponse, error) {
	out := new(GetFreeTokenResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/GetFreeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetStakingContractInfo(ctx context.Context, in *StakingContractInfoRequest, opts ...grpc.CallOption) (*StakingContractInfoResponse, error) {
	out := new(StakingContractInfoResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/GetStakingContractInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	FetchAccountState(context.Context, *FetchAccountStateRequest) (*FetchAccountStateResponse, error)
	GetFreeToken(context.Context, *GetFreeTokenRequest) (*GetFreeTokenResponse, error)
	GetStakingContractInfo(context.Context, *StakingContractInfoRequest) (*StakingContractInfoResponse, error)
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_FetchAccountState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAccountStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).FetchAccountState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/FetchAccountState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).FetchAccountState(ctx, req.(*FetchAccountStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetFreeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetFreeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/GetFreeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetFreeToken(ctx, req.(*GetFreeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetStakingContractInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingContractInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetStakingContractInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/GetStakingContractInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetStakingContractInfo(ctx, req.(*StakingContractInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAccountState",
			Handler:    _ClientService_FetchAccountState_Handler,
		},
		{
			MethodName: "GetFreeToken",
			Handler:    _ClientService_GetFreeToken_Handler,
		},
		{
			MethodName: "GetStakingContractInfo",
			Handler:    _ClientService_GetStakingContractInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}
