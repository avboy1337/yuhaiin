package mapper

import (
	"encoding/json"
	"strings"

	"github.com/Asutorufa/yuhaiin/pkg/net/interfaces/proxy"
)

var (
	_        uint8 = 0
	last     uint8 = 1
	wildcard uint8 = 2
)

type domainNode[T any] struct {
	Symbol uint8                     `json:"symbol"`
	Mark   T                         `json:"mark"`
	Child  map[string]*domainNode[T] `json:"child"`
}

type domainStr struct {
	domain string
	aft    int
	pre    int
}

func newDomainStr(domain string) *domainStr {
	return &domainStr{
		domain: domain,
		aft:    len(domain),
		pre:    strings.LastIndexByte(domain, '.') + 1,
	}
}

func (d *domainStr) hasNext() bool {
	return d.aft >= 0
}

func (d *domainStr) last() bool {
	return d.pre == 0
}

func (d *domainStr) next() bool {
	d.aft = d.pre - 1
	if d.aft < 0 {
		return false
	}
	d.pre = strings.LastIndexByte(d.domain[:d.aft], '.') + 1
	return true
}

var valueEmpty = string([]byte{0x03})

func (d *domainStr) str() string {
	if d.pre == d.aft {
		return valueEmpty
	}
	return d.domain[d.pre:d.aft]
}

func s[T any](root *domainNode[T], domain string) (resp T, ok bool) {
	s := root
	z := newDomainStr(domain)
	first, asterisk := true, false

	for {
		if !z.hasNext() || s == nil {
			return
		}

		if r, okk := s.Child[z.str()]; okk {
			if r.Symbol != 0 {
				if r.Symbol == wildcard {
					resp, ok = r.Mark, true
				}

				if r.Symbol == last && z.last() {
					return r.Mark, true
				}
			}

			s, first, _ = r, false, z.next()
			continue
		}

		if !first {
			return
		}

		if !asterisk {
			s, asterisk = s.Child["*"], true
		} else {
			z.next()
		}
	}
}

func insert[T any](root *domainNode[T], domain string, mark T) {
	z := newDomainStr(domain)
	for z.hasNext() {
		if z.last() && domain[0] == '*' {
			root.Symbol, root.Mark = wildcard, mark
			break
		}

		if root.Child == nil {
			root.Child = make(map[string]*domainNode[T])
		}

		if root.Child[z.str()] == nil {
			root.Child[z.str()] = &domainNode[T]{}
		}

		root = root.Child[z.str()]

		if z.last() {
			root.Symbol, root.Mark = last, mark
		}

		z.next()
	}
}

type domain[T any] struct {
	Root         *domainNode[T] `json:"root"`          // for example.com, example.*
	WildcardRoot *domainNode[T] `json:"wildcard_root"` // for *.example.com, *.example.*
}

func (d *domain[T]) Insert(domain string, mark T) {
	if len(domain) == 0 {
		return
	}

	if domain[0] == '*' {
		insert(d.WildcardRoot, domain, mark)
	} else {
		insert(d.Root, domain, mark)
	}
}

func (d *domain[T]) Search(domain proxy.Address) (mark T, ok bool) {
	mark, ok = s(d.Root, domain.Hostname())
	if ok {
		return
	}
	return s(d.WildcardRoot, domain.Hostname())
}

func (d *domain[T]) Marshal() ([]byte, error) {
	return json.MarshalIndent(d, "", "  ")
}

func (d *domain[T]) Clear() error {
	d.Root = &domainNode[T]{Child: map[string]*domainNode[T]{}}
	d.WildcardRoot = &domainNode[T]{Child: map[string]*domainNode[T]{}}
	return nil
}

func NewDomainMapper[T any]() *domain[T] {
	return &domain[T]{
		Root:         &domainNode[T]{Child: map[string]*domainNode[T]{}},
		WildcardRoot: &domainNode[T]{Child: map[string]*domainNode[T]{}},
	}
}
