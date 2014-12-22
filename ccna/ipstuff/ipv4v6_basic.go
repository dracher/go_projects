package ipstuff

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	ipv4Pattern = `\d{1,3}\\.\d{1,3}\\.\d{1,3}\\.\d{1,3}`
	ipv6Pattern = `[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}:[0-9a-fA-F]{1,4}`
)

type IPAddress struct {
	Address string
}

func (ipd *IPAddress) IsValidIPv4OrIPv6() (error, int) {
	res, _ := regexp.MatchString(ipv4Pattern, ipd.Address)
	if res == true {

		return nil, 4
	}
	fmt.Printf("%s is not fit IPv4 pattern, continue to test fit IPv6 pattern\n", ipd.Address)

	res6, _ := regexp.MatchString(ipv6Pattern, ipd.Address)
	if res6 == true {
		return nil, 6
	}
	fmt.Printf("%s neither IPv4 nor IPv6", ipd.Address)
	err := errors.New("Invalid input")
	return err, 0
}
