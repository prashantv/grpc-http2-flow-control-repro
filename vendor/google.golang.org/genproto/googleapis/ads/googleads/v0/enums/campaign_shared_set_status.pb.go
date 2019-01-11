// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/campaign_shared_set_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enum listing the possible campaign shared set statuses.
type CampaignSharedSetStatusEnum_CampaignSharedSetStatus int32

const (
	// Not specified.
	CampaignSharedSetStatusEnum_UNSPECIFIED CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignSharedSetStatusEnum_UNKNOWN CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 1
	// The campaign shared set is enabled.
	CampaignSharedSetStatusEnum_ENABLED CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 2
	// The campaign shared set is removed and can no longer be used.
	CampaignSharedSetStatusEnum_REMOVED CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 3
)

var CampaignSharedSetStatusEnum_CampaignSharedSetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}
var CampaignSharedSetStatusEnum_CampaignSharedSetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x CampaignSharedSetStatusEnum_CampaignSharedSetStatus) String() string {
	return proto.EnumName(CampaignSharedSetStatusEnum_CampaignSharedSetStatus_name, int32(x))
}
func (CampaignSharedSetStatusEnum_CampaignSharedSetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_campaign_shared_set_status_82017bcab3cc753a, []int{0, 0}
}

// Container for enum describing types of campaign shared set statuses.
type CampaignSharedSetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignSharedSetStatusEnum) Reset()         { *m = CampaignSharedSetStatusEnum{} }
func (m *CampaignSharedSetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetStatusEnum) ProtoMessage()    {}
func (*CampaignSharedSetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_shared_set_status_82017bcab3cc753a, []int{0}
}
func (m *CampaignSharedSetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetStatusEnum.Unmarshal(m, b)
}
func (m *CampaignSharedSetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetStatusEnum.Marshal(b, m, deterministic)
}
func (dst *CampaignSharedSetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetStatusEnum.Merge(dst, src)
}
func (m *CampaignSharedSetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetStatusEnum.Size(m)
}
func (m *CampaignSharedSetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignSharedSetStatusEnum)(nil), "google.ads.googleads.v0.enums.CampaignSharedSetStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus", CampaignSharedSetStatusEnum_CampaignSharedSetStatus_name, CampaignSharedSetStatusEnum_CampaignSharedSetStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/campaign_shared_set_status.proto", fileDescriptor_campaign_shared_set_status_82017bcab3cc753a)
}

var fileDescriptor_campaign_shared_set_status_82017bcab3cc753a = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xe4, 0xc4, 0xdc, 0x82, 0xc4, 0xcc, 0xf4, 0xbc, 0xf8, 0xe2, 0x8c, 0xc4,
	0xa2, 0xd4, 0x94, 0xf8, 0xe2, 0xd4, 0x92, 0xf8, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x26, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x7e, 0xbd,
	0x32, 0x03, 0x3d, 0xb0, 0x7e, 0xa5, 0x02, 0x2e, 0x69, 0x67, 0xa8, 0x11, 0xc1, 0x60, 0x13, 0x82,
	0x53, 0x4b, 0x82, 0xc1, 0xfa, 0x5d, 0xf3, 0x4a, 0x73, 0x95, 0x02, 0xb9, 0xc4, 0x71, 0x48, 0x0b,
	0xf1, 0x73, 0x71, 0x87, 0xfa, 0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba, 0xba, 0x08, 0x30,
	0x08, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x82, 0x38, 0xae,
	0x7e, 0x8e, 0x4e, 0x3e, 0xae, 0x2e, 0x02, 0x4c, 0x20, 0x4e, 0x90, 0xab, 0xaf, 0x7f, 0x98, 0xab,
	0x8b, 0x00, 0xb3, 0xd3, 0x59, 0x46, 0x2e, 0xc5, 0xe4, 0xfc, 0x5c, 0x3d, 0xbc, 0xee, 0x72, 0x92,
	0xc1, 0x61, 0x6d, 0x00, 0xc8, 0x53, 0x01, 0x8c, 0x51, 0x4e, 0x50, 0xed, 0xe9, 0xf9, 0x39, 0x89,
	0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x2f, 0xc3, 0x82, 0xa9, 0x20,
	0xb3, 0x18, 0x47, 0xa8, 0x59, 0x83, 0xc9, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e, 0xab, 0x98, 0x64,
	0xdd, 0x21, 0x46, 0x39, 0xa6, 0x14, 0xeb, 0x41, 0x98, 0x20, 0x56, 0x98, 0x81, 0x1e, 0x28, 0x04,
	0x8a, 0x4f, 0xc1, 0xe4, 0x63, 0x1c, 0x53, 0x8a, 0x63, 0xe0, 0xf2, 0x31, 0x61, 0x06, 0x31, 0x60,
	0xf9, 0x24, 0x36, 0xb0, 0xa5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xcf, 0xee, 0xb5,
	0xa9, 0x01, 0x00, 0x00,
}
