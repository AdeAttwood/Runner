// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package commands

import (
	"errors"
	"strings"

	"github.com/AdeAttwood/Runner/pkg/config"
	"github.com/AdeAttwood/Runner/pkg/console"
	"github.com/AdeAttwood/Runner/pkg/runner"
)

var spinner *console.Spinner

func progress(event runner.RunnerEvent) {
	switch event.Type {
	case runner.Start:
		spinner = console.NewSpinner(event.Task.Title)
		spinner.Start()
	case runner.End:
		if event.Result.Error == nil {
			spinner.Success()
		} else {
			spinner.Error()
		}

		spinner = nil
	}
}

func Run(application Application, config config.Config) error {
	if len(application.Arguments) == 0 {
		return errors.New("Missing task to run")
	}

	tasks, err := config.BuildTaskList(application.Arguments[0])
	if err != nil {
		return err
	}

	result := runner.Run(tasks, progress)
	if result.Error != nil {
		console.WriteLine("")
		console.WriteLine(result.Error.Error())
		console.WriteLine(strings.TrimSpace(result.Output))
		return result.Error
	}

	return nil
}
