//go:build mage

package main

import (
	"os"
	"os/exec"
)

// Test runs the package test suite.
func Test() error {
	return goCommand("test", "./...")
}

// Vet runs Go's static analyzer.
func Vet() error {
	return goCommand("vet", "./...")
}

// Check formats source code, then runs vet and tests.
func Check() error {
	if err := goCommand("fmt", "./..."); err != nil {
		return err
	}
	if err := Vet(); err != nil {
		return err
	}
	return Test()
}

func goCommand(args ...string) error {
	command := exec.Command("go", args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
