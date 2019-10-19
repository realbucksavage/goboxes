package core

import (
	"os/exec"
	"strings"
)

// ExecSubcommand executes a VBoxManage subcommand and returns output as string
func ExecSubcommand(commands ...string) (string, error) {
	bytes, err := exec.Command("VBoxManage", commands...).Output()
	return strings.TrimSpace(string(bytes)), err
}
