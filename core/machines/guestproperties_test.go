package machines

import (
	"reflect"
	"testing"
)

func TestGuestProperty_FromString(t *testing.T) {
	testCases := []struct {
		info string
		gp   GuestProperty
	}{
		{
			info: "Name: /VirtualBox/HostInfo/GUI/LanguageID, value: en_US, timestamp: 1554430700004960000, flags:",
			gp: GuestProperty{
				Name:      "/VirtualBox/HostInfo/GUI/LanguageID",
				Value:     "en_US",
				Timestamp: "1554430700004960000",
				Flags:     nil,
			},
		},
		{
			info: "Name: /VirtualBox/VMInfo/ResumeCounter, value: 5, timestamp: 1554409878043994000, flags: TRANSIENT, RDONLYGUEST",
			gp: GuestProperty{
				Name:      "/VirtualBox/VMInfo/ResumeCounter",
				Value:     "5",
				Timestamp: "1554409878043994000",
				Flags:     []string{"TRANSIENT", "RDONLYGUEST"},
			},
		},
	}

	for _, tc := range testCases {
		gp := GuestProperty{}.FromString(tc.info)
		if !reflect.DeepEqual(gp, tc.gp) {
			t.Fatalf("Expected %#v but got %#v", tc.gp, gp)
		}
	}
}
