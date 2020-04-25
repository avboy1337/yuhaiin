wget https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/apple.china.conf -O apple.china.conf
wget https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/google.china.conf -O google.china.conf
wget https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf -O accelerated-domains.china.conf
cat custom.conf > yuhaiin.conf
cat accelerated-domains.china.conf | sed 's/server=\///g' | sed 's/\/114\.114\.114\.114/ direct/g' >> yuhaiin.conf
cat google.china.conf | sed 's/server=\///g' | sed 's/\/114\.114\.114\.114/ direct/g' >> yuhaiin.conf
cat apple.china.conf | sed 's/server=\///g' | sed 's/\/114\.114\.114\.114/ direct/g' >> yuhaiin.conf
cat ../cn/cn.acl | sed 's/$/ direct/g' >> yuhaiin.conf
cat ../common/lan.acl | sed 's/$/ direct/g' >> yuhaiin.conf
cat abroad.conf | sed 's/$/ proxy/g' >> yuhaiin.conf
rm apple.china.conf google.china.conf accelerated-domains.china.conf