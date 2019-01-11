// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/string_format_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible string format errors.
type StringFormatErrorEnum_StringFormatError int32

const (
	// Enum unspecified.
	StringFormatErrorEnum_UNSPECIFIED StringFormatErrorEnum_StringFormatError = 0
	// The received error code is not known in this version.
	StringFormatErrorEnum_UNKNOWN StringFormatErrorEnum_StringFormatError = 1
	// The input string value contains disallowed characters.
	StringFormatErrorEnum_ILLEGAL_CHARS StringFormatErrorEnum_StringFormatError = 2
	// The input string value is invalid for the associated field.
	StringFormatErrorEnum_INVALID_FORMAT StringFormatErrorEnum_StringFormatError = 3
)

var StringFormatErrorEnum_StringFormatError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ILLEGAL_CHARS",
	3: "INVALID_FORMAT",
}
var StringFormatErrorEnum_StringFormatError_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ILLEGAL_CHARS":  2,
	"INVALID_FORMAT": 3,
}

func (x StringFormatErrorEnum_StringFormatError) String() string {
	return proto.EnumName(StringFormatErrorEnum_StringFormatError_name, int32(x))
}
func (StringFormatErrorEnum_StringFormatError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_string_format_error_507abb9d06602679, []int{0, 0}
}

// Container for enum describing possible string format errors.
type StringFormatErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringFormatErrorEnum) Reset()         { *m = StringFormatErrorEnum{} }
func (m *StringFormatErrorEnum) String() string { return proto.CompactTextString(m) }
func (*StringFormatErrorEnum) ProtoMessage()    {}
func (*StringFormatErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_string_format_error_507abb9d06602679, []int{0}
}
func (m *StringFormatErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringFormatErrorEnum.Unmarshal(m, b)
}
func (m *StringFormatErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringFormatErrorEnum.Marshal(b, m, deterministic)
}
func (dst *StringFormatErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringFormatErrorEnum.Merge(dst, src)
}
func (m *StringFormatErrorEnum) XXX_Size() int {
	return xxx_messageInfo_StringFormatErrorEnum.Size(m)
}
func (m *StringFormatErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_StringFormatErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_StringFormatErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StringFormatErrorEnum)(nil), "google.ads.googleads.v0.errors.StringFormatErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.StringFormatErrorEnum_StringFormatError", StringFormatErrorEnum_StringFormatError_name, StringFormatErrorEnum_StringFormatError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/string_format_error.proto", fileDescriptor_string_format_error_507abb9d06602679)
}

var fileDescriptor_string_format_error_507abb9d06602679 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xe2, 0x92, 0xa2, 0xcc, 0xbc, 0xf4, 0xf8, 0xb4, 0xfc, 0xa2, 0xdc,
	0xc4, 0x92, 0x78, 0xb0, 0xa0, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44, 0xb9, 0x5e,
	0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xa7, 0x5e, 0x99, 0x81, 0x1e, 0x44, 0xa7, 0x52, 0x21, 0x97, 0x68,
	0x30, 0x58, 0xb3, 0x1b, 0x58, 0xaf, 0x2b, 0x48, 0xd4, 0x35, 0xaf, 0x34, 0x57, 0x29, 0x82, 0x4b,
	0x10, 0x43, 0x42, 0x88, 0x9f, 0x8b, 0x3b, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9, 0xd3, 0xcd, 0xd3,
	0xd5, 0x45, 0x80, 0x41, 0x88, 0x9b, 0x8b, 0x3d, 0xd4, 0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80,
	0x51, 0x48, 0x90, 0x8b, 0xd7, 0xd3, 0xc7, 0xc7, 0xd5, 0xdd, 0xd1, 0x27, 0xde, 0xd9, 0xc3, 0x31,
	0x28, 0x58, 0x80, 0x49, 0x48, 0x88, 0x8b, 0xcf, 0xd3, 0x2f, 0xcc, 0xd1, 0xc7, 0xd3, 0x25, 0xde,
	0xcd, 0x3f, 0xc8, 0xd7, 0x31, 0x44, 0x80, 0xd9, 0xe9, 0x0c, 0x23, 0x97, 0x52, 0x72, 0x7e, 0xae,
	0x1e, 0x7e, 0x97, 0x39, 0x89, 0x61, 0x58, 0x1f, 0x00, 0xf2, 0x51, 0x00, 0x63, 0x94, 0x0b, 0x54,
	0x67, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0xd8,
	0xbf, 0xb0, 0xd0, 0x29, 0xc8, 0x2c, 0xc6, 0x15, 0x58, 0xd6, 0x10, 0x6a, 0x11, 0x13, 0xb3, 0xbb,
	0xa3, 0xe3, 0x2a, 0x26, 0x39, 0x77, 0x88, 0x61, 0x8e, 0x29, 0xc5, 0x7a, 0x10, 0x26, 0x88, 0x15,
	0x66, 0xa0, 0x07, 0xb6, 0xb2, 0xf8, 0x14, 0x4c, 0x41, 0x8c, 0x63, 0x4a, 0x71, 0x0c, 0x5c, 0x41,
	0x4c, 0x98, 0x41, 0x0c, 0x44, 0x41, 0x12, 0x1b, 0xd8, 0x62, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x92, 0xc6, 0x00, 0xa4, 0x01, 0x00, 0x00,
}
