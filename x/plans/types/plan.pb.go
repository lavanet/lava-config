// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/plans/plan.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lavanet/lava/x/spec/types"
	io "io"
	math "math"
	math_bits "math/bits"
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

// The geolocation values are encoded as bits in a bitmask, with two special values:
// GLS is set to 0 so it will be restrictive with the AND operator.
// GL is set to -1 so it will be permissive with the AND operator.
type Geolocation int32

const (
	Geolocation_GLS Geolocation = 0
	Geolocation_USC Geolocation = 1
	Geolocation_EU  Geolocation = 2
	Geolocation_USE Geolocation = 4
	Geolocation_USW Geolocation = 8
	Geolocation_AF  Geolocation = 16
	Geolocation_AS  Geolocation = 32
	Geolocation_AU  Geolocation = 64
	Geolocation_GL  Geolocation = 65535
)

var Geolocation_name = map[int32]string{
	0:     "GLS",
	1:     "USC",
	2:     "EU",
	4:     "USE",
	8:     "USW",
	16:    "AF",
	32:    "AS",
	64:    "AU",
	65535: "GL",
}

var Geolocation_value = map[string]int32{
	"GLS": 0,
	"USC": 1,
	"EU":  2,
	"USE": 4,
	"USW": 8,
	"AF":  16,
	"AS":  32,
	"AU":  64,
	"GL":  65535,
}

func (x Geolocation) String() string {
	return proto.EnumName(Geolocation_name, int32(x))
}

func (Geolocation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_64c3707a3b09a2e5, []int{0}
}

type Plan struct {
	Index                    string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Block                    uint64     `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	Price                    types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	AllowOveruse             bool       `protobuf:"varint,8,opt,name=allow_overuse,json=allowOveruse,proto3" json:"allow_overuse,omitempty"`
	OveruseRate              uint64     `protobuf:"varint,9,opt,name=overuse_rate,json=overuseRate,proto3" json:"overuse_rate,omitempty"`
	Description              string     `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Type                     string     `protobuf:"bytes,12,opt,name=type,proto3" json:"type,omitempty"`
	AnnualDiscountPercentage uint64     `protobuf:"varint,13,opt,name=annual_discount_percentage,json=annualDiscountPercentage,proto3" json:"annual_discount_percentage,omitempty"`
	PlanPolicy               Policy     `protobuf:"bytes,14,opt,name=plan_policy,json=planPolicy,proto3" json:"plan_policy"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_64c3707a3b09a2e5, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Plan) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Plan) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func (m *Plan) GetAllowOveruse() bool {
	if m != nil {
		return m.AllowOveruse
	}
	return false
}

func (m *Plan) GetOveruseRate() uint64 {
	if m != nil {
		return m.OveruseRate
	}
	return 0
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Plan) GetAnnualDiscountPercentage() uint64 {
	if m != nil {
		return m.AnnualDiscountPercentage
	}
	return 0
}

func (m *Plan) GetPlanPolicy() Policy {
	if m != nil {
		return m.PlanPolicy
	}
	return Policy{}
}

func init() {
	proto.RegisterEnum("lavanet.lava.plans.Geolocation", Geolocation_name, Geolocation_value)
	proto.RegisterType((*Plan)(nil), "lavanet.lava.plans.Plan")
}

func init() { proto.RegisterFile("lavanet/lava/plans/plan.proto", fileDescriptor_64c3707a3b09a2e5) }

var fileDescriptor_64c3707a3b09a2e5 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xcd, 0x26, 0x4e, 0xb2, 0xf5, 0xa6, 0xc8, 0xb2, 0x7a, 0x30, 0x91, 0xd8, 0x2e, 0x20, 0x50,
	0xc4, 0xc1, 0xab, 0x82, 0xb8, 0x71, 0xa0, 0x0d, 0x25, 0x52, 0x54, 0x89, 0x28, 0x51, 0x84, 0x04,
	0x87, 0x95, 0xe3, 0x58, 0xc1, 0xc2, 0x5d, 0xaf, 0x76, 0x9d, 0xd0, 0xfe, 0x00, 0x67, 0x3e, 0x83,
	0x4f, 0xe9, 0xb1, 0x47, 0x4e, 0x08, 0x25, 0x1f, 0x52, 0x64, 0x7b, 0x41, 0x54, 0xf4, 0xb2, 0x33,
	0xf3, 0xde, 0xd3, 0xce, 0x9b, 0xf1, 0xc0, 0x07, 0x8a, 0x6d, 0x58, 0x2e, 0x4c, 0x6a, 0x63, 0x5a,
	0x28, 0x96, 0x57, 0xee, 0x4b, 0x8b, 0x52, 0x1b, 0x8d, 0x71, 0x4d, 0x53, 0x1b, 0xa9, 0xa3, 0xfb,
	0x07, 0x2b, 0xbd, 0xd2, 0x8e, 0x4e, 0x6d, 0xe6, 0x95, 0xfd, 0x98, 0xeb, 0xea, 0x5c, 0x57, 0xe9,
	0x82, 0x55, 0x22, 0xdd, 0x1c, 0x2d, 0x84, 0x61, 0x47, 0x29, 0xd7, 0xb2, 0xfe, 0x53, 0xff, 0xe9,
	0xad, 0x46, 0x55, 0x21, 0x78, 0xca, 0x0a, 0x99, 0x71, 0xad, 0x94, 0xe0, 0x46, 0xea, 0x3f, 0xba,
	0xc3, 0xbb, 0x0c, 0x69, 0x25, 0xf9, 0xa5, 0x17, 0x3c, 0xfa, 0xda, 0x82, 0x60, 0xa2, 0x58, 0x8e,
	0x0f, 0x60, 0x5b, 0xe6, 0x4b, 0x71, 0x41, 0x82, 0x24, 0x18, 0xec, 0x4d, 0x7d, 0x61, 0xd1, 0x85,
	0xd2, 0xfc, 0x33, 0x69, 0x25, 0xc1, 0x00, 0x4c, 0x7d, 0x81, 0x5f, 0xc2, 0x76, 0x51, 0x4a, 0x2e,
	0x08, 0x48, 0x82, 0x41, 0xf4, 0xfc, 0x3e, 0xf5, 0x6e, 0xa9, 0x75, 0x4b, 0x6b, 0xb7, 0x74, 0xa8,
	0x65, 0x7e, 0x02, 0xae, 0x7e, 0x1e, 0x36, 0xa6, 0x5e, 0x8d, 0x1f, 0xc3, 0x7d, 0xa6, 0x94, 0xfe,
	0x92, 0xe9, 0x8d, 0x28, 0xd7, 0x95, 0x20, 0x61, 0x12, 0x0c, 0xc2, 0x69, 0xcf, 0x81, 0xef, 0x3c,
	0x86, 0x1f, 0xc2, 0x5e, 0x4d, 0x67, 0x25, 0x33, 0x82, 0xec, 0xb9, 0xc6, 0x51, 0x8d, 0x4d, 0x99,
	0x11, 0x38, 0x81, 0xd1, 0x52, 0x54, 0xbc, 0x94, 0x85, 0x9d, 0x94, 0x44, 0xce, 0xf0, 0xbf, 0x10,
	0xc6, 0x10, 0x98, 0xcb, 0x42, 0x90, 0x9e, 0xa3, 0x5c, 0x8e, 0x5f, 0xc1, 0x3e, 0xcb, 0xf3, 0x35,
	0x53, 0xd9, 0x52, 0x56, 0x5c, 0xaf, 0x73, 0x93, 0x15, 0xa2, 0xe4, 0x22, 0x37, 0x6c, 0x25, 0xc8,
	0xbe, 0x6b, 0x43, 0xbc, 0xe2, 0x4d, 0x2d, 0x98, 0xfc, 0xe5, 0xf1, 0x31, 0x8c, 0xec, 0xf6, 0x32,
	0xbf, 0x3c, 0x72, 0xcf, 0x0d, 0xde, 0xa7, 0xff, 0x3f, 0x28, 0x9d, 0x38, 0x45, 0x3d, 0x39, 0xb4,
	0x98, 0x47, 0xc6, 0x20, 0x84, 0x28, 0x1a, 0x83, 0xb0, 0x89, 0x5a, 0x63, 0x10, 0xb6, 0x51, 0x67,
	0x0c, 0xc2, 0x0e, 0xea, 0x8e, 0x41, 0xd8, 0x45, 0xe1, 0xb3, 0x8f, 0x30, 0x1a, 0x09, 0xad, 0x34,
	0x67, 0x6e, 0x82, 0x2e, 0x6c, 0x8d, 0xce, 0x66, 0xa8, 0x61, 0x93, 0xf9, 0x6c, 0x88, 0x02, 0xdc,
	0x81, 0xcd, 0xd3, 0x39, 0x6a, 0x7a, 0xe0, 0x14, 0x01, 0x9f, 0xbc, 0x47, 0xa1, 0x65, 0x8e, 0xdf,
	0x22, 0xe4, 0xe2, 0x0c, 0x25, 0x2e, 0xce, 0xd1, 0x6b, 0x1c, 0xc2, 0xe6, 0xe8, 0x0c, 0xdd, 0xdc,
	0xb4, 0x4e, 0x86, 0xdf, 0xb7, 0x71, 0x70, 0xb5, 0x8d, 0x83, 0xeb, 0x6d, 0x1c, 0xfc, 0xda, 0xc6,
	0xc1, 0xb7, 0x5d, 0xdc, 0xb8, 0xde, 0xc5, 0x8d, 0x1f, 0xbb, 0xb8, 0xf1, 0xe1, 0xc9, 0x4a, 0x9a,
	0x4f, 0xeb, 0x05, 0xe5, 0xfa, 0x3c, 0xbd, 0x75, 0x2f, 0x17, 0xf5, 0xc5, 0xd8, 0xfd, 0x55, 0x8b,
	0x8e, 0xbb, 0x98, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xb8, 0xdb, 0x97, 0xe5, 0x02,
	0x00, 0x00,
}

func (this *Plan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plan)
	if !ok {
		that2, ok := that.(Plan)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Block != that1.Block {
		return false
	}
	if !this.Price.Equal(&that1.Price) {
		return false
	}
	if this.AllowOveruse != that1.AllowOveruse {
		return false
	}
	if this.OveruseRate != that1.OveruseRate {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.AnnualDiscountPercentage != that1.AnnualDiscountPercentage {
		return false
	}
	if !this.PlanPolicy.Equal(&that1.PlanPolicy) {
		return false
	}
	return true
}
func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PlanPolicy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	if m.AnnualDiscountPercentage != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.AnnualDiscountPercentage))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x5a
	}
	if m.OveruseRate != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.OveruseRate))
		i--
		dAtA[i] = 0x48
	}
	if m.AllowOveruse {
		i--
		if m.AllowOveruse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Block != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovPlan(uint64(m.Block))
	}
	l = m.Price.Size()
	n += 1 + l + sovPlan(uint64(l))
	if m.AllowOveruse {
		n += 2
	}
	if m.OveruseRate != 0 {
		n += 1 + sovPlan(uint64(m.OveruseRate))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.AnnualDiscountPercentage != 0 {
		n += 1 + sovPlan(uint64(m.AnnualDiscountPercentage))
	}
	l = m.PlanPolicy.Size()
	n += 1 + l + sovPlan(uint64(l))
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowOveruse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowOveruse = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OveruseRate", wireType)
			}
			m.OveruseRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OveruseRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualDiscountPercentage", wireType)
			}
			m.AnnualDiscountPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnnualDiscountPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PlanPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
