//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Deps() error {
	return sh.RunV("go", "get", "github.com/wailsapp/wails/v2/cmd/wails@latest")
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(Deps)
	return sh.RunV("wails", "build", "-clean", "-upx")
}

func Install() error {
	if err := sh.RunV("sudo", "install", "-Dm00644", "dist/ulid-tools.desktop", "/usr/local/share/applications/ulid-tools.desktop"); err != nil {
		return err
	}
	if err := sh.RunV("sudo", "install", "-Dm00755", "build/bin/ulid-tools", "/usr/local/bin/ulid-tools"); err != nil {
		return err
	}
	return sh.RunV("sudo", "install", "-Dm00644", "dist/ulid-tools.png", "/usr/local/share/pixmaps/ulid-tools.png")
}

// Uninstall will take care of removing any desktop entries and binaries from your system.
func Uninstall() error {
	if err := sh.RunV("sudo", "rm", "/usr/local/share/applications/ulid-tools.desktop"); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("sequel is likely not installed: %v", err)
		}
		return err
	}
	if err := sh.RunV("sudo", "rm", "/usr/local/bin/ulid-tools"); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("ulid-tools is likely not installed: %v", err)
		}
		return err
	}
	err := sh.RunV("sudo", "rm", "/usr/local/share/pixmaps/ulid-tools.png")
	if os.IsNotExist(err) {
		return fmt.Errorf("ulid-tools is likely not installed: %v", err)
	}
	return err
}
