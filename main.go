package main

import (
	"os"
	"os/exec"
)

func main() {
	// Simple wrapper to run a command provided via CLI arguments
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin // Connect standard input
	cmd.Stdout = os.Stdout // Connect standard output
	cmd.Stderr = os.Stderr // Connect standard error

	if err := cmd.Run(); err != nil {
		os.Exit(1) // Exit with error code if command fails
	}
}