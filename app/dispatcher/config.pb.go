package dispatcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionConfig struct {
	AllowPassiveConnection bool `protobuf:"varint,1,opt,name=allow_passive_connection,json=allowPassiveConnection" json:"allow_passive_connection,omitempty"`
}

func (m *SessionConfig) Reset()                    { *m = SessionConfig{} }
func (m *SessionConfig) String() string            { return proto.CompactTextString(m) }
func (*SessionConfig) ProtoMessage()               {}
func (*SessionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionConfig) GetAllowPassiveConnection() bool {
	if m != nil {
		return m.AllowPassiveConnection
	}
	return false
}

type Config struct {
	Settings *SessionConfig `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetSettings() *SessionConfig {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionConfig)(nil), "v2ray.core.app.dispatcher.SessionConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.dispatcher.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/app/dispatcher/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x4f,
	0xc9, 0x2c, 0x2e, 0x48, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x84, 0xa9, 0x2d, 0x4a, 0xd5, 0x4b, 0x2c, 0x28,
	0xd0, 0x43, 0xa8, 0x53, 0xf2, 0xe4, 0xe2, 0x0d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x73, 0x06,
	0xeb, 0x10, 0xb2, 0xe0, 0x92, 0x48, 0xcc, 0xc9, 0xc9, 0x2f, 0x8f, 0x2f, 0x48, 0x2c, 0x2e, 0xce,
	0x2c, 0x4b, 0x8d, 0x4f, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x08, 0x12, 0x03, 0xcb, 0x07, 0x40, 0xa4, 0x9d, 0xe1, 0xb2, 0x4a, 0x7e, 0x5c,
	0x6c, 0x50, 0x33, 0x5c, 0xb8, 0x38, 0x8a, 0x53, 0x4b, 0x4a, 0x32, 0xf3, 0xd2, 0x8b, 0xc1, 0x7a,
	0xb8, 0x8d, 0x34, 0xf4, 0x70, 0x3a, 0x41, 0x0f, 0xc5, 0xfe, 0x20, 0xb8, 0x4e, 0xa7, 0x10, 0x2e,
	0xd9, 0xe4, 0xfc, 0x5c, 0xdc, 0x1a, 0x9d, 0xb8, 0x21, 0x5a, 0x02, 0x40, 0x7e, 0x8c, 0xe2, 0x42,
	0x48, 0xac, 0x62, 0x92, 0x0c, 0x33, 0x0a, 0x4a, 0xac, 0xd4, 0x73, 0x06, 0x69, 0x72, 0x2c, 0x28,
	0xd0, 0x73, 0x81, 0xcb, 0x25, 0xb1, 0x81, 0x83, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4d,
	0xf7, 0x84, 0x36, 0x40, 0x01, 0x00, 0x00,
}
