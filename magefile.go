//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Test runs the package test suite.
func Test() error {
	fmt.Println("Running unit tests...")
	return goCommand("test", "./...")
}

// Vet runs Go's static analyzer.
func Vet() error {
	fmt.Println("Vetting the source code...")
	return goCommand("vet", "./...")
}

// ActionLint checks GitHub Actions workflows.
func ActionLint() error {

	fmt.Println("Linting the github workflow files...")
	return goCommand("run", "github.com/rhysd/actionlint/cmd/actionlint@v1.7.12", "-color")
}

// Fmt formats Go source code using "go fmt".
func Fmt() error {
	fmt.Println("Formatting the source code...")
	return goCommand("fmt", "./...")

}

// GolangciLint runs the golangci-lint checker.
func GolangciLint() error {

	fmt.Println("Linting the source code with golangci-lint...")
	return goCommand("run", "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2", "run")

}

// Lint runs various linters: It formats and lints the on go code and lints the workflow files.
func Lint() error {
	if err := Fmt(); err != nil {
		return err
	}
	if err := Vet(); err != nil {
		return err
	}
	if err := ActionLint(); err != nil {
		return err
	}
	return GolangciLint()
}

// Check runs linters and unit tests.
func Check() error {

	if err := Lint(); err != nil {
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
