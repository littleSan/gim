// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: logic/device.int.proto

package logicpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId   uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`      // 设备id
	UserId     uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`            // 用户id
	Token      string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`                             // 秘钥
	ConnAddr   string `protobuf:"bytes,4,opt,name=conn_addr,json=connAddr,proto3" json:"conn_addr,omitempty"`       // 服务器地址
	ClientAddr string `protobuf:"bytes,5,opt,name=client_addr,json=clientAddr,proto3" json:"client_addr,omitempty"` // 客户端地址
}

func (x *ConnSignInRequest) Reset() {
	*x = ConnSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_device_int_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnSignInRequest) ProtoMessage() {}

func (x *ConnSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logic_device_int_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnSignInRequest.ProtoReflect.Descriptor instead.
func (*ConnSignInRequest) Descriptor() ([]byte, []int) {
	return file_logic_device_int_proto_rawDescGZIP(), []int{0}
}

func (x *ConnSignInRequest) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *ConnSignInRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ConnSignInRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnSignInRequest) GetConnAddr() string {
	if x != nil {
		return x.ConnAddr
	}
	return ""
}

func (x *ConnSignInRequest) GetClientAddr() string {
	if x != nil {
		return x.ClientAddr
	}
	return ""
}

type OfflineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`            // 用户id
	DeviceId   uint64 `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`      // 设备id
	ClientAddr string `protobuf:"bytes,3,opt,name=client_addr,json=clientAddr,proto3" json:"client_addr,omitempty"` // 客户端地址
}

func (x *OfflineRequest) Reset() {
	*x = OfflineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_device_int_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineRequest) ProtoMessage() {}

func (x *OfflineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logic_device_int_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineRequest.ProtoReflect.Descriptor instead.
func (*OfflineRequest) Descriptor() ([]byte, []int) {
	return file_logic_device_int_proto_rawDescGZIP(), []int{1}
}

func (x *OfflineRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OfflineRequest) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *OfflineRequest) GetClientAddr() string {
	if x != nil {
		return x.ClientAddr
	}
	return ""
}

type GetDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetDeviceRequest) Reset() {
	*x = GetDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_device_int_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceRequest) ProtoMessage() {}

func (x *GetDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logic_device_int_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return file_logic_device_int_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeviceRequest) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

type GetDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *GetDeviceReply) Reset() {
	*x = GetDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_device_int_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceReply) ProtoMessage() {}

func (x *GetDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_logic_device_int_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceReply.ProtoReflect.Descriptor instead.
func (*GetDeviceReply) Descriptor() ([]byte, []int) {
	return file_logic_device_int_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceReply) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId      uint64     `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`               // 设备id
	UserId        uint64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                     // 用户id
	Type          DeviceType `protobuf:"varint,3,opt,name=type,proto3,enum=logic.DeviceType" json:"type,omitempty"`                 // 设备类型
	Brand         string     `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`                                      // 手机厂商
	Model         string     `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`                                      // 机型
	SystemVersion string     `protobuf:"bytes,6,opt,name=system_version,json=systemVersion,proto3" json:"system_version,omitempty"` // 系统版本
	SdkVersion    string     `protobuf:"bytes,7,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`          // SDK版本
	Status        int32      `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`                                   // 在线状态，0：不在线；1：在线
	ConnAddr      string     `protobuf:"bytes,9,opt,name=conn_addr,json=connAddr,proto3" json:"conn_addr,omitempty"`                // 服务端连接地址
	ClientAddr    string     `protobuf:"bytes,10,opt,name=client_addr,json=clientAddr,proto3" json:"client_addr,omitempty"`         // 客户端地址
	CreateTime    int64      `protobuf:"varint,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`        // 创建时间
	UpdateTime    int64      `protobuf:"varint,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`        // 更新时间
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_device_int_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_logic_device_int_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_logic_device_int_proto_rawDescGZIP(), []int{4}
}

func (x *Device) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *Device) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Device) GetType() DeviceType {
	if x != nil {
		return x.Type
	}
	return DeviceType_DT_DEFAULT
}

func (x *Device) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Device) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Device) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *Device) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *Device) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Device) GetConnAddr() string {
	if x != nil {
		return x.ConnAddr
	}
	return ""
}

func (x *Device) GetClientAddr() string {
	if x != nil {
		return x.ClientAddr
	}
	return ""
}

func (x *Device) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Device) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type ServerStopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnAddr string `protobuf:"bytes,1,opt,name=conn_addr,json=connAddr,proto3" json:"conn_addr,omitempty"`
}

func (x *ServerStopRequest) Reset() {
	*x = ServerStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_device_int_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStopRequest) ProtoMessage() {}

func (x *ServerStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logic_device_int_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStopRequest.ProtoReflect.Descriptor instead.
func (*ServerStopRequest) Descriptor() ([]byte, []int) {
	return file_logic_device_int_proto_rawDescGZIP(), []int{5}
}

func (x *ServerStopRequest) GetConnAddr() string {
	if x != nil {
		return x.ConnAddr
	}
	return ""
}

var File_logic_device_int_proto protoreflect.FileDescriptor

var file_logic_device_int_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x22, 0x67, 0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x22, 0x2f, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x37,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x25, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xf1, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x11, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x32, 0x89, 0x02,
	0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x38, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x15, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x0a, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x6d,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logic_device_int_proto_rawDescOnce sync.Once
	file_logic_device_int_proto_rawDescData = file_logic_device_int_proto_rawDesc
)

func file_logic_device_int_proto_rawDescGZIP() []byte {
	file_logic_device_int_proto_rawDescOnce.Do(func() {
		file_logic_device_int_proto_rawDescData = protoimpl.X.CompressGZIP(file_logic_device_int_proto_rawDescData)
	})
	return file_logic_device_int_proto_rawDescData
}

var file_logic_device_int_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_logic_device_int_proto_goTypes = []interface{}{
	(*ConnSignInRequest)(nil), // 0: logic.ConnSignInRequest
	(*OfflineRequest)(nil),    // 1: logic.OfflineRequest
	(*GetDeviceRequest)(nil),  // 2: logic.GetDeviceRequest
	(*GetDeviceReply)(nil),    // 3: logic.GetDeviceReply
	(*Device)(nil),            // 4: logic.Device
	(*ServerStopRequest)(nil), // 5: logic.ServerStopRequest
	(DeviceType)(0),           // 6: logic.DeviceType
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_logic_device_int_proto_depIdxs = []int32{
	4, // 0: logic.GetDeviceReply.device:type_name -> logic.Device
	6, // 1: logic.Device.type:type_name -> logic.DeviceType
	0, // 2: logic.DeviceIntService.ConnSignIn:input_type -> logic.ConnSignInRequest
	1, // 3: logic.DeviceIntService.Offline:input_type -> logic.OfflineRequest
	2, // 4: logic.DeviceIntService.GetDevice:input_type -> logic.GetDeviceRequest
	5, // 5: logic.DeviceIntService.ServerStop:input_type -> logic.ServerStopRequest
	7, // 6: logic.DeviceIntService.ConnSignIn:output_type -> google.protobuf.Empty
	7, // 7: logic.DeviceIntService.Offline:output_type -> google.protobuf.Empty
	3, // 8: logic.DeviceIntService.GetDevice:output_type -> logic.GetDeviceReply
	7, // 9: logic.DeviceIntService.ServerStop:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logic_device_int_proto_init() }
func file_logic_device_int_proto_init() {
	if File_logic_device_int_proto != nil {
		return
	}
	file_logic_device_ext_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_logic_device_int_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnSignInRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logic_device_int_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logic_device_int_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logic_device_int_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logic_device_int_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logic_device_int_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStopRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logic_device_int_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logic_device_int_proto_goTypes,
		DependencyIndexes: file_logic_device_int_proto_depIdxs,
		MessageInfos:      file_logic_device_int_proto_msgTypes,
	}.Build()
	File_logic_device_int_proto = out.File
	file_logic_device_int_proto_rawDesc = nil
	file_logic_device_int_proto_goTypes = nil
	file_logic_device_int_proto_depIdxs = nil
}
