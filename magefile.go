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

// GolangciLint runs the golangci-lint checker.
func GolangciLint() error {

	return goCommand("run", "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2", "run")

}

// Check ilints and formats the source code, then runs tests.
func Check() error {
	if err := Fmt(); err != nil {
		return err
	}
	if err := Vet(); err != nil {
		return err
	}
	if err := Actionlint(); err != nil {
		return err
	}
	if err := GolangciLint(); err != nil {
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
