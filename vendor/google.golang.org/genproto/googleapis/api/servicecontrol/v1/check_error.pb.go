// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/servicecontrol/v1/check_error.proto

package servicecontrol // import "google.golang.org/genproto/googleapis/api/servicecontrol/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Error codes for Check responses.
type CheckError_Code int32

const (
	// This is never used in `CheckResponse`.
	CheckError_ERROR_CODE_UNSPECIFIED CheckError_Code = 0
	// The consumer's project id was not found.
	// Same as [google.rpc.Code.NOT_FOUND][].
	CheckError_NOT_FOUND CheckError_Code = 5
	// The consumer doesn't have access to the specified resource.
	// Same as [google.rpc.Code.PERMISSION_DENIED][].
	CheckError_PERMISSION_DENIED CheckError_Code = 7
	// Quota check failed. Same as [google.rpc.Code.RESOURCE_EXHAUSTED][].
	CheckError_RESOURCE_EXHAUSTED CheckError_Code = 8
	// The consumer hasn't activated the service.
	CheckError_SERVICE_NOT_ACTIVATED CheckError_Code = 104
	// The consumer cannot access the service because billing is disabled.
	CheckError_BILLING_DISABLED CheckError_Code = 107
	// The consumer's project has been marked as deleted (soft deletion).
	CheckError_PROJECT_DELETED CheckError_Code = 108
	// The consumer's project number or id does not represent a valid project.
	CheckError_PROJECT_INVALID CheckError_Code = 114
	// The IP address of the consumer is invalid for the specific consumer
	// project.
	CheckError_IP_ADDRESS_BLOCKED CheckError_Code = 109
	// The referer address of the consumer request is invalid for the specific
	// consumer project.
	CheckError_REFERER_BLOCKED CheckError_Code = 110
	// The client application of the consumer request is invalid for the
	// specific consumer project.
	CheckError_CLIENT_APP_BLOCKED CheckError_Code = 111
	// The API targeted by this request is invalid for the specified consumer
	// project.
	CheckError_API_TARGET_BLOCKED CheckError_Code = 122
	// The consumer's API key is invalid.
	CheckError_API_KEY_INVALID CheckError_Code = 105
	// The consumer's API Key has expired.
	CheckError_API_KEY_EXPIRED CheckError_Code = 112
	// The consumer's API Key was not found in config record.
	CheckError_API_KEY_NOT_FOUND CheckError_Code = 113
	// The backend server for looking up project id/number is unavailable.
	CheckError_NAMESPACE_LOOKUP_UNAVAILABLE CheckError_Code = 300
	// The backend server for checking service status is unavailable.
	CheckError_SERVICE_STATUS_UNAVAILABLE CheckError_Code = 301
	// The backend server for checking billing status is unavailable.
	CheckError_BILLING_STATUS_UNAVAILABLE CheckError_Code = 302
)

var CheckError_Code_name = map[int32]string{
	0:   "ERROR_CODE_UNSPECIFIED",
	5:   "NOT_FOUND",
	7:   "PERMISSION_DENIED",
	8:   "RESOURCE_EXHAUSTED",
	104: "SERVICE_NOT_ACTIVATED",
	107: "BILLING_DISABLED",
	108: "PROJECT_DELETED",
	114: "PROJECT_INVALID",
	109: "IP_ADDRESS_BLOCKED",
	110: "REFERER_BLOCKED",
	111: "CLIENT_APP_BLOCKED",
	122: "API_TARGET_BLOCKED",
	105: "API_KEY_INVALID",
	112: "API_KEY_EXPIRED",
	113: "API_KEY_NOT_FOUND",
	300: "NAMESPACE_LOOKUP_UNAVAILABLE",
	301: "SERVICE_STATUS_UNAVAILABLE",
	302: "BILLING_STATUS_UNAVAILABLE",
}
var CheckError_Code_value = map[string]int32{
	"ERROR_CODE_UNSPECIFIED":       0,
	"NOT_FOUND":                    5,
	"PERMISSION_DENIED":            7,
	"RESOURCE_EXHAUSTED":           8,
	"SERVICE_NOT_ACTIVATED":        104,
	"BILLING_DISABLED":             107,
	"PROJECT_DELETED":              108,
	"PROJECT_INVALID":              114,
	"IP_ADDRESS_BLOCKED":           109,
	"REFERER_BLOCKED":              110,
	"CLIENT_APP_BLOCKED":           111,
	"API_TARGET_BLOCKED":           122,
	"API_KEY_INVALID":              105,
	"API_KEY_EXPIRED":              112,
	"API_KEY_NOT_FOUND":            113,
	"NAMESPACE_LOOKUP_UNAVAILABLE": 300,
	"SERVICE_STATUS_UNAVAILABLE":   301,
	"BILLING_STATUS_UNAVAILABLE":   302,
}

func (x CheckError_Code) String() string {
	return proto.EnumName(CheckError_Code_name, int32(x))
}
func (CheckError_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_check_error_13686b1a5f512ccf, []int{0, 0}
}

// Defines the errors to be returned in
// [google.api.servicecontrol.v1.CheckResponse.check_errors][google.api.servicecontrol.v1.CheckResponse.check_errors].
type CheckError struct {
	// The error code.
	Code CheckError_Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.api.servicecontrol.v1.CheckError_Code" json:"code,omitempty"`
	// Free-form text providing details on the error cause of the error.
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckError) Reset()         { *m = CheckError{} }
func (m *CheckError) String() string { return proto.CompactTextString(m) }
func (*CheckError) ProtoMessage()    {}
func (*CheckError) Descriptor() ([]byte, []int) {
	return fileDescriptor_check_error_13686b1a5f512ccf, []int{0}
}
func (m *CheckError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckError.Unmarshal(m, b)
}
func (m *CheckError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckError.Marshal(b, m, deterministic)
}
func (dst *CheckError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckError.Merge(dst, src)
}
func (m *CheckError) XXX_Size() int {
	return xxx_messageInfo_CheckError.Size(m)
}
func (m *CheckError) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckError.DiscardUnknown(m)
}

var xxx_messageInfo_CheckError proto.InternalMessageInfo

func (m *CheckError) GetCode() CheckError_Code {
	if m != nil {
		return m.Code
	}
	return CheckError_ERROR_CODE_UNSPECIFIED
}

func (m *CheckError) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckError)(nil), "google.api.servicecontrol.v1.CheckError")
	proto.RegisterEnum("google.api.servicecontrol.v1.CheckError_Code", CheckError_Code_name, CheckError_Code_value)
}

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/check_error.proto", fileDescriptor_check_error_13686b1a5f512ccf)
}

var fileDescriptor_check_error_13686b1a5f512ccf = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x49, 0x29, 0x83, 0x59, 0x82, 0x05, 0xc3, 0xaa, 0x51, 0x55, 0xa2, 0xec, 0xb4, 0x0b,
	0x89, 0x06, 0x47, 0x4e, 0xae, 0xfd, 0x76, 0x78, 0xcd, 0x12, 0xcb, 0x4e, 0xaa, 0xc1, 0xc5, 0x0a,
	0x69, 0x94, 0x45, 0xeb, 0xe2, 0x92, 0x56, 0x3d, 0x70, 0xe6, 0xc3, 0x70, 0x00, 0x3e, 0x02, 0x9f,
	0x8b, 0xe3, 0xe4, 0x76, 0xfd, 0x27, 0x4d, 0x3b, 0xe6, 0xe7, 0xdf, 0xfb, 0xbc, 0xca, 0xab, 0x07,
	0x79, 0x85, 0x31, 0xc5, 0x38, 0xf7, 0xd3, 0x49, 0xe9, 0x4f, 0xf3, 0x7a, 0x5e, 0x66, 0x79, 0x66,
	0xaa, 0x59, 0x6d, 0xc6, 0xfe, 0xfc, 0xd4, 0xcf, 0xae, 0xf2, 0xec, 0x5a, 0xe7, 0x75, 0x6d, 0x6a,
	0x6f, 0x52, 0x9b, 0x99, 0xc1, 0x9d, 0xa5, 0xef, 0xa5, 0x93, 0xd2, 0xdb, 0xf5, 0xbd, 0xf9, 0x69,
	0xbb, 0xb3, 0x95, 0x96, 0x56, 0x95, 0x99, 0xa5, 0xb3, 0xd2, 0x54, 0xd3, 0xe5, 0xec, 0xf1, 0xaf,
	0x26, 0x42, 0xd4, 0x26, 0x82, 0x0d, 0xc4, 0x04, 0x35, 0x33, 0x33, 0xca, 0x8f, 0x9c, 0xae, 0x73,
	0xf2, 0xe2, 0xc3, 0x7b, 0xef, 0xa1, 0x64, 0x6f, 0x33, 0xe7, 0x51, 0x33, 0xca, 0xe5, 0x62, 0x14,
	0xb7, 0xd0, 0xde, 0x28, 0x9f, 0xa5, 0xe5, 0xf8, 0xa8, 0xd1, 0x75, 0x4e, 0xf6, 0xe5, 0xdd, 0xd7,
	0xf1, 0xbf, 0xc7, 0xa8, 0x69, 0x35, 0xdc, 0x46, 0x2d, 0x90, 0x32, 0x92, 0x9a, 0x46, 0x0c, 0x74,
	0x12, 0x2a, 0x01, 0x94, 0xf7, 0x39, 0x30, 0xf7, 0x11, 0x7e, 0x8e, 0xf6, 0xc3, 0x28, 0xd6, 0xfd,
	0x28, 0x09, 0x99, 0xfb, 0x04, 0x1f, 0xa2, 0x97, 0x02, 0xe4, 0x05, 0x57, 0x8a, 0x47, 0xa1, 0x66,
	0x10, 0x5a, 0xeb, 0x29, 0x6e, 0x21, 0x2c, 0x41, 0x45, 0x89, 0xa4, 0xa0, 0xe1, 0xf2, 0x33, 0x49,
	0x54, 0x0c, 0xcc, 0x7d, 0x86, 0xdf, 0xa0, 0x43, 0x05, 0x72, 0xc8, 0x29, 0x68, 0x9b, 0x42, 0x68,
	0xcc, 0x87, 0xc4, 0x3e, 0x5d, 0xe1, 0xd7, 0xc8, 0xed, 0xf1, 0x20, 0xe0, 0xe1, 0x99, 0x66, 0x5c,
	0x91, 0x5e, 0x00, 0xcc, 0xbd, 0xc6, 0xaf, 0xd0, 0x81, 0x90, 0xd1, 0x39, 0xd0, 0x58, 0x33, 0x08,
	0xc0, 0xaa, 0xe3, 0x6d, 0xc8, 0xc3, 0x21, 0x09, 0x38, 0x73, 0x6b, 0xbb, 0x92, 0x0b, 0x4d, 0x18,
	0x93, 0xa0, 0x94, 0xee, 0x05, 0x11, 0x1d, 0x00, 0x73, 0x6f, 0xac, 0x2c, 0xa1, 0x0f, 0x12, 0xe4,
	0x1a, 0x56, 0x56, 0xa6, 0x01, 0x87, 0x30, 0xd6, 0x44, 0x88, 0x35, 0x37, 0x96, 0x13, 0xc1, 0x75,
	0x4c, 0xe4, 0x19, 0xc4, 0x6b, 0xfe, 0xc3, 0x86, 0x58, 0x3e, 0x80, 0x2f, 0xeb, 0x8d, 0xe5, 0x36,
	0x84, 0x4b, 0xc1, 0x25, 0x30, 0x77, 0x62, 0x0f, 0xb2, 0x82, 0x9b, 0x3b, 0x7d, 0xc7, 0xef, 0x50,
	0x27, 0x24, 0x17, 0xa0, 0x04, 0xa1, 0xa0, 0x83, 0x28, 0x1a, 0x24, 0x42, 0x27, 0x21, 0x19, 0x12,
	0x1e, 0xd8, 0x5f, 0x75, 0x7f, 0x37, 0xf0, 0x5b, 0xd4, 0x5e, 0xdd, 0x46, 0xc5, 0x24, 0x4e, 0xd4,
	0x8e, 0xf0, 0x67, 0x21, 0xac, 0x2e, 0x74, 0x8f, 0xf0, 0xb7, 0xd1, 0xfb, 0xe9, 0xa0, 0x6e, 0x66,
	0x6e, 0x1e, 0xec, 0x44, 0xef, 0x60, 0x53, 0x0a, 0x61, 0x0b, 0x26, 0x9c, 0xaf, 0xe7, 0x77, 0x03,
	0x85, 0x19, 0xa7, 0x55, 0xe1, 0x99, 0xba, 0xf0, 0x8b, 0xbc, 0x5a, 0xd4, 0xcf, 0x5f, 0x3e, 0xa5,
	0x93, 0x72, 0x7a, 0x7f, 0xdb, 0x3f, 0xed, 0x92, 0xff, 0x8e, 0xf3, 0x6d, 0x6f, 0x31, 0xf9, 0xf1,
	0x36, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf4, 0xc8, 0x44, 0x26, 0x03, 0x00, 0x00,
}
