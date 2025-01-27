//go:build linux || darwin
// +build linux darwin

package main

import (
	"fmt"

	"github.com/Asutorufa/yuhaiin/pkg/net/interfaces/proxy"
	"github.com/Asutorufa/yuhaiin/pkg/net/netlink"
	"github.com/Asutorufa/yuhaiin/pkg/utils/yerror"
)

func init() {
	processDumper = processDumperImpl{}
}

type processDumperImpl struct{}

func (processDumperImpl) ProcessName(network string, src, _ proxy.Address) (string, error) {
	if src.Type() != proxy.IP {
		return "", fmt.Errorf("source address is not ip")
	}
	return netlink.FindProcessName(network, yerror.Ignore(src.IP()), src.Port().Port())
}
