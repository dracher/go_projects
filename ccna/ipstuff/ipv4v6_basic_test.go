package ipstuff

import (
	"testing"
)

var (
	ipv4Valid   = []string{"192.168.1.1", "10.66.10.88", "4.4.4.4", "113.114.115.116"}
	ipv4Invalid = []string{"1..2.3", "1111", "", "192.168.123", "311.123.22.1"}
	ipv4Special = []string{"0.0.0.0", "255.255.255.255"}
)

func TestIPv4IPv6Distinguish(t *testing.T) {

	for _, ipd := range ipv4Invalid {
		t.Log(ipd)
		i := IPAddress{ipd}

		err, tp := i.IsValidIPv4OrIPv6()

		if err != nil || tp == 0 {
			t.Error(err)
		} else {
			t.Logf("%s is IPv%d\n", i.Address, tp)
		}
	}
}
