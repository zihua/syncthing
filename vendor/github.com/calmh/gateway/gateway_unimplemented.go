// +build !darwin,!linux,!windows

package gateway

import (
	"fmt"
	"net"
	"runtime"
)

func DiscoverGateway() (ip net.IP, err error) {
	err = fmt.Errorf("DiscoverGateway not implemented for OS %s", runtime.GOOS)
	return
}
