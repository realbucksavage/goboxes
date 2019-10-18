package machines

import (
	"fmt"
	"strings"

	"github.com/realbucksavage/goboxes/core"
)

var poweroffState = "poweroff"
var poweronState = "running"

// VirtualMachine represents a Virtual Machine
type VirtualMachine struct {
	UUID    string
	Name    string
	OsType  string
	VMState string

	guestPoperties []GuestProperty
}

// ReadVM reads the information of the VM identified by uuid in a VirtualMachine
// object.
func ReadVM(uuid string) VirtualMachine {
	m := stateInfoMap(uuid)

	return VirtualMachine{
		UUID:    m["UUID"],
		Name:    m["name"],
		OsType:  m["ostype"],
		VMState: m["VMState"],
	}
}

// PowerOn starts a Virtual Mchine. This will return an error if the VM is already
// in the ON state.
func (v *VirtualMachine) PowerOn() error {
	if v.VMState == poweronState {
		return fmt.Errorf("%s {%s} already in running state", v.Name, v.UUID)
	}

	core.ExecSubcommand("startvm", v.Name, "--type", "headless")
	v.VMState = poweronState

	return nil
}

// PowerOff shuts a Virtual Machine down. This will return an error if the VM is
// already in the OFF state.
func (v *VirtualMachine) PowerOff() error {
	if v.VMState == poweroffState {
		return fmt.Errorf("%s {%s} is already in poweroff state", v.Name, v.UUID)
	}

	core.ExecSubcommand("controlvm", v.Name, poweroffState)
	v.VMState = poweroffState

	return nil
}

// GetProperties retuns the properties attached to the guest VM.
// TODO: implement a way to cache these properties
func (v *VirtualMachine) GetProperties() ([]GuestProperty, error) {
	info, err := core.ExecSubcommand("guestproperty", "enumerate", v.UUID)
	if err != nil {
		return nil, err
	}

	props := []GuestProperty{}

	if info == "" {
		return props, nil
	}

	for _, line := range strings.Split(info, "\n") {
		props = append(props, GuestProperty{}.FromString(line))
	}

	v.guestPoperties = props

	return v.guestPoperties, nil
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
