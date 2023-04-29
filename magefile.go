//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Setup will ensure wails is installed.
// It will also run the doctor command to ensure everything is setup correctly.
func Setup() error {
	if err := sh.RunV("go", "get", "github.com/wailsapp/wails/v2/cmd/wails@latest"); err != nil {
		return err
	}

	return sh.RunV("wails", "doctor")
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	return sh.RunV("wails", "build", "-clean", "-upx")
}

type Linux mg.Namespace

// Install will take care of installing the desktop entry and binaries to your system.
// It will also install the icon to the correct location.
// This will require sudo privileges.
func (Linux) Install() error {
	if err := sh.RunV("sudo", "install", "-Dm00644", "dist/ulid-tools.desktop", "/usr/local/share/applications/ulid-tools.desktop"); err != nil {
		return err
	}
	if err := sh.RunV("sudo", "install", "-Dm00755", "build/bin/ulid-tools", "/usr/local/bin/ulid-tools"); err != nil {
		return err
	}
	return sh.RunV("sudo", "install", "-Dm00644", "dist/ulid-tools.png", "/usr/local/share/pixmaps/ulid-tools.png")
}

// Uninstall will take care of removing any desktop entries and binaries from your system.
// This will require sudo privileges.
// If you have manually removed the binaries or desktop entry, this will return an error.
func (Linux) Uninstall() error {
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
