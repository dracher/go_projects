package ipstuff

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ipv4Pattern = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
)

type IPAddress struct {
	Address string
}

func (ipd *IPAddress) IsValidIPv4OrIPv6() (error, int) {
	res, _ := regexp.MatchString(ipv4Pattern, ipd.Address)
	if res == true {
		for _, i := range strings.Split(ipd.Address, ".") {
			i_int, _ := strconv.ParseInt(i, 10, 16)
			if i_int > 255 {
				err := errors.New("ipv4 address can not greater than 255")
				return err, 0
			}
		}
		return nil, 4
	}
	fmt.Printf("%s is not fit IPv4 pattern, continue to test fit IPv6 pattern\n", ipd.Address)

	res6, _ := regexp.MatchString(ipv6Pattern, ipd.Address)
	if res6 == true {
		return nil, 6
	}
	fmt.Printf("%s neither IPv4 nor IPv6\n", ipd.Address)
	err := errors.New("Invalid input")
	return err, 0
}
