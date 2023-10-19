//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	BUILD_DIR = "build"
)

// Runs go mod download
func Download() error {
	return sh.Run("go", "mod", "download")
}

// Installs the dev dependencies.
// Dependencies are: mage, golangci-lint
func InstallDevDependencies() error {
	tools := []string{
		"github.com/magefile/mage",
		"github.com/golangci/golangci-lint/cmd/golangci-lint",
		"github.com/cosmtrek/air",
	}

	for _, tool := range tools {
		if err := sh.Run("go", "install", tool+"@latest"); err != nil {
			return err
		}
	}

	return nil
}

// Runs go mod download and then installs the binary.
func Build() error {
	mg.Deps(Download)

	return sh.Run("go", "build", "./cmd/server", "-o", BUILD_DIR)
}

// Run go tests.
func Test() error {
	mg.Deps(Download)

	return sh.Run("go", "test", "-coverprofile", BUILD_DIR+"/cover.out", "-v", "./...")
}

// Lint the code.
func Lint() error {
	mg.Deps(InstallDevDependencies)
	return sh.Run("golangci-lint", "run", "./...")
}

// Starts the development server with air featuring live reload.
func Dev() error {
	mg.Deps(Download)
	return sh.Run("air")
}

// Opens the coverage report in the browser.
func Coverage() error {
	mg.Deps(Test)
	return sh.Run("go", "tool", "cover", "-html="+BUILD_DIR+"/cover.out")
}
