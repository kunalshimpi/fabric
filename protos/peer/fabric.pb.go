// Code generated by protoc-gen-go.
// source: peer/fabric.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PeerEndpoint_Type int32

const (
	PeerEndpoint_UNDEFINED     PeerEndpoint_Type = 0
	PeerEndpoint_VALIDATOR     PeerEndpoint_Type = 1
	PeerEndpoint_NON_VALIDATOR PeerEndpoint_Type = 2
)

var PeerEndpoint_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "VALIDATOR",
	2: "NON_VALIDATOR",
}
var PeerEndpoint_Type_value = map[string]int32{
	"UNDEFINED":     0,
	"VALIDATOR":     1,
	"NON_VALIDATOR": 2,
}

func (x PeerEndpoint_Type) String() string {
	return proto.EnumName(PeerEndpoint_Type_name, int32(x))
}
func (PeerEndpoint_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2, 0} }

type PeerAddress struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *PeerAddress) Reset()                    { *m = PeerAddress{} }
func (m *PeerAddress) String() string            { return proto.CompactTextString(m) }
func (*PeerAddress) ProtoMessage()               {}
func (*PeerAddress) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type PeerID struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PeerID) Reset()                    { *m = PeerID{} }
func (m *PeerID) String() string            { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()               {}
func (*PeerID) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type PeerEndpoint struct {
	ID      *PeerID           `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Address string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type    PeerEndpoint_Type `protobuf:"varint,3,opt,name=type,enum=protos.PeerEndpoint_Type" json:"type,omitempty"`
	PkiID   []byte            `protobuf:"bytes,4,opt,name=pkiID,proto3" json:"pkiID,omitempty"`
}

func (m *PeerEndpoint) Reset()                    { *m = PeerEndpoint{} }
func (m *PeerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PeerEndpoint) ProtoMessage()               {}
func (*PeerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *PeerEndpoint) GetID() *PeerID {
	if m != nil {
		return m.ID
	}
	return nil
}

type PeersMessage struct {
	Peers []*PeerEndpoint `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *PeersMessage) Reset()                    { *m = PeersMessage{} }
func (m *PeersMessage) String() string            { return proto.CompactTextString(m) }
func (*PeersMessage) ProtoMessage()               {}
func (*PeersMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *PeersMessage) GetPeers() []*PeerEndpoint {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PeersAddresses struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *PeersAddresses) Reset()                    { *m = PeersAddresses{} }
func (m *PeersAddresses) String() string            { return proto.CompactTextString(m) }
func (*PeersAddresses) ProtoMessage()               {}
func (*PeersAddresses) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	Height            uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	CurrentBlockHash  []byte `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash []byte `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
}

func (m *BlockchainInfo) Reset()                    { *m = BlockchainInfo{} }
func (m *BlockchainInfo) String() string            { return proto.CompactTextString(m) }
func (*BlockchainInfo) ProtoMessage()               {}
func (*BlockchainInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func init() {
	proto.RegisterType((*PeerAddress)(nil), "protos.PeerAddress")
	proto.RegisterType((*PeerID)(nil), "protos.PeerID")
	proto.RegisterType((*PeerEndpoint)(nil), "protos.PeerEndpoint")
	proto.RegisterType((*PeersMessage)(nil), "protos.PeersMessage")
	proto.RegisterType((*PeersAddresses)(nil), "protos.PeersAddresses")
	proto.RegisterType((*BlockchainInfo)(nil), "protos.BlockchainInfo")
	proto.RegisterEnum("protos.PeerEndpoint_Type", PeerEndpoint_Type_name, PeerEndpoint_Type_value)
}

func init() { proto.RegisterFile("peer/fabric.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x76, 0x73, 0xb9, 0x93, 0xcc, 0x5d, 0xc3, 0xdd, 0x52, 0x24, 0x42, 0x91, 0x90, 0xa7, 0x58,
	0x35, 0x81, 0x13, 0x11, 0x7c, 0xbb, 0x92, 0x13, 0x03, 0x9a, 0xca, 0x52, 0x7d, 0xf0, 0x45, 0x72,
	0xc9, 0x34, 0x09, 0x6d, 0xb3, 0xcb, 0x6e, 0x4e, 0xb8, 0x57, 0x7f, 0x9c, 0xbf, 0x4b, 0x76, 0xb7,
	0xf1, 0x84, 0xf6, 0x29, 0xf3, 0x7d, 0xf3, 0x7d, 0x93, 0x99, 0xd9, 0x81, 0x95, 0x40, 0x94, 0xe9,
	0x75, 0xb9, 0x93, 0x5d, 0x95, 0x08, 0xc9, 0x07, 0x4e, 0x67, 0xe6, 0xa3, 0xa2, 0x77, 0x30, 0xff,
	0x8a, 0x28, 0x37, 0x75, 0x2d, 0x51, 0x29, 0x4a, 0xc1, 0x6d, 0xb9, 0x1a, 0x02, 0x12, 0x92, 0xd8,
	0x63, 0x26, 0xd6, 0x9c, 0xe0, 0x72, 0x08, 0x9c, 0x90, 0xc4, 0x53, 0x66, 0xe2, 0xe8, 0x0c, 0x66,
	0xda, 0x96, 0x67, 0x3a, 0xdb, 0x97, 0x77, 0x38, 0x3a, 0x74, 0x1c, 0xfd, 0x21, 0xb0, 0xd0, 0xe9,
	0x6d, 0x5f, 0x0b, 0xde, 0xf5, 0x03, 0x7d, 0x01, 0x4e, 0x9e, 0x19, 0xc9, 0x7c, 0xed, 0xdb, 0x0e,
	0x54, 0x62, 0x0b, 0x30, 0x27, 0xcf, 0x68, 0x00, 0x4f, 0x4b, 0xdb, 0x81, 0xf9, 0x8b, 0xc7, 0x46,
	0x48, 0xdf, 0x80, 0x3b, 0x1c, 0x04, 0x06, 0x93, 0x90, 0xc4, 0xfe, 0xfa, 0xf9, 0xff, 0xde, 0xb1,
	0x7a, 0x72, 0x75, 0x10, 0xc8, 0x8c, 0x8c, 0x9e, 0xc2, 0x54, 0xdc, 0x74, 0x79, 0x16, 0xb8, 0x21,
	0x89, 0x17, 0xcc, 0x82, 0xe8, 0x3d, 0xb8, 0x5a, 0x43, 0x4f, 0xc0, 0xfb, 0x56, 0x64, 0xdb, 0x8f,
	0x79, 0xb1, 0xcd, 0x96, 0x4f, 0x34, 0xfc, 0xbe, 0xf9, 0x9c, 0x67, 0x9b, 0xab, 0x4b, 0xb6, 0x24,
	0x74, 0x05, 0x27, 0xc5, 0x65, 0xf1, 0xf3, 0x48, 0x39, 0xd1, 0x07, 0x3b, 0x87, 0xfa, 0x82, 0x4a,
	0x95, 0x0d, 0xd2, 0x73, 0x98, 0xea, 0x55, 0xaa, 0x80, 0x84, 0x93, 0x78, 0xbe, 0x3e, 0x7d, 0xac,
	0x1d, 0x66, 0x25, 0x51, 0x02, 0xbe, 0xf1, 0xde, 0xaf, 0x16, 0x15, 0x3d, 0x03, 0xaf, 0x1c, 0x81,
	0xa9, 0xe0, 0xb1, 0x23, 0x11, 0xfd, 0x26, 0xe0, 0x5f, 0xdc, 0xf2, 0xea, 0xa6, 0x6a, 0xcb, 0xae,
	0xcf, 0xfb, 0x6b, 0x4e, 0x9f, 0xc1, 0xac, 0xc5, 0xae, 0x69, 0xed, 0x7b, 0xb8, 0xec, 0x1e, 0xd1,
	0x73, 0x58, 0x56, 0x7b, 0x29, 0xb1, 0x1f, 0x8c, 0xe1, 0x53, 0xa9, 0x5a, 0xb3, 0xb7, 0x05, 0x7b,
	0xc0, 0xd3, 0xd7, 0xb0, 0x12, 0x12, 0x7f, 0x75, 0x7c, 0xaf, 0x8e, 0xe2, 0x89, 0x11, 0x3f, 0x4c,
	0x5c, 0xbc, 0xfa, 0xf1, 0xb2, 0xe9, 0x86, 0x76, 0xbf, 0x4b, 0x2a, 0x7e, 0x97, 0xb6, 0x07, 0x81,
	0xf2, 0x16, 0xeb, 0xe6, 0xdf, 0xf5, 0xa4, 0x76, 0xe0, 0x54, 0x8f, 0xb8, 0xb3, 0x37, 0xf4, 0xf6,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x30, 0x01, 0x3f, 0x5f, 0x02, 0x00, 0x00,
}
