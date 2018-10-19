// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync_trigger.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	sync_trigger.proto

It has these top-level messages:
	SyncTrigger
*/
package pb

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

// 同步消息触发
type SyncTrigger struct {
	Sequence int64 `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *SyncTrigger) Reset()                    { *m = SyncTrigger{} }
func (m *SyncTrigger) String() string            { return proto.CompactTextString(m) }
func (*SyncTrigger) ProtoMessage()               {}
func (*SyncTrigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SyncTrigger) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*SyncTrigger)(nil), "pb.SyncTrigger")
}

func init() { proto.RegisterFile("sync_trigger.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 85 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xae, 0xcc, 0x4b,
	0x8e, 0x2f, 0x29, 0xca, 0x4c, 0x4f, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x52, 0xd2, 0xe4, 0xe2, 0x0e, 0xae, 0xcc, 0x4b, 0x0e, 0x81, 0x48, 0x08, 0x49, 0x71,
	0x71, 0x14, 0xa7, 0x16, 0x96, 0xa6, 0xe6, 0x25, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07,
	0xc1, 0xf9, 0x49, 0x6c, 0x60, 0x5d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x38, 0x6a,
	0x72, 0x4b, 0x00, 0x00, 0x00,
}
