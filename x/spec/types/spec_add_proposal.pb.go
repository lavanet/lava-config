// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/spec/spec_add_proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type SpecAddProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Specs       []Spec `protobuf:"bytes,3,rep,name=specs,proto3" json:"specs"`
}

func (m *SpecAddProposal) Reset()      { *m = SpecAddProposal{} }
func (*SpecAddProposal) ProtoMessage() {}
func (*SpecAddProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_90ccba27a5c90c8a, []int{0}
}
func (m *SpecAddProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpecAddProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpecAddProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpecAddProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecAddProposal.Merge(m, src)
}
func (m *SpecAddProposal) XXX_Size() int {
	return m.Size()
}
func (m *SpecAddProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecAddProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SpecAddProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SpecAddProposal)(nil), "lavanet.lava.spec.SpecAddProposal")
}

func init() {
	proto.RegisterFile("lavanet/lava/spec/spec_add_proposal.proto", fileDescriptor_90ccba27a5c90c8a)
}

var fileDescriptor_90ccba27a5c90c8a = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xc5, 0x05, 0xa9, 0xc9, 0x60, 0x22, 0x3e, 0x31, 0x25,
	0x25, 0xbe, 0xa0, 0x28, 0xbf, 0x20, 0xbf, 0x38, 0x31, 0x47, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f,
	0x48, 0x10, 0xaa, 0x54, 0x0f, 0x44, 0xeb, 0x81, 0x54, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83,
	0x65, 0xf5, 0x41, 0x2c, 0x88, 0x42, 0x29, 0x19, 0xec, 0x66, 0x42, 0x64, 0x95, 0x3a, 0x18, 0xb9,
	0xf8, 0x83, 0x0b, 0x52, 0x93, 0x1d, 0x53, 0x52, 0x02, 0xa0, 0x16, 0x08, 0x89, 0x70, 0xb1, 0x96,
	0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x0a, 0x5c,
	0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x60, 0x39,
	0x64, 0x21, 0x21, 0x63, 0x2e, 0x56, 0x90, 0xc9, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0xe2, 0x7a, 0x18, 0x4e, 0xd4, 0x03, 0x59, 0xe5, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x44,
	0xad, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86, 0x19, 0x0b, 0xe4, 0x19, 0x9c, 0x9c, 0x56, 0x3c, 0x92,
	0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x95, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x14, 0x1f, 0x55, 0x40, 0xfc, 0x54, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xab, 0xcf, 0x87,
	0x10, 0x49, 0x01, 0x00, 0x00,
}

func (this *SpecAddProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecAddProposal)
	if !ok {
		that2, ok := that.(SpecAddProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Specs) != len(that1.Specs) {
		return false
	}
	for i := range this.Specs {
		if !this.Specs[i].Equal(&that1.Specs[i]) {
			return false
		}
	}
	return true
}
func (m *SpecAddProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecAddProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpecAddProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Specs) > 0 {
		for iNdEx := len(m.Specs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Specs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpecAddProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintSpecAddProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSpecAddProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpecAddProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpecAddProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SpecAddProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSpecAddProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovSpecAddProposal(uint64(l))
	}
	if len(m.Specs) > 0 {
		for _, e := range m.Specs {
			l = e.Size()
			n += 1 + l + sovSpecAddProposal(uint64(l))
		}
	}
	return n
}

func sovSpecAddProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpecAddProposal(x uint64) (n int) {
	return sovSpecAddProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpecAddProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpecAddProposal
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
			return fmt.Errorf("proto: SpecAddProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecAddProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpecAddProposal
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
				return ErrInvalidLengthSpecAddProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpecAddProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpecAddProposal
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
				return ErrInvalidLengthSpecAddProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpecAddProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Specs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpecAddProposal
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
				return ErrInvalidLengthSpecAddProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpecAddProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Specs = append(m.Specs, Spec{})
			if err := m.Specs[len(m.Specs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpecAddProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpecAddProposal
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
func skipSpecAddProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpecAddProposal
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
					return 0, ErrIntOverflowSpecAddProposal
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
					return 0, ErrIntOverflowSpecAddProposal
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
				return 0, ErrInvalidLengthSpecAddProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpecAddProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpecAddProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpecAddProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpecAddProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpecAddProposal = fmt.Errorf("proto: unexpected end of group")
)
