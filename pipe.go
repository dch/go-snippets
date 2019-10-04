package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
)

func main() {
    cmdName := "fish"
    cmdArgs := []string{"-c", "for f in (seq 1 10); echo $f; sleep 1; end"}

    cmd := exec.Command(cmdName, cmdArgs...)
    cmdReader, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Fprintln(os.Stderr, "pipefail", err)
        os.Exit(1)
    }

    scanner := bufio.NewScanner(cmdReader)
    go func() {
        for scanner.Scan() {
            fmt.Printf("pipe stdout | %s\n", scanner.Text())
        }
    }()

    err = cmd.Start()
    if err != nil {
        fmt.Fprintln(os.Stderr, "startfail", err)
        os.Exit(1)
    }

    err = cmd.Wait()
    if err != nil {
        fmt.Fprintln(os.Stderr, "waitfail", err)
        os.Exit(1)
    }
}
