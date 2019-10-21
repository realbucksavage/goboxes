package machines

import (
	"regexp"
	"strings"
)

// GuestProperty represents a guest VMs property
type GuestProperty struct {
	Name      string
	Value     string
	Timestamp string
	Flags     []string
}

var expression = regexp.MustCompile(`Name\:.(.*),.value\:.(.*),.timestamp\:.(.*),.flags\:(.*)`)

// FromString takes a line from the output of VBoxManage guestproperty enumerate
// and converts it to a GuestProperty object.
func (g GuestProperty) FromString(info string) GuestProperty {
	m := expression.FindStringSubmatch(info)

	var flags []string
	flagString := strings.TrimSpace(m[4])
	if flagString != "" {
		flags = strings.Split(flagString, ", ")
	}

	return GuestProperty{
		Name:      m[1],
		Value:     m[2],
		Timestamp: m[3],
		Flags:     flags,
	}
}
