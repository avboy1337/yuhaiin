package pool

import (
	"bytes"
	"math/bits"
	"sync"

	"github.com/Asutorufa/yuhaiin/pkg/utils/syncmap"
)

type Pool interface {
	GetBytes(size int) []byte
	PutBytes(b []byte)

	GetBuffer() *bytes.Buffer
	PutBuffer(b *bytes.Buffer)
}

const DefaultSize = 16 * 0x400

var DefaultPool Pool = &pool{}

func GetBytes(size int) []byte  { return DefaultPool.GetBytes(size) }
func PutBytes(b []byte)         { DefaultPool.PutBytes(b) }
func GetBuffer() *bytes.Buffer  { return DefaultPool.GetBuffer() }
func PutBuffer(b *bytes.Buffer) { DefaultPool.PutBuffer(b) }

var poolMap syncmap.SyncMap[int, *sync.Pool]

type pool struct{}

func buffPool(size int) *sync.Pool {
	if v, ok := poolMap.Load(size); ok {
		return v
	}

	p := &sync.Pool{New: func() any { return make([]byte, size) }}
	poolMap.Store(size, p)
	return p
}

func (pool) GetBytes(size int) []byte {
	l := bits.Len(uint(size)) - 1
	if size != 1<<l {
		size = 1 << (l + 1)
	}
	return buffPool(size).Get().([]byte)
}

func (pool) PutBytes(b []byte) {
	l := bits.Len(uint(len(b))) - 1
	buffPool(1 << l).Put(b) //lint:ignore SA6002 ignore temporarily
}

var bufpool = sync.Pool{New: func() any { return bytes.NewBuffer(nil) }}

func (pool) GetBuffer() *bytes.Buffer { return bufpool.Get().(*bytes.Buffer) }
func (pool) PutBuffer(b *bytes.Buffer) {
	if b != nil {
		b.Reset()
		bufpool.Put(b)
	}
}