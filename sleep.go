package main

import (
    "context"
    "fmt"
    "os/exec"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    fmt.Println("started")

    if err := exec.CommandContext(ctx,"fish", "-c", "for f in (seq 1 10); echo $f; sleep 1; end").Run(); err != nil {
        fmt.Println("stopped:", err)
    }
}
