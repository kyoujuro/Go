package main

import (
  "os"
  "os/exec"
)

func main() {
  cmd := exec.Command("ping", "localhost")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
