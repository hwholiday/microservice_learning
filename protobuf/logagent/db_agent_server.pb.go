// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db_agent_server.proto

/*
Package logagent is a generated protocol buffer package.

It is generated from these files:
	db_agent_server.proto

It has these top-level messages:
	Log
*/
package logagent

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

type Log struct {
	Time     int64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Data     string `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Filename string `protobuf:"bytes,4,opt,name=filename" json:"filename,omitempty"`
	Line     string `protobuf:"bytes,5,opt,name=line" json:"line,omitempty"`
	Method   string `protobuf:"bytes,6,opt,name=method" json:"method,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Log) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Log) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Log) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *Log) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func init() {
	proto.RegisterType((*Log)(nil), "logagent.Log")
}

func init() { proto.RegisterFile("db_agent_server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xbd, 0xaa, 0xc2, 0x40,
	0x10, 0x46, 0xd9, 0x9b, 0x1f, 0x72, 0xa7, 0x1c, 0x54, 0x06, 0xab, 0x60, 0x95, 0xca, 0xc6, 0xd7,
	0xb0, 0xca, 0x0b, 0x84, 0x0d, 0x19, 0x63, 0x20, 0xbb, 0x23, 0xe3, 0xe2, 0x4b, 0xf8, 0xd2, 0xb2,
	0x13, 0xb1, 0xfb, 0xce, 0xe1, 0x14, 0x1f, 0xec, 0xa7, 0x71, 0xf0, 0x33, 0xc7, 0x34, 0x3c, 0x59,
	0x5f, 0xac, 0xe7, 0x87, 0x4a, 0x12, 0x6c, 0x56, 0x99, 0x4d, 0x9f, 0xde, 0x0e, 0x8a, 0xab, 0xcc,
	0x88, 0x50, 0xa6, 0x25, 0x30, 0xb9, 0xd6, 0x75, 0x45, 0x6f, 0x1b, 0x77, 0x50, 0xb1, 0xaa, 0x28,
	0xfd, 0xb5, 0xae, 0xfb, 0xef, 0x37, 0xc8, 0xe5, 0xe4, 0x93, 0xa7, 0xc2, 0xa4, 0x6d, 0x3c, 0x42,
	0x73, 0x5b, 0x56, 0x8e, 0x3e, 0x30, 0x95, 0xe6, 0x7f, 0x9c, 0xfb, 0x75, 0x89, 0x4c, 0xd5, 0xd6,
	0xe7, 0x8d, 0x07, 0xa8, 0x03, 0xa7, 0xbb, 0x4c, 0x54, 0x9b, 0xfd, 0xd2, 0x58, 0xdb, 0xbd, 0xcb,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x4c, 0x62, 0x68, 0xb7, 0x00, 0x00, 0x00,
}
