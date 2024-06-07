package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run `ant install` command
	installCmd := exec.Command("ant", "install")
	installOutput, installErr := installCmd.CombinedOutput()
	if installErr != nil {
		fmt.Println("Error running 'ant install':", installErr)
		return
	}
	fmt.Println("Output of 'ant install':", string(installOutput))

	// Run `ant --version` command
	versionCmd := exec.Command("ant", "--version")
	versionOutput, versionErr := versionCmd.CombinedOutput()
	if versionErr != nil {
		fmt.Println("Error running 'ant --version':", versionErr)
		return
	}
	fmt.Println("Output of 'ant --version':", string(versionOutput))
}
