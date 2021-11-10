package main

import (
	"os"
	"os/exec"
)

// cat /etc/os-release | grep -oP '"\K[^"\047]+(?=["\047])'

func main() {
	printVersion := "cat"
	osRelease := "/etc/os-release"

    cmd := exec.Command(printVersion, osRelease)

    // open the out file for writing
    outfile, err := os.Create("../linux")
    if err != nil {
        panic(err)
    }
    defer outfile.Close()
    cmd.Stdout = outfile

    err = cmd.Start(); if err != nil {
        panic(err)
    }
    cmd.Wait()
}
