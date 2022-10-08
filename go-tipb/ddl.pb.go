// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ddl.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DDLType int32

const (
	DDLType_TypeDDLIndex DDLType = 0
)

var DDLType_name = map[int32]string{
	0: "TypeDDLIndex",
}
var DDLType_value = map[string]int32{
	"TypeDDLIndex": 0,
}

func (x DDLType) Enum() *DDLType {
	p := new(DDLType)
	*p = x
	return p
}
func (x DDLType) String() string {
	return proto.EnumName(DDLType_name, int32(x))
}
func (x *DDLType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DDLType_value, data, "DDLType")
	if err != nil {
		return err
	}
	*x = DDLType(value)
	return nil
}
func (DDLType) EnumDescriptor() ([]byte, []int) { return fileDescriptorDdl, []int{0} }

type DDLRequest struct {
	TableId                int64         `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	Columns                []*ColumnInfo `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	Ranges                 []KeyRange    `protobuf:"bytes,3,rep,name=ranges" json:"ranges"`
	PrimaryPrefixColumnIds []int64       `protobuf:"varint,4,rep,name=primary_prefix_column_ids,json=primaryPrefixColumnIds" json:"primary_prefix_column_ids,omitempty"`
	Unique                 *bool         `protobuf:"varint,5,opt,name=unique" json:"unique,omitempty"`
	TableInfo              *TableInfo    `protobuf:"bytes,6,opt,name=table_info,json=tableInfo" json:"table_info,omitempty"`
	IndexInfo              *IndexInfo    `protobuf:"bytes,7,opt,name=index_info,json=indexInfo" json:"index_info,omitempty"`
	PrimaryColumnIds       []int64       `protobuf:"varint,8,rep,name=primary_column_ids,json=primaryColumnIds" json:"primary_column_ids,omitempty"`
	XXX_unrecognized       []byte        `json:"-"`
}

func (m *DDLRequest) Reset()                    { *m = DDLRequest{} }
func (m *DDLRequest) String() string            { return proto.CompactTextString(m) }
func (*DDLRequest) ProtoMessage()               {}
func (*DDLRequest) Descriptor() ([]byte, []int) { return fileDescriptorDdl, []int{0} }

func (m *DDLRequest) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *DDLRequest) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *DDLRequest) GetRanges() []KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

func (m *DDLRequest) GetPrimaryPrefixColumnIds() []int64 {
	if m != nil {
		return m.PrimaryPrefixColumnIds
	}
	return nil
}

func (m *DDLRequest) GetUnique() bool {
	if m != nil && m.Unique != nil {
		return *m.Unique
	}
	return false
}

func (m *DDLRequest) GetTableInfo() *TableInfo {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *DDLRequest) GetIndexInfo() *IndexInfo {
	if m != nil {
		return m.IndexInfo
	}
	return nil
}

func (m *DDLRequest) GetPrimaryColumnIds() []int64 {
	if m != nil {
		return m.PrimaryColumnIds
	}
	return nil
}

type DDLResponse struct {
	Error            *Error   `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Keys             [][]byte `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
	Values           [][]byte `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DDLResponse) Reset()                    { *m = DDLResponse{} }
func (m *DDLResponse) String() string            { return proto.CompactTextString(m) }
func (*DDLResponse) ProtoMessage()               {}
func (*DDLResponse) Descriptor() ([]byte, []int) { return fileDescriptorDdl, []int{1} }

func (m *DDLResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DDLResponse) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *DDLResponse) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*DDLRequest)(nil), "tipb.DDLRequest")
	proto.RegisterType((*DDLResponse)(nil), "tipb.DDLResponse")
	proto.RegisterEnum("tipb.DDLType", DDLType_name, DDLType_value)
}
func (m *DDLRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DDLRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintDdl(dAtA, i, uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDdl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Ranges) > 0 {
		for _, msg := range m.Ranges {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintDdl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.PrimaryPrefixColumnIds) > 0 {
		for _, num := range m.PrimaryPrefixColumnIds {
			dAtA[i] = 0x20
			i++
			i = encodeVarintDdl(dAtA, i, uint64(num))
		}
	}
	if m.Unique != nil {
		dAtA[i] = 0x28
		i++
		if *m.Unique {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.TableInfo != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDdl(dAtA, i, uint64(m.TableInfo.Size()))
		n1, err := m.TableInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.IndexInfo != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDdl(dAtA, i, uint64(m.IndexInfo.Size()))
		n2, err := m.IndexInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.PrimaryColumnIds) > 0 {
		for _, num := range m.PrimaryColumnIds {
			dAtA[i] = 0x40
			i++
			i = encodeVarintDdl(dAtA, i, uint64(num))
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DDLResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DDLResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDdl(dAtA, i, uint64(m.Error.Size()))
		n3, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDdl(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.Values) > 0 {
		for _, b := range m.Values {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintDdl(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDdl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DDLRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDdl(uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovDdl(uint64(l))
		}
	}
	if len(m.Ranges) > 0 {
		for _, e := range m.Ranges {
			l = e.Size()
			n += 1 + l + sovDdl(uint64(l))
		}
	}
	if len(m.PrimaryPrefixColumnIds) > 0 {
		for _, e := range m.PrimaryPrefixColumnIds {
			n += 1 + sovDdl(uint64(e))
		}
	}
	if m.Unique != nil {
		n += 2
	}
	if m.TableInfo != nil {
		l = m.TableInfo.Size()
		n += 1 + l + sovDdl(uint64(l))
	}
	if m.IndexInfo != nil {
		l = m.IndexInfo.Size()
		n += 1 + l + sovDdl(uint64(l))
	}
	if len(m.PrimaryColumnIds) > 0 {
		for _, e := range m.PrimaryColumnIds {
			n += 1 + sovDdl(uint64(e))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DDLResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovDdl(uint64(l))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			l = len(b)
			n += 1 + l + sovDdl(uint64(l))
		}
	}
	if len(m.Values) > 0 {
		for _, b := range m.Values {
			l = len(b)
			n += 1 + l + sovDdl(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDdl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDdl(x uint64) (n int) {
	return sovDdl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DDLRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDdl
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
			return fmt.Errorf("proto: DDLRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DDLRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ranges = append(m.Ranges, KeyRange{})
			if err := m.Ranges[len(m.Ranges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDdl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PrimaryPrefixColumnIds = append(m.PrimaryPrefixColumnIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDdl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthDdl
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDdl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PrimaryPrefixColumnIds = append(m.PrimaryPrefixColumnIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryPrefixColumnIds", wireType)
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Unique = &b
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TableInfo == nil {
				m.TableInfo = &TableInfo{}
			}
			if err := m.TableInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IndexInfo == nil {
				m.IndexInfo = &IndexInfo{}
			}
			if err := m.IndexInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDdl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PrimaryColumnIds = append(m.PrimaryColumnIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDdl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthDdl
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDdl
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PrimaryColumnIds = append(m.PrimaryColumnIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryColumnIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDdl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDdl
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
func (m *DDLResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDdl
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
			return fmt.Errorf("proto: DDLResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DDLResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, make([]byte, postIndex-iNdEx))
			copy(m.Keys[len(m.Keys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDdl
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, make([]byte, postIndex-iNdEx))
			copy(m.Values[len(m.Values)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDdl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDdl
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
func skipDdl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDdl
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
					return 0, ErrIntOverflowDdl
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
					return 0, ErrIntOverflowDdl
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
				return 0, ErrInvalidLengthDdl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDdl
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
				next, err := skipDdl(dAtA[start:])
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
	ErrInvalidLengthDdl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDdl   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ddl.proto", fileDescriptorDdl) }

var fileDescriptorDdl = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x75, 0x9a, 0xa4, 0xe3, 0x08, 0xa2, 0x15, 0x54, 0xa6, 0x48, 0xa9, 0xe9, 0xc9,
	0x54, 0x95, 0x91, 0x72, 0xe3, 0x5a, 0xcc, 0x21, 0x22, 0x87, 0x6a, 0xd5, 0x23, 0x52, 0xe4, 0x78,
	0x37, 0x61, 0x85, 0xb3, 0xbb, 0xdd, 0xb5, 0x51, 0xfd, 0x26, 0x3c, 0x52, 0x8f, 0x3c, 0x01, 0x42,
	0xe1, 0xc2, 0x63, 0xa0, 0xfd, 0xe3, 0x2a, 0xa7, 0x9d, 0xf9, 0xe6, 0xa7, 0x9d, 0x6f, 0x3e, 0x38,
	0xa3, 0xb4, 0xce, 0x95, 0x96, 0x8d, 0xc4, 0xc3, 0x86, 0xab, 0xcd, 0xc5, 0xab, 0x9d, 0xdc, 0x49,
	0x27, 0x7c, 0xb0, 0x95, 0x9f, 0x5d, 0x4c, 0x4d, 0xf5, 0x8d, 0xed, 0xcb, 0xe7, 0x8e, 0xd5, 0xac,
	0x6a, 0x7c, 0x77, 0xf5, 0xef, 0x04, 0xa0, 0x28, 0x56, 0x84, 0x3d, 0xb4, 0xcc, 0x34, 0xf8, 0x12,
	0x26, 0x4d, 0xb9, 0xa9, 0xd9, 0x9a, 0xd3, 0x04, 0xa5, 0x28, 0x8b, 0x6e, 0x87, 0x4f, 0xbf, 0x2f,
	0x07, 0x64, 0xec, 0xd4, 0x25, 0xc5, 0xd7, 0x30, 0xae, 0x64, 0xdd, 0xee, 0x85, 0x49, 0x4e, 0xd2,
	0x28, 0x8b, 0x17, 0xb3, 0xdc, 0x6e, 0xce, 0x3f, 0x39, 0x71, 0x29, 0xb6, 0x92, 0xf4, 0x00, 0xbe,
	0x81, 0x91, 0x2e, 0xc5, 0x8e, 0x99, 0x24, 0x72, 0xe8, 0x0b, 0x8f, 0x7e, 0x61, 0x1d, 0xb1, 0x72,
	0xf8, 0x3a, 0x30, 0xf8, 0x23, 0xbc, 0x51, 0x9a, 0xef, 0x4b, 0xdd, 0xad, 0x95, 0x66, 0x5b, 0xfe,
	0xb8, 0xf6, 0xff, 0xac, 0x39, 0x35, 0xc9, 0x30, 0x8d, 0xb2, 0x88, 0x9c, 0x07, 0xe0, 0xce, 0xcd,
	0xc3, 0x4a, 0x6a, 0xf0, 0x39, 0x8c, 0x5a, 0xc1, 0x1f, 0x5a, 0x96, 0x9c, 0xa6, 0x28, 0x9b, 0x90,
	0xd0, 0xe1, 0x1c, 0x20, 0x5c, 0x23, 0xb6, 0x32, 0x19, 0xa5, 0x28, 0x8b, 0x17, 0x2f, 0xbd, 0x89,
	0x7b, 0x77, 0x8f, 0xb5, 0x7b, 0xd6, 0xf4, 0xa5, 0xe5, 0xb9, 0xa0, 0xec, 0xd1, 0xf3, 0xe3, 0x63,
	0x7e, 0x69, 0x75, 0xcf, 0xf3, 0xbe, 0xc4, 0x37, 0x80, 0x7b, 0xcb, 0x47, 0x5e, 0x27, 0xce, 0xeb,
	0x2c, 0x4c, 0x9e, 0x5d, 0x5e, 0x7d, 0x85, 0xd8, 0x25, 0x6d, 0x94, 0x14, 0x86, 0xe1, 0x77, 0x70,
	0xca, 0xb4, 0x96, 0xda, 0xe5, 0x1c, 0x2f, 0x62, 0xbf, 0xe7, 0xb3, 0x95, 0x88, 0x9f, 0x60, 0x0c,
	0xc3, 0xef, 0xac, 0xf3, 0x49, 0x4f, 0x89, 0xab, 0xed, 0xad, 0x3f, 0xca, 0xba, 0x0d, 0xa1, 0x4e,
	0x49, 0xe8, 0xae, 0xdf, 0xc2, 0xb8, 0x28, 0x56, 0xf7, 0x9d, 0x62, 0x78, 0x06, 0x53, 0xfb, 0x16,
	0xc5, 0xca, 0xb9, 0x9e, 0x0d, 0x6e, 0xdf, 0x3f, 0x1d, 0xe6, 0xe8, 0xd7, 0x61, 0x8e, 0xfe, 0x1c,
	0xe6, 0xe8, 0xe7, 0xdf, 0xf9, 0x00, 0x5e, 0x57, 0x72, 0x9f, 0x2b, 0x2e, 0x76, 0x55, 0xa9, 0xf2,
	0x86, 0xd3, 0x8d, 0x5b, 0x7f, 0x87, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xee, 0x2b, 0xcd, 0xf6,
	0x54, 0x02, 0x00, 0x00,
}
