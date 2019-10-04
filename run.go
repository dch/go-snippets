package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("fish", "-c", "for f in (seq 1 10); echo $f; sleep 1; end")
    fmt.Println("started")
    err := cmd.Run()
    fmt.Println("stopped:", err)
}
