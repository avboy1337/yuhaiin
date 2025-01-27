package dns

import (
	"net/netip"
	"testing"

	s5c "github.com/Asutorufa/yuhaiin/pkg/net/proxy/socks5/client"
	rr "github.com/Asutorufa/yuhaiin/pkg/net/resolver"
	"github.com/Asutorufa/yuhaiin/pkg/protos/config/dns"
	"github.com/Asutorufa/yuhaiin/pkg/utils/assert"
)

func TestDOH(t *testing.T) {
	rr.Bootstrap = &rr.System{DisableIPv6: true}
	s, err := netip.ParsePrefix("223.5.5.5/24")
	assert.NoError(t, err)
	s5Dialer := s5c.Dial("127.0.0.1", "1080", "", "")

	configMap := map[string]Config{
		"google": {
			Type:   dns.Type_doh,
			Host:   "dns.google",
			IPv6:   true,
			Subnet: s,
			Dialer: s5Dialer,
		},
		"iijJP": {
			Type:       dns.Type_doh,
			Host:       "103.2.57.5",
			Servername: "public.dns.iij.jp",
			IPv6:       true,
			Subnet:     s,
		},
		"cloudflare": {
			Type:   dns.Type_doh,
			Host:   "cloudflare-dns.com",
			Subnet: s,
			IPv6:   true,
			Dialer: s5Dialer,
		},
		"quad9": {
			Type:   dns.Type_doh,
			Host:   "9.9.9.9",
			Subnet: s,
			IPv6:   true,
		},
		"a.passcloud": {
			Type:   dns.Type_doh,
			Host:   "a.passcloud.xyz",
			Subnet: s,
			IPv6:   true,
		},
		"adguard": {
			Type: dns.Type_doh,
			Host: "https://unfiltered.adguard-dns.com/dns-query",
			IPv6: true,
		},
		"ali": {
			Type:   dns.Type_doh,
			Host:   "223.5.5.5",
			Subnet: s,
		},
		"dns.pub": {
			Type:   dns.Type_doh,
			Host:   "dns.pub",
			Subnet: s,
			IPv6:   true,
		},
		"sm2.dnspub": {
			Type:   dns.Type_doh,
			Host:   "sm2.doh.pub",
			Subnet: s,
			IPv6:   true,
		},
		"360": {
			Type:   dns.Type_doh,
			Host:   "doh.360.cn",
			Subnet: s,
			IPv6:   true,
		},
		"dnssb": {
			Type:   dns.Type_doh,
			Host:   "doh.dns.sb",
			Subnet: s,
			IPv6:   true,
			Dialer: s5Dialer,
		},
		"opendns": {
			Type:   dns.Type_doh,
			Host:   "doh.opendns.com",
			Subnet: s,
			IPv6:   true,
			Dialer: s5Dialer,
		},
	}

	d, err := New(configMap["google"])
	assert.NoError(t, err)

	t.Log(d.LookupIP("plasma"))
	t.Log(d.LookupIP("fonts.gstatic.com"))
	t.Log(d.LookupIP("dc.services.visualstudio.com")) // -> will error, but not found reason
	t.Log(d.LookupIP("i2.hdslb.com"))
	t.Log(d.LookupIP("www.baidu.com"))
	t.Log(d.LookupIP("push.services.mozilla.com"))
	t.Log(d.LookupIP("www.google.com"))
	t.Log(d.LookupIP("www.pixiv.net"))
	t.Log(d.LookupIP("s1.hdslb.com"))
	t.Log(d.LookupIP("dns.nextdns.io"))
	t.Log(d.LookupIP("bilibili.com"))
	t.Log(d.LookupIP("test-ipv6.com"))
	t.Log(d.LookupIP("ss1.bdstatic.com"))
	t.Log(d.LookupIP("www.twitter.com"))
	t.Log(d.LookupIP("www.facebook.com"))
	t.Log(d.LookupIP("yahoo.co.jp"))
	t.Log(d.LookupIP("115-235-111-150.dhost.00cdn.com"))
}

func TestGetURL(t *testing.T) {
	dnsQStr := "https://8.8.8.8/dns-query"
	assert.Equal(t, getUrlAndHost("https://8.8.8.8"), dnsQStr)
	assert.Equal(t, getUrlAndHost("8.8.8.8"), dnsQStr)

	t.Log(getUrlAndHost("https://"))
	t.Log(getUrlAndHost("/dns-query"))
}
