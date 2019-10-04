package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "github.com/briandowns/jail"
)

func main() {
    o := &jail.Opts{
        Version:  uint32(2),
        Path:     "/",
        Name:     "maestro",
        Hostname: "maestro.local",
        IP4:      "100.64.0.10",
        Chdir:    false,
    }
    jid, err := jail.Jail(o)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("JID: %d - / directory listing in jail", jid)
    files, err := ioutil.ReadDir(".")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, f := range files {
        fmt.Println(f.Name())
    }
}
