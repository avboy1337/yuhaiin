// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package channel provides the implemention of channel-based data-link layer
// endpoints. Such endpoints allow injection of inbound packets and store
// outbound packets in a channel.

package tun

import (
	"log"

	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

type writer interface {
	Write([]byte) tcpip.Error
	WritePacket(*stack.PacketBuffer) tcpip.Error
	WritePackets(stack.PacketBufferList) (int, tcpip.Error)
}

type inbound interface {
	stop()
	dispatch() (bool, tcpip.Error)
}

var _ stack.InjectableLinkEndpoint = (*Endpoint)(nil)

// Endpoint is link layer endpoint that stores outbound packets in a channel
// and allows injection of inbound packets.
type Endpoint struct {
	dispatcher         stack.NetworkDispatcher
	mtu                uint32
	linkAddr           tcpip.LinkAddress
	LinkEPCapabilities stack.LinkEndpointCapabilities

	writer  writer
	inbound inbound
	closed  chan struct{}
}

// New creates a new channel endpoint.
func NewEndpoint(w writer, mtu uint32, linkAddr tcpip.LinkAddress) *Endpoint {
	return &Endpoint{
		mtu:      mtu,
		linkAddr: linkAddr,
		writer:   w,
	}
}

func (e *Endpoint) SetInbound(i inbound) { e.inbound = i }

// Close closes e. Further packet injections will return an error, and all pending
// packets are discarded. Close may be called concurrently with WritePackets.
func (e *Endpoint) Close() {
	if e.closed == nil {
		return
	}

	select {
	case <-e.closed:
		return
	default:
	}

	e.inbound.stop()
	<-e.closed
	e.closed = nil
}

// InjectInbound injects an inbound packet.
func (e *Endpoint) InjectInbound(protocol tcpip.NetworkProtocolNumber, pkt *stack.PacketBuffer) {
	e.dispatcher.DeliverNetworkPacket("", "", protocol, pkt)
}

// InjectOutbound writes a fully formed outbound packet directly to the
// link.
//
// dest is used by endpoints with multiple raw destinations.
func (e *Endpoint) InjectOutbound(dest tcpip.Address, packet []byte) tcpip.Error {
	return e.writer.Write(packet)
}

// Attach saves the stack network-layer dispatcher for use later when packets
// are injected.
func (e *Endpoint) Attach(dispatcher stack.NetworkDispatcher) {
	if dispatcher == nil && e.IsAttached() {
		e.Close()
	}

	if dispatcher != nil && !e.IsAttached() {
		e.closed = make(chan struct{})
		go func() {
			defer close(e.closed)
			for {
				cont, err := e.inbound.dispatch()
				if err != nil || !cont {
					log.Println("dispatch exit:", err)
					break
				}
			}
		}()
	}

	e.dispatcher = dispatcher
}

// IsAttached implements stack.LinkEndpoint.IsAttached.
func (e *Endpoint) IsAttached() bool {
	return e.dispatcher != nil
}

// MTU implements stack.LinkEndpoint.MTU. It returns the value initialized
// during construction.
func (e *Endpoint) MTU() uint32 {
	return e.mtu
}

// Capabilities implements stack.LinkEndpoint.Capabilities.
func (e *Endpoint) Capabilities() stack.LinkEndpointCapabilities {
	return e.LinkEPCapabilities
}

// MaxHeaderLength returns the maximum size of the link layer header. Given it
// doesn't have a header, it just returns 0.
func (*Endpoint) MaxHeaderLength() uint16 {
	return 0
}

// LinkAddress returns the link address of this endpoint.
func (e *Endpoint) LinkAddress() tcpip.LinkAddress {
	return e.linkAddr
}

// WritePackets stores outbound packets into the channel.
// Multiple concurrent calls are permitted.
// func (e *Endpoint) WritePackets(pkts stack.PacketBufferList) (int, tcpip.Error) {
// 	iovecs := make([]unix.Iovec, 0, 47)
// 	for _, pkt := range pkts.AsSlice() {
// 		/*
// 			    old method
// 				sendBuffer := utils.GetBuffer()
// 				defer utils.PutBuffer(sendBuffer)
// 				sendBuffer.Reset()
// 				sendBuffer.Write(pkt.NetworkHeader().View())
// 				sendBuffer.Write(pkt.TransportHeader().View())
// 				sendBuffer.Write(pkt.Data().AsRange().ToOwnedView())
// 		*/

// 		for _, s := range pkt.Slices() {
// 			iovecs = append(iovecs, rawfile.IovecFromBytes(s))
// 		}
// 	}
// 	if err := rawfile.NonBlockingWriteIovec(e.fd, iovecs); err != nil {
// 		return 0, err
// 	}
// 	return pkts.Len(), nil
// }

func (e *Endpoint) WritePacket(_ stack.RouteInfo, _ tcpip.NetworkProtocolNumber, pkt *stack.PacketBuffer) tcpip.Error {
	return e.writer.WritePacket(pkt)
}

func (e *Endpoint) WritePackets(_ stack.RouteInfo, pkts stack.PacketBufferList, _ tcpip.NetworkProtocolNumber) (int, tcpip.Error) {
	return e.writer.WritePackets(pkts)
}

func (e *Endpoint) WriteRawPacket(pkt *stack.PacketBuffer) tcpip.Error {
	return e.writer.WritePacket(pkt)
}

// Wait implements stack.LinkEndpoint.Wait.
func (*Endpoint) Wait() {}

// ARPHardwareType implements stack.LinkEndpoint.ARPHardwareType.
func (*Endpoint) ARPHardwareType() header.ARPHardwareType { return header.ARPHardwareNone }

// AddHeader implements stack.LinkEndpoint.AddHeader.
// func (*Endpoint) AddHeader(*stack.PacketBuffer) {}
func (*Endpoint) AddHeader(tcpip.LinkAddress, tcpip.LinkAddress, tcpip.NetworkProtocolNumber, *stack.PacketBuffer) {
}
