// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ETHBACKENDClient is the client API for ETHBACKEND service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ETHBACKENDClient interface {
	Etherbase(ctx context.Context, in *EtherbaseRequest, opts ...grpc.CallOption) (*EtherbaseReply, error)
	NetVersion(ctx context.Context, in *NetVersionRequest, opts ...grpc.CallOption) (*NetVersionReply, error)
	// ProtocolVersion returns the Ethereum protocol version number (e.g. 66 for ETH66).
	ProtocolVersion(ctx context.Context, in *ProtocolVersionRequest, opts ...grpc.CallOption) (*ProtocolVersionReply, error)
	// ClientVersion returns the Ethereum client version string using node name convention (e.g. TurboGeth/v2021.03.2-alpha/Linux).
	ClientVersion(ctx context.Context, in *ClientVersionRequest, opts ...grpc.CallOption) (*ClientVersionReply, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (ETHBACKEND_SubscribeClient, error)
	// GetWork returns a work package for external miner.
	//
	// The work package consists of 3 strings:
	//   result[0] - 32 bytes hex encoded current block header pow-hash
	//   result[1] - 32 bytes hex encoded seed hash used for DAG
	//   result[2] - 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	//   result[3] - hex encoded block number
	GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkReply, error)
	// SubmitWork can be used by external miner to submit their POW solution.
	// It returns an indication if the work was accepted.
	// Note either an invalid solution, a stale work a non-existent work will return false.
	SubmitWork(ctx context.Context, in *SubmitWorkRequest, opts ...grpc.CallOption) (*SubmitWorkReply, error)
	// SubmitHashRate can be used for remote miners to submit their hash rate.
	// This enables the node to report the combined hash rate of all miners
	// which submit work through this node.
	//
	// It accepts the miner hash rate and an identifier which must be unique
	// between nodes.
	SubmitHashRate(ctx context.Context, in *SubmitHashRateRequest, opts ...grpc.CallOption) (*SubmitHashRateReply, error)
	// GetHashRate returns the current hashrate for local CPU miner and remote miner.
	GetHashRate(ctx context.Context, in *GetHashRateRequest, opts ...grpc.CallOption) (*GetHashRateReply, error)
	// Mining returns an indication if this node is currently mining and it's mining configuration
	Mining(ctx context.Context, in *MiningRequest, opts ...grpc.CallOption) (*MiningReply, error)
}

type eTHBACKENDClient struct {
	cc grpc.ClientConnInterface
}

func NewETHBACKENDClient(cc grpc.ClientConnInterface) ETHBACKENDClient {
	return &eTHBACKENDClient{cc}
}

func (c *eTHBACKENDClient) Etherbase(ctx context.Context, in *EtherbaseRequest, opts ...grpc.CallOption) (*EtherbaseReply, error) {
	out := new(EtherbaseReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/Etherbase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) NetVersion(ctx context.Context, in *NetVersionRequest, opts ...grpc.CallOption) (*NetVersionReply, error) {
	out := new(NetVersionReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/NetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) ProtocolVersion(ctx context.Context, in *ProtocolVersionRequest, opts ...grpc.CallOption) (*ProtocolVersionReply, error) {
	out := new(ProtocolVersionReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/ProtocolVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) ClientVersion(ctx context.Context, in *ClientVersionRequest, opts ...grpc.CallOption) (*ClientVersionReply, error) {
	out := new(ClientVersionReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/ClientVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (ETHBACKEND_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ETHBACKEND_ServiceDesc.Streams[0], "/remote.ETHBACKEND/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eTHBACKENDSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ETHBACKEND_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type eTHBACKENDSubscribeClient struct {
	grpc.ClientStream
}

func (x *eTHBACKENDSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eTHBACKENDClient) GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkReply, error) {
	out := new(GetWorkReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/GetWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) SubmitWork(ctx context.Context, in *SubmitWorkRequest, opts ...grpc.CallOption) (*SubmitWorkReply, error) {
	out := new(SubmitWorkReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/SubmitWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) SubmitHashRate(ctx context.Context, in *SubmitHashRateRequest, opts ...grpc.CallOption) (*SubmitHashRateReply, error) {
	out := new(SubmitHashRateReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/SubmitHashRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) GetHashRate(ctx context.Context, in *GetHashRateRequest, opts ...grpc.CallOption) (*GetHashRateReply, error) {
	out := new(GetHashRateReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/GetHashRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eTHBACKENDClient) Mining(ctx context.Context, in *MiningRequest, opts ...grpc.CallOption) (*MiningReply, error) {
	out := new(MiningReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/Mining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ETHBACKENDServer is the server API for ETHBACKEND service.
// All implementations must embed UnimplementedETHBACKENDServer
// for forward compatibility
type ETHBACKENDServer interface {
	Etherbase(context.Context, *EtherbaseRequest) (*EtherbaseReply, error)
	NetVersion(context.Context, *NetVersionRequest) (*NetVersionReply, error)
	// ProtocolVersion returns the Ethereum protocol version number (e.g. 66 for ETH66).
	ProtocolVersion(context.Context, *ProtocolVersionRequest) (*ProtocolVersionReply, error)
	// ClientVersion returns the Ethereum client version string using node name convention (e.g. TurboGeth/v2021.03.2-alpha/Linux).
	ClientVersion(context.Context, *ClientVersionRequest) (*ClientVersionReply, error)
	Subscribe(*SubscribeRequest, ETHBACKEND_SubscribeServer) error
	// GetWork returns a work package for external miner.
	//
	// The work package consists of 3 strings:
	//   result[0] - 32 bytes hex encoded current block header pow-hash
	//   result[1] - 32 bytes hex encoded seed hash used for DAG
	//   result[2] - 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	//   result[3] - hex encoded block number
	GetWork(context.Context, *GetWorkRequest) (*GetWorkReply, error)
	// SubmitWork can be used by external miner to submit their POW solution.
	// It returns an indication if the work was accepted.
	// Note either an invalid solution, a stale work a non-existent work will return false.
	SubmitWork(context.Context, *SubmitWorkRequest) (*SubmitWorkReply, error)
	// SubmitHashRate can be used for remote miners to submit their hash rate.
	// This enables the node to report the combined hash rate of all miners
	// which submit work through this node.
	//
	// It accepts the miner hash rate and an identifier which must be unique
	// between nodes.
	SubmitHashRate(context.Context, *SubmitHashRateRequest) (*SubmitHashRateReply, error)
	// GetHashRate returns the current hashrate for local CPU miner and remote miner.
	GetHashRate(context.Context, *GetHashRateRequest) (*GetHashRateReply, error)
	// Mining returns an indication if this node is currently mining and it's mining configuration
	Mining(context.Context, *MiningRequest) (*MiningReply, error)
	mustEmbedUnimplementedETHBACKENDServer()
}

// UnimplementedETHBACKENDServer must be embedded to have forward compatible implementations.
type UnimplementedETHBACKENDServer struct {
}

func (UnimplementedETHBACKENDServer) Etherbase(context.Context, *EtherbaseRequest) (*EtherbaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Etherbase not implemented")
}
func (UnimplementedETHBACKENDServer) NetVersion(context.Context, *NetVersionRequest) (*NetVersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetVersion not implemented")
}
func (UnimplementedETHBACKENDServer) ProtocolVersion(context.Context, *ProtocolVersionRequest) (*ProtocolVersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProtocolVersion not implemented")
}
func (UnimplementedETHBACKENDServer) ClientVersion(context.Context, *ClientVersionRequest) (*ClientVersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientVersion not implemented")
}
func (UnimplementedETHBACKENDServer) Subscribe(*SubscribeRequest, ETHBACKEND_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedETHBACKENDServer) GetWork(context.Context, *GetWorkRequest) (*GetWorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWork not implemented")
}
func (UnimplementedETHBACKENDServer) SubmitWork(context.Context, *SubmitWorkRequest) (*SubmitWorkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitWork not implemented")
}
func (UnimplementedETHBACKENDServer) SubmitHashRate(context.Context, *SubmitHashRateRequest) (*SubmitHashRateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitHashRate not implemented")
}
func (UnimplementedETHBACKENDServer) GetHashRate(context.Context, *GetHashRateRequest) (*GetHashRateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashRate not implemented")
}
func (UnimplementedETHBACKENDServer) Mining(context.Context, *MiningRequest) (*MiningReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mining not implemented")
}
func (UnimplementedETHBACKENDServer) mustEmbedUnimplementedETHBACKENDServer() {}

// UnsafeETHBACKENDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ETHBACKENDServer will
// result in compilation errors.
type UnsafeETHBACKENDServer interface {
	mustEmbedUnimplementedETHBACKENDServer()
}

func RegisterETHBACKENDServer(s grpc.ServiceRegistrar, srv ETHBACKENDServer) {
	s.RegisterService(&ETHBACKEND_ServiceDesc, srv)
}

func _ETHBACKEND_Etherbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EtherbaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).Etherbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/Etherbase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).Etherbase(ctx, req.(*EtherbaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_NetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).NetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/NetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).NetVersion(ctx, req.(*NetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_ProtocolVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).ProtocolVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/ProtocolVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).ProtocolVersion(ctx, req.(*ProtocolVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_ClientVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).ClientVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/ClientVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).ClientVersion(ctx, req.(*ClientVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ETHBACKENDServer).Subscribe(m, &eTHBACKENDSubscribeServer{stream})
}

type ETHBACKEND_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type eTHBACKENDSubscribeServer struct {
	grpc.ServerStream
}

func (x *eTHBACKENDSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

func _ETHBACKEND_GetWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).GetWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/GetWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).GetWork(ctx, req.(*GetWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_SubmitWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).SubmitWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/SubmitWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).SubmitWork(ctx, req.(*SubmitWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_SubmitHashRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitHashRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).SubmitHashRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/SubmitHashRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).SubmitHashRate(ctx, req.(*SubmitHashRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_GetHashRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHashRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).GetHashRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/GetHashRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).GetHashRate(ctx, req.(*GetHashRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ETHBACKEND_Mining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ETHBACKENDServer).Mining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.ETHBACKEND/Mining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ETHBACKENDServer).Mining(ctx, req.(*MiningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ETHBACKEND_ServiceDesc is the grpc.ServiceDesc for ETHBACKEND service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ETHBACKEND_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.ETHBACKEND",
	HandlerType: (*ETHBACKENDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Etherbase",
			Handler:    _ETHBACKEND_Etherbase_Handler,
		},
		{
			MethodName: "NetVersion",
			Handler:    _ETHBACKEND_NetVersion_Handler,
		},
		{
			MethodName: "ProtocolVersion",
			Handler:    _ETHBACKEND_ProtocolVersion_Handler,
		},
		{
			MethodName: "ClientVersion",
			Handler:    _ETHBACKEND_ClientVersion_Handler,
		},
		{
			MethodName: "GetWork",
			Handler:    _ETHBACKEND_GetWork_Handler,
		},
		{
			MethodName: "SubmitWork",
			Handler:    _ETHBACKEND_SubmitWork_Handler,
		},
		{
			MethodName: "SubmitHashRate",
			Handler:    _ETHBACKEND_SubmitHashRate_Handler,
		},
		{
			MethodName: "GetHashRate",
			Handler:    _ETHBACKEND_GetHashRate_Handler,
		},
		{
			MethodName: "Mining",
			Handler:    _ETHBACKEND_Mining_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ETHBACKEND_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote/ethbackend.proto",
}