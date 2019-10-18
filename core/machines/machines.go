package machines

import (
	"fmt"
	"strings"

	"github.com/realbucksavage/goboxes/core"
)

var poweroffState = "poweroff"
var poweronState = "running"

type VirtualMachine struct {
	UUID    string
	Name    string
	OsType  string
	VMState string

	guestPoperties []GuestProperty
}

func ReadVm(uuid string) VirtualMachine {
	m := stateInfoMap(uuid)

	return VirtualMachine{
		UUID:    m["UUID"],
		Name:    m["name"],
		OsType:  m["ostype"],
		VMState: m["VMState"],
	}
}

func (v *VirtualMachine) PowerOn() error {
	if v.VMState == poweronState {
		return fmt.Errorf("%s {%s} already in running state", v.Name, v.UUID)
	}

	core.ExecSubcommand("startvm", v.Name, "--type", "headless")
	v.VMState = poweronState

	return nil
}

func (v *VirtualMachine) PowerOff() error {
	if v.VMState == poweroffState {
		return fmt.Errorf("%s {%s} is already in poweroff state", v.Name, v.UUID)
	}

	core.ExecSubcommand("controlvm", v.Name, poweroffState)
	v.VMState = poweroffState

	return nil
}

func (v *VirtualMachine) GetProperties() []GuestProperty {
	if len(v.guestPoperties) == 0 {
		v.RefreshGuestProperties()
	}

	return v.guestPoperties
}

func (v *VirtualMachine) RefreshGuestProperties() error {
	info, err := core.ExecSubcommand("guestproperty", "enumerate", v.UUID)
	if err != nil {
		return err
	}

	if info == "" {
		return nil
	}

	props := []GuestProperty{}

	for _, line := range strings.Split(info, "\n") {
		props = append(props, GuestProperty{}.FromString(line))
	}

	v.guestPoperties = props

	return nil
}

func stateInfo(uuid string) (string, error) {
	return core.ExecSubcommand("showvminfo", uuid, "--machinereadable")
}

func stateInfoMap(uuid string) map[string]string {
	i, _ := stateInfo(uuid)
	m := make(map[string]string)
	for _, info := range strings.Split(i, "\n") {
		i := strings.Split(info, "=")

		m[i[0]] = strings.ReplaceAll(i[1], "\"", "")
	}

	return m
}
