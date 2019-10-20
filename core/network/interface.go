package network

import "strings"

// Interface struct represnts a VM network interface
type Interface struct {
	Name              string
	GUID              string
	DHCPEnabled       bool
	IPAddress         string
	NetMask           string
	IPV6Address       string
	IPV6NetMaskPfxLen string
	HardwareAddress   string
	MediumType        string
	Wireless          bool
	Status            string
	NetworkName       string
}

// ReadInterface takes a single chunk of interface data produced by
// <VBoxManage list ...> and converts it into Interface type.
func ReadInterface(info string) Interface {
	m := make(map[string]string)
	for _, line := range strings.Split(info, "\n") {
		i := strings.Split(line, ":")

		m[i[0]] = strings.TrimSpace(i[1])
	}

	dhcp := m["DHCP"] == "Enabled"
	wl := m["Wireless"] == "Yes"

	return Interface{
		Name:              m["Name"],
		GUID:              m["GUID"],
		DHCPEnabled:       dhcp,
		IPAddress:         m["IPAddress"],
		NetMask:           m["NetworkMask"],
		IPV6Address:       m["IPV6Address"],
		IPV6NetMaskPfxLen: m["IPV6NetworkMaskPrefixLength"],
		HardwareAddress:   m["HardwareAddress"],
		MediumType:        m["MediumType"],
		Wireless:          wl,
		Status:            m["Status"],
		NetworkName:       m["VBoxNetworkName"],
	}
}
