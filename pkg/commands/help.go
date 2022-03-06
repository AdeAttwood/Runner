// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package commands

import (
	"fmt"

	"github.com/AdeAttwood/Runner/pkg/config"
)

func Help(application Application, config config.Config) error {
	fmt.Print(`
Usage: run [options] <command>

Options:

  -cwd <string> Overrides the current working directory where your 'runnner.jsonnet' file is
  -h            Display this message

Commands:

  help         Display this message
  run <string> Runs a command from your config
`)

	if len(config.Tasks) > 0 {
		List(application, config)
	}
	return nil
}
