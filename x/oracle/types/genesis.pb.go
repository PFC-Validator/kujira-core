// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params                        Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeederDelegations             []FeederDelegation             `protobuf:"bytes,2,rep,name=feeder_delegations,json=feederDelegations,proto3" json:"feeder_delegations"`
	ExchangeRates                 ExchangeRateTuples             `protobuf:"bytes,3,rep,name=exchange_rates,json=exchangeRates,proto3,castrepeated=ExchangeRateTuples" json:"exchange_rates"`
	MissCounters                  []MissCounter                  `protobuf:"bytes,4,rep,name=miss_counters,json=missCounters,proto3" json:"miss_counters"`
	AggregateExchangeRatePrevotes []AggregateExchangeRatePrevote `protobuf:"bytes,5,rep,name=aggregate_exchange_rate_prevotes,json=aggregateExchangeRatePrevotes,proto3" json:"aggregate_exchange_rate_prevotes"`
	AggregateExchangeRateVotes    []AggregateExchangeRateVote    `protobuf:"bytes,6,rep,name=aggregate_exchange_rate_votes,json=aggregateExchangeRateVotes,proto3" json:"aggregate_exchange_rate_votes"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce0b3a2b4a184fc3, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeederDelegations() []FeederDelegation {
	if m != nil {
		return m.FeederDelegations
	}
	return nil
}

func (m *GenesisState) GetExchangeRates() ExchangeRateTuples {
	if m != nil {
		return m.ExchangeRates
	}
	return nil
}

func (m *GenesisState) GetMissCounters() []MissCounter {
	if m != nil {
		return m.MissCounters
	}
	return nil
}

func (m *GenesisState) GetAggregateExchangeRatePrevotes() []AggregateExchangeRatePrevote {
	if m != nil {
		return m.AggregateExchangeRatePrevotes
	}
	return nil
}

func (m *GenesisState) GetAggregateExchangeRateVotes() []AggregateExchangeRateVote {
	if m != nil {
		return m.AggregateExchangeRateVotes
	}
	return nil
}

// FeederDelegation is the address for where oracle feeder authority are
// delegated to. By default this struct is only used at genesis to feed in
// default feeder addresses.
type FeederDelegation struct {
	FeederAddress    string `protobuf:"bytes,1,opt,name=feeder_address,json=feederAddress,proto3" json:"feeder_address,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *FeederDelegation) Reset()         { *m = FeederDelegation{} }
func (m *FeederDelegation) String() string { return proto.CompactTextString(m) }
func (*FeederDelegation) ProtoMessage()    {}
func (*FeederDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce0b3a2b4a184fc3, []int{1}
}
func (m *FeederDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeederDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeederDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeederDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeederDelegation.Merge(m, src)
}
func (m *FeederDelegation) XXX_Size() int {
	return m.Size()
}
func (m *FeederDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeederDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_FeederDelegation proto.InternalMessageInfo

func (m *FeederDelegation) GetFeederAddress() string {
	if m != nil {
		return m.FeederAddress
	}
	return ""
}

func (m *FeederDelegation) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

// MissCounter defines an miss counter and validator address pair used in
// oracle module's genesis state
type MissCounter struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	MissCounter      uint64 `protobuf:"varint,2,opt,name=miss_counter,json=missCounter,proto3" json:"miss_counter,omitempty"`
}

func (m *MissCounter) Reset()         { *m = MissCounter{} }
func (m *MissCounter) String() string { return proto.CompactTextString(m) }
func (*MissCounter) ProtoMessage()    {}
func (*MissCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce0b3a2b4a184fc3, []int{2}
}
func (m *MissCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MissCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MissCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissCounter.Merge(m, src)
}
func (m *MissCounter) XXX_Size() int {
	return m.Size()
}
func (m *MissCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MissCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MissCounter proto.InternalMessageInfo

func (m *MissCounter) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MissCounter) GetMissCounter() uint64 {
	if m != nil {
		return m.MissCounter
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kujira.oracle.GenesisState")
	proto.RegisterType((*FeederDelegation)(nil), "kujira.oracle.FeederDelegation")
	proto.RegisterType((*MissCounter)(nil), "kujira.oracle.MissCounter")
}

func init() { proto.RegisterFile("oracle/genesis.proto", fileDescriptor_ce0b3a2b4a184fc3) }

var fileDescriptor_ce0b3a2b4a184fc3 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x86, 0xb3, 0x6d, 0x0c, 0x38, 0x49, 0x4a, 0x3b, 0xb6, 0x10, 0x16, 0xba, 0x89, 0x01, 0x21,
	0x50, 0xc8, 0xd2, 0xf6, 0x09, 0x1a, 0xad, 0x5e, 0x09, 0x65, 0x2d, 0x5e, 0x08, 0xb2, 0x9c, 0xec,
	0x9e, 0xac, 0xab, 0xc9, 0xce, 0x3a, 0x67, 0x12, 0xaa, 0x4f, 0xe1, 0x73, 0x78, 0xed, 0x43, 0xf4,
	0xb2, 0x97, 0x5e, 0xa9, 0x24, 0x2f, 0x22, 0x3b, 0x33, 0x36, 0xdb, 0xb5, 0x11, 0xaf, 0x76, 0x39,
	0xff, 0x7f, 0xfe, 0xef, 0x30, 0x67, 0x86, 0xed, 0x0b, 0x09, 0xd1, 0x14, 0xfd, 0x04, 0x33, 0xa4,
	0x94, 0x86, 0xb9, 0x14, 0x4a, 0xf0, 0xf6, 0x87, 0xf9, 0xfb, 0x54, 0xc2, 0xd0, 0x88, 0xee, 0x7e,
	0x22, 0x12, 0xa1, 0x15, 0xbf, 0xf8, 0x33, 0x26, 0xf7, 0x91, 0x6d, 0x35, 0x1f, 0x5b, 0xf4, 0x22,
	0x41, 0x33, 0x41, 0xfe, 0x18, 0x08, 0xfd, 0xc5, 0xf1, 0x18, 0x15, 0x1c, 0xfb, 0x91, 0x48, 0x33,
	0xa3, 0xf7, 0xbf, 0xd5, 0x59, 0xeb, 0x85, 0x61, 0xbd, 0x52, 0xa0, 0x90, 0x9f, 0xb2, 0x46, 0x0e,
	0x12, 0x66, 0xd4, 0x71, 0x7a, 0xce, 0xa0, 0x79, 0x72, 0x30, 0xbc, 0xc3, 0x1e, 0x5e, 0x68, 0x71,
	0x54, 0xbf, 0xfe, 0xd1, 0xad, 0x05, 0xd6, 0xca, 0x2f, 0x19, 0x9f, 0x20, 0xc6, 0x28, 0xc3, 0x18,
	0xa7, 0x98, 0x80, 0x4a, 0x45, 0x46, 0x9d, 0xad, 0xde, 0xf6, 0xa0, 0x79, 0xd2, 0xad, 0x04, 0x3c,
	0xd7, 0xc6, 0x67, 0xb7, 0x3e, 0x1b, 0xb5, 0x37, 0xa9, 0xd4, 0x89, 0x47, 0x6c, 0x07, 0xaf, 0xa2,
	0x77, 0x90, 0x25, 0x18, 0x4a, 0x50, 0x48, 0x9d, 0x6d, 0x9d, 0xd8, 0xab, 0x24, 0x9e, 0x5b, 0x53,
	0x00, 0x0a, 0x2f, 0xe7, 0xf9, 0x14, 0x47, 0x6e, 0x11, 0xf9, 0xf5, 0x67, 0x97, 0xff, 0x25, 0x51,
	0xd0, 0xc6, 0x52, 0x8d, 0xf8, 0x39, 0x6b, 0xcf, 0x52, 0xa2, 0x30, 0x12, 0xf3, 0x4c, 0xa1, 0xa4,
	0x4e, 0x5d, 0x33, 0xdc, 0x0a, 0xe3, 0x65, 0x4a, 0xf4, 0xd4, 0x58, 0xec, 0xc0, 0xad, 0xd9, 0xba,
	0x44, 0xfc, 0x33, 0xeb, 0x41, 0x92, 0xc8, 0x62, 0x76, 0x0c, 0xef, 0x4c, 0x1d, 0xe6, 0x12, 0x17,
	0xa2, 0x98, 0xfe, 0x81, 0x4e, 0x3e, 0xaa, 0x24, 0x9f, 0xfd, 0x69, 0x2b, 0xcf, 0x7a, 0x61, 0x7a,
	0x2c, 0xea, 0x10, 0xfe, 0xe1, 0x21, 0xfe, 0x91, 0x1d, 0x6e, 0x62, 0x1b, 0x70, 0x43, 0x83, 0x07,
	0xff, 0x03, 0x7e, 0xbd, 0xa6, 0xba, 0xb0, 0xc9, 0x40, 0xfd, 0x09, 0xdb, 0xad, 0xee, 0x91, 0x3f,
	0x61, 0x3b, 0xf6, 0x12, 0x40, 0x1c, 0x4b, 0x24, 0x73, 0x83, 0x1e, 0x06, 0x6d, 0x53, 0x3d, 0x33,
	0x45, 0x7e, 0xc4, 0xf6, 0x16, 0x30, 0x4d, 0x63, 0x50, 0x62, 0xed, 0xdc, 0xd2, 0xce, 0xdd, 0x5b,
	0xc1, 0x9a, 0xfb, 0x6f, 0x59, 0xb3, 0x74, 0xf2, 0xf7, 0xf7, 0x3a, 0xf7, 0xf7, 0xf2, 0xc7, 0xac,
	0x55, 0xde, 0xac, 0x66, 0xd4, 0x83, 0x66, 0x69, 0x6d, 0x23, 0xff, 0x7a, 0xe9, 0x39, 0x37, 0x4b,
	0xcf, 0xf9, 0xb5, 0xf4, 0x9c, 0x2f, 0x2b, 0xaf, 0x76, 0xb3, 0xf2, 0x6a, 0xdf, 0x57, 0x5e, 0xed,
	0xcd, 0x81, 0x39, 0x2b, 0xff, 0xca, 0x3e, 0x27, 0x5f, 0x7d, 0xca, 0x91, 0xc6, 0x0d, 0xfd, 0x6a,
	0x4e, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x72, 0x37, 0x9b, 0xa7, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AggregateExchangeRateVotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRateVotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRateVotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRatePrevotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRatePrevotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MissCounters) > 0 {
		for iNdEx := len(m.MissCounters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissCounters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ExchangeRates) > 0 {
		for iNdEx := len(m.ExchangeRates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExchangeRates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeederDelegations) > 0 {
		for iNdEx := len(m.FeederDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeederDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeederDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeederDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeederDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FeederAddress) > 0 {
		i -= len(m.FeederAddress)
		copy(dAtA[i:], m.FeederAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.FeederAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MissCounter))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeederDelegations) > 0 {
		for _, e := range m.FeederDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExchangeRates) > 0 {
		for _, e := range m.ExchangeRates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissCounters) > 0 {
		for _, e := range m.MissCounters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for _, e := range m.AggregateExchangeRatePrevotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRateVotes) > 0 {
		for _, e := range m.AggregateExchangeRateVotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *FeederDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeederAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *MissCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MissCounter != 0 {
		n += 1 + sovGenesis(uint64(m.MissCounter))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederDelegations = append(m.FeederDelegations, FeederDelegation{})
			if err := m.FeederDelegations[len(m.FeederDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeRates = append(m.ExchangeRates, ExchangeRateTuple{})
			if err := m.ExchangeRates[len(m.ExchangeRates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissCounters = append(m.MissCounters, MissCounter{})
			if err := m.MissCounters[len(m.MissCounters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRatePrevotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateExchangeRatePrevotes = append(m.AggregateExchangeRatePrevotes, AggregateExchangeRatePrevote{})
			if err := m.AggregateExchangeRatePrevotes[len(m.AggregateExchangeRatePrevotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRateVotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateExchangeRateVotes = append(m.AggregateExchangeRateVotes, AggregateExchangeRateVote{})
			if err := m.AggregateExchangeRateVotes[len(m.AggregateExchangeRateVotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *FeederDelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: FeederDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeederDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *MissCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: MissCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounter", wireType)
			}
			m.MissCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
