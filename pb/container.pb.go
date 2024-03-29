// Code generated by protoc-gen-go. DO NOT EDIT.
// source: container.proto

package pb

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

type Port struct {
	IP                   string   `protobuf:"bytes,1,opt,name=iP,proto3" json:"iP,omitempty"`
	PrivatePort          int32    `protobuf:"varint,2,opt,name=privatePort,proto3" json:"privatePort,omitempty"`
	PublicPort           int32    `protobuf:"varint,3,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{0}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Port) GetPrivatePort() int32 {
	if m != nil {
		return m.PrivatePort
	}
	return 0
}

func (m *Port) GetPublicPort() int32 {
	if m != nil {
		return m.PublicPort
	}
	return 0
}

func (m *Port) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type MountPoint struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Mode                 string   `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`
	RW                   bool     `protobuf:"varint,6,opt,name=rW,proto3" json:"rW,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MountPoint) Reset()         { *m = MountPoint{} }
func (m *MountPoint) String() string { return proto.CompactTextString(m) }
func (*MountPoint) ProtoMessage()    {}
func (*MountPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{1}
}

func (m *MountPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountPoint.Unmarshal(m, b)
}
func (m *MountPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountPoint.Marshal(b, m, deterministic)
}
func (m *MountPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountPoint.Merge(m, src)
}
func (m *MountPoint) XXX_Size() int {
	return xxx_messageInfo_MountPoint.Size(m)
}
func (m *MountPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_MountPoint.DiscardUnknown(m)
}

var xxx_messageInfo_MountPoint proto.InternalMessageInfo

func (m *MountPoint) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MountPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MountPoint) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *MountPoint) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *MountPoint) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *MountPoint) GetRW() bool {
	if m != nil {
		return m.RW
	}
	return false
}

type Container struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image                string        `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Command              string        `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	Created              int64         `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	State                string        `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Status               string        `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Ports                []*Port       `protobuf:"bytes,8,rep,name=ports,proto3" json:"ports,omitempty"`
	Mounts               []*MountPoint `protobuf:"bytes,9,rep,name=mounts,proto3" json:"mounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Container) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Container) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Container) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Container) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Container) GetMounts() []*MountPoint {
	if m != nil {
		return m.Mounts
	}
	return nil
}

type GetContainerRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContainerRequest) Reset()         { *m = GetContainerRequest{} }
func (m *GetContainerRequest) String() string { return proto.CompactTextString(m) }
func (*GetContainerRequest) ProtoMessage()    {}
func (*GetContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{3}
}

func (m *GetContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainerRequest.Unmarshal(m, b)
}
func (m *GetContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainerRequest.Marshal(b, m, deterministic)
}
func (m *GetContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainerRequest.Merge(m, src)
}
func (m *GetContainerRequest) XXX_Size() int {
	return xxx_messageInfo_GetContainerRequest.Size(m)
}
func (m *GetContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainerRequest proto.InternalMessageInfo

func (m *GetContainerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GetContainerRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type GetContainerResponse struct {
	Container            *Container `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetContainerResponse) Reset()         { *m = GetContainerResponse{} }
func (m *GetContainerResponse) String() string { return proto.CompactTextString(m) }
func (*GetContainerResponse) ProtoMessage()    {}
func (*GetContainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{4}
}

func (m *GetContainerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContainerResponse.Unmarshal(m, b)
}
func (m *GetContainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContainerResponse.Marshal(b, m, deterministic)
}
func (m *GetContainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContainerResponse.Merge(m, src)
}
func (m *GetContainerResponse) XXX_Size() int {
	return xxx_messageInfo_GetContainerResponse.Size(m)
}
func (m *GetContainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetContainerResponse proto.InternalMessageInfo

func (m *GetContainerResponse) GetContainer() *Container {
	if m != nil {
		return m.Container
	}
	return nil
}

type StartContainerRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartContainerRequest) Reset()         { *m = StartContainerRequest{} }
func (m *StartContainerRequest) String() string { return proto.CompactTextString(m) }
func (*StartContainerRequest) ProtoMessage()    {}
func (*StartContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{5}
}

func (m *StartContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartContainerRequest.Unmarshal(m, b)
}
func (m *StartContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartContainerRequest.Marshal(b, m, deterministic)
}
func (m *StartContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartContainerRequest.Merge(m, src)
}
func (m *StartContainerRequest) XXX_Size() int {
	return xxx_messageInfo_StartContainerRequest.Size(m)
}
func (m *StartContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartContainerRequest proto.InternalMessageInfo

func (m *StartContainerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StartContainerRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type StopContainerRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopContainerRequest) Reset()         { *m = StopContainerRequest{} }
func (m *StopContainerRequest) String() string { return proto.CompactTextString(m) }
func (*StopContainerRequest) ProtoMessage()    {}
func (*StopContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{6}
}

func (m *StopContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopContainerRequest.Unmarshal(m, b)
}
func (m *StopContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopContainerRequest.Marshal(b, m, deterministic)
}
func (m *StopContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopContainerRequest.Merge(m, src)
}
func (m *StopContainerRequest) XXX_Size() int {
	return xxx_messageInfo_StopContainerRequest.Size(m)
}
func (m *StopContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopContainerRequest proto.InternalMessageInfo

func (m *StopContainerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StopContainerRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type ErrorStatus struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorStatus) Reset()         { *m = ErrorStatus{} }
func (m *ErrorStatus) String() string { return proto.CompactTextString(m) }
func (*ErrorStatus) ProtoMessage()    {}
func (*ErrorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afe31759757e49a, []int{7}
}

func (m *ErrorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorStatus.Unmarshal(m, b)
}
func (m *ErrorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorStatus.Marshal(b, m, deterministic)
}
func (m *ErrorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorStatus.Merge(m, src)
}
func (m *ErrorStatus) XXX_Size() int {
	return xxx_messageInfo_ErrorStatus.Size(m)
}
func (m *ErrorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorStatus proto.InternalMessageInfo

func (m *ErrorStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Port)(nil), "Port")
	proto.RegisterType((*MountPoint)(nil), "MountPoint")
	proto.RegisterType((*Container)(nil), "Container")
	proto.RegisterType((*GetContainerRequest)(nil), "GetContainerRequest")
	proto.RegisterType((*GetContainerResponse)(nil), "GetContainerResponse")
	proto.RegisterType((*StartContainerRequest)(nil), "StartContainerRequest")
	proto.RegisterType((*StopContainerRequest)(nil), "StopContainerRequest")
	proto.RegisterType((*ErrorStatus)(nil), "ErrorStatus")
}

func init() { proto.RegisterFile("container.proto", fileDescriptor_7afe31759757e49a) }

var fileDescriptor_7afe31759757e49a = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x8a, 0xd4, 0x40,
	0x14, 0x25, 0xe9, 0x4e, 0xf7, 0xe4, 0x66, 0x1c, 0xa5, 0xec, 0x1e, 0x8a, 0x11, 0x24, 0xc4, 0x85,
	0x59, 0x65, 0xd1, 0x8a, 0xa0, 0xb3, 0x11, 0x1d, 0x71, 0x25, 0x34, 0xd5, 0x8b, 0x01, 0x77, 0xd5,
	0x49, 0xa1, 0x05, 0x93, 0x54, 0xac, 0xba, 0x19, 0xf0, 0x37, 0xfc, 0x1f, 0x3f, 0xc8, 0xbf, 0x90,
	0xaa, 0x3c, 0x1b, 0xb2, 0xea, 0xdd, 0x3d, 0xf7, 0x79, 0xce, 0xa9, 0x04, 0x9e, 0xe6, 0xaa, 0x42,
	0x2e, 0x2b, 0xa1, 0xb3, 0x5a, 0x2b, 0x54, 0xc9, 0x03, 0x2c, 0xf7, 0x4a, 0x23, 0xb9, 0x02, 0x5f,
	0xee, 0xa9, 0x17, 0x7b, 0x69, 0xc8, 0x7c, 0xb9, 0x27, 0x31, 0x44, 0xb5, 0x96, 0x8f, 0x1c, 0x85,
	0x2d, 0x53, 0x3f, 0xf6, 0xd2, 0x80, 0x4d, 0x53, 0xe4, 0x25, 0x40, 0xdd, 0x1c, 0x1f, 0x64, 0xee,
	0x1a, 0x16, 0xae, 0x61, 0x92, 0x21, 0x04, 0x96, 0xf8, 0xbb, 0x16, 0x74, 0xe9, 0x76, 0xba, 0x38,
	0xf9, 0xe3, 0x01, 0x7c, 0x53, 0x4d, 0x85, 0x7b, 0x25, 0xab, 0xb1, 0xc5, 0x1b, 0x5b, 0x6c, 0xae,
	0xe2, 0xa5, 0x70, 0x17, 0x43, 0xe6, 0x62, 0x72, 0x0d, 0x2b, 0xa3, 0x1a, 0x9d, 0x0b, 0x77, 0x26,
	0x64, 0x1d, 0xb2, 0x24, 0x0b, 0x61, 0x50, 0x56, 0x1c, 0xa5, 0xaa, 0xba, 0x4b, 0xd3, 0x94, 0xdd,
	0x56, 0xaa, 0x42, 0xd0, 0xa0, 0xdd, 0x66, 0x63, 0x2b, 0x55, 0xdf, 0xd3, 0x55, 0xec, 0xa5, 0x17,
	0xcc, 0xd7, 0xf7, 0xc9, 0x3f, 0x0f, 0xc2, 0xcf, 0xbd, 0x2d, 0xce, 0x88, 0x62, 0x30, 0xa2, 0x98,
	0xe5, 0xb3, 0x81, 0x40, 0x96, 0xfc, 0x47, 0x4f, 0xa7, 0x05, 0x84, 0xc2, 0x3a, 0x57, 0x65, 0xc9,
	0xab, 0xa2, 0x63, 0xd2, 0x43, 0x57, 0xd1, 0x82, 0xa3, 0x28, 0x1c, 0x91, 0x05, 0xeb, 0xa1, 0xdd,
	0x64, 0x90, 0xa3, 0x70, 0x74, 0x42, 0xd6, 0x02, 0xa7, 0x17, 0x39, 0x36, 0x86, 0xae, 0x3b, 0xbd,
	0x0e, 0x91, 0x17, 0x10, 0xd4, 0x4a, 0xa3, 0xa1, 0x17, 0xf1, 0x22, 0x8d, 0x76, 0x41, 0x66, 0x8d,
	0x66, 0x6d, 0x8e, 0xbc, 0x82, 0x55, 0x69, 0xad, 0x35, 0x34, 0x74, 0xd5, 0x28, 0x1b, 0x9d, 0x66,
	0x5d, 0x29, 0x79, 0x0f, 0xcf, 0xbf, 0x0a, 0x1c, 0xd4, 0x32, 0xf1, 0xab, 0x11, 0xa6, 0x7d, 0xfd,
	0xbb, 0x41, 0xf4, 0x9d, 0x15, 0xfd, 0x53, 0x19, 0xec, 0x45, 0xdb, 0x38, 0xf9, 0x08, 0x9b, 0xd3,
	0x51, 0x53, 0xab, 0xca, 0x08, 0x92, 0x42, 0x38, 0x7c, 0x54, 0x6e, 0x45, 0xb4, 0x83, 0x6c, 0x6c,
	0x1b, 0x8b, 0xc9, 0x2d, 0x6c, 0x0f, 0xc8, 0xf5, 0x79, 0xe7, 0x3f, 0xc0, 0xe6, 0x80, 0xaa, 0x3e,
	0x6b, 0xf6, 0x35, 0x44, 0x5f, 0xb4, 0x56, 0xfa, 0xd0, 0xda, 0x48, 0x61, 0x5d, 0x0a, 0x63, 0xec,
	0x03, 0xb6, 0x73, 0x3d, 0xdc, 0xfd, 0xf5, 0xe0, 0xd9, 0x70, 0xe1, 0x20, 0xf4, 0xa3, 0xcc, 0x05,
	0xb9, 0x85, 0xcb, 0xa9, 0x70, 0xb2, 0xc9, 0x66, 0x2c, 0xbc, 0xd9, 0x66, 0xb3, 0xee, 0xbc, 0x83,
	0xab, 0x53, 0xcd, 0xe4, 0x3a, 0x9b, 0x35, 0xe1, 0xe6, 0x32, 0x9b, 0x72, 0x7c, 0x0b, 0x4f, 0x4e,
	0xe4, 0x92, 0x6d, 0x36, 0x27, 0xff, 0x74, 0xea, 0xd3, 0xf2, 0xbb, 0x5f, 0x1f, 0x8f, 0x2b, 0xf7,
	0x6b, 0xbf, 0xf9, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x0c, 0x09, 0x02, 0xed, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContainerServiceClient is the client API for ContainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContainerServiceClient interface {
	GetContainer(ctx context.Context, in *GetContainerRequest, opts ...grpc.CallOption) (*GetContainerResponse, error)
	StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*ErrorStatus, error)
	StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*ErrorStatus, error)
}

type containerServiceClient struct {
	cc *grpc.ClientConn
}

func NewContainerServiceClient(cc *grpc.ClientConn) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) GetContainer(ctx context.Context, in *GetContainerRequest, opts ...grpc.CallOption) (*GetContainerResponse, error) {
	out := new(GetContainerResponse)
	err := c.cc.Invoke(ctx, "/ContainerService/GetContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*ErrorStatus, error) {
	out := new(ErrorStatus)
	err := c.cc.Invoke(ctx, "/ContainerService/StartContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*ErrorStatus, error) {
	out := new(ErrorStatus)
	err := c.cc.Invoke(ctx, "/ContainerService/StopContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServiceServer is the server API for ContainerService service.
type ContainerServiceServer interface {
	GetContainer(context.Context, *GetContainerRequest) (*GetContainerResponse, error)
	StartContainer(context.Context, *StartContainerRequest) (*ErrorStatus, error)
	StopContainer(context.Context, *StopContainerRequest) (*ErrorStatus, error)
}

// UnimplementedContainerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContainerServiceServer struct {
}

func (*UnimplementedContainerServiceServer) GetContainer(ctx context.Context, req *GetContainerRequest) (*GetContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainer not implemented")
}
func (*UnimplementedContainerServiceServer) StartContainer(ctx context.Context, req *StartContainerRequest) (*ErrorStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartContainer not implemented")
}
func (*UnimplementedContainerServiceServer) StopContainer(ctx context.Context, req *StopContainerRequest) (*ErrorStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopContainer not implemented")
}

func RegisterContainerServiceServer(s *grpc.Server, srv ContainerServiceServer) {
	s.RegisterService(&_ContainerService_serviceDesc, srv)
}

func _ContainerService_GetContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).GetContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContainerService/GetContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).GetContainer(ctx, req.(*GetContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_StartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).StartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContainerService/StartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).StartContainer(ctx, req.(*StartContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_StopContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).StopContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContainerService/StopContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).StopContainer(ctx, req.(*StopContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContainerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContainer",
			Handler:    _ContainerService_GetContainer_Handler,
		},
		{
			MethodName: "StartContainer",
			Handler:    _ContainerService_StartContainer_Handler,
		},
		{
			MethodName: "StopContainer",
			Handler:    _ContainerService_StopContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "container.proto",
}
