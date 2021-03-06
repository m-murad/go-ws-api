// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	device.proto

It has these top-level messages:
	ID
	UpdateDevice
	Device
	Devices
	Empty
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Device_DeviceType int32

const (
	Device_unknown Device_DeviceType = 0
	Device_onOff   Device_DeviceType = 1
	Device_dimmer  Device_DeviceType = 2
	Device_sensor  Device_DeviceType = 3
)

var Device_DeviceType_name = map[int32]string{
	0: "unknown",
	1: "onOff",
	2: "dimmer",
	3: "sensor",
}
var Device_DeviceType_value = map[string]int32{
	"unknown": 0,
	"onOff":   1,
	"dimmer":  2,
	"sensor":  3,
}

func (x Device_DeviceType) String() string {
	return proto.EnumName(Device_DeviceType_name, int32(x))
}
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ID struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateDevice struct {
	Id    int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Value int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *UpdateDevice) Reset()                    { *m = UpdateDevice{} }
func (m *UpdateDevice) String() string            { return proto.CompactTextString(m) }
func (*UpdateDevice) ProtoMessage()               {}
func (*UpdateDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateDevice) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateDevice) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Device struct {
	Id       int32             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Hardware string            `protobuf:"bytes,2,opt,name=hardware" json:"hardware,omitempty"`
	Name     string            `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Location string            `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Type     Device_DeviceType `protobuf:"varint,5,opt,name=type,enum=pb.Device_DeviceType" json:"type,omitempty"`
	Unit     string            `protobuf:"bytes,6,opt,name=unit" json:"unit,omitempty"`
	State    int32             `protobuf:"varint,7,opt,name=state" json:"state,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Device) GetHardware() string {
	if m != nil {
		return m.Hardware
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Device) GetType() Device_DeviceType {
	if m != nil {
		return m.Type
	}
	return Device_unknown
}

func (m *Device) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Device) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

type Devices struct {
	Device []*Device `protobuf:"bytes,1,rep,name=device" json:"device,omitempty"`
}

func (m *Devices) Reset()                    { *m = Devices{} }
func (m *Devices) String() string            { return proto.CompactTextString(m) }
func (*Devices) ProtoMessage()               {}
func (*Devices) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Devices) GetDevice() []*Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*ID)(nil), "pb.ID")
	proto.RegisterType((*UpdateDevice)(nil), "pb.UpdateDevice")
	proto.RegisterType((*Device)(nil), "pb.Device")
	proto.RegisterType((*Devices)(nil), "pb.Devices")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterEnum("pb.Device_DeviceType", Device_DeviceType_name, Device_DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceService service

type DeviceServiceClient interface {
	// List all registered devices
	GetAllDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Devices, error)
	// Get a device by ID
	GetDeviceByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Device, error)
	// Update a device's state
	SwitchDevice(ctx context.Context, in *UpdateDevice, opts ...grpc.CallOption) (*Device, error)
	// Register a new device
	RegisterDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
}

type deviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceServiceClient(cc *grpc.ClientConn) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) GetAllDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Devices, error) {
	out := new(Devices)
	err := grpc.Invoke(ctx, "/pb.DeviceService/GetAllDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/pb.DeviceService/GetDeviceByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) SwitchDevice(ctx context.Context, in *UpdateDevice, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/pb.DeviceService/SwitchDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) RegisterDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/pb.DeviceService/RegisterDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceService service

type DeviceServiceServer interface {
	// List all registered devices
	GetAllDevices(context.Context, *Empty) (*Devices, error)
	// Get a device by ID
	GetDeviceByID(context.Context, *ID) (*Device, error)
	// Update a device's state
	SwitchDevice(context.Context, *UpdateDevice) (*Device, error)
	// Register a new device
	RegisterDevice(context.Context, *Device) (*Device, error)
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_GetAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/GetAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetAllDevices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/GetDeviceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_SwitchDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDevice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).SwitchDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/SwitchDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).SwitchDevice(ctx, req.(*UpdateDevice))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllDevices",
			Handler:    _DeviceService_GetAllDevices_Handler,
		},
		{
			MethodName: "GetDeviceByID",
			Handler:    _DeviceService_GetDeviceByID_Handler,
		},
		{
			MethodName: "SwitchDevice",
			Handler:    _DeviceService_SwitchDevice_Handler,
		},
		{
			MethodName: "RegisterDevice",
			Handler:    _DeviceService_RegisterDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}

func init() { proto.RegisterFile("device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x72, 0xd3, 0x30,
	0x14, 0x85, 0xb1, 0xf3, 0x47, 0x6f, 0x7f, 0x08, 0x22, 0xb4, 0x1e, 0x4f, 0x17, 0xc6, 0x1b, 0x4a,
	0x07, 0xc7, 0x34, 0xb0, 0xca, 0xb0, 0x68, 0x4a, 0x3a, 0x9d, 0x2c, 0x18, 0x98, 0x94, 0xce, 0xb0,
	0x55, 0xec, 0x1b, 0x47, 0xc5, 0x96, 0x34, 0xb6, 0x92, 0xd4, 0xd3, 0xe9, 0x86, 0x47, 0x80, 0x1d,
	0x0b, 0xde, 0x81, 0x67, 0xe1, 0x15, 0x78, 0x07, 0xb6, 0x8c, 0xe5, 0x94, 0xa6, 0x2d, 0x5d, 0xb0,
	0x92, 0xae, 0x74, 0xce, 0xa7, 0x23, 0xeb, 0x1a, 0xd6, 0x42, 0x9c, 0xb1, 0x00, 0xdb, 0x32, 0x15,
	0x4a, 0x10, 0x53, 0x8e, 0xec, 0xed, 0x48, 0x88, 0x28, 0x46, 0x9f, 0x4a, 0xe6, 0x53, 0xce, 0x85,
	0xa2, 0x8a, 0x09, 0x9e, 0x95, 0x0a, 0xfb, 0xb9, 0x1e, 0x02, 0x2f, 0x42, 0xee, 0x65, 0x73, 0x1a,
	0x45, 0x98, 0xfa, 0x42, 0x6a, 0xc5, 0x6d, 0xb5, 0xdb, 0x02, 0x73, 0xd0, 0x27, 0x1b, 0x60, 0xb2,
	0xd0, 0x32, 0x1c, 0x63, 0xa7, 0x36, 0x34, 0x59, 0xe8, 0xbe, 0x82, 0xb5, 0x13, 0x19, 0x52, 0x85,
	0x7d, 0x7d, 0xf6, 0xcd, 0x7d, 0xd2, 0x82, 0xda, 0x8c, 0xc6, 0x53, 0xb4, 0x4c, 0xbd, 0x54, 0x16,
	0xee, 0x6f, 0x03, 0xea, 0x77, 0x18, 0x6c, 0xb8, 0x3f, 0xa1, 0x69, 0x38, 0xa7, 0x69, 0xe9, 0x59,
	0x19, 0xfe, 0xad, 0x09, 0x81, 0x2a, 0xa7, 0x09, 0x5a, 0x15, 0xbd, 0xae, 0xe7, 0x85, 0x3e, 0x16,
	0x81, 0x4e, 0x6a, 0x55, 0x4b, 0xfd, 0x65, 0x4d, 0x9e, 0x41, 0x55, 0xe5, 0x12, 0xad, 0x9a, 0x63,
	0xec, 0x6c, 0x74, 0x1e, 0xb7, 0xe5, 0xa8, 0x5d, 0x9e, 0xba, 0x18, 0x3e, 0xe4, 0x12, 0x87, 0x5a,
	0x52, 0xa0, 0xa7, 0x9c, 0x29, 0xab, 0x5e, 0xa2, 0x8b, 0x79, 0x91, 0x3d, 0x53, 0x54, 0xa1, 0xd5,
	0x28, 0xb3, 0xeb, 0xc2, 0x7d, 0x0d, 0x70, 0xe5, 0x26, 0xab, 0xd0, 0x98, 0xf2, 0x4f, 0x5c, 0xcc,
	0x79, 0xf3, 0x1e, 0x59, 0x81, 0x9a, 0xe0, 0xef, 0xc6, 0xe3, 0xa6, 0x41, 0x00, 0xea, 0x21, 0x4b,
	0x12, 0x4c, 0x9b, 0x66, 0x31, 0xcf, 0x90, 0x67, 0x22, 0x6d, 0x56, 0x5c, 0x0f, 0x1a, 0xa5, 0x3b,
	0x23, 0x2e, 0xd4, 0xcb, 0x07, 0xb3, 0x0c, 0xa7, 0xb2, 0xb3, 0xda, 0x81, 0xab, 0x7c, 0xc3, 0xc5,
	0x8e, 0xdb, 0x80, 0xda, 0x61, 0x22, 0x55, 0xde, 0xf9, 0x61, 0xc2, 0x7a, 0xb9, 0x77, 0x8c, 0xa9,
	0xfe, 0x70, 0xfb, 0xb0, 0x7e, 0x84, 0xaa, 0x17, 0xc7, 0x97, 0xbc, 0x95, 0xc2, 0xaf, 0xd5, 0xf6,
	0xea, 0x15, 0x2a, 0x73, 0xb7, 0x3e, 0xff, 0xfc, 0xf5, 0xd5, 0x7c, 0x48, 0x1e, 0xe8, 0x36, 0x98,
	0xed, 0xf9, 0xe1, 0xc2, 0xd0, 0xd3, 0x84, 0x52, 0x76, 0x90, 0x0f, 0xfa, 0xa4, 0x5e, 0xd8, 0x06,
	0x7d, 0x7b, 0x29, 0x89, 0xbb, 0xad, 0xdd, 0x9b, 0xa4, 0x75, 0xc3, 0xed, 0x9f, 0xb3, 0xf0, 0x82,
	0x9c, 0xc0, 0xda, 0xf1, 0x9c, 0xa9, 0x60, 0xb2, 0x78, 0xcd, 0x66, 0xe1, 0x5c, 0x6e, 0x88, 0x6b,
	0xac, 0xa7, 0x9a, 0xf5, 0xc4, 0xdd, 0xfe, 0x17, 0xcb, 0x3f, 0xd7, 0xbd, 0x71, 0xd1, 0x35, 0x76,
	0x49, 0x1f, 0x36, 0x86, 0x18, 0xb1, 0x4c, 0x61, 0xba, 0x00, 0x2f, 0x61, 0xae, 0x21, 0x6d, 0x8d,
	0x6c, 0xb9, 0x37, 0x2f, 0xd7, 0x35, 0x76, 0x0f, 0xbe, 0x9b, 0x5f, 0x7a, 0xdf, 0x4c, 0xa2, 0xe0,
	0x51, 0xcf, 0x39, 0x60, 0xca, 0x11, 0x63, 0xe7, 0x70, 0x86, 0x69, 0xae, 0x26, 0x8c, 0x47, 0xee,
	0x47, 0x68, 0x45, 0xc3, 0xf7, 0x6f, 0xbc, 0x23, 0xaa, 0x70, 0x4e, 0x73, 0x47, 0xa6, 0xe2, 0x14,
	0x03, 0x45, 0xda, 0x13, 0xa5, 0x64, 0xd6, 0xf5, 0xfd, 0x88, 0xa9, 0xc9, 0x74, 0xd4, 0x0e, 0x44,
	0xe2, 0x47, 0xa9, 0x0c, 0x3c, 0x0c, 0x44, 0x96, 0x67, 0x0a, 0x17, 0x65, 0x54, 0xfa, 0xec, 0x26,
	0x17, 0x1c, 0xf7, 0xf1, 0x8c, 0x26, 0x32, 0xc6, 0x42, 0xdd, 0xa9, 0xec, 0xb5, 0x5f, 0xec, 0x56,
	0x0c, 0xb3, 0xda, 0x69, 0x52, 0x29, 0x63, 0x56, 0x76, 0xa3, 0x7f, 0x9a, 0x09, 0xde, 0xd9, 0x5c,
	0x5e, 0x39, 0xf3, 0xc6, 0x42, 0x78, 0x09, 0x4b, 0xb0, 0x7b, 0x4b, 0xd9, 0xbd, 0x43, 0x99, 0x0e,
	0x60, 0xeb, 0xad, 0x48, 0xd1, 0xa1, 0x23, 0x31, 0x55, 0xce, 0xf2, 0x15, 0xfe, 0x37, 0xfa, 0xa8,
	0xae, 0xff, 0xec, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x94, 0x66, 0xbf, 0x39, 0x04,
	0x00, 0x00,
}
