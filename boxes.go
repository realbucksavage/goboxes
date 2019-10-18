package goboxes

import (
	"github.com/realbucksavage/goboxes/commands/list"
)

// List function gives access to all possible <list> commands.
func List() list.Commands {
	return list.Commands{}
}
