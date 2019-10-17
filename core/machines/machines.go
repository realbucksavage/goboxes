package machines

import (
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

func (v *VirtualMachine) PowerOn() {
	if v.VMState == "running" {
		return
	}

	core.ExecSubcommand("startvm", v.Name, "--type", "headless")
}

func (v *VirtualMachine) PowerOff() {
	if v.VMState == "poweroff" {
		return
	}

	core.ExecSubcommand("controlvm", v.Name, "poweroff")
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
