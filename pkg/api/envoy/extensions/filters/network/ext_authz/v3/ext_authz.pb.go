// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/network/ext_authz/v3/ext_authz.proto

package envoy_extensions_filters_network_ext_authz_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// External Authorization filter calls out to an external service over the
// gRPC Authorization API defined by
// :ref:`CheckRequest <envoy_api_msg_service.auth.v3.CheckRequest>`.
// A failed check will cause this filter to close the TCP connection.
type ExtAuthz struct {
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The external authorization gRPC service configuration.
	// The default timeout is set to 200ms by this filter.
	GrpcService *v3.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. When it is set to true, Envoy will also allow traffic in case of
	// communication failure between authorization service and the proxy.
	// Defaults to false.
	FailureModeAllow bool `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Specifies if the peer certificate is sent to the external service.
	//
	// When this field is true, Envoy will include the peer X.509 certificate, if available, in the
	// :ref:`certificate<envoy_api_field_service.auth.v3.AttributeContext.Peer.certificate>`.
	IncludePeerCertificate bool     `protobuf:"varint,4,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_c579db07d85696d8, []int{0}
}
func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return m.Size()
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

func (m *ExtAuthz) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ExtAuthz) GetGrpcService() *v3.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func (m *ExtAuthz) GetIncludePeerCertificate() bool {
	if m != nil {
		return m.IncludePeerCertificate
	}
	return false
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.extensions.filters.network.ext_authz.v3.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/ext_authz/v3/ext_authz.proto", fileDescriptor_c579db07d85696d8)
}

var fileDescriptor_c579db07d85696d8 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0xae, 0xd3, 0x30,
	0x18, 0x85, 0xe5, 0x70, 0x75, 0xb9, 0xb8, 0x20, 0x55, 0x19, 0x20, 0xaa, 0x44, 0x94, 0xb2, 0x90,
	0x01, 0x1c, 0xd1, 0x2c, 0x15, 0x88, 0xa1, 0x01, 0xc4, 0x84, 0x14, 0x95, 0x07, 0x88, 0x5c, 0xe7,
	0x4f, 0xb0, 0x08, 0x76, 0xe4, 0x38, 0x69, 0xca, 0xc4, 0xc8, 0x33, 0x20, 0x9e, 0x84, 0x1d, 0xa9,
	0x23, 0xbc, 0x01, 0xea, 0x63, 0x30, 0x21, 0xc7, 0xa9, 0x02, 0x82, 0x85, 0x2d, 0xce, 0xe7, 0x73,
	0x7c, 0xfe, 0xff, 0xe0, 0xa7, 0x20, 0x3a, 0x79, 0x88, 0xa0, 0xd7, 0x20, 0x1a, 0x2e, 0x45, 0x13,
	0x15, 0xbc, 0xd2, 0xa0, 0x9a, 0x48, 0x80, 0xde, 0x4b, 0xf5, 0xd6, 0xa0, 0x8c, 0xb6, 0xfa, 0xcd,
	0xfb, 0xa8, 0x8b, 0xa7, 0x03, 0xa9, 0x95, 0xd4, 0xd2, 0x7d, 0x38, 0xc8, 0xc9, 0x24, 0x27, 0xa3,
	0x9c, 0x8c, 0x72, 0x32, 0x29, 0xba, 0x78, 0x71, 0xdf, 0xbe, 0xc6, 0xa4, 0x28, 0x78, 0x19, 0x31,
	0xa9, 0xc0, 0x98, 0x96, 0xaa, 0x66, 0x59, 0x03, 0xaa, 0xe3, 0x0c, 0xac, 0xef, 0xe2, 0x6e, 0x9b,
	0xd7, 0x34, 0xa2, 0x42, 0x48, 0x4d, 0xf5, 0x10, 0xab, 0xd1, 0x54, 0xb7, 0xcd, 0x88, 0x97, 0x7f,
	0xe1, 0x0e, 0x94, 0x79, 0x9f, 0x8b, 0x72, 0xbc, 0x72, 0xa7, 0xa3, 0x15, 0xcf, 0xa9, 0x86, 0xe8,
	0xfc, 0x61, 0xc1, 0xbd, 0xcf, 0x0e, 0xbe, 0x7a, 0xd1, 0xeb, 0x8d, 0xc9, 0xe4, 0x86, 0x78, 0x66,
	0x8c, 0xb3, 0x5a, 0x41, 0xc1, 0x7b, 0x0f, 0x05, 0x28, 0xbc, 0x91, 0x5c, 0xff, 0x99, 0x5c, 0x28,
	0x27, 0x40, 0x5b, 0x6c, 0x58, 0x3a, 0x20, 0xf7, 0x39, 0xbe, 0xf9, 0x7b, 0x4e, 0xcf, 0x09, 0x50,
	0x38, 0x5b, 0x2d, 0x89, 0x5d, 0x80, 0x9d, 0x88, 0x98, 0x89, 0x48, 0x17, 0x93, 0x97, 0xaa, 0x66,
	0xaf, 0xed, 0xc5, 0xed, 0xac, 0x9c, 0x0e, 0xee, 0x03, 0xec, 0x16, 0x94, 0x57, 0xad, 0x82, 0xec,
	0x9d, 0xcc, 0x21, 0xa3, 0x55, 0x25, 0xf7, 0xde, 0xb5, 0x00, 0x85, 0x57, 0xdb, 0xf9, 0x48, 0x5e,
	0xc9, 0x1c, 0x36, 0xe6, 0xbf, 0xbb, 0xc6, 0x1e, 0x17, 0xac, 0x6a, 0x73, 0xc8, 0x6a, 0x00, 0x95,
	0x31, 0x50, 0x9a, 0x17, 0x9c, 0x51, 0x0d, 0xde, 0xc5, 0xa0, 0xb9, 0x3d, 0xf2, 0x14, 0x40, 0x3d,
	0x9b, 0xe8, 0xe3, 0xf5, 0xa7, 0xaf, 0x1f, 0xfd, 0x18, 0x3f, 0xfa, 0x23, 0x9d, 0xad, 0xe6, 0x5f,
	0xcd, 0xac, 0xc8, 0x79, 0x23, 0xc9, 0xee, 0x78, 0xf2, 0xd1, 0xb7, 0x93, 0x8f, 0x7e, 0x9c, 0x7c,
	0xf4, 0xe5, 0xc3, 0xf1, 0xfb, 0xa5, 0x33, 0x47, 0xf8, 0x09, 0x97, 0x76, 0xd2, 0x5a, 0xc9, 0xfe,
	0x40, 0xfe, 0xab, 0xf5, 0xe4, 0xd6, 0xd9, 0x3c, 0x35, 0x05, 0xa4, 0x68, 0x77, 0x39, 0x34, 0x11,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x6c, 0xbb, 0x01, 0x7d, 0x02, 0x00, 0x00,
}

func (m *ExtAuthz) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtAuthz) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtAuthz) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IncludePeerCertificate {
		i--
		if m.IncludePeerCertificate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.FailureModeAllow {
		i--
		if m.FailureModeAllow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.GrpcService != nil {
		{
			size, err := m.GrpcService.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExtAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintExtAuthz(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtAuthz) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovExtAuthz(uint64(l))
	}
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovExtAuthz(uint64(l))
	}
	if m.FailureModeAllow {
		n += 2
	}
	if m.IncludePeerCertificate {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExtAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtAuthz(x uint64) (n int) {
	return sovExtAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtAuthz) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtAuthz
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
			return fmt.Errorf("proto: ExtAuthz: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtAuthz: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
				return ErrInvalidLengthExtAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
				return ErrInvalidLengthExtAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &v3.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeAllow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
			m.FailureModeAllow = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncludePeerCertificate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
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
			m.IncludePeerCertificate = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipExtAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExtAuthz
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExtAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtAuthz
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
					return 0, ErrIntOverflowExtAuthz
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
					return 0, ErrIntOverflowExtAuthz
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
				return 0, ErrInvalidLengthExtAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtAuthz = fmt.Errorf("proto: unexpected end of group")
)
