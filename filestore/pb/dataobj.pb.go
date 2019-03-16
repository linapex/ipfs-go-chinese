
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:42</date>
//</624460175868170240>

//由原基因gogo生成的代码。不要编辑。
//来源：filestore/pb/dataobj.proto

package datastore_pb

import proto "gx/ipfs/QmdxUuburamoF6zF9qjeQC4WYcWGbWuRmdLacMEsW8ioD8/gogo-protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.GoGoProtoPackageIsVersion2 //请升级proto包

type DataObj struct {
	FilePath             string   `protobuf:"bytes,1,opt,name=FilePath" json:"FilePath"`
	Offset               uint64   `protobuf:"varint,2,opt,name=Offset" json:"Offset"`
	Size_                uint64   `protobuf:"varint,3,opt,name=Size" json:"Size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataObj) Reset()         { *m = DataObj{} }
func (m *DataObj) String() string { return proto.CompactTextString(m) }
func (*DataObj) ProtoMessage()    {}
func (*DataObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataobj_216c555249812eeb, []int{0}
}
func (m *DataObj) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataObj.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DataObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataObj.Merge(dst, src)
}
func (m *DataObj) XXX_Size() int {
	return m.Size()
}
func (m *DataObj) XXX_DiscardUnknown() {
	xxx_messageInfo_DataObj.DiscardUnknown(m)
}

var xxx_messageInfo_DataObj proto.InternalMessageInfo

func (m *DataObj) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *DataObj) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DataObj) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func init() {
	proto.RegisterType((*DataObj)(nil), "datastore.pb.DataObj")
}
func (m *DataObj) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataObj) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintDataobj(dAtA, i, uint64(len(m.FilePath)))
	i += copy(dAtA[i:], m.FilePath)
	dAtA[i] = 0x10
	i++
	i = encodeVarintDataobj(dAtA, i, uint64(m.Offset))
	dAtA[i] = 0x18
	i++
	i = encodeVarintDataobj(dAtA, i, uint64(m.Size_))
	return i, nil
}

func encodeVarintDataobj(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DataObj) Size() (n int) {
	var l int
	_ = l
	l = len(m.FilePath)
	n += 1 + l + sovDataobj(uint64(l))
	n += 1 + sovDataobj(uint64(m.Offset))
	n += 1 + sovDataobj(uint64(m.Size_))
	return n
}

func sovDataobj(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDataobj(x uint64) (n int) {
	return sovDataobj(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataObj) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataobj
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataObj: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataObj: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataobj
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataobj
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataobj
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataobj
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDataobj(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDataobj
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
func skipDataobj(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataobj
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
					return 0, ErrIntOverflowDataobj
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataobj
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDataobj
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDataobj
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDataobj(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDataobj = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataobj   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("filestore/pb/dataobj.proto", fileDescriptor_dataobj_216c555249812eeb) }

var fileDescriptor_dataobj_216c555249812eeb = []byte{
//gzip文件描述符或协议的151字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcb, 0xcc, 0x49,
	0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2f, 0x48, 0xd2, 0x4f, 0x49, 0x2c, 0x49, 0xcc, 0x4f, 0xca,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x71, 0xc1, 0x72, 0x7a, 0x05, 0x49, 0x4a,
	0xc9, 0x5c, 0xec, 0x2e, 0x89, 0x25, 0x89, 0xfe, 0x49, 0x59, 0x42, 0x0a, 0x5c, 0x1c, 0x6e, 0x99,
	0x39, 0xa9, 0x01, 0x89, 0x25, 0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x2c, 0x27, 0xee,
	0xc9, 0x33, 0x04, 0xc1, 0x45, 0x85, 0x64, 0xb8, 0xd8, 0xfc, 0xd3, 0xd2, 0x8a, 0x53, 0x4b, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x58, 0xa0, 0xf2, 0x50, 0x31, 0x21, 0x09, 0x2e, 0x96, 0xe0, 0xcc, 0xaa,
	0x54, 0x09, 0x66, 0x24, 0x39, 0xb0, 0x88, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0xe7, 0x83, 0xa2, 0xa1, 0x00, 0x00, 0x00,
}

