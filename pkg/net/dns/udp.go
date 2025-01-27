package dns

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/Asutorufa/yuhaiin/pkg/net/interfaces/dns"
	"github.com/Asutorufa/yuhaiin/pkg/net/nat"
	pdns "github.com/Asutorufa/yuhaiin/pkg/protos/config/dns"
	"github.com/Asutorufa/yuhaiin/pkg/protos/statistic"
	"github.com/Asutorufa/yuhaiin/pkg/utils/syncmap"
)

func init() {
	Register(pdns.Type_udp, NewDoU)
	Register(pdns.Type_reserve, NewDoU)
}

type udp struct {
	*client

	packetConn net.PacketConn
	mu         sync.Mutex
	bufChanMap syncmap.SyncMap[[2]byte, *bufChan]
}

func (u *udp) Close() error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.packetConn != nil {
		err := u.packetConn.Close()
		u.packetConn = nil
		return err
	}
	return nil
}

func (u *udp) handleResponse() {
	defer u.Close()

	for {
		buf := make([]byte, nat.MaxSegmentSize)
		n, _, err := u.packetConn.ReadFrom(buf)
		if err != nil {
			return
		}

		if n < 2 {
			continue
		}

		c, ok := u.bufChanMap.Load([2]byte(buf[:2]))
		if !ok {
			continue
		}

		c.Send(buf[:n])
	}
}

func (u *udp) initPacketConn() (net.PacketConn, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.packetConn != nil {
		return u.packetConn, nil
	}

	addr, err := ParseAddr(statistic.Type_udp, u.config.Host, "53")
	if err != nil {
		return nil, fmt.Errorf("parse addr failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	addr.WithContext(ctx)

	conn, err := u.config.Dialer.PacketConn(addr)
	if err != nil {
		return nil, fmt.Errorf("get packetConn failed: %w", err)
	}

	u.packetConn = conn
	go u.handleResponse()

	return conn, nil
}

type bufChan struct {
	closed  bool
	mu      sync.Mutex
	bufChan chan []byte
}

func (b *bufChan) Send(buf []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}
	b.bufChan <- buf
}

func (b *bufChan) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.closed = true
	close(b.bufChan)
}

func NewDoU(config Config) (dns.DNS, error) {
	addr, err := ParseAddr(statistic.Type_udp, config.Host, "53")
	if err != nil {
		return nil, fmt.Errorf("parse addr failed: %w", err)
	}

	udp := &udp{}

	udp.client = NewClient(config, func(req []byte) ([]byte, error) {

		packetConn, err := udp.initPacketConn()
		if err != nil {
			return nil, err
		}
		id := [2]byte{req[0], req[1]}

	_retry:
		_, ok := udp.bufChanMap.Load([2]byte(req[:2]))
		if ok {
			rand.Read(req[:2])
			goto _retry
		}

		bchan, _ := udp.bufChanMap.LoadOrStore([2]byte(req[:2]), &bufChan{bufChan: make(chan []byte)})
		defer func() {
			udp.bufChanMap.Delete([2]byte(req[:2]))
			bchan.Close()
		}()

		_, err = packetConn.WriteTo(req, addr)
		if err != nil {
			return nil, err
		}

		select {
		case <-time.After(time.Second * 10):
			return nil, fmt.Errorf("timeout")
		case data := <-bchan.bufChan:
			data[0] = id[0]
			data[1] = id[1]
			return data, nil
		}
	})

	return udp, nil
}
