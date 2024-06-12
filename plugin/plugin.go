// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// Args provides plugin execution arguments.
type Args struct {
	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// Goals defines the ant command goals
	Goals string `envconfig:"PLUGIN_GOALS" required:"true"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// Split the Goals string into individual arguments
	goalArgs := strings.Fields(args.Goals)

	// Run the command with the specified goals
	cmd := exec.Command("ant", goalArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Error running 'ant %s': %s", args.Goals, err)
		logrus.Error(string(output))
		return fmt.Errorf("error running 'ant %s': %w", args.Goals, err)
	}
	logrus.Infof("Output of 'ant %s': %s", args.Goals, string(output))

	return nil
}
