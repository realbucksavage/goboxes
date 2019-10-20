# GoBoxes [![Codacy Badge](https://api.codacy.com/project/badge/Grade/cfeb88ce4434454b9138fe7f71bba1bb)](https://www.codacy.com/manual/jgodara/goboxes?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=realbucksavage/goboxes&amp;utm_campaign=Badge_Grade)

Easily control your VirtualBox VMs

## Basic Usage

```golang
import (
  "fmt"
  "github.com/realbucksavage/goboxes"
  "github.com/realbucksavage/goboxes/core/machines"
)

func main() {
  // List all vms (true -> only show running)
  vms, _ = goboxes.List().Vms(false)
  for uuid, name := range vms {
    fmt.Printf("VM UUID: %s ; Name: %s\n", uuid, name)
  }

  // Get a Virtual Machine by UUID
  uuid := "92f681e7-ec80-4a9d-9554-961face7c3f9"
  vm := machines.ReadVm(uuid)

  fmt.Printf("Power state of %s is %s\n", vm.Name, vm.VMState)

  // Starts the VM
  must(vm.PowerOn())

  // Stop the VM
  must(vm.PowerOff())

  // List network interfaces
  ifs, _ := goboxes.List().Interfaces().Bridged()
  for _, iface := range ifs {
    fmt.Println(iface.Name)
  }
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}
```
