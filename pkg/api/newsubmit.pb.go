// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/api/newsubmit.proto

package api

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "k8s.io/api/core/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("pkg/api/newsubmit.proto", fileDescriptor_9d4608f63c9100a9) }

var fileDescriptor_9d4608f63c9100a9 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8d, 0x31, 0x4e, 0xc5, 0x30,
	0x0c, 0x86, 0x13, 0x90, 0x18, 0x18, 0x11, 0x12, 0xd2, 0x03, 0x79, 0x60, 0x27, 0x11, 0x62, 0x80,
	0xf5, 0xb1, 0xb1, 0x71, 0x85, 0xa4, 0xcf, 0x84, 0x50, 0x12, 0x47, 0xa9, 0x0b, 0x62, 0xe3, 0x08,
	0x9c, 0x83, 0x93, 0x30, 0x76, 0xec, 0x08, 0xe9, 0x45, 0x50, 0xd3, 0x4a, 0x6f, 0xf3, 0x6f, 0x7f,
	0x9f, 0xff, 0xe3, 0xb3, 0xd4, 0x3a, 0x6d, 0x92, 0xd7, 0x11, 0xdf, 0xbb, 0xde, 0x06, 0xcf, 0x2a,
	0x65, 0x62, 0x3a, 0x39, 0x34, 0xc9, 0x6f, 0xce, 0x1d, 0x91, 0x7b, 0x45, 0x5d, 0x57, 0xb6, 0x7f,
	0xd2, 0x18, 0x12, 0x7f, 0x2c, 0xc4, 0xe6, 0xb2, 0xbd, 0xeb, 0x94, 0xa7, 0x6a, 0x37, 0x94, 0x51,
	0xbf, 0x5d, 0x6b, 0x87, 0x11, 0xb3, 0x61, 0xdc, 0xad, 0xcc, 0xc5, 0xfa, 0x60, 0x66, 0x4c, 0x8c,
	0xc4, 0x86, 0x3d, 0xc5, 0x6e, 0xbd, 0x5e, 0x39, 0xcf, 0xcf, 0xbd, 0x55, 0x0d, 0x05, 0xed, 0xc8,
	0xd1, 0xbe, 0x67, 0x4e, 0x35, 0xd4, 0x69, 0xc1, 0xef, 0x6f, 0xc7, 0x3f, 0x10, 0x9f, 0x05, 0xe4,
	0x4f, 0x01, 0x39, 0x14, 0x90, 0xbf, 0x05, 0xe4, 0xd7, 0x04, 0x62, 0x98, 0x40, 0x8c, 0x13, 0x88,
	0xef, 0x83, 0xd3, 0x6d, 0x0e, 0x66, 0x67, 0x1e, 0x33, 0xbd, 0x60, 0xc3, 0xea, 0x81, 0xd4, 0x36,
	0x79, 0x7b, 0x54, 0xfd, 0x9b, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x8c, 0x5d, 0x89, 0xed,
	0x00, 0x00, 0x00,
}