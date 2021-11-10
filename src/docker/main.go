package main

import (
	"os"
	"os/exec"
)

func main() {
	program := "docker"
	checkVersion := "--version"

    cmd := exec.Command(program, checkVersion)

    // open the out file for writing
    outfile, err := os.Create("../docker")
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