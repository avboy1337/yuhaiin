package tcpip

import (
	"encoding/binary"
	"net/netip"

	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/tun/tun2socket/checksum"
)

type IPProtocol = byte

// IPProtocol type
const (
	ICMP   IPProtocol = 0x01
	TCP    IPProtocol = 0x06
	UDP    IPProtocol = 0x11
	ICMPv6 IPProtocol = 0x3a
)

const (
	FlagDontFragment = 1 << 1
	FlagMoreFragment = 1 << 2
)

const IPv4HeaderSize = 20

type IPv4Packet []byte

func (p IPv4Packet) TotalLen() uint16 {
	return binary.BigEndian.Uint16(p[2:])
}

func (p IPv4Packet) SetTotalLength(length uint16) {
	binary.BigEndian.PutUint16(p[2:], length)
}

func (p IPv4Packet) HeaderLen() uint16 {
	return uint16(p[0]&0xf) * 4
}

func (p IPv4Packet) SetHeaderLen(length uint16) {
	p[0] &= 0xF0
	p[0] |= byte(length / 4)
}

func (p IPv4Packet) TypeOfService() byte {
	return p[1]
}

func (p IPv4Packet) SetTypeOfService(tos byte) {
	p[1] = tos
}

func (p IPv4Packet) Identification() uint16 {
	return binary.BigEndian.Uint16(p[4:])
}

func (p IPv4Packet) SetIdentification(id uint16) {
	binary.BigEndian.PutUint16(p[4:], id)
}

func (p IPv4Packet) FragmentOffset() uint16 {
	return binary.BigEndian.Uint16([]byte{p[6] & 0x7, p[7]}) * 8
}

func (p IPv4Packet) SetFragmentOffset(offset uint32) {
	flags := p.Flags()
	binary.BigEndian.PutUint16(p[6:], uint16(offset/8))
	p.SetFlags(flags)
}

func (p IPv4Packet) DataLen() uint16 {
	return p.TotalLen() - p.HeaderLen()
}

func (p IPv4Packet) Payload() []byte {
	return p[p.HeaderLen():p.TotalLen()]
}

func (p IPv4Packet) Protocol() IPProtocol {
	return p[9]
}

func (p IPv4Packet) SetProtocol(protocol IPProtocol) {
	p[9] = protocol
}

func (p IPv4Packet) Flags() byte {
	return p[6] >> 5
}

func (p IPv4Packet) SetFlags(flags byte) {
	p[6] &= 0x1F
	p[6] |= flags << 5
}

func (p IPv4Packet) SourceIP() netip.Addr {
	return netip.AddrFrom4([4]byte{p[12], p[13], p[14], p[15]})
}

func (p IPv4Packet) SetSourceIP(ip netip.Addr) {
	if ip.Is4() {
		copy(p[12:16], ip.AsSlice())
	}
}

func (p IPv4Packet) DestinationIP() netip.Addr {
	return netip.AddrFrom4([4]byte{p[16], p[17], p[18], p[19]})
}

func (p IPv4Packet) SetDestinationIP(ip netip.Addr) {
	if ip.Is4() {
		copy(p[16:20], ip.AsSlice())
	}
}

func (p IPv4Packet) Checksum() uint16 {
	return binary.BigEndian.Uint16(p[10:])
}

func (p IPv4Packet) SetChecksum(sum [2]byte) {
	p[10] = sum[0]
	p[11] = sum[1]
}

func (p IPv4Packet) TimeToLive() byte {
	return p[8]
}

func (p IPv4Packet) SetTimeToLive(ttl byte) {
	p[8] = ttl
}

func (p IPv4Packet) ResetChecksum() {
	p.SetChecksum(checksum.ZeroChecksum)
	p.SetChecksum(checksum.Checksum(0, p[:p.HeaderLen()]))
}

// PseudoSum for tcp checksum
func (p IPv4Packet) PseudoSum() uint32 {
	sum := checksum.Sum(p[12:20])
	sum += uint32(p.Protocol())
	sum += uint32(p.DataLen())
	return sum
}

func (p IPv4Packet) Valid() bool {
	return len(p) >= IPv4HeaderSize && uint16(len(p)) >= p.TotalLen()
}

func (p IPv4Packet) DecTimeToLive() {
	p[8] = p[8] - uint8(1)
}
