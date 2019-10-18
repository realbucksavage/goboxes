package machines

import (
	"fmt"
	"strings"

	"github.com/realbucksavage/goboxes/core"
)

type VirtualMachine struct {
	UUID    string
	Name    string
	OsType  string
	VMState string
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
	if v.VMState == "running" {
		return fmt.Errorf("%s {%s} already in running state", v.Name, v.UUID)
	}

	core.ExecSubcommand("startvm", v.Name, "--type", "headless")

	return nil
}

func (v *VirtualMachine) PowerOff() error {
	if v.VMState == "poweroff" {
		return fmt.Errorf("%s {%s} is already in poweroff state", v.Name, v.UUID)
	}

	core.ExecSubcommand("controlvm", v.Name, "poweroff")

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
