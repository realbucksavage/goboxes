package machines

import "regexp"

type GuestProperty struct {
	Name      string
	Value     string
	Timestamp string
	Flags     []string
}

var expression = regexp.MustCompile(`Name\:.(.*),.value\:.(.*),\stimestamp\:.(.*),.flags\:(.*)`)

func (g GuestProperty) FromString(info string) GuestProperty {
	m := expression.FindStringSubmatch(info)

	return GuestProperty{
		Name:      m[1],
		Value:     m[2],
		Timestamp: m[3],
	}
}
