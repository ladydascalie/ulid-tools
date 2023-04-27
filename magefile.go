//go:build mage
// +build mage

package main

import (
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

func Deps() error {
	cmd := exec.Command("go", "install", "github.com/wailsapp/wails/v2/cmd/wails@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(Deps)
	cmd := exec.Command("wails", "build", "-clean", "-upx")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
