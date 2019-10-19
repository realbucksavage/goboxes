package list

import (
	"regexp"
	"strings"

	"github.com/realbucksavage/goboxes/core"
)

// Commands struct gives access to all VBoxManage <list> commands
// TODO: implement commands
type Commands struct{}

var machineMatcher = regexp.MustCompile(`\"(.*)\"\s\{(.*)\}`)

// Vms returns all VMs in UUID - Name key value pair. The UUID may be passed to
// machines.ReadVM(string) to get a VM instance.
func (l Commands) Vms(running bool) (map[string]string, error) {
	cmdName := "vms"
	if running {
		cmdName = "runningvms"
	}

	info, err := core.ExecSubcommand("list", cmdName)
	if err != nil {
		return nil, err
	}

	vms := make(map[string]string)

	if info == "" {
		// No VMs
		return vms, nil
	}

	for _, line := range strings.Split(info, "\n") {
		m := machineMatcher.FindStringSubmatch(line)

		vms[m[2]] = m[1]
	}

	return vms, nil
}
