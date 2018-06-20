// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha1/l4_fault.proto

package v1alpha1 // import "istio.io/api/routing/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// (-- Faults can be injected into the connections from downstream by the
// Envoy, for testing the failure recovery capabilities of downstream
// services.  Faults include aborting the connection from downstream
// service, delaying proxying of connection to the destination
// service, and throttling the bandwidth of the connection (either
// end). Bandwidth throttling for failure testing should not be confused
// with the rate limiting policy enforcement provided by the Mixer
// component. L4 fault injection is not supported at the moment. --)
type L4FaultInjection struct {
	// Unlike Http services, we have very little context for raw Tcp|Udp
	// connections. We could throttle bandwidth of the connections (slow down
	// the connection) and/or abruptly reset (terminate) the Tcp connection
	// after it has been established.
	// We first throttle (if set) and then terminate the connection.
	Throttle             *L4FaultInjection_Throttle  `protobuf:"bytes,1,opt,name=throttle,proto3" json:"throttle,omitempty"`
	Terminate            *L4FaultInjection_Terminate `protobuf:"bytes,2,opt,name=terminate,proto3" json:"terminate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *L4FaultInjection) Reset()         { *m = L4FaultInjection{} }
func (m *L4FaultInjection) String() string { return proto.CompactTextString(m) }
func (*L4FaultInjection) ProtoMessage()    {}
func (*L4FaultInjection) Descriptor() ([]byte, []int) {
	return fileDescriptor_l4_fault_a9fe5b227cbbe295, []int{0}
}
func (m *L4FaultInjection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L4FaultInjection.Unmarshal(m, b)
}
func (m *L4FaultInjection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L4FaultInjection.Marshal(b, m, deterministic)
}
func (dst *L4FaultInjection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L4FaultInjection.Merge(dst, src)
}
func (m *L4FaultInjection) XXX_Size() int {
	return xxx_messageInfo_L4FaultInjection.Size(m)
}
func (m *L4FaultInjection) XXX_DiscardUnknown() {
	xxx_messageInfo_L4FaultInjection.DiscardUnknown(m)
}

var xxx_messageInfo_L4FaultInjection proto.InternalMessageInfo

func (m *L4FaultInjection) GetThrottle() *L4FaultInjection_Throttle {
	if m != nil {
		return m.Throttle
	}
	return nil
}

func (m *L4FaultInjection) GetTerminate() *L4FaultInjection_Terminate {
	if m != nil {
		return m.Terminate
	}
	return nil
}

// Bandwidth throttling for Tcp and Udp connections
type L4FaultInjection_Throttle struct {
	// percentage of connections to throttle.
	Percent float32 `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
	// bandwidth limit in "bits" per second between downstream and Envoy
	DownstreamLimitBps int64 `protobuf:"varint,2,opt,name=downstream_limit_bps,json=downstreamLimitBps,proto3" json:"downstream_limit_bps,omitempty"`
	// bandwidth limits in "bits" per second between Envoy and upstream
	UpstreamLimitBps int64 `protobuf:"varint,3,opt,name=upstream_limit_bps,json=upstreamLimitBps,proto3" json:"upstream_limit_bps,omitempty"`
	// Types that are valid to be assigned to ThrottleAfter:
	//	*L4FaultInjection_Throttle_ThrottleAfterPeriod
	//	*L4FaultInjection_Throttle_ThrottleAfterBytes
	ThrottleAfter isL4FaultInjection_Throttle_ThrottleAfter `protobuf_oneof:"throttle_after"`
	// Stop throttling after the given duration. If not set, the connection
	// will be throttled for its lifetime.
	ThrottleForPeriod    *duration.Duration `protobuf:"bytes,6,opt,name=throttle_for_period,json=throttleForPeriod,proto3" json:"throttle_for_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L4FaultInjection_Throttle) Reset()         { *m = L4FaultInjection_Throttle{} }
func (m *L4FaultInjection_Throttle) String() string { return proto.CompactTextString(m) }
func (*L4FaultInjection_Throttle) ProtoMessage()    {}
func (*L4FaultInjection_Throttle) Descriptor() ([]byte, []int) {
	return fileDescriptor_l4_fault_a9fe5b227cbbe295, []int{0, 0}
}
func (m *L4FaultInjection_Throttle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L4FaultInjection_Throttle.Unmarshal(m, b)
}
func (m *L4FaultInjection_Throttle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L4FaultInjection_Throttle.Marshal(b, m, deterministic)
}
func (dst *L4FaultInjection_Throttle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L4FaultInjection_Throttle.Merge(dst, src)
}
func (m *L4FaultInjection_Throttle) XXX_Size() int {
	return xxx_messageInfo_L4FaultInjection_Throttle.Size(m)
}
func (m *L4FaultInjection_Throttle) XXX_DiscardUnknown() {
	xxx_messageInfo_L4FaultInjection_Throttle.DiscardUnknown(m)
}

var xxx_messageInfo_L4FaultInjection_Throttle proto.InternalMessageInfo

type isL4FaultInjection_Throttle_ThrottleAfter interface {
	isL4FaultInjection_Throttle_ThrottleAfter()
}

type L4FaultInjection_Throttle_ThrottleAfterPeriod struct {
	ThrottleAfterPeriod *duration.Duration `protobuf:"bytes,4,opt,name=throttle_after_period,json=throttleAfterPeriod,proto3,oneof"`
}
type L4FaultInjection_Throttle_ThrottleAfterBytes struct {
	ThrottleAfterBytes float64 `protobuf:"fixed64,5,opt,name=throttle_after_bytes,json=throttleAfterBytes,proto3,oneof"`
}

func (*L4FaultInjection_Throttle_ThrottleAfterPeriod) isL4FaultInjection_Throttle_ThrottleAfter() {}
func (*L4FaultInjection_Throttle_ThrottleAfterBytes) isL4FaultInjection_Throttle_ThrottleAfter()  {}

func (m *L4FaultInjection_Throttle) GetThrottleAfter() isL4FaultInjection_Throttle_ThrottleAfter {
	if m != nil {
		return m.ThrottleAfter
	}
	return nil
}

func (m *L4FaultInjection_Throttle) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *L4FaultInjection_Throttle) GetDownstreamLimitBps() int64 {
	if m != nil {
		return m.DownstreamLimitBps
	}
	return 0
}

func (m *L4FaultInjection_Throttle) GetUpstreamLimitBps() int64 {
	if m != nil {
		return m.UpstreamLimitBps
	}
	return 0
}

func (m *L4FaultInjection_Throttle) GetThrottleAfterPeriod() *duration.Duration {
	if x, ok := m.GetThrottleAfter().(*L4FaultInjection_Throttle_ThrottleAfterPeriod); ok {
		return x.ThrottleAfterPeriod
	}
	return nil
}

func (m *L4FaultInjection_Throttle) GetThrottleAfterBytes() float64 {
	if x, ok := m.GetThrottleAfter().(*L4FaultInjection_Throttle_ThrottleAfterBytes); ok {
		return x.ThrottleAfterBytes
	}
	return 0
}

func (m *L4FaultInjection_Throttle) GetThrottleForPeriod() *duration.Duration {
	if m != nil {
		return m.ThrottleForPeriod
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*L4FaultInjection_Throttle) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _L4FaultInjection_Throttle_OneofMarshaler, _L4FaultInjection_Throttle_OneofUnmarshaler, _L4FaultInjection_Throttle_OneofSizer, []interface{}{
		(*L4FaultInjection_Throttle_ThrottleAfterPeriod)(nil),
		(*L4FaultInjection_Throttle_ThrottleAfterBytes)(nil),
	}
}

func _L4FaultInjection_Throttle_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*L4FaultInjection_Throttle)
	// throttle_after
	switch x := m.ThrottleAfter.(type) {
	case *L4FaultInjection_Throttle_ThrottleAfterPeriod:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ThrottleAfterPeriod); err != nil {
			return err
		}
	case *L4FaultInjection_Throttle_ThrottleAfterBytes:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.ThrottleAfterBytes))
	case nil:
	default:
		return fmt.Errorf("L4FaultInjection_Throttle.ThrottleAfter has unexpected type %T", x)
	}
	return nil
}

func _L4FaultInjection_Throttle_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*L4FaultInjection_Throttle)
	switch tag {
	case 4: // throttle_after.throttle_after_period
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(duration.Duration)
		err := b.DecodeMessage(msg)
		m.ThrottleAfter = &L4FaultInjection_Throttle_ThrottleAfterPeriod{msg}
		return true, err
	case 5: // throttle_after.throttle_after_bytes
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ThrottleAfter = &L4FaultInjection_Throttle_ThrottleAfterBytes{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _L4FaultInjection_Throttle_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*L4FaultInjection_Throttle)
	// throttle_after
	switch x := m.ThrottleAfter.(type) {
	case *L4FaultInjection_Throttle_ThrottleAfterPeriod:
		s := proto.Size(x.ThrottleAfterPeriod)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *L4FaultInjection_Throttle_ThrottleAfterBytes:
		n += 1 // tag and wire
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Abruptly reset (terminate) the Tcp connection after it has been
// established, emulating remote server crash or link failure.
type L4FaultInjection_Terminate struct {
	// percentage of established Tcp connections to be terminated/reset
	Percent float32 `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
	// TODO: see if it makes sense to create a generic Duration type to
	// express time interval related configs.
	TerminateAfterPeriod *duration.Duration `protobuf:"bytes,2,opt,name=terminate_after_period,json=terminateAfterPeriod,proto3" json:"terminate_after_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L4FaultInjection_Terminate) Reset()         { *m = L4FaultInjection_Terminate{} }
func (m *L4FaultInjection_Terminate) String() string { return proto.CompactTextString(m) }
func (*L4FaultInjection_Terminate) ProtoMessage()    {}
func (*L4FaultInjection_Terminate) Descriptor() ([]byte, []int) {
	return fileDescriptor_l4_fault_a9fe5b227cbbe295, []int{0, 1}
}
func (m *L4FaultInjection_Terminate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L4FaultInjection_Terminate.Unmarshal(m, b)
}
func (m *L4FaultInjection_Terminate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L4FaultInjection_Terminate.Marshal(b, m, deterministic)
}
func (dst *L4FaultInjection_Terminate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L4FaultInjection_Terminate.Merge(dst, src)
}
func (m *L4FaultInjection_Terminate) XXX_Size() int {
	return xxx_messageInfo_L4FaultInjection_Terminate.Size(m)
}
func (m *L4FaultInjection_Terminate) XXX_DiscardUnknown() {
	xxx_messageInfo_L4FaultInjection_Terminate.DiscardUnknown(m)
}

var xxx_messageInfo_L4FaultInjection_Terminate proto.InternalMessageInfo

func (m *L4FaultInjection_Terminate) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *L4FaultInjection_Terminate) GetTerminateAfterPeriod() *duration.Duration {
	if m != nil {
		return m.TerminateAfterPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*L4FaultInjection)(nil), "istio.routing.v1alpha1.L4FaultInjection")
	proto.RegisterType((*L4FaultInjection_Throttle)(nil), "istio.routing.v1alpha1.L4FaultInjection.Throttle")
	proto.RegisterType((*L4FaultInjection_Terminate)(nil), "istio.routing.v1alpha1.L4FaultInjection.Terminate")
}

func init() {
	proto.RegisterFile("routing/v1alpha1/l4_fault.proto", fileDescriptor_l4_fault_a9fe5b227cbbe295)
}

var fileDescriptor_l4_fault_a9fe5b227cbbe295 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x4e, 0xc2, 0x30,
	0x18, 0x85, 0xd9, 0x10, 0x84, 0x9a, 0x18, 0xac, 0x48, 0x26, 0x89, 0x42, 0xbc, 0xe2, 0xc2, 0x74,
	0x82, 0xbc, 0x80, 0x8b, 0x21, 0x90, 0x60, 0x20, 0x8b, 0x57, 0xde, 0x2c, 0x1d, 0x74, 0x50, 0x33,
	0xd6, 0xa6, 0xeb, 0x30, 0xbe, 0x87, 0xcf, 0xe8, 0x73, 0x98, 0x95, 0x75, 0x08, 0x1a, 0xf1, 0xb2,
	0xf9, 0xcf, 0x77, 0xfe, 0xd3, 0xd3, 0x82, 0x96, 0x60, 0x89, 0xa4, 0xd1, 0xc2, 0x5e, 0x77, 0x71,
	0xc8, 0x97, 0xb8, 0x6b, 0x87, 0x7d, 0x2f, 0xc0, 0x49, 0x28, 0x11, 0x17, 0x4c, 0x32, 0xd8, 0xa0,
	0xb1, 0xa4, 0x0c, 0x65, 0x32, 0xa4, 0x65, 0xcd, 0xeb, 0x05, 0x63, 0x8b, 0x90, 0xd8, 0x4a, 0xe5,
	0x27, 0x81, 0x3d, 0x4f, 0x04, 0x96, 0x94, 0x45, 0x1b, 0xee, 0xe6, 0xa3, 0x04, 0x6a, 0xe3, 0xfe,
	0x20, 0x75, 0x1a, 0x45, 0xaf, 0x64, 0x96, 0x8e, 0xe0, 0x13, 0xa8, 0xc8, 0xa5, 0x60, 0x52, 0x86,
	0xc4, 0x32, 0xda, 0x46, 0xe7, 0xa4, 0xd7, 0x45, 0xbf, 0xfb, 0xa3, 0x7d, 0x16, 0x3d, 0x67, 0xa0,
	0x9b, 0x5b, 0xc0, 0x29, 0xa8, 0x4a, 0x22, 0x56, 0x34, 0xc2, 0x92, 0x58, 0xa6, 0xf2, 0xeb, 0xfd,
	0xdf, 0x4f, 0x93, 0xee, 0xd6, 0xa4, 0xf9, 0x69, 0x82, 0x8a, 0x5e, 0x04, 0x2d, 0x70, 0xcc, 0x89,
	0x98, 0x91, 0x48, 0xaa, 0xb0, 0xa6, 0xab, 0x8f, 0xf0, 0x0e, 0xd4, 0xe7, 0xec, 0x2d, 0x8a, 0xa5,
	0x20, 0x78, 0xe5, 0x85, 0x74, 0x45, 0xa5, 0xe7, 0xf3, 0x58, 0x65, 0x28, 0xba, 0x70, 0x3b, 0x1b,
	0xa7, 0x23, 0x87, 0xc7, 0xf0, 0x16, 0xc0, 0x84, 0xff, 0xd0, 0x17, 0x95, 0xbe, 0xa6, 0x27, 0xb9,
	0x7a, 0x02, 0x2e, 0xf4, 0x25, 0x3d, 0x1c, 0x48, 0x22, 0x3c, 0x4e, 0x04, 0x65, 0x73, 0xeb, 0x48,
	0x5d, 0xf2, 0x12, 0x6d, 0xca, 0x47, 0xba, 0x7c, 0xf4, 0x98, 0x95, 0x3f, 0x2c, 0xb8, 0xe7, 0x9a,
	0x7c, 0x48, 0xc1, 0xa9, 0xe2, 0x60, 0x0f, 0xd4, 0xf7, 0x0c, 0xfd, 0x77, 0x49, 0x62, 0xab, 0xd4,
	0x36, 0x3a, 0xc6, 0xb0, 0xe0, 0xc2, 0x1d, 0xc8, 0x49, 0x67, 0x70, 0x04, 0x72, 0x2b, 0x2f, 0x60,
	0x79, 0x84, 0xf2, 0x81, 0x08, 0xee, 0x99, 0xa6, 0x06, 0x2c, 0x5b, 0xef, 0xd4, 0xc0, 0xe9, 0xee,
	0xfa, 0xe6, 0x1a, 0x54, 0xf3, 0x07, 0xf8, 0xa3, 0xe8, 0x09, 0x68, 0xe4, 0x8f, 0xb3, 0xdb, 0x84,
	0x79, 0x28, 0x46, 0x3d, 0x07, 0xbf, 0x15, 0xe1, 0xb4, 0x5e, 0xae, 0x36, 0x1f, 0x84, 0x32, 0x1b,
	0x73, 0x6a, 0xef, 0x7f, 0x7f, 0xbf, 0xac, 0x9c, 0xee, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20,
	0xfd, 0xab, 0xcb, 0x19, 0x03, 0x00, 0x00,
}
