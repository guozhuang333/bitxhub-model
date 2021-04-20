// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_meshplus_bitxhub_kit_types "github.com/meshplus/bitxhub-kit/types"
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

type Block struct {
	BlockHeader  *BlockHeader                                `protobuf:"bytes,1,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	Transactions *Transactions                               `protobuf:"bytes,2,opt,name=transactions,proto3,customtype=Transactions" json:"transactions,omitempty"`
	BlockHash    *github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"block_hash,omitempty"`
	Signature    []byte                                      `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Extra        []byte                                      `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetBlockHeader() *BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

func (m *Block) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type BlockHeader struct {
	Number      uint64                                       `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	StateRoot   *github_com_meshplus_bitxhub_kit_types.Hash  `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"state_root,omitempty"`
	TxRoot      *github_com_meshplus_bitxhub_kit_types.Hash  `protobuf:"bytes,3,opt,name=tx_root,json=txRoot,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"tx_root,omitempty"`
	ReceiptRoot *github_com_meshplus_bitxhub_kit_types.Hash  `protobuf:"bytes,4,opt,name=receipt_root,json=receiptRoot,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"receipt_root,omitempty"`
	ParentHash  *github_com_meshplus_bitxhub_kit_types.Hash  `protobuf:"bytes,5,opt,name=parent_hash,json=parentHash,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"parent_hash,omitempty"`
	Timestamp   int64                                        `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Version     []byte                                       `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Bloom       *github_com_meshplus_bitxhub_kit_types.Bloom `protobuf:"bytes,8,opt,name=Bloom,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Bloom" json:"Bloom,omitempty"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x8a, 0x13, 0x31,
	0x14, 0xc6, 0x3b, 0xdb, 0x7f, 0x6e, 0x66, 0x40, 0x09, 0x2a, 0x61, 0x91, 0x69, 0xd9, 0xab, 0xa2,
	0xec, 0x0c, 0xac, 0x3e, 0x41, 0x41, 0xdc, 0x1b, 0x11, 0x83, 0xf7, 0x25, 0xa9, 0xb1, 0x09, 0xbb,
	0x33, 0x09, 0xc9, 0x19, 0x19, 0xdf, 0xc2, 0xc7, 0xf2, 0x72, 0x2f, 0xa5, 0x17, 0x65, 0x69, 0x5f,
	0xc1, 0x07, 0x90, 0x9e, 0xcc, 0xb2, 0xf5, 0xce, 0xed, 0x5d, 0xbe, 0x2f, 0xe7, 0xfc, 0xf2, 0x1d,
	0x4e, 0x48, 0x2a, 0x6f, 0xec, 0xf2, 0xba, 0x70, 0xde, 0x82, 0xa5, 0x27, 0x4e, 0x9e, 0xbd, 0x90,
	0xad, 0x5e, 0x80, 0x17, 0x75, 0x10, 0x4b, 0x30, 0xb6, 0x8e, 0x57, 0x67, 0x17, 0x2b, 0x03, 0xba,
	0x91, 0xc5, 0xd2, 0x56, 0xe5, 0xca, 0xae, 0x6c, 0x89, 0xb6, 0x6c, 0xbe, 0xa1, 0x42, 0x81, 0xa7,
	0x58, 0x7e, 0xfe, 0x27, 0x21, 0xc3, 0xf9, 0x9e, 0x4c, 0x2f, 0x49, 0x86, 0x4f, 0x2c, 0xb4, 0x12,
	0x5f, 0x95, 0x67, 0xc9, 0x34, 0x99, 0xa5, 0x97, 0x4f, 0x0b, 0x27, 0x0b, 0x2c, 0xb8, 0x42, 0x9b,
	0xc7, 0x1c, 0x51, 0xd0, 0x77, 0x24, 0x3b, 0x48, 0x10, 0xd8, 0xc9, 0x34, 0x99, 0x65, 0xf3, 0x67,
	0xeb, 0xcd, 0x24, 0xfb, 0x72, 0xe0, 0xf3, 0x7f, 0xaa, 0xe8, 0x47, 0x42, 0xba, 0x97, 0x44, 0xd0,
	0xac, 0x8f, 0x3d, 0xc5, 0x7a, 0x33, 0x79, 0x7d, 0x10, 0xbd, 0x52, 0x41, 0xbb, 0x9b, 0x26, 0x94,
	0xd2, 0x40, 0xab, 0x1b, 0x79, 0x71, 0x6d, 0xa0, 0x84, 0x1f, 0x4e, 0x85, 0xe2, 0x4a, 0x04, 0xcd,
	0x4f, 0x63, 0x0c, 0x11, 0x34, 0x7d, 0x45, 0x4e, 0x83, 0x59, 0xd5, 0x02, 0x1a, 0xaf, 0xd8, 0x60,
	0x4f, 0xe3, 0x0f, 0x06, 0x7d, 0x4e, 0x86, 0xaa, 0x05, 0x2f, 0xd8, 0x10, 0x6f, 0xa2, 0x38, 0xbf,
	0xeb, 0x93, 0xf4, 0x60, 0x2a, 0xfa, 0x92, 0x8c, 0xea, 0xa6, 0x92, 0xdd, 0xd8, 0x03, 0xde, 0xa9,
	0x7d, 0xd4, 0x00, 0x02, 0xd4, 0xc2, 0x5b, 0x0b, 0xdd, 0x78, 0x8f, 0x8e, 0x8a, 0x04, 0x6e, 0x2d,
	0xd0, 0x0f, 0x64, 0x0c, 0x6d, 0x64, 0x1d, 0x37, 0xf6, 0x08, 0x5a, 0x04, 0x7d, 0x26, 0x99, 0x57,
	0x4b, 0x65, 0x1c, 0x44, 0xda, 0xe0, 0x28, 0x5a, 0xda, 0x31, 0x10, 0xf9, 0x89, 0xa4, 0x4e, 0x78,
	0x55, 0x43, 0x5c, 0xcb, 0xf0, 0x28, 0x22, 0x89, 0x88, 0xfb, 0xbd, 0x80, 0xa9, 0x54, 0x00, 0x51,
	0x39, 0x36, 0x9a, 0x26, 0xb3, 0x3e, 0x7f, 0x30, 0x28, 0x23, 0xe3, 0xef, 0xca, 0x07, 0x63, 0x6b,
	0x36, 0xc6, 0xcd, 0xdc, 0x4b, 0xfa, 0x1e, 0x7f, 0xa4, 0xad, 0xd8, 0x13, 0x8c, 0x50, 0xae, 0x37,
	0x93, 0x37, 0xff, 0x17, 0x01, 0xdb, 0x78, 0xec, 0x9e, 0xb3, 0x5f, 0xdb, 0x3c, 0xb9, 0xdd, 0xe6,
	0xc9, 0xdd, 0x36, 0x4f, 0x7e, 0xee, 0xf2, 0xde, 0xed, 0x2e, 0xef, 0xfd, 0xde, 0xe5, 0x3d, 0x39,
	0xc2, 0xaf, 0xff, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x9a, 0x3c, 0xd7, 0x53, 0x03,
	0x00, 0x00,
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHash != nil {
		{
			size := m.BlockHash.Size()
			i -= size
			if _, err := m.BlockHash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Transactions != nil {
		{
			size := m.Transactions.Size()
			i -= size
			if _, err := m.Transactions.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeader != nil {
		{
			size, err := m.BlockHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bloom != nil {
		{
			size := m.Bloom.Size()
			i -= size
			if _, err := m.Bloom.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.ParentHash != nil {
		{
			size := m.ParentHash.Size()
			i -= size
			if _, err := m.ParentHash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ReceiptRoot != nil {
		{
			size := m.ReceiptRoot.Size()
			i -= size
			if _, err := m.ReceiptRoot.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TxRoot != nil {
		{
			size := m.TxRoot.Size()
			i -= size
			if _, err := m.TxRoot.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.StateRoot != nil {
		{
			size := m.StateRoot.Size()
			i -= size
			if _, err := m.StateRoot.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Number != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeader != nil {
		l = m.BlockHeader.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Transactions != nil {
		l = m.Transactions.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.BlockHash != nil {
		l = m.BlockHash.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *BlockHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovBlock(uint64(m.Number))
	}
	if m.StateRoot != nil {
		l = m.StateRoot.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.TxRoot != nil {
		l = m.TxRoot.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.ReceiptRoot != nil {
		l = m.ReceiptRoot.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.ParentHash != nil {
		l = m.ParentHash.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Bloom != nil {
		l = m.Bloom.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockHeader == nil {
				m.BlockHeader = &BlockHeader{}
			}
			if err := m.BlockHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v Transactions
			m.Transactions = &v
			if err := m.Transactions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Hash
			m.BlockHash = &v
			if err := m.BlockHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Hash
			m.StateRoot = &v
			if err := m.StateRoot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Hash
			m.TxRoot = &v
			if err := m.TxRoot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Hash
			m.ReceiptRoot = &v
			if err := m.ReceiptRoot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Hash
			m.ParentHash = &v
			if err := m.ParentHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = append(m.Version[:0], dAtA[iNdEx:postIndex]...)
			if m.Version == nil {
				m.Version = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_meshplus_bitxhub_kit_types.Bloom
			m.Bloom = &v
			if err := m.Bloom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
