package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Update package list
	// (Uncomment and modify if necessary)
	// updateCmd := exec.Command("apt-get", "update")
	// updateOutput, updateErr := updateCmd.CombinedOutput()
	// if updateErr != nil {
	// 	fmt.Println("Error running 'apt-get update':", updateErr)
	// 	fmt.Println(string(updateOutput))
	// 	return
	// }
	// fmt.Println("Output of 'apt-get update':", string(updateOutput))

	// Install Ant
	// installCmd := exec.Command("apt-get", "install", "-y", "ant")
	// installOutput, installErr := installCmd.CombinedOutput()
	// if installErr != nil {
	// 	fmt.Println("Error running 'apt-get install -y ant':", installErr)
	// 	fmt.Println(string(installOutput))
	// 	return
	// }
	// fmt.Println("Output of 'apt-get install -y ant':", string(installOutput))

	// Run `ant -p` command to list available targets
	listTargetsCmd := exec.Command("ant", "-p")
	listTargetsOutput, listTargetsErr := listTargetsCmd.CombinedOutput()
	if listTargetsErr != nil {
		fmt.Println("Error running 'ant -p':", listTargetsErr)
		fmt.Println(string(listTargetsOutput))
		return
	}
	fmt.Println("Output of 'ant -p':", string(listTargetsOutput))

	// Run `ant build` command to build the project
	buildCmd := exec.Command("ant", "build")
	buildOutput, buildErr := buildCmd.CombinedOutput()
	if buildErr != nil {
		fmt.Println("Error running 'ant build':", buildErr)
		fmt.Println(string(buildOutput))
		return
	}
	fmt.Println("Output of 'ant build':", string(buildOutput))

	// Run `ant run` command to run the sample application
	runCmd := exec.Command("ant", "run")
	runOutput, runErr := runCmd.CombinedOutput()
	if runErr != nil {
		fmt.Println("Error running 'ant run':", runErr)
		fmt.Println(string(runOutput))
		return
	}
	fmt.Println("Output of 'ant run':", string(runOutput))

	// Run `ant --version` command
	versionCmd := exec.Command("ant", "--version")
	versionOutput, versionErr := versionCmd.CombinedOutput()
	if versionErr != nil {
		fmt.Println("Error running 'ant --version':", versionErr)
		fmt.Println(string(versionOutput))
		return
	}
	fmt.Println("Output of 'ant --version':", string(versionOutput))
}
