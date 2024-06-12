// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os/exec"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// Run `ant -p` command to list available targets
	listTargetsCmd := exec.Command("ant", "-p")
	listTargetsOutput, listTargetsErr := listTargetsCmd.CombinedOutput()
	if listTargetsErr != nil {
		fmt.Println("Error running 'ant -p':", listTargetsErr)
		fmt.Println(string(listTargetsOutput))
		return fmt.Errorf("error running 'ant -p': %w", listTargetsErr)
	}
	fmt.Println("Output of 'ant -p':", string(listTargetsOutput))

	// Run `ant build` command to build the project
	buildCmd := exec.Command("ant", "build")
	buildOutput, buildErr := buildCmd.CombinedOutput()
	if buildErr != nil {
		fmt.Println("Error running 'ant build':", buildErr)
		fmt.Println(string(buildOutput))
		return fmt.Errorf("error running 'ant build': %w", buildErr)
	}
	fmt.Println("Output of 'ant build':", string(buildOutput))

	// Run `ant run` command to run the sample application
	runCmd := exec.Command("ant", "run")
	runOutput, runErr := runCmd.CombinedOutput()
	if runErr != nil {
		fmt.Println("Error running 'ant run':", runErr)
		fmt.Println(string(runOutput))
		return fmt.Errorf("error running 'ant run': %w", runErr)
	}
	fmt.Println("Output of 'ant run':", string(runOutput))

	// Run `ant --version` command
	versionCmd := exec.Command("ant", "-version")
	versionOutput, versionErr := versionCmd.CombinedOutput()
	if versionErr != nil {
		fmt.Println("Error running 'ant -version':", versionErr)
		fmt.Println(string(versionOutput))
		return fmt.Errorf("error running 'ant -version': %w", versionErr)
	}
	fmt.Println("Output of 'ant -version':", string(versionOutput))

	return nil
}
