package interfaces

import (
	"strings"

	"github.com/realbucksavage/goboxes/core"
	"github.com/realbucksavage/goboxes/core/network"
)

// Commands struct gives access to listable network interfaces
type Commands struct{}

// HostOnly is the equivalent of "hostonlyifs"
func (l Commands) HostOnly() ([]network.Interface, error) {
	return listIfacecType("hostonlyifs")
}

// Bridged is the equivalent of "bridgedifs"
func (l Commands) Bridged() ([]network.Interface, error) {
	return listIfacecType("bridgedifs")
}

func listIfacecType(iftype string) ([]network.Interface, error) {
	info, err := core.ExecSubcommand("list", iftype)
	if err != nil {
		return nil, err
	}

	ifaces := []network.Interface{}
	rawIfaces := strings.Split(info, "\n\n")

	for _, i := range rawIfaces {
		ifaces = append(ifaces, network.ReadInterface(i))
	}

	return ifaces, nil
}
