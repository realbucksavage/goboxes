package list

import (
	"regexp"
	"strings"

	"github.com/realbucksavage/goboxes/core"
)

type ListCommands struct{}

var machineMatcher = regexp.MustCompile(`\"(.*)\"\s\{(.*)\}`)

func (l ListCommands) Vms(running bool) (map[string]string, error) {
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
