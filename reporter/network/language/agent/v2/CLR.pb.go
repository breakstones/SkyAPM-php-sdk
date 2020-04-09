// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/CLR.proto

package v2

import (
	fmt "fmt"
	v2 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v2"
	proto "github.com/golang/protobuf/proto"
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

type CLRMetric struct {
	Time                 int64      `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu                  *v2.CPU    `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gc                   *ClrGC     `protobuf:"bytes,3,opt,name=gc,proto3" json:"gc,omitempty"`
	Thread               *ClrThread `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CLRMetric) Reset()         { *m = CLRMetric{} }
func (m *CLRMetric) String() string { return proto.CompactTextString(m) }
func (*CLRMetric) ProtoMessage()    {}
func (*CLRMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{0}
}

func (m *CLRMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLRMetric.Unmarshal(m, b)
}
func (m *CLRMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLRMetric.Marshal(b, m, deterministic)
}
func (m *CLRMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLRMetric.Merge(m, src)
}
func (m *CLRMetric) XXX_Size() int {
	return xxx_messageInfo_CLRMetric.Size(m)
}
func (m *CLRMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_CLRMetric.DiscardUnknown(m)
}

var xxx_messageInfo_CLRMetric proto.InternalMessageInfo

func (m *CLRMetric) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *CLRMetric) GetCpu() *v2.CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CLRMetric) GetGc() *ClrGC {
	if m != nil {
		return m.Gc
	}
	return nil
}

func (m *CLRMetric) GetThread() *ClrThread {
	if m != nil {
		return m.Thread
	}
	return nil
}

type ClrGC struct {
	Gen0CollectCount     int64    `protobuf:"varint,1,opt,name=Gen0CollectCount,proto3" json:"Gen0CollectCount,omitempty"`
	Gen1CollectCount     int64    `protobuf:"varint,2,opt,name=Gen1CollectCount,proto3" json:"Gen1CollectCount,omitempty"`
	Gen2CollectCount     int64    `protobuf:"varint,3,opt,name=Gen2CollectCount,proto3" json:"Gen2CollectCount,omitempty"`
	HeapMemory           int64    `protobuf:"varint,4,opt,name=HeapMemory,proto3" json:"HeapMemory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClrGC) Reset()         { *m = ClrGC{} }
func (m *ClrGC) String() string { return proto.CompactTextString(m) }
func (*ClrGC) ProtoMessage()    {}
func (*ClrGC) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{1}
}

func (m *ClrGC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrGC.Unmarshal(m, b)
}
func (m *ClrGC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrGC.Marshal(b, m, deterministic)
}
func (m *ClrGC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrGC.Merge(m, src)
}
func (m *ClrGC) XXX_Size() int {
	return xxx_messageInfo_ClrGC.Size(m)
}
func (m *ClrGC) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrGC.DiscardUnknown(m)
}

var xxx_messageInfo_ClrGC proto.InternalMessageInfo

func (m *ClrGC) GetGen0CollectCount() int64 {
	if m != nil {
		return m.Gen0CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen1CollectCount() int64 {
	if m != nil {
		return m.Gen1CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen2CollectCount() int64 {
	if m != nil {
		return m.Gen2CollectCount
	}
	return 0
}

func (m *ClrGC) GetHeapMemory() int64 {
	if m != nil {
		return m.HeapMemory
	}
	return 0
}

type ClrThread struct {
	AvailableCompletionPortThreads int32    `protobuf:"varint,1,opt,name=AvailableCompletionPortThreads,proto3" json:"AvailableCompletionPortThreads,omitempty"`
	AvailableWorkerThreads         int32    `protobuf:"varint,2,opt,name=AvailableWorkerThreads,proto3" json:"AvailableWorkerThreads,omitempty"`
	MaxCompletionPortThreads       int32    `protobuf:"varint,3,opt,name=MaxCompletionPortThreads,proto3" json:"MaxCompletionPortThreads,omitempty"`
	MaxWorkerThreads               int32    `protobuf:"varint,4,opt,name=MaxWorkerThreads,proto3" json:"MaxWorkerThreads,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *ClrThread) Reset()         { *m = ClrThread{} }
func (m *ClrThread) String() string { return proto.CompactTextString(m) }
func (*ClrThread) ProtoMessage()    {}
func (*ClrThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{2}
}

func (m *ClrThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrThread.Unmarshal(m, b)
}
func (m *ClrThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrThread.Marshal(b, m, deterministic)
}
func (m *ClrThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrThread.Merge(m, src)
}
func (m *ClrThread) XXX_Size() int {
	return xxx_messageInfo_ClrThread.Size(m)
}
func (m *ClrThread) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrThread.DiscardUnknown(m)
}

var xxx_messageInfo_ClrThread proto.InternalMessageInfo

func (m *ClrThread) GetAvailableCompletionPortThreads() int32 {
	if m != nil {
		return m.AvailableCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetAvailableWorkerThreads() int32 {
	if m != nil {
		return m.AvailableWorkerThreads
	}
	return 0
}

func (m *ClrThread) GetMaxCompletionPortThreads() int32 {
	if m != nil {
		return m.MaxCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetMaxWorkerThreads() int32 {
	if m != nil {
		return m.MaxWorkerThreads
	}
	return 0
}

func init() {
	proto.RegisterType((*CLRMetric)(nil), "skywalking.network.protocol.common.CLRMetric")
	proto.RegisterType((*ClrGC)(nil), "skywalking.network.protocol.common.ClrGC")
	proto.RegisterType((*ClrThread)(nil), "skywalking.network.protocol.common.ClrThread")
}

func init() { proto.RegisterFile("common/CLR.proto", fileDescriptor_a10d56830892247a) }

var fileDescriptor_a10d56830892247a = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0x87, 0xb1, 0x9d, 0x04, 0xa6, 0xdd, 0x04, 0x0d, 0x86, 0xc9, 0x45, 0x08, 0x66, 0xb0, 0x6c,
	0x30, 0x3b, 0xc9, 0x60, 0x90, 0xdd, 0x6d, 0xa6, 0x4d, 0x2f, 0xe2, 0x62, 0xd4, 0x96, 0x40, 0xef,
	0x14, 0x55, 0x38, 0xc6, 0xb2, 0x64, 0x14, 0xe5, 0xdf, 0x2b, 0xb5, 0xaf, 0xd4, 0xb7, 0xe8, 0x0b,
	0x94, 0xc8, 0x8e, 0x1b, 0x93, 0x84, 0xe4, 0xca, 0xe6, 0xfc, 0xbe, 0x4f, 0xc7, 0xc7, 0x47, 0xa0,
	0x49, 0x44, 0x9a, 0x0a, 0xee, 0xf9, 0x63, 0xe4, 0x66, 0x52, 0x28, 0x01, 0x9d, 0x79, 0xb2, 0x59,
	0x61, 0x96, 0xc4, 0x3c, 0x72, 0x39, 0x55, 0x2b, 0x21, 0x93, 0x3c, 0x21, 0x82, 0xb9, 0x39, 0xdd,
	0xfa, 0x52, 0x58, 0xf9, 0x23, 0x8f, 0x9d, 0x57, 0x03, 0x7c, 0xf2, 0xc7, 0x28, 0xa0, 0x4a, 0xc6,
	0x04, 0x42, 0x50, 0x53, 0x71, 0x4a, 0x6d, 0xa3, 0x63, 0x74, 0x2d, 0xa4, 0xdf, 0xe1, 0x10, 0x58,
	0x24, 0x5b, 0xd8, 0x66, 0xc7, 0xe8, 0x7e, 0x1e, 0x7c, 0x77, 0xcf, 0x37, 0x72, 0xfd, 0xf0, 0x01,
	0x6d, 0x1d, 0x38, 0x04, 0x66, 0x44, 0x6c, 0x4b, 0x9b, 0x3f, 0x2e, 0x32, 0x99, 0x1c, 0xf9, 0xc8,
	0x8c, 0x08, 0xbc, 0x02, 0x0d, 0x35, 0x93, 0x14, 0x3f, 0xd9, 0x35, 0xad, 0xff, 0xba, 0x50, 0xbf,
	0xd7, 0x12, 0x2a, 0x64, 0xe7, 0xc5, 0x00, 0x75, 0x7d, 0x28, 0xfc, 0x09, 0x9a, 0x23, 0xca, 0x7b,
	0xbe, 0x60, 0x8c, 0x12, 0xe5, 0x8b, 0x05, 0x57, 0xc5, 0x98, 0x07, 0xf5, 0x82, 0xed, 0x57, 0x58,
	0xb3, 0x64, 0xfb, 0x47, 0xd8, 0x41, 0x85, 0xb5, 0x4a, 0xb6, 0x52, 0x87, 0x6d, 0x00, 0x6e, 0x28,
	0xce, 0x02, 0x9a, 0x0a, 0xb9, 0xd1, 0x83, 0x59, 0x68, 0xaf, 0xe2, 0xbc, 0x6d, 0x97, 0xb1, 0x9b,
	0x01, 0x5e, 0x83, 0xf6, 0xbf, 0x25, 0x8e, 0x19, 0x9e, 0x32, 0xea, 0x8b, 0x34, 0x63, 0x54, 0xc5,
	0x82, 0x87, 0x42, 0xaa, 0x1c, 0x98, 0xeb, 0xef, 0xaf, 0xa3, 0x33, 0x14, 0xfc, 0x03, 0xbe, 0x96,
	0xc4, 0x44, 0xc8, 0x84, 0xca, 0x9d, 0x6f, 0x6a, 0xff, 0x44, 0x0a, 0xff, 0x02, 0x3b, 0xc0, 0xeb,
	0xe3, 0x9d, 0x2d, 0x6d, 0x9e, 0xcc, 0xb7, 0x7f, 0x25, 0xc0, 0xeb, 0x6a, 0xb7, 0x9a, 0x76, 0x0e,
	0xea, 0xff, 0x57, 0xa0, 0x27, 0x64, 0xe4, 0xe2, 0x0c, 0x93, 0x19, 0xdd, 0x5f, 0x33, 0xce, 0xd2,
	0x72, 0xd5, 0x0c, 0xf3, 0x68, 0x81, 0x23, 0xea, 0xe2, 0x88, 0x72, 0x15, 0x1a, 0x8f, 0xdf, 0x3e,
	0x40, 0xaf, 0x80, 0xbc, 0x1d, 0xe4, 0x69, 0xc8, 0x5b, 0x0e, 0x9e, 0xcd, 0xd6, 0x5d, 0xb2, 0x99,
	0x14, 0xe7, 0xdd, 0xe6, 0x58, 0x58, 0xdc, 0x9a, 0x69, 0x43, 0xdf, 0x9f, 0xdf, 0xef, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0xab, 0x81, 0x70, 0x4f, 0x03, 0x00, 0x00,
}