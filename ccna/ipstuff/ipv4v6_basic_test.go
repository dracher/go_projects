package ipstuff

import (
	"testing"
)

func TestIPv4IPv6Distinguish(t *testing.T) {

	// ipv4 := "10.66.11.22"
	ipv6 := "2001:0db8:0000:0000:0000:ff00:0042:8329"

	ipt := IPAddress{ipv6}
	err, tp := ipt.IsValidIPv4OrIPv6()

	if err != nil || tp == 0 {
		t.Error(err)
	} else {
		t.Logf("%s is IPv%d\n", ipt.Address, tp)
	}
}
