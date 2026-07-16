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

// Actionlint checks GitHub Actions workflows.
func Actionlint() error {
	return goCommand("run", "github.com/rhysd/actionlint/cmd/actionlint@v1.7.12", "-color")
}

// Fmt formats Go source code using "go fmt".
func Fmt() error {

	return goCommand("fmt", "./...")

}

// Check formats source code, then runs vet and tests.
func Check() error {
	if err := Fmt(); err != nil {
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
